package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/emmm1245/EdgeCopilotFix/internal/logger"
	"github.com/emmm1245/EdgeCopilotFix/internal/paths"
)

// ReadConfig 读取配置文件
func ReadConfig(path string) (map[string]interface{}, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("读取配置文件失败: %v", err)
	}

	var config map[string]interface{}
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("解析 JSON 失败: %v", err)
	}

	return config, nil
}

// UpdateCountry 修改配置中的国家代码
// 返回 (是否已修改, 错误)
func UpdateCountry(config map[string]interface{}) (bool, error) {
	currentCountry, exists := config["variations_country"]
	if !exists {
		// 如果键不存在，添加它
		config["variations_country"] = "US"
		return true, nil
	}

	countryStr, ok := currentCountry.(string)
	if !ok {
		return false, fmt.Errorf("variations_country 值类型异常")
	}

	if countryStr == "US" {
		return false, nil // 已经是 US，无需修改
	}

	config["variations_country"] = "US"
	return true, nil
}

// SaveConfig 保存配置文件
func SaveConfig(path string, config map[string]interface{}) error {
	// 使用缩进格式化 JSON
	data, err := json.MarshalIndent(config, "", "   ")
	if err != nil {
		return fmt.Errorf("序列化 JSON 失败: %v", err)
	}

	// 写入文件
	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("写入配置文件失败: %v", err)
	}

	return nil
}

// FixEdgeVersion 修复指定版本的 Edge Copilot 配置
func FixEdgeVersion(edgeVersion paths.EdgeVersion) error {
	logger.Info(fmt.Sprintf("正在修复 %s...", edgeVersion.DisplayName))
	
	// 获取配置文件路径
	configPath := paths.GetLocalStatePath(edgeVersion.UserDataPath)
	
	// 检查文件是否存在
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return fmt.Errorf("配置文件不存在: %s", configPath)
	}

	// 读取配置
	config, err := ReadConfig(configPath)
	if err != nil {
		return err
	}

	// 检查并更新国家代码
	_, err = UpdateCountry(config)
	if err != nil {
		return err
	}

	// 保存配置
	if err := SaveConfig(configPath, config); err != nil {
		return err
	}

	logger.Success(fmt.Sprintf("%s 修复完成！", edgeVersion.DisplayName))
	return nil
}

// FixAllEdgeVersions 修复所有检测到的 Edge 版本
func FixAllEdgeVersions() error {
	// 获取所有 Edge 路径
	edgeVersions, err := paths.GetAllEdgePaths()
	if err != nil {
		return err
	}

	logger.Info(fmt.Sprintf("检测到 %d 个 Edge 版本:", len(edgeVersions)))
	for i, ev := range edgeVersions {
		fmt.Printf("  %d. %s\n", i+1, ev.DisplayName)
	}
	fmt.Println()

	// 逐个修复
	successCount := 0
	for _, ev := range edgeVersions {
		if err := FixEdgeVersion(ev); err != nil {
			logger.Warning(fmt.Sprintf("%s 修复失败: %v", ev.DisplayName, err))
		} else {
			successCount++
		}
		fmt.Println()
	}

	if successCount == 0 {
		return fmt.Errorf("所有版本修复失败")
	}

	logger.Success(fmt.Sprintf("成功修复 %d/%d 个 Edge 版本", successCount, len(edgeVersions)))
	return nil
}
