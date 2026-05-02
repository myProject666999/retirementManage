package models

import (
	"time"

	"gorm.io/gorm"
)

const (
	RoleSales        = "sales"
	RoleHR           = "hr"
	RoleService      = "service"
	RoleCatering     = "catering"
	RoleFinance      = "finance"
	RoleSuperAdmin   = "super_admin"
)

type Role struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" gorm:"unique;not null;size:50"`
	Description string         `json:"description" gorm:"size:200"`
	Code        string         `json:"code" gorm:"unique;not null;size:50"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

var DefaultRoles = []Role{
	{Name: "销售员", Description: "营销管理", Code: RoleSales},
	{Name: "人事", Description: "人员管理", Code: RoleHR},
	{Name: "服务员", Description: "服务管理", Code: RoleService},
	{Name: "餐饮员", Description: "餐饮管理", Code: RoleCatering},
	{Name: "财务人员", Description: "费用管理", Code: RoleFinance},
	{Name: "超级管理员", Description: "系统管理员", Code: RoleSuperAdmin},
}
