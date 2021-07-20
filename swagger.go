package main

import (
	"PAAS-TA-PORTAL-V3/docs"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"net"
)

func main() {
	//// programmatically set swagger info
	iFaces, _ := net.Interfaces()
	// handle err
	docs.SwaggerInfo.Host = "localhost:2222"
	for _, i := range iFaces {
		addrs, _ := i.Addrs()
		for _, addr := range addrs {
			switch v := addr.(type) {
			case *net.IPNet:
				var ip net.IP
				if v.IP.String() != "127.0.0.1" {
					ip = v.IP
					docs.SwaggerInfo.Host = ip.String() + ":2222"
				}
			}
		}
	}
	docs.SwaggerInfo.Title = "PaaS-TA Portal CF V3 API"
	docs.SwaggerInfo.Description = "This is a PaaS-TA Portal Server."
	docs.SwaggerInfo.Version = "3.0"
	docs.SwaggerInfo.BasePath = "/v3"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r := gin.New()
	// use ginSwagger middleware to serve the API docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()
}
