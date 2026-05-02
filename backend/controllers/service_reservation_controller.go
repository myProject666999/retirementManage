package controllers

import (
	"net/http"
	"strconv"
	"time"

	"retirementManage/database"
	"retirementManage/models"

	"github.com/gin-gonic/gin"
)

type CreateServiceReservationRequest struct {
	ElderID       uint       `json:"elder_id" binding:"required"`
	ServiceItemID uint      `json:"service_item_id" binding:"required"`
	ServiceTime   *time.Time `json:"service_time"`
	ServicePlace  string     `json:"service_place"`
	Quantity      int        `json:"quantity"`
	Amount        float64    `json:"amount"`
	Remarks       string     `json:"remarks"`
}

type UpdateServiceReservationRequest struct {
	ServiceTime   *time.Time `json:"service_time"`
	ServicePlace  string     `json:"service_place"`
	Quantity      int        `json:"quantity"`
	Amount        float64    `json:"amount"`
	Status        *int       `json:"status"`
	ServiceBy     uint       `json:"service_by"`
	ServiceResult string     `json:"service_result"`
	Remarks       string     `json:"remarks"`
}

func GetServiceReservations(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	statusStr := c.DefaultQuery("status", "")
	elderIDStr := c.DefaultQuery("elder_id", "")

	offset := (page - 1) * pageSize

	var reservations []models.ServiceReservation
	var total int64

	query := database.DB.Model(&models.ServiceReservation{}).Preload("Elder").Preload("ServiceItem")
	if statusStr != "" {
		status, _ := strconv.Atoi(statusStr)
		query = query.Where("status = ?", status)
	}
	if elderIDStr != "" {
		elderID, _ := strconv.Atoi(elderIDStr)
		query = query.Where("elder_id = ?", elderID)
	}

	query.Count(&total)

	if err := query.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&reservations).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "获取服务预定列表失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data": gin.H{
			"list":      reservations,
			"total":     total,
			"page":      page,
			"page_size": pageSize,
		},
	})
}

func GetServiceReservation(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var reservation models.ServiceReservation
	if err := database.DB.Preload("Elder").Preload("ServiceItem").First(&reservation, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "服务预定不存在",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data":    reservation,
	})
}

func CreateServiceReservation(c *gin.Context) {
	var req CreateServiceReservationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误: " + err.Error(),
		})
		return
	}

	reservation := models.ServiceReservation{
		ElderID:       req.ElderID,
		ServiceItemID: req.ServiceItemID,
		ServiceTime:   req.ServiceTime,
		ServicePlace:  req.ServicePlace,
		Quantity:      req.Quantity,
		Amount:        req.Amount,
		Remarks:       req.Remarks,
		Status:        0,
	}

	if err := database.DB.Create(&reservation).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "创建服务预定失败",
		})
		return
	}

	database.DB.Preload("Elder").Preload("ServiceItem").First(&reservation, reservation.ID)

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "创建成功",
		"data":    reservation,
	})
}

func UpdateServiceReservation(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var req UpdateServiceReservationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var reservation models.ServiceReservation
	if err := database.DB.First(&reservation, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "服务预定不存在",
		})
		return
	}

	updates := make(map[string]interface{})
	if req.ServiceTime != nil {
		updates["service_time"] = req.ServiceTime
	}
	if req.ServicePlace != "" {
		updates["service_place"] = req.ServicePlace
	}
	if req.Quantity > 0 {
		updates["quantity"] = req.Quantity
	}
	if req.Amount >= 0 {
		updates["amount"] = req.Amount
	}
	if req.Status != nil {
		updates["status"] = *req.Status
	}
	if req.ServiceBy > 0 {
		updates["service_by"] = req.ServiceBy
	}
	if req.ServiceResult != "" {
		updates["service_result"] = req.ServiceResult
	}
	if req.Remarks != "" {
		updates["remarks"] = req.Remarks
	}

	if err := database.DB.Model(&reservation).Updates(updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "更新服务预定失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "更新成功",
	})
}

func DeleteServiceReservation(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var reservation models.ServiceReservation
	if err := database.DB.First(&reservation, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "服务预定不存在",
		})
		return
	}

	if err := database.DB.Delete(&reservation).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "删除服务预定失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "删除成功",
	})
}
