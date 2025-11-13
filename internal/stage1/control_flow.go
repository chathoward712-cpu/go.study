package stage1

import (
	"fmt"
	"math/rand"
	"time"
)

// DemoControlFlow 演示控制流语句
func DemoControlFlow() {
	fmt.Println("\n=== 控制流语句演示 ===")

	// 1. if-else 语句
	fmt.Println("\n1. if-else 语句：")
	demoIfElse()

	// 2. for 循环
	fmt.Println("\n2. for 循环：")
	demoForLoops()

	// 3. switch 语句
	fmt.Println("\n3. switch 语句：")
	demoSwitch()

	// 4. 循环控制语句
	fmt.Println("\n4. 循环控制语句：")
	demoLoopControl()

	// 5. 标签和跳转
	fmt.Println("\n5. 标签和跳转：")
	demoLabelsAndJumps()
}

// demoIfElse 演示 if-else 语句
func demoIfElse() {
	// 基本 if 语句
	age := 18
	if age >= 18 {
		fmt.Printf("年龄 %d：成年人\n", age)
	}

	// if-else 语句
	score := 85
	if score >= 90 {
		fmt.Printf("分数 %d：优秀\n", score)
	} else if score >= 80 {
		fmt.Printf("分数 %d：良好\n", score)
	} else if score >= 70 {
		fmt.Printf("分数 %d：中等\n", score)
	} else if score >= 60 {
		fmt.Printf("分数 %d：及格\n", score)
	} else {
		fmt.Printf("分数 %d：不及格\n", score)
	}

	// if 语句的初始化
	if num := rand.Intn(100); num > 50 {
		fmt.Printf("随机数 %d 大于 50\n", num)
	} else {
		fmt.Printf("随机数 %d 小于等于 50\n", num)
	}
	// 注意：num 变量只在 if 语句块中有效

	// 复杂条件判断
	temperature := 25
	humidity := 60

	if temperature > 30 && humidity > 70 {
		fmt.Println("天气：炎热潮湿")
	} else if temperature > 30 && humidity <= 70 {
		fmt.Println("天气：炎热干燥")
	} else if temperature <= 30 && humidity > 70 {
		fmt.Println("天气：温和潮湿")
	} else {
		fmt.Println("天气：温和干燥")
	}

	// 检查零值
	var name string
	var count int
	var ptr *int

	if name == "" {
		fmt.Println("字符串为空")
	}

	if count == 0 {
		fmt.Println("计数为零")
	}

	if ptr == nil {
		fmt.Println("指针为nil")
	}
}

// demoForLoops 演示 for 循环的各种形式
func demoForLoops() {
	// 1. 传统的三部分 for 循环
	fmt.Println("传统 for 循环:")
	for i := 0; i < 5; i++ {
		fmt.Printf("  i = %d\n", i)
	}

	// 2. while 风格的 for 循环
	fmt.Println("while 风格:")
	j := 0
	for j < 3 {
		fmt.Printf("  j = %d\n", j)
		j++
	}

	// 3. 无限循环（需要用 break 退出）
	fmt.Println("无限循环（计数到3退出）:")
	k := 0
	for {
		if k >= 3 {
			break
		}
		fmt.Printf("  k = %d\n", k)
		k++
	}

	// 4. range 循环 - 遍历切片
	fmt.Println("遍历切片:")
	fruits := []string{"苹果", "香蕉", "橙子"}
	for index, fruit := range fruits {
		fmt.Printf("  索引 %d: %s\n", index, fruit)
	}

	// 只要值，不要索引
	fmt.Println("只要值:")
	for _, fruit := range fruits {
		fmt.Printf("  水果: %s\n", fruit)
	}

	// 只要索引，不要值
	fmt.Println("只要索引:")
	for index := range fruits {
		fmt.Printf("  索引: %d\n", index)
	}

	// 5. range 循环 - 遍历映射
	fmt.Println("遍历映射:")
	ages := map[string]int{
		"Alice": 25,
		"Bob":   30,
		"Carol": 35,
	}
	for name, age := range ages {
		fmt.Printf("  %s: %d岁\n", name, age)
	}

	// 6. range 循环 - 遍历字符串
	fmt.Println("遍历字符串:")
	text := "Go语言"
	for i, r := range text {
		fmt.Printf("  位置 %d: %c (Unicode: %d)\n", i, r, r)
	}

	// 7. range 循环 - 遍历通道（channel）
	fmt.Println("遍历通道:")
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch) // 关闭通道，否则 range 会阻塞

	for value := range ch {
		fmt.Printf("  从通道接收: %d\n", value)
	}

	// 8. 嵌套循环
	fmt.Println("嵌套循环（乘法表）:")
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("  %d × %d = %d\n", i, j, i*j)
		}
	}
}

