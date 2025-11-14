package stage4

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

// RunStage4 运行第4阶段演示
func RunStage4() {
	fmt.Println("第4阶段：并发编程")

	// Goroutine基础演示
	DemoGoroutines()

	// Channel演示
	DemoChannels()

	// Select多路复用演示
	DemoSelect()

	// 互斥锁演示
	DemoMutex()

	// Context上下文演示
	DemoContext()

	// 并发模式演示
	DemoConcurrencyPatterns()
}

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
	fmt.Println("\n=== Channel基础演示 ===")

	// 1. 基本Channel使用
	fmt.Println("\n1. 基本Channel使用：")
	demoBasicChannel()

	// 2. 缓冲Channel
	fmt.Println("\n2. 缓冲Channel：")
	demoBufferedChannel()

	// 3. Channel方向
	fmt.Println("\n3. Channel方向：")
	demoChannelDirection()

	// 4. Channel关闭
	fmt.Println("\n4. Channel关闭：")
	demoChannelClose()

	// 5. Range遍历Channel
	fmt.Println("\n5. Range遍历Channel：")
	demoChannelRange()

	// 6. Channel模式
	fmt.Println("\n6. Channel模式：")
	demoChannelPatterns()
}

// demoBasicChannel 演示基本Channel使用
func demoBasicChannel() {
	// 1. 创建无缓冲channel
	ch := make(chan string)

	// 2. 在goroutine中发送数据
	go func() {
		time.Sleep(50 * time.Millisecond)
		ch <- "Hello from goroutine!"
	}()

	// 3. 接收数据
	message := <-ch
	fmt.Printf("接收到消息: %s\n", message)

	// 4. 双向通信
	fmt.Println("\n双向通信:")
	requestCh := make(chan string)
	responseCh := make(chan string)

	// 启动服务goroutine
	go func() {
		request := <-requestCh
		fmt.Printf("服务端收到请求: %s\n", request)
		responseCh <- "处理完成: " + request
	}()

	// 发送请求
	requestCh <- "计算任务"
	response := <-responseCh
	fmt.Printf("客户端收到响应: %s\n", response)

	// 5. 同步使用
	fmt.Println("\n同步使用:")
	done := make(chan bool)

	go func() {
		fmt.Println("执行异步任务...")
		time.Sleep(100 * time.Millisecond)
		fmt.Println("异步任务完成")
		done <- true
	}()

	fmt.Println("等待异步任务完成...")
	<-done
	fmt.Println("主程序继续执行")
}

// demoBufferedChannel 演示缓冲Channel
func demoBufferedChannel() {
	// 1. 创建缓冲channel
	ch := make(chan int, 3)

	fmt.Printf("Channel容量: %d, 当前长度: %d\n", cap(ch), len(ch))

	// 2. 发送数据（不会阻塞）
	ch <- 1
	ch <- 2
	ch <- 3

	fmt.Printf("发送3个数据后 - 容量: %d, 当前长度: %d\n", cap(ch), len(ch))

	// 3. 接收数据
	for i := 0; i < 3; i++ {
		value := <-ch
		fmt.Printf("接收到: %d, 剩余长度: %d\n", value, len(ch))
	}

	// 4. 非阻塞发送和接收
	fmt.Println("\n非阻塞操作:")

	// 缓冲channel的非阻塞发送
	select {
	case ch <- 100:
		fmt.Println("成功发送100")
	default:
		fmt.Println("发送失败，channel已满")
	}

	// 非阻塞接收
	select {
	case value := <-ch:
		fmt.Printf("成功接收: %d\n", value)
	default:
		fmt.Println("接收失败，channel为空")
	}

	// 5. 生产者-消费者模式
	fmt.Println("\n生产者-消费者模式:")
	demoProducerConsumer()
}

