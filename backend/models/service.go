package models

import (
	"time"

	"gorm.io/gorm"
)

type ServiceItem struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" gorm:"not null;size:100"`
	Code        string         `json:"code" gorm:"size:50"`
	Type        int            `json:"type"`
	Description string         `json:"description" gorm:"type:text"`
	Price       float64        `json:"price" gorm:"type:decimal(10,2)"`
	Duration    int            `json:"duration" gorm:"default:0"`
	Unit        string         `json:"unit" gorm:"size:20"`
	Status      int            `json:"status" gorm:"default:1"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

type CareLevel struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" gorm:"not null;size:50"`
	Code        string         `json:"code" gorm:"size:50"`
	Level       int            `json:"level"`
	Description string         `json:"description" gorm:"type:text"`
	Price       float64        `json:"price" gorm:"type:decimal(10,2)"`
	Status      int            `json:"status" gorm:"default:1"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

type ServiceReservation struct {
	ID             uint           `json:"id" gorm:"primaryKey"`
	ElderID        uint           `json:"elder_id"`
	Elder          Elder          `json:"elder" gorm:"foreignKey:ElderID"`
	ServiceItemID  uint           `json:"service_item_id"`
	ServiceItem    ServiceItem    `json:"service_item" gorm:"foreignKey:ServiceItemID"`
	ServiceTime    *time.Time     `json:"service_time"`
	ServicePlace   string         `json:"service_place" gorm:"size:100"`
	Quantity       int            `json:"quantity" gorm:"default:1"`
	Amount         float64        `json:"amount" gorm:"type:decimal(10,2)"`
	Status         int            `json:"status" gorm:"default:0"`
	ServiceBy      uint           `json:"service_by"`
	ServiceResult  string         `json:"service_result" gorm:"type:text"`
	Remarks        string         `json:"remarks" gorm:"type:text"`
	CreatedBy      uint           `json:"created_by"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"-" gorm:"index"`
}
