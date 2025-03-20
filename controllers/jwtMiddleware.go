package controllers

import (
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	_ "github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

// Custom Middleware JWT
func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// รับค่า JWT Token จาก Header
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Missing token"})
		}

		// ดึง Token ออกจาก "Bearer <token>"
		var tokenString string
		if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
			tokenString = authHeader[7:]
		} else {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token format"})
		}

		// ตรวจสอบ JWT Token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(API_KEY), nil
		})
		if err != nil || !token.Valid {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid or expired token"})
		}

		// ดึงข้อมูลจาก Claims และเก็บใน Context
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Set("username", claims["username"])
		} else {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid claims"})
		}

		return next(c)
	}
}

// ✅ ฟังก์ชันสร้าง JWT Token (สำหรับทดสอบ)
func GenerateJWT() string {
	claims := jwt.MapClaims{
		"username": "testuser",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, _ := token.SignedString([]byte(API_KEY))
	return t
}
