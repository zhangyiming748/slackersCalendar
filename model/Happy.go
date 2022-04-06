package model

import (
	"fmt"
	"github.com/nosixtools/solarlunar"
	"sort"
	"strconv"
	"strings"
	"time"
)

func init() {
	NewYear.SetChineseName("元旦")
	NewYear.SetDate("01-01")
	NewYear.SetSubDay(allInSolar(NewYear.GetDate()))
	Countdown = append(Countdown, NewYear)

	ValentinesDay.SetChineseName("情人节")
	ValentinesDay.SetDate("02-14")
	ValentinesDay.SetSubDay(allInSolar(ValentinesDay.GetDate()))
	Countdown = append(Countdown, ValentinesDay)

	WomensDay.SetChineseName("妇女节")
	WomensDay.SetDate("03-08")
	WomensDay.SetSubDay(allInSolar(WomensDay.GetDate()))
	Countdown = append(Countdown, WomensDay)

	ArborDay.SetChineseName("植树节")
	ArborDay.SetDate("03-12")
	ArborDay.SetSubDay(allInSolar(ArborDay.GetDate()))
	Countdown = append(Countdown, ArborDay)

	AprilFoolsDay.SetChineseName("愚人节")
	AprilFoolsDay.SetDate("04-01")
	AprilFoolsDay.SetSubDay(allInSolar(AprilFoolsDay.GetDate()))
	Countdown = append(Countdown, AprilFoolsDay)

	LaborDay.SetChineseName("劳动节")
	LaborDay.SetDate("05-01")
	LaborDay.SetSubDay(allInSolar(LaborDay.GetDate()))
	Countdown = append(Countdown, LaborDay)

	YouthDay.SetChineseName("青年节")
	YouthDay.SetDate("05-04")
	YouthDay.SetSubDay(allInSolar(YouthDay.GetDate()))
	Countdown = append(Countdown, YouthDay)

	ChildrensDay.SetChineseName("儿童节")
	ChildrensDay.SetDate("06-01")
	ChildrensDay.SetSubDay(allInSolar(ChildrensDay.GetDate()))
	Countdown = append(Countdown, ChildrensDay)

	NationalDay.SetChineseName("国庆节")
	NationalDay.SetDate("10-01")
	NationalDay.SetSubDay(allInSolar(NationalDay.GetDate()))
	Countdown = append(Countdown, NationalDay)

	NationalPovertyReliefDay.SetChineseName("国家扶贫日")
	NationalPovertyReliefDay.SetDate("10-17")
	NationalPovertyReliefDay.SetSubDay(allInSolar(NationalPovertyReliefDay.GetDate()))
	Countdown = append(Countdown, NationalPovertyReliefDay)

	Halloween.SetChineseName("万圣夜")
	Halloween.SetDate("10-31")
	Halloween.SetSubDay(allInSolar(Halloween.GetDate()))
	Countdown = append(Countdown, Halloween)

	Double11.SetChineseName("双十一")
	Double11.SetDate("11-11")
	Double11.SetSubDay(allInSolar(Double11.GetDate()))
	Countdown = append(Countdown, Double11)

	NationalConstitutionDay.SetChineseName("国家宪法日")
	NationalConstitutionDay.SetDate("12-04")
	NationalConstitutionDay.SetSubDay(allInSolar(NationalConstitutionDay.GetDate()))
	Countdown = append(Countdown, NationalConstitutionDay)

	ThanksgivingDay.SetChineseName("感恩节")
	ThanksgivingDay.SetDate("11-25")
	ThanksgivingDay.SetSubDay(allInSolar(ThanksgivingDay.GetDate()))
	Countdown = append(Countdown, ThanksgivingDay)

	Double12.SetChineseName("双十二")
	Double12.SetDate("12-12")
	Double12.SetSubDay(allInSolar(Double12.GetDate()))
	Countdown = append(Countdown, Double12)

	NationalMemorialDay.SetChineseName("国家公祭日")
	NationalMemorialDay.SetDate("12-13")
	NationalMemorialDay.SetSubDay(allInSolar(NationalMemorialDay.GetDate()))
	Countdown = append(Countdown, NationalMemorialDay)

	ChristmasEve.SetChineseName("平安夜")
	ChristmasEve.SetDate("12-24")
	ChristmasEve.SetSubDay(allInSolar(ChristmasEve.GetDate()))
	Countdown = append(Countdown, ChristmasEve)

	Christmas.SetChineseName("圣诞节")
	Christmas.SetDate("12-25")
	Christmas.SetSubDay(allInSolar(Christmas.GetDate()))
	Countdown = append(Countdown, Christmas)

	ChineseNewYear.SetChineseName("除夕")
	ChineseNewYear.SetDate("01-01")
	//ToDo 农历的春节时间暂时使用当年1号实现,年底如果发现问题再更改
	ChineseNewYear.SetSubDay(allInSpring(ChineseNewYear.GetDate()))
	Countdown = append(Countdown, ChineseNewYear)

	LanternFestival.SetChineseName("元宵节")
	LanternFestival.SetDate("01-15")
	LanternFestival.SetSubDay(allInLuna(LanternFestival.GetDate()))
	Countdown = append(Countdown, LanternFestival)

	ChingMingFestival.SetChineseName("清明节")
	ChingMingFestival.SetDate("03-05")
	ChingMingFestival.SetSubDay(allInLuna(ChingMingFestival.GetDate()))
	Countdown = append(Countdown, ChingMingFestival)

	MothersDay.SetChineseName("母亲节")
	MothersDay.SetDate("04-08")
	MothersDay.SetSubDay(allInLuna(MothersDay.GetDate()))
	Countdown = append(Countdown, MothersDay)

	DragonBoatFestival.SetChineseName("端午节")
	DragonBoatFestival.SetDate("05-05")
	DragonBoatFestival.SetSubDay(allInLuna(DragonBoatFestival.GetDate()))
	Countdown = append(Countdown, DragonBoatFestival)

	SingleDogDay.SetChineseName("七夕节")
	SingleDogDay.SetDate("07-07")
	SingleDogDay.SetSubDay(allInLuna(SingleDogDay.GetDate()))
	Countdown = append(Countdown, SingleDogDay)

	MidyearFestival.SetChineseName("中元节")
	MidyearFestival.SetDate("07-15")
	MidyearFestival.SetSubDay(allInLuna(MidyearFestival.GetDate()))
	Countdown = append(Countdown, MidyearFestival)

	MidAutumnFestival.SetChineseName("中秋节")
	MidAutumnFestival.SetDate("08-15")
	MidAutumnFestival.SetSubDay(allInLuna(MidAutumnFestival.GetDate()))
	Countdown = append(Countdown, MidAutumnFestival)

	DoubleNinthFestival.SetChineseName("重阳节")
	DoubleNinthFestival.SetDate("09-09")
	DoubleNinthFestival.SetSubDay(allInLuna(DoubleNinthFestival.GetDate()))
	Countdown = append(Countdown, DoubleNinthFestival)

	LabaFestival.SetChineseName("腊八节")
	LabaFestival.SetDate("12-08")
	LabaFestival.SetSubDay(allInLuna(LabaFestival.GetDate()))
	Countdown = append(Countdown, LabaFestival)

	WinterOlympics.SetChineseName("2022冬奥会")
	WinterOlympics.SetDate("2022-02-04")
	WinterOlympics.SetSubDay(allInThisYear(WinterOlympics.GetDate()))
	Countdown = append(Countdown, WinterOlympics)

	sort.Sort(HappySlice(Countdown))
}
func allInSolar(date string) int {
	day := strings.Join([]string{thisYear(), date}, "-")
	zone := strings.Join([]string{day, "0800"}, "+")
	ret, _ := time.Parse("2006-01-02-0700", zone)
	unsub := ret.Sub(time.Now())
	if unsub < 0 {
		return int(unsub.Hours())/24 + Year
	}
	return int(unsub.Hours()) / 24
}

