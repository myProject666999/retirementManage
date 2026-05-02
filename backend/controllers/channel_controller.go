package controllers

import (
	"net/http"
	"strconv"

	"retirementManage/database"
	"retirementManage/models"

	"github.com/gin-gonic/gin"
)

type CreateChannelRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

type UpdateChannelRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      *int   `json:"status"`
}

func GetChannels(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	name := c.DefaultQuery("name", "")
	statusStr := c.DefaultQuery("status", "")

	offset := (page - 1) * pageSize

	var channels []models.Channel
	var total int64

	query := database.DB.Model(&models.Channel{})
	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}
	if statusStr != "" {
		status, _ := strconv.Atoi(statusStr)
		query = query.Where("status = ?", status)
	}

	query.Count(&total)

	if err := query.Offset(offset).Limit(pageSize).Find(&channels).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "获取渠道列表失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data": gin.H{
			"list":      channels,
			"total":     total,
			"page":      page,
			"page_size": pageSize,
		},
	})
}

func GetAllChannels(c *gin.Context) {
	var channels []models.Channel
	if err := database.DB.Where("status = 1").Find(&channels).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "获取渠道列表失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data":    channels,
	})
}

func GetChannel(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var channel models.Channel
	if err := database.DB.First(&channel, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "渠道不存在",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data":    channel,
	})
}

func CreateChannel(c *gin.Context) {
	var req CreateChannelRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误: " + err.Error(),
		})
		return
	}

	var existingChannel models.Channel
	if database.DB.Where("name = ?", req.Name).First(&existingChannel).Error == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "渠道名称已存在",
		})
		return
	}

	channel := models.Channel{
		Name:        req.Name,
		Description: req.Description,
		Status:      1,
	}

	if err := database.DB.Create(&channel).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "创建渠道失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "创建成功",
		"data":    channel,
	})
}

func UpdateChannel(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var req UpdateChannelRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var channel models.Channel
	if err := database.DB.First(&channel, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "渠道不存在",
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
	if req.Status != nil {
		updates["status"] = *req.Status
	}

	if err := database.DB.Model(&channel).Updates(updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "更新渠道失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "更新成功",
	})
}

func DeleteChannel(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var channel models.Channel
	if err := database.DB.First(&channel, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "渠道不存在",
		})
		return
	}

	if err := database.DB.Delete(&channel).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "删除渠道失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "删除成功",
	})
}
