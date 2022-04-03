package service

import (
	"context"
	"errors"
	"github.com/Xanonymous-GitHub/sxcctw/internal/repository"
	"github.com/Xanonymous-GitHub/sxcctw/pkg/proto/pb"
	"github.com/sirupsen/logrus"
	"time"
)

type UrlService interface {
	GetOriginUrl(shortenedId string) (*string, error)
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

func (s *urlService) GetOriginUrl(shortenedId string) (*string, error) {
	result, err := s.urlRepository.GetOriginUrlWith(context.Background(), shortenedId)
	if err != nil {
		s.logger.Errorln(err)
		return nil, err
	}

	if result.Status == pb.RecordStatus_EXPIRED {
		msg := "record expired"
		s.logger.Warningln(msg)
		return nil, errors.New(msg)
	}

	if result.Status == pb.RecordStatus_NOTFOUND {
		msg := "record not found"
		s.logger.Warningln(msg)
		return nil, errors.New(msg)
	}

	return &result.OriginUrl, nil
}

func (s *urlService) CreateRecord(originUrl string, expireAt time.Time) (*string, error) {
	result, err := s.urlRepository.CreateRecordWith(context.Background(), originUrl, expireAt)
	if err != nil {
		s.logger.Errorln(err)
		return nil, err
	}

	return &result.ShortenedId, nil
}
