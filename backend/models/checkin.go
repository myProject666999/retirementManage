package models

import (
	"time"

	"gorm.io/gorm"
)

type Building struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" gorm:"not null;size:50"`
	Description string         `json:"description" gorm:"size:200"`
	FloorCount  int            `json:"floor_count" gorm:"default:1"`
	Status      int            `json:"status" gorm:"default:1"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

type RoomType struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" gorm:"not null;size:50"`
	Description string         `json:"description" gorm:"size:200"`
	BedCount    int            `json:"bed_count" gorm:"default:1"`
	Price       float64        `json:"price" gorm:"type:decimal(10,2)"`
	Status      int            `json:"status" gorm:"default:1"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

type Room struct {
	ID         uint           `json:"id" gorm:"primaryKey"`
	BuildingID uint           `json:"building_id"`
	Building   Building       `json:"building" gorm:"foreignKey:BuildingID"`
	RoomTypeID uint           `json:"room_type_id"`
	RoomType   RoomType       `json:"room_type" gorm:"foreignKey:RoomTypeID"`
	RoomNumber string         `json:"room_number" gorm:"size:20"`
	Floor      int            `json:"floor"`
	Status     int            `json:"status" gorm:"default:1"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"index"`
}

type Bed struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	RoomID      uint           `json:"room_id"`
	Room        Room           `json:"room" gorm:"foreignKey:RoomID"`
	BedNumber   string         `json:"bed_number" gorm:"size:20"`
	Status      int            `json:"status" gorm:"default:0"`
	ElderID     uint           `json:"elder_id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

type Contract struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	ContractNo   string         `json:"contract_no" gorm:"size:50;unique"`
	ElderID      uint           `json:"elder_id"`
	Elder        Elder          `json:"elder" gorm:"foreignKey:ElderID"`
	BedID        uint           `json:"bed_id"`
	Bed          Bed            `json:"bed" gorm:"foreignKey:BedID"`
	StartDate    *time.Time     `json:"start_date"`
	EndDate      *time.Time     `json:"end_date"`
	Amount       float64        `json:"amount" gorm:"type:decimal(10,2)"`
	Deposit      float64        `json:"deposit" gorm:"type:decimal(10,2)"`
	Status       int            `json:"status" gorm:"default:0"`
	Remarks      string         `json:"remarks" gorm:"type:text"`
	CreatedBy    uint           `json:"created_by"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index"`
}

type Outing struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	ElderID     uint           `json:"elder_id"`
	Elder       Elder          `json:"elder" gorm:"foreignKey:ElderID"`
	OutTime     *time.Time     `json:"out_time"`
	ReturnTime  *time.Time     `json:"return_time"`
	Companion   string         `json:"companion" gorm:"size:50"`
	Relation    string         `json:"relation" gorm:"size:50"`
	Phone       string         `json:"phone" gorm:"size:20"`
	Purpose     string         `json:"purpose" gorm:"size:200"`
	Status      int            `json:"status" gorm:"default:0"`
	Remarks     string         `json:"remarks" gorm:"type:text"`
	CreatedBy   uint           `json:"created_by"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

type Visit struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	ElderID     uint           `json:"elder_id"`
	Elder       Elder          `json:"elder" gorm:"foreignKey:ElderID"`
	VisitorName string         `json:"visitor_name" gorm:"size:50"`
	Relation    string         `json:"relation" gorm:"size:50"`
	Phone       string         `json:"phone" gorm:"size:20"`
	IDCard      string         `json:"id_card" gorm:"size:18"`
	VisitTime   *time.Time     `json:"visit_time"`
	LeaveTime   *time.Time     `json:"leave_time"`
	Status      int            `json:"status" gorm:"default:0"`
	Remarks     string         `json:"remarks" gorm:"type:text"`
	CreatedBy   uint           `json:"created_by"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

type Accident struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	ElderID     uint           `json:"elder_id"`
	Elder       Elder          `json:"elder" gorm:"foreignKey:ElderID"`
	AccidentTime *time.Time    `json:"accident_time"`
	Location    string         `json:"location" gorm:"size:100"`
	Type        int            `json:"type"`
	Title       string         `json:"title" gorm:"size:100"`
	Description string         `json:"description" gorm:"type:text"`
	Handler     string         `json:"handler" gorm:"size:50"`
	Result      string         `json:"result" gorm:"type:text"`
	Status      int            `json:"status" gorm:"default:0"`
	CreatedBy   uint           `json:"created_by"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

type Checkout struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	ContractID  uint           `json:"contract_id"`
	Contract    Contract       `json:"contract" gorm:"foreignKey:ContractID"`
	ElderID     uint           `json:"elder_id"`
	Elder       Elder          `json:"elder" gorm:"foreignKey:ElderID"`
	CheckoutDate *time.Time    `json:"checkout_date"`
	Reason      string         `json:"reason" gorm:"type:text"`
	RefundAmount float64       `json:"refund_amount" gorm:"type:decimal(10,2)"`
	Status      int            `json:"status" gorm:"default:0"`
	AuditBy     uint           `json:"audit_by"`
	AuditTime   *time.Time     `json:"audit_time"`
	AuditRemark string         `json:"audit_remark" gorm:"type:text"`
	CreatedBy   uint           `json:"created_by"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}
