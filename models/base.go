package models

import "time"

type BaseModel struct {
	ID        uint64    `json:"id",gorm:"primaryKey,autoIncrement"`
	CreatedAt time.Time `json:"createdAt"`
}
