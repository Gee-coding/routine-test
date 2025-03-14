package models

import (
	"time"
)

type User struct {
	ID        uint           `gorm:"primaryKey"`
	Username  string         `gorm:"unique;not null"`
	Email     string         `gorm:"unique;not null"`
	Password  string         `gorm:"not null"`
	Role      string         `gorm:"type:enum('admin', 'employee', 'patient');not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}


type Employee struct {
	ID         uint   `gorm:"primaryKey"`
	UserID     uint   `gorm:"uniqueIndex"` // เชื่อมกับ User
	FirstName  string `gorm:"not null"`
	LastName   string `gorm:"not null"`
	Position   string `gorm:"not null"` // ตำแหน่ง เช่น หมอ, พยาบาล
	Department string // แผนก เช่น ศัลยกรรม, อายุรกรรม
	Phone      string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Admin struct {
	ID       uint   `gorm:"primaryKey"`
	UserID   uint   `gorm:"uniqueIndex"`
	FullName string `gorm:"not null"`
	Role     string `gorm:"not null;default:'admin'"`
}

type Patient struct {
	ID          uint   `gorm:"primaryKey"`
	UserID      uint   `gorm:"uniqueIndex"`
	FirstName   string `gorm:"not null"`
	LastName    string `gorm:"not null"`
	DateOfBirth time.Time
	Gender      string `gorm:"type:enum('male', 'female', 'other');not null"`
	Address     string
	Phone       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}