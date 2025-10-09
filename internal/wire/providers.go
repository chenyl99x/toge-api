package wire

import (
	"git.lulumia.fun/root/toge-api/internal/app"
	"git.lulumia.fun/root/toge-api/internal/handler"
	"git.lulumia.fun/root/toge-api/internal/repository"
	"git.lulumia.fun/root/toge-api/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// ProviderSet 是 wire 的提供者集合
var ProviderSet = wire.NewSet(
	// Repository 层
	repository.NewUserRepository,
	repository.NewPersonRepository,
	repository.NewNationRepository,
	repository.NewVersionRepository,
	repository.NewArtifactSetRepository,
	repository.NewArtifactRepository,

	// Service 层
	service.NewUserService,
	service.NewPersonService,
	service.NewNationService,
	service.NewVersionService,
	service.NewArtifactSetService,
	service.NewArtifactService,
	// Handler 层
	handler.NewAuthHandler,
	handler.NewHealthHandler,
	handler.NewTimezoneHandler,
	handler.NewUserHandler,
	handler.NewPersonHandler,
	handler.NewNationHandler,
	handler.NewVersionHandler,
	handler.NewArtifactSetHandler,
	handler.NewArtifactHandler,

	// 提供 gin 引擎
	ProvideGinEngine,

	// 提供应用实例
	app.NewApp,
	app.InitializeDatabase,
)

// ProvideGinEngine 提供 gin 引擎
func ProvideGinEngine() *gin.Engine {
	// 使用 gin.New() 而不是 gin.Default() 来避免默认的日志中间件
	engine := gin.New()

	// 只添加必要的中间件，不包含默认的日志中间件
	// 我们使用自定义的日志中间件来替代
	return engine
}