//计算特殊日期 如奥运会
func allInThisYear(date string) int {
	day := date
	zone := strings.Join([]string{day, "0800"}, "+")
	ret, _ := time.Parse("2006-01-02-0700", zone)
	unsub := ret.Sub(time.Now())
	if unsub < 0 {
		return int(unsub.Hours())/24 + Year
	}
	return int(unsub.Hours()) / 24
}

//计算农历时间
func allInLuna(date string) int {
	day := strings.Join([]string{thisYear(), date}, "-")
	convert := solarlunar.LunarToSolar(day, false)
	zone := strings.Join([]string{convert, "0800"}, "+")
	ret, _ := time.Parse("2006-01-02-0700", zone)
	unsub := ret.Sub(time.Now())
	if unsub < 0 {
		return int(unsub.Hours())/24 + Year
	}
	return int(unsub.Hours()) / 24
}
func allInSpring(date string) int {
	day := strings.Join([]string{thisYear(), date}, "-")
	convert := solarlunar.LunarToSolar(day, false)
	ret, _ := time.Parse("2006-01-02", convert)
	unsub := ret.Sub(time.Now())
	if unsub < 0 {
		day := strings.Join([]string{nextYear(), date}, "-")
		convert := solarlunar.LunarToSolar(day, false)
		ret, _ := time.Parse("2006-01-02", convert)
		unsub := ret.Sub(time.Now())
		return int(unsub.Hours()) / 24
	}
	return int(unsub.Hours())/24 - 1
}

