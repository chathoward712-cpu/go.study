# go.study - Go 学习项目

一个从零开始学习 Go 语言的项目，包含基础概念、常见模式和最佳实践的实例代码。

## 项目结构

```
go.study/
├── internal/              # 内部包（不能被其他项目导入）
│   ├── stage1/           # 第1阶段：基础语法
│   ├── stage2/           # 第2阶段：数据结构
│   ├── stage3/           # 第3阶段：面向对象
│   ├── stage4/           # 第4阶段：并发编程
│   └── stage5/           # 第5阶段：模块化与工程实践
├── main.go               # 主程序入口
├── go.mod                # Go 模块定义
└── README.md             # 本文件
```

## 快速开始

### 运行学习演示程序

```bash
# 运行当前阶段的演示（第5阶段：模块化与工程实践）
go run main.go

# 或者先构建再运行
go build -o go-study .
./go-study
```

### 运行所有测试

```bash
go test ./...
```

### 查看测试覆盖率

```bash
go test -cover ./...
```

### 代码格式化

```bash
go fmt ./...
```

### 依赖管理

```bash
# 整理依赖
go mod tidy

# 下载依赖
go mod download

# 验证依赖
go mod verify
```

## 学习阶段

### 第1阶段：基础语法 (internal/stage1/)
- 基本数据类型：整数、浮点数、布尔值、字符串
- 变量声明和常量定义
- 运算符和表达式
- 控制流：if、for、switch
- 函数定义和调用

### 第2阶段：数据结构 (internal/stage2/)
- 数组和切片 (Arrays & Slices)
- 映射 (Maps)
- 结构体 (Structs)
- 指针 (Pointers)
- 方法 (Methods)

### 第3阶段：面向对象 (internal/stage3/)
- 接口 (Interfaces)
- 类型断言和类型选择
- 嵌入和组合
- 多态性
- 错误处理

### 第4阶段：并发编程 (internal/stage4/)
- Goroutines：轻量级线程
- Channels：通道通信
- Select：多路复用
- 同步原语：Mutex、WaitGroup
- 并发模式和最佳实践

### 第5阶段：模块化与工程实践 (internal/stage5/)
- Go 模块系统
- 包管理和依赖管理
- 单元测试和基准测试
- 文档生成
- 构建和部署
- CI/CD 集成

## 使用方法

### 1. 克隆项目

```bash
git clone <repository-url>
cd go.study
```

### 2. 安装 Go

确保已安装 Go 1.19 或更高版本：

```bash
go version
```

### 3. 运行演示

```bash
# 运行完整的学习演示
go run main.go

# 运行特定阶段（需要修改 main.go 中的导入）
# 例如运行第1阶段
# import "github.com/howard/go.study/internal/stage1"
# stage1.RunStage1()
```

### 4. 构建可执行文件

```bash
# 构建到当前目录
go build .

# 构建到指定位置
go build -o bin/go-study .

# 交叉编译
GOOS=linux GOARCH=amd64 go build -o go-study-linux .
GOOS=windows GOARCH=amd64 go build -o go-study.exe .
```

### 5. 运行测试

```bash
# 运行所有测试
go test ./...

# 运行特定包的测试
go test -v ./internal/stage1
go test -v ./pkg/utils

# 运行基准测试
go test -bench=. ./...

# 查看测试覆盖率
go test -cover ./...
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

### 6. 代码质量检查

```bash
# 格式化代码
go fmt ./...

# 静态分析
go vet ./...

# 运行 golint（需要安装）
golint ./...
```

### 7. 文档生成

```bash
# 查看包文档
go doc ./internal/stage1

# 启动本地文档服务器
godoc -http=:6060
# 访问 http://localhost:6060
```

## 开发环境设置

### 1. Go 环境配置

```bash
# 设置 Go 代理（中国用户推荐）
go env -w GOPROXY=https://goproxy.cn,direct
go env -w GOSUMDB=sum.golang.org

# 查看 Go 环境配置
go env
```

### 2. 编辑器配置

推荐使用支持 Go 的编辑器：
- **VS Code** + Go 扩展
- **GoLand** (JetBrains)
- **Vim/Neovim** + vim-go
- **Emacs** + go-mode

### 3. 开发工具

```bash
# 安装常用工具
go install golang.org/x/tools/cmd/goimports@latest
go install golang.org/x/lint/golint@latest
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

## 开发规范

遵循 [Effective Go](https://golang.org/doc/effective_go) 和 [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)。

### 命名规范
- **包名**：小写单数形式 (`stage1`, `utils`)
- **导出标识符**：大写开头 (`RunStage1`, `DemoBasics`)
- **非导出标识符**：小写开头 (`demoVariables`, `makeCounter`)
- **接口名**：以 `-er` 结尾 (`Reader`, `Writer`, `Shape`)
- **常量**：驼峰命名或全大写 (`MaxSize`, `DEFAULT_TIMEOUT`)

### 代码组织
- 使用 `internal/` 目录存放内部包
- 主程序入口为 `main.go`
- 一个目录一个包，包名与目录名一致
- 按学习阶段组织代码结构

### 测试规范
- 测试文件以 `_test.go` 结尾
- 测试函数以 `Test` 开头
- 基准测试函数以 `Benchmark` 开头
- 示例函数以 `Example` 开头
- 使用表驱动测试处理多个测试用例

## 学习建议

### 1. 循序渐进
按照阶段顺序学习，每个阶段都有完整的示例和说明：
1. 基础语法 → 2. 数据结构 → 3. 面向对象 → 4. 并发编程 → 5. 工程实践

### 2. 动手实践
- 运行每个阶段的演示代码
- 修改代码观察结果变化
- 尝试编写自己的示例

### 3. 阅读源码
- 查看 `internal/` 目录下的实现
- 理解每个函数的设计思路
- 学习代码组织和命名规范

### 4. 编写测试
- 为自己的代码编写测试
- 学习使用 `testing` 包
- 实践 TDD（测试驱动开发）

## 常见问题

### Q: 如何切换到不同的学习阶段？
A: 修改 `main.go` 文件中的导入和调用：
```go
// 切换到第1阶段
import "github.com/howard/go.study/internal/stage1"
func main() {
    stage1.RunStage1()
}
```

### Q: 如何添加新的演示内容？
A: 在对应的 stage 目录下添加新的 `.go` 文件，并在主函数中调用。

### Q: 测试失败怎么办？
A: 检查 Go 版本（需要 1.19+），运行 `go mod tidy` 更新依赖。

### Q: 如何贡献代码？
A: 
1. Fork 项目
2. 创建特性分支
3. 提交更改
4. 发起 Pull Request

## 参考资源

### 官方资源
- [Go 官方网站](https://golang.org)
- [Go 语言规范](https://golang.org/ref/spec)
- [Effective Go](https://golang.org/doc/effective_go)
- [Go 标准库](https://pkg.go.dev/std)

### 学习资源
- [Go by Example](https://gobyexample.com)
- [A Tour of Go](https://tour.golang.org)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- [Go 语言圣经](https://gopl.io)

### 工具和库
- [Awesome Go](https://awesome-go.com) - Go 生态系统
- [Go Report Card](https://goreportcard.com) - 代码质量检查
- [pkg.go.dev](https://pkg.go.dev) - Go 包文档

## 许可证

本项目采用 MIT 许可证，详见 LICENSE 文件。

## 贡献

欢迎提交 Issue 和 Pull Request 来改进这个学习项目！
