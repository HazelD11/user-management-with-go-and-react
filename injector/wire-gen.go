package injector

import (
	"github.com/AustinNick/coba-golang/controller"
	"github.com/AustinNick/coba-golang/repository"
	"github.com/AustinNick/coba-golang/service"
	"gorm.io/gorm"
)

func InitUser(db *gorm.DB) controller.UserController {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)
	return userController
}