// demoSwitch 演示 switch 语句
func demoSwitch() {
	// 1. 基本 switch 语句
	fmt.Println("基本 switch:")
	day := 3
	switch day {
	case 1:
		fmt.Println("  星期一")
	case 2:
		fmt.Println("  星期二")
	case 3:
		fmt.Println("  星期三")
	case 4:
		fmt.Println("  星期四")
	case 5:
		fmt.Println("  星期五")
	case 6, 7: // 多个值
		fmt.Println("  周末")
	default:
		fmt.Println("  无效的日期")
	}

	// 2. switch 语句的初始化
	fmt.Println("带初始化的 switch:")
	switch hour := time.Now().Hour(); {
	case hour < 6:
		fmt.Println("  凌晨")
	case hour < 12:
		fmt.Println("  上午")
	case hour < 18:
		fmt.Println("  下午")
	default:
		fmt.Println("  晚上")
	}

	// 3. 表达式 switch
	fmt.Println("表达式 switch:")
	score := 85
	switch {
	case score >= 90:
		fmt.Println("  等级: A")
	case score >= 80:
		fmt.Println("  等级: B")
	case score >= 70:
		fmt.Println("  等级: C")
	case score >= 60:
		fmt.Println("  等级: D")
	default:
		fmt.Println("  等级: F")
	}

	// 4. 类型 switch
	fmt.Println("类型 switch:")
	var value interface{} = "Hello"
	switch v := value.(type) {
	case string:
		fmt.Printf("  字符串: %s (长度: %d)\n", v, len(v))
	case int:
		fmt.Printf("  整数: %d\n", v)
	case bool:
		fmt.Printf("  布尔值: %t\n", v)
	default:
		fmt.Printf("  未知类型: %T\n", v)
	}

	// 5. fallthrough 关键字
	fmt.Println("fallthrough 演示:")
	grade := 'B'
	switch grade {
	case 'A':
		fmt.Println("  优秀")
		fallthrough
	case 'B':
		fmt.Println("  良好")
		fallthrough
	case 'C':
		fmt.Println("  及格")
	case 'D':
		fmt.Println("  不及格")
	}
}

// demoLoopControl 演示循环控制语句
func demoLoopControl() {
	// 1. break 语句
	fmt.Println("break 语句:")
	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Printf("  遇到 %d，跳出循环\n", i)
			break
		}
		fmt.Printf("  i = %d\n", i)
	}

	// 2. continue 语句
	fmt.Println("continue 语句:")
	for i := 0; i < 5; i++ {
		if i == 2 {
			fmt.Printf("  跳过 %d\n", i)
			continue
		}
		fmt.Printf("  i = %d\n", i)
	}

	// 3. 在嵌套循环中使用 break 和 continue
	fmt.Println("嵌套循环中的控制:")
	for i := 0; i < 3; i++ {
		fmt.Printf("外层循环 i = %d\n", i)
		for j := 0; j < 3; j++ {
			if j == 1 {
				fmt.Printf("  跳过内层 j = %d\n", j)
				continue
			}
			if i == 1 && j == 2 {
				fmt.Printf("  内层 break，j = %d\n", j)
				break
			}
			fmt.Printf("  内层循环 j = %d\n", j)
		}
	}
}

// demoLabelsAndJumps 演示标签和跳转
func demoLabelsAndJumps() {
	// 1. 标签与 break
	fmt.Println("标签与 break:")
OuterLoop:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				fmt.Printf("  在 i=%d, j=%d 处跳出外层循环\n", i, j)
				break OuterLoop
			}
			fmt.Printf("  i=%d, j=%d\n", i, j)
		}
	}

	// 2. 标签与 continue
	fmt.Println("标签与 continue:")
OuterLoop2:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j == 1 {
				fmt.Printf("  在 i=%d, j=%d 处继续外层循环\n", i, j)
				continue OuterLoop2
			}
			fmt.Printf("  i=%d, j=%d\n", i, j)
		}
	}

	// 3. goto 语句（不推荐使用，但了解一下）
	fmt.Println("goto 语句演示:")
	i := 0

Start:
	if i < 3 {
		fmt.Printf("  goto 循环: i = %d\n", i)
		i++
		goto Start
	}

	// 4. 实际应用：错误处理中的 goto
	fmt.Println("错误处理中的 goto:")
	if err := processStep1(); err != nil {
		goto Cleanup
	}

	if err := processStep2(); err != nil {
		goto Cleanup
	}

	if err := processStep3(); err != nil {
		goto Cleanup
	}

	fmt.Println("  所有步骤成功完成")
	return

Cleanup:
	fmt.Println("  执行清理操作")
}

// 辅助函数用于演示错误处理
func processStep1() error {
	fmt.Println("  执行步骤1")
	return nil
}

func processStep2() error {
	fmt.Println("  执行步骤2")
	return nil
}

func processStep3() error {
	fmt.Println("  执行步骤3")
	return nil
}
