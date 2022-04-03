package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

type createRecordRequest struct {
	OriginUrl string    `json:"originUrl,omitempty"`
	ExpireAt  time.Time `json:"expireAt"`
}

type createRecordResponse struct {
	ShortenedID string `json:"shortenedID,omitempty"`
}

func (h *handler) HandleCreateRecord(ctx *gin.Context) {
	var req *createRecordRequest

	err := ctx.BindJSON(&req)
	if err != nil {
		h.logger.Errorln(err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	shortenedID, err := h.urlService.CreateRecord(strings.TrimSpace(req.OriginUrl), req.ExpireAt)
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
