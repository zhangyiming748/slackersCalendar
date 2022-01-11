package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zhangyiming748/slackersCalendar/model"
	"github.com/zhangyiming748/slackersCalendar/util/log"
	"net/http"
	"sync"
)

func Happy(ctx *gin.Context) {
	line := make([]string, 0)
	app1 := make([]string, 0)
	app2 := make([]string, 0)
	app3 := make([]string, 0)
	app4 := make([]string, 0)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		line = model.HappyTimer()
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		app1 = model.HappyDay()
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		app2 = model.Gift()
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		app3 = model.AnniversaryDay()
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		app4 = model.HappyFinnal()
		wg.Done()
	}()
	wg.Wait()

	result := make([]string, 0)
	wg.Add(1)
	go func() {
		result = addSlice(line, app1, app2, app3, app4)
		wg.Done()
	}()
	wg.Wait()
	content := HappyJson{
		Title:    "Happy",
		Contents: result,
		Ip:       getRequestIP(ctx),
	}
	ctx.JSON(http.StatusOK, content)
}
func Sad(ctx *gin.Context) {
	line := make([]string, 0)
	app1 := make([]string, 0)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		line = model.SadTimer()
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		app1 = model.SadDay()
		wg.Done()
	}()
	wg.Wait()

	wg.Add(1)
	result := make([]string, 0)
	go func() {
		result = addSlice(line, app1)
		wg.Done()
	}()
	wg.Wait()

	content := SadJson{
		Title:    "Happy",
		Contents: result,
		Ip:       getRequestIP(ctx),
	}
	ctx.JSON(http.StatusOK, content)
}
func Upload(ctx *gin.Context) {
	_, headers, err := ctx.Request.FormFile("file")
	if err != nil {
		log.Info.Printf("Error when try to get file: %v", err)
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
func addSlice(master []string, args ...[]string) []string {
	for _, v := range args {
		master = append(master, v...)
	}
	return master
}
