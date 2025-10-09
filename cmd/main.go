// @title           toge API
// @version         1.0
// @description     This is a toge server API documentation.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
package main

import (
	"log"
	"os"

	"git.lulumia.fun/root/toge-api/internal/app"
	"git.lulumia.fun/root/toge-api/internal/wire"
	"git.lulumia.fun/root/toge-api/pkg/config"
	"git.lulumia.fun/root/toge-api/pkg/logger"

	_ "git.lulumia.fun/root/toge-api/docs"

	"github.com/gin-gonic/gin"
)

func main() {
	// 获取环境变量，默认为 dev
	env := os.Getenv("ENV")
	if env == "" {
		env = "dev"
	}

	// 加载配置文件
	if err := config.LoadConfig(env); err != nil {
		log.Fatal("Failed to load config:", err)
	}

	// 验证日志配置
	if err := logger.ValidateLogConfig(); err != nil {
		log.Fatal("Invalid log config:", err)
	}

	// 设置 Gin 模式
	gin.SetMode(config.GlobalConfig.App.Mode)

	// 初始化日志
	app.InitializeLogger()

	// 使用 Wire 初始化应用
	appInstance, err := wire.InitializeApp()
	if err != nil {
		log.Fatal("Failed to initialize app:", err)
	}

	// 初始化数据库
	if err := app.InitializeDatabase(); err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	// 初始化 Redis
	if err := app.InitializeRedis(); err != nil {
		log.Fatal("Failed to initialize Redis:", err)
	}

	// 初始化时区
	if err := app.InitializeTimezone(); err != nil {
		log.Fatal("Failed to initialize timezone:", err)
	}

	// 设置路由
	appInstance.SetupRoutes()

	// 启动服务器
	if err := appInstance.Run(); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
