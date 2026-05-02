package controllers

import (
	"net/http"
	"strconv"

	"retirementManage/database"
	"retirementManage/models"

	"github.com/gin-gonic/gin"
)

type CreateRoomTypeRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	BedCount    int     `json:"bed_count"`
	Price       float64 `json:"price" binding:"required"`
}

type UpdateRoomTypeRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	BedCount    int     `json:"bed_count"`
	Price       float64 `json:"price"`
	Status      *int    `json:"status"`
}

func GetRoomTypes(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	name := c.DefaultQuery("name", "")
	statusStr := c.DefaultQuery("status", "")

	offset := (page - 1) * pageSize

	var roomTypes []models.RoomType
	var total int64

	query := database.DB.Model(&models.RoomType{})
	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}
	if statusStr != "" {
		status, _ := strconv.Atoi(statusStr)
		query = query.Where("status = ?", status)
	}

	query.Count(&total)

	if err := query.Offset(offset).Limit(pageSize).Find(&roomTypes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "获取房间类型列表失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data": gin.H{
			"list":      roomTypes,
			"total":     total,
			"page":      page,
			"page_size": pageSize,
		},
	})
}

func GetAllRoomTypes(c *gin.Context) {
	var roomTypes []models.RoomType
	if err := database.DB.Where("status = 1").Find(&roomTypes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "获取房间类型列表失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data":    roomTypes,
	})
}

func GetRoomType(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var roomType models.RoomType
	if err := database.DB.First(&roomType, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "房间类型不存在",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data":    roomType,
	})
}

func CreateRoomType(c *gin.Context) {
	var req CreateRoomTypeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误: " + err.Error(),
		})
		return
	}

	var existingRoomType models.RoomType
	if database.DB.Where("name = ?", req.Name).First(&existingRoomType).Error == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "房间类型名称已存在",
		})
		return
	}

	roomType := models.RoomType{
		Name:        req.Name,
		Description: req.Description,
		BedCount:    req.BedCount,
		Price:       req.Price,
		Status:      1,
	}

	if err := database.DB.Create(&roomType).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "创建房间类型失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "创建成功",
		"data":    roomType,
	})
}

func UpdateRoomType(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var req UpdateRoomTypeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var roomType models.RoomType
	if err := database.DB.First(&roomType, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "房间类型不存在",
		})
		return
	}

	updates := make(map[string]interface{})
	if req.Name != "" {
		updates["name"] = req.Name
	}
	if req.Description != "" {
		updates["description"] = req.Description
	}
	if req.BedCount > 0 {
		updates["bed_count"] = req.BedCount
	}
	if req.Price >= 0 {
		updates["price"] = req.Price
	}
	if req.Status != nil {
		updates["status"] = *req.Status
	}

	if err := database.DB.Model(&roomType).Updates(updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "更新房间类型失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "更新成功",
	})
}

func DeleteRoomType(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var roomType models.RoomType
	if err := database.DB.First(&roomType, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "房间类型不存在",
		})
		return
	}

	if err := database.DB.Delete(&roomType).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "删除房间类型失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "删除成功",
	})
}
