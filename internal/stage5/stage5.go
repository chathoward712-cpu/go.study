package stage5

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// RunStage5 运行第5阶段演示
func RunStage5() {
	fmt.Println("第5阶段：模块化与工程实践")

	// 模块管理演示
	DemoModules()

	// 包管理演示
	DemoPackages()

	// 依赖管理演示
	DemoDependencies()

	// 测试演示
	DemoTesting()

	// 文档演示
	DemoDocumentation()

	// 构建部署演示
	DemoBuildDeploy()
}

// DemoModules 演示模块管理
func DemoModules() {
	fmt.Println("\n=== 模块管理演示 ===")

	// 1. Go模块基础
	fmt.Println("\n1. Go模块基础：")
	demoGoModBasics()

	// 2. 模块版本管理
	fmt.Println("\n2. 模块版本管理：")
	demoVersionManagement()

	// 3. 模块结构分析
	fmt.Println("\n3. 模块结构分析：")
	demoModuleStructure()

	// 4. 工作区模式
	fmt.Println("\n4. 工作区模式：")
	demoWorkspaceMode()
}

// demoGoModBasics 演示Go模块基础
func demoGoModBasics() {
	fmt.Println("Go模块系统基础概念:")

	// 读取当前项目的go.mod文件
	if content, err := os.ReadFile("go.mod"); err == nil {
		fmt.Println("\n当前项目的go.mod文件:")
		fmt.Println(string(content))
	} else {
		fmt.Printf("无法读取go.mod文件: %v\n", err)
	}

	// 展示模块路径概念
	fmt.Println("\n模块路径概念:")
	fmt.Println("- 模块路径是模块的唯一标识符")
	fmt.Println("- 通常是代码仓库的URL")
	fmt.Println("- 例如: github.com/howard/go.study")

	// 展示语义版本
	fmt.Println("\n语义版本控制:")
	fmt.Println("- 主版本号.次版本号.修订号 (例如: v1.2.3)")
	fmt.Println("- 主版本号: 不兼容的API修改")
	fmt.Println("- 次版本号: 向后兼容的功能性新增")
	fmt.Println("- 修订号: 向后兼容的问题修正")

	// 展示模块命令
	fmt.Println("\n常用模块命令:")
	demoModuleCommands()
}

// demoModuleCommands 演示模块命令
func demoModuleCommands() {
	commands := []struct {
		cmd  string
		desc string
	}{
		{"go mod init <module-path>", "初始化新模块"},
		{"go mod tidy", "添加缺失的模块，删除未使用的模块"},
		{"go mod download", "下载模块到本地缓存"},
		{"go mod verify", "验证依赖项的完整性"},
		{"go mod graph", "打印模块依赖图"},
		{"go mod why <package>", "解释为什么需要某个包"},
		{"go list -m all", "列出所有模块"},
		{"go list -m -versions <module>", "列出模块的可用版本"},
	}

	for _, cmd := range commands {
		fmt.Printf("  %-30s - %s\n", cmd.cmd, cmd.desc)
	}
}

// demoVersionManagement 演示版本管理
func demoVersionManagement() {
	fmt.Println("版本管理策略:")

	// 版本选择规则
	fmt.Println("\n版本选择规则:")
	fmt.Println("1. 最小版本选择 (Minimal Version Selection)")
	fmt.Println("2. 选择满足所有约束的最低版本")
	fmt.Println("3. 确保构建的可重现性")

	// 版本约束示例
	fmt.Println("\n版本约束示例:")
	constraints := []struct {
		constraint string
		meaning    string
	}{
		{"v1.2.3", "精确版本"},
		{">=v1.2.0", "大于等于指定版本"},
		{"<v2.0.0", "小于指定版本"},
		{"~v1.2.3", "补丁级别兼容 (>=v1.2.3, <v1.3.0)"},
		{"^v1.2.3", "次版本兼容 (>=v1.2.3, <v2.0.0)"},
	}

	for _, c := range constraints {
		fmt.Printf("  %-12s - %s\n", c.constraint, c.meaning)
	}

	// 主版本升级
	fmt.Println("\n主版本升级:")
	fmt.Println("- v0和v1: 导入路径不变")
	fmt.Println("- v2+: 导入路径需要包含版本后缀")
	fmt.Println("  例如: github.com/user/repo/v2")
}

// demoModuleStructure 演示模块结构
func demoModuleStructure() {
	fmt.Println("分析当前项目模块结构:")

	// 分析项目结构
	analyzeProjectStructure(".")

	// 展示推荐的项目布局
	fmt.Println("\n推荐的Go项目布局:")
	showRecommendedLayout()
}

// analyzeProjectStructure 分析项目结构
func analyzeProjectStructure(root string) {
	fmt.Printf("\n项目根目录: %s\n", root)

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 跳过隐藏文件和目录
		if strings.HasPrefix(info.Name(), ".") {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		// 跳过vendor目录
		if info.IsDir() && info.Name() == "vendor" {
			return filepath.SkipDir
		}

		// 计算相对路径
		relPath, _ := filepath.Rel(root, path)
		if relPath == "." {
			return nil
		}

		// 显示目录结构
		depth := strings.Count(relPath, string(filepath.Separator))
		indent := strings.Repeat("  ", depth)

		if info.IsDir() {
			fmt.Printf("%s[dir]  %s/\n", indent, info.Name())
		} else {
			fmt.Printf("%s[file] %s\n", indent, info.Name())
		}

		return nil
	})

	if err != nil {
		fmt.Printf("分析项目结构时出错: %v\n", err)
	}
}

