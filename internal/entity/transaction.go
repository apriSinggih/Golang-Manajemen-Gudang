package entity

import (
	"time"

	"gorm.io/gorm"
)

// Transaction struct untuk tabel Transactions
type Transaction struct {
	ID              int            `gorm:"primaryKey;autoIncrement"`
	InventoryID     int            `gorm:"not null;foreignKey:InventoryRefer;constraint:OnDelete:CASCADE"`
	TransactionType string         `gorm:"type:enum('in','out');not null"`
	Quantity        int            `gorm:"not null"`
	UserID          int            `gorm:"not null;foreignKey:UserRefer;constraint:OnDelete:CASCADE"`
	CreatedAt       time.Time      `gorm:"autoCreateTime"`
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}
