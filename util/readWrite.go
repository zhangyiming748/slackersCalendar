package util

import (
	"fmt"
	"github.com/zhangyiming748/slackersCalendar/util/log"
)

func ReadAndWrite(s string)string  {
	log.Debug.Println(s)
	ret:=fmt.Sprint(s)
	return ret
}
