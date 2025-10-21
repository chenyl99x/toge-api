package migrate

import (
	"fmt"
	"time"

	"github.com/chenyl99x/toge-api/internal/model"
	"github.com/chenyl99x/toge-api/pkg/database"
	"github.com/chenyl99x/toge-api/pkg/logger"
)

// MigrationFunc 迁移函数类型
type MigrationFunc func() error

// Migration 迁移结构
type Migration struct {
	Version     string
	Description string
	Up          MigrationFunc
	Down        MigrationFunc
}

// 迁移列表
var migrations = []Migration{
	{
		Version:     "010",
		Description: "Create initial tables",
		Up: func() error {
			return database.DB.AutoMigrate(
				&model.User{},
			)
		},
		Down: func() error {
			return database.DB.Migrator().DropTable(
				&model.User{},
			)
		},
	},
}

// RunMigrations 执行所有未应用的迁移
func RunMigrations() error {
	// 首先确保 migrations 表存在
	if err := database.DB.AutoMigrate(&model.Migration{}); err != nil {
		return fmt.Errorf("failed to create migrations table: %v", err)
	}

	// 获取已应用的迁移
	appliedMigrations, err := getAppliedMigrations()
	if err != nil {
		return fmt.Errorf("failed to get applied migrations: %v", err)
	}

	// 执行未应用的迁移
	for _, migration := range migrations {
		if isApplied(appliedMigrations, migration.Version) {
			logger.Info("Migration already applied", "version", migration.Version)
			continue
		}

		logger.Info("Applying migration", "version", migration.Version, "description", migration.Description)

		// 开始事务
		tx := database.DB.Begin()
		if tx.Error != nil {
			return fmt.Errorf("failed to start transaction: %v", tx.Error)
		}

		// 执行迁移
		if err := migration.Up(); err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to apply migration %s: %v", migration.Version, err)
		}

		// 记录迁移
		migrationRecord := &model.Migration{
			Version:     migration.Version,
			Description: migration.Description,
			AppliedAt:   time.Now(),
		}

		if err := tx.Create(migrationRecord).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to record migration %s: %v", migration.Version, err)
		}

		// 提交事务
		if err := tx.Commit().Error; err != nil {
			return fmt.Errorf("failed to commit migration %s: %v", migration.Version, err)
		}

		logger.Info("Migration applied successfully", "version", migration.Version)
	}

	return nil
}

// RollbackMigration 回滚指定版本的迁移
func RollbackMigration(version string) error {
	// 查找迁移
	var migration *Migration
	for _, m := range migrations {
		if m.Version == version {
			migration = &m
			break
		}
	}

	if migration == nil {
		return fmt.Errorf("migration %s not found", version)
	}

	// 检查是否已应用
	appliedMigrations, err := getAppliedMigrations()
	if err != nil {
		return fmt.Errorf("failed to get applied migrations: %v", err)
	}

	if !isApplied(appliedMigrations, version) {
		return fmt.Errorf("migration %s is not applied", version)
	}

	logger.Info("Rolling back migration", "version", version)

	// 开始事务
	tx := database.DB.Begin()
	if tx.Error != nil {
		return fmt.Errorf("failed to start transaction: %v", tx.Error)
	}

	// 执行回滚
	if err := migration.Down(); err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to rollback migration %s: %v", version, err)
	}

	// 删除迁移记录
	if err := tx.Where("version = ?", version).Delete(&model.Migration{}).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to delete migration record %s: %v", version, err)
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("failed to commit rollback %s: %v", version, err)
	}

	logger.Info("Migration rolled back successfully", "version", version)
	return nil
}

// GetMigrationStatus 获取迁移状态
func GetMigrationStatus() ([]model.Migration, error) {
	var migrations []model.Migration
	if err := database.DB.Order("applied_at").Find(&migrations).Error; err != nil {
		return nil, err
	}
	return migrations, nil
}

// getAppliedMigrations 获取已应用的迁移版本
func getAppliedMigrations() ([]string, error) {
	var versions []string
	if err := database.DB.Model(&model.Migration{}).Pluck("version", &versions).Error; err != nil {
		return nil, err
	}
	return versions, nil
}

// isApplied 检查迁移是否已应用
func isApplied(appliedMigrations []string, version string) bool {
	for _, applied := range appliedMigrations {
		if applied == version {
			return true
		}
	}
	return false
}

// DropTables 删除所有表（谨慎使用）
func DropTables() error {
	return database.DB.Migrator().DropTable(
		&model.User{},
		&model.Migration{},
	)
}

// ResetDatabase 重置数据库（删除并重新创建表）
func ResetDatabase() error {
	if err := DropTables(); err != nil {
		return err
	}
	return RunMigrations()
}
