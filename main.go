package main

import (
	"routine-test/controllers"

	_ "github.com/joho/godotenv"
)

func main() {

	controllers.Middleware()

}
