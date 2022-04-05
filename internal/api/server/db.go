package server

import (
	"context"
	"github.com/Xanonymous-GitHub/sxcctw/internal/repository"
	"github.com/Xanonymous-GitHub/sxcctw/internal/service"
	db_client "github.com/Xanonymous-GitHub/sxcctw/pkg/gorm"
	"github.com/Xanonymous-GitHub/sxcctw/pkg/proto/pb"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type recordServer struct {
	pb.UnimplementedRecordServiceServer
	logger        *logrus.Logger
	db            *gorm.DB
	recordService service.RecordService
}

func RegisterRecordServiceServer(server *grpc.Server, client db_client.DBClient, logger *logrus.Logger) {
	recordRepository := repository.CreateNewRecordRepositoryWith(client.DB(), logger)
	recordService := service.CreateRecordServiceWith(recordRepository, logger)

	pb.RegisterRecordServiceServer(server, &recordServer{
		db:            client.DB(),
		logger:        logger,
		recordService: recordService,
	})
}

func (r *recordServer) CreateRecord(_ context.Context, req *pb.CreateRecordRequest) (*pb.CreateRecordResponse, error) {
	return r.recordService.CreateRecord(req)
}

func (r *recordServer) GetOriginUrl(_ context.Context, req *pb.GetOriginUrlRequest) (*pb.GetOriginUrlResponse, error) {
	return r.recordService.GetOriginUrl(req)
}
