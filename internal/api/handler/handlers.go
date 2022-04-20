package handler

import (
	"github.com/Xanonymous-GitHub/sxcctw/internal/repository"
	"github.com/Xanonymous-GitHub/sxcctw/internal/service"
	"github.com/Xanonymous-GitHub/sxcctw/pkg/env"
	"github.com/Xanonymous-GitHub/sxcctw/pkg/proto/pb"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func RegisterRestApiHandlers(recordSvcClient pb.RecordServiceClient, router *gin.Engine, logger *logrus.Logger) error {
	urlRepository := repository.CreateUrlRepositoryWith(recordSvcClient, logger)
	urlService := service.CreateUrlServiceWith(urlRepository, logger)

	handler := CreateNewRestApiHandlerWith(urlService, logger)

	routerGroup := router.Group(env.ApiRootPath)
	{
		routerGroup.GET("/url", handler.HandleGetOriginUrl)
		routerGroup.POST("/url", handler.HandleCreateRecord)
	}

	router.Any("/s/:id", handler.HandleRedirect)

	return nil
}

type RestApiHandler interface {
	HandleGetOriginUrl(*gin.Context)
	HandleCreateRecord(*gin.Context)
	HandleRedirect(*gin.Context)
}

type handler struct {
	urlService service.UrlService
	logger     *logrus.Logger
}

func CreateNewRestApiHandlerWith(urlService service.UrlService, logger *logrus.Logger) RestApiHandler {
	return &handler{
		urlService: urlService,
		logger:     logger,
	}
}
