package models

import (
	"time"

	"gorm.io/gorm"
)

type Dish struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" gorm:"not null;size:100"`
	Code        string         `json:"code" gorm:"size:50"`
	Type        int            `json:"type"`
	Price       float64        `json:"price" gorm:"type:decimal(10,2)"`
	Unit        string         `json:"unit" gorm:"size:20"`
	Description string         `json:"description" gorm:"type:text"`
	Image       string         `json:"image" gorm:"size:255"`
	IsHot       int            `json:"is_hot" gorm:"default:0"`
	Status      int            `json:"status" gorm:"default:1"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

type MealPackage struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" gorm:"not null;size:100"`
	Code        string         `json:"code" gorm:"size:50"`
	Type        int            `json:"type"`
	Price       float64        `json:"price" gorm:"type:decimal(10,2)"`
	Description string         `json:"description" gorm:"type:text"`
	Image       string         `json:"image" gorm:"size:255"`
	Status      int            `json:"status" gorm:"default:1"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
	Items       []PackageItem  `json:"items" gorm:"foreignKey:PackageID"`
}

type PackageItem struct {
	ID         uint           `json:"id" gorm:"primaryKey"`
	PackageID  uint           `json:"package_id"`
	DishID     uint           `json:"dish_id"`
	Dish       Dish           `json:"dish" gorm:"foreignKey:DishID"`
	Quantity   int            `json:"quantity" gorm:"default:1"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"index"`
}

type Order struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	OrderNo     string         `json:"order_no" gorm:"size:50;unique"`
	ElderID     uint           `json:"elder_id"`
	Elder       Elder          `json:"elder" gorm:"foreignKey:ElderID"`
	OrderTime   *time.Time     `json:"order_time"`
	DeliveryTime *time.Time    `json:"delivery_time"`
	DeliveryPlace string        `json:"delivery_place" gorm:"size:100"`
	TotalAmount float64        `json:"total_amount" gorm:"type:decimal(10,2)"`
	PaidAmount  float64        `json:"paid_amount" gorm:"type:decimal(10,2)"`
	Status      int            `json:"status" gorm:"default:0"`
	Remarks     string         `json:"remarks" gorm:"type:text"`
	CreatedBy   uint           `json:"created_by"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
	Items       []OrderItem    `json:"items" gorm:"foreignKey:OrderID"`
}

type OrderItem struct {
	ID         uint           `json:"id" gorm:"primaryKey"`
	OrderID    uint           `json:"order_id"`
	DishID     uint           `json:"dish_id"`
	Dish       Dish           `json:"dish" gorm:"foreignKey:DishID"`
	Quantity   int            `json:"quantity" gorm:"default:1"`
	Price      float64        `json:"price" gorm:"type:decimal(10,2)"`
	Amount     float64        `json:"amount" gorm:"type:decimal(10,2)"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"index"`
}
