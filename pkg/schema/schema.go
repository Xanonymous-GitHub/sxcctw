package schema

import (
	"time"
)

type Record struct {
	ID          uint64 `gorm:"primaryKey"`
	ShortenedId string `gorm:"primaryKey"`
	OriginUrl   string
	CreatedAt   time.Time
	ExpiredAt   time.Time
}
