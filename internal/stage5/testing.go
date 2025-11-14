package stage5

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

// DemoTesting 演示测试
func DemoTesting() {
	fmt.Println("\n=== 测试演示 ===")

	// 1. 单元测试基础
	fmt.Println("\n1. 单元测试基础：")
	demoUnitTestBasics()

	// 2. 表格驱动测试
	fmt.Println("\n2. 表格驱动测试：")
	demoTableDrivenTests()

	// 3. 基准测试
	fmt.Println("\n3. 基准测试：")
	demoBenchmarkTests()

	// 4. 示例测试
	fmt.Println("\n4. 示例测试：")
	demoExampleTests()

	// 5. 测试覆盖率
	fmt.Println("\n5. 测试覆盖率：")
	demoTestCoverage()

	// 6. 测试最佳实践
	fmt.Println("\n6. 测试最佳实践：")
	demoTestingBestPractices()
}

// demoUnitTestBasics 演示单元测试基础
func demoUnitTestBasics() {
	fmt.Println("Go单元测试基础:")
	fmt.Println("- 测试文件以 _test.go 结尾")
	fmt.Println("- 测试函数以 Test 开头")
	fmt.Println("- 测试函数接受 *testing.T 参数")
	fmt.Println("- 使用 go test 命令运行测试")

	fmt.Println("\n基本测试示例:")
	testExample := `package math

import "testing"

// 被测试的函数
func Add(a, b int) int {
    return a + b
}

// 测试函数
func TestAdd(t *testing.T) {
    result := Add(2, 3)
    expected := 5
    
    if result != expected {
        t.Errorf("Add(2, 3) = %d; want %d", result, expected)
    }
}

// 测试多个用例
func TestAddMultiple(t *testing.T) {
    tests := []struct {
        a, b, want int
    }{
        {1, 2, 3},
        {0, 0, 0},
        {-1, 1, 0},
        {10, -5, 5},
    }
    
    for _, tt := range tests {
        if got := Add(tt.a, tt.b); got != tt.want {
            t.Errorf("Add(%d, %d) = %d; want %d", 
                tt.a, tt.b, got, tt.want)
        }
    }
}`
	fmt.Println(testExample)

	fmt.Println("\n常用测试方法:")
	testMethods := []struct {
		method string
		desc   string
	}{
		{"t.Error()", "记录错误但继续执行"},
		{"t.Errorf()", "格式化错误信息"},
		{"t.Fatal()", "记录错误并停止测试"},
		{"t.Fatalf()", "格式化错误信息并停止"},
		{"t.Log()", "记录日志信息"},
		{"t.Logf()", "格式化日志信息"},
		{"t.Skip()", "跳过测试"},
		{"t.Skipf()", "格式化跳过信息"},
	}

	for _, method := range testMethods {
		fmt.Printf("  %-12s - %s\n", method.method, method.desc)
	}
}

// demoTableDrivenTests 演示表格驱动测试
func demoTableDrivenTests() {
	fmt.Println("表格驱动测试模式:")
	fmt.Println("- 使用结构体切片定义测试用例")
	fmt.Println("- 循环执行所有测试用例")
	fmt.Println("- 便于添加新的测试用例")
	fmt.Println("- 测试逻辑清晰，数据与逻辑分离")

	fmt.Println("\n表格驱动测试示例:")
	tableDrivenExample := `func TestStringLength(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        expected int
    }{
        {"empty string", "", 0},
        {"single char", "a", 1},
        {"normal string", "hello", 5},
        {"unicode string", "你好", 2},
        {"mixed string", "hello世界", 7},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := len([]rune(tt.input)); got != tt.expected {
                t.Errorf("len(%q) = %d; want %d", 
                    tt.input, got, tt.expected)
            }
        })
    }
}`
	fmt.Println(tableDrivenExample)

	fmt.Println("\n子测试的优势:")
	fmt.Println("- 使用 t.Run() 创建子测试")
	fmt.Println("- 每个用例独立运行")
	fmt.Println("- 可以单独运行特定用例")
	fmt.Println("- 更好的错误报告")
	fmt.Println("- 支持并行测试")
}

