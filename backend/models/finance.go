package models

import (
	"time"

	"gorm.io/gorm"
)

type Recharge struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	RechargeNo  string         `json:"recharge_no" gorm:"size:50;unique"`
	ElderID     uint           `json:"elder_id"`
	Elder       Elder          `json:"elder" gorm:"foreignKey:ElderID"`
	Amount      float64        `json:"amount" gorm:"type:decimal(10,2)"`
	GiftAmount  float64        `json:"gift_amount" gorm:"type:decimal(10,2)"`
	TotalAmount float64        `json:"total_amount" gorm:"type:decimal(10,2)"`
	PayMethod   int            `json:"pay_method"`
	PayTime     *time.Time     `json:"pay_time"`
	Status      int            `json:"status" gorm:"default:0"`
	Remarks     string         `json:"remarks" gorm:"type:text"`
	CreatedBy   uint           `json:"created_by"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

type ExpenseRecord struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	RecordNo    string         `json:"record_no" gorm:"size:50;unique"`
	ElderID     uint           `json:"elder_id"`
	Elder       Elder          `json:"elder" gorm:"foreignKey:ElderID"`
	Type        int            `json:"type"`
	Title       string         `json:"title" gorm:"size:100"`
	Amount      float64        `json:"amount" gorm:"type:decimal(10,2)"`
	Balance     float64        `json:"balance" gorm:"type:decimal(10,2)"`
	RelatedID   uint           `json:"related_id"`
	RelatedType string         `json:"related_type" gorm:"size:50"`
	Description string         `json:"description" gorm:"type:text"`
	Status      int            `json:"status" gorm:"default:1"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

type Account struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	ElderID     uint           `json:"elder_id" gorm:"unique"`
	Elder       Elder          `json:"elder" gorm:"foreignKey:ElderID"`
	Balance     float64        `json:"balance" gorm:"type:decimal(10,2);default:0"`
	TotalRecharge float64      `json:"total_recharge" gorm:"type:decimal(10,2);default:0"`
	TotalExpense float64       `json:"total_expense" gorm:"type:decimal(10,2);default:0"`
	Status      int            `json:"status" gorm:"default:1"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}
