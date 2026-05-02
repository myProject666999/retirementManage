package controllers

import (
	"net/http"
	"strconv"

	"retirementManage/database"
	"retirementManage/models"

	"github.com/gin-gonic/gin"
)

type CreateDishRequest struct {
	Name        string  `json:"name" binding:"required"`
	Code        string  `json:"code"`
	Type        int     `json:"type"`
	Price       float64 `json:"price" binding:"required"`
	Unit        string  `json:"unit"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	IsHot       int     `json:"is_hot"`
}

type UpdateDishRequest struct {
	Name        string  `json:"name"`
	Code        string  `json:"code"`
	Type        int     `json:"type"`
	Price       float64 `json:"price"`
	Unit        string  `json:"unit"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	IsHot       int     `json:"is_hot"`
	Status      *int    `json:"status"`
}

func GetDishes(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	name := c.DefaultQuery("name", "")
	typeStr := c.DefaultQuery("type", "")
	statusStr := c.DefaultQuery("status", "")

	offset := (page - 1) * pageSize

	var dishes []models.Dish
	var total int64

	query := database.DB.Model(&models.Dish{})
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

	if err := query.Offset(offset).Limit(pageSize).Find(&dishes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "获取菜品列表失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data": gin.H{
			"list":      dishes,
			"total":     total,
			"page":      page,
			"page_size": pageSize,
		},
	})
}

func GetAllDishes(c *gin.Context) {
	var dishes []models.Dish
	if err := database.DB.Where("status = 1").Find(&dishes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "获取菜品列表失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data":    dishes,
	})
}

func GetDish(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var dish models.Dish
	if err := database.DB.First(&dish, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "菜品不存在",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data":    dish,
	})
}

func CreateDish(c *gin.Context) {
	var req CreateDishRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误: " + err.Error(),
		})
		return
	}

	dish := models.Dish{
		Name:        req.Name,
		Code:        req.Code,
		Type:        req.Type,
		Price:       req.Price,
		Unit:        req.Unit,
		Description: req.Description,
		Image:       req.Image,
		IsHot:       req.IsHot,
		Status:      1,
	}

	if err := database.DB.Create(&dish).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "创建菜品失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "创建成功",
		"data":    dish,
	})
}

func UpdateDish(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var req UpdateDishRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var dish models.Dish
	if err := database.DB.First(&dish, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "菜品不存在",
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
	if req.Unit != "" {
		updates["unit"] = req.Unit
	}
	if req.Description != "" {
		updates["description"] = req.Description
	}
	if req.Image != "" {
		updates["image"] = req.Image
	}
	if req.IsHot != 0 {
		updates["is_hot"] = req.IsHot
	}
	if req.Status != nil {
		updates["status"] = *req.Status
	}

	if err := database.DB.Model(&dish).Updates(updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "更新菜品失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "更新成功",
	})
}

func DeleteDish(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var dish models.Dish
	if err := database.DB.First(&dish, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "菜品不存在",
		})
		return
	}

	if err := database.DB.Delete(&dish).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "删除菜品失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "删除成功",
	})
}
