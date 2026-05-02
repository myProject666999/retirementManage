package controllers

import (
	"net/http"
	"strconv"

	"retirementManage/database"
	"retirementManage/models"

	"github.com/gin-gonic/gin"
)

type CreateCareLevelRequest struct {
	Name        string  `json:"name" binding:"required"`
	Code        string  `json:"code"`
	Level       int     `json:"level"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required"`
}

type UpdateCareLevelRequest struct {
	Name        string  `json:"name"`
	Code        string  `json:"code"`
	Level       int     `json:"level"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Status      *int    `json:"status"`
}

func GetCareLevels(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	name := c.DefaultQuery("name", "")
	statusStr := c.DefaultQuery("status", "")

	offset := (page - 1) * pageSize

	var levels []models.CareLevel
	var total int64

	query := database.DB.Model(&models.CareLevel{})
	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}
	if statusStr != "" {
		status, _ := strconv.Atoi(statusStr)
		query = query.Where("status = ?", status)
	}

	query.Count(&total)

	if err := query.Order("level ASC").Offset(offset).Limit(pageSize).Find(&levels).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "获取护理等级列表失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data": gin.H{
			"list":      levels,
			"total":     total,
			"page":      page,
			"page_size": pageSize,
		},
	})
}

func GetAllCareLevels(c *gin.Context) {
	var levels []models.CareLevel
	if err := database.DB.Where("status = 1").Order("level ASC").Find(&levels).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "获取护理等级列表失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data":    levels,
	})
}

func GetCareLevel(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var level models.CareLevel
	if err := database.DB.First(&level, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "护理等级不存在",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data":    level,
	})
}

func CreateCareLevel(c *gin.Context) {
	var req CreateCareLevelRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误: " + err.Error(),
		})
		return
	}

	level := models.CareLevel{
		Name:        req.Name,
		Code:        req.Code,
		Level:       req.Level,
		Description: req.Description,
		Price:       req.Price,
		Status:      1,
	}

	if err := database.DB.Create(&level).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "创建护理等级失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "创建成功",
		"data":    level,
	})
}

func UpdateCareLevel(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var req UpdateCareLevelRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var level models.CareLevel
	if err := database.DB.First(&level, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "护理等级不存在",
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
	if req.Level > 0 {
		updates["level"] = req.Level
	}
	if req.Description != "" {
		updates["description"] = req.Description
	}
	if req.Price >= 0 {
		updates["price"] = req.Price
	}
	if req.Status != nil {
		updates["status"] = *req.Status
	}

	if err := database.DB.Model(&level).Updates(updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "更新护理等级失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "更新成功",
	})
}

func DeleteCareLevel(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var level models.CareLevel
	if err := database.DB.First(&level, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "护理等级不存在",
		})
		return
	}

	if err := database.DB.Delete(&level).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "删除护理等级失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "删除成功",
	})
}
