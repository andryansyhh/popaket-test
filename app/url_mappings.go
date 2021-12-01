package app

import (
	"popaket/auth"
	config "popaket/config"
	"popaket/handler"
	"popaket/service"
	"popaket/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	DB                 *gorm.DB = config.Connection()
	authService                 = auth.NewService()
	userService                 = service.NewUserService(userRepository)
	userRepository              = storage.NewDao(DB)
	userHandler                 = handler.NewUserHandler(userService, authService)
	logisticservice             = service.NewLogisticService(logisticRepository)
	logisticRepository          = storage.NewDao(DB)
	logisticHandler             = handler.NewLogisticHandler(logisticservice)
)

func RegisterApi(r *gin.Engine) {
	api := r.Group("/popaket")
	{
		// user
		api.POST("/users/register", userHandler.RegisterUserHandler)
		api.POST("/users/login", userHandler.LoginUserHandler)

		//product
		api.GET("/logistic", logisticHandler.ShowAllLogistiHandler)
		api.GET("logistic/filter", logisticHandler.ShowAllLogistiByParamHandler)

	}
}
