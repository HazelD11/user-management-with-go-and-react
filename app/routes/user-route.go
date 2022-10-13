package route

import (
	"github.com/AustinNick/coba-golang/injector"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewUserRoute(db *gorm.DB, route *gin.Engine) {
	userController := injector.InitUser(db)
	userRoute := route.Group("api/user")
	userRoute.GET("/", userController.All)
	userRoute.POST("/", userController.Insert)
}
