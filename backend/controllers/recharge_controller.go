package controllers

import (
	"net/http"
	"strconv"
	"time"

	"retirementManage/database"
	"retirementManage/models"

	"github.com/gin-gonic/gin"
)

type CreateRechargeRequest struct {
	ElderID     uint    `json:"elder_id" binding:"required"`
	Amount      float64 `json:"amount" binding:"required"`
	GiftAmount  float64 `json:"gift_amount"`
	PayMethod   int     `json:"pay_method" binding:"required"`
	Remarks     string  `json:"remarks"`
}

func GetRecharges(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	rechargeNo := c.DefaultQuery("recharge_no", "")
	elderName := c.DefaultQuery("elder_name", "")
	payMethodStr := c.DefaultQuery("pay_method", "")
	statusStr := c.DefaultQuery("status", "")

	offset := (page - 1) * pageSize

	var recharges []models.Recharge
	var total int64

	query := database.DB.Model(&models.Recharge{}).Preload("Elder")
	if rechargeNo != "" {
		query = query.Where("recharge_no LIKE ?", "%"+rechargeNo+"%")
	}
	if payMethodStr != "" {
		payMethod, _ := strconv.Atoi(payMethodStr)
		query = query.Where("pay_method = ?", payMethod)
	}
	if statusStr != "" {
		status, _ := strconv.Atoi(statusStr)
		query = query.Where("status = ?", status)
	}

	query.Count(&total)

	if err := query.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&recharges).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "获取充值记录失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data": gin.H{
			"list":      recharges,
			"total":     total,
			"page":      page,
			"page_size": pageSize,
		},
	})
}

func GetRecharge(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var recharge models.Recharge
	if err := database.DB.Preload("Elder").First(&recharge, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "充值记录不存在",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data":    recharge,
	})
}

func CreateRecharge(c *gin.Context) {
	var req CreateRechargeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误: " + err.Error(),
		})
		return
	}

	now := time.Now()
	rechargeNo := "RC" + now.Format("20060102") + "0001"

	recharge := models.Recharge{
		RechargeNo:  rechargeNo,
		ElderID:     req.ElderID,
		Amount:      req.Amount,
		GiftAmount:  req.GiftAmount,
		TotalAmount: req.Amount + req.GiftAmount,
		PayMethod:   req.PayMethod,
		Status:      0,
		Remarks:     req.Remarks,
	}

	if err := database.DB.Create(&recharge).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "创建充值记录失败",
		})
		return
	}

	database.DB.Preload("Elder").First(&recharge, recharge.ID)

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "创建成功",
		"data":    recharge,
	})
}

func ConfirmRecharge(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var recharge models.Recharge
	if err := database.DB.First(&recharge, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "充值记录不存在",
		})
		return
	}

	if recharge.Status != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "该充值订单状态不正确",
		})
		return
	}

	now := time.Now()
	if err := database.DB.Model(&recharge).Updates(map[string]interface{}{
		"status":   1,
		"pay_time": now,
	}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "确认充值失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "确认支付成功",
	})
}

func CancelRecharge(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var recharge models.Recharge
	if err := database.DB.First(&recharge, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "充值记录不存在",
		})
		return
	}

	if recharge.Status != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "该充值订单状态不正确",
		})
		return
	}

	if err := database.DB.Model(&recharge).Update("status", 2).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "取消充值失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "已取消订单",
	})
}
