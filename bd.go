package main

import (
    "log"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
//    "github.com/jinzhu/gorm/dialects/postgres"
    "github.com/FantLab/go_backend/bd"
)

var db *gorm.DB

func init() {
    //open a db connection
    db, err := gorm.Open("mysql", "root:desarito@tcp(127.0.0.1:3307)/fantlab?charset=utf8&parseTime=True&loc=Local")
    if err != nil {
        log.Fatalf("failed to connect database %v", err)
    }
    defer db.Close()
    db.LogMode(true)

    //Migrate the schema
    db.AutoMigrate(&bd.FLUser{})
}