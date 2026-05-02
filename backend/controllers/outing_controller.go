package controllers

import (
	"net/http"
	"strconv"
	"time"

	"retirementManage/database"
	"retirementManage/models"

	"github.com/gin-gonic/gin"
)

type CreateOutingRequest struct {
	ElderID    uint       `json:"elder_id" binding:"required"`
	OutTime    *time.Time `json:"out_time"`
	Companion  string     `json:"companion"`
	Relation   string     `json:"relation"`
	Phone      string     `json:"phone"`
	Purpose    string     `json:"purpose"`
	Remarks    string     `json:"remarks"`
}

type UpdateOutingRequest struct {
	ReturnTime *time.Time `json:"return_time"`
	Status     *int       `json:"status"`
	Remarks    string     `json:"remarks"`
}

func GetOutings(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	statusStr := c.DefaultQuery("status", "")

	offset := (page - 1) * pageSize

	var outings []models.Outing
	var total int64

	query := database.DB.Model(&models.Outing{}).Preload("Elder")
	if statusStr != "" {
		status, _ := strconv.Atoi(statusStr)
		query = query.Where("status = ?", status)
	}

	query.Count(&total)

	if err := query.Order("out_time DESC").Offset(offset).Limit(pageSize).Find(&outings).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "获取外出登记列表失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data": gin.H{
			"list":      outings,
			"total":     total,
			"page":      page,
			"page_size": pageSize,
		},
	})
}

func GetOuting(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var outing models.Outing
	if err := database.DB.Preload("Elder").First(&outing, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "外出登记不存在",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data":    outing,
	})
}

func CreateOuting(c *gin.Context) {
	var req CreateOutingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误: " + err.Error(),
		})
		return
	}

	outing := models.Outing{
		ElderID:   req.ElderID,
		OutTime:   req.OutTime,
		Companion: req.Companion,
		Relation:  req.Relation,
		Phone:     req.Phone,
		Purpose:   req.Purpose,
		Remarks:   req.Remarks,
		Status:    0,
	}

	if err := database.DB.Create(&outing).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "创建外出登记失败",
		})
		return
	}

	database.DB.Preload("Elder").First(&outing, outing.ID)

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "创建成功",
		"data":    outing,
	})
}

func UpdateOuting(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var req UpdateOutingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var outing models.Outing
	if err := database.DB.First(&outing, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "外出登记不存在",
		})
		return
	}

	updates := make(map[string]interface{})
	if req.ReturnTime != nil {
		updates["return_time"] = req.ReturnTime
	}
	if req.Status != nil {
		updates["status"] = *req.Status
	}
	if req.Remarks != "" {
		updates["remarks"] = req.Remarks
	}

	if err := database.DB.Model(&outing).Updates(updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "更新外出登记失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "更新成功",
	})
}
