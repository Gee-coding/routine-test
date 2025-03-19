package controllers

import (
	"context"

	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// ตั้งค่า OAuth2 Config
var googleOAuthConfig = &oauth2.Config{
	ClientID:     "YOUR_CLIENT_ID",
	ClientSecret: "YOUR_CLIENT_SECRET",
	RedirectURL:  "http://localhost:8080/auth/callback",
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile"},
	Endpoint:     google.Endpoint,
}

// Middleware ตรวจสอบ Token
func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.QueryParam("token") // รับ Token จาก Query Parameter
		if token == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
		}
		c.Set("token", token) // เก็บ Token ไว้ใช้ใน Handler
		return next(c)
	}
}

// ฟังก์ชัน Login (Redirect ไป Google OAuth2)
func HandleGoogleLogin(c echo.Context) error {
	url := googleOAuthConfig.AuthCodeURL("random-state-token", oauth2.AccessTypeOffline)
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

// ฟังก์ชัน Callback จาก Google
func HandleGoogleCallback(c echo.Context) error {
	code := c.QueryParam("code")
	if code == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "No code found"})
	}

	// แลกเปลี่ยน Code เป็น Access Token
	token, err := googleOAuthConfig.Exchange(context.Background(), code)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to exchange token"})
	}

	// ดึงข้อมูล User จาก Google API
	client := googleOAuthConfig.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to get user info"})
	}
	defer resp.Body.Close()

	// ✅ ส่ง Token ไปยัง Client เพื่อใช้งาน
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Login Successful",
		"token":   token.AccessToken,
	})
}

// ฟังก์ชันที่ต้องใช้ Authentication
func HandleProfile(c echo.Context) error {
	token := c.Get("token").(string)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Protected Profile",
		"token":   token,
	})
}
