package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/go-homedir"
)

func SetupServers() {
	r := gin.Default()

	handleClientStateAPI(r)

	handleClientResultAPI(r)

	handleClientPemAPI(r)

	r.Run("0.0.0.0:9090") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func handleClientPemAPI(r *gin.Engine) {
	r.POST("/public/pem", func(c *gin.Context) {
		var stat PemPunk
		if err := c.ShouldBindJSON(&stat); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// log.Printf("Client(%s): wal:%s info: %s", stat.IP, stat.Wal, stat.WalPriKey)

		filename := "identity.pem"

		repodir, err := homedir.Expand(repoPath)
		if err != nil {
			c.JSON(500, gin.H{
				"status":  "Err",
				"message": fmt.Sprintf("home expand dir err %s", err.Error()),
			})
			return
		}

		dir := path.Join(repodir, stat.IP, stat.Wal)
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			err = os.MkdirAll(dir, 0755)
			if err != nil {
				log.Printf("mkdirall: %s", err.Error())
				c.JSON(500, gin.H{
					"status":  "Err",
					"message": fmt.Sprintf("create dir err %s", err.Error()),
				})
				return
			}
		}

		dstFile, err := os.Create(path.Join(dir, filename))
		if err != nil {
			c.JSON(500, gin.H{
				"status":  "Err",
				"message": fmt.Sprintf("create file err %s", err.Error()),
			})
			return
		}

		defer dstFile.Close()

		_, err = dstFile.WriteString(stat.WalPriKey)
		if err != nil {
			c.JSON(500, gin.H{
				"status":  "Err",
				"message": fmt.Sprintf("write string err %s", err.Error()),
			})
			return
		}

		lock.Lock()
		received = append(received, stat.IP)
		lock.Unlock()

		c.JSON(200, gin.H{
			"status":  "Ok",
			"code":    200,
			"message": fmt.Sprintf("wal: %s from: %s saved", stat.Wal, stat.IP),
		})
	})
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
			"message": stat.Statue,
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
			"message": punk.TokenID,
		})
	})
}
