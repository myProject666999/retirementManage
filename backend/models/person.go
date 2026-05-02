package models

import (
	"time"

	"gorm.io/gorm"
)

type Elder struct {
	ID              uint           `json:"id" gorm:"primaryKey"`
	Name            string         `json:"name" gorm:"size:50"`
	Gender          int            `json:"gender"`
	BirthDate       *time.Time     `json:"birth_date"`
	Age             int            `json:"age"`
	IDCard          string         `json:"id_card" gorm:"size:18;unique"`
	Phone           string         `json:"phone" gorm:"size:20"`
	MaritalStatus   int            `json:"marital_status"`
	Education       string         `json:"education" gorm:"size:50"`
	Occupation      string         `json:"occupation" gorm:"size:50"`
	NativePlace     string         `json:"native_place" gorm:"size:100"`
	Address         string         `json:"address" gorm:"size:200"`
	HealthStatus    int            `json:"health_status"`
	AllergyHistory  string         `json:"allergy_history" gorm:"type:text"`
	MedicalHistory  string         `json:"medical_history" gorm:"type:text"`
	CareLevelID     uint           `json:"care_level_id"`
	CareLevel       CareLevel      `json:"care_level" gorm:"foreignKey:CareLevelID"`
	Status          int            `json:"status" gorm:"default:0"`
	Remarks         string         `json:"remarks" gorm:"type:text"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"-" gorm:"index"`
	Contacts        []Contact      `json:"contacts" gorm:"foreignKey:ElderID"`
}

type Contact struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	ElderID     uint           `json:"elder_id"`
	Elder       Elder          `json:"elder" gorm:"foreignKey:ElderID"`
	Name        string         `json:"name" gorm:"size:50"`
	Relation    string         `json:"relation" gorm:"size:50"`
	Phone       string         `json:"phone" gorm:"size:20"`
	WorkUnit    string         `json:"work_unit" gorm:"size:100"`
	Address     string         `json:"address" gorm:"size:200"`
	IsEmergency int            `json:"is_emergency" gorm:"default:0"`
	Status      int            `json:"status" gorm:"default:1"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

type Employee struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	Name         string         `json:"name" gorm:"size:50"`
	Gender       int            `json:"gender"`
	BirthDate    *time.Time     `json:"birth_date"`
	IDCard       string         `json:"id_card" gorm:"size:18;unique"`
	Phone        string         `json:"phone" gorm:"size:20"`
	Email        string         `json:"email" gorm:"size:100"`
	DepartmentID uint           `json:"department_id"`
	Department   Department     `json:"department" gorm:"foreignKey:DepartmentID"`
	Position     string         `json:"position" gorm:"size:50"`
	EntryDate    *time.Time     `json:"entry_date"`
	LeaveDate    *time.Time     `json:"leave_date"`
	Status       int            `json:"status" gorm:"default:1"`
	Remarks      string         `json:"remarks" gorm:"type:text"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index"`
}

type Department struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" gorm:"not null;size:50"`
	Description string         `json:"description" gorm:"size:200"`
	ParentID    uint           `json:"parent_id" gorm:"default:0"`
	Status      int            `json:"status" gorm:"default:1"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}
