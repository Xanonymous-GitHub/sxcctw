package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (h *handler) HandleRedirect(ctx *gin.Context) {
	id := strings.TrimSpace(ctx.Param("id"))
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

	ctx.Redirect(http.StatusFound, *originUrl)
}