// demoProducerConsumer 演示生产者-消费者模式
func demoProducerConsumer() {
	buffer := make(chan int, 5)
	var wg sync.WaitGroup

	// 生产者
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			buffer <- i
			fmt.Printf("生产者: 生产 %d\n", i)
			time.Sleep(20 * time.Millisecond)
		}
		close(buffer)
		fmt.Println("生产者: 完成生产")
	}()

	// 消费者
	wg.Add(1)
	go func() {
		defer wg.Done()
		for item := range buffer {
			fmt.Printf("消费者: 消费 %d\n", item)
			time.Sleep(50 * time.Millisecond)
		}
		fmt.Println("消费者: 完成消费")
	}()

	wg.Wait()
}

// demoChannelDirection 演示Channel方向
func demoChannelDirection() {
	// 双向channel
	ch := make(chan string, 1)

	// 只发送channel
	go sendOnly(ch)

	// 只接收channel
	go receiveOnly(ch)

	time.Sleep(100 * time.Millisecond)

	// 管道模式
	fmt.Println("\n管道模式:")
	demoPipeline()
}

// sendOnly 只能发送的channel
func sendOnly(ch chan<- string) {
	ch <- "只发送channel的消息"
	fmt.Println("发送完成")
}

// receiveOnly 只能接收的channel
func receiveOnly(ch <-chan string) {
	message := <-ch
	fmt.Printf("接收到: %s\n", message)
}

// demoPipeline 演示管道模式
func demoPipeline() {
	// 创建管道
	numbers := make(chan int)
	squares := make(chan int)

	// 第一阶段：生成数字
	go func() {
		defer close(numbers)
		for i := 1; i <= 5; i++ {
			numbers <- i
			fmt.Printf("生成数字: %d\n", i)
		}
	}()

	// 第二阶段：计算平方
	go func() {
		defer close(squares)
		for num := range numbers {
			square := num * num
			squares <- square
			fmt.Printf("计算平方: %d -> %d\n", num, square)
		}
	}()

	// 第三阶段：输出结果
	fmt.Println("最终结果:")
	for square := range squares {
		fmt.Printf("平方值: %d\n", square)
	}
}

// demoChannelClose 演示Channel关闭
func demoChannelClose() {
	ch := make(chan int, 3)

	// 发送数据
	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i
			fmt.Printf("发送: %d\n", i)
		}
		close(ch) // 关闭channel
		fmt.Println("Channel已关闭")
	}()

	// 接收数据，检查channel是否关闭
	for {
		value, ok := <-ch
		if !ok {
			fmt.Println("Channel已关闭，退出接收")
			break
		}
		fmt.Printf("接收: %d\n", value)
	}

	// 演示向已关闭的channel发送数据会panic
	fmt.Println("\n关闭状态检查:")
	testCh := make(chan int)
	close(testCh)

	// 从已关闭的channel接收（安全）
	value, ok := <-testCh
	fmt.Printf("从已关闭channel接收: value=%d, ok=%t\n", value, ok)
}

// demoChannelRange 演示Range遍历Channel
func demoChannelRange() {
	ch := make(chan string, 3)

	// 发送数据并关闭
	go func() {
		fruits := []string{"苹果", "香蕉", "橙子", "葡萄"}
		for _, fruit := range fruits {
			ch <- fruit
			fmt.Printf("发送水果: %s\n", fruit)
			time.Sleep(30 * time.Millisecond)
		}
		close(ch)
	}()

	// 使用range遍历channel
	fmt.Println("使用range遍历channel:")
	for fruit := range ch {
		fmt.Printf("收到水果: %s\n", fruit)
	}

	fmt.Println("遍历完成")
}

// demoChannelPatterns 演示Channel模式
func demoChannelPatterns() {
	// 1. Fan-out模式（一个输入，多个输出）
	fmt.Println("Fan-out模式:")
	demoFanOut()

	// 2. Fan-in模式（多个输入，一个输出）
	fmt.Println("\nFan-in模式:")
	demoFanIn()

	// 3. 工作池模式
	fmt.Println("\n工作池模式:")
	demoWorkerPool()
}

