package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"product/received/api"
)

var (
	PORT        string
	Gin         *gin.Engine
	RestHandler *api.RestHandler
)

func main() {
	//define routes
	Gin.POST("/api/product/products", RestHandler.InsertProduct)
	Gin.PUT("/api/product/products/id/:product-id", RestHandler.UpdateProduct)
	Gin.GET("/api/product/products/id/:product-id", RestHandler.GetProductByID)
	Gin.GET("/api/product/products", RestHandler.GetAllProducts)

	//run gin
	err := Gin.Run(PORT)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println("-> Server running on ", PORT)
}
