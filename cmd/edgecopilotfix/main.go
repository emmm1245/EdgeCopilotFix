// 执行地区配置修复，并使用 Copilot 兼容参数重新启动 Edge。
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
	if err := process.WaitForEdgeExit(); err != nil {
		logger.Error(err.Error())
		waitForExit()
		os.Exit(1)
	}

	fmt.Println()

	// 第二步：修复所有检测到的 Edge 版本配置文件
	fixedVersions, err := config.FixAllEdgeVersions()
	if err != nil {
		logger.Error(err.Error())
		waitForExit()
		os.Exit(1)
	}

	fmt.Println()
	startedCount := 0
	for _, edgeVersion := range fixedVersions {
		if err := process.StartEdge(edgeVersion); err != nil {
			logger.Warning(err.Error())
		} else {
			startedCount++
			logger.Success(fmt.Sprintf("已使用 Copilot 兼容参数启动 %s", edgeVersion.DisplayName))
		}
	}
	logger.Info("请检查 Copilot 侧边栏是否恢复；如果没有恢复，请重新运行本工具，多重试几次该操作。")
	logger.Info("Edge 155.0.4283.60 (正式版本) (64 位) 亲测有效。")
	logger.Info("兼容参数仅对本次启动生效；通过普通快捷方式启动后若再次消失，请重新运行本工具。")
	if startedCount == 0 {
		logger.Error("未能启动任何 Edge 版本，Copilot 兼容参数尚未生效。")
		waitForExit()
		os.Exit(1)
	}

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
