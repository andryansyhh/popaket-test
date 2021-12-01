package main

import (
	"popaket/app"
	"popaket/docs"
)

func main() {

	docs.SwaggerInfo.Title = "Popaket API documentation"
	docs.SwaggerInfo.Description = "Popaket API documentation"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:1231"
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	app.StartApplication()
}
