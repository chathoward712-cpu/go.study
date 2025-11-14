package stage4

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// DemoConcurrencyPatterns 演示并发模式
func DemoConcurrencyPatterns() {
	fmt.Println("\n=== 并发模式演示 ===")

	// 1. 生产者-消费者模式
	fmt.Println("\n1. 生产者-消费者模式：")
	demoProducerConsumerPattern()

	// 2. 发布-订阅模式
	fmt.Println("\n2. 发布-订阅模式：")
	demoPubSubPattern()

	// 3. 工作池模式
	fmt.Println("\n3. 工作池模式：")
	demoWorkerPoolPattern()

	// 4. 管道模式
	fmt.Println("\n4. 管道模式：")
	demoPipelinePattern()

	// 5. 扇入扇出模式
	fmt.Println("\n5. 扇入扇出模式：")
	demoFanInFanOutPattern()

	// 6. 限流模式
	fmt.Println("\n6. 限流模式：")
	demoRateLimitingPattern()

	// 7. 超时模式
	fmt.Println("\n7. 超时模式：")
	demoTimeoutPattern()
}

// demoProducerConsumerPattern 演示生产者-消费者模式
func demoProducerConsumerPattern() {
	buffer := make(chan int, 5)
	var wg sync.WaitGroup

	// 启动生产者
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(buffer)

		for i := 1; i <= 10; i++ {
			fmt.Printf("生产者: 生产商品 %d\n", i)
			buffer <- i
			time.Sleep(100 * time.Millisecond)
		}
		fmt.Println("生产者: 生产完成")
	}()

	// 启动多个消费者
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			for item := range buffer {
				fmt.Printf("消费者%d: 消费商品 %d\n", id, item)
				time.Sleep(150 * time.Millisecond)
			}
			fmt.Printf("消费者%d: 消费完成\n", id)
		}(i)
	}

	wg.Wait()
}

// PubSub 发布订阅系统
type PubSub struct {
	mu          sync.RWMutex
	subscribers map[string][]chan string
}

// NewPubSub 创建新的发布订阅系统
func NewPubSub() *PubSub {
	return &PubSub{
		subscribers: make(map[string][]chan string),
	}
}

// Subscribe 订阅主题
func (ps *PubSub) Subscribe(topic string) <-chan string {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	ch := make(chan string, 10)
	ps.subscribers[topic] = append(ps.subscribers[topic], ch)
	return ch
}

// Publish 发布消息
func (ps *PubSub) Publish(topic, message string) {
	ps.mu.RLock()
	defer ps.mu.RUnlock()

	for _, ch := range ps.subscribers[topic] {
		select {
		case ch <- message:
		default:
			// 如果channel满了，跳过这个订阅者
		}
	}
}

// demoPubSubPattern 演示发布-订阅模式
func demoPubSubPattern() {
	pubsub := NewPubSub()
	var wg sync.WaitGroup

	// 订阅者1 - 订阅新闻
	wg.Add(1)
	go func() {
		defer wg.Done()
		newsCh := pubsub.Subscribe("news")

		for i := 0; i < 3; i++ {
			select {
			case msg := <-newsCh:
				fmt.Printf("新闻订阅者: 收到 %s\n", msg)
			case <-time.After(500 * time.Millisecond):
				fmt.Println("新闻订阅者: 超时退出")
				return
			}
		}
	}()

	// 订阅者2 - 订阅体育
	wg.Add(1)
	go func() {
		defer wg.Done()
		sportsCh := pubsub.Subscribe("sports")

		for i := 0; i < 2; i++ {
			select {
			case msg := <-sportsCh:
				fmt.Printf("体育订阅者: 收到 %s\n", msg)
			case <-time.After(500 * time.Millisecond):
				fmt.Println("体育订阅者: 超时退出")
				return
			}
		}
	}()

	// 订阅者3 - 订阅新闻和体育
	wg.Add(1)
	go func() {
		defer wg.Done()
		newsCh := pubsub.Subscribe("news")
		sportsCh := pubsub.Subscribe("sports")

		for i := 0; i < 4; i++ {
			select {
			case msg := <-newsCh:
				fmt.Printf("综合订阅者: 收到新闻 %s\n", msg)
			case msg := <-sportsCh:
				fmt.Printf("综合订阅者: 收到体育 %s\n", msg)
			case <-time.After(500 * time.Millisecond):
				fmt.Println("综合订阅者: 超时退出")
				return
			}
		}
	}()

	// 发布者
	go func() {
		time.Sleep(50 * time.Millisecond)
		pubsub.Publish("news", "重要新闻1")

		time.Sleep(100 * time.Millisecond)
		pubsub.Publish("sports", "体育新闻1")

		time.Sleep(100 * time.Millisecond)
		pubsub.Publish("news", "重要新闻2")

		time.Sleep(100 * time.Millisecond)
		pubsub.Publish("sports", "体育新闻2")

		time.Sleep(100 * time.Millisecond)
		pubsub.Publish("news", "重要新闻3")
	}()

	wg.Wait()
}

