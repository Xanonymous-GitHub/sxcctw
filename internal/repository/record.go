package repository

import (
	"github.com/Xanonymous-GitHub/sxcctw/pkg/schema"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type RecordRepository interface {
	IsIdUsed(id uint64) bool
	SaveRecord(record schema.Record) error
	LoadRecord(id uint64) (*schema.Record, error)
}

type repository struct {
	db     *gorm.DB
	logger *logrus.Logger
}

func CreateNewRecordRepositoryWith(
	db *gorm.DB,
	logger *logrus.Logger,
) RecordRepository {
	return &repository{
		db:     db,
		logger: logger,
	}
}

func (r *repository) IsIdUsed(id uint64) bool {
	var isExists bool

	if err := r.db.Model(&schema.Record{}).
		Select("count(*) > 0").
		Where("id = ?", id).
		Find(&isExists).
		Error; err != nil {
		r.logger.Errorln(err)
		return false
	}

	return isExists
}

func (r *repository) SaveRecord(record schema.Record) error {
	result := r.db.Create(record)

	err := result.Error
	if err != nil {
		r.logger.Errorln(err)
		return err
	}

	return nil
}

func (r *repository) LoadRecord(id uint64) (*schema.Record, error) {
	var storedRecord *schema.Record

	if err := r.db.Model(&schema.Record{}).
		Where("id = ?", id).
		Find(&storedRecord).
		Error; err != nil {
		return nil, err
	}

	return storedRecord, nil
}
