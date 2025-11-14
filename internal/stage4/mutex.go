package stage4

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// DemoMutex 演示互斥锁
func DemoMutex() {
	fmt.Println("\n=== 互斥锁演示 ===")

	// 1. 基本Mutex使用
	fmt.Println("\n1. 基本Mutex使用：")
	demoBasicMutex()

	// 2. RWMutex读写锁
	fmt.Println("\n2. RWMutex读写锁：")
	demoRWMutex()

	// 3. 原子操作
	fmt.Println("\n3. 原子操作：")
	demoAtomicOperations()

	// 4. 条件变量
	fmt.Println("\n4. 条件变量：")
	demoCondition()

	// 5. Once单次执行
	fmt.Println("\n5. Once单次执行：")
	demoOnce()

	// 6. 同步原语比较
	fmt.Println("\n6. 同步原语比较：")
	demoSyncComparison()
}

// Counter 计数器结构体
type Counter struct {
	mu    sync.Mutex
	value int
}

// Increment 增加计数
func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// Value 获取当前值
func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

// demoBasicMutex 演示基本Mutex使用
func demoBasicMutex() {
	// 1. 不使用锁的竞态条件
	fmt.Println("不使用锁的竞态条件:")
	unsafeCounter := 0
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				unsafeCounter++ // 竞态条件
			}
		}()
	}
	wg.Wait()
	fmt.Printf("不安全计数器结果: %d (期望: 10000)\n", unsafeCounter)

	// 2. 使用Mutex保护
	fmt.Println("\n使用Mutex保护:")
	safeCounter := &Counter{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				safeCounter.Increment()
			}
		}()
	}
	wg.Wait()
	fmt.Printf("安全计数器结果: %d (期望: 10000)\n", safeCounter.Value())

	// 3. 死锁演示（注释掉避免程序卡死）
	fmt.Println("\n死锁预防:")
	demoDeadlockPrevention()
}

// demoDeadlockPrevention 演示死锁预防
func demoDeadlockPrevention() {
	var mu1, mu2 sync.Mutex

	// 正确的锁顺序
	fmt.Println("正确的锁顺序:")
	var wg sync.WaitGroup

	// Goroutine 1
	wg.Add(1)
	go func() {
		defer wg.Done()
		mu1.Lock()
		fmt.Println("Goroutine 1: 获得锁1")
		time.Sleep(10 * time.Millisecond)

		mu2.Lock()
		fmt.Println("Goroutine 1: 获得锁2")
		mu2.Unlock()
		mu1.Unlock()
		fmt.Println("Goroutine 1: 释放所有锁")
	}()

	// Goroutine 2 - 相同的锁顺序
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(5 * time.Millisecond) // 稍微延迟

		mu1.Lock()
		fmt.Println("Goroutine 2: 获得锁1")
		time.Sleep(10 * time.Millisecond)

		mu2.Lock()
		fmt.Println("Goroutine 2: 获得锁2")
		mu2.Unlock()
		mu1.Unlock()
		fmt.Println("Goroutine 2: 释放所有锁")
	}()

	wg.Wait()
	fmt.Println("所有goroutine完成，无死锁")
}

// Cache 缓存结构体
type Cache struct {
	mu   sync.RWMutex
	data map[string]string
}

// NewCache 创建新缓存
func NewCache() *Cache {
	return &Cache{
		data: make(map[string]string),
	}
}

// Get 读取缓存
func (c *Cache) Get(key string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, ok := c.data[key]
	return value, ok
}

// Set 设置缓存
func (c *Cache) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
}

// demoRWMutex 演示读写锁
func demoRWMutex() {
	cache := NewCache()
	var wg sync.WaitGroup

	// 写入数据
	fmt.Println("写入初始数据:")
	for i := 0; i < 5; i++ {
		key := fmt.Sprintf("key%d", i)
		value := fmt.Sprintf("value%d", i)
		cache.Set(key, value)
		fmt.Printf("设置 %s = %s\n", key, value)
	}

	// 启动多个读取者
	fmt.Println("\n启动多个读取者:")
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 3; j++ {
				key := fmt.Sprintf("key%d", j)
				if value, ok := cache.Get(key); ok {
					fmt.Printf("读取者%d: %s = %s\n", id, key, value)
				}
				time.Sleep(10 * time.Millisecond)
			}
		}(i)
	}

	// 启动写入者
	fmt.Println("启动写入者:")
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 5; i < 8; i++ {
			key := fmt.Sprintf("key%d", i)
			value := fmt.Sprintf("value%d", i)
			cache.Set(key, value)
			fmt.Printf("写入者: 设置 %s = %s\n", key, value)
			time.Sleep(20 * time.Millisecond)
		}
	}()

	wg.Wait()
	fmt.Println("读写操作完成")
}