// showRecommendedLayout 展示推荐的项目布局
func showRecommendedLayout() {
	layout := `
标准Go项目布局:
/
├── cmd/                    # 主应用程序
│   └── myapp/
│       └── main.go
├── internal/               # 私有应用程序和库代码
│   ├── app/
│   ├── pkg/
│   └── ...
├── pkg/                    # 外部应用程序可以使用的库代码
│   └── ...
├── api/                    # API定义文件
├── web/                    # Web应用程序特定的组件
├── configs/                # 配置文件模板或默认配置
├── init/                   # 系统初始化配置
├── scripts/                # 构建、安装、分析等脚本
├── build/                  # 打包和持续集成
├── deployments/            # 部署配置和模板
├── test/                   # 额外的外部测试应用程序和测试数据
├── docs/                   # 设计和用户文档
├── tools/                  # 项目的支持工具
├── examples/               # 应用程序或公共库的示例
├── third_party/            # 外部辅助工具、分叉代码和其他第三方工具
├── githooks/               # Git钩子
├── assets/                 # 与存储库一起使用的其他资产
├── website/                # 项目网站数据
├── README.md
├── LICENSE
├── Makefile
└── go.mod
`
	fmt.Println(layout)
}

// demoWorkspaceMode 演示工作区模式
func demoWorkspaceMode() {
	fmt.Println("Go工作区模式 (Go 1.18+):")

	fmt.Println("\n工作区的优势:")
	fmt.Println("- 同时开发多个相关模块")
	fmt.Println("- 本地替换远程依赖")
	fmt.Println("- 简化多模块项目的开发")

	fmt.Println("\n工作区命令:")
	workspaceCommands := []struct {
		cmd  string
		desc string
	}{
		{"go work init", "初始化工作区"},
		{"go work use <dir>", "添加模块到工作区"},
		{"go work edit", "编辑go.work文件"},
		{"go work sync", "同步工作区构建列表"},
	}

	for _, cmd := range workspaceCommands {
		fmt.Printf("  %-20s - %s\n", cmd.cmd, cmd.desc)
	}

	// 检查是否存在go.work文件
	if _, err := os.Stat("go.work"); err == nil {
		fmt.Println("\n当前目录存在go.work文件")
		if content, err := os.ReadFile("go.work"); err == nil {
			fmt.Println("go.work内容:")
			fmt.Println(string(content))
		}
	} else {
		fmt.Println("\n当前目录不存在go.work文件")
		fmt.Println("这是一个单模块项目")
	}
}

// DemoPackages 演示包管理
func DemoPackages() {
	fmt.Println("\n=== 包管理演示 ===")

	// 1. 包的基本概念
	fmt.Println("\n1. 包的基本概念：")
	demoPackageBasics()

	// 2. 包的可见性
	fmt.Println("\n2. 包的可见性：")
	demoPackageVisibility()

	// 3. 包的导入
	fmt.Println("\n3. 包的导入：")
	demoPackageImports()

	// 4. 包的初始化
	fmt.Println("\n4. 包的初始化：")
	demoPackageInit()

	// 5. 内部包
	fmt.Println("\n5. 内部包：")
	demoInternalPackages()
}

// demoPackageBasics 演示包的基本概念
func demoPackageBasics() {
	fmt.Println("Go包的基本概念:")
	fmt.Println("- 包是Go代码组织的基本单位")
	fmt.Println("- 每个Go文件都属于一个包")
	fmt.Println("- 包名通常与目录名相同")
	fmt.Println("- main包是特殊的，用于创建可执行程序")

	fmt.Println("\n包的命名规范:")
	fmt.Println("- 使用小写字母")
	fmt.Println("- 简短且有意义")
	fmt.Println("- 避免下划线和驼峰命名")
	fmt.Println("- 避免与标准库包名冲突")

	// 分析当前项目的包结构
	fmt.Println("\n当前项目的包结构:")
	analyzePackageStructure("internal")
}

// analyzePackageStructure 分析包结构
func analyzePackageStructure(root string) {
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(info.Name(), ".go") {
			// 解析Go文件获取包名
			if packageName := getPackageName(path); packageName != "" {
				relPath, _ := filepath.Rel(".", path)
				fmt.Printf("  %s -> package %s\n", relPath, packageName)
			}
		}

		return nil
	})

	if err != nil {
		fmt.Printf("分析包结构时出错: %v\n", err)
	}
}

// getPackageName 获取Go文件的包名
func getPackageName(filename string) string {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, filename, nil, parser.PackageClauseOnly)
	if err != nil {
		return ""
	}

	if file.Name != nil {
		return file.Name.Name
	}

	return ""
}

