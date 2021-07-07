package main

import (
	"PAAS-TA-PORTAL-V3/docs"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func main() {
	//// programmatically set swagger info
	docs.SwaggerInfo.Title = "PaaS-TA Portal CF V3 API"
	docs.SwaggerInfo.Description = "This is a PaaS-TA Portal Server."
	docs.SwaggerInfo.Version = "2.0"
	docs.SwaggerInfo.BasePath = "/v3"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	docs.SwaggerInfo.Host = "localhost:2222"
	r := gin.New()

	// use ginSwagger middleware to serve the API docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()
}
