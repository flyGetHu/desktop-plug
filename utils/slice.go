package utils

import "strings"

// ContainsInSlice 判断字符串是否在 slice 中
func ContainsInSlice(items []string, item string) bool {
	for _, eachItem := range items {
		if strings.ToLower(eachItem) == strings.ToLower(item) {
			return true
		}
	}
	return false
}