// demoPackageVisibility 演示包的可见性
func demoPackageVisibility() {
	fmt.Println("Go包的可见性规则:")
	fmt.Println("- 大写字母开头的标识符是导出的（公开的）")
	fmt.Println("- 小写字母开头的标识符是未导出的（私有的）")
	fmt.Println("- 只有导出的标识符可以被其他包访问")

	fmt.Println("\n可见性示例:")
	visibilityExamples := []struct {
		name       string
		visibility string
		accessible string
	}{
		{"PublicFunction", "导出", "其他包可访问"},
		{"privateFunction", "未导出", "仅包内访问"},
		{"PublicStruct", "导出", "其他包可访问"},
		{"privateStruct", "未导出", "仅包内访问"},
		{"PublicVar", "导出", "其他包可访问"},
		{"privateVar", "未导出", "仅包内访问"},
	}

	for _, example := range visibilityExamples {
		fmt.Printf("  %-20s - %-8s - %s\n", example.name, example.visibility, example.accessible)
	}

	fmt.Println("\n结构体字段的可见性:")
	fmt.Println(`type User struct {
    Name    string // 导出字段，其他包可访问
    age     int    // 未导出字段，仅包内访问
    Email   string // 导出字段，其他包可访问
}`)
}

// demoPackageImports 演示包的导入
func demoPackageImports() {
	fmt.Println("包的导入方式:")

	importTypes := []struct {
		syntax  string
		desc    string
		example string
	}{
		{
			"import \"package\"",
			"标准导入",
			"import \"fmt\"",
		},
		{
			"import alias \"package\"",
			"别名导入",
			"import f \"fmt\"",
		},
		{
			"import . \"package\"",
			"点导入（不推荐）",
			"import . \"fmt\"",
		},
		{
			"import _ \"package\"",
			"空白导入",
			"import _ \"database/sql\"",
		},
	}

	for _, imp := range importTypes {
		fmt.Printf("\n%s:\n", imp.desc)
		fmt.Printf("  语法: %s\n", imp.syntax)
		fmt.Printf("  示例: %s\n", imp.example)
	}

	fmt.Println("\n导入路径规则:")
	fmt.Println("- 标准库包：直接使用包名 (如 \"fmt\", \"os\")")
	fmt.Println("- 第三方包：完整的模块路径 (如 \"github.com/user/repo\")")
	fmt.Println("- 本地包：相对于模块根的路径")

	fmt.Println("\n导入分组建议:")
	fmt.Println(`import (
    // 标准库
    "fmt"
    "os"
    "strings"
    
    // 第三方库
    "github.com/gorilla/mux"
    "github.com/lib/pq"
    
    // 本地包
    "github.com/howard/go.study/internal/stage1"
    "github.com/howard/go.study/internal/stage2"
)`)
}

// demoPackageInit 演示包的初始化
func demoPackageInit() {
	fmt.Println("包的初始化过程:")
	fmt.Println("1. 导入包的依赖")
	fmt.Println("2. 初始化包级变量")
	fmt.Println("3. 执行init函数")
	fmt.Println("4. 每个包只初始化一次")

	fmt.Println("\ninit函数特点:")
	fmt.Println("- 无参数，无返回值")
	fmt.Println("- 一个包可以有多个init函数")
	fmt.Println("- 按照文件名字典序执行")
	fmt.Println("- 在main函数之前执行")

	fmt.Println("\ninit函数示例:")
	fmt.Println(`package config

import "log"

var AppConfig *Config

func init() {
    log.Println("初始化配置...")
    AppConfig = loadConfig()
}

func init() {
    log.Println("验证配置...")
    validateConfig(AppConfig)
}`)

	fmt.Println("\n初始化顺序:")
	fmt.Println("- 深度优先的依赖初始化")
	fmt.Println("- 同一包内按文件名排序")
	fmt.Println("- 同一文件内按出现顺序")
}

// demoInternalPackages 演示内部包
func demoInternalPackages() {
	fmt.Println("内部包 (internal) 的特殊性:")
	fmt.Println("- internal目录下的包只能被其父目录及子目录导入")
	fmt.Println("- 提供了包级别的访问控制")
	fmt.Println("- 防止外部包导入内部实现")

	fmt.Println("\n内部包示例结构:")
	fmt.Println(`project/
├── cmd/
│   └── app/
│       └── main.go          # 可以导入 project/internal/...
├── internal/
│   ├── auth/
│   │   └── auth.go          # 只能被 project/ 下的包导入
│   └── database/
│       └── db.go            # 只能被 project/ 下的包导入
├── pkg/
│   └── api/
│       └── api.go           # 可以导入 project/internal/...
└── third_party/
    └── external.go          # 不能导入 project/internal/...`)

	fmt.Println("\n内部包的优势:")
	fmt.Println("- 隐藏实现细节")
	fmt.Println("- 防止API滥用")
	fmt.Println("- 更好的模块化设计")
	fmt.Println("- 减少向后兼容性负担")

	// 检查当前项目是否使用了internal包
	if _, err := os.Stat("internal"); err == nil {
		fmt.Println("\n当前项目使用了internal包结构")
		fmt.Println("这是一个良好的实践，有助于代码组织和封装")
	}
}
