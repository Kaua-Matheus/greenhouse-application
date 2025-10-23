package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	router.GET("/", func(c *gin.Context) {
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

	// Post
	router.POST("/data", func (c *gin.Context) {
		data := database.SensorData{};
		
		if err := c.BindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": http.StatusBadRequest,
				"message": "Error: Bind JSON doesn't work as well",
			})
		}

		database.AddData(connection, data);
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"data": "Data registered successfully",
		})
	})

	// Put
	router.PUT("/data/:id", func (c *gin.Context){
		data := database.SensorData{}

		id_str := c.Param("id");
		id, err := uuid.Parse(id_str); if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": http.StatusBadRequest,
				"message": "Error: Invalid UUID format",
			})
		}

		if err := c.BindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": http.StatusBadRequest,
				"message": "Error: Bind JSON doesn't work as well",
			})
		}

		err = database.UpdateData(connection, id,  data); if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": http.StatusInternalServerError,
				"message": "Error: Couldn't update the data",
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"data": "Data updated successfully",
		})

	})

	// Delete
	router.DELETE("/data/:id", func (c *gin.Context){
		id_str := c.Param("id");
		id, err := uuid.Parse(id_str); if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": http.StatusBadRequest,
				"message": "Error: Invalid UUID format",
			})
			return;
		}

		if err := database.DeleteData(connection, id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": http.StatusInternalServerError,
				"message": "Error: Couldn't delete the data",
				"error": err,
			})
			return;
		}

		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"data": "Data deleted successfully",
		})
	})

	router.Run();
}