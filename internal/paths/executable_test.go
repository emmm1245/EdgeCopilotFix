// 验证 Windows 各渠道的 Edge 安装路径选择和缺失安装的处理。
package paths

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func TestGetExecutablePathWindows(t *testing.T) {
	if runtime.GOOS != "windows" {
		t.Skip("Windows 安装路径测试")
	}
	for _, channel := range []string{"Edge", "Edge Beta", "Edge Dev", "Edge SxS"} {
		t.Run(channel, func(t *testing.T) {
			root := t.TempDir()
			t.Setenv("ProgramFiles(x86)", filepath.Join(root, "Program Files (x86)"))
			t.Setenv("ProgramFiles", filepath.Join(root, "Program Files"))
			t.Setenv("LOCALAPPDATA", filepath.Join(root, "Local AppData"))
			version := EdgeVersion{
				UserDataPath: filepath.Join(os.Getenv("LOCALAPPDATA"), "Microsoft", channel, "User Data"),
				DisplayName:  channel,
			}
			if _, err := GetExecutablePath(version); err == nil {
				t.Fatal("没有安装文件时应返回错误")
			}
			for _, installRoot := range []string{os.Getenv("LOCALAPPDATA"), os.Getenv("ProgramFiles"), os.Getenv("ProgramFiles(x86)")} {
				want := filepath.Join(installRoot, "Microsoft", channel, "Application", "msedge.exe")
				if err := os.MkdirAll(filepath.Dir(want), 0755); err != nil {
					t.Fatal(err)
				}
				if err := os.WriteFile(want, nil, 0644); err != nil {
					t.Fatal(err)
				}
				got, err := GetExecutablePath(version)
				if err != nil || got != want {
					t.Fatalf("路径 = %q, 错误 = %v; 期望 %q", got, err, want)
				}
			}
		})
	}
}
