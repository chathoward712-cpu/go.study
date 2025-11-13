package stage4

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// DemoGoroutines 演示Goroutine基础
func DemoGoroutines() {
	fmt.Println("\n=== Goroutine基础演示 ===")

	// 1. 基本Goroutine使用
	fmt.Println("\n1. 基本Goroutine使用：")
	demoBasicGoroutine()

	// 2. 多个Goroutine
	fmt.Println("\n2. 多个Goroutine：")
	demoMultipleGoroutines()

	// 3. Goroutine与匿名函数
	fmt.Println("\n3. Goroutine与匿名函数：")
	demoGoroutineWithAnonymousFunc()

	// 4. Goroutine的生命周期
	fmt.Println("\n4. Goroutine的生命周期：")
	demoGoroutineLifecycle()

	// 5. WaitGroup同步
	fmt.Println("\n5. WaitGroup同步：")
	demoWaitGroup()

	// 6. Goroutine泄漏预防
	fmt.Println("\n6. Goroutine泄漏预防：")
	demoGoroutineLeakPrevention()

	// 7. 运行时信息
	fmt.Println("\n7. 运行时信息：")
	demoRuntimeInfo()
}

// demoBasicGoroutine 演示基本Goroutine使用
func demoBasicGoroutine() {
	// 1. 普通函数调用
	fmt.Println("普通函数调用:")
	sayHello("World")

	// 2. Goroutine调用
	fmt.Println("\nGoroutine调用:")
	go sayHello("Goroutine")

	// 等待一下，让goroutine有时间执行
	time.Sleep(100 * time.Millisecond)

	// 3. 对比执行顺序
	fmt.Println("\n执行顺序对比:")
	fmt.Println("主线程: 开始")

	go func() {
		fmt.Println("Goroutine: 异步执行")
	}()

	fmt.Println("主线程: 继续执行")
	time.Sleep(50 * time.Millisecond)
	fmt.Println("主线程: 结束")
}

// sayHello 简单的问候函数
func sayHello(name string) {
	for i := 0; i < 3; i++ {
		fmt.Printf("Hello, %s! (%d)\n", name, i+1)
		time.Sleep(10 * time.Millisecond)
	}
}

// demoMultipleGoroutines 演示多个Goroutine
func demoMultipleGoroutines() {
	fmt.Println("启动多个Goroutine:")

	// 启动多个goroutine
	for i := 1; i <= 5; i++ {
		go func(id int) {
			fmt.Printf("Goroutine %d: 开始执行\n", id)
			time.Sleep(time.Duration(id*10) * time.Millisecond)
			fmt.Printf("Goroutine %d: 执行完成\n", id)
		}(i) // 注意：传递参数避免闭包陷阱
	}

	// 等待所有goroutine完成
	time.Sleep(100 * time.Millisecond)

	// 演示闭包陷阱
	fmt.Println("\n闭包陷阱示例:")
	fmt.Println("错误的方式:")
	for i := 1; i <= 3; i++ {
		go func() {
			fmt.Printf("错误: Goroutine %d\n", i) // 可能都打印4
		}()
	}
	time.Sleep(50 * time.Millisecond)

	fmt.Println("正确的方式:")
	for i := 1; i <= 3; i++ {
		go func(id int) {
			fmt.Printf("正确: Goroutine %d\n", id)
		}(i)
	}
	time.Sleep(50 * time.Millisecond)
}

