package app

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/vilchykau/golangtest/api"
	_ "github.com/vilchykau/golangtest/docs"
	"github.com/vilchykau/golangtest/middleware"
)

type OnShutdownExec func()

type App struct {
	onShutdownE []OnShutdownExec
}

func (a *App) Start() {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	// driver :=
	// 	drivers.NewPgDriver(drivers.PgConfig{DnsString: "User ID=habrpguser;Password=pgpwd4habr;Host=127.0.0.1;Port=5432;Database=habrdb;Pooling=true;Min Pool Size=0;Max Pool Size=100;Connection Lifetime=0;"})
	// driver.Connect(ctx)

	r.Use(middleware.CorsMiddleware())
	r.Use(middleware.Db())
	api.InitGroups(r)
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Run(":8001")
}

func (a *App) AddOnShutdown(shut OnShutdownExec) {
	a.onShutdownE = append(a.onShutdownE, shut)
}

func (a *App) ShutDown() {
	for _, on := range a.onShutdownE {
		on()
	}
}
