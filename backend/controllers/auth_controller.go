package controllers

import (
	"net/http"

	"retirementManage/database"
	"retirementManage/models"
	"retirementManage/utils"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var user models.User
	if err := database.DB.Preload("Role").Where("username = ?", req.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "用户名或密码错误",
		})
		return
	}

	if !user.CheckPassword(req.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "用户名或密码错误",
		})
		return
	}

	if user.Status != 1 {
		c.JSON(http.StatusForbidden, gin.H{
			"code":    403,
			"message": "账户已被禁用",
		})
		return
	}

	token, err := utils.GenerateToken(user.ID, user.Username, user.RoleID, user.Role.Code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "生成token失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"message": "登录成功",
		"data": gin.H{
			"token": token,
			"user": gin.H{
				"id":         user.ID,
				"username":   user.Username,
				"real_name":  user.RealName,
				"phone":      user.Phone,
				"email":      user.Email,
				"role_id":    user.RoleID,
				"role_name":  user.Role.Name,
				"role_code":  user.Role.Code,
			},
		},
	})
}

func GetUserInfo(c *gin.Context) {
	userID := c.GetUint("user_id")

	var user models.User
	if err := database.DB.Preload("Role").First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "用户不存在",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data": gin.H{
			"id":         user.ID,
			"username":   user.Username,
			"real_name":  user.RealName,
			"phone":      user.Phone,
			"email":      user.Email,
			"role_id":    user.RoleID,
			"role_name":  user.Role.Name,
			"role_code":  user.Role.Code,
		},
	})
}

func Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "退出登录成功",
	})
}
