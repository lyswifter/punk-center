package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupServers() {
	r := gin.Default()

	handleClientStateAPI(r)
	handleClientResultAPI(r)

	r.Run("0.0.0.0:9090") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func handleClientStateAPI(r *gin.Engine) {
	r.POST("/public/stat", func(c *gin.Context) {
		var stat StatuePunk
		if err := c.ShouldBindJSON(&stat); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		log.Printf("Client(%s): stat: %s ", stat.IP, stat.Statue)

		c.JSON(200, gin.H{
			"status":  "Ok",
			"code":    200,
			"message": "",
		})
	})
}

func handleClientResultAPI(r *gin.Engine) {
	r.POST("/public/ret", func(c *gin.Context) {
		var punk ResultPunk
		if err := c.ShouldBindJSON(&punk); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		log.Printf("message: %+v", punk)

		err := saveInfo(punk)
		if err != nil {
			c.JSON(500, gin.H{
				"status":  "Err",
				"message": fmt.Sprintf("saveInfo err %s", err.Error()),
			})
			return
		}

		c.JSON(200, gin.H{
			"status":  "Ok",
			"code":    200,
			"message": "",
		})
	})
}
