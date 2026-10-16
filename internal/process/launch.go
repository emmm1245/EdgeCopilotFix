// 等待 Edge 完全退出，并使用 Copilot 兼容参数启动浏览器。
package process

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/emmm1245/EdgeCopilotFix/internal/paths"
)

func WaitForEdgeExit() error {
	deadline := time.Now().Add(5 * time.Second)
	for {
		processes, err := CheckEdgeProcesses()
		if err != nil {
			return err
		}
		if len(processes) == 0 {
			return nil
		}
		if time.Now().After(deadline) {
			return fmt.Errorf("Edge 后台进程仍在运行，请完全退出 Edge 后重试")
		}
		time.Sleep(200 * time.Millisecond)
	}
}

func StartEdge(edgeVersion paths.EdgeVersion) error {
	executablePath, err := paths.GetExecutablePath(edgeVersion)
	if err != nil {
		return err
	}
	cmd := exec.Command(executablePath, "--user-data-dir="+edgeVersion.UserDataPath, "--disable-features=EdgeCopilotMsaEligibility")
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("启动 %s 失败: %w", edgeVersion.DisplayName, err)
	}
	return cmd.Process.Release()
}
