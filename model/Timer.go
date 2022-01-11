package model

import (
	"github.com/zhangyiming748/slackersCalendar/util"
	"strconv"
	"strings"
	"time"
)

/*
下午好,摸鱼人!
今天是2022年1月10日 星期一,还有四天周末 下午 4点36分08秒
周一不摸鱼,脑子有问题
周二不摸鱼,上班没心情
周三不摸鱼,你是头傻驴
周四不摸鱼,谁请肯德基
周五不摸鱼,浪费好心情
周六不摸鱼,活着真多余
【摸鱼办】提醒您
工作再累,一定不要忘记摸鱼哦!
有事没事起身去茶水间,去厕所,去廊道走走
吃饭时间就吃饭
午休时间就午休
别老在工位上坐着,钱是老板的,但命是自己的
*/
func HappyTimer() []string {
	result := make([]string, 0)
	nowDay := time.Now().Format("今天是2006年1月2日")
	nowTime := time.Now().Format("PM 3点04分05秒")
	week := time.Now().Weekday().String()
	hour := time.Now().Format("15")
	inthour, _ := strconv.Atoi(hour)
	nowTime = strings.Replace(nowTime, "AM", "上午", -1)
	nowTime = strings.Replace(nowTime, "PM", "下午", -1)
	if inthour >= 6 && inthour <= 11 {
		result = append(result, util.ReadAndWrite("早上好,摸鱼人!"))
	}
	if inthour == 12 {
		result = append(result, util.ReadAndWrite("中午好,摸鱼人!"))
	}
	if inthour >= 13 && inthour <= 17 {
		result = append(result, util.ReadAndWrite("下午好,摸鱼人!"))
	}
	if inthour >= 18 && inthour <= 22 {
		result = append(result, util.ReadAndWrite("晚上好,摸鱼人!"))
	}
	if inthour >= 23 && inthour <= 5 {
		result = append(result, util.ReadAndWrite("晚上抓紧睡觉,休息好,第二天才有精神摸鱼!"))
	}
	result = append(result, util.ReadAndWrite(nowDay))
	result = append(result, util.ReadAndWrite(HappyWeekDayMap[week]))
	result = append(result, util.ReadAndWrite(nowTime))
	result = append(result, util.ReadAndWrite("周一不摸鱼,脑子有问题"))
	result = append(result, util.ReadAndWrite("周二不摸鱼,上班没心情"))
	result = append(result, util.ReadAndWrite("周三不摸鱼,你是头傻驴"))
	result = append(result, util.ReadAndWrite("周四不摸鱼,谁请肯德基"))
	result = append(result, util.ReadAndWrite("周五不摸鱼,浪费好心情"))
	result = append(result, util.ReadAndWrite("周六不摸鱼,活着真多余"))

	if inthour == 15 {
		result = append(result, util.ReadAndWrite("喂!三点几咧!做......做撚啊做!"))
		result = append(result, util.ReadAndWrite("饮茶先啊!"))
		result = append(result, util.ReadAndWrite("三点几 饮......饮茶先啊!"))
		result = append(result, util.ReadAndWrite("做咁多都冇用嘅,老细唔锡你嘅咧!"))
		result = append(result, util.ReadAndWrite("喂!饮下茶先啊!三点几咧!"))
		result = append(result, util.ReadAndWrite("做碌鸠啊做!"))
	}
	if inthour == 18 {
		result = append(result, util.ReadAndWrite("喂!朋友!"))
		result = append(result, util.ReadAndWrite("做乜嘢咁多啦?差唔多七点咧!"))
		result = append(result, util.ReadAndWrite("放工啦!唔使做咁多啦!"))
		result = append(result, util.ReadAndWrite("做咁多,钱带去边度?"))
		result = append(result, util.ReadAndWrite("差唔多七点咧!放工!"))
		result = append(result, util.ReadAndWrite("落去茶室,饮下靓靓嘅杯......麦啤酒、黑啤酒,OK?"))
		result = append(result, util.ReadAndWrite("Happy下,唔使做咁多."))
		result = append(result, util.ReadAndWrite("死咗都冇用咧,银纸冇得带去咧!"))
		result = append(result, util.ReadAndWrite("Happy下,饮酒,OK?"))
	}
	result = append(result, util.ReadAndWrite("【摸鱼办】提醒您"))
	result = append(result, util.ReadAndWrite("工作再累,一定不要忘记摸鱼哦!"))
	result = append(result, util.ReadAndWrite("有事没事起身去茶水间,去厕所,去廊道走走"))
	result = append(result, util.ReadAndWrite("吃饭时间就吃饭"))
	result = append(result, util.ReadAndWrite("午休时间就午休"))
	result = append(result, util.ReadAndWrite("别老在工位上坐着,钱是老板的,但命是自己的"))

	return result
}

/*
下午好,继续卷!
今天是2022年1月10日 下午4点39分52秒 星期一
【内卷办】提醒您
*/
func SadTimer() []string {
	result := make([]string, 0)
	nowDay := time.Now().Format("今天是2006年1月2日")
	nowTime := time.Now().Format("PM3点04分05秒")
	week := time.Now().Weekday().String()
	hour := time.Now().Format("15")
	inthour, _ := strconv.Atoi(hour)
	nowTime = strings.Replace(nowTime, "AM", "上午", -1)
	nowTime = strings.Replace(nowTime, "PM", "下午", -1)
	if inthour >= 6 && inthour <= 11 {
		result = append(result, util.ReadAndWrite("早上好,开始内卷的一天!"))
	}
	if inthour == 12 {
		result = append(result, util.ReadAndWrite("中午好,别人休息的时候正是你内卷的好时候!"))
	}
	if inthour >= 13 && inthour <= 17 {
		result = append(result, util.ReadAndWrite("下午好,继续卷!"))
	}
	if inthour >= 18 && inthour <= 22 {
		result = append(result, util.ReadAndWrite("晚上好,坚持住!"))
	}
	if inthour >= 23 && inthour <= 5 {
		result = append(result, util.ReadAndWrite("晚上抓紧卷,不睡觉你才能获得比别人更多的时间!"))
	}
	result = append(result, util.ReadAndWrite(nowDay))
	result = append(result, util.ReadAndWrite(nowTime))
	result = append(result, util.ReadAndWrite(SadWeekDayMap[week]))
	result = append(result, util.ReadAndWrite("【内卷办】提醒您"))
	return result
}
