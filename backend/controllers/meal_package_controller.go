package controllers

import (
	"net/http"
	"strconv"

	"retirementManage/database"
	"retirementManage/models"

	"github.com/gin-gonic/gin"
)

type CreateMealPackageRequest struct {
	Name        string  `json:"name" binding:"required"`
	Code        string  `json:"code"`
	Type        int     `json:"type"`
	Price       float64 `json:"price" binding:"required"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
}

type UpdateMealPackageRequest struct {
	Name        string  `json:"name"`
	Code        string  `json:"code"`
	Type        int     `json:"type"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Status      *int    `json:"status"`
}

func GetMealPackages(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	name := c.DefaultQuery("name", "")
	typeStr := c.DefaultQuery("type", "")
	statusStr := c.DefaultQuery("status", "")

	offset := (page - 1) * pageSize

	var packages []models.MealPackage
	var total int64

	query := database.DB.Model(&models.MealPackage{}).Preload("Items.Dish")
	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}
	if typeStr != "" {
		typeVal, _ := strconv.Atoi(typeStr)
		query = query.Where("type = ?", typeVal)
	}
	if statusStr != "" {
		status, _ := strconv.Atoi(statusStr)
		query = query.Where("status = ?", status)
	}

	query.Count(&total)

	if err := query.Offset(offset).Limit(pageSize).Find(&packages).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "获取套餐列表失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data": gin.H{
			"list":      packages,
			"total":     total,
			"page":      page,
			"page_size": pageSize,
		},
	})
}

func GetAllMealPackages(c *gin.Context) {
	var packages []models.MealPackage
	if err := database.DB.Where("status = 1").Preload("Items.Dish").Find(&packages).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "获取套餐列表失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data":    packages,
	})
}

func GetMealPackage(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var pkg models.MealPackage
	if err := database.DB.Preload("Items.Dish").First(&pkg, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "套餐不存在",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data":    pkg,
	})
}

func CreateMealPackage(c *gin.Context) {
	var req CreateMealPackageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误: " + err.Error(),
		})
		return
	}

	pkg := models.MealPackage{
		Name:        req.Name,
		Code:        req.Code,
		Type:        req.Type,
		Price:       req.Price,
		Description: req.Description,
		Image:       req.Image,
		Status:      1,
	}

	if err := database.DB.Create(&pkg).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "创建套餐失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "创建成功",
		"data":    pkg,
	})
}

func UpdateMealPackage(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var req UpdateMealPackageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var pkg models.MealPackage
	if err := database.DB.First(&pkg, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "套餐不存在",
		})
		return
	}

	updates := make(map[string]interface{})
	if req.Name != "" {
		updates["name"] = req.Name
	}
	if req.Code != "" {
		updates["code"] = req.Code
	}
	if req.Type != 0 {
		updates["type"] = req.Type
	}
	if req.Price >= 0 {
		updates["price"] = req.Price
	}
	if req.Description != "" {
		updates["description"] = req.Description
	}
	if req.Image != "" {
		updates["image"] = req.Image
	}
	if req.Status != nil {
		updates["status"] = *req.Status
	}

	if err := database.DB.Model(&pkg).Updates(updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "更新套餐失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "更新成功",
	})
}

func DeleteMealPackage(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var pkg models.MealPackage
	if err := database.DB.First(&pkg, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "套餐不存在",
		})
		return
	}

	if err := database.DB.Delete(&pkg).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "删除套餐失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "删除成功",
	})
}
