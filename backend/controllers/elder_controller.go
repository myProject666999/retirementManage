package controllers

import (
	"net/http"
	"strconv"
	"time"

	"retirementManage/database"
	"retirementManage/models"

	"github.com/gin-gonic/gin"
)

type CreateElderRequest struct {
	Name           string     `json:"name" binding:"required"`
	Gender         int        `json:"gender"`
	BirthDate      *time.Time `json:"birth_date"`
	Age            int        `json:"age"`
	IDCard         string     `json:"id_card" binding:"required"`
	Phone          string     `json:"phone"`
	MaritalStatus  int        `json:"marital_status"`
	Education      string     `json:"education"`
	Occupation     string     `json:"occupation"`
	NativePlace    string     `json:"native_place"`
	Address        string     `json:"address"`
	HealthStatus   int        `json:"health_status"`
	AllergyHistory string     `json:"allergy_history"`
	MedicalHistory string     `json:"medical_history"`
	CareLevelID    uint       `json:"care_level_id"`
	Remarks        string     `json:"remarks"`
}

type UpdateElderRequest struct {
	Name           string     `json:"name"`
	Gender         int        `json:"gender"`
	BirthDate      *time.Time `json:"birth_date"`
	Age            int        `json:"age"`
	Phone          string     `json:"phone"`
	MaritalStatus  int        `json:"marital_status"`
	Education      string     `json:"education"`
	Occupation     string     `json:"occupation"`
	NativePlace    string     `json:"native_place"`
	Address        string     `json:"address"`
	HealthStatus   int        `json:"health_status"`
	AllergyHistory string     `json:"allergy_history"`
	MedicalHistory string     `json:"medical_history"`
	CareLevelID    uint       `json:"care_level_id"`
	Status         *int       `json:"status"`
	Remarks        string     `json:"remarks"`
}

func GetElders(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	name := c.DefaultQuery("name", "")
	phone := c.DefaultQuery("phone", "")
	statusStr := c.DefaultQuery("status", "")

	offset := (page - 1) * pageSize

	var elders []models.Elder
	var total int64

	query := database.DB.Model(&models.Elder{}).Preload("CareLevel")
	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}
	if phone != "" {
		query = query.Where("phone LIKE ?", "%"+phone+"%")
	}
	if statusStr != "" {
		status, _ := strconv.Atoi(statusStr)
		query = query.Where("status = ?", status)
	}

	query.Count(&total)

	if err := query.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&elders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "获取长者列表失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data": gin.H{
			"list":      elders,
			"total":     total,
			"page":      page,
			"page_size": pageSize,
		},
	})
}

func GetAllElders(c *gin.Context) {
	var elders []models.Elder
	if err := database.DB.Where("status IN ?", []int{0, 1}).Preload("CareLevel").Find(&elders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "获取长者列表失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data":    elders,
	})
}

func GetElder(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var elder models.Elder
	if err := database.DB.Preload("CareLevel").Preload("Contacts").First(&elder, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "长者不存在",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data":    elder,
	})
}

func CreateElder(c *gin.Context) {
	var req CreateElderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误: " + err.Error(),
		})
		return
	}

	var existingElder models.Elder
	if database.DB.Where("id_card = ?", req.IDCard).First(&existingElder).Error == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "身份证号已存在",
		})
		return
	}

	elder := models.Elder{
		Name:           req.Name,
		Gender:         req.Gender,
		BirthDate:      req.BirthDate,
		Age:            req.Age,
		IDCard:         req.IDCard,
		Phone:          req.Phone,
		MaritalStatus:  req.MaritalStatus,
		Education:      req.Education,
		Occupation:     req.Occupation,
		NativePlace:    req.NativePlace,
		Address:        req.Address,
		HealthStatus:   req.HealthStatus,
		AllergyHistory: req.AllergyHistory,
		MedicalHistory: req.MedicalHistory,
		CareLevelID:    req.CareLevelID,
		Remarks:        req.Remarks,
		Status:         0,
	}

	if err := database.DB.Create(&elder).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "创建长者档案失败",
		})
		return
	}

	database.DB.Preload("CareLevel").First(&elder, elder.ID)

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "创建成功",
		"data":    elder,
	})
}

func UpdateElder(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var req UpdateElderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var elder models.Elder
	if err := database.DB.First(&elder, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "长者不存在",
		})
		return
	}

	updates := make(map[string]interface{})
	if req.Name != "" {
		updates["name"] = req.Name
	}
	if req.Gender != 0 {
		updates["gender"] = req.Gender
	}
	if req.BirthDate != nil {
		updates["birth_date"] = req.BirthDate
	}
	if req.Age > 0 {
		updates["age"] = req.Age
	}
	if req.Phone != "" {
		updates["phone"] = req.Phone
	}
	if req.MaritalStatus != 0 {
		updates["marital_status"] = req.MaritalStatus
	}
	if req.Education != "" {
		updates["education"] = req.Education
	}
	if req.Occupation != "" {
		updates["occupation"] = req.Occupation
	}
	if req.NativePlace != "" {
		updates["native_place"] = req.NativePlace
	}
	if req.Address != "" {
		updates["address"] = req.Address
	}
	if req.HealthStatus != 0 {
		updates["health_status"] = req.HealthStatus
	}
	if req.AllergyHistory != "" {
		updates["allergy_history"] = req.AllergyHistory
	}
	if req.MedicalHistory != "" {
		updates["medical_history"] = req.MedicalHistory
	}
	if req.CareLevelID > 0 {
		updates["care_level_id"] = req.CareLevelID
	}
	if req.Status != nil {
		updates["status"] = *req.Status
	}
	if req.Remarks != "" {
		updates["remarks"] = req.Remarks
	}

	if err := database.DB.Model(&elder).Updates(updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "更新长者档案失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "更新成功",
	})
}

func DeleteElder(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var elder models.Elder
	if err := database.DB.First(&elder, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "长者不存在",
		})
		return
	}

	if err := database.DB.Delete(&elder).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "删除长者档案失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "删除成功",
	})
}
