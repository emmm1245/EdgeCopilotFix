// 定位各平台、各渠道的 Edge 浏览器可执行文件。
package paths

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func GetExecutablePath(edgeVersion EdgeVersion) (string, error) {
	var candidates []string
	switch runtime.GOOS {
	case "windows":
		channelDir := filepath.Base(filepath.Dir(edgeVersion.UserDataPath))
		for _, root := range []string{os.Getenv("ProgramFiles(x86)"), os.Getenv("ProgramFiles"), os.Getenv("LOCALAPPDATA")} {
			if root != "" {
				candidates = append(candidates, filepath.Join(root, "Microsoft", channelDir, "Application", "msedge.exe"))
			}
		}
	case "darwin":
		appName := map[string]string{
			"stable": "Microsoft Edge", "beta": "Microsoft Edge Beta",
			"dev": "Microsoft Edge Dev", "canary": "Microsoft Edge Canary",
		}[edgeVersion.Version]
		roots := []string{"/Applications"}
		if home, err := os.UserHomeDir(); err == nil {
			roots = append(roots, filepath.Join(home, "Applications"))
		}
		for _, root := range roots {
			candidates = append(candidates, filepath.Join(root, appName+".app", "Contents", "MacOS", appName))
		}
	case "linux":
		names := map[string][]string{
			"stable": {"microsoft-edge-stable", "microsoft-edge"},
			"beta":   {"microsoft-edge-beta"},
			"dev":    {"microsoft-edge-dev"},
			"canary": {"microsoft-edge-canary"},
		}[edgeVersion.Version]
		for _, name := range names {
			if path, err := exec.LookPath(name); err == nil {
				return path, nil
			}
		}
	}

	for _, path := range candidates {
		if info, err := os.Stat(path); err == nil && !info.IsDir() {
			return path, nil
		}
	}
	return "", fmt.Errorf("未找到 %s 的可执行文件，请确认浏览器已安装", edgeVersion.DisplayName)
}
