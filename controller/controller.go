package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zhangyiming748/slackersCalendar/model"
	"github.com/zhangyiming748/slackersCalendar/util/log"
	"net/http"
	"strconv"
	"strings"
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

type file struct {
	F_type string `json:"f_type"`
	F_size string `json:"f_size"`
	F_err  string `json:"f_err"`
}

func Upload(ctx *gin.Context) {
	_, headers, err := ctx.Request.FormFile("file")
	if err != nil {
		log.Info.Printf("Error when try to get file: %v", err)
	}
	var result file
	//获取上传文件的类型
	if headers.Header.Get("Content-Type") != "text/plain" {
		result = file{
			F_size: strconv.FormatInt(headers.Size, 10),
			F_err:  "只允许上传txt文本文件",
			F_type: headers.Header.Get("Content-Type"),
		}
		ctx.JSON(http.StatusUnsupportedMediaType, result)
		return
	}
	//headers.Size 获取文件大小
	if headers.Size > 1024*1024*2 {
		m_size := float64(headers.Size) / 1024 / 1024
		result = file{
			F_size: strings.Join([]string{strconv.FormatFloat(m_size, 'f', 2, 64), "M"}, ""),
			//TODO 更改错误返回文字内容
			F_err:  "文件太大了(2M以下)",
			F_type: headers.Header.Get("Content-Type"),
		}
		ctx.JSON(http.StatusRequestEntityTooLarge, result)
		return
	}
	if err := ctx.SaveUploadedFile(headers, "./users.txt"); err != nil {
		result = file{
			F_err: "上传文件失败,目录是否不可写?",
		}
		ctx.JSON(http.StatusRequestEntityTooLarge, result)
		return
	}
	result = file{
		F_size: strconv.Itoa(int(headers.Size)),
		F_type: headers.Header.Get("Content-Type"),
	}
	ctx.JSON(http.StatusOK, result)
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
