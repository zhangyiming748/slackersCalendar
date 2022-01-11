package model

import (
	"sort"
	"strconv"
	"strings"
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
func SadDay()[]string {
	result :=make([]string,0)

	for _, v := range Lyingflat {
		if v.GetSubDay() < 0 {
			continue
		}
		if v.GetSubDay() == 0 || v.GetSubDay() == 365 {
			//log.Debug.Printf("明天是%v\n", v.GetChineseName())
			continue
		}

		result = append(result,strings.Join([]string{"距离",v.GetChineseName(),"还有",strconv.Itoa(v.GetSubDay()),"天"},""))
	}
	return result
}