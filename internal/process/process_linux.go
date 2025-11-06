//go:build linux

package process

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/emmm1245/EdgeCopilotFix/internal/logger"
)

// CheckEdgeProcesses 检查是否有 Edge 进程正在运行
func CheckEdgeProcesses() ([]string, error) {
	// 在 Linux 上，Edge 进程名称通常是 msedge
	cmd := exec.Command("pgrep", "-l", "msedge")
	output, err := cmd.Output()
	
	// 如果没有找到进程，pgrep 会返回错误，这是正常的
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			if exitErr.ExitCode() == 1 {
				// 退出码 1 表示没有找到进程
				return []string{}, nil
			}
		}
		return nil, fmt.Errorf("检测进程失败: %v", err)
	}

	outputStr := string(output)
	var processes []string

	// 解析输出，每行格式为 "PID 进程名称"
	lines := strings.Split(outputStr, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" {
			processes = append(processes, line)
		}
	}

	return processes, nil
}

// KillEdgeProcesses 终止所有 Edge 进程
func KillEdgeProcesses() error {
	// 使用 pkill 终止所有 msedge 进程
	cmd := exec.Command("pkill", "msedge")
	err := cmd.Run()
	
	// pkill 如果没找到进程会返回退出码 1，这不算错误
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			if exitErr.ExitCode() == 1 {
				// 没有进程被终止，但不是错误
				return nil
			}
		}
		return fmt.Errorf("终止进程失败: %v", err)
	}
	
	return nil
}

// AskUserConfirmation 询问用户是否确认终止进程
func AskUserConfirmation() bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("检测到Edge进程正在运行，是否要终止所有Edge进程？(y/n): \n")
	fmt.Print("\033[31m请确保保存了重要的浏览数据\033[0m\n")
	
	response, err := reader.ReadString('\n')
	if err != nil {
		logger.Error(fmt.Sprintf("读取输入失败: %v", err))
		return false
	}

	response = strings.TrimSpace(strings.ToLower(response))
	return response == "y" || response == "yes"
}

// HandleEdgeProcesses 处理 Edge 进程（检测、询问、终止）
func HandleEdgeProcesses() error {
	logger.Info("检测 Edge 进程...")
	
	processes, err := CheckEdgeProcesses()
	if err != nil {
		return err
	}

	if len(processes) == 0 {
		logger.Info("未检测到运行中的 Edge 进程。")
		return nil
	}

	logger.Warning(fmt.Sprintf("检测到 %d 个 Edge 进程", len(processes)))
	for i, proc := range processes {
		fmt.Printf("  %d. %s\n", i+1, proc)
	}
	fmt.Println()

	if !AskUserConfirmation() {
		return fmt.Errorf("用户取消操作")
	}

	logger.Info("终止 Edge 进程...")
	if err := KillEdgeProcesses(); err != nil {
		return err
	}

	logger.Success("已终止所有 Edge 进程。")
	return nil
}

