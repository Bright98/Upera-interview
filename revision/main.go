package product

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"revision/api"
)

var (
	PORT        string
	Gin         *gin.Engine
	RestHandler *api.RestHandler
)

func main() {
	//define routes

	//run gin
	err := Gin.Run(PORT)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println("-> Server running on ", PORT)
}
