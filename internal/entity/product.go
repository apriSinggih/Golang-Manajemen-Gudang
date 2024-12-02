package entity

import (
	"time"

	"gorm.io/gorm"
)

// Product struct untuk tabel Products
type Product struct {
	ID          int            `gorm:"primaryKey;autoIncrement"`
	Name        string         `gorm:"size:100;not null"`
	Description string         `gorm:"type:text"`
	SKU         string         `gorm:"size:50;unique;not null"`
	SupplierID  int            `gorm:"foreignKey:SupplierRefer;constraint:OnDelete:SET NULL"`
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
