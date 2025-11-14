package stage5

import (
	"fmt"
	"go/ast"
	"go/doc"
	"go/parser"
	"go/token"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// DemoDocumentation 演示文档
func DemoDocumentation() {
	fmt.Println("\n=== 文档演示 ===")

	// 1. Go文档基础
	fmt.Println("\n1. Go文档基础：")
	demoDocumentationBasics()

	// 2. 文档注释规范
	fmt.Println("\n2. 文档注释规范：")
	demoDocumentationComments()

	// 3. godoc工具
	fmt.Println("\n3. godoc工具：")
	demoGodocTool()

	// 4. 文档示例
	fmt.Println("\n4. 文档示例：")
	demoDocumentationExamples()

	// 5. 文档最佳实践
	fmt.Println("\n5. 文档最佳实践：")
	demoDocumentationBestPractices()
}

// demoDocumentationBasics 演示文档基础
func demoDocumentationBasics() {
	fmt.Println("Go文档系统特点:")
	fmt.Println("- 文档即代码，与源码紧密结合")
	fmt.Println("- 使用注释生成文档")
	fmt.Println("- 支持HTML和文本格式")
	fmt.Println("- 自动提取示例代码")

	fmt.Println("\n文档类型:")
	docTypes := []struct {
		docType string
		desc    string
	}{
		{"包文档", "package声明前的注释"},
		{"函数文档", "函数声明前的注释"},
		{"类型文档", "类型声明前的注释"},
		{"变量文档", "变量声明前的注释"},
		{"常量文档", "常量声明前的注释"},
		{"示例文档", "Example函数的输出"},
	}

	for _, docType := range docTypes {
		fmt.Printf("  %-8s - %s\n", docType.docType, docType.desc)
	}

	fmt.Println("\n文档工具:")
	fmt.Println("- go doc: 命令行文档查看")
	fmt.Println("- godoc: Web服务器文档")
	fmt.Println("- pkg.go.dev: 在线文档平台")
	fmt.Println("- IDE集成: 编辑器内文档显示")
}

// demoDocumentationComments 演示文档注释规范
func demoDocumentationComments() {
	fmt.Println("文档注释规范:")

	fmt.Println("\n1. 包文档:")
	packageDocExample := `// Package calculator provides basic arithmetic operations.
//
// This package implements addition, subtraction, multiplication,
// and division operations for integers and floating-point numbers.
//
// Example usage:
//
//	result := calculator.Add(2, 3)
//	fmt.Println(result) // Output: 5
//
package calculator`
	fmt.Println(packageDocExample)

	fmt.Println("\n2. 函数文档:")
	functionDocExample := `// Add returns the sum of two integers.
//
// It takes two integer parameters and returns their sum.
// This function handles integer overflow by returning the
// mathematical result without error checking.
//
// Example:
//
//	sum := Add(10, 20)
//	fmt.Println(sum) // Output: 30
//
func Add(a, b int) int {
    return a + b
}`
	fmt.Println(functionDocExample)

	fmt.Println("\n3. 类型文档:")
	typeDocExample := `// User represents a user in the system.
//
// A User contains basic information such as name, email,
// and creation timestamp. All fields are required except
// for the optional LastLogin field.
type User struct {
    // Name is the user's full name
    Name string
    
    // Email is the user's email address
    Email string
    
    // CreatedAt is when the user was created
    CreatedAt time.Time
    
    // LastLogin is the last login time (optional)
    LastLogin *time.Time
}`
	fmt.Println(typeDocExample)

	fmt.Println("\n文档注释规则:")
	rules := []string{
		"以被文档化的标识符名称开头",
		"使用完整的句子",
		"第一句话应该是简洁的摘要",
		"使用现在时态",
		"避免冗余信息",
		"包含使用示例",
		"解释参数和返回值",
		"说明错误条件",
	}

	for i, rule := range rules {
		fmt.Printf("  %d. %s\n", i+1, rule)
	}
}

// demoGodocTool 演示godoc工具
func demoGodocTool() {
	fmt.Println("godoc工具使用:")

	fmt.Println("\ngo doc命令:")
	goDocCommands := []struct {
		cmd  string
		desc string
	}{
		{"go doc", "显示当前包的文档"},
		{"go doc fmt", "显示fmt包的文档"},
		{"go doc fmt.Println", "显示特定函数的文档"},
		{"go doc -all fmt", "显示包的所有文档"},
		{"go doc -short fmt", "显示简短文档"},
		{"go doc -u fmt", "包含未导出的标识符"},
	}

	for _, cmd := range goDocCommands {
		fmt.Printf("  %-20s - %s\n", cmd.cmd, cmd.desc)
	}

	fmt.Println("\ngodoc服务器:")
	fmt.Println("  godoc -http=:6060        # 启动本地文档服务器")
	fmt.Println("  访问 http://localhost:6060 查看文档")

	fmt.Println("\ngo doc示例 - 查看fmt包:")
	if cmd := exec.Command("go", "doc", "fmt"); cmd != nil {
		if output, err := cmd.Output(); err == nil {
			lines := strings.Split(string(output), "\n")
			fmt.Println("  输出示例:")
			for i, line := range lines {
				if i < 10 && line != "" { // 只显示前10行
					fmt.Printf("    %s\n", line)
				}
			}
			if len(lines) > 10 {
				fmt.Printf("    ... (共%d行)\n", len(lines))
			}
		} else {
			fmt.Printf("  执行go doc命令失败: %v\n", err)
		}
	}

	// 分析当前项目的文档
	fmt.Println("\n分析当前项目的文档:")
	analyzeProjectDocumentation("internal")
}

// analyzeProjectDocumentation 分析项目文档
func analyzeProjectDocumentation(root string) {
	var totalFiles, documentedFiles int
	var packages []string

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(info.Name(), ".go") && !strings.HasSuffix(info.Name(), "_test.go") {
			totalFiles++

			// 检查文件是否有文档
			if hasDocumentation(path) {
				documentedFiles++
			}

			// 收集包名
			if packageName := getPackageName(path); packageName != "" {
				found := false
				for _, pkg := range packages {
					if pkg == packageName {
						found = true
						break
					}
				}
				if !found {
					packages = append(packages, packageName)
				}
			}
		}

		return nil
	})

	if err != nil {
		fmt.Printf("  分析文档时出错: %v\n", err)
		return
	}

	fmt.Printf("  总文件数: %d\n", totalFiles)
	fmt.Printf("  有文档的文件: %d\n", documentedFiles)
	if totalFiles > 0 {
		fmt.Printf("  文档覆盖率: %.1f%%\n", float64(documentedFiles)/float64(totalFiles)*100)
	}
	fmt.Printf("  包数量: %d\n", len(packages))
	fmt.Printf("  包列表: %v\n", packages)
}

