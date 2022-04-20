package utils

import "strconv"

//ReverseStr 反转字符串
func ReverseStr(s string) string {
	a := []rune(s)
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return string(a)
}

// StrIsNum 校验字符串是否为数字
func StrIsNum(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}
