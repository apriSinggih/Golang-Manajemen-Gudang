package entity

import (
	"time"
	"gorm.io/gorm"
)

// User struct untuk tabel Users
type User struct {
	ID        int            `gorm:"primaryKey;autoIncrement"`
	Name      string         `gorm:"size:100;not null"`
	Email     string         `gorm:"size:100;unique;not null"`
	Password  string         `gorm:"size:255;not null"`
	Role      string         `gorm:"type:enum('admin','staff');not null"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