// demoBenchmarkTests 演示基准测试
func demoBenchmarkTests() {
	fmt.Println("基准测试 (Benchmark):")
	fmt.Println("- 函数名以 Benchmark 开头")
	fmt.Println("- 接受 *testing.B 参数")
	fmt.Println("- 使用 go test -bench 运行")
	fmt.Println("- 测量函数执行性能")

	fmt.Println("\n基准测试示例:")
	benchmarkExample := `func BenchmarkStringConcat(b *testing.B) {
    for i := 0; i < b.N; i++ {
        var s string
        for j := 0; j < 100; j++ {
            s += "hello"
        }
    }
}

func BenchmarkStringBuilder(b *testing.B) {
    for i := 0; i < b.N; i++ {
        var sb strings.Builder
        for j := 0; j < 100; j++ {
            sb.WriteString("hello")
        }
        _ = sb.String()
    }
}

func BenchmarkStringJoin(b *testing.B) {
    strs := make([]string, 100)
    for i := range strs {
        strs[i] = "hello"
    }
    
    b.ResetTimer() // 重置计时器
    for i := 0; i < b.N; i++ {
        _ = strings.Join(strs, "")
    }
}`
	fmt.Println(benchmarkExample)

	fmt.Println("\n基准测试命令:")
	benchCommands := []struct {
		cmd  string
		desc string
	}{
		{"go test -bench=.", "运行所有基准测试"},
		{"go test -bench=BenchmarkFunc", "运行特定基准测试"},
		{"go test -bench=. -benchmem", "显示内存分配统计"},
		{"go test -bench=. -count=5", "运行5次取平均值"},
		{"go test -bench=. -benchtime=10s", "运行10秒"},
		{"go test -bench=. -cpu=1,2,4", "指定CPU核数"},
	}

	for _, cmd := range benchCommands {
		fmt.Printf("  %-30s - %s\n", cmd.cmd, cmd.desc)
	}

	fmt.Println("\n基准测试结果解读:")
	fmt.Println("  BenchmarkFunc-8    1000000    1234 ns/op    456 B/op    7 allocs/op")
	fmt.Println("  ├─ 函数名-CPU核数")
	fmt.Println("  ├─ 执行次数")
	fmt.Println("  ├─ 每次操作耗时")
	fmt.Println("  ├─ 每次操作分配字节数")
	fmt.Println("  └─ 每次操作分配次数")
}

// demoExampleTests 演示示例测试
func demoExampleTests() {
	fmt.Println("示例测试 (Example):")
	fmt.Println("- 函数名以 Example 开头")
	fmt.Println("- 包含 // Output: 注释")
	fmt.Println("- 既是测试也是文档")
	fmt.Println("- 会出现在 godoc 中")

	fmt.Println("\n示例测试示例:")
	exampleTest := `func ExampleAdd() {
    result := Add(2, 3)
    fmt.Println(result)
    // Output: 5
}

func ExampleAdd_negative() {
    result := Add(-1, -2)
    fmt.Println(result)
    // Output: -3
}

func ExampleStringReverse() {
    s := "hello"
    reversed := Reverse(s)
    fmt.Printf("Original: %s, Reversed: %s", s, reversed)
    // Output: Original: hello, Reversed: olleh
}`
	fmt.Println(exampleTest)

	fmt.Println("\n示例测试的特点:")
	fmt.Println("- 验证输出是否与期望一致")
	fmt.Println("- 可以包含多行输出")
	fmt.Println("- 支持无序输出 (// Unordered output:)")
	fmt.Println("- 可以测试包级别的示例")
	fmt.Println("- 自动包含在文档中")
}

// demoTestCoverage 演示测试覆盖率
func demoTestCoverage() {
	fmt.Println("测试覆盖率分析:")
	fmt.Println("- 衡量测试的完整性")
	fmt.Println("- 识别未测试的代码")
	fmt.Println("- 帮助提高代码质量")

	fmt.Println("\n覆盖率命令:")
	coverageCommands := []struct {
		cmd  string
		desc string
	}{
		{"go test -cover", "显示覆盖率百分比"},
		{"go test -coverprofile=cover.out", "生成覆盖率文件"},
		{"go tool cover -html=cover.out", "生成HTML覆盖率报告"},
		{"go tool cover -func=cover.out", "按函数显示覆盖率"},
		{"go test -covermode=count", "统计执行次数"},
		{"go test -coverpkg=./...", "包含所有包的覆盖率"},
	}

	for _, cmd := range coverageCommands {
		fmt.Printf("  %-35s - %s\n", cmd.cmd, cmd.desc)
	}

	fmt.Println("\n覆盖率模式:")
	coverageModes := []struct {
		mode string
		desc string
	}{
		{"set", "是否执行过（默认）"},
		{"count", "执行次数"},
		{"atomic", "原子计数（并发安全）"},
	}

	for _, mode := range coverageModes {
		fmt.Printf("  %-8s - %s\n", mode.mode, mode.desc)
	}

	fmt.Println("\n覆盖率最佳实践:")
	fmt.Println("- 目标覆盖率通常在80-90%")
	fmt.Println("- 100%覆盖率不一定意味着完美测试")
	fmt.Println("- 关注关键业务逻辑的覆盖")
	fmt.Println("- 结合代码审查和静态分析")
}