// demoFanOut 演示Fan-out模式
func demoFanOut() {
	input := make(chan int)
	output1 := make(chan int)
	output2 := make(chan int)

	// 输入数据
	go func() {
		defer close(input)
		for i := 1; i <= 6; i++ {
			input <- i
		}
	}()

	// Fan-out：分发到两个输出
	go func() {
		defer close(output1)
		defer close(output2)

		for num := range input {
			if num%2 == 0 {
				output1 <- num
			} else {
				output2 <- num
			}
		}
	}()

	// 处理输出
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for num := range output1 {
			fmt.Printf("偶数处理器: %d\n", num)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for num := range output2 {
			fmt.Printf("奇数处理器: %d\n", num)
		}
	}()

	wg.Wait()
}

// demoFanIn 演示Fan-in模式
func demoFanIn() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	output := make(chan string)

	// 输入源1
	go func() {
		defer close(ch1)
		for i := 1; i <= 3; i++ {
			ch1 <- fmt.Sprintf("源1-消息%d", i)
			time.Sleep(50 * time.Millisecond)
		}
	}()

	// 输入源2
	go func() {
		defer close(ch2)
		for i := 1; i <= 3; i++ {
			ch2 <- fmt.Sprintf("源2-消息%d", i)
			time.Sleep(70 * time.Millisecond)
		}
	}()

	// Fan-in：合并多个输入
	go func() {
		defer close(output)
		var wg sync.WaitGroup

		wg.Add(1)
		go func() {
			defer wg.Done()
			for msg := range ch1 {
				output <- msg
			}
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
			for msg := range ch2 {
				output <- msg
			}
		}()

		wg.Wait()
	}()

	// 处理合并后的输出
	for msg := range output {
		fmt.Printf("合并输出: %s\n", msg)
	}
}

// demoWorkerPool 演示工作池模式
func demoWorkerPool() {
	const numWorkers = 3
	const numJobs = 10

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// 启动工作者
	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results)
	}

	// 发送任务
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// 收集结果
	for r := 1; r <= numJobs; r++ {
		result := <-results
		fmt.Printf("任务结果: %d\n", result)
	}
}

// worker 工作者函数
func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("工作者 %d 开始任务 %d\n", id, job)
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		result := job * 2
		fmt.Printf("工作者 %d 完成任务 %d，结果: %d\n", id, job, result)
		results <- result
	}
}

// DemoSelect 演示Select多路复用
func DemoSelect() {
	fmt.Println("\n=== Select多路复用演示 ===")

	// 1. 基本Select使用
	fmt.Println("\n1. 基本Select使用：")
	demoBasicSelect()

	// 2. Select超时控制
	fmt.Println("\n2. Select超时控制：")
	demoSelectTimeout()

	// 3. 非阻塞Select
	fmt.Println("\n3. 非阻塞Select：")
	demoNonBlockingSelect()

	// 4. Select随机选择
	fmt.Println("\n4. Select随机选择：")
	demoSelectRandom()

	// 5. Select与Channel关闭
	fmt.Println("\n5. Select与Channel关闭：")
	demoSelectWithClose()

	// 6. 复杂Select模式
	fmt.Println("\n6. 复杂Select模式：")
	demoComplexSelect()
}

// demoBasicSelect 演示基本Select使用
func demoBasicSelect() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	// 启动两个goroutine
	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "来自channel1的消息"
	}()

	go func() {
		time.Sleep(150 * time.Millisecond)
		ch2 <- "来自channel2的消息"
	}()

	// 使用select等待第一个可用的channel
	select {
	case msg1 := <-ch1:
		fmt.Printf("收到: %s\n", msg1)
	case msg2 := <-ch2:
		fmt.Printf("收到: %s\n", msg2)
	}

	// 接收剩余的消息
	select {
	case msg1 := <-ch1:
		fmt.Printf("收到剩余: %s\n", msg1)
	case msg2 := <-ch2:
		fmt.Printf("收到剩余: %s\n", msg2)
	}
}

