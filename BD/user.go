package main

type User struct {
    gorm.Model
    ID      int
    Login   string `gorm:"size:16"`
    Password string `gorm:size:64"`
}
