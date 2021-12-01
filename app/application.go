package app

import (
	"fmt"
	"os"

	"popaket/handler"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

var (
	router = gin.Default()
)

func StartApplication() {
	router.Use(handler.CORSMiddleware())
	RegisterApi(router)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := os.Getenv("APP_PORT")
	router.Run(fmt.Sprintf(":%s", port))

}
