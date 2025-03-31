package routes

import (
	"log"
	"net/http"
	"rest-api/models"
	"rest-api/utils"

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

func login(ctx *gin.Context) {
	var loginReq models.LoginRequest
	
	err := ctx.ShouldBindJSON(&loginReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}

	user := models.User{
		Email:    loginReq.Email,
		Password: loginReq.Password,
	}

	err = user.ValidateCredentials()
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.ID, user.Name, user.Email)
	if err != nil {
		log.Println("Error generating token:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not generate token. Try again later."})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
}

func getUser(ctx *gin.Context) {}
func getUsers(ctx *gin.Context) {}

func createUser(ctx *gin.Context) {}

func updateUser(ctx *gin.Context) {}

func deleteUser(ctx *gin.Context) {}