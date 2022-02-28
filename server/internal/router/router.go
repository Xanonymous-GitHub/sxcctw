package router

import (
	"github.com/Xanonymous-GitHub/sxcctw/db/pkg/vp"
	"github.com/Xanonymous-GitHub/sxcctw/server/internal/api"
	"github.com/gin-gonic/gin"
)

func routerMode() string {
	isDebugMode := vp.Cvp.GetBool("debug")

	if isDebugMode {
		return gin.DebugMode
	}

	return gin.ReleaseMode
}

func NewRouter() *gin.Engine {
	mode := routerMode()
	gin.SetMode(mode)

	gin.ForceConsoleColor()

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	const (
		apiRootPath   = "/api"
		recordApiPath = apiRootPath + "/record"
	)

	// Application APIs
	recordApiHandlers := api.NewRecordHandler()
	recordApiGroup := r.Group(recordApiPath)
	{
		recordApiGroup.POST("", recordApiHandlers.Create)
		recordApiGroup.GET("", recordApiHandlers.Get)
	}

	return r
}
