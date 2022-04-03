package server

import (
	"github.com/Xanonymous-GitHub/sxcctw/pkg/env"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func RegisterRestApiHandlers(router *gin.Engine, logger *logrus.Logger) error {
	handler := CreateNewRestApiHandlerWith(logger)

	routerGroup := router.Group(env.ApiRootPath)
	{
		routerGroup.GET("/url", handler.HandleGetOriginUrl)
		routerGroup.POST("/url", handler.HandleGetOriginUrl)
	}

	return nil
}

type RestApiHandler interface {
	HandleGetOriginUrl(*gin.Context)
	HandleCreateRecord(*gin.Context)
}

type handler struct {
	logger *logrus.Logger
}

func CreateNewRestApiHandlerWith(logger *logrus.Logger) RestApiHandler {
	return &handler{
		logger: logger,
	}
}

func (h *handler) HandleGetOriginUrl(ctx *gin.Context) {

}

func (h handler) HandleCreateRecord(ctx *gin.Context) {

}
