package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	// loadEnvVariables()
	initGateway()
}

func initGateway() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	port := 9000
	if err := router.Run(fmt.Sprintf(":%d", port)); err != nil {
		log.Fatalf("cannot start api gateway on port %d, error is :%s", port, err.Error())
	}
}


