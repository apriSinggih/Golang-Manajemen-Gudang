package entity

import (
	"time"

	"gorm.io/gorm"
)

// Customer struct untuk tabel Customers (Opsional)
type Customer struct {
	ID        int            `gorm:"primaryKey;autoIncrement"`
	Name      string         `gorm:"size:100;not null"`
	Contact   string         `gorm:"size:50"`
	Address   string         `gorm:"type:text"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
