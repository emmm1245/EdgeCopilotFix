# EdgeCopilotFix

修复 Microsoft Edge 浏览器中因地区限制或个人账号资格检查导致 Copilot 侧边栏入口消失的问题。

## 特性

- **跨平台支持**：支持 Windows、macOS 和 Linux
- **多版本支持**：可检测并修复 Stable、Beta、Dev、Canary 所有版本
- **新版兼容启动**：保留原有地区配置修复，并携带参数启动 Edge，关闭个人账号的 Copilot 资格检查

## 本次修复说明

**Edge 155.0.4283.60 (正式版本) (64 位) 亲测有效。** 此版本已实测恢复 Copilot 入口，其他版本和平台的本次兼容启动修复尚未实测。

如果没有恢复，请完全退出 Edge（包括后台进程），重新运行本工具，再多重试几次该操作。

本次修复使用 `--disable-features=EdgeCopilotMsaEligibility`【关闭个人账号的 Copilot 资格检查】启动 Edge。参数仅对本次启动生效；通过普通快捷方式启动后若入口再次消失，请重新运行本工具。本工具不修改浏览器快捷方式。

## 使用方法

### 方式一：直接运行可执行文件

1. 从 [Release](../../releases) 页面下载对应平台的程序：

   - Windows: `EdgeCopilotFixTool-windows-amd64.exe`
   - macOS (Intel): `EdgeCopilotFixTool-darwin-amd64`
   - macOS (Apple Silicon): `EdgeCopilotFixTool-darwin-arm64`
   - Linux: `EdgeCopilotFixTool-linux-amd64`
2. 运行程序：

   - **Windows**: 双击 `.exe` 文件
   - **macOS/Linux**: 在终端中运行
     ```bash
     chmod +x EdgeCopilotFixTool-*
     ./EdgeCopilotFixTool-*
     ```
3. 根据提示操作（如果检测到 Edge 进程会要求确认终止）
4. 工具会等待 Edge 完全退出、保存地区配置，再使用兼容参数启动修复成功的 Edge 版本
5. 检查 Copilot 侧边栏是否恢复；如果没有恢复，请重新运行本工具，多重试几次该操作

### 方式二：从源码编译

```bash
# 克隆仓库
git clone https://github.com/emmm1245/EdgeCopilotFix.git
cd EdgeCopilotFix

# 编译当前平台版本
go build -o EdgeCopilotFix cmd/edgecopilotfix/main.go

# 运行
./EdgeCopilotFix
```

### 跨平台编译

#### Windows 系统

```batch
# 编译所有平台版本
build.bat
```

#### Linux/macOS 系统

```bash
# 给脚本执行权限
chmod +x build.sh

# 编译所有平台版本
./build.sh
```

#### 手动指定平台编译

**Windows PowerShell:**

```powershell
# Windows (64位)
$env:GOOS="windows"; $env:GOARCH="amd64"; go build -o EdgeCopilotFixTool-windows-amd64.exe cmd/edgecopilotfix/main.go

# macOS (Intel)
$env:GOOS="darwin"; $env:GOARCH="amd64"; go build -o EdgeCopilotFixTool-darwin-amd64 cmd/edgecopilotfix/main.go

# macOS (Apple Silicon)
$env:GOOS="darwin"; $env:GOARCH="arm64"; go build -o EdgeCopilotFixTool-darwin-arm64 cmd/edgecopilotfix/main.go

# Linux (64位)
$env:GOOS="linux"; $env:GOARCH="amd64"; go build -o EdgeCopilotFixTool-linux-amd64 cmd/edgecopilotfix/main.go
```

**Linux/macOS Bash:**

```bash
# Windows (64位)
GOOS=windows GOARCH=amd64 go build -o EdgeCopilotFixTool-windows-amd64.exe cmd/edgecopilotfix/main.go

# macOS (Intel)
GOOS=darwin GOARCH=amd64 go build -o EdgeCopilotFixTool-darwin-amd64 cmd/edgecopilotfix/main.go

# macOS (Apple Silicon)
GOOS=darwin GOARCH=arm64 go build -o EdgeCopilotFixTool-darwin-arm64 cmd/edgecopilotfix/main.go

# Linux (64位)
GOOS=linux GOARCH=amd64 go build -o EdgeCopilotFixTool-linux-amd64 cmd/edgecopilotfix/main.go
```

## 支持的平台

| 操作系统      | 架构                  | 状态           |
| ------------- | --------------------- | -------------- |
| Windows 10/11 | amd64                 | 已测试         |
| macOS         | amd64 (Intel)         | 已实现，未测试 |
| macOS         | arm64 (Apple Silicon) | 已实现，未测试 |
| Linux         | amd64                 | 已实现，未测试 |

## 支持的 Edge 版本

- Edge Stable (稳定版)
- Edge Beta (测试版)
- Edge Dev (开发版)
- Edge Canary (金丝雀版)

程序会自动检测系统中所有已安装的 Edge 版本并逐个修复。

## 注意事项

- 使用前会提示终止 Edge 进程，请确保保存了重要的浏览数据
- 工具会使用兼容参数启动 Edge，请在此次启动的窗口中检查 Copilot 侧边栏
- 如果没有恢复，请完全退出 Edge 后重新运行本工具，多重试几次该操作
- 兼容参数仅对本次启动生效，通过普通快捷方式启动后若再次消失，请重新运行本工具
- macOS/Linux 用户需要给程序添加执行权限

## 系统要求

- Windows 10/11、macOS 或 Linux
- Microsoft Edge 浏览器（任意版本）
- 如果从源码编译，需要 Go 1.16 或更高版本

## License

MIT License