//获取当前年份
func thisYear() string {
	ret := fmt.Sprint(time.Now().Format("2006"))
	return ret
}
func nextYear() string {
	this := thisYear()
	i, _ := strconv.Atoi(this)
	i += 1
	s := strconv.Itoa(i)
	return s
}

/*
上班是帮老板赚钱,摸鱼是赚老板的钱!
该休息就休息,该放松就放松
距离除夕还有20天
距离2022冬奥会还有24天
距离情人节还有34天
距离元宵节还有35天
距离妇女节还有56天
距离植树节还有60天
距离愚人节还有80天
距离清明节还有84天
距离劳动节还有110天
距离青年节还有113天
距离母亲节还有117天
距离儿童节还有141天
距离端午节还有143天
距离七夕节还有205天
距离中元节还有213天
距离中秋节还有242天
距离国庆节还有263天
距离重阳节还有266天
距离国家扶贫日还有279天
距离万圣夜还有293天
距离双十一还有304天
距离感恩节还有318天
距离国家宪法日还有327天
距离双十二还有335天
距离国家公祭日还有336天
距离平安夜还有347天
距离圣诞节还有348天
距离腊八节还有353天
距离元旦还有356天
*/
func HappyDay() []string {
	result := make([]string, 0)
	result = append(result, "上班是帮老板赚钱,摸鱼是赚老板的钱!")
	result = append(result, "该休息就休息,该放松就放松")
	for _, v := range Countdown {
		if v.GetSubDay() < 1 {
			switch v.GetChineseName() {
			case "七夕节":
				result = append(result, fmt.Sprintf("今天是七夕节.看着走进电影院的情侣,我笑了笑,并掏出黑卡把今天所有的双数座都买了下来.它们的表情特别有趣,从刚走进电影院的欣喜,再到我扬言买下所有双数座的嘲讽,直到我掏出黑卡的惊讶与失望.每一个表情都令我十分欣慰.有钱人的快乐往往就是这么朴实无华且枯燥"))
			case "情人节":
				result = append(result, fmt.Sprintf("昨晚去看电影,电影院里坐满了一对对的情侣,我这个单身狗独自坐在第一排.就在准备熄灯开始放映前,工作人员在广播里说:\"哪位观众是川L88888的宾利车主,请去车库挪下车\"于是我果断地起身并在所有人的注视下走出了放映厅,到厕所抽了根烟"))
			default:
				result = append(result, fmt.Sprintf("今天是%v", v.GetChineseName()))
			}
		}
		result = append(result, strings.Join([]string{"距离", v.GetChineseName(), "还有", strconv.Itoa(v.GetSubDay()), "天"}, ""))
	}
	return result
}

/*
最后,祝愿天下所有摸鱼人,都能愉快的渡过每一天
*/
func HappyFinnal() []string {
	result := make([]string, 0)
	result = append(result, "最后,祝愿天下所有摸鱼人,都能愉快的渡过每一天")
	return result
}
