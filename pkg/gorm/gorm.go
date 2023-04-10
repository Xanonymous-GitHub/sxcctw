package gorm

import (
	"context"
	"fmt"
	"github.com/Xanonymous-GitHub/sxcctw/pkg/env"
	gmt "github.com/wei840222/gorm-otel"
	"go.uber.org/fx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/url"
)

// NewGORMMySQLClient return new *gorm.DB with mysql configs with fx.Lifecycle
func NewGORMMySQLClient(lc fx.Lifecycle) (*gorm.DB, error) {
	db, err := NewGORMMySQLClientWithoutLC()
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	if err := db.Use(gmt.New(
		gmt.WithLogResult(env.IsDebugMode),
		gmt.WithSqlParameters(env.IsDebugMode))); err != nil {
		return nil, err
	}

	lc.Append(fx.Hook{
		OnStop: func(context.Context) error {
			return sqlDB.Close()
		},
	})
	return db, nil
}

// NewGORMMySQLClientWithoutLC return new *gorm.DB with mysql configs
func NewGORMMySQLClientWithoutLC() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=%s",
		env.DBUserName,
		env.DBPassword,
		env.DBHost,
		env.DBPort,
		env.DBName,
		url.QueryEscape("'+8:00'"),
	)

	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