// Job 工作任务
type Job struct {
	ID   int
	Data string
}

// Result 工作结果
type Result struct {
	Job    Job
	Output string
	Error  error
}

// WorkerPool 工作池
type WorkerPool struct {
	jobs    chan Job
	results chan Result
	workers int
}

// NewWorkerPool 创建工作池
func NewWorkerPool(workers int) *WorkerPool {
	return &WorkerPool{
		jobs:    make(chan Job, 100),
		results: make(chan Result, 100),
		workers: workers,
	}
}

// Start 启动工作池
func (wp *WorkerPool) Start() {
	for i := 0; i < wp.workers; i++ {
		go wp.worker(i)
	}
}

// worker 工作者
func (wp *WorkerPool) worker(id int) {
	for job := range wp.jobs {
		fmt.Printf("工作者%d: 开始处理任务%d\n", id, job.ID)

		// 模拟工作
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)

		result := Result{
			Job:    job,
			Output: fmt.Sprintf("处理结果: %s", job.Data),
		}

		wp.results <- result
		fmt.Printf("工作者%d: 完成任务%d\n", id, job.ID)
	}
}

// AddJob 添加任务
func (wp *WorkerPool) AddJob(job Job) {
	wp.jobs <- job
}

// GetResult 获取结果
func (wp *WorkerPool) GetResult() Result {
	return <-wp.results
}

// Close 关闭工作池
func (wp *WorkerPool) Close() {
	close(wp.jobs)
}

// demoWorkerPoolPattern 演示工作池模式
func demoWorkerPoolPattern() {
	// 创建工作池
	pool := NewWorkerPool(3)
	pool.Start()

	// 添加任务
	go func() {
		for i := 1; i <= 8; i++ {
			job := Job{
				ID:   i,
				Data: fmt.Sprintf("数据%d", i),
			}
			pool.AddJob(job)
		}
		pool.Close()
	}()

	// 收集结果
	for i := 1; i <= 8; i++ {
		result := pool.GetResult()
		fmt.Printf("收到结果: 任务%d -> %s\n", result.Job.ID, result.Output)
	}
}

// demoPipelinePattern 演示管道模式
func demoPipelinePattern() {
	// 第一阶段：生成数字
	numbers := make(chan int)
	go func() {
		defer close(numbers)
		for i := 1; i <= 10; i++ {
			numbers <- i
			fmt.Printf("生成: %d\n", i)
		}
	}()

	// 阶段2：计算平方
	squares := make(chan int)
	go func() {
		defer close(squares)
		for num := range numbers {
			square := num * num
			squares <- square
			fmt.Printf("平方: %d -> %d\n", num, square)
		}
	}()

	// 阶段3：过滤偶数
	evens := make(chan int)
	go func() {
		defer close(evens)
		for square := range squares {
			if square%2 == 0 {
				evens <- square
				fmt.Printf("过滤偶数: %d\n", square)
			}
		}
	}()

	// 阶段4：输出结果
	fmt.Println("最终结果:")
	for even := range evens {
		fmt.Printf("输出: %d\n", even)
	}
}

// demoFanInFanOutPattern 演示扇入扇出模式
func demoFanInFanOutPattern() {
	// 输入源
	input := make(chan int)

	// Fan-out: 分发到多个处理器
	processor1 := make(chan int)
	processor2 := make(chan int)
	processor3 := make(chan int)

	// 分发器
	go func() {
		defer close(processor1)
		defer close(processor2)
		defer close(processor3)

		for num := range input {
			switch num % 3 {
			case 0:
				processor1 <- num
			case 1:
				processor2 <- num
			case 2:
				processor3 <- num
			}
		}
	}()

	// 处理器
	output1 := make(chan string)
	output2 := make(chan string)
	output3 := make(chan string)

	go func() {
		defer close(output1)
		for num := range processor1 {
			result := fmt.Sprintf("处理器1: %d*2=%d", num, num*2)
			output1 <- result
		}
	}()

	go func() {
		defer close(output2)
		for num := range processor2 {
			result := fmt.Sprintf("处理器2: %d*3=%d", num, num*3)
			output2 <- result
		}
	}()

	go func() {
		defer close(output3)
		for num := range processor3 {
			result := fmt.Sprintf("处理器3: %d*4=%d", num, num*4)
			output3 <- result
		}
	}()

	// Fan-in: 合并结果
	output := make(chan string)
	var wg sync.WaitGroup

	wg.Add(3)
	go func() {
		defer wg.Done()
		for result := range output1 {
			output <- result
		}
	}()

	go func() {
		defer wg.Done()
		for result := range output2 {
			output <- result
		}
	}()

	go func() {
		defer wg.Done()
		for result := range output3 {
			output <- result
		}
	}()

	go func() {
		wg.Wait()
		close(output)
	}()

	// 发送输入数据
	go func() {
		defer close(input)
		for i := 1; i <= 9; i++ {
			input <- i
		}
	}()

	// 收集结果
	for result := range output {
		fmt.Println(result)
	}
}

