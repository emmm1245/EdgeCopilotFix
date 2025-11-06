# EdgeCopilotFix

自动修复因地区限制导致 Microsoft Edge 浏览器中 Copilot 功能不可用的问题。

## 使用方法

### 方式一：直接运行可执行文件

1. 从 Release 页面下载 `EdgeCopilotFix.exe`
2. 双击运行程序
3. 根据提示操作（如果检测到 Edge 进程会要求确认终止）
4. 等待修复完成后重新启动 Edge 浏览器

### 方式二：从源码编译

```bash
# 克隆仓库
git clone https://github.com/emmm1245/EdgeCopilotFix.git
cd EdgeCopilotFix

# 编译程序
go build -o EdgeCopilotFix.exe ./cmd/edgecopilotfix

# 运行
./EdgeCopilotFix.exe
```

## 注意事项

- 本工具目前仅适用于 Windows 系统
- 使用前会提示终止 Edge 进程，请确保保存了重要的浏览数据
- 修复完成后需要重新启动 Edge 浏览器才能生效

## 系统要求

- Windows 10/11
- Microsoft Edge 浏览器

## License

MIT License
