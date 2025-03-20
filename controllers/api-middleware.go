package controllers

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var API_KEY string

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	API_KEY = os.Getenv("API_KEY")
	if API_KEY == "" {
		log.Fatal("API_KEY is not set")
	}
}

func Middleware() {

	// Create a new instance of Echo
	e := echo.New()

	e.Use(middleware.CORS())
	// ✅ เพิ่ม Global Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	_publicAPI := e.Group("/api/v1")

	_privateAPI := e.Group("/api/v1", echojwt.WithConfig(echojwt.Config{ // JWT
		//Claims:     &models.JwtCustomClaims{},
		SigningKey: []byte(API_KEY),
	}))

	_, _ = _privateAPI, _publicAPI

	e.Logger.Fatal(e.Start(":8080"))
}


