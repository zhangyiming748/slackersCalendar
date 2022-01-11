package slackersCalendar

import (
	"github.com/zhangyiming748/slackersCalendar/http"
	"github.com/zhangyiming748/slackersCalendar/model"
)

func Happy() {
	model.HappyTimer()
	model.HappyDay()
	model.Gift()
	model.AnniversaryDay()
	model.HappyFinnal()
}
func Sad() {
	model.SadTimer()
	model.SadDay()
}
func WebServer(port int) {
	http.ShowWeb(port)
}
