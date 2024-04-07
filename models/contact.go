package models

type Contact struct {
    ID    uint   `json:"id" gorm:"primary_key"`
    Name  string `json:"name" binding:"required"`
    Phone string `json:"phone" binding:"required"`
    Email string `json:"email"`
}
