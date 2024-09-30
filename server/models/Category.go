package models

type Category struct {
	CategoryName        string `gorm:"type:varchar(255);not null"`
	CategoryDescription string `gorm:"type:varchar(255);not null"`
}
