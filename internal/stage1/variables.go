package stage1

import "fmt"

// DemoVariablesAndConstants 演示变量和常量的使用
func DemoVariablesAndConstants() {
	fmt.Println("\n=== 变量和常量演示 ===")

	// 1. 变量声明的几种方式
	fmt.Println("\n1. 变量声明方式：")

	// 方式1：var 关键字声明
	var name string
	name = "Go语言"
	fmt.Printf("方式1 - var声明: %s\n", name)

	// 方式2：var 声明并初始化
	var age int = 25
	fmt.Printf("方式2 - var声明并初始化: %d\n", age)

	// 方式3：类型推断
	var score = 95.5
	fmt.Printf("方式3 - 类型推断: %.1f (类型: %T)\n", score, score)

	// 方式4：短变量声明（最常用）
	city := "北京"
	fmt.Printf("方式4 - 短变量声明: %s\n", city)

	// 2. 多变量声明
	fmt.Println("\n2. 多变量声明：")
	var x, y int = 10, 20
	fmt.Printf("多变量声明: x=%d, y=%d\n", x, y)

	a, b, c := 1, 2.5, "hello"
	fmt.Printf("多变量短声明: a=%d, b=%.1f, c=%s\n", a, b, c)

	// 3. 零值
	fmt.Println("\n3. 零值演示：")
	var (
		defaultInt    int
		defaultFloat  float64
		defaultBool   bool
		defaultString string
	)
	fmt.Printf("int零值: %d\n", defaultInt)
	fmt.Printf("float64零值: %.1f\n", defaultFloat)
	fmt.Printf("bool零值: %t\n", defaultBool)
	fmt.Printf("string零值: '%s' (长度: %d)\n", defaultString, len(defaultString))

	// 4. 常量
	fmt.Println("\n4. 常量演示：")
	const pi = 3.14159
	const greeting = "Hello, World!"
	fmt.Printf("常量pi: %.5f\n", pi)
	fmt.Printf("常量greeting: %s\n", greeting)

	// 5. 常量组
	fmt.Println("\n5. 常量组：")
	const (
		Monday    = 1
		Tuesday   = 2
		Wednesday = 3
	)
	fmt.Printf("星期一: %d, 星期二: %d, 星期三: %d\n", Monday, Tuesday, Wednesday)

	// 6. iota 枚举器
	fmt.Println("\n6. iota 枚举器：")
	const (
		Red   = iota // 0
		Green        // 1
		Blue         // 2
	)
	fmt.Printf("Red: %d, Green: %d, Blue: %d\n", Red, Green, Blue)

	const (
		_  = iota             // 跳过0
		KB = 1 << (10 * iota) // 1024
		MB                    // 1048576
		GB                    // 1073741824
	)
	fmt.Printf("KB: %d, MB: %d, GB: %d\n", KB, MB, GB)
}
