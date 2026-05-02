package controllers

import (
	"net/http"
	"strconv"

	"retirementManage/database"
	"retirementManage/models"

	"github.com/gin-gonic/gin"
)

func GetExpenseRecords(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	recordNo := c.DefaultQuery("record_no", "")
	elderName := c.DefaultQuery("elder_name", "")
	typeStr := c.DefaultQuery("type", "")

	offset := (page - 1) * pageSize

	var records []models.ExpenseRecord
	var total int64

	query := database.DB.Model(&models.ExpenseRecord{}).Preload("Elder")
	if recordNo != "" {
		query = query.Where("record_no LIKE ?", "%"+recordNo+"%")
	}
	if typeStr != "" {
		typeVal, _ := strconv.Atoi(typeStr)
		query = query.Where("type = ?", typeVal)
	}

	query.Count(&total)

	if err := query.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&records).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "获取消费记录失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data": gin.H{
			"list":      records,
			"total":     total,
			"page":      page,
			"page_size": pageSize,
		},
	})
}

func GetExpenseRecord(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var record models.ExpenseRecord
	if err := database.DB.Preload("Elder").First(&record, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "消费记录不存在",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data":    record,
	})
}
