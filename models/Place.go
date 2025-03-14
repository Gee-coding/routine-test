package models

type Company struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"not null"`
	Address  string
	Phone    string
	Email    string `gorm:"unique"`
	Employees []Employee `gorm:"foreignKey:CompanyID"`
}

type Room struct {
	ID       uint   `gorm:"primaryKey"`
	RoomType string `gorm:"type:enum('ICU', 'General', 'Surgery', 'Maternity');not null"`
	Capacity int
	WardID   uint `gorm:"index"`
}

type Ward struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	Location  string
	Rooms     []Room `gorm:"foreignKey:WardID"`
}