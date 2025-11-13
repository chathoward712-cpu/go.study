package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/howard/go.study/internal/stage1"
	"github.com/howard/go.study/internal/stage2"
	"github.com/howard/go.study/internal/stage3"
	"github.com/howard/go.study/internal/stage4"
	"github.com/howard/go.study/internal/stage5"
)

func main() {
	fmt.Println("╔════════════════════════════════════════╗")
	fmt.Println("║   Go 语言完整学习课程 - go.study       ║")
	fmt.Println("║   From Zero to Go Web Developer        ║")
	fmt.Println("╚════════════════════════════════════════╝")
	fmt.Println()

	for {
		displayMenu()
		choice := getUserInput("请选择要学习的阶段 (1-5) 或 0 退出: ")

		switch choice {
		case 1:
			runStage1()
		case 2:
			runStage2()
		case 3:
			runStage3()
		case 4:
			runStage4()
		case 5:
			runStage5()
		case 0:
			fmt.Println("谢谢使用 go.study，再见！")
			return
		default:
			fmt.Println("无效选择，请重试")
		}
		fmt.Println()
	}
}

func displayMenu() {
	fmt.Println("\n=== 学习阶段选择 ===")
	fmt.Println("1. 第1阶段：Go 基础语法与类型系统")
	fmt.Println("2. 第2阶段：集合数据结构与方法系统")
	fmt.Println("3. 第3阶段：接口与设计模式")
	fmt.Println("4. 第4阶段：并发编程")
	fmt.Println("5. 第5阶段：模块与工程化")
	fmt.Println("0. 退出")
}

func runStage1() {
	fmt.Println("\n╔════════════════════════════════════════╗")
	fmt.Println("║ 第1阶段：Go 基础语法与类型系统         ║")
	fmt.Println("╚════════════════════════════════════════╝")

	for {
		fmt.Println("\n--- 子模块选择 ---")
		fmt.Println("1. 变量和常量")
		fmt.Println("2. 类型系统（基本类型）")
		fmt.Println("3. 函数定义与调用")
		fmt.Println("4. 高阶函数与闭包")
		fmt.Println("5. 控制流语句")
		fmt.Println("6. 指针")
		fmt.Println("0. 返回上级菜单")

		choice := getUserInput("选择子模块: ")
		switch choice {
		case 1:
			stage1.DemoVariablesAndConstants()
		case 2:
			stage1.DemoNumericTypes()
			stage1.DemoStringTypes()
			stage1.DemoBoolType()
		case 3:
			stage1.DemoFunctions()
		case 4:
			stage1.DemoHigherOrderFunctions()
			stage1.DemoClosure()
		case 5:
			stage1.DemoControlFlow()
		case 6:
			stage1.DemoPointers()
			stage1.DemoPointersAdvanced()
		case 0:
			return
		default:
			fmt.Println("无效选择")
		}
	}
}

func runStage2() {
	fmt.Println("\n╔════════════════════════════════════════╗")
	fmt.Println("║ 第2阶段：集合数据结构与方法系统       ║")
	fmt.Println("╚════════════════════════════════════════╝")

	for {
		fmt.Println("\n--- 子模块选择 ---")
		fmt.Println("1. 数组与切片")
		fmt.Println("2. 映射 (Map)")
		fmt.Println("3. 字符串操作")
		fmt.Println("4. 结构体")
		fmt.Println("5. 方法系统")
		fmt.Println("6. 构造函数与嵌入")
		fmt.Println("0. 返回上级菜单")

		choice := getUserInput("选择子模块: ")
		switch choice {
		case 1:
			stage2.DemoArrays()
			stage2.DemoSlices()
		case 2:
			stage2.DemoMaps()
		case 3:
			stage2.DemoStringOperations()
		case 4:
			stage2.DemoStructs()
		case 5:
			stage2.DemoMethods()
		case 6:
			stage2.DemoConstructor()
			stage2.DemoEmbedding()
		case 0:
			return
		default:
			fmt.Println("无效选择")
		}
	}
}

func runStage3() {
	fmt.Println("\n╔════════════════════════════════════════╗")
	fmt.Println("║ 第3阶段：接口与设计模式               ║")
	fmt.Println("╚════════════════════════════════════════╝")

	for {
		fmt.Println("\n--- 子模块选择 ---")
		fmt.Println("1. 接口基础")
		fmt.Println("2. 类型断言与类型转换")
		fmt.Println("3. 工厂模式")
		fmt.Println("4. 策略模式")
		fmt.Println("5. 观察者模式")
		fmt.Println("6. 装饰器模式")
		fmt.Println("0. 返回上级菜单")

		choice := getUserInput("选择子模块: ")
		switch choice {
		case 1:
			stage3.DemoInterfaces()
			stage3.DemoPrinterInterface()
		case 2:
			stage3.DemoTypeAssertion()
			stage3.DemoEmptyInterface()
		case 3:
			stage3.DemoFactoryPattern()
		case 4:
			stage3.DemoStrategyPattern()
		case 5:
			stage3.DemoObserverPattern()
		case 6:
			stage3.DemoDecoratorPattern()
		case 0:
			return
		default:
			fmt.Println("无效选择")
		}
	}
}

func runStage4() {
	fmt.Println("\n╔════════════════════════════════════════╗")
	fmt.Println("║ 第4阶段：并发编程                     ║")
	fmt.Println("╚════════════════════════════════════════╝")

	for {
		fmt.Println("\n--- 子模块选择 ---")
		fmt.Println("1. Goroutine 基础")
		fmt.Println("2. Channel 基础")
		fmt.Println("3. Select 多路复用")
		fmt.Println("4. 互斥锁 (Mutex)")
		fmt.Println("5. sync.Once")
		fmt.Println("6. Channel 模式与 Pipeline")
		fmt.Println("0. 返回上级菜单")

		choice := getUserInput("选择子模块: ")
		switch choice {
		case 1:
			stage4.DemoGoroutines()
		case 2:
			stage4.DemoChannels()
		case 3:
			stage4.DemoSelect()
		case 4:
			stage4.DemoMutex()
		case 5:
			stage4.DemoOnce()
		case 6:
			stage4.DemoChannelPatterns()
		case 0:
			return
		default:
			fmt.Println("无效选择")
		}
	}
}

func runStage5() {
	fmt.Println("\n╔════════════════════════════════════════╗")
	fmt.Println("║ 第5阶段：模块与工程化                 ║")
	fmt.Println("╚════════════════════════════════════════╝")

	for {
		fmt.Println("\n--- 子模块选择 ---")
		fmt.Println("1. 错误处理")
		fmt.Println("2. 字符串转换")
		fmt.Println("3. 日志系统")
		fmt.Println("4. 模块组织")
		fmt.Println("5. Go 模块管理 (go.mod)")
		fmt.Println("0. 返回上级菜单")

		choice := getUserInput("选择子模块: ")
		switch choice {
		case 1:
			stage5.DemoErrors()
		case 2:
			stage5.DemoStringConversion()
		case 3:
			stage5.DemoLogging()
		case 4:
			stage5.DemoModuleOrganization()
		case 5:
			stage5.DemoGoMod()
		case 0:
			return
		default:
			fmt.Println("无效选择")
		}
	}
}

func getUserInput(prompt string) int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	choice, _ := strconv.Atoi(strings.TrimSpace(input))
	return choice
}
