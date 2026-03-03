package controllers

import (
    "net/http"

    "github.com/gin-gonic/gin"

    "test/auth"
    "test/database"
    "test/middlewares"
)

type AuthRequest struct {
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required,min=6"`
}

func Register(c *gin.Context) {
    var req AuthRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := auth.Register(database.DB, req.Email, req.Password); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "registered successfully"})
}

func Login(c *gin.Context) {
    var req AuthRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user, err := auth.Login(database.DB, req.Email, req.Password)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
    }

    token, err := middlewares.GenerateToken(user.ID)
    if err != nil {
    c.JSON(500, gin.H{
        "error": err.Error(),
    })
    return
}


    c.JSON(http.StatusOK, gin.H{
        "token": token,
    })
}
