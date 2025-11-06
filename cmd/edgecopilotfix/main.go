package main

import (
	"fmt"
	"os"

	"github.com/emmm1245/EdgeCopilotFix/internal/config"
	"github.com/emmm1245/EdgeCopilotFix/internal/logger"
	"github.com/emmm1245/EdgeCopilotFix/internal/process"
)

func main() {
	// 输出欢迎信息
	fmt.Println("======================")
	fmt.Println("Edge Copilot Fix Tool")
	fmt.Println("======================")
	fmt.Println()

	// 第一步：处理 Edge 进程
	if err := process.HandleEdgeProcesses(); err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	fmt.Println()

	// 第二步：修复配置文件
	if err := config.FixCopilotConfig(); err != nil {
		// 如果是"无需修改"的特殊情况，显示信息但不作为错误
		if err.Error() == "配置已经是 US，无需修改" {
			logger.Info("配置已是最新状态。")
			fmt.Println()
			logger.Success("修复完成！")
		} else {
			logger.Error(err.Error())
			os.Exit(1)
		}
	} else {
		logger.Success("修复完成！")
		fmt.Println()
		logger.Info("请重新启动 Edge 浏览器。")
	}

	fmt.Println()
	fmt.Println("按 Enter 键退出...")
	fmt.Scanln()
}

