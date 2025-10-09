package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"git.lulumia.fun/root/toge-api/pkg/config"
	"git.lulumia.fun/root/toge-api/pkg/database"
	"git.lulumia.fun/root/toge-api/pkg/logger"
	"git.lulumia.fun/root/toge-api/pkg/migrate"
)

func main() {
	var (
		env     = flag.String("env", "dev", "Environment (dev, test, production)")
		action  = flag.String("action", "up", "Action: up, down, status, reset")
		version = flag.String("version", "", "Migration version (for down action)")
	)
	flag.Parse()

	// 加载配置
	if err := config.LoadConfig(*env); err != nil {
		log.Fatal("Failed to load config:", err)
	}

	// 验证日志配置
	if err := logger.ValidateLogConfig(); err != nil {
		log.Fatal("Invalid log config:", err)
	}

	// 初始化日志
	logger.InitLogger()

	// 初始化数据库连接
	if err := database.InitDB(); err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	// 执行迁移操作
	switch *action {
	case "up":
		if err := migrate.RunMigrations(); err != nil {
			log.Fatal("Failed to run migrations:", err)
		}
		fmt.Println("Migrations completed successfully")

	case "down":
		if *version == "" {
			log.Fatal("Version is required for down action")
		}
		if err := migrate.RollbackMigration(*version); err != nil {
			log.Fatal("Failed to rollback migration:", err)
		}
		fmt.Printf("Migration %s rolled back successfully\n", *version)

	case "status":
		migrations, err := migrate.GetMigrationStatus()
		if err != nil {
			log.Fatal("Failed to get migration status:", err)
		}

		fmt.Println("Migration Status:")
		fmt.Println("=================")
		if len(migrations) == 0 {
			fmt.Println("No migrations applied")
		} else {
			for _, m := range migrations {
				fmt.Printf("Version: %s, Description: %s, Applied: %s\n",
					m.Version, m.Description, m.AppliedAt.Format("2006-01-02 15:04:05"))
			}
		}

	case "reset":
		fmt.Print("This will drop all tables and recreate them. Are you sure? (y/N): ")
		var confirm string
		fmt.Scanln(&confirm)
		if confirm != "y" && confirm != "Y" {
			fmt.Println("Operation cancelled")
			os.Exit(0)
		}

		if err := migrate.ResetDatabase(); err != nil {
			log.Fatal("Failed to reset database:", err)
		}
		fmt.Println("Database reset successfully")

	default:
		fmt.Println("Usage:")
		fmt.Println("  go run cmd/migrate/main.go -action=up                    # Apply all pending migrations")
		fmt.Println("  go run cmd/migrate/main.go -action=down -version=001     # Rollback specific migration")
		fmt.Println("  go run cmd/migrate/main.go -action=status                # Show migration status")
		fmt.Println("  go run cmd/migrate/main.go -action=reset                 # Reset database (DANGEROUS)")
		fmt.Println("")
		fmt.Println("Options:")
		fmt.Println("  -env string     Environment (dev, test, production) (default: dev)")
		fmt.Println("  -action string  Action: up, down, status, reset (default: up)")
		fmt.Println("  -version string Migration version (required for down action)")
		os.Exit(1)
	}
}
