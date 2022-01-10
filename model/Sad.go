package model

import (
	"github.com/zhangyiming748/slackersCalendar/util/log"
	"sort"
)

//输入sad列表
func init() {
	CollegeEntranceExamination.SetChineseName("高考")
	CollegeEntranceExamination.SetDate("06-07")
	CollegeEntranceExamination.SetSubDay(allInSolar(CollegeEntranceExamination.GetDate()))
	Lyingflat = append(Lyingflat, CollegeEntranceExamination)

	sort.Sort(SadSlice(Lyingflat))
}
/*
距离高考还有147天
 */
func SadDay() {
	for _, v := range Lyingflat {
		if v.GetSubDay() == 0 || v.GetSubDay() == 365 {
			log.Debug.Printf("\t%v\n", v.GetChineseName())
		}
	}
	for _, v := range Lyingflat {
		if v.GetSubDay() < 0 {
			continue
		}
		if v.GetSubDay() == 0 || v.GetSubDay() == 365 {
			//log.Debug.Printf("明天是%v\n", v.GetChineseName())
			continue
		}
		log.Debug.Printf("距离%v还有%v天\n", v.GetChineseName(), v.GetSubDay())
	}
}