package logger

import (
	"fmt"
)

// Info 输出信息日志
func Info(msg string) {
	fmt.Printf("\033[36m[信息] %s\033[0m\n", msg)
}

// Success 输出成功日志
func Success(msg string) {
	fmt.Printf("\033[32m[成功] %s\033[0m\n", msg)
}

// Error 输出错误日志
func Error(msg string) {
	fmt.Printf("\033[31m[错误] %s\033[0m\n", msg)
}

// Warning 输出警告日志
func Warning(msg string) {
	fmt.Printf("\033[33m[警告] %s\033[0m\n", msg)
}


