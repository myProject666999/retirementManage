package models

import (
	"time"

	"gorm.io/gorm"
)

type Channel struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" gorm:"not null;size:50"`
	Description string         `json:"description" gorm:"size:200"`
	Status      int            `json:"status" gorm:"default:1"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

type Consultation struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" gorm:"size:50"`
	Phone       string         `json:"phone" gorm:"size:20"`
	Gender      int            `json:"gender"`
	Age         int            `json:"age"`
	ChannelID   uint           `json:"channel_id"`
	Channel     Channel        `json:"channel" gorm:"foreignKey:ChannelID"`
	Content     string         `json:"content" gorm:"type:text"`
	FollowUp    string         `json:"follow_up" gorm:"type:text"`
	Status      int            `json:"status" gorm:"default:0"`
	ConsultTime *time.Time     `json:"consult_time"`
	CreatedBy   uint           `json:"created_by"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

type Reservation struct {
	ID             uint           `json:"id" gorm:"primaryKey"`
	Name           string         `json:"name" gorm:"size:50"`
	Phone          string         `json:"phone" gorm:"size:20"`
	Gender         int            `json:"gender"`
	Age            int            `json:"age"`
	IDCard         string         `json:"id_card" gorm:"size:18"`
	RoomTypeID     uint           `json:"room_type_id"`
	RoomType       RoomType       `json:"room_type" gorm:"foreignKey:RoomTypeID"`
	ReservationDate *time.Time    `json:"reservation_date"`
	DepositAmount  float64        `json:"deposit_amount" gorm:"type:decimal(10,2)"`
	PaidAmount     float64        `json:"paid_amount" gorm:"type:decimal(10,2)"`
	Remarks        string         `json:"remarks" gorm:"type:text"`
	Status         int            `json:"status" gorm:"default:0"`
	CreatedBy      uint           `json:"created_by"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"-" gorm:"index"`
}
