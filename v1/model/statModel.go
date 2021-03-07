package model

import (
	"time"

	// "github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Stat struct {
	ID            uint16    `json:"id"`
	DataDate      time.Time `json:"data_date"`
	TotalReturn1y float32   `json:"total_return_1y" sql:"type:decimal(20,6);"`
	// Std1Y         float32         `json:"std_1y" gorm:"std_1y"`
	FundID    uint16
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// TableName stat
func (Stat) TableName() string {
	return "stat"
}
