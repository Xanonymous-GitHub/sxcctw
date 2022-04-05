package service

import (
	"context"
	"github.com/Xanonymous-GitHub/sxcctw/internal/repository"
	"github.com/Xanonymous-GitHub/sxcctw/pkg/proto/pb"
	"github.com/sirupsen/logrus"
	"time"
)

type UrlService interface {
	GetOriginUrl(shortenedId string) (*string, *pb.RecordStatus, error)
	CreateRecord(originUrl string, expireAt time.Time) (*string, error)
}

type urlService struct {
	urlRepository repository.UrlRepository
	logger        *logrus.Logger
}

func CreateUrlServiceWith(
	urlRepository repository.UrlRepository,
	logger *logrus.Logger,
) UrlService {
	return &urlService{
		urlRepository: urlRepository,
		logger:        logger,
	}
}

func (s *urlService) GetOriginUrl(shortenedId string) (*string, *pb.RecordStatus, error) {
	result, err := s.urlRepository.GetOriginUrlWith(context.Background(), shortenedId)
	if err != nil {
		s.logger.Errorln(err)
		return nil, nil, err
	}

	return &result.OriginUrl, &result.Status, nil
}

func (s *urlService) CreateRecord(originUrl string, expireAt time.Time) (*string, error) {
	result, err := s.urlRepository.CreateRecordWith(context.Background(), originUrl, expireAt)
	if err != nil {
		s.logger.Errorln(err)
		return nil, err
	}

	return &result.ShortenedId, nil
}
