package model

import (
	"fmt"
	"github.com/zhangyiming748/slackersCalendar/util"
	"github.com/zhangyiming748/slackersCalendar/util/log"
	"sort"
	"strconv"
	"strings"
	"time"
)

func init() {

	QVOD.SetName("快播关闭")
	QVOD.SetDate("2014年4月16日")
	QVOD.SetSubDay(CountDay(QVOD.GetDate()))
	AS = append(AS, QVOD)

	Apple.SetName("Iphone发布")
	Apple.SetDate("2007年1月9日")
	Apple.SetSubDay(CountDay(Apple.GetDate()))
	AS = append(AS, Apple)

	sort.Sort(AS)
}

func CountDay(date string) int {
	ret, _ := time.Parse("2006年1月2日", date)
	old := fmt.Sprintf("%v", ret)
	oldYear := strings.Split(old, "-")[0]
	thisYear := fmt.Sprint(time.Now().Format("2006"))
	oi, _ := strconv.Atoi(oldYear)
	ti, _ := strconv.Atoi(thisYear)
	sub := ti - oi
	//fmt.Printf("old is %v\nthis is %v\noldYear is %v\nsub is %v\n", old, thisYear, oldYear, sub)
	return sub
}

/*
距离快播关闭已经过去了8周年	但你还欠王欣一个年费会员
距离Iphone发布已经过去了15周年
*/
func AnniversaryDay() []string {
	//fmt.Println("P.S.")
	result := make([]string, 0)
	for _, v := range AS {
		//fmt.Printf("距离%v已经过去了%v周年\n", v.GetName(), v.GetSubDay())
		switch v.GetName() {
		case "快播关闭":
			result = append(result, util.ReadAndWrite(strings.Join([]string{"距离", v.GetName(), "已经过去了", strconv.Itoa(v.GetSubDay()), "年,但你还欠王欣一个年费会员"}, "")))
			log.Debug.Printf("距离%v已经过去了%v周年\t但你还欠王欣一个年费会员\n", v.GetName(), v.GetSubDay())
		default:
			result = append(result, util.ReadAndWrite(strings.Join([]string{"距离", v.GetName(), "已经过去了", strconv.Itoa(v.GetSubDay()), "年"}, "")))
		}
	}
	return result
}
