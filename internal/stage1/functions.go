package stage1

import (
	"fmt"
)

// DemoFunctions 演示函数定义与调用
func DemoFunctions() {
	fmt.Println("\n=== 函数定义与调用演示 ===")

	// 1. 基本函数定义和调用
	fmt.Println("\n1. 基本函数：")
	greet("Alice")
	greet("Bob")

	// 2. 带返回值的函数
	fmt.Println("\n2. 带返回值的函数：")
	sum := add(10, 20)
	fmt.Printf("10 + 20 = %d\n", sum)

	// 3. 多返回值函数
	fmt.Println("\n3. 多返回值函数：")
	quotient, remainder := divide(17, 5)
	fmt.Printf("17 ÷ 5 = %d 余 %d\n", quotient, remainder)

	// 4. 命名返回值
	fmt.Println("\n4. 命名返回值：")
	area, perimeter := rectangleStats(5, 3)
	fmt.Printf("矩形(5x3) - 面积: %d, 周长: %d\n", area, perimeter)

	// 5. 可变参数函数
	fmt.Println("\n5. 可变参数函数：")
	fmt.Printf("求和(1,2,3): %d\n", sumAll(1, 2, 3))
	fmt.Printf("求和(1,2,3,4,5): %d\n", sumAll(1, 2, 3, 4, 5))

	numbers := []int{10, 20, 30}
	fmt.Printf("求和切片[10,20,30]: %d\n", sumAll(numbers...))

	// 6. 函数作为值
	fmt.Println("\n6. 函数作为值：")
	var operation func(int, int) int
	operation = add
	fmt.Printf("函数变量调用 add(5, 3): %d\n", operation(5, 3))

	operation = multiply
	fmt.Printf("函数变量调用 multiply(5, 3): %d\n", operation(5, 3))

	// 7. 匿名函数
	fmt.Println("\n7. 匿名函数：")
	square := func(x int) int {
		return x * x
	}
	fmt.Printf("匿名函数 square(4): %d\n", square(4))

	// 立即执行的匿名函数
	result := func(a, b int) int {
		return a*a + b*b
	}(3, 4)
	fmt.Printf("立即执行匿名函数 (3² + 4²): %d\n", result)

	// 8. 递归函数
	fmt.Println("\n8. 递归函数：")
	fmt.Printf("阶乘 5! = %d\n", factorial(5))
	fmt.Printf("斐波那契数列第10项: %d\n", fibonacci(10))

	// 9. defer 语句
	fmt.Println("\n9. defer 语句演示：")
	demoDefer()

	// 10. 错误处理
	fmt.Println("\n10. 错误处理：")
	result1, err := safeDivide(10, 2)
	if err != nil {
		fmt.Printf("错误: %v\n", err)
	} else {
		fmt.Printf("10 ÷ 2 = %.2f\n", result1)
	}

	result2, err := safeDivide(10, 0)
	if err != nil {
		fmt.Printf("错误: %v\n", err)
	} else {
		fmt.Printf("10 ÷ 0 = %.2f\n", result2)
	}
}

// greet 简单的问候函数
func greet(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

// add 两数相加
func add(a, b int) int {
	return a + b
}

// multiply 两数相乘
func multiply(a, b int) int {
	return a * b
}

// divide 除法运算，返回商和余数
func divide(dividend, divisor int) (int, int) {
	quotient := dividend / divisor
	remainder := dividend % divisor
	return quotient, remainder
}

// rectangleStats 计算矩形的面积和周长（命名返回值）
func rectangleStats(length, width int) (area, perimeter int) {
	area = length * width
	perimeter = 2 * (length + width)
	return // 裸返回
}

// sumAll 可变参数函数，计算所有参数的和
func sumAll(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// factorial 计算阶乘（递归）
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

// fibonacci 计算斐波那契数列第n项（递归）
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// demoDefer 演示defer语句
func demoDefer() {
	fmt.Println("  函数开始")

	defer fmt.Println("  defer 1: 最后执行")
	defer fmt.Println("  defer 2: 倒数第二执行")
	defer fmt.Println("  defer 3: 倒数第三执行")

	fmt.Println("  函数中间")

	// defer 语句按LIFO（后进先出）顺序执行
	for i := 1; i <= 3; i++ {
		defer fmt.Printf("  循环defer %d\n", i)
	}

	fmt.Println("  函数即将结束")
}

// safeDivide 安全除法，返回结果和错误
func safeDivide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("除数不能为零")
	}
	return a / b, nil
}

