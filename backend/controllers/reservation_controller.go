package controllers

import (
	"net/http"
	"strconv"
	"time"

	"retirementManage/database"
	"retirementManage/models"

	"github.com/gin-gonic/gin"
)

type CreateReservationRequest struct {
	Name            string     `json:"name" binding:"required"`
	Phone           string     `json:"phone" binding:"required"`
	Gender          int        `json:"gender"`
	Age             int        `json:"age"`
	IDCard          string     `json:"id_card"`
	RoomTypeID      uint       `json:"room_type_id"`
	ReservationDate *time.Time `json:"reservation_date"`
	DepositAmount   float64    `json:"deposit_amount"`
	Remarks         string     `json:"remarks"`
}

type UpdateReservationRequest struct {
	Name            string     `json:"name"`
	Phone           string     `json:"phone"`
	Gender          int        `json:"gender"`
	Age             int        `json:"age"`
	IDCard          string     `json:"id_card"`
	RoomTypeID      uint       `json:"room_type_id"`
	ReservationDate *time.Time `json:"reservation_date"`
	DepositAmount   float64    `json:"deposit_amount"`
	PaidAmount      float64    `json:"paid_amount"`
	Remarks         string     `json:"remarks"`
	Status          *int       `json:"status"`
}

func GetReservations(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	name := c.DefaultQuery("name", "")
	phone := c.DefaultQuery("phone", "")
	statusStr := c.DefaultQuery("status", "")

	offset := (page - 1) * pageSize

	var reservations []models.Reservation
	var total int64

	query := database.DB.Model(&models.Reservation{}).Preload("RoomType")
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

	query.Count(&total)

	if err := query.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&reservations).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "获取预定列表失败",
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

func GetReservation(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var reservation models.Reservation
	if err := database.DB.Preload("RoomType").First(&reservation, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "预定记录不存在",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data":    reservation,
	})
}

func CreateReservation(c *gin.Context) {
	var req CreateReservationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误: " + err.Error(),
		})
		return
	}

	reservation := models.Reservation{
		Name:            req.Name,
		Phone:           req.Phone,
		Gender:          req.Gender,
		Age:             req.Age,
		IDCard:          req.IDCard,
		RoomTypeID:      req.RoomTypeID,
		ReservationDate: req.ReservationDate,
		DepositAmount:   req.DepositAmount,
		Remarks:         req.Remarks,
		Status:          0,
	}

	if err := database.DB.Create(&reservation).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "创建预定记录失败",
		})
		return
	}

	database.DB.Preload("RoomType").First(&reservation, reservation.ID)

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "创建成功",
		"data":    reservation,
	})
}

func UpdateReservation(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var req UpdateReservationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var reservation models.Reservation
	if err := database.DB.First(&reservation, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "预定记录不存在",
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
	if req.IDCard != "" {
		updates["id_card"] = req.IDCard
	}
	if req.RoomTypeID > 0 {
		updates["room_type_id"] = req.RoomTypeID
	}
	if req.ReservationDate != nil {
		updates["reservation_date"] = req.ReservationDate
	}
	if req.DepositAmount >= 0 {
		updates["deposit_amount"] = req.DepositAmount
	}
	if req.PaidAmount >= 0 {
		updates["paid_amount"] = req.PaidAmount
	}
	if req.Remarks != "" {
		updates["remarks"] = req.Remarks
	}
	if req.Status != nil {
		updates["status"] = *req.Status
	}

	if err := database.DB.Model(&reservation).Updates(updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "更新预定记录失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "更新成功",
	})
}

func DeleteReservation(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var reservation models.Reservation
	if err := database.DB.First(&reservation, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "预定记录不存在",
		})
		return
	}

	if err := database.DB.Delete(&reservation).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "删除预定记录失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "删除成功",
	})
}
