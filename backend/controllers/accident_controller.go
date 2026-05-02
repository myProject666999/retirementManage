package controllers

import (
	"net/http"
	"strconv"
	"time"

	"retirementManage/database"
	"retirementManage/models"

	"github.com/gin-gonic/gin"
)

type CreateAccidentRequest struct {
	ElderID      uint       `json:"elder_id" binding:"required"`
	AccidentTime *time.Time `json:"accident_time"`
	Location     string     `json:"location"`
	Type         int        `json:"type"`
	Title        string     `json:"title" binding:"required"`
	Description  string     `json:"description"`
	Handler      string     `json:"handler"`
	Remarks      string     `json:"remarks"`
}

type UpdateAccidentRequest struct {
	AccidentTime *time.Time `json:"accident_time"`
	Location     string     `json:"location"`
	Type         int        `json:"type"`
	Title        string     `json:"title"`
	Description  string     `json:"description"`
	Handler      string     `json:"handler"`
	Result       string     `json:"result"`
	Status       *int       `json:"status"`
	Remarks      string     `json:"remarks"`
}

func GetAccidents(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	statusStr := c.DefaultQuery("status", "")

	offset := (page - 1) * pageSize

	var accidents []models.Accident
	var total int64

	query := database.DB.Model(&models.Accident{}).Preload("Elder")
	if statusStr != "" {
		status, _ := strconv.Atoi(statusStr)
		query = query.Where("status = ?", status)
	}

	query.Count(&total)

	if err := query.Order("accident_time DESC").Offset(offset).Limit(pageSize).Find(&accidents).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "获取事故登记列表失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data": gin.H{
			"list":      accidents,
			"total":     total,
			"page":      page,
			"page_size": pageSize,
		},
	})
}

func GetAccident(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var accident models.Accident
	if err := database.DB.Preload("Elder").First(&accident, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "事故登记不存在",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data":    accident,
	})
}

func CreateAccident(c *gin.Context) {
	var req CreateAccidentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误: " + err.Error(),
		})
		return
	}

	accident := models.Accident{
		ElderID:      req.ElderID,
		AccidentTime: req.AccidentTime,
		Location:     req.Location,
		Type:         req.Type,
		Title:        req.Title,
		Description:  req.Description,
		Handler:      req.Handler,
		Remarks:      req.Remarks,
		Status:       0,
	}

	if err := database.DB.Create(&accident).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "创建事故登记失败",
		})
		return
	}

	database.DB.Preload("Elder").First(&accident, accident.ID)

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "创建成功",
		"data":    accident,
	})
}

func UpdateAccident(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var req UpdateAccidentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var accident models.Accident
	if err := database.DB.First(&accident, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "事故登记不存在",
		})
		return
	}

	updates := make(map[string]interface{})
	if req.AccidentTime != nil {
		updates["accident_time"] = req.AccidentTime
	}
	if req.Location != "" {
		updates["location"] = req.Location
	}
	if req.Type != 0 {
		updates["type"] = req.Type
	}
	if req.Title != "" {
		updates["title"] = req.Title
	}
	if req.Description != "" {
		updates["description"] = req.Description
	}
	if req.Handler != "" {
		updates["handler"] = req.Handler
	}
	if req.Result != "" {
		updates["result"] = req.Result
	}
	if req.Status != nil {
		updates["status"] = *req.Status
	}
	if req.Remarks != "" {
		updates["remarks"] = req.Remarks
	}

	if err := database.DB.Model(&accident).Updates(updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "更新事故登记失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "更新成功",
	})
}
