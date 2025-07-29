package util

import "fmt"

func GetPrefixNum(s string) string {
	if s == "" {
		return ""
	}
	index := len(s)
	for i, c := range s {
		if !IsNum(c) {
			index = i
			break
		}
	}
	return s[0:index]
}

func IsNum(c int32) bool {
	b := false
	if c >= 48 && c <= 57 {
		b = true
	}
	return b
}

func PrintArrString(arr []string) {
	for _, a := range arr {
		fmt.Println(a)
	}

}
