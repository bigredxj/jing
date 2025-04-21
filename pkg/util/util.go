package util

import (
	"path/filepath"
	"runtime"
)

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

func GetWorkDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := filepath.Dir(b)
	rootDir := filepath.Join(d, "../..") // 假设根目录在当前包的上一级目录
	absRootDir, _ := filepath.Abs(rootDir)
	return absRootDir
}
