package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Happy(ctx *gin.Context) {
	content := HappyJson{
		Title:    "Happy",
		Contents: []string{},
		Ip:       getRequestIP(ctx),
	}
	ctx.JSON(http.StatusOK, content)
}
func Sad(ctx *gin.Context) {
	content := SadJson{
		Title:    "Happy",
		Contents: []string{},
		Ip:       getRequestIP(ctx),
	}
	ctx.JSON(http.StatusOK, content)
}
func Upload(ctx *gin.Context) {
	_, headers, err := ctx.Request.FormFile("file")
	if err != nil {
		log.Printf("Error when try to get file: %v", err)
	}
	//headers.Size 获取文件大小
	if headers.Size > 1024*1024*2 {
		fmt.Println("文件太大了")
		return
	}
	//headers.Header.Get("Content-Type")获取上传文件的类型
	if headers.Header.Get("Content-Type") != "text/plain" {
		fmt.Println("只允许上传txt文本文件")
		return
	}

	if err := ctx.SaveUploadedFile(headers, "./users.txt"); err != nil {
		return
	}
	ctx.String(http.StatusOK, headers.Filename)
}
func getRequestIP(c *gin.Context) string {
	reqIP := c.ClientIP()
	if reqIP == "::1" {
		reqIP = "127.0.0.1"
	}
	return reqIP
}
