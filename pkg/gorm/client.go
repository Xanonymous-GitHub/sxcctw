package gorm

import (
	"errors"
	"github.com/Xanonymous-GitHub/sxcctw/pkg/env"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"reflect"
)

type DBClient interface {
	DB() *gorm.DB
	PrepareTables(schema []any) error
}

type dbClient struct {
	db     *gorm.DB
	logger *logrus.Logger
}

func CreateNewDBClientWith(db *gorm.DB, logger *logrus.Logger) DBClient {
	return &dbClient{
		db:     db,
		logger: logger,
	}
}

func (d *dbClient) DB() *gorm.DB {
	if d != nil {
		return d.db
	}

	return nil
}

func (d *dbClient) createTablesProd(schemas []any) error {
	if d == nil {
		msg := "dbClient is nil"
		d.logger.Errorln(msg)
		return errors.New(msg)
	}

	// TODO(TU): use go generic type to replace the `[]any` type.
	for _, model := range schemas {
		if !d.db.Migrator().HasTable(model) {
			d.logger.Warnf("The table %v not exist, creating...", reflect.TypeOf(model))
			if err := d.db.Migrator().CreateTable(model); err != nil {
				d.logger.Errorln(err)
				return err
			}
		}
		d.logger.Infof("The table %v existed, skip create", reflect.TypeOf(model))
	}

	return nil
}

func (d *dbClient) PrepareTables(schemas []any) error {
	if d == nil {
		msg := "dbClient is nil"
		d.logger.Errorln(msg)
		return errors.New(msg)
	}

	// TODO(TU): use go generic type to replace the `[]any` type.
	if env.IsDebugMode {
		if err := d.db.Debug().AutoMigrate(schemas...); err != nil {
			d.logger.Errorln(err)
			return err
		}
	} else {
		if err := d.createTablesProd(schemas); err != nil {
			d.logger.Errorln(err)
			return err
		}
	}

	return nil
}
