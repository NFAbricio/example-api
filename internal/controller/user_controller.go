package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/NFAbricio/example-api/users"
)

type UserController struct {
	userService users.Usecase
}

func NewUserController(us users.Usecase) *UserController {
	return &UserController{userService: us}
}

func (uc *UserController) Create(ctx *gin.Context) {
	var requestParams users.User
	//the JSON will be binded in a User struct 
	if err := ctx.BindJSON(&requestParams); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid body content"})
		return
	}

	err := uc.userService.Create(&requestParams)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "user created suscessfully"})
}

func (uc *UserController) Get(ctx *gin.Context) {
	var requestParams users.User
	//the JSON will be binded in a User struct 
	if err := ctx.BindJSON(&requestParams); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid body content"})
		return
	}

	err := uc.userService.Create(&requestParams)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "user created suscessfully"})

}

// func (uc *UserController) Update(ctx *gin.Context) {}

// func (uc *UserController) Delete(ctx *gin.Context) {}

// func (uc *UserController) Auth(ctx *gin.Context) {}
