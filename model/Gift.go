package model

import (
	"bufio"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"
)

/*
铁打的身体,磁铁打的床
*/
func Gift() []string {
	result := make([]string, 0)
	src := "./users.txt"
	if isExist(src) {
		users := readLine(src)
		for _, user := range users {
			result = append(result, strings.Join([]string{"帝豪全体技师预祝大客户", user, "节日快乐,期待您的下次光临"}, ""))
			result = append(result, strings.Join([]string{"豪庭全体技师预祝大客户", user, "节日快乐,期待您的下次光临"}, ""))
			result = append(result, strings.Join([]string{"君怡全体技师预祝大客户", user, "节日快乐,期待您的下次光临"}, ""))
			result = append(result, strings.Join([]string{"情思全体技师预祝大客户", user, "节日快乐,期待您的下次光临"}, ""))
			result = append(result, strings.Join([]string{"蓝天全体技师预祝大客户", user, "节日快乐,期待您的下次光临"}, ""))
			result = append(result, strings.Join([]string{"年代全体技师预祝大客户", user, "节日快乐,期待您的下次光临"}, ""))
			result = append(result, strings.Join([]string{"5号馆全体技师预祝大客户", user, "节日快乐,期待您的下次光临"}, ""))
			result = append(result, strings.Join([]string{"享乐汇全体技师预祝大客户", user, "节日快乐,期待您的下次光临"}, ""))
			result = append(result, strings.Join([]string{"新海员全体技师预祝大客户", user, "节日快乐,期待您的下次光临"}, ""))
			result = append(result, strings.Join([]string{"新太阳全体技师预祝大客户", user, "节日快乐,期待您的下次光临"}, ""))
			result = append(result, strings.Join([]string{"水悦轩全体技师预祝大客户", user, "节日快乐,期待您的下次光临"}, ""))
			result = append(result, strings.Join([]string{"金海湾全体技师预祝大客户", user, "节日快乐,期待您的下次光临"}, ""))
			result = append(result, strings.Join([]string{"水玲珑全体技师预祝大客户", user, "节日快乐,期待您的下次光临"}, ""))
			result = append(result, strings.Join([]string{"水疗会全体技师预祝大客户", user, "节日快乐,期待您的下次光临"}, ""))
			result = append(result, strings.Join([]string{"水善坊全体技师预祝大客户", user, "节日快乐,期待您的下次光临"}, ""))
			result = append(result, strings.Join([]string{"水云间全体技师预祝大客户", user, "节日快乐,期待您的下次光临"}, ""))
			result = append(result, strings.Join([]string{"嘉联华全体技师预祝大客户", user, "节日快乐,期待您的下次光临"}, ""))
			result = append(result, strings.Join([]string{"新境界全体技师预祝大客户", user, "节日快乐,期待您的下次光临"}, ""))
			result = append(result, strings.Join([]string{"新名仕全体技师预祝大客户", user, "节日快乐,期待您的下次光临"}, ""))
			result = append(result, strings.Join([]string{"龙湖阁全体技师预祝大客户", user, "节日快乐,期待您的下次光临"}, ""))
			result = append(result, strings.Join([]string{"紫汕坊全体技师预祝大客户", user, "节日快乐,期待您的下次光临"}, ""))
			result = append(result, strings.Join([]string{"新美龙全体技师预祝大客户", user, "节日快乐,期待您的下次光临"}, ""))
			result = append(result, strings.Join([]string{"云顶天境全体技师预祝大客户", user, "节日快乐,期待您的下次光临"}, ""))
			result = append(result, strings.Join([]string{"鹏发理疗全体技师预祝大客户", user, "节日快乐,期待您的下次光临"}, ""))
			result = append(result, strings.Join([]string{"衡山水疗全体技师预祝大客户", user, "节日快乐,期待您的下次光临"}, ""))
			result = append(result, strings.Join([]string{"龙湖夜总会全体妈咪预祝大客户", user, "节日快乐,期待您的下次光临"}, ""))
			result = append(result, strings.Join([]string{"四季酒店全体技师预祝大客户", user, "节日快乐,期待您的下次光临"}, ""))
			result = append(result, strings.Join([]string{"八号公馆全体技师预祝大客户", user, "节日快乐,期待您的下次光临"}, ""))
		}
	}
	rand.Seed(time.Now().Unix())
	//fmt.Println(words[rand.Intn(len(words))])
	result = append(result, words[rand.Intn(len(words))])
	return result
}
func readLine(src string) []string {
	fi, err := os.Open(src)
	if err != nil {
		//log.Info.Printf("打开用户目录文件失败: %s\n", err)
		return []string{}
	}
	defer func() {
		if err := fi.Close(); err != nil {
			//log.Info.Printf("关闭用户目录文件失败: %s\n", err)
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
		//log.Info.Printf("读取到的用户(%s)\n", string(a))
	}
	return names
}
func isExist(fname string) bool {
	_, err := os.Stat(fname)
	if err == nil {
		//log.Info.Println("File exist")
		return true
	}
	if os.IsNotExist(err) {
		//log.Info.Println("File not exist")
		return false
	}
	return false
}
