package entity

import (
	"time"

	"gorm.io/gorm"
)

// Warehouse struct untuk tabel Warehouses
type Warehouse struct {
	ID        int            `gorm:"primaryKey;autoIncrement"`
	Name      string         `gorm:"size:100;not null"`
	Location  string         `gorm:"type:text;not null"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
