package stage4

import (
	"context"
	"fmt"
	"time"
)

// DemoContext 演示Context上下文
func DemoContext() {
	fmt.Println("\n=== Context上下文演示 ===")

	// 1. 基本Context使用
	fmt.Println("\n1. 基本Context使用：")
	demoBasicContext()

	// 2. Context取消
	fmt.Println("\n2. Context取消：")
	demoContextCancel()

	// 3. Context超时
	fmt.Println("\n3. Context超时：")
	demoContextTimeout()

	// 4. Context截止时间
	fmt.Println("\n4. Context截止时间：")
	demoContextDeadline()

	// 5. Context值传递
	fmt.Println("\n5. Context值传递：")
	demoContextValue()

	// 6. Context最佳实践
	fmt.Println("\n6. Context最佳实践：")
	demoContextBestPractices()
}

// demoBasicContext 演示基本Context使用
func demoBasicContext() {
	// 1. 背景Context
	ctx := context.Background()
	fmt.Printf("背景Context: %v\n", ctx)

	// 2. TODO Context
	todoCtx := context.TODO()
	fmt.Printf("TODO Context: %v\n", todoCtx)

	// 3. 基本的Context传递
	fmt.Println("\n基本Context传递:")
	processRequest(ctx, "用户请求")
}

// processRequest 处理请求
func processRequest(ctx context.Context, request string) {
	fmt.Printf("处理请求: %s\n", request)

	// 检查context是否被取消
	select {
	case <-ctx.Done():
		fmt.Printf("请求被取消: %v\n", ctx.Err())
		return
	default:
		fmt.Println("请求处理完成")
	}
}

// demoContextCancel 演示Context取消
func demoContextCancel() {
	// 创建可取消的context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // 确保资源清理

	// 启动工作goroutine
	go func() {
		for i := 1; i <= 10; i++ {
			select {
			case <-ctx.Done():
				fmt.Printf("工作被取消: %v\n", ctx.Err())
				return
			default:
				fmt.Printf("执行工作 %d\n", i)
				time.Sleep(100 * time.Millisecond)
			}
		}
		fmt.Println("工作正常完成")
	}()

	// 等待一段时间后取消
	time.Sleep(350 * time.Millisecond)
	fmt.Println("发送取消信号")
	cancel()

	// 等待goroutine结束
	time.Sleep(100 * time.Millisecond)
}

// demoContextTimeout 演示Context超时
func demoContextTimeout() {
	// 创建带超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	// 启动可能耗时的操作
	result := make(chan string, 1)

	go func() {
		// 模拟耗时操作
		time.Sleep(300 * time.Millisecond)
		result <- "操作完成"
	}()

	// 等待结果或超时
	select {
	case res := <-result:
		fmt.Printf("收到结果: %s\n", res)
	case <-ctx.Done():
		fmt.Printf("操作超时: %v\n", ctx.Err())
	}

	// 演示不同的超时场景
	fmt.Println("\n不同超时场景:")
	testTimeoutScenarios()
}

// testTimeoutScenarios 测试不同超时场景
func testTimeoutScenarios() {
	scenarios := []struct {
		name     string
		timeout  time.Duration
		workTime time.Duration
	}{
		{"快速完成", 200 * time.Millisecond, 100 * time.Millisecond},
		{"刚好超时", 150 * time.Millisecond, 150 * time.Millisecond},
		{"明显超时", 100 * time.Millisecond, 200 * time.Millisecond},
	}

	for _, scenario := range scenarios {
		fmt.Printf("\n场景: %s\n", scenario.name)

		ctx, cancel := context.WithTimeout(context.Background(), scenario.timeout)

		done := make(chan bool, 1)
		go func() {
			time.Sleep(scenario.workTime)
			done <- true
		}()

		select {
		case <-done:
			fmt.Println("任务完成")
		case <-ctx.Done():
			fmt.Printf("任务超时: %v\n", ctx.Err())
		}

		cancel()
	}
}

