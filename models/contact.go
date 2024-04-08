package models

import(
    "gorm.io/gorm"
)

type Contact struct {
    gorm.Model
    Name  string `json:"name" binding:"required"`
    Phone string `json:"phone" binding:"required"`
    Email string `json:"email"`
    Address  string `json:"address"`
}
