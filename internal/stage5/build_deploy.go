package stage5

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

// DemoBuildDeploy 演示构建部署
func DemoBuildDeploy() {
	fmt.Println("\n=== 构建部署演示 ===")

	// 1. 构建基础
	fmt.Println("\n1. 构建基础：")
	demoBuildBasics()

	// 2. 交叉编译
	fmt.Println("\n2. 交叉编译：")
	demoCrossCompilation()

	// 3. 构建优化
	fmt.Println("\n3. 构建优化：")
	demoBuildOptimization()

	// 4. 部署策略
	fmt.Println("\n4. 部署策略：")
	demoDeploymentStrategies()

	// 5. 容器化部署
	fmt.Println("\n5. 容器化部署：")
	demoContainerization()

	// 6. CI/CD集成
	fmt.Println("\n6. CI/CD集成：")
	demoCICD()
}

// demoBuildBasics 演示构建基础
func demoBuildBasics() {
	fmt.Println("Go构建系统基础:")
	fmt.Println("- go build: 编译包和依赖")
	fmt.Println("- go install: 编译并安装包")
	fmt.Println("- go run: 编译并运行程序")
	fmt.Println("- go clean: 清理构建文件")

	fmt.Println("\n基本构建命令:")
	buildCommands := []struct {
		cmd  string
		desc string
	}{
		{"go build", "构建当前目录的包"},
		{"go build .", "构建当前目录的包"},
		{"go build ./...", "构建当前目录及子目录的所有包"},
		{"go build -o myapp", "指定输出文件名"},
		{"go build -v", "显示详细构建信息"},
		{"go build -x", "显示执行的命令"},
		{"go build -race", "启用竞态检测"},
		{"go build -tags=prod", "使用构建标签"},
	}

	for _, cmd := range buildCommands {
		fmt.Printf("  %-25s - %s\n", cmd.cmd, cmd.desc)
	}

	fmt.Println("\n构建标签示例:")
	buildTagsExample := `// +build prod

package config

const (
    Debug = false
    LogLevel = "error"
)

// +build !prod

package config

const (
    Debug = true
    LogLevel = "debug"
)`
	fmt.Println(buildTagsExample)

	// 演示构建当前项目
	fmt.Println("\n构建当前项目:")
	demoBuildCurrentProject()
}

// demoBuildCurrentProject 演示构建当前项目
func demoBuildCurrentProject() {
	fmt.Println("  尝试构建当前项目...")

	// 检查是否有main包
	hasMain := false
	if files, err := filepath.Glob("*.go"); err == nil {
		for _, file := range files {
			if content, err := os.ReadFile(file); err == nil {
				if strings.Contains(string(content), "package main") {
					hasMain = true
					break
				}
			}
		}
	}

	if !hasMain {
		// 检查cmd目录
		if _, err := os.Stat("cmd"); err == nil {
			fmt.Println("  发现cmd目录，查找main包...")
			err := filepath.Walk("cmd", func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if !info.IsDir() && strings.HasSuffix(info.Name(), ".go") {
					if content, err := os.ReadFile(path); err == nil {
						if strings.Contains(string(content), "package main") {
							hasMain = true
							fmt.Printf("  找到main包: %s\n", path)
							return filepath.SkipDir
						}
					}
				}
				return nil
			})
			if err != nil {
				fmt.Printf("  搜索main包时出错: %v\n", err)
			}
		}
	}

	if hasMain {
		// 尝试构建
		if cmd := exec.Command("go", "build", "-v"); cmd != nil {
			if output, err := cmd.CombinedOutput(); err != nil {
				fmt.Printf("  构建失败: %v\n", err)
				if len(output) > 0 {
					fmt.Printf("  输出: %s\n", string(output))
				}
			} else {
				fmt.Println("  构建成功!")
				if len(output) > 0 {
					fmt.Printf("  构建信息: %s\n", string(output))
				}
			}
		}
	} else {
		fmt.Println("  当前项目没有main包，无法构建可执行文件")
		fmt.Println("  这是一个库项目，可以使用 go build ./... 检查编译")

		// 尝试编译检查
		if cmd := exec.Command("go", "build", "./..."); cmd != nil {
			if output, err := cmd.CombinedOutput(); err != nil {
				fmt.Printf("  编译检查失败: %v\n", err)
				if len(output) > 0 {
					fmt.Printf("  输出: %s\n", string(output))
				}
			} else {
				fmt.Println("  编译检查通过!")
			}
		}
	}
}

