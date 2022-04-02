package schema

import (
	"time"
)

type Record struct {
	ID          uint64 `gorm:"primarykey"`
	ShortenedId string `gorm:"primarykey"`
	OriginUrl   string
	CreatedAt   time.Time
	ExpiredAt   time.Time
}
