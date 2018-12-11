package bd

import "github.com/jinzhu/gorm"

type FLUser struct {
    gorm.Model
    Login   string `gorm:"size:16;unique;not null"`
    Password string `gorm:"size:128;not null"`
    Email    string `gorm:"size:50;not null"`
}