// demoCrossCompilation 演示交叉编译
func demoCrossCompilation() {
	fmt.Println("Go交叉编译:")
	fmt.Println("- 支持多种操作系统和架构")
	fmt.Println("- 使用GOOS和GOARCH环境变量")
	fmt.Println("- 无需安装目标平台的工具链")

	fmt.Println("\n支持的平台:")
	platforms := []struct {
		os   string
		arch string
		desc string
	}{
		{"linux", "amd64", "Linux 64位"},
		{"linux", "386", "Linux 32位"},
		{"linux", "arm64", "Linux ARM64"},
		{"windows", "amd64", "Windows 64位"},
		{"windows", "386", "Windows 32位"},
		{"darwin", "amd64", "macOS Intel"},
		{"darwin", "arm64", "macOS Apple Silicon"},
		{"freebsd", "amd64", "FreeBSD 64位"},
	}

	for _, platform := range platforms {
		fmt.Printf("  %-8s/%-6s - %s\n", platform.os, platform.arch, platform.desc)
	}

	fmt.Println("\n交叉编译命令示例:")
	crossCompileExamples := []string{
		"GOOS=linux GOARCH=amd64 go build -o myapp-linux-amd64",
		"GOOS=windows GOARCH=amd64 go build -o myapp-windows-amd64.exe",
		"GOOS=darwin GOARCH=arm64 go build -o myapp-darwin-arm64",
	}

	for _, example := range crossCompileExamples {
		fmt.Printf("  %s\n", example)
	}

	fmt.Println("\n当前平台信息:")
	fmt.Printf("  GOOS: %s\n", runtime.GOOS)
	fmt.Printf("  GOARCH: %s\n", runtime.GOARCH)
	fmt.Printf("  NumCPU: %d\n", runtime.NumCPU())

	// 显示所有支持的平台
	fmt.Println("\n查看所有支持的平台:")
	if cmd := exec.Command("go", "tool", "dist", "list"); cmd != nil {
		if output, err := cmd.Output(); err == nil {
			lines := strings.Split(strings.TrimSpace(string(output)), "\n")
			fmt.Printf("  共支持 %d 个平台组合\n", len(lines))
			fmt.Println("  前10个平台:")
			for i, line := range lines {
				if i < 10 {
					fmt.Printf("    %s\n", line)
				}
			}
			if len(lines) > 10 {
				fmt.Println("    ...")
			}
		}
	}
}

// demoBuildOptimization 演示构建优化
func demoBuildOptimization() {
	fmt.Println("构建优化技术:")

	fmt.Println("\n1. 编译器优化:")
	compilerOptimizations := []struct {
		flag string
		desc string
	}{
		{"-ldflags='-s -w'", "去除符号表和调试信息"},
		{"-trimpath", "移除文件系统路径"},
		{"-buildmode=pie", "生成位置无关可执行文件"},
		{"-race", "启用竞态检测（调试用）"},
		{"-msan", "启用内存清理检测"},
	}

	for _, opt := range compilerOptimizations {
		fmt.Printf("  %-25s - %s\n", opt.flag, opt.desc)
	}

	fmt.Println("\n2. 链接器优化:")
	linkerOptimizations := []struct {
		flag string
		desc string
	}{
		{"-X main.version=1.0.0", "设置字符串变量值"},
		{"-X main.buildTime=$(date)", "设置构建时间"},
		{"-extldflags '-static'", "静态链接"},
		{"-linkmode external", "使用外部链接器"},
	}

	for _, opt := range linkerOptimizations {
		fmt.Printf("  %-30s - %s\n", opt.flag, opt.desc)
	}

	fmt.Println("\n3. 构建模式:")
	buildModes := []struct {
		mode string
		desc string
	}{
		{"exe", "可执行文件（默认）"},
		{"pie", "位置无关可执行文件"},
		{"c-archive", "C静态库"},
		{"c-shared", "C动态库"},
		{"shared", "Go共享库"},
		{"plugin", "Go插件"},
	}

	for _, mode := range buildModes {
		fmt.Printf("  %-12s - %s\n", mode.mode, mode.desc)
	}

	fmt.Println("\n4. 优化示例:")
	optimizationExample := `# 生产环境构建
go build -ldflags="-s -w -X main.version=1.0.0 -X main.buildTime=$(date)" \
         -trimpath \
         -o myapp

# 调试构建
go build -race -o myapp-debug

# 静态链接构建
CGO_ENABLED=0 go build -ldflags="-s -w" -o myapp-static`
	fmt.Println(optimizationExample)
}