// demoTestingBestPractices 演示测试最佳实践
func demoTestingBestPractices() {
	fmt.Println("测试最佳实践:")

	fmt.Println("\n1. 测试命名:")
	fmt.Println("- 使用描述性的测试名称")
	fmt.Println("- 包含被测试的功能和场景")
	fmt.Println("- 例如: TestUserService_CreateUser_WithValidData")

	fmt.Println("\n2. 测试结构 (AAA模式):")
	fmt.Println("- Arrange: 准备测试数据和环境")
	fmt.Println("- Act: 执行被测试的操作")
	fmt.Println("- Assert: 验证结果")

	fmt.Println("\n3. 测试隔离:")
	fmt.Println("- 每个测试应该独立")
	fmt.Println("- 不依赖其他测试的执行顺序")
	fmt.Println("- 使用 setup 和 teardown")

	fmt.Println("\n4. 测试数据:")
	fmt.Println("- 使用有意义的测试数据")
	fmt.Println("- 避免魔法数字")
	fmt.Println("- 考虑边界条件")

	fmt.Println("\n5. 错误测试:")
	fmt.Println("- 测试正常路径和异常路径")
	fmt.Println("- 验证错误类型和错误消息")
	fmt.Println("- 使用 testify 等断言库")

	fmt.Println("\n6. 并发测试:")
	fmt.Println("- 使用 t.Parallel() 并行执行")
	fmt.Println("- 注意共享状态的竞态条件")
	fmt.Println("- 使用 race detector")

	fmt.Println("\n7. 测试工具:")
	testingTools := []struct {
		tool string
		desc string
	}{
		{"testify", "断言和模拟库"},
		{"gomock", "生成模拟对象"},
		{"ginkgo", "BDD测试框架"},
		{"httptest", "HTTP测试工具"},
		{"goleak", "检测goroutine泄漏"},
	}

	for _, tool := range testingTools {
		fmt.Printf("  %-12s - %s\n", tool.tool, tool.desc)
	}

	// 尝试运行项目中的测试
	fmt.Println("\n运行当前项目的测试:")
	runProjectTests()
}

// runProjectTests 运行项目测试
func runProjectTests() {
	// 查找测试文件
	testFiles := findTestFiles(".")
	if len(testFiles) == 0 {
		fmt.Println("  当前项目中没有找到测试文件")
		fmt.Println("  建议为关键功能添加单元测试")
		return
	}

	fmt.Printf("  找到 %d 个测试文件:\n", len(testFiles))
	for _, file := range testFiles {
		fmt.Printf("    %s\n", file)
	}

	// 尝试运行测试
	fmt.Println("\n  尝试运行测试...")
	if cmd := exec.Command("go", "test", "./..."); cmd != nil {
		cmd.Dir = "."
		if output, err := cmd.CombinedOutput(); err != nil {
			fmt.Printf("  测试执行失败: %v\n", err)
			if len(output) > 0 {
				fmt.Printf("  输出: %s\n", string(output))
			}
		} else {
			fmt.Printf("  测试结果:\n%s", string(output))
		}
	}
}

// findTestFiles 查找测试文件
func findTestFiles(root string) []string {
	var testFiles []string

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 跳过vendor和隐藏目录
		if info.IsDir() && (strings.HasPrefix(info.Name(), ".") || info.Name() == "vendor") {
			return filepath.SkipDir
		}

		// 查找测试文件
		if !info.IsDir() && strings.HasSuffix(info.Name(), "_test.go") {
			relPath, _ := filepath.Rel(root, path)
			testFiles = append(testFiles, relPath)
		}

		return nil
	})

	if err != nil {
		fmt.Printf("查找测试文件时出错: %v\n", err)
	}

	return testFiles
}
