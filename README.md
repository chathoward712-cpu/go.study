# go.study - Go 学习项目

一个从零开始学习 Go 语言的项目，包含基础概念、常见模式和最佳实践的实例代码。

## 项目结构

```
go.study/
├── cmd/                    # 可执行程序入口
│   └── hello/             # Hello 示例程序
├── internal/              # 内部包（不能被其他项目导入）
│   └── basics/           # Go 基础概念演示
├── pkg/                   # 公共包（可被其他项目导入）
│   └── utils/            # 工具函数
├── go.mod                 # Go 模块定义
└── README.md              # 本文件
```

## 快速开始

### 运行示例程序

```bash
go run ./cmd/hello
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

## 学习内容

### 1. 基础类型 (internal/basics/basics.go)
- 整数、浮点数、布尔值、字符、字符串
- 类型声明和类型转换

### 2. 复合数据类型
- 切片 (Slices)：动态数组
- 映射 (Maps)：键值对集合

### 3. 函数和闭包
- 函数定义和调用
- 高阶函数（函数作为参数）
- 闭包：函数返回函数

### 4. 接口 (Interfaces)
- 接口定义
- 实现接口
- 多态

### 5. 错误处理
- error 类型
- 自定义错误

## 测试

项目使用 Go 标准库的 `testing` 包编写测试。

### 运行特定包的测试

```bash
go test -v ./internal/basics
go test -v ./pkg/utils
```

### 运行基准测试

```bash
go test -bench=. ./internal/basics
```

## 开发规范

遵循 [Effective Go](https://golang.org/doc/effective_go) 和 [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)。

关键规范：
- 包名使用小写单数形式：`basics`, `utils`
- 导出的名称以大写开头：`Max`, `Reverse`
- 非导出的名称以小写开头：`makeCounter`
- 接口名以 `-er` 结尾：`Reader`, `Writer`, `Shape`
- 使用表驱动测试处理多个测试用例

## 资源

- [官方文档](https://golang.org)
- [Effective Go](https://golang.org/doc/effective_go)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- [标准库](https://pkg.go.dev/std)

## 提示

- 所有 `*_test.go` 文件与被测试的代码在同一个包中
- 测试函数以 `Test` 开头，参数为 `*testing.T`
- 使用 `-v` 标志查看详细的测试输出
- 使用 `-run` 标志运行特定的测试：`go test -v -run TestMax ./pkg/utils`
