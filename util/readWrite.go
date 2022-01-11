package util

import "fmt"

func ReadAndWrite(s string)string  {
	fmt.Println(s)
	ret:=fmt.Sprint(s)
	return ret
}
