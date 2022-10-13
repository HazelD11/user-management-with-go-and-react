package controller

import (
	"fmt"
	"net/http"

	"github.com/AustinNick/coba-golang/helper"
	"github.com/AustinNick/coba-golang/model/web"
	"github.com/AustinNick/coba-golang/service"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	All(ctx *gin.Context)
	// FindById(ctx *gin.Context)
	// FindByEmail(ctx *gin.Context)
	Insert(ctx *gin.Context)
	// Update(ctx *gin.Context)
	// Delete(ctx *gin.Context)
}

type userController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &userController{
		userService: userService,
	}
}

func (ctrl *userController) All(ctx *gin.Context) {
	var data = ctrl.userService.All()
	webResponse := helper.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Errors: "",
		Data:   data,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

// func FindById(ctx *gin.Context) {
// 	var id = ctx.Param("id")
// 	var user, err = userService.FindById(id)
// 	if err != nil {
// 		ctx.JSON(404, err.Error())
// 	} else {
// 		ctx.JSON(200, user)
// 	}
// }

// func FindByEmail(ctx *gin.Context) {
// 	var email = ctx.Param("email")
// 	var user = userService.FindByEmail(email)
// 	ctx.JSON(200, user)
// }

func (ctrl *userController) Insert(ctx *gin.Context) {
	var userCreateRequest web.UserCreateRequest
	var errMsg string
	if err := ctx.BindJSON(&userCreateRequest); err != nil {
		return
	}
	data, err := ctrl.userService.Create(userCreateRequest)
	if err != nil {
		errMsg = "Failed to create user"
	}
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Errors: errMsg,
		Data:   data,
	}
	ctx.JSON(http.StatusOK, webResponse)
	fmt.Println("ini id:", data.ID)
}

// func Update(ctx *gin.Context) {

// }

// func Delete(ctx *gin.Context) {

// }