// DemoHigherOrderFunctions 演示高阶函数
func DemoHigherOrderFunctions() {
	fmt.Println("\n=== 高阶函数演示 ===")

	// 1. 函数作为参数
	fmt.Println("\n1. 函数作为参数：")
	numbers := []int{1, 2, 3, 4, 5}

	// 使用不同的函数处理数组
	doubled := mapInts(numbers, func(x int) int { return x * 2 })
	squared := mapInts(numbers, func(x int) int { return x * x })

	fmt.Printf("原数组: %v\n", numbers)
	fmt.Printf("翻倍: %v\n", doubled)
	fmt.Printf("平方: %v\n", squared)

	// 2. 过滤函数
	fmt.Println("\n2. 过滤函数：")
	evens := filterInts(numbers, func(x int) bool { return x%2 == 0 })
	odds := filterInts(numbers, func(x int) bool { return x%2 == 1 })
	greaterThan3 := filterInts(numbers, func(x int) bool { return x > 3 })

	fmt.Printf("偶数: %v\n", evens)
	fmt.Printf("奇数: %v\n", odds)
	fmt.Printf("大于3: %v\n", greaterThan3)

	// 3. 归约函数
	fmt.Println("\n3. 归约函数：")
	sum := reduceInts(numbers, 0, func(acc, x int) int { return acc + x })
	product := reduceInts(numbers, 1, func(acc, x int) int { return acc * x })
	max := reduceInts(numbers, numbers[0], func(acc, x int) int {
		if x > acc {
			return x
		}
		return acc
	})

	fmt.Printf("求和: %d\n", sum)
	fmt.Printf("求积: %d\n", product)
	fmt.Printf("最大值: %d\n", max)

	// 4. 函数组合
	fmt.Println("\n4. 函数组合：")
	addOne := func(x int) int { return x + 1 }
	multiplyByTwo := func(x int) int { return x * 2 }

	// 组合函数：先加1，再乘2
	composed := compose(multiplyByTwo, addOne)
	result := composed(5) // (5 + 1) * 2 = 12
	fmt.Printf("组合函数 (5 + 1) * 2 = %d\n", result)

	// 5. 柯里化
	fmt.Println("\n5. 柯里化：")
	addCurried := curry(func(a, b int) int { return a + b })
	add10 := addCurried(10)

	fmt.Printf("柯里化加法 add10(5): %d\n", add10(5))
	fmt.Printf("柯里化加法 add10(15): %d\n", add10(15))

	// 6. 函数工厂
	fmt.Println("\n6. 函数工厂：")
	multiplier3 := createMultiplier(3)
	multiplier5 := createMultiplier(5)

	fmt.Printf("3倍数生成器: %d\n", multiplier3(4))
	fmt.Printf("5倍数生成器: %d\n", multiplier5(4))
}

