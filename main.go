package main

import (
	"routine-test/controllers"

	_ "github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	// API_KEY := os.Getenv("API_KEY")
	// if API_KEY == "" {
	// 	log.Fatal("API_KEY is not set")
	// }

	// Create a new instance of Echo
	e := echo.New()

	e.Use(middleware.CORS())
	// ✅ เพิ่ม Global Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// // สร้างกลุ่ม public API
	// _publicAPI := e.Group("/api/v1")

	// // สร้างกลุ่ม private API
	// config := middleware.JWTConfig{
	// 	Claims:     &models.JwtCustomClaims{},
	// 	SigningKey: []byte(API_KEY),
	// }
	// _privateAPI := e.Group("/api/v1")
	// _privateAPI.Use(middleware.JWTWithConfig(config))

	// ✅ เส้นทางสำหรับ Login
	e.GET("/auth/login", controllers.HandleGoogleLogin)
	e.GET("/auth/callback", controllers.HandleGoogleCallback)

	// ✅ เส้นทางที่ต้องการ Authentication
	private := e.Group("/profile")
	private.Use(controllers.AuthMiddleware) // ใช้ Middleware ตรวจสอบ Token
	private.GET("", controllers.HandleProfile)

	// Start Server
	e.Logger.Fatal(e.Start(":8080"))

}
