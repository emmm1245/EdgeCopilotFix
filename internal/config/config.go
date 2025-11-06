package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// GetLocalStatePath 获取 Edge Local State 配置文件路径
func GetLocalStatePath() (string, error) {
	localAppData := os.Getenv("LOCALAPPDATA")
	if localAppData == "" {
		return "", fmt.Errorf("无法获取 LOCALAPPDATA 环境变量")
	}

	localStatePath := filepath.Join(localAppData, "Microsoft", "Edge", "User Data", "Local State")
	return localStatePath, nil
}

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

// FixCopilotConfig 修复 Copilot 配置的完整流程
func FixCopilotConfig() error {
	// 获取配置文件路径
	configPath, err := GetLocalStatePath()
	if err != nil {
		return err
	}

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
	modified, err := UpdateCountry(config)
	if err != nil {
		return err
	}

	if !modified {
		return fmt.Errorf("配置已经是 US，无需修改")
	}

	// 保存配置
	if err := SaveConfig(configPath, config); err != nil {
		return err
	}

	return nil
}

