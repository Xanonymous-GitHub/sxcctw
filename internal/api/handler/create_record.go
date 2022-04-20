package handler

import (
	"github.com/Xanonymous-GitHub/sxcctw/pkg/env"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type createRecordRequest struct {
	OriginUrl *string    `json:"originUrl,omitempty"`
	ExpireAt  *time.Time `json:"expireAt,omitempty"`
}

type createRecordResponse struct {
	ShortenedID string `json:"shortenedID,omitempty"`
}

type errResponse struct {
	Msg string `json:"msg,omitempty"`
}

const (
	HTTPS = "https"
	HTTP  = "http"
)

func (h *handler) HandleCreateRecord(ctx *gin.Context) {
	var req *createRecordRequest

	err := ctx.BindJSON(&req)
	if err != nil {
		h.logger.Errorln(err)
		msg := "Invalid Request data format"
		ctx.JSON(http.StatusBadRequest, &errResponse{Msg: msg})
		return
	}

	if req.OriginUrl == nil {
		msg := "originURL is not provided"
		h.logger.Warningln(msg)
		ctx.JSON(http.StatusBadRequest, &errResponse{Msg: msg})
		return
	}

	originUrl := strings.TrimSpace(*req.OriginUrl)
	if originUrl == "" {
		msg := "originURL is empty"
		h.logger.Warningln(msg)
		ctx.JSON(http.StatusBadRequest, &errResponse{Msg: msg})
		return
	}

	validOriginUrl, err := url.ParseRequestURI(originUrl)
	if err != nil {
		msg := "originUrl is invalid"
		h.logger.Warningln(msg)
		ctx.JSON(http.StatusBadRequest, &errResponse{Msg: msg})
		return
	}
	if validOriginUrl.Scheme != HTTPS && validOriginUrl.Scheme != HTTP {
		msg := "originUrl must be a http/https URL"
		h.logger.Warningln(msg)
		ctx.JSON(http.StatusBadRequest, &errResponse{Msg: msg})
		return
	}
	if validOriginUrl.Host == "" {
		msg := "originUrl must have host"
		h.logger.Warningln(msg)
		ctx.JSON(http.StatusBadRequest, &errResponse{Msg: msg})
		return
	}
	if validOriginUrl.Host == env.ShortenServerHost {
		msg := "originUrl has already be shortened"
		h.logger.Warningln(msg)
		ctx.JSON(http.StatusBadRequest, &errResponse{Msg: msg})
		return
	}

	var expireAt time.Time

	if req.ExpireAt == nil {
		h.logger.Infoln("ExpireAt is empty, will set default expiration to 2 weeks.")
		expireAt = time.Now().AddDate(0, 0, 14)
	} else {
		expireAt = *req.ExpireAt
	}

	shortenedID, err := h.urlService.CreateRecord(validOriginUrl.String(), expireAt)
	if err != nil {
		h.logger.Errorln(err)
		msg := "Error occurred when creating shortened URL, error code: 1"
		ctx.JSON(http.StatusInternalServerError, &errResponse{Msg: msg})
		return
	}
	if shortenedID == nil {
		h.logger.Errorln("shortenedID is nil")
		msg := "Error occurred when creating shortened URL, error code: 2"
		ctx.JSON(http.StatusInternalServerError, &errResponse{Msg: msg})
		return
	}

	resp := &createRecordResponse{ShortenedID: *shortenedID}
	ctx.JSON(http.StatusCreated, resp)
}
