package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("กำลังทำงาน...")
	time.Sleep(2 * time.Second) // จำลองการทำงาน 2 วินาที
	fmt.Println("เสร็จสิ้น!")
	done <- true // ส่งค่าบอกว่าเสร็จแล้ว
}
func main() {

	// Connect to MySQL Database
	// dsn := "user:password@tcp(localhost:3306)/hospital_db?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	log.Fatal("Failed to connect to database:", err)
	// }

	// fmt.Println("Database connected successfully!")

	// // Auto Migrate Tables
	// db.AutoMigrate(&models.User{}, &models.Employee{}, &models.Admin{}, &models.Patient{}, &models.Company{}, &models.Room{}, &models.Ward{})

	// fmt.Println("Database Migrated Successfully!")

	done := make(chan bool) // สร้างช่องทางสื่อสาร
	go worker(done) // เรียกใช้ Goroutine
	time.Sleep(1 * time.Second) 
	fmt.Println("routine 2")

	<-done // รอรับค่าจาก Channel ก่อนจะจบโปรแกรม
}
