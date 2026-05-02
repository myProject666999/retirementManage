package utils

import (
	"time"

	"retirementManage/config"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	RoleID   uint   `json:"role_id"`
	RoleCode string `json:"role_code"`
	jwt.StandardClaims
}

func GenerateToken(userID uint, username string, roleID uint, roleCode string) (string, error) {
	secret := []byte(config.AppConfig.JWT.Secret)
	expireTime := time.Now().Add(time.Second * time.Duration(config.AppConfig.JWT.Expire))

	claims := Claims{
		UserID:   userID,
		Username: username,
		RoleID:   roleID,
		RoleCode: roleCode,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "retirement-manage",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

func ParseToken(tokenString string) (*Claims, error) {
	secret := []byte(config.AppConfig.JWT.Secret)

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.NewValidationError("无效的token", jwt.ValidationErrorClaimsInvalid)
}
