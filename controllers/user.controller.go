package controllers

import (
	"example.com/aorts/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService services.UserService
}

func New(userservice services.UserService) UserController {
	return UserController{
		UserService: userservice,
	}
}

func (u *UserController) CreateUser(ctx *gin.Context) {
	ctx.JSON(200, "")
}
func (u *UserController) GetAll(ctx *gin.Context) {
	ctx.JSON(200, "")
}

func (u *UserController) GetUser(ctx *gin.Context) {
	ctx.JSON(200, "")
}

func (u *UserController) UpdateUser(ctx *gin.Context) {
	ctx.JSON(200, "")
}

func (u *UserController) DeleteUser(ctx *gin.Context) {
	ctx.JSON(200, "")
}

func (uc *UserController) RegisterUserRoutes(rg *gin.RouterGroup) {
	userroute := rg.Group("/user")
	userroute.POST("/create", uc.CreateUser)
	userroute.GET("/get/:name", uc.GetUser)
	userroute.GET("/getAll", uc.GetAll)
	userroute.PATCH("/update", uc.UpdateUser)
	userroute.DELETE("/delete", uc.DeleteUser)
}
