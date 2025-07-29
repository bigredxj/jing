package util

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func GetWorkDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := filepath.Dir(b)
	rootDir := filepath.Join(d, "../..") // 假设根目录在当前包的上一级目录
	absRootDir, _ := filepath.Abs(rootDir)
	return absRootDir
}

func WriteOutput(path string, arr []string) {
	f, err := os.Create(path)
	defer f.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, r := range arr {
		//fmt.Println(r)
		f.WriteString(r + "\n")
	}

}

func AppendToFile(path string, value string) {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close() // 确保在函数返回时关闭文件

	// 使用bufio.Writer来提高写入性能
	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(value + "\n")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	writer.Flush() // 确
}

func ReadLinesFromFile(path string) []string {
	file, _ := os.Open(path)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