// demoAtomicOperations 演示原子操作
func demoAtomicOperations() {
	// 1. 原子计数器
	fmt.Println("原子计数器:")
	var atomicCounter int64
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&atomicCounter, 1)
			}
		}()
	}
	wg.Wait()
	fmt.Printf("原子计数器结果: %d\n", atomic.LoadInt64(&atomicCounter))

	// 2. 原子交换
	fmt.Println("\n原子交换:")
	var value int64 = 100
	fmt.Printf("初始值: %d\n", value)

	old := atomic.SwapInt64(&value, 200)
	fmt.Printf("交换后: 新值=%d, 旧值=%d\n", value, old)

	// 3. 比较并交换
	fmt.Println("\n比较并交换:")
	var cas int64 = 200

	// 成功的CAS
	if atomic.CompareAndSwapInt64(&cas, 200, 300) {
		fmt.Printf("CAS成功: %d -> 300\n", 200)
	}

	// 失败的CAS
	if !atomic.CompareAndSwapInt64(&cas, 200, 400) {
		fmt.Printf("CAS失败: 期望200，实际%d\n", cas)
	}

	// 4. 原子指针操作
	fmt.Println("\n原子指针操作:")
	demoAtomicPointer()
}

// demoAtomicPointer 演示原子指针操作
func demoAtomicPointer() {
	type Config struct {
		Name    string
		Version int
	}

	var configPtr atomic.Value

	// 初始配置
	initialConfig := &Config{Name: "App", Version: 1}
	configPtr.Store(initialConfig)

	// 读取配置
	config := configPtr.Load().(*Config)
	fmt.Printf("当前配置: %+v\n", config)

	// 更新配置
	newConfig := &Config{Name: "App", Version: 2}
	configPtr.Store(newConfig)

	// 再次读取
	config = configPtr.Load().(*Config)
	fmt.Printf("更新后配置: %+v\n", config)
}

// demoCondition 演示条件变量
func demoCondition() {
	var mu sync.Mutex
	cond := sync.NewCond(&mu)
	ready := false
	data := make([]int, 0)

	// 消费者
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			mu.Lock()
			defer mu.Unlock()

			// 等待条件满足
			for !ready {
				fmt.Printf("消费者%d: 等待数据准备\n", id)
				cond.Wait()
			}

			// 处理数据
			fmt.Printf("消费者%d: 处理数据 %v\n", id, data)
		}(i)
	}

	// 生产者
	go func() {
		time.Sleep(100 * time.Millisecond)

		mu.Lock()
		// 准备数据
		data = []int{1, 2, 3, 4, 5}
		ready = true
		fmt.Println("生产者: 数据准备完成")
		mu.Unlock()

		// 通知所有等待的goroutine
		cond.Broadcast()
	}()

	wg.Wait()
	fmt.Println("条件变量演示完成")
}

// demoOnce 演示Once单次执行
func demoOnce() {
	var once sync.Once
	var initialized bool

	initialize := func() {
		fmt.Println("执行初始化操作...")
		time.Sleep(50 * time.Millisecond)
		initialized = true
		fmt.Println("初始化完成")
	}

	var wg sync.WaitGroup

	// 启动多个goroutine尝试初始化
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Goroutine %d: 尝试初始化\n", id)
			once.Do(initialize)
			fmt.Printf("Goroutine %d: 初始化状态 = %t\n", id, initialized)
		}(i)
	}

	wg.Wait()
	fmt.Println("Once演示完成")
}

// demoSyncComparison 演示同步原语比较
func demoSyncComparison() {
	const iterations = 100000

	// 1. Mutex性能测试
	fmt.Println("Mutex性能测试:")
	start := time.Now()
	var mu sync.Mutex
	var mutexCounter int

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < iterations/10; j++ {
				mu.Lock()
				mutexCounter++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	mutexTime := time.Since(start)
	fmt.Printf("Mutex: %v, 结果: %d\n", mutexTime, mutexCounter)

	// 2. 原子操作性能测试
	fmt.Println("\n原子操作性能测试:")
	start = time.Now()
	var atomicCounter int64

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < iterations/10; j++ {
				atomic.AddInt64(&atomicCounter, 1)
			}
		}()
	}
	wg.Wait()
	atomicTime := time.Since(start)
	fmt.Printf("原子操作: %v, 结果: %d\n", atomicTime, atomicCounter)

	// 3. Channel性能测试
	fmt.Println("\nChannel性能测试:")
	start = time.Now()
	ch := make(chan int, 1000)
	var channelCounter int

	// 发送者
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(ch)
		for i := 0; i < iterations; i++ {
			ch <- 1
		}
	}()

	// 接收者
	wg.Add(1)
	go func() {
		defer wg.Done()
		for range ch {
			channelCounter++
		}
	}()

	wg.Wait()
	channelTime := time.Since(start)
	fmt.Printf("Channel: %v, 结果: %d\n", channelTime, channelCounter)

	// 性能比较
	fmt.Printf("\n性能比较 (相对于Mutex):\n")
	fmt.Printf("Mutex: 1.00x\n")
	fmt.Printf("原子操作: %.2fx\n", float64(mutexTime)/float64(atomicTime))
	fmt.Printf("Channel: %.2fx\n", float64(mutexTime)/float64(channelTime))
}
