package main

import {
    "fmt"

    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
    "github.com/jinzhu/gorm/dialects/mysql"
//    "github.com/jinzhu/gorm/dialects/postgres"
}

func main() {
    r := gin.Default()

    db, err := gorm.Open("mysql", "root:desarito@127.0.0.1:3307/fantlab?charset=utf8&parseTime=True&loc=Local")
    defer db.Close()

    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })

    r.Run()
}