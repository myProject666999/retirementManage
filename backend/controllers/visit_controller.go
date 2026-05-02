package controllers

import (
	"net/http"
	"strconv"
	"time"

	"retirementManage/database"
	"retirementManage/models"

	"github.com/gin-gonic/gin"
)

type CreateVisitRequest struct {
	ElderID     uint       `json:"elder_id" binding:"required"`
	VisitorName string     `json:"visitor_name" binding:"required"`
	Relation    string     `json:"relation"`
	Phone       string     `json:"phone"`
	IDCard      string     `json:"id_card"`
	VisitTime   *time.Time `json:"visit_time"`
	Remarks     string     `json:"remarks"`
}

type UpdateVisitRequest struct {
	LeaveTime *time.Time `json:"leave_time"`
	Status    *int       `json:"status"`
	Remarks   string     `json:"remarks"`
}

func GetVisits(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	statusStr := c.DefaultQuery("status", "")

	offset := (page - 1) * pageSize

	var visits []models.Visit
	var total int64

	query := database.DB.Model(&models.Visit{}).Preload("Elder")
	if statusStr != "" {
		status, _ := strconv.Atoi(statusStr)
		query = query.Where("status = ?", status)
	}

	query.Count(&total)

	if err := query.Order("visit_time DESC").Offset(offset).Limit(pageSize).Find(&visits).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "获取来访登记列表失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data": gin.H{
			"list":      visits,
			"total":     total,
			"page":      page,
			"page_size": pageSize,
		},
	})
}

func GetVisit(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var visit models.Visit
	if err := database.DB.Preload("Elder").First(&visit, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "来访登记不存在",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data":    visit,
	})
}

func CreateVisit(c *gin.Context) {
	var req CreateVisitRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误: " + err.Error(),
		})
		return
	}

	visit := models.Visit{
		ElderID:     req.ElderID,
		VisitorName: req.VisitorName,
		Relation:    req.Relation,
		Phone:       req.Phone,
		IDCard:      req.IDCard,
		VisitTime:   req.VisitTime,
		Remarks:     req.Remarks,
		Status:      0,
	}

	if err := database.DB.Create(&visit).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "创建来访登记失败",
		})
		return
	}

	database.DB.Preload("Elder").First(&visit, visit.ID)

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "创建成功",
		"data":    visit,
	})
}

func UpdateVisit(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var req UpdateVisitRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var visit models.Visit
	if err := database.DB.First(&visit, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "来访登记不存在",
		})
		return
	}

	updates := make(map[string]interface{})
	if req.LeaveTime != nil {
		updates["leave_time"] = req.LeaveTime
	}
	if req.Status != nil {
		updates["status"] = *req.Status
	}
	if req.Remarks != "" {
		updates["remarks"] = req.Remarks
	}

	if err := database.DB.Model(&visit).Updates(updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "更新来访登记失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "更新成功",
	})
}
