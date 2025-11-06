package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/emmm1245/EdgeCopilotFix/internal/config"
	"github.com/emmm1245/EdgeCopilotFix/internal/logger"
	"github.com/emmm1245/EdgeCopilotFix/internal/process"
)

func main() {
	// 输出欢迎信息
	fmt.Println("======================")
	fmt.Println("Edge Copilot Fix Tool")
	fmt.Println("======================")
	fmt.Printf("操作系统: %s\n", getOSName())
	fmt.Println()

	// 第一步：处理 Edge 进程
	if err := process.HandleEdgeProcesses(); err != nil {
		logger.Error(err.Error())
		waitForExit()
		os.Exit(1)
	}

	fmt.Println()

	// 第二步：修复所有检测到的 Edge 版本配置文件
	if err := config.FixAllEdgeVersions(); err != nil {
		logger.Error(err.Error())
		waitForExit()
		os.Exit(1)
	}

	fmt.Println()
	logger.Success("所有操作完成！")
	logger.Info("请重新启动 Edge 浏览器以使更改生效。")

	waitForExit()
}

// getOSName 获取操作系统的友好名称
func getOSName() string {
	switch runtime.GOOS {
	case "windows":
		return "Windows"
	case "darwin":
		return "macOS"
	case "linux":
		return "Linux"
	default:
		return runtime.GOOS
	}
}

// waitForExit 等待用户按 Enter 键退出
func waitForExit() {
	fmt.Println()
	fmt.Println("按 Enter 键退出...")
	fmt.Scanln()
}
