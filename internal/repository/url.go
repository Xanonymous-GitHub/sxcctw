package repository

import (
	"context"
	"github.com/Xanonymous-GitHub/sxcctw/pkg/proto/pb"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type UrlRepository interface {
	GetOriginUrlWith(ctx context.Context, shortenedId string) (*pb.GetOriginUrlResponse, error)
	CreateRecordWith(ctx context.Context, originUrl string, expireAt time.Time) (*pb.CreateRecordResponse, error)
}

type urlRepository struct {
	logger              *logrus.Logger
	recordServiceClient pb.RecordServiceClient
}

func CreateUrlRepositoryWith(
	recordServiceClient pb.RecordServiceClient,
	logger *logrus.Logger,
) UrlRepository {
	return &urlRepository{
		recordServiceClient: recordServiceClient,
		logger:              logger,
	}
}

func (r *urlRepository) GetOriginUrlWith(ctx context.Context, shortenedId string) (*pb.GetOriginUrlResponse, error) {
	req := &pb.GetOriginUrlRequest{ShortenedId: shortenedId}
	resp, err := r.recordServiceClient.GetOriginUrl(ctx, req)
	if err != nil {
		r.logger.Errorln(err)
		return nil, err
	}

	return resp, nil
}

func (r *urlRepository) CreateRecordWith(ctx context.Context, originUrl string, expireAt time.Time) (*pb.CreateRecordResponse, error) {
	req := &pb.CreateRecordRequest{
		OriginUrl: originUrl,
		ExpireAt:  timestamppb.New(expireAt),
	}
	resp, err := r.recordServiceClient.CreateRecord(ctx, req)
	if err != nil {
		r.logger.Errorln(err)
		return nil, err
	}

	return resp, nil
}
