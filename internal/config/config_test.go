// 验证仅返回配置修复成功的 Edge 版本供后续兼容启动使用。
package config

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func TestFixAllEdgeVersionsReturnsSuccessfulVersions(t *testing.T) {
	if runtime.GOOS != "windows" {
		t.Skip("使用 Windows 用户数据目录隔离配置测试")
	}
	for _, country := range []string{"CN", "US"} {
		t.Run(country, func(t *testing.T) {
			t.Setenv("LOCALAPPDATA", t.TempDir())
			stableDir := filepath.Join(os.Getenv("LOCALAPPDATA"), "Microsoft", "Edge", "User Data")
			betaDir := filepath.Join(os.Getenv("LOCALAPPDATA"), "Microsoft", "Edge Beta", "User Data")
			for _, dir := range []string{stableDir, betaDir} {
				if err := os.MkdirAll(dir, 0755); err != nil {
					t.Fatal(err)
				}
			}
			stablePath := filepath.Join(stableDir, "Local State")
			if err := os.WriteFile(stablePath, []byte(`{"variations_country":"`+country+`","keep":true}`), 0644); err != nil {
				t.Fatal(err)
			}
			if err := os.WriteFile(filepath.Join(betaDir, "Local State"), []byte("invalid JSON"), 0644); err != nil {
				t.Fatal(err)
			}
			versions, err := FixAllEdgeVersions()
			if err != nil {
				t.Fatal(err)
			}
			if len(versions) != 1 || versions[0].UserDataPath != stableDir {
				t.Fatalf("应仅返回成功修复的稳定版，实际为 %v", versions)
			}
			saved, err := ReadConfig(stablePath)
			if err != nil {
				t.Fatal(err)
			}
			if saved["variations_country"] != "US" || saved["keep"] != true {
				t.Fatalf("地区应为 US 且保留其他配置，实际为 %v", saved)
			}
		})
	}
}