// hasDocumentation 检查文件是否有文档
func hasDocumentation(filename string) bool {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		return false
	}

	// 检查包文档
	if file.Doc != nil && len(file.Doc.List) > 0 {
		return true
	}

	// 检查导出的声明是否有文档
	for _, decl := range file.Decls {
		switch d := decl.(type) {
		case *ast.FuncDecl:
			if d.Name.IsExported() && d.Doc != nil && len(d.Doc.List) > 0 {
				return true
			}
		case *ast.GenDecl:
			if d.Doc != nil && len(d.Doc.List) > 0 {
				return true
			}
			for _, spec := range d.Specs {
				switch s := spec.(type) {
				case *ast.TypeSpec:
					if s.Name.IsExported() && s.Doc != nil && len(s.Doc.List) > 0 {
						return true
					}
				case *ast.ValueSpec:
					for _, name := range s.Names {
						if name.IsExported() && s.Doc != nil && len(s.Doc.List) > 0 {
							return true
						}
					}
				}
			}
		}
	}

	return false
}

// demoDocumentationExamples 演示文档示例
func demoDocumentationExamples() {
	fmt.Println("文档示例最佳实践:")

	fmt.Println("\n1. 包级别示例:")
	packageExampleDoc := `// Package strings implements simple functions to manipulate UTF-8 encoded strings.
//
// For information about UTF-8 strings in Go, see https://blog.golang.org/strings.
package strings`
	fmt.Println(packageExampleDoc)

	fmt.Println("\n2. 函数示例:")
	functionExampleDoc := `// Contains reports whether substr is within s.
func Contains(s, substr string) bool`
	fmt.Println(functionExampleDoc)

	fmt.Println("\n3. 复杂函数示例:")
	complexFunctionDoc := `// Replace returns a copy of the string s with the first n
// non-overlapping instances of old replaced by new.
// If old is empty, it matches at the beginning of the string
// and after each UTF-8 sequence, yielding up to k+1 replacements
// for a k-rune string.
// If n < 0, there is no limit on the number of replacements.
func Replace(s, old, new string, n int) string`
	fmt.Println(complexFunctionDoc)

	fmt.Println("\n4. 类型和方法示例:")
	typeMethodDoc := `// Reader implements the io.Reader, io.ReaderAt, io.WriterTo, io.Seeker,
// io.ByteScanner, and io.RuneScanner interfaces by reading from
// a string. The zero value for Reader operates like a Reader of an empty string.
type Reader struct {
    s        string
    i        int64 // current reading index
    prevRune int   // index of previous rune; or < 0
}

// Read implements the io.Reader interface.
func (r *Reader) Read(b []byte) (n int, err error)`
	fmt.Println(typeMethodDoc)
}

