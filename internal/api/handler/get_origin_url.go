package handler

import (
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

	originUrl, err := h.urlService.GetOriginUrl(id)
	if err != nil {
		h.logger.Errorln(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if originUrl == nil {
		h.logger.Warningln("origin url is empty")
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	resp := &getOriginUrlResponse{OriginUrl: *originUrl}
	ctx.JSON(http.StatusOK, resp)
}
