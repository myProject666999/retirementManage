package controllers

import (
	"net/http"
	"strconv"
	"time"

	"retirementManage/database"
	"retirementManage/models"

	"github.com/gin-gonic/gin"
)

func GetCheckouts(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	elderName := c.DefaultQuery("elder_name", "")
	statusStr := c.DefaultQuery("status", "")

	offset := (page - 1) * pageSize

	var checkouts []models.Checkout
	var total int64

	query := database.DB.Model(&models.Checkout{}).Preload("Elder").Preload("Contract")
	if statusStr != "" {
		status, _ := strconv.Atoi(statusStr)
		query = query.Where("status = ?", status)
	}

	query.Count(&total)

	if err := query.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&checkouts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "获取退住申请列表失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data": gin.H{
			"list":      checkouts,
			"total":     total,
			"page":      page,
			"page_size": pageSize,
		},
	})
}

func GetCheckout(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var checkout models.Checkout
	if err := database.DB.Preload("Elder").Preload("Contract").First(&checkout, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "退住申请不存在",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data":    checkout,
	})
}

func ApproveCheckout(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var checkout models.Checkout
	if err := database.DB.First(&checkout, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "退住申请不存在",
		})
		return
	}

	if checkout.Status != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "该申请状态不正确",
		})
		return
	}

	now := time.Now()
	if err := database.DB.Model(&checkout).Updates(map[string]interface{}{
		"status":      1,
		"audit_time":  now,
	}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "审核失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "审核通过",
	})
}

func RejectCheckout(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var req struct {
		AuditRemark string `json:"audit_remark"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var checkout models.Checkout
	if err := database.DB.First(&checkout, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "退住申请不存在",
		})
		return
	}

	if checkout.Status != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "该申请状态不正确",
		})
		return
	}

	now := time.Now()
	if err := database.DB.Model(&checkout).Updates(map[string]interface{}{
		"status":       2,
		"audit_time":   now,
		"audit_remark": req.AuditRemark,
	}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "审核失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "已拒绝申请",
	})
}
