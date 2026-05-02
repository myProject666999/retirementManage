package controllers

import (
	"net/http"
	"strconv"

	"retirementManage/database"
	"retirementManage/models"

	"github.com/gin-gonic/gin"
)

type CreateServiceItemRequest struct {
	Name        string  `json:"name" binding:"required"`
	Code        string  `json:"code"`
	Type        int     `json:"type"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required"`
	Duration    int     `json:"duration"`
	Unit        string  `json:"unit"`
}

type UpdateServiceItemRequest struct {
	Name        string  `json:"name"`
	Code        string  `json:"code"`
	Type        int     `json:"type"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Duration    int     `json:"duration"`
	Unit        string  `json:"unit"`
	Status      *int    `json:"status"`
}

func GetServiceItems(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	name := c.DefaultQuery("name", "")
	typeStr := c.DefaultQuery("type", "")
	statusStr := c.DefaultQuery("status", "")

	offset := (page - 1) * pageSize

	var items []models.ServiceItem
	var total int64

	query := database.DB.Model(&models.ServiceItem{})
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

	if err := query.Offset(offset).Limit(pageSize).Find(&items).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "获取服务项目列表失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data": gin.H{
			"list":      items,
			"total":     total,
			"page":      page,
			"page_size": pageSize,
		},
	})
}

func GetAllServiceItems(c *gin.Context) {
	var items []models.ServiceItem
	if err := database.DB.Where("status = 1").Find(&items).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "获取服务项目列表失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data":    items,
	})
}

func GetServiceItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var item models.ServiceItem
	if err := database.DB.First(&item, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "服务项目不存在",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data":    item,
	})
}

func CreateServiceItem(c *gin.Context) {
	var req CreateServiceItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误: " + err.Error(),
		})
		return
	}

	item := models.ServiceItem{
		Name:        req.Name,
		Code:        req.Code,
		Type:        req.Type,
		Description: req.Description,
		Price:       req.Price,
		Duration:    req.Duration,
		Unit:        req.Unit,
		Status:      1,
	}

	if err := database.DB.Create(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "创建服务项目失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "创建成功",
		"data":    item,
	})
}

func UpdateServiceItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var req UpdateServiceItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var item models.ServiceItem
	if err := database.DB.First(&item, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "服务项目不存在",
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
	if req.Description != "" {
		updates["description"] = req.Description
	}
	if req.Price >= 0 {
		updates["price"] = req.Price
	}
	if req.Duration > 0 {
		updates["duration"] = req.Duration
	}
	if req.Unit != "" {
		updates["unit"] = req.Unit
	}
	if req.Status != nil {
		updates["status"] = *req.Status
	}

	if err := database.DB.Model(&item).Updates(updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "更新服务项目失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "更新成功",
	})
}

func DeleteServiceItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var item models.ServiceItem
	if err := database.DB.First(&item, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "服务项目不存在",
		})
		return
	}

	if err := database.DB.Delete(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "删除服务项目失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "删除成功",
	})
}
