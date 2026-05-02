package controllers

import (
	"net/http"
	"strconv"

	"retirementManage/database"
	"retirementManage/models"

	"github.com/gin-gonic/gin"
)

type CreateBuildingRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	FloorCount  int    `json:"floor_count"`
}

type UpdateBuildingRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	FloorCount  int    `json:"floor_count"`
	Status      *int   `json:"status"`
}

func GetBuildings(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	name := c.DefaultQuery("name", "")
	statusStr := c.DefaultQuery("status", "")

	offset := (page - 1) * pageSize

	var buildings []models.Building
	var total int64

	query := database.DB.Model(&models.Building{})
	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}
	if statusStr != "" {
		status, _ := strconv.Atoi(statusStr)
		query = query.Where("status = ?", status)
	}

	query.Count(&total)

	if err := query.Offset(offset).Limit(pageSize).Find(&buildings).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "获取楼栋列表失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data": gin.H{
			"list":      buildings,
			"total":     total,
			"page":      page,
			"page_size": pageSize,
		},
	})
}

func GetAllBuildings(c *gin.Context) {
	var buildings []models.Building
	if err := database.DB.Where("status = 1").Find(&buildings).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "获取楼栋列表失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data":    buildings,
	})
}

func GetBuilding(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var building models.Building
	if err := database.DB.First(&building, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "楼栋不存在",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data":    building,
	})
}

func CreateBuilding(c *gin.Context) {
	var req CreateBuildingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误: " + err.Error(),
		})
		return
	}

	var existingBuilding models.Building
	if database.DB.Where("name = ?", req.Name).First(&existingBuilding).Error == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "楼栋名称已存在",
		})
		return
	}

	building := models.Building{
		Name:        req.Name,
		Description: req.Description,
		FloorCount:  req.FloorCount,
		Status:      1,
	}

	if err := database.DB.Create(&building).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "创建楼栋失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "创建成功",
		"data":    building,
	})
}

func UpdateBuilding(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var req UpdateBuildingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var building models.Building
	if err := database.DB.First(&building, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "楼栋不存在",
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
	if req.FloorCount > 0 {
		updates["floor_count"] = req.FloorCount
	}
	if req.Status != nil {
		updates["status"] = *req.Status
	}

	if err := database.DB.Model(&building).Updates(updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "更新楼栋失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "更新成功",
	})
}

func DeleteBuilding(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var building models.Building
	if err := database.DB.First(&building, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "楼栋不存在",
		})
		return
	}

	if err := database.DB.Delete(&building).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "删除楼栋失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "删除成功",
	})
}