// demoDocumentationBestPractices 演示文档最佳实践
func demoDocumentationBestPractices() {
	fmt.Println("文档编写最佳实践:")

	fmt.Println("\n1. 内容原则:")
	contentPrinciples := []string{
		"简洁明了，避免冗余",
		"使用标准的英语语法",
		"第一句话是关键摘要",
		"解释'什么'和'为什么'，而不仅仅是'如何'",
		"包含使用示例",
		"说明边界条件和错误情况",
		"保持文档与代码同步",
	}

	for i, principle := range contentPrinciples {
		fmt.Printf("  %d. %s\n", i+1, principle)
	}

	fmt.Println("\n2. 格式规范:")
	fmt.Println("- 使用标准的Go注释格式")
	fmt.Println("- 代码示例使用缩进")
	fmt.Println("- 使用空行分隔段落")
	fmt.Println("- 链接使用完整URL")

	fmt.Println("\n3. 示例代码:")
	fmt.Println("- 提供可运行的示例")
	fmt.Println("- 使用Example函数")
	fmt.Println("- 包含预期输出")
	fmt.Println("- 展示典型用法")

	fmt.Println("\n4. 文档工具集成:")
	fmt.Println("- 使用go doc查看文档")
	fmt.Println("- 集成到IDE中")
	fmt.Println("- 发布到pkg.go.dev")
	fmt.Println("- 生成静态文档")

	fmt.Println("\n5. 文档维护:")
	fmt.Println("- 代码审查时检查文档")
	fmt.Println("- 定期更新过时文档")
	fmt.Println("- 收集用户反馈")
	fmt.Println("- 使用文档生成工具")

	fmt.Println("\n6. 常见错误:")
	commonMistakes := []string{
		"文档与实现不一致",
		"过于技术化，缺乏使用示例",
		"忽略错误处理说明",
		"文档过于简单或过于复杂",
		"没有说明参数约束",
		"缺少包级别的概述",
	}

	for i, mistake := range commonMistakes {
		fmt.Printf("  %d. %s\n", i+1, mistake)
	}
}
