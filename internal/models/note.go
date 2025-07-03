package models

import (
	"time"

)

type Note struct {
	ID          string    `gorm:"primaryKey,size:6" json:"id"`
	Title       string    `gorm:"not null" json:"title"`
	Description string    `gorm:"not null" json:"description"`
	Url         *string    `gorm:"" json:"url"`
	Date        time.Time    `gorm:"not null" json:"date"`
	Time        time.Time    `gorm:"not null" json:"time"`
	Warn bool `gorm:"not null,default:true" json:"warn"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
