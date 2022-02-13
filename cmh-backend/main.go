package main

import (
	"cmh-backend/endpoints"
	"cmh-backend/middlewares"
	"cmh-backend/wss"

	"github.com/gin-gonic/gin"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)
	// gin.DisableConsoleColor()

	// model.ConnectToDB("mongodb://cmh:cmh@65.0.96.94:27017/admin?retryWrites=true&w=majority")
	// datastream.ConnectToMqttBroker("test.mosquitto.org", 1883)

	router := gin.New()
	router.Use(middlewares.DisableCors)
	router.GET("/ping", gin.Logger(), func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	endpoints.InitializeDataSourceEndPoints(router)

	wss.IntializeWss(router)

	router.Run()
}
