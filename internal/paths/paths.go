package paths

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

// EdgeVersion 表示一个 Edge 版本及其数据路径
type EdgeVersion struct {
	Version      string // stable, canary, dev, beta
	UserDataPath string
	DisplayName  string // 用于显示的名称
}

// GetAllEdgePaths 获取所有可用的 Edge 数据路径
func GetAllEdgePaths() ([]EdgeVersion, error) {
	var paths []EdgeVersion

	switch runtime.GOOS {
	case "windows":
		paths = getWindowsPaths()
	case "darwin":
		paths = getDarwinPaths()
	case "linux":
		paths = getLinuxPaths()
	default:
		return nil, fmt.Errorf("不支持的操作系统: %s", runtime.GOOS)
	}

	// 过滤出实际存在的路径
	var existingPaths []EdgeVersion
	for _, p := range paths {
		if _, err := os.Stat(p.UserDataPath); err == nil {
			existingPaths = append(existingPaths, p)
		}
	}

	if len(existingPaths) == 0 {
		return nil, fmt.Errorf("未找到任何 Edge 浏览器数据目录")
	}

	return existingPaths, nil
}

// getWindowsPaths 获取 Windows 平台的路径
func getWindowsPaths() []EdgeVersion {
	localAppData := os.Getenv("LOCALAPPDATA")
	if localAppData == "" {
		return nil
	}

	return []EdgeVersion{
		{
			Version:      "stable",
			UserDataPath: filepath.Join(localAppData, "Microsoft", "Edge", "User Data"),
			DisplayName:  "Edge Stable",
		},
		{
			Version:      "beta",
			UserDataPath: filepath.Join(localAppData, "Microsoft", "Edge Beta", "User Data"),
			DisplayName:  "Edge Beta",
		},
		{
			Version:      "dev",
			UserDataPath: filepath.Join(localAppData, "Microsoft", "Edge Dev", "User Data"),
			DisplayName:  "Edge Dev",
		},
		{
			Version:      "canary",
			UserDataPath: filepath.Join(localAppData, "Microsoft", "Edge SxS", "User Data"),
			DisplayName:  "Edge Canary",
		},
	}
}

// getDarwinPaths 获取 macOS 平台的路径
func getDarwinPaths() []EdgeVersion {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil
	}

	return []EdgeVersion{
		{
			Version:      "stable",
			UserDataPath: filepath.Join(homeDir, "Library", "Application Support", "Microsoft Edge"),
			DisplayName:  "Edge Stable",
		},
		{
			Version:      "beta",
			UserDataPath: filepath.Join(homeDir, "Library", "Application Support", "Microsoft Edge Beta"),
			DisplayName:  "Edge Beta",
		},
		{
			Version:      "dev",
			UserDataPath: filepath.Join(homeDir, "Library", "Application Support", "Microsoft Edge Dev"),
			DisplayName:  "Edge Dev",
		},
		{
			Version:      "canary",
			UserDataPath: filepath.Join(homeDir, "Library", "Application Support", "Microsoft Edge Canary"),
			DisplayName:  "Edge Canary",
		},
	}
}

// getLinuxPaths 获取 Linux 平台的路径
func getLinuxPaths() []EdgeVersion {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil
	}

	return []EdgeVersion{
		{
			Version:      "stable",
			UserDataPath: filepath.Join(homeDir, ".config", "microsoft-edge"),
			DisplayName:  "Edge Stable",
		},
		{
			Version:      "beta",
			UserDataPath: filepath.Join(homeDir, ".config", "microsoft-edge-beta"),
			DisplayName:  "Edge Beta",
		},
		{
			Version:      "dev",
			UserDataPath: filepath.Join(homeDir, ".config", "microsoft-edge-dev"),
			DisplayName:  "Edge Dev",
		},
		{
			Version:      "canary",
			UserDataPath: filepath.Join(homeDir, ".config", "microsoft-edge-canary"),
			DisplayName:  "Edge Canary",
		},
	}
}

// GetLocalStatePath 获取指定用户数据目录的 Local State 文件路径
func GetLocalStatePath(userDataPath string) string {
	return filepath.Join(userDataPath, "Local State")
}

// GetLastVersionPath 获取 Last Version 文件路径
func GetLastVersionPath(userDataPath string) string {
	return filepath.Join(userDataPath, "Last Version")
}

