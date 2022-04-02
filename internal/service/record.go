package service

import (
	"github.com/Xanonymous-GitHub/sxcctw/internal/repository"
	"github.com/Xanonymous-GitHub/sxcctw/pkg/base62"
	"github.com/Xanonymous-GitHub/sxcctw/pkg/proto/pb"
	"github.com/Xanonymous-GitHub/sxcctw/pkg/schema"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"math/rand"
	"time"
)

type RecordService interface {
	CreateRecord(*pb.CreateRecordRequest) (*pb.CreateRecordResponse, error)
	GetOriginUrl(*pb.GetOriginUrlRequest) (*pb.GetOriginUrlResponse, error)
}

type service struct {
	recordRepository repository.RecordRepository
	logger           *logrus.Logger
}

func CreateRecordServiceWith(
	recordRepository repository.RecordRepository,
	logger *logrus.Logger,
) RecordService {
	return &service{
		recordRepository: recordRepository,
		logger:           logger,
	}
}

func isExpired(u time.Time) bool {
	return time.Now().After(u)
}

func recordStatusOf(isExpired bool) pb.RecordStatus {
	if isExpired {
		return pb.RecordStatus_EXPIRED
	}

	return pb.RecordStatus_NORMAL
}

func (s *service) CreateRecord(req *pb.CreateRecordRequest) (*pb.CreateRecordResponse, error) {
	// Generate a unique id for SQL to record each short link,
	// since gorm's id auto-generation process will be executed after Create operation called,
	// and we encode this unique ID to be the shortened URL id, so we can not use gorm's auto ID.
	// The reason why the shortened URLs are always has a same length is their ID are all have same length.
	newID := rand.Uint64()
	for used := true; used; used = s.recordRepository.IsIdUsed(newID) {
		newID = rand.Uint64()
	}

	shortenedID := base62.Encode(newID)

	newRecord := &schema.Record{
		ID:          newID,
		ShortenedId: shortenedID,
		OriginUrl:   req.OriginUrl,
		ExpiredAt:   req.ExpireAt.AsTime(),
	}

	err := s.recordRepository.SaveRecord(*newRecord)
	if err != nil {
		s.logger.Errorln(err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.CreateRecordResponse{ShortenedId: shortenedID}, nil
}

func (s *service) GetOriginUrl(req *pb.GetOriginUrlRequest) (*pb.GetOriginUrlResponse, error) {
	decodedId, err := base62.Decode(req.ShortenedId)
	if err != nil {
		s.logger.Errorln(err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	storedRecord, err := s.recordRepository.LoadRecord(decodedId)
	if err != nil {
		s.logger.Errorln(err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	expired := isExpired(storedRecord.ExpiredAt)
	isExist := storedRecord.ID != 0 && storedRecord.ID == decodedId

	var recordStatus pb.RecordStatus
	if !isExist {
		recordStatus = pb.RecordStatus_NOTFOUND
	} else {
		recordStatus = recordStatusOf(expired)
	}

	return &pb.GetOriginUrlResponse{
		OriginUrl: storedRecord.OriginUrl,
		Status:    recordStatus,
	}, nil
}