// demoContextDeadline 演示Context截止时间
func demoContextDeadline() {
	// 设置截止时间为当前时间后300ms
	deadline := time.Now().Add(300 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	fmt.Printf("设置截止时间: %v\n", deadline.Format("15:04:05.000"))

	// 检查截止时间
	if dl, ok := ctx.Deadline(); ok {
		fmt.Printf("Context截止时间: %v\n", dl.Format("15:04:05.000"))
		fmt.Printf("剩余时间: %v\n", time.Until(dl))
	}

	// 执行任务直到截止时间
	ticker := time.NewTicker(50 * time.Millisecond)
	defer ticker.Stop()

	for i := 1; ; i++ {
		select {
		case <-ctx.Done():
			fmt.Printf("达到截止时间，停止执行: %v\n", ctx.Err())
			return
		case <-ticker.C:
			fmt.Printf("执行任务 %d，当前时间: %v\n", i, time.Now().Format("15:04:05.000"))
		}
	}
}

// demoContextValue 演示Context值传递
func demoContextValue() {
	// 定义context key类型
	type contextKey string

	const (
		userIDKey    contextKey = "userID"
		requestIDKey contextKey = "requestID"
		traceIDKey   contextKey = "traceID"
	)

	// 创建带值的context
	ctx := context.Background()
	ctx = context.WithValue(ctx, userIDKey, "user123")
	ctx = context.WithValue(ctx, requestIDKey, "req456")
	ctx = context.WithValue(ctx, traceIDKey, "trace789")

	// 传递context到不同的函数
	fmt.Println("Context值传递:")
	handleUserRequest(ctx)
}

// handleUserRequest 处理用户请求
func handleUserRequest(ctx context.Context) {
	// 从context中获取值
	userID := ctx.Value("userID")
	requestID := ctx.Value("requestID")
	traceID := ctx.Value("traceID")

	fmt.Printf("处理用户请求 - UserID: %v, RequestID: %v, TraceID: %v\n",
		userID, requestID, traceID)

	// 调用其他服务
	callExternalService(ctx)

	// 记录日志
	logRequest(ctx)
}

// callExternalService 调用外部服务
func callExternalService(ctx context.Context) {
	traceID := ctx.Value("traceID")
	fmt.Printf("调用外部服务 - TraceID: %v\n", traceID)

	// 模拟服务调用
	time.Sleep(50 * time.Millisecond)
	fmt.Println("外部服务调用完成")
}

// logRequest 记录请求日志
func logRequest(ctx context.Context) {
	userID := ctx.Value("userID")
	requestID := ctx.Value("requestID")

	fmt.Printf("记录日志 - UserID: %v, RequestID: %v\n", userID, requestID)
}

// demoContextBestPractices 演示Context最佳实践
func demoContextBestPractices() {
	fmt.Println("Context最佳实践演示:")

	// 1. 链式Context
	fmt.Println("\n1. 链式Context:")
	demoContextChaining()

	// 2. Context传播
	fmt.Println("\n2. Context传播:")
	demoContextPropagation()

	// 3. 错误处理
	fmt.Println("\n3. 错误处理:")
	demoContextErrorHandling()
}

// demoContextChaining 演示Context链式使用
func demoContextChaining() {
	// 创建基础context
	baseCtx := context.Background()

	// 添加超时
	timeoutCtx, cancel1 := context.WithTimeout(baseCtx, 500*time.Millisecond)
	defer cancel1()

	// 添加取消功能
	cancelCtx, cancel2 := context.WithCancel(timeoutCtx)
	defer cancel2()

	// 添加值
	valueCtx := context.WithValue(cancelCtx, "operation", "chain-demo")

	// 使用链式context
	go func() {
		for i := 1; i <= 10; i++ {
			select {
			case <-valueCtx.Done():
				fmt.Printf("链式操作被中断: %v\n", valueCtx.Err())
				return
			default:
				operation := valueCtx.Value("operation")
				fmt.Printf("执行 %v 步骤 %d\n", operation, i)
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	// 提前取消
	time.Sleep(250 * time.Millisecond)
	cancel2()
	time.Sleep(100 * time.Millisecond)
}

// demoContextPropagation 演示Context传播
func demoContextPropagation() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	// 启动多层调用
	if err := serviceA(ctx); err != nil {
		fmt.Printf("服务调用失败: %v\n", err)
	}
}

// serviceA 服务A
func serviceA(ctx context.Context) error {
	fmt.Println("服务A: 开始处理")

	// 检查context状态
	select {
	case <-ctx.Done():
		return fmt.Errorf("服务A被取消: %w", ctx.Err())
	default:
	}

	// 调用服务B
	if err := serviceB(ctx); err != nil {
		return fmt.Errorf("服务A调用服务B失败: %w", err)
	}

	fmt.Println("服务A: 处理完成")
	return nil
}

// serviceB 服务B
func serviceB(ctx context.Context) error {
	fmt.Println("服务B: 开始处理")

	// 模拟处理时间
	timer := time.NewTimer(200 * time.Millisecond)
	defer timer.Stop()

	select {
	case <-ctx.Done():
		return fmt.Errorf("服务B被取消: %w", ctx.Err())
	case <-timer.C:
		fmt.Println("服务B: 处理完成")
		return nil
	}
}

// demoContextErrorHandling 演示Context错误处理
func demoContextErrorHandling() {
	// 测试不同的错误类型
	testCases := []struct {
		name string
		ctx  context.Context
	}{
		{
			name: "取消错误",
			ctx: func() context.Context {
				ctx, cancel := context.WithCancel(context.Background())
				cancel()
				return ctx
			}(),
		},
		{
			name: "超时错误",
			ctx: func() context.Context {
				ctx, cancel := context.WithTimeout(context.Background(), -1*time.Second)
				defer cancel()
				return ctx
			}(),
		},
		{
			name: "截止时间错误",
			ctx: func() context.Context {
				ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(-1*time.Second))
				defer cancel()
				return ctx
			}(),
		},
	}

	for _, tc := range testCases {
		fmt.Printf("\n测试 %s:\n", tc.name)

		select {
		case <-tc.ctx.Done():
			err := tc.ctx.Err()
			switch err {
			case context.Canceled:
				fmt.Println("Context被取消")
			case context.DeadlineExceeded:
				fmt.Println("Context超时")
			default:
				fmt.Printf("其他错误: %v\n", err)
			}
		default:
			fmt.Println("Context仍然有效")
		}
	}
}