// demoSelectTimeout 演示Select超时控制
func demoSelectTimeout() {
	ch := make(chan string)

	// 启动一个慢速的goroutine
	go func() {
		time.Sleep(200 * time.Millisecond)
		ch <- "慢速消息"
	}()

	// 使用select实现超时
	select {
	case msg := <-ch:
		fmt.Printf("收到消息: %s\n", msg)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("操作超时")
	}

	// 演示不同的超时时间
	fmt.Println("\n不同超时时间测试:")
	testTimeouts := []time.Duration{50 * time.Millisecond, 300 * time.Millisecond}

	for i, timeout := range testTimeouts {
		ch := make(chan string)

		go func() {
			time.Sleep(150 * time.Millisecond)
			ch <- fmt.Sprintf("消息%d", i+1)
		}()

		select {
		case msg := <-ch:
			fmt.Printf("超时%v: 收到 %s\n", timeout, msg)
		case <-time.After(timeout):
			fmt.Printf("超时%v: 操作超时\n", timeout)
		}
	}
}

// demoNonBlockingSelect 演示非阻塞Select
func demoNonBlockingSelect() {
	ch := make(chan string, 1)

	// 非阻塞发送
	select {
	case ch <- "非阻塞发送":
		fmt.Println("成功发送消息")
	default:
		fmt.Println("发送失败，channel已满")
	}

	// 非阻塞接收
	select {
	case msg := <-ch:
		fmt.Printf("成功接收: %s\n", msg)
	default:
		fmt.Println("接收失败，channel为空")
	}

	// 再次尝试非阻塞接收
	select {
	case msg := <-ch:
		fmt.Printf("成功接收: %s\n", msg)
	default:
		fmt.Println("接收失败，channel为空")
	}

	// 演示非阻塞的实际应用
	fmt.Println("\n非阻塞轮询:")
	demoNonBlockingPolling()
}

// demoNonBlockingPolling 演示非阻塞轮询
func demoNonBlockingPolling() {
	dataCh := make(chan int, 5)
	controlCh := make(chan bool)

	// 数据生产者
	go func() {
		for i := 1; i <= 10; i++ {
			dataCh <- i
			time.Sleep(50 * time.Millisecond)
		}
		controlCh <- true
	}()

	// 非阻塞轮询消费者
	for {
		select {
		case data := <-dataCh:
			fmt.Printf("处理数据: %d\n", data)
		case <-controlCh:
			fmt.Println("收到停止信号")
			// 处理剩余数据
			for {
				select {
				case data := <-dataCh:
					fmt.Printf("处理剩余数据: %d\n", data)
				default:
					fmt.Println("所有数据处理完成")
					return
				}
			}
		default:
			fmt.Println("暂无数据，执行其他任务...")
			time.Sleep(30 * time.Millisecond)
		}
	}
}

// demoSelectRandom 演示Select随机选择
func demoSelectRandom() {
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)
	ch3 := make(chan string, 1)

	// 同时向所有channel发送数据
	ch1 <- "Channel 1"
	ch2 <- "Channel 2"
	ch3 <- "Channel 3"

	// Select会随机选择一个可用的case
	fmt.Println("随机选择测试（运行多次）:")
	for i := 0; i < 5; i++ {
		// 重新填充channel
		select {
		case <-ch1:
		default:
		}
		select {
		case <-ch2:
		default:
		}
		select {
		case <-ch3:
		default:
		}

		ch1 <- "Channel 1"
		ch2 <- "Channel 2"
		ch3 <- "Channel 3"

		// 随机选择
		select {
		case msg := <-ch1:
			fmt.Printf("第%d次选择: %s\n", i+1, msg)
		case msg := <-ch2:
			fmt.Printf("第%d次选择: %s\n", i+1, msg)
		case msg := <-ch3:
			fmt.Printf("第%d次选择: %s\n", i+1, msg)
		}
	}
}

