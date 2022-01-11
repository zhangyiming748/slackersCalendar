package main

import (
	"github.com/zhangyiming748/slackersCalendar/model"
)

func main() {
	Web()
	All()
	AllHappy()
	AllSad()
}
func Web() {
	WebServer(3306)
}

func AllHappy() {
	model.HappyTimer()
	model.HappyDay()
	model.Gift()
	model.AnniversaryDay()
	model.HappyFinnal()
}
func AllSad() {
	model.SadTimer()
	model.SadDay()
}
func All() {
	model.HappyTimer()
	model.HappyDay()
	model.Gift()
	model.AnniversaryDay()
	model.HappyFinnal()
	model.SadTimer()
	model.SadDay()
}
