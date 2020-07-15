package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
	"path/filepath"
)

const (
	UPLOAD_DIR      = "/tmp"
	URL_PREFIX      = "http://localhost:8080/download/"
	USER_BASIC_AUTH = true
	USER_NAME       = "vearne"
	PASSWORD        = "shuai"
)


func DealUpload(c *gin.Context) {
	// single file
	file, _ := c.FormFile("file")
	log.Println(file.Filename)
	// Upload the file to specific dst.
	dst := filepath.Join(UPLOAD_DIR, file.Filename)
	c.SaveUploadedFile(file, dst)
	c.String(http.StatusOK, "DownloadPath: %s\n", URL_PREFIX+file.Filename)
}

func main() {
	router := gin.Default()
	if USER_BASIC_AUTH {
		router.Use(gin.BasicAuth(gin.Accounts{
			USER_NAME:    PASSWORD, //用户名：密码
		}))
	}
	router.StaticFS("/download", http.Dir(UPLOAD_DIR))
	router.POST("/upload", DealUpload)
	router.Run(":8080")
}
