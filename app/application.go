package app

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"popaket/handler"
)

var (
	router = gin.Default()
)

func StartApplication() {
	router.Use(handler.CORSMiddleware())
	RegisterApi(router)

	port := os.Getenv("APP_PORT")
	router.Run(fmt.Sprintf(":%s", port))

}