// RateLimiter 限流器
type RateLimiter struct {
	tokens chan struct{}
	ticker *time.Ticker
}

// NewRateLimiter 创建限流器
func NewRateLimiter(rate int, burst int) *RateLimiter {
	rl := &RateLimiter{
		tokens: make(chan struct{}, burst),
		ticker: time.NewTicker(time.Second / time.Duration(rate)),
	}

	// 初始化令牌
	for i := 0; i < burst; i++ {
		rl.tokens <- struct{}{}
	}

	// 定期添加令牌
	go func() {
		for range rl.ticker.C {
			select {
			case rl.tokens <- struct{}{}:
			default:
				// 令牌桶已满
			}
		}
	}()

	return rl
}

// Allow 检查是否允许请求
func (rl *RateLimiter) Allow() bool {
	select {
	case <-rl.tokens:
		return true
	default:
		return false
	}
}

// Stop 停止限流器
func (rl *RateLimiter) Stop() {
	rl.ticker.Stop()
}

// demoRateLimitingPattern 演示限流模式
func demoRateLimitingPattern() {
	// 创建限流器：每秒2个请求，突发3个
	limiter := NewRateLimiter(2, 3)
	defer limiter.Stop()

	// 模拟请求
	for i := 1; i <= 10; i++ {
		if limiter.Allow() {
			fmt.Printf("请求%d: 通过\n", i)
		} else {
			fmt.Printf("请求%d: 被限流\n", i)
		}
		time.Sleep(200 * time.Millisecond)
	}
}

// demoTimeoutPattern 演示超时模式
func demoTimeoutPattern() {
	// 1. 简单超时
	fmt.Println("简单超时:")
	demoSimpleTimeout()

	// 2. 可取消的超时
	fmt.Println("\n可取消的超时:")
	demoCancellableTimeout()

	// 3. 级联超时
	fmt.Println("\n级联超时:")
	demoCascadingTimeout()
}

// demoSimpleTimeout 演示简单超时
func demoSimpleTimeout() {
	result := make(chan string, 1)

	go func() {
		// 模拟耗时操作
		time.Sleep(300 * time.Millisecond)
		result <- "操作完成"
	}()

	select {
	case res := <-result:
		fmt.Printf("结果: %s\n", res)
	case <-time.After(200 * time.Millisecond):
		fmt.Println("操作超时")
	}
}

// demoCancellableTimeout 演示可取消的超时
func demoCancellableTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	result := make(chan string, 1)

	go func() {
		select {
		case <-ctx.Done():
			fmt.Println("操作被取消")
			return
		case <-time.After(300 * time.Millisecond):
			result <- "操作完成"
		}
	}()

	select {
	case res := <-result:
		fmt.Printf("结果: %s\n", res)
	case <-ctx.Done():
		fmt.Printf("超时: %v\n", ctx.Err())
	}
}

// demoCascadingTimeout 演示级联超时
func demoCascadingTimeout() {
	// 总超时时间
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	// 第一步操作
	if err := stepOne(ctx); err != nil {
		fmt.Printf("第一步失败: %v\n", err)
		return
	}

	// 第二步操作
	if err := stepTwo(ctx); err != nil {
		fmt.Printf("第二步失败: %v\n", err)
		return
	}

	fmt.Println("所有步骤完成")
}

// stepOne 第一步操作
func stepOne(ctx context.Context) error {
	// 为这一步设置更短的超时
	stepCtx, cancel := context.WithTimeout(ctx, 200*time.Millisecond)
	defer cancel()

	done := make(chan bool, 1)
	go func() {
		time.Sleep(150 * time.Millisecond)
		done <- true
	}()

	select {
	case <-done:
		fmt.Println("第一步完成")
		return nil
	case <-stepCtx.Done():
		return fmt.Errorf("第一步超时: %w", stepCtx.Err())
	}
}

// stepTwo 第二步操作
func stepTwo(ctx context.Context) error {
	stepCtx, cancel := context.WithTimeout(ctx, 200*time.Millisecond)
	defer cancel()

	done := make(chan bool, 1)
	go func() {
		time.Sleep(150 * time.Millisecond)
		done <- true
	}()

	select {
	case <-done:
		fmt.Println("第二步完成")
		return nil
	case <-stepCtx.Done():
		return fmt.Errorf("第二步超时: %w", stepCtx.Err())
	}
}
