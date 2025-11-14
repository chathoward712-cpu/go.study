package stage5

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// DemoDependencies 演示依赖管理
func DemoDependencies() {
	fmt.Println("\n=== 依赖管理演示 ===")

	// 1. 依赖管理基础
	fmt.Println("\n1. 依赖管理基础：")
	demoDependencyBasics()

	// 2. 依赖版本控制
	fmt.Println("\n2. 依赖版本控制：")
	demoDependencyVersioning()

	// 3. 依赖解析
	fmt.Println("\n3. 依赖解析：")
	demoDependencyResolution()

	// 4. 依赖安全
	fmt.Println("\n4. 依赖安全：")
	demoDependencySecurity()

	// 5. 依赖优化
	fmt.Println("\n5. 依赖优化：")
	demoDependencyOptimization()
}

// demoDependencyBasics 演示依赖管理基础
func demoDependencyBasics() {
	fmt.Println("Go依赖管理的核心概念:")
	fmt.Println("- go.mod: 模块定义文件")
	fmt.Println("- go.sum: 依赖校验和文件")
	fmt.Println("- 模块缓存: $GOPATH/pkg/mod")
	fmt.Println("- 代理服务: GOPROXY环境变量")

	// 读取go.mod文件
	if content, err := os.ReadFile("go.mod"); err == nil {
		fmt.Println("\n当前项目的go.mod:")
		lines := strings.Split(string(content), "\n")
		for i, line := range lines {
			if i < 10 { // 只显示前10行
				fmt.Printf("  %s\n", line)
			}
		}
		if len(lines) > 10 {
			fmt.Printf("  ... (共%d行)\n", len(lines))
		}
	}

	// 读取go.sum文件
	if content, err := os.ReadFile("go.sum"); err == nil {
		fmt.Println("\ngo.sum文件存在，包含依赖校验和")
		lines := strings.Split(string(content), "\n")
		fmt.Printf("  共%d个校验和记录\n", len(lines)-1)
	} else {
		fmt.Println("\ngo.sum文件不存在（项目可能没有外部依赖）")
	}
}

// demoDependencyVersioning 演示依赖版本控制
func demoDependencyVersioning() {
	fmt.Println("依赖版本控制策略:")

	fmt.Println("\n语义版本控制 (SemVer):")
	fmt.Println("- MAJOR.MINOR.PATCH (例如: v1.2.3)")
	fmt.Println("- MAJOR: 不兼容的API变更")
	fmt.Println("- MINOR: 向后兼容的功能添加")
	fmt.Println("- PATCH: 向后兼容的错误修复")

	fmt.Println("\nGo模块版本规则:")
	versionRules := []struct {
		version string
		rule    string
	}{
		{"v0.x.x", "开发版本，API可能不稳定"},
		{"v1.x.x", "稳定版本，保证向后兼容"},
		{"v2+.x.x", "主版本升级，需要新的导入路径"},
		{"+incompatible", "非模块化的v2+版本"},
		{"pseudo-version", "基于commit的版本"},
	}

	for _, rule := range versionRules {
		fmt.Printf("  %-15s - %s\n", rule.version, rule.rule)
	}

	fmt.Println("\n版本选择示例:")
	fmt.Println("  go get github.com/user/repo@v1.2.3    # 精确版本")
	fmt.Println("  go get github.com/user/repo@latest    # 最新版本")
	fmt.Println("  go get github.com/user/repo@v1        # v1的最新版本")
	fmt.Println("  go get github.com/user/repo@master    # 特定分支")
	fmt.Println("  go get github.com/user/repo@commit    # 特定提交")
}

// demoDependencyResolution 演示依赖解析
func demoDependencyResolution() {
	fmt.Println("依赖解析机制:")

	fmt.Println("\n最小版本选择 (MVS):")
	fmt.Println("- 选择满足所有约束的最低版本")
	fmt.Println("- 确保构建的可重现性")
	fmt.Println("- 避免依赖地狱问题")

	fmt.Println("\n依赖解析过程:")
	steps := []string{
		"1. 读取go.mod文件中的直接依赖",
		"2. 递归解析间接依赖",
		"3. 应用最小版本选择算法",
		"4. 检查版本兼容性",
		"5. 生成最终的依赖图",
	}

	for _, step := range steps {
		fmt.Printf("  %s\n", step)
	}

	// 演示依赖图查看
	fmt.Println("\n查看依赖图的命令:")
	fmt.Println("  go mod graph                    # 显示依赖图")
	fmt.Println("  go list -m all                 # 列出所有依赖")
	fmt.Println("  go mod why <package>           # 解释为什么需要某个包")
	fmt.Println("  go list -m -versions <module>  # 列出模块的可用版本")

	// 尝试执行go list -m all命令
	if cmd := exec.Command("go", "list", "-m", "all"); cmd != nil {
		if output, err := cmd.Output(); err == nil {
			fmt.Println("\n当前项目的依赖列表:")
			lines := strings.Split(string(output), "\n")
			for i, line := range lines {
				if i < 5 && line != "" { // 只显示前5个依赖
					fmt.Printf("  %s\n", line)
				}
			}
			if len(lines) > 6 {
				fmt.Printf("  ... (共%d个依赖)\n", len(lines)-1)
			}
		}
	}
}

