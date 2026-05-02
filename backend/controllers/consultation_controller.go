package controllers

import (
	"net/http"
	"strconv"
	"time"

	"retirementManage/database"
	"retirementManage/models"

	"github.com/gin-gonic/gin"
)

type CreateConsultationRequest struct {
	Name        string     `json:"name" binding:"required"`
	Phone       string     `json:"phone" binding:"required"`
	Gender      int        `json:"gender"`
	Age         int        `json:"age"`
	ChannelID   uint       `json:"channel_id"`
	Content     string     `json:"content"`
	ConsultTime *time.Time `json:"consult_time"`
}

type UpdateConsultationRequest struct {
	Name        string     `json:"name"`
	Phone       string     `json:"phone"`
	Gender      int        `json:"gender"`
	Age         int        `json:"age"`
	ChannelID   uint       `json:"channel_id"`
	Content     string     `json:"content"`
	FollowUp    string     `json:"follow_up"`
	Status      *int       `json:"status"`
	ConsultTime *time.Time `json:"consult_time"`
}

func GetConsultations(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	name := c.DefaultQuery("name", "")
	phone := c.DefaultQuery("phone", "")
	statusStr := c.DefaultQuery("status", "")
	channelIDStr := c.DefaultQuery("channel_id", "")

	offset := (page - 1) * pageSize

	var consultations []models.Consultation
	var total int64

	query := database.DB.Model(&models.Consultation{}).Preload("Channel")
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
	if channelIDStr != "" {
		channelID, _ := strconv.Atoi(channelIDStr)
		query = query.Where("channel_id = ?", channelID)
	}

	query.Count(&total)

	if err := query.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&consultations).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "获取咨询列表失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data": gin.H{
			"list":      consultations,
			"total":     total,
			"page":      page,
			"page_size": pageSize,
		},
	})
}

func GetConsultation(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var consultation models.Consultation
	if err := database.DB.Preload("Channel").First(&consultation, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "咨询记录不存在",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data":    consultation,
	})
}

func CreateConsultation(c *gin.Context) {
	var req CreateConsultationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误: " + err.Error(),
		})
		return
	}

	consultation := models.Consultation{
		Name:        req.Name,
		Phone:       req.Phone,
		Gender:      req.Gender,
		Age:         req.Age,
		ChannelID:   req.ChannelID,
		Content:     req.Content,
		ConsultTime: req.ConsultTime,
		Status:      0,
	}

	if err := database.DB.Create(&consultation).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "创建咨询记录失败",
		})
		return
	}

	database.DB.Preload("Channel").First(&consultation, consultation.ID)

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "创建成功",
		"data":    consultation,
	})
}

func UpdateConsultation(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var req UpdateConsultationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var consultation models.Consultation
	if err := database.DB.First(&consultation, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "咨询记录不存在",
		})
		return
	}

	updates := make(map[string]interface{})
	if req.Name != "" {
		updates["name"] = req.Name
	}
	if req.Phone != "" {
		updates["phone"] = req.Phone
	}
	if req.Gender != 0 {
		updates["gender"] = req.Gender
	}
	if req.Age > 0 {
		updates["age"] = req.Age
	}
	if req.ChannelID > 0 {
		updates["channel_id"] = req.ChannelID
	}
	if req.Content != "" {
		updates["content"] = req.Content
	}
	if req.FollowUp != "" {
		updates["follow_up"] = req.FollowUp
	}
	if req.Status != nil {
		updates["status"] = *req.Status
	}
	if req.ConsultTime != nil {
		updates["consult_time"] = req.ConsultTime
	}

	if err := database.DB.Model(&consultation).Updates(updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "更新咨询记录失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "更新成功",
	})
}

func DeleteConsultation(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var consultation models.Consultation
	if err := database.DB.First(&consultation, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "咨询记录不存在",
		})
		return
	}

	if err := database.DB.Delete(&consultation).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "删除咨询记录失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "删除成功",
	})
}
