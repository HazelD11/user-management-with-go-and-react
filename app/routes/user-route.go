git package route

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewUserRoute(db *gorm.DB, router *gin.Engine) {
	// userController := controller.NewUserController(db)
	// userRouter := router.Group("/user")
	// {
	// 	userRouter.POST("/register", userController.Register)
	// 	userRouter.POST("/login", userController.Login)
	// }
}
