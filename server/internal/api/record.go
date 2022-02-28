package api

import (
	"context"
	"github.com/Xanonymous-GitHub/sxcctw/db/pkg/proto/pb"
	"github.com/Xanonymous-GitHub/sxcctw/server/internal/grpcsvc"
	"github.com/Xanonymous-GitHub/sxcctw/server/internal/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type RecordApiHandler interface {
	Create(*gin.Context)
	Get(*gin.Context)
}

type recordHandler struct{}

func NewRecordHandler() RecordApiHandler {
	return &recordHandler{}
}

func (r *recordHandler) Create(ctx *gin.Context) {
	var createRecordRequest *pb.CreateRecordRequest
	err := ctx.BindJSON(&createRecordRequest)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c := context.Background()
	createRecordResponse, err := service.CreateRecordService(c, grpcsvc.RecordSvcClient, createRecordRequest)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusCreated, createRecordResponse)
}

func (r *recordHandler) Get(ctx *gin.Context) {
	var getRecordRequest *pb.GetOriginUrlRequest
	err := ctx.BindJSON(&getRecordRequest)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c := context.Background()
	getRecordResponse, err := service.GetRecordService(c, grpcsvc.RecordSvcClient, getRecordRequest)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, getRecordResponse)
}
