package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	Serve();
}

func Serve(){
	router := gin.Default();

	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"message": "A API está funcionando.",
		})
	});

	router.Run();
}
