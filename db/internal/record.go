package internal

import (
	"github.com/Xanonymous-GitHub/sxcctw/db/pkg/proto/pb"
	"github.com/Xanonymous-GitHub/sxcctw/db/pkg/schema"
	"gorm.io/gorm"
	"math/rand"
	"time"
)

func isIdUsed(db *gorm.DB, testId uint64) (bool, error) {
	var isExists bool
	err := db.Model(&schema.Record{}).
		Select("count(*) > 0").
		Where("id = ?", testId).
		Find(&isExists).
		Error
	if err != nil {
		return false, err
	}
	return isExists, nil
}

func recordStatus(isExpired bool) pb.RecordStatus {
	if isExpired {
		return pb.RecordStatus_EXPIRED
	}
	return pb.RecordStatus_NORMAL
}

func SaveRecord(db *gorm.DB, newRecordRequest *pb.CreateRecordRequest) (*pb.CreateRecordResponse, error) {
	// Generate a unique id for SQL to record each short link,
	// since gorm's id auto-generation process will be executed after Create operation called,
	// and we encode this unique ID to be the shortened URL id, so we can not use gorm's auto ID.
	// The reason why the shortened URLs are always has a same length is their ID are all have same length.
	id := rand.Uint64()
	var err error
	for used := true; used; used, err = isIdUsed(db, id) {
		if err != nil {
			return nil, err
		}
		id = rand.Uint64()
	}

	shortenedId := Encode(id)

	newRecord := &schema.Record{
		ID:          id,
		ShortenedId: shortenedId,
		OriginUrl:   newRecordRequest.OriginUrl,
		ExpiredAt:   newRecordRequest.ExpireAt.AsTime(),
	}

	result := db.Create(newRecord)

	err = result.Error
	if err != nil {
		return nil, err
	}

	return &pb.CreateRecordResponse{ShortenedId: shortenedId}, nil
}

func LoadRecord(db *gorm.DB, loadRecordRequest *pb.GetOriginUrlRequest) (*pb.GetOriginUrlResponse, error) {
	decodedId, err := Decode(loadRecordRequest.ShortenedId)
	if err != nil {
		return nil, err
	}

	var storedRecord *schema.Record
	//var isExist bool
	err = db.Model(&schema.Record{}).
		Where("id = ?", decodedId).
		Find(&storedRecord).
		Error
	if err != nil {
		return nil, err
	}

	now := time.Now()
	isExpired := now.After(storedRecord.ExpiredAt)
	isExist := storedRecord.ID != 0 && storedRecord.ID == decodedId

	var status pb.RecordStatus
	if !isExist {
		status = pb.RecordStatus_NOTFOUND
	} else {
		status = recordStatus(isExpired)
	}

	return &pb.GetOriginUrlResponse{
		OriginUrl: storedRecord.OriginUrl,
		Status:    status,
	}, nil
}
