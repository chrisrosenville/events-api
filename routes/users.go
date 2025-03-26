package routes

import (
	"net/http"
	"rest-api/models"

	"github.com/gin-gonic/gin"
)

func signup(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
        return
    }

	err = user.Save()
	if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create user. Try again later."})
        return
    }

	ctx.JSON(http.StatusCreated, gin.H{"message": "User successfully created"})
}

func getUser(ctx *gin.Context) {}
func getUsers(ctx *gin.Context) {}

func createUser(ctx *gin.Context) {}

func updateUser(ctx *gin.Context) {}

func deleteUser(ctx *gin.Context) {}