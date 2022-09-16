package models

type Person struct {
	ID        uint `gorm:"primaryKey"`
	FirstName string
	LastName  string
	Age       uint64
}
