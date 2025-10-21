package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kaua-matheus/greenhouse-application/database"
)

func main(){
	Serve();
}

func Serve(){
	router := gin.Default();
	connection, err := database.NewConnection(); if err != nil {
		fmt.Printf("An error occurs trying to create a connection: %s", err);
	}

	// Get
	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"message": "A API está funcionando.",
		})
	});

	router.GET("/data", func (c *gin.Context) {
		data, err := database.GetAllData(connection); if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": http.StatusInternalServerError,
				"message": "An error occurs trying to get all the data",
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"data": data,
		})
	})

	router.Run();
}