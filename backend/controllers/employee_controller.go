package controllers

import (
	"net/http"
	"strconv"
	"time"

	"retirementManage/database"
	"retirementManage/models"

	"github.com/gin-gonic/gin"
)

type CreateEmployeeRequest struct {
	Name         string     `json:"name" binding:"required"`
	Gender       int        `json:"gender"`
	BirthDate    *time.Time `json:"birth_date"`
	IDCard       string     `json:"id_card" binding:"required"`
	Phone        string     `json:"phone"`
	Email        string     `json:"email"`
	DepartmentID uint       `json:"department_id"`
	Position     string     `json:"position"`
	EntryDate    *time.Time `json:"entry_date"`
	Remarks      string     `json:"remarks"`
}

type UpdateEmployeeRequest struct {
	Name         string     `json:"name"`
	Gender       int        `json:"gender"`
	BirthDate    *time.Time `json:"birth_date"`
	Phone        string     `json:"phone"`
	Email        string     `json:"email"`
	DepartmentID uint       `json:"department_id"`
	Position     string     `json:"position"`
	EntryDate    *time.Time `json:"entry_date"`
	LeaveDate    *time.Time `json:"leave_date"`
	Status       *int       `json:"status"`
	Remarks      string     `json:"remarks"`
}

func GetEmployees(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	name := c.DefaultQuery("name", "")
	phone := c.DefaultQuery("phone", "")
	statusStr := c.DefaultQuery("status", "")

	offset := (page - 1) * pageSize

	var employees []models.Employee
	var total int64

	query := database.DB.Model(&models.Employee{}).Preload("Department")
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

	if err := query.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&employees).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "获取员工列表失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data": gin.H{
			"list":      employees,
			"total":     total,
			"page":      page,
			"page_size": pageSize,
		},
	})
}

func GetEmployee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var employee models.Employee
	if err := database.DB.Preload("Department").First(&employee, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "员工不存在",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data":    employee,
	})
}

func CreateEmployee(c *gin.Context) {
	var req CreateEmployeeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误: " + err.Error(),
		})
		return
	}

	var existingEmployee models.Employee
	if database.DB.Where("id_card = ?", req.IDCard).First(&existingEmployee).Error == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "身份证号已存在",
		})
		return
	}

	employee := models.Employee{
		Name:         req.Name,
		Gender:       req.Gender,
		BirthDate:    req.BirthDate,
		IDCard:       req.IDCard,
		Phone:        req.Phone,
		Email:        req.Email,
		DepartmentID: req.DepartmentID,
		Position:     req.Position,
		EntryDate:    req.EntryDate,
		Remarks:      req.Remarks,
		Status:       1,
	}

	if err := database.DB.Create(&employee).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "创建员工失败",
		})
		return
	}

	database.DB.Preload("Department").First(&employee, employee.ID)

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "创建成功",
		"data":    employee,
	})
}

func UpdateEmployee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var req UpdateEmployeeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var employee models.Employee
	if err := database.DB.First(&employee, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "员工不存在",
		})
		return
	}

	updates := make(map[string]interface{})
	if req.Name != "" {
		updates["name"] = req.Name
	}
	if req.Gender != 0 {
		updates["gender"] = req.Gender
	}
	if req.BirthDate != nil {
		updates["birth_date"] = req.BirthDate
	}
	if req.Phone != "" {
		updates["phone"] = req.Phone
	}
	if req.Email != "" {
		updates["email"] = req.Email
	}
	if req.DepartmentID > 0 {
		updates["department_id"] = req.DepartmentID
	}
	if req.Position != "" {
		updates["position"] = req.Position
	}
	if req.EntryDate != nil {
		updates["entry_date"] = req.EntryDate
	}
	if req.LeaveDate != nil {
		updates["leave_date"] = req.LeaveDate
	}
	if req.Status != nil {
		updates["status"] = *req.Status
	}
	if req.Remarks != "" {
		updates["remarks"] = req.Remarks
	}

	if err := database.DB.Model(&employee).Updates(updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "更新员工失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "更新成功",
	})
}

func DeleteEmployee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var employee models.Employee
	if err := database.DB.First(&employee, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "员工不存在",
		})
		return
	}

	if err := database.DB.Delete(&employee).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "删除员工失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "删除成功",
	})
}