// mapInts 对整数切片应用函数
func mapInts(slice []int, fn func(int) int) []int {
	result := make([]int, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}

// filterInts 过滤整数切片
func filterInts(slice []int, predicate func(int) bool) []int {
	var result []int
	for _, v := range slice {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

// reduceInts 归约整数切片
func reduceInts(slice []int, initial int, fn func(int, int) int) int {
	result := initial
	for _, v := range slice {
		result = fn(result, v)
	}
	return result
}

// compose 组合两个函数
func compose(f, g func(int) int) func(int) int {
	return func(x int) int {
		return f(g(x))
	}
}

// curry 柯里化二元函数
func curry(fn func(int, int) int) func(int) func(int) int {
	return func(a int) func(int) int {
		return func(b int) int {
			return fn(a, b)
		}
	}
}

// createMultiplier 创建乘法器函数
func createMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

// DemoClosure 演示闭包
func DemoClosure() {
	fmt.Println("\n=== 闭包演示 ===")

	// 1. 基本闭包
	fmt.Println("\n1. 基本闭包：")
	counter := createCounter()
	fmt.Printf("计数器: %d\n", counter())
	fmt.Printf("计数器: %d\n", counter())
	fmt.Printf("计数器: %d\n", counter())

	// 创建另一个独立的计数器
	counter2 := createCounter()
	fmt.Printf("计数器2: %d\n", counter2())
	fmt.Printf("原计数器: %d\n", counter())

	// 2. 带参数的闭包
	fmt.Println("\n2. 带参数的闭包：")
	adder := createAdder(10)
	fmt.Printf("加法器(+10): %d\n", adder(5))
	fmt.Printf("加法器(+10): %d\n", adder(3))

	// 3. 修改外部变量的闭包
	fmt.Println("\n3. 修改外部变量的闭包：")
	balance := 100.0
	withdraw := createWithdrawFunction(&balance)

	fmt.Printf("初始余额: %.2f\n", balance)
	success := withdraw(30)
	fmt.Printf("取款30: %t, 余额: %.2f\n", success, balance)
	success = withdraw(80)
	fmt.Printf("取款80: %t, 余额: %.2f\n", success, balance)

	// 4. 闭包捕获循环变量
	fmt.Println("\n4. 闭包捕获循环变量：")

	// 错误的方式（所有闭包都会捕获最后的i值）
	fmt.Println("错误的方式:")
	var funcs1 []func() int
	for i := 0; i < 3; i++ {
		funcs1 = append(funcs1, func() int {
			return i // 捕获的是循环变量i的引用
		})
	}
	for j, f := range funcs1 {
		fmt.Printf("  函数%d: %d\n", j, f())
	}

	// 正确的方式1：使用参数传递
	fmt.Println("正确的方式1（参数传递）:")
	var funcs2 []func() int
	for i := 0; i < 3; i++ {
		funcs2 = append(funcs2, func(val int) func() int {
			return func() int {
				return val
			}
		}(i))
	}
	for j, f := range funcs2 {
		fmt.Printf("  函数%d: %d\n", j, f())
	}

	// 正确的方式2：使用局部变量
	fmt.Println("正确的方式2（局部变量）:")
	var funcs3 []func() int
	for i := 0; i < 3; i++ {
		val := i // 创建局部变量
		funcs3 = append(funcs3, func() int {
			return val
		})
	}
	for j, f := range funcs3 {
		fmt.Printf("  函数%d: %d\n", j, f())
	}

	// 5. 闭包实现装饰器模式
	fmt.Println("\n5. 闭包实现装饰器模式：")

	// 原始函数
	slowFunction := func(name string) string {
		return fmt.Sprintf("处理 %s", name)
	}

	// 添加日志装饰器
	loggedFunction := withLogging(slowFunction)

	// 添加计时装饰器
	timedFunction := withTiming(loggedFunction)

	result := timedFunction("重要任务")
	fmt.Printf("最终结果: %s\n", result)

	// 6. 闭包实现缓存
	fmt.Println("\n6. 闭包实现缓存：")

	// 创建带缓存的斐波那契函数
	fibWithCache := createCachedFibonacci()

	fmt.Printf("斐波那契(10): %d\n", fibWithCache(10))
	fmt.Printf("斐波那契(15): %d\n", fibWithCache(15))
	fmt.Printf("斐波那契(10): %d (从缓存获取)\n", fibWithCache(10))
}

// createCounter 创建计数器闭包
func createCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// createAdder 创建加法器闭包
func createAdder(base int) func(int) int {
	return func(x int) int {
		return base + x
	}
}

// createWithdrawFunction 创建取款函数闭包
func createWithdrawFunction(balance *float64) func(float64) bool {
	return func(amount float64) bool {
		if *balance >= amount {
			*balance -= amount
			return true
		}
		return false
	}
}

// withLogging 日志装饰器
func withLogging(fn func(string) string) func(string) string {
	return func(input string) string {
		fmt.Printf("  [LOG] 开始处理: %s\n", input)
		result := fn(input)
		fmt.Printf("  [LOG] 处理完成: %s\n", input)
		return result
	}
}

// withTiming 计时装饰器
func withTiming(fn func(string) string) func(string) string {
	return func(input string) string {
		fmt.Printf("  [TIMER] 开始计时\n")
		result := fn(input)
		fmt.Printf("  [TIMER] 执行完成\n")
		return result
	}
}

// createCachedFibonacci 创建带缓存的斐波那契函数
func createCachedFibonacci() func(int) int {
	cache := make(map[int]int)

	var fib func(int) int
	fib = func(n int) int {
		if n <= 1 {
			return n
		}

		// 检查缓存
		if val, exists := cache[n]; exists {
			fmt.Printf("    从缓存获取 fib(%d) = %d\n", n, val)
			return val
		}

		// 计算并缓存
		result := fib(n-1) + fib(n-2)
		cache[n] = result
		fmt.Printf("    计算并缓存 fib(%d) = %d\n", n, result)
		return result
	}

	return fib
}