// demoGoroutineWithAnonymousFunc 演示Goroutine与匿名函数
func demoGoroutineWithAnonymousFunc() {
	// 1. 简单匿名函数
	go func() {
		fmt.Println("匿名函数Goroutine执行")
	}()

	// 2. 带参数的匿名函数
	message := "Hello from anonymous goroutine"
	go func(msg string) {
		fmt.Printf("带参数的匿名函数: %s\n", msg)
	}(message)

	// 3. 带返回值的匿名函数（通过channel返回）
	resultChan := make(chan int)
	go func() {
		result := 42 * 2
		resultChan <- result
	}()

	// 4. 复杂的匿名函数
	go func() {
		for i := 1; i <= 3; i++ {
			fmt.Printf("复杂匿名函数: 步骤 %d\n", i)
			time.Sleep(20 * time.Millisecond)
		}
	}()

	// 等待结果
	result := <-resultChan
	fmt.Printf("匿名函数计算结果: %d\n", result)

	time.Sleep(100 * time.Millisecond)
}

// demoGoroutineLifecycle 演示Goroutine的生命周期
func demoGoroutineLifecycle() {
	fmt.Printf("主程序开始，当前Goroutine数量: %d\n", runtime.NumGoroutine())

	// 创建一个有生命周期的goroutine
	done := make(chan bool)

	go func() {
		fmt.Println("长期运行的Goroutine开始")
		for i := 1; i <= 5; i++ {
			fmt.Printf("长期Goroutine: 工作 %d\n", i)
			time.Sleep(50 * time.Millisecond)
		}
		fmt.Println("长期运行的Goroutine结束")
		done <- true
	}()

	fmt.Printf("创建Goroutine后，当前数量: %d\n", runtime.NumGoroutine())

	// 创建一些短期goroutine
	for i := 1; i <= 3; i++ {
		go func(id int) {
			fmt.Printf("短期Goroutine %d执行\n", id)
		}(i)
	}

	fmt.Printf("创建更多Goroutine后，当前数量: %d\n", runtime.NumGoroutine())

	// 等待长期goroutine完成
	<-done
	time.Sleep(50 * time.Millisecond) // 等待短期goroutine完成

	fmt.Printf("所有Goroutine完成后，当前数量: %d\n", runtime.NumGoroutine())
}

// demoWaitGroup 演示WaitGroup同步
func demoWaitGroup() {
	var wg sync.WaitGroup

	fmt.Println("使用WaitGroup同步多个Goroutine:")

	// 启动多个工作goroutine
	for i := 1; i <= 5; i++ {
		wg.Add(1) // 增加等待计数

		go func(id int) {
			defer wg.Done() // 完成时减少计数

			fmt.Printf("工作者 %d: 开始工作\n", id)

			// 模拟不同的工作时间
			workTime := time.Duration(id*20) * time.Millisecond
			time.Sleep(workTime)

			fmt.Printf("工作者 %d: 工作完成 (耗时 %v)\n", id, workTime)
		}(i)
	}

	fmt.Println("等待所有工作者完成...")
	wg.Wait() // 等待所有goroutine完成
	fmt.Println("所有工作者都完成了!")

	// 演示WaitGroup的错误用法预防
	fmt.Println("\nWaitGroup最佳实践:")
	demoWaitGroupBestPractices()
}

// demoWaitGroupBestPractices 演示WaitGroup最佳实践
func demoWaitGroupBestPractices() {
	var wg sync.WaitGroup

	// 最佳实践：在启动goroutine之前调用Add
	tasks := []string{"任务A", "任务B", "任务C"}

	for _, task := range tasks {
		wg.Add(1) // 在goroutine启动前Add

		go func(taskName string) {
			defer wg.Done() // 使用defer确保Done被调用

			fmt.Printf("执行 %s\n", taskName)
			time.Sleep(30 * time.Millisecond)
			fmt.Printf("%s 完成\n", taskName)
		}(task) // 传递参数避免闭包问题
	}

	wg.Wait()
	fmt.Println("所有任务完成")
}