// demoDependencySecurity 演示依赖安全
func demoDependencySecurity() {
	fmt.Println("依赖安全管理:")

	fmt.Println("\n安全检查工具:")
	securityTools := []struct {
		tool string
		desc string
	}{
		{"go mod verify", "验证依赖的完整性"},
		{"go list -m -u all", "检查可更新的依赖"},
		{"govulncheck", "扫描已知漏洞"},
		{"go mod download", "预下载依赖到缓存"},
	}

	for _, tool := range securityTools {
		fmt.Printf("  %-20s - %s\n", tool.tool, tool.desc)
	}

	fmt.Println("\n校验和验证:")
	fmt.Println("- go.sum文件记录所有依赖的校验和")
	fmt.Println("- 防止依赖被篡改")
	fmt.Println("- 确保构建的一致性")

	fmt.Println("\n代理和镜像:")
	fmt.Println("- GOPROXY: 模块代理服务器")
	fmt.Println("- GOSUMDB: 校验和数据库")
	fmt.Println("- GOPRIVATE: 私有模块配置")

	fmt.Println("\n环境变量示例:")
	fmt.Println("  export GOPROXY=https://proxy.golang.org,direct")
	fmt.Println("  export GOSUMDB=sum.golang.org")
	fmt.Println("  export GOPRIVATE=*.corp.example.com")

	// 显示当前的Go环境变量
	if cmd := exec.Command("go", "env", "GOPROXY", "GOSUMDB"); cmd != nil {
		if output, err := cmd.Output(); err == nil {
			fmt.Println("\n当前Go环境配置:")
			lines := strings.Split(strings.TrimSpace(string(output)), "\n")
			if len(lines) >= 2 {
				fmt.Printf("  GOPROXY=%s\n", lines[0])
				fmt.Printf("  GOSUMDB=%s\n", lines[1])
			}
		}
	}
}

// demoDependencyOptimization 演示依赖优化
func demoDependencyOptimization() {
	fmt.Println("依赖优化策略:")

	fmt.Println("\n依赖清理:")
	fmt.Println("- go mod tidy: 添加缺失的依赖，移除未使用的依赖")
	fmt.Println("- go mod download: 预下载依赖")
	fmt.Println("- go clean -modcache: 清理模块缓存")

	fmt.Println("\n构建优化:")
	buildOptimizations := []struct {
		flag string
		desc string
	}{
		{"-mod=readonly", "只读模式，不修改go.mod"},
		{"-mod=vendor", "使用vendor目录"},
		{"-mod=mod", "允许修改go.mod（默认）"},
		{"-trimpath", "移除文件系统路径"},
		{"-ldflags", "链接器标志"},
	}

	for _, opt := range buildOptimizations {
		fmt.Printf("  %-15s - %s\n", opt.flag, opt.desc)
	}

	fmt.Println("\nVendor模式:")
	fmt.Println("- go mod vendor: 创建vendor目录")
	fmt.Println("- 将所有依赖复制到项目中")
	fmt.Println("- 适用于离线构建或严格控制依赖")

	fmt.Println("\n依赖分析:")
	fmt.Println("- go list -deps: 列出所有依赖包")
	fmt.Println("- go mod graph | grep <module>: 查找特定依赖")
	fmt.Println("- go list -m -json all: JSON格式的依赖信息")

	fmt.Println("\n最佳实践:")
	bestPractices := []string{
		"定期运行 go mod tidy",
		"使用固定版本而非latest",
		"定期更新依赖到安全版本",
		"监控依赖的安全漏洞",
		"避免过多的间接依赖",
		"使用go.mod的replace指令进行本地开发",
	}

	for i, practice := range bestPractices {
		fmt.Printf("  %d. %s\n", i+1, practice)
	}
}
