package server

import (
	"github.com/Xanonymous-GitHub/sxcctw/internal/repository"
	"github.com/Xanonymous-GitHub/sxcctw/internal/service"
	"github.com/Xanonymous-GitHub/sxcctw/pkg/env"
	"github.com/Xanonymous-GitHub/sxcctw/pkg/proto/pb"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func RegisterRestApiHandlers(recordSvcClient *pb.RecordServiceClient, router *gin.Engine, logger *logrus.Logger) error {
	urlRepository := repository.CreateUrlRepositoryWith(*recordSvcClient, logger)
	urlService := service.CreateUrlServiceWith(urlRepository, logger)

	handler := CreateNewRestApiHandlerWith(urlService, logger)

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
	urlService service.UrlService
	logger     *logrus.Logger
}

func CreateNewRestApiHandlerWith(urlService service.UrlService, logger *logrus.Logger) RestApiHandler {
	return &handler{
		urlService: urlService,
		logger:     logger,
	}
}

type createRecordRequest struct {
	OriginUrl string    `json:"originUrl,omitempty"`
	ExpireAt  time.Time `json:"expireAt"`
}

type createRecordResponse struct {
	ShortenedID string `json:"shortenedID,omitempty"`
}

func (h *handler) HandleGetOriginUrl(ctx *gin.Context) {

}

func (h *handler) HandleCreateRecord(ctx *gin.Context) {
	var req *createRecordRequest

	err := ctx.BindJSON(&req)
	if err != nil {
		h.logger.Errorln(err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	shortenedID, err := h.urlService.CreateRecord(req.OriginUrl, req.ExpireAt)
	if err != nil {
		h.logger.Errorln(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if shortenedID == nil {
		h.logger.Errorln("shortenedID is nil")
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	resp := &createRecordResponse{ShortenedID: *shortenedID}
	ctx.JSON(http.StatusCreated, resp)
}
