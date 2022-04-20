package handler

import (
	"github.com/Xanonymous-GitHub/sxcctw/pkg/proto/pb"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type getOriginUrlResponse struct {
	OriginUrl string `json:"originUrl,omitempty"`
}

func (h *handler) HandleGetOriginUrl(ctx *gin.Context) {
	id := strings.TrimSpace(ctx.Query("id"))
	if id == "" {
		h.logger.Warningln("id is empty")
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	originUrl, status, err := h.urlService.GetOriginUrl(id)
	if err != nil {
		h.logger.Errorln(err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if originUrl == nil {
		h.logger.Warningln("origin url is empty")
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if status == nil {
		h.logger.Warningln("status is nil")
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if *status == pb.RecordStatus_EXPIRED {
		ctx.AbortWithStatus(http.StatusGone)
		return
	}

	if *status == pb.RecordStatus_NOTFOUND {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	resp := &getOriginUrlResponse{OriginUrl: *originUrl}
	ctx.JSON(http.StatusOK, resp)
}
