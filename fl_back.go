package main

import (
    "fmt"
    "log"
    "net/http"
    "strings"

    "github.com/gin-gonic/contrib/sessions"
    "github.com/gin-gonic/gin"

    "github.com/FantLab/go_backend/API"
)

var user API.User

func main() {
    r := gin.Default()

    store := sessions.NewCookieStore([]byte("cp8y3c58942ych589"))
    r.Use(sessions.Sessions("mysession", store))

    r.POST("/login", login)
    r.GET("/logout", logout)

    private := r.Group("/private")
    {
        private.GET("/", private1)
        private.GET("/two", private2)
    }

    private.Use(AuthRequired())

    fmt.Println("Starting")

    r.Run(":8080")
}


func AuthRequired() gin.HandlerFunc {
    return func(c *gin.Context) {
        session := sessions.Default(c)
        user := session.Get("user")
        if user == nil {
            // You'd normally redirect to login page
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
        } else {
            // Continue down the chain to handler etc
            c.Next()
        }
    }
}

func login(c *gin.Context) {
    session := sessions.Default(c)
    var user API.User

    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if strings.Trim(user.Login, " ") == "" || strings.Trim(user.Pass, " ") == "" {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Parameters can't be empty"})
        return
    }

    if user.Login == "hello" && user.Pass == "123" {
        session.Set("user", user.Login) //In real world usage you'd set this to the users ID
        err := session.Save()

        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate session token"})
        } else {
            c.JSON(http.StatusOK, gin.H{"message": "Successfully authenticated user"})
        }
    } else {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
    }
}

func logout(c *gin.Context) {
    session := sessions.Default(c)
    user := session.Get("user")
    if user == nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
    } else {
        log.Println(user)
        session.Delete("user")
        session.Save()
        c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
    }
}

func private1(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"hello": user})
}

func private2(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"hello": "Logged in user"})
}