// demoGoroutineLeakPrevention 演示Goroutine泄漏预防
func demoGoroutineLeakPrevention() {
	fmt.Printf("演示前Goroutine数量: %d\n", runtime.NumGoroutine())

	// 1. 使用context控制goroutine生命周期
	fmt.Println("使用channel控制Goroutine:")

	stop := make(chan bool)
	done := make(chan bool)

	go func() {
		defer func() { done <- true }()

		ticker := time.NewTicker(20 * time.Millisecond)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				fmt.Println("定期任务执行中...")
			case <-stop:
				fmt.Println("收到停止信号，Goroutine退出")
				return
			}
		}
	}()

	// 让goroutine运行一段时间
	time.Sleep(100 * time.Millisecond)

	// 发送停止信号
	stop <- true
	<-done // 等待goroutine完全退出

	fmt.Printf("演示后Goroutine数量: %d\n", runtime.NumGoroutine())

	// 2. 演示超时控制
	fmt.Println("\n超时控制示例:")
	demoTimeoutControl()
}

// demoTimeoutControl 演示超时控制
func demoTimeoutControl() {
	timeout := time.After(50 * time.Millisecond)
	done := make(chan bool)

	go func() {
		// 模拟一个可能很慢的操作
		time.Sleep(100 * time.Millisecond)
		done <- true
	}()

	select {
	case <-done:
		fmt.Println("操作完成")
	case <-timeout:
		fmt.Println("操作超时")
	}
}

// demoRuntimeInfo 演示运行时信息
func demoRuntimeInfo() {
	fmt.Printf("CPU核心数: %d\n", runtime.NumCPU())
	fmt.Printf("当前Goroutine数量: %d\n", runtime.NumGoroutine())
	fmt.Printf("Go版本: %s\n", runtime.Version())
	fmt.Printf("操作系统: %s\n", runtime.GOOS)
	fmt.Printf("架构: %s\n", runtime.GOARCH)

	// 设置使用的CPU核心数
	fmt.Printf("GOMAXPROCS (当前): %d\n", runtime.GOMAXPROCS(0))

	// 内存统计
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("分配的内存: %d KB\n", m.Alloc/1024)
	fmt.Printf("总分配的内存: %d KB\n", m.TotalAlloc/1024)
	fmt.Printf("系统内存: %d KB\n", m.Sys/1024)
	fmt.Printf("GC次数: %d\n", m.NumGC)

	// 演示Goroutine调度
	fmt.Println("\nGoroutine调度演示:")
	demoGoroutineScheduling()
}

// demoGoroutineScheduling 演示Goroutine调度
func demoGoroutineScheduling() {
	var wg sync.WaitGroup

	// CPU密集型任务
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			fmt.Printf("CPU任务 %d 开始\n", id)

			// 模拟CPU密集型工作
			count := 0
			for j := 0; j < 1000000; j++ {
				count++
				if j%200000 == 0 {
					runtime.Gosched() // 主动让出CPU
				}
			}

			fmt.Printf("CPU任务 %d 完成，计算结果: %d\n", id, count)
		}(i)
	}

	// I/O密集型任务
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			fmt.Printf("I/O任务 %d 开始\n", id)

			// 模拟I/O等待
			time.Sleep(50 * time.Millisecond)

			fmt.Printf("I/O任务 %d 完成\n", id)
		}(i)
	}

	wg.Wait()
	fmt.Println("所有调度任务完成")
}

// DemoChannels 演示Channel基础
func DemoChannels() {
	fmt.Println("第4阶段 - Channel基础演示（待实现）")
}

// DemoSelect 演示Select多路复用
func DemoSelect() {
	fmt.Println("第4阶段 - Select多路复用演示（待实现）")
}

// DemoMutex 演示互斥锁
func DemoMutex() {
	fmt.Println("第4阶段 - 互斥锁演示（待实现）")
}

// DemoOnce 演示sync.Once
func DemoOnce() {
	fmt.Println("第4阶段 - sync.Once演示（待实现）")
}

// DemoChannelPatterns 演示Channel模式与Pipeline
func DemoChannelPatterns() {
	fmt.Println("第4阶段 - Channel模式与Pipeline演示（待实现）")
}
