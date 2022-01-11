package controller

import (
	"bufio"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zhangyiming748/slackersCalendar/model"
	"github.com/zhangyiming748/slackersCalendar/util/log"
	"io"

	"net/http"
	"os"
	"os/exec"

)

func Happy(ctx *gin.Context) {
	line:=model.HappyTimer()
	model.HappyDay()
	model.Gift()
	model.AnniversaryDay()
	model.HappyFinnal()
	content := HappyJson{
		Title:    "Happy",
		Contents: line,
		Ip:       getRequestIP(ctx),
	}
	ctx.JSON(http.StatusOK, content)
}
func Sad(ctx *gin.Context) {
	model.SadTimer()
	model.SadDay()
	info:=readLine("./info.txt")
	content := SadJson{
		Title:    "Happy",
		Contents: info,
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
func readLine(src string) []string {
	fi, err := os.Open(src)
	if err != nil {
		log.Info.Printf("打开用户目录文件失败: %s\n", err)
		return []string{}
	}
	defer func() {
		if err := fi.Close(); err != nil {
			log.Info.Printf("关闭用户目录文件失败: %s\n", err)
		}
	}()
	names := []string{}
	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		names = append(names, string(a))
		//log.Printf("读取到的用户(%s)\n", string(a))
	}
	return names
}
//ToDo 每次接收请求之前删除文件
func removeAll(){
	cmd := exec.Command("find", ".", "-name", "info.txt", "-exec", "rm",	"{}", "\\;")
	//log.Info.Printf("生成的命令是:%s", cmd)
	// 命令的错误输出和标准输出都连接到同一个管道
	//stdout, err := cmd.StdoutPipe()
	//cmd.Stderr = cmd.Stdout
	//if err != nil {
	//	log.Info.Printf("cmd.StdoutPipe产生的错误:%v", err)
	//}
	if err := cmd.Start(); err != nil {
		log.Info.Printf("cmd.Run产生的错误:%v", err)
	}
	// 从管道中实时获取输出并打印到终端
	//for {
	//	tmp := make([]byte, 1024)
	//	_, err := stdout.Read(tmp)
	//	//写成输出日志
	//	log.Info.Println(string(tmp))
	//	if err != nil {
	//		break
	//	}
	//}
	if err := cmd.Wait(); err != nil {
		log.Info.Println("命令执行中有错误产生", err)
	}
}