// demoDeploymentStrategies 演示部署策略
func demoDeploymentStrategies() {
	fmt.Println("部署策略:")

	fmt.Println("\n1. 传统部署:")
	fmt.Println("- 直接部署二进制文件")
	fmt.Println("- 使用systemd等服务管理")
	fmt.Println("- 配置文件和日志管理")
	fmt.Println("- 进程监控和重启")

	fmt.Println("\n2. 容器化部署:")
	fmt.Println("- Docker容器")
	fmt.Println("- Kubernetes编排")
	fmt.Println("- 镜像版本管理")
	fmt.Println("- 滚动更新")

	fmt.Println("\n3. 云原生部署:")
	fmt.Println("- 无服务器函数")
	fmt.Println("- 容器即服务")
	fmt.Println("- 托管Kubernetes")
	fmt.Println("- 自动扩缩容")

	fmt.Println("\n4. 部署最佳实践:")
	deploymentPractices := []string{
		"使用版本标签",
		"健康检查端点",
		"优雅关闭处理",
		"配置外部化",
		"日志结构化",
		"监控和告警",
		"备份和恢复",
		"安全加固",
	}

	for i, practice := range deploymentPractices {
		fmt.Printf("  %d. %s\n", i+1, practice)
	}
}

// demoContainerization 演示容器化部署
func demoContainerization() {
	fmt.Println("容器化部署:")

	fmt.Println("\n1. Dockerfile示例:")
	dockerfileExample := `# 多阶段构建
FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o main .

# 运行阶段
FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /app/main .

EXPOSE 8080
CMD ["./main"]`
	fmt.Println(dockerfileExample)

	fmt.Println("\n2. Docker命令:")
	dockerCommands := []struct {
		cmd  string
		desc string
	}{
		{"docker build -t myapp .", "构建镜像"},
		{"docker run -p 8080:8080 myapp", "运行容器"},
		{"docker push myapp:latest", "推送镜像"},
		{"docker-compose up", "使用compose启动"},
	}

	for _, cmd := range dockerCommands {
		fmt.Printf("  %-30s - %s\n", cmd.cmd, cmd.desc)
	}

	fmt.Println("\n3. Kubernetes部署:")
	k8sExample := `apiVersion: apps/v1
kind: Deployment
metadata:
  name: myapp
spec:
  replicas: 3
  selector:
    matchLabels:
      app: myapp
  template:
    metadata:
      labels:
        app: myapp
    spec:
      containers:
      - name: myapp
        image: myapp:latest
        ports:
        - containerPort: 8080
        env:
        - name: ENV
          value: "production"
        resources:
          requests:
            memory: "64Mi"
            cpu: "250m"
          limits:
            memory: "128Mi"
            cpu: "500m"`
	fmt.Println(k8sExample)
}

// demoCICD 演示CI/CD集成
func demoCICD() {
	fmt.Println("CI/CD集成:")

	fmt.Println("\n1. GitHub Actions示例:")
	githubActionsExample := `name: Build and Deploy

on:
  push:
    branches: [ main ]
  pull_request:
    branches: [ main ]

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v3
    
    - name: Set up Go
      uses: actions/setup-go@v3
      with:
        go-version: 1.21
    
    - name: Test
      run: go test -v ./...
    
    - name: Build
      run: go build -v ./...

  build:
    needs: test
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v3
    
    - name: Build Docker image
      run: docker build -t myapp .
    
    - name: Push to registry
      run: docker push myapp:latest`
	fmt.Println(githubActionsExample)

	fmt.Println("\n2. GitLab CI示例:")
	gitlabCIExample := `stages:
  - test
  - build
  - deploy

test:
  stage: test
  image: golang:1.21
  script:
    - go test -v ./...

build:
  stage: build
  image: docker:latest
  script:
    - docker build -t myapp .
    - docker push myapp:latest

deploy:
  stage: deploy
  script:
    - kubectl apply -f k8s/`
	fmt.Println(gitlabCIExample)

	fmt.Println("\n3. CI/CD最佳实践:")
	cicdPractices := []string{
		"自动化测试",
		"代码质量检查",
		"安全扫描",
		"依赖检查",
		"多环境部署",
		"回滚机制",
		"监控集成",
		"通知机制",
	}

	for i, practice := range cicdPractices {
		fmt.Printf("  %d. %s\n", i+1, practice)
	}

	fmt.Println("\n4. 部署工具:")
	deploymentTools := []struct {
		tool string
		desc string
	}{
		{"GitHub Actions", "GitHub集成CI/CD"},
		{"GitLab CI", "GitLab集成CI/CD"},
		{"Jenkins", "开源CI/CD平台"},
		{"Docker", "容器化平台"},
		{"Kubernetes", "容器编排"},
		{"Helm", "Kubernetes包管理"},
		{"Terraform", "基础设施即代码"},
		{"Ansible", "配置管理"},
	}

	for _, tool := range deploymentTools {
		fmt.Printf("  %-15s - %s\n", tool.tool, tool.desc)
	}
}
