package main

import (
	"fmt"
	"log"

	"retirementManage/config"
	"retirementManage/database"
	"retirementManage/models"
	"retirementManage/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitConfig()

	gin.SetMode(config.AppConfig.Server.Mode)

	database.InitDB()

	err := database.DB.AutoMigrate(
		&models.Role{},
		&models.User{},
		&models.Channel{},
		&models.Consultation{},
		&models.Reservation{},
		&models.Building{},
		&models.RoomType{},
		&models.Room{},
		&models.Bed{},
		&models.Contract{},
		&models.Outing{},
		&models.Visit{},
		&models.Accident{},
		&models.Checkout{},
		&models.CareLevel{},
		&models.Elder{},
		&models.Contact{},
		&models.Department{},
		&models.Employee{},
		&models.ServiceItem{},
		&models.ServiceReservation{},
		&models.Dish{},
		&models.MealPackage{},
		&models.PackageItem{},
		&models.Order{},
		&models.OrderItem{},
		&models.Account{},
		&models.Recharge{},
		&models.ExpenseRecord{},
	)
	if err != nil {
		log.Fatalf("数据库迁移失败: %v", err)
	}
	log.Println("数据库迁移成功")

	initDefaultData()

	r := routers.SetupRouter()

	addr := fmt.Sprintf(":%d", config.AppConfig.Server.Port)
	log.Printf("服务器启动在 %s", addr)

	if err := r.Run(addr); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}

func initDefaultData() {
	for _, role := range models.DefaultRoles {
		var existingRole models.Role
		if database.DB.Where("code = ?", role.Code).First(&existingRole).Error != nil {
			database.DB.Create(&role)
			log.Printf("创建默认角色: %s", role.Name)
		}
	}

	var superAdminRole models.Role
	if err := database.DB.Where("code = ?", models.RoleSuperAdmin).First(&superAdminRole).Error; err != nil {
		log.Printf("超级管理员角色不存在，跳过默认管理员创建")
		return
	}

	var adminUser models.User
	if database.DB.Where("username = ?", "admin").First(&adminUser).Error == nil {
		log.Println("默认管理员已存在")
		return
	}

	adminUser = models.User{
		Username: "admin",
		Password: "admin123",
		RealName: "超级管理员",
		Phone:    "13800138000",
		Email:    "admin@example.com",
		RoleID:   superAdminRole.ID,
		Status:   1,
	}

	if err := adminUser.HashPassword(); err != nil {
		log.Printf("密码加密失败: %v", err)
		return
	}

	if err := database.DB.Create(&adminUser).Error; err != nil {
		log.Printf("创建默认管理员失败: %v", err)
		return
	}

	log.Println("创建默认管理员成功: admin / admin123")
}
