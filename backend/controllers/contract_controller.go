package controllers

import (
	"net/http"
	"strconv"
	"time"

	"retirementManage/database"
	"retirementManage/models"

	"github.com/gin-gonic/gin"
)

type CreateContractRequest struct {
	ContractNo string     `json:"contract_no"`
	ElderID    uint       `json:"elder_id" binding:"required"`
	BedID      uint       `json:"bed_id" binding:"required"`
	StartDate  *time.Time `json:"start_date"`
	EndDate    *time.Time `json:"end_date"`
	Amount     float64    `json:"amount"`
	Deposit    float64    `json:"deposit"`
	Remarks    string     `json:"remarks"`
}

type UpdateContractRequest struct {
	StartDate *time.Time `json:"start_date"`
	EndDate   *time.Time `json:"end_date"`
	Amount    float64    `json:"amount"`
	Deposit   float64    `json:"deposit"`
	Status    *int       `json:"status"`
	Remarks   string     `json:"remarks"`
}

func GetContracts(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	contractNo := c.DefaultQuery("contract_no", "")
	statusStr := c.DefaultQuery("status", "")

	offset := (page - 1) * pageSize

	var contracts []models.Contract
	var total int64

	query := database.DB.Model(&models.Contract{}).Preload("Elder").Preload("Bed.Room.RoomType")
	if contractNo != "" {
		query = query.Where("contract_no LIKE ?", "%"+contractNo+"%")
	}
	if statusStr != "" {
		status, _ := strconv.Atoi(statusStr)
		query = query.Where("status = ?", status)
	}

	query.Count(&total)

	if err := query.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&contracts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "获取合同列表失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data": gin.H{
			"list":      contracts,
			"total":     total,
			"page":      page,
			"page_size": pageSize,
		},
	})
}

func GetContract(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var contract models.Contract
	if err := database.DB.Preload("Elder").Preload("Bed.Room.RoomType").First(&contract, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "合同不存在",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data":    contract,
	})
}

func CreateContract(c *gin.Context) {
	var req CreateContractRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误: " + err.Error(),
		})
		return
	}

	now := time.Now()
	contractNo := "HT" + now.Format("20060102") + "0001"

	if req.ContractNo != "" {
		contractNo = req.ContractNo
	}

	contract := models.Contract{
		ContractNo: contractNo,
		ElderID:    req.ElderID,
		BedID:      req.BedID,
		StartDate:  req.StartDate,
		EndDate:    req.EndDate,
		Amount:     req.Amount,
		Deposit:    req.Deposit,
		Remarks:    req.Remarks,
		Status:     0,
	}

	if err := database.DB.Create(&contract).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "创建合同失败",
		})
		return
	}

	database.DB.Preload("Elder").Preload("Bed.Room.RoomType").First(&contract, contract.ID)

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "创建成功",
		"data":    contract,
	})
}

func UpdateContract(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var req UpdateContractRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var contract models.Contract
	if err := database.DB.First(&contract, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "合同不存在",
		})
		return
	}

	updates := make(map[string]interface{})
	if req.StartDate != nil {
		updates["start_date"] = req.StartDate
	}
	if req.EndDate != nil {
		updates["end_date"] = req.EndDate
	}
	if req.Amount >= 0 {
		updates["amount"] = req.Amount
	}
	if req.Deposit >= 0 {
		updates["deposit"] = req.Deposit
	}
	if req.Status != nil {
		updates["status"] = *req.Status
	}
	if req.Remarks != "" {
		updates["remarks"] = req.Remarks
	}

	if err := database.DB.Model(&contract).Updates(updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "更新合同失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "更新成功",
	})
}
