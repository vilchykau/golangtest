package api

import (
	"github.com/gin-gonic/gin"
	"github.com/vilchykau/golangtest/internal/api/subs"
)

func initSubs(engine *gin.Engine) {
	g := engine.Group("/api/v1/subs/")
	subs.InitGroup(g)
}

func InitGroups(engine *gin.Engine) {
	initSubs(engine)
}
