package main

import (
	"fmt"

	_ "github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	// Create a new instance of Echo
	e := echo.New()

	e.Use(middleware.CORS())
	// ✅ เพิ่ม Global Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	_publicAPI := e.Group("/api/v1")

	fmt.Println(_publicAPI)

	// Start Server
	e.Logger.Fatal(e.Start(":8080"))

}