// demoSelectWithClose 演示Select与Channel关闭
func demoSelectWithClose() {
	ch1 := make(chan int)
	ch2 := make(chan string)
	done := make(chan bool)

	// 数据发送者
	go func() {
		for i := 1; i <= 5; i++ {
			ch1 <- i
			time.Sleep(50 * time.Millisecond)
		}
		close(ch1)
	}()

	go func() {
		messages := []string{"A", "B", "C"}
		for _, msg := range messages {
			ch2 <- msg
			time.Sleep(70 * time.Millisecond)
		}
		close(ch2)
	}()

	// 监控goroutine
	go func() {
		time.Sleep(400 * time.Millisecond)
		done <- true
	}()

	// 使用select处理多个channel的关闭
	ch1Open := true
	ch2Open := true

	for ch1Open || ch2Open {
		select {
		case num, ok := <-ch1:
			if !ok {
				fmt.Println("Channel 1 已关闭")
				ch1Open = false
			} else {
				fmt.Printf("从Channel 1收到: %d\n", num)
			}
		case msg, ok := <-ch2:
			if !ok {
				fmt.Println("Channel 2 已关闭")
				ch2Open = false
			} else {
				fmt.Printf("从Channel 2收到: %s\n", msg)
			}
		case <-done:
			fmt.Println("超时，强制退出")
			return
		}
	}

	fmt.Println("所有channel都已关闭")
}

// demoComplexSelect 演示复杂Select模式
func demoComplexSelect() {
	// 心跳监控系统
	fmt.Println("心跳监控系统:")
	demoHeartbeatMonitor()

	// 请求合并系统
	fmt.Println("\n请求合并系统:")
	demoRequestBatcher()
}

// demoHeartbeatMonitor 演示心跳监控
func demoHeartbeatMonitor() {
	heartbeat := make(chan bool)
	shutdown := make(chan bool)

	// 心跳发送者
	go func() {
		ticker := time.NewTicker(100 * time.Millisecond)
		defer ticker.Stop()

		for i := 0; i < 5; i++ {
			select {
			case <-ticker.C:
				heartbeat <- true
				fmt.Printf("发送心跳 %d\n", i+1)
			case <-shutdown:
				fmt.Println("心跳发送者收到关闭信号")
				return
			}
		}
	}()

	// 心跳监控者
	go func() {
		timeout := time.NewTimer(150 * time.Millisecond)
		defer timeout.Stop()

		for {
			timeout.Reset(150 * time.Millisecond)

			select {
			case <-heartbeat:
				fmt.Println("收到心跳，系统正常")
			case <-timeout.C:
				fmt.Println("心跳超时，系统异常！")
				shutdown <- true
				return
			case <-shutdown:
				fmt.Println("监控者收到关闭信号")
				return
			}
		}
	}()

	time.Sleep(600 * time.Millisecond)
	close(shutdown)
}

// demoRequestBatcher 演示请求合并
func demoRequestBatcher() {
	requests := make(chan string, 10)
	batchSize := 3
	batchTimeout := 200 * time.Millisecond

	// 请求发送者
	go func() {
		defer close(requests)
		for i := 1; i <= 8; i++ {
			requests <- fmt.Sprintf("请求%d", i)
			time.Sleep(50 * time.Millisecond)
		}
	}()

	// 批处理器
	var batch []string
	batchTimer := time.NewTimer(batchTimeout)
	batchTimer.Stop()

	for {
		select {
		case req, ok := <-requests:
			if !ok {
				// 处理最后一批
				if len(batch) > 0 {
					fmt.Printf("处理最后一批: %v\n", batch)
				}
				return
			}

			batch = append(batch, req)

			// 启动或重置定时器
			if len(batch) == 1 {
				batchTimer.Reset(batchTimeout)
			}

			// 检查是否达到批大小
			if len(batch) >= batchSize {
				fmt.Printf("批大小达到，处理批次: %v\n", batch)
				batch = nil
				batchTimer.Stop()
			}

		case <-batchTimer.C:
			// 超时，处理当前批次
			if len(batch) > 0 {
				fmt.Printf("超时处理批次: %v\n", batch)
				batch = nil
			}
		}
	}
}
