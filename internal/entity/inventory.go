package entity

import (
	"time"

	"gorm.io/gorm"
)

// Inventory struct untuk tabel Inventory
type Inventory struct {
	ID          int            `gorm:"primaryKey;autoIncrement"`
	WarehouseID int            `gorm:"not null;foreignKey:WarehouseRefer;constraint:OnDelete:CASCADE"`
	ProductID   int            `gorm:"not null;foreignKey:ProductRefer;constraint:OnDelete:CASCADE"`
	Quantity    int            `gorm:"default:0"`
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
