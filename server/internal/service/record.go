package service

import (
	"context"
	"github.com/Xanonymous-GitHub/sxcctw/db/pkg/proto/pb"
)

func CreateRecordService(ctx context.Context, client pb.RecordServiceClient, request *pb.CreateRecordRequest) (*pb.CreateRecordResponse, error) {
	created, err := client.CreateRecord(ctx, request)
	if err != nil {
		return nil, err
	}

	return created, nil
}

func GetRecordService(ctx context.Context, client pb.RecordServiceClient, request *pb.GetOriginUrlRequest) (*pb.GetOriginUrlResponse, error) {
	queried, err := client.GetOriginUrl(ctx, request)
	if err != nil {
		return nil, err
	}

	return queried, nil
}
