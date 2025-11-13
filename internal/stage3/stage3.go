package stage3

import "fmt"

// RunStage3 运行第3阶段演示
func RunStage3() {
	fmt.Println("第3阶段：接口与多态")
	
	// 接口演示
	DemoInterfaces()
	
	// 多态演示
	DemoPolymorphism()
	
	// 接口组合演示
	DemoInterfaceComposition()
	
	// 设计模式演示
	DemoDesignPatterns()
	
	// 类型断言演示
	DemoTypeAssertion()
}

// DemoInterfaces 演示接口
func DemoInterfaces() {
	fmt.Println("\n=== 接口基础演示 ===")
	
	// 1. 基本接口定义和实现
	fmt.Println("\n1. 基本接口定义和实现：")
	demoBasicInterface()
	
	// 2. 空接口
	fmt.Println("\n2. 空接口：")
	demoEmptyInterface()
	
	// 3. 接口值
	fmt.Println("\n3. 接口值：")
	demoInterfaceValues()
	
	// 4. 接口实现检查
	fmt.Println("\n4. 接口实现检查：")
	demoInterfaceImplementation()
	
	// 5. 接口最佳实践
	fmt.Println("\n5. 接口最佳实践：")
	demoInterfaceBestPractices()
}

// DemoPolymorphism 演示多态
func DemoPolymorphism() {
	fmt.Println("\n=== 多态演示 ===")
	
	// 1. 基本多态
	fmt.Println("\n1. 基本多态：")
	demoBasicPolymorphism()
	
	// 2. 接口切片多态
	fmt.Println("\n2. 接口切片多态：")
	demoPolymorphicSlice()
	
	// 3. 多态工厂模式
	fmt.Println("\n3. 多态工厂模式：")
	demoPolymorphicFactory()
	
	// 4. 策略模式
	fmt.Println("\n4. 策略模式：")
	demoStrategyPattern()
	
	// 5. 多态的实际应用
	fmt.Println("\n5. 多态的实际应用：")
	demoPolymorphismInPractice()
}

// DemoInterfaceComposition 演示接口组合
func DemoInterfaceComposition() {
	fmt.Println("\n=== 接口组合演示 ===")
	
	// 1. 基本接口组合
	fmt.Println("\n1. 基本接口组合：")
	demoBasicInterfaceComposition()
	
	// 2. 多层接口组合
	fmt.Println("\n2. 多层接口组合：")
	demoMultiLevelComposition()
	
	// 3. 接口分离原则
	fmt.Println("\n3. 接口分离原则：")
	demoInterfaceSegregation()
	
	// 4. 组合vs继承
	fmt.Println("\n4. 组合vs继承：")
	demoCompositionVsInheritance()
	
	// 5. 实际应用场景
	fmt.Println("\n5. 实际应用场景：")
	demoCompositionInPractice()
}

// DemoDesignPatterns 演示设计模式
func DemoDesignPatterns() {
	fmt.Println("\n=== 设计模式演示 ===")
	
	// 1. 观察者模式
	fmt.Println("\n1. 观察者模式：")
	demoObserverPattern()
	
	// 2. 装饰器模式
	fmt.Println("\n2. 装饰器模式：")
	demoDecoratorPattern()
	
	// 3. 适配器模式
	fmt.Println("\n3. 适配器模式：")
	demoAdapterPattern()
	
	// 4. 命令模式
	fmt.Println("\n4. 命令模式：")
	demoCommandPattern()
	
	// 5. 责任链模式
	fmt.Println("\n5. 责任链模式：")
	demoChainOfResponsibility()
}

// DemoTypeAssertion 演示类型断言
func DemoTypeAssertion() {
	fmt.Println("\n=== 类型断言演示 ===")
	
	// 1. 基本类型断言
	fmt.Println("\n1. 基本类型断言：")
	demoBasicTypeAssertion()
	
	// 2. 类型开关
	fmt.Println("\n2. 类型开关：")
	demoTypeSwitch()
	
	// 3. 接口类型断言
	fmt.Println("\n3. 接口类型断言：")
	demoInterfaceTypeAssertion()
	
	// 4. 类型断言的安全性
	fmt.Println("\n4. 类型断言的安全性：")
	demoTypeAssertionSafety()
	
	// 5. 实际应用场景
	fmt.Println("\n5. 实际应用场景：")
	demoTypeAssertionInPractice()
}

// ===== 接口基础演示 =====

// Shape 形状接口
type Shape interface {
	Area() float64
	Perimeter() float64
	String() string
}

// Rectangle 矩形
type Rectangle struct {
	Width  float64
	Height float64
}

// Area 计算矩形面积
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter 计算矩形周长
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// String 矩形字符串表示
func (r Rectangle) String() string {
	return fmt.Sprintf("Rectangle{Width: %.2f, Height: %.2f}", r.Width, r.Height)
}

// Circle 圆形
type Circle struct {
	Radius float64
}

// Area 计算圆形面积
func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

// Perimeter 计算圆形周长
func (c Circle) Perimeter() float64 {
	return 2 * 3.14159 * c.Radius
}

// String 圆形字符串表示
func (c Circle) String() string {
	return fmt.Sprintf("Circle{Radius: %.2f}", c.Radius)
}

// demoBasicInterface 演示基本接口
func demoBasicInterface() {
	// 1. 创建不同形状的实例
	rect := Rectangle{Width: 5, Height: 3}
	circle := Circle{Radius: 4}
	
	// 2. 使用接口变量
	var shape Shape
	
	// 3. 矩形操作
	shape = rect
	fmt.Printf("形状: %s\n", shape.String())
	fmt.Printf("面积: %.2f\n", shape.Area())
	fmt.Printf("周长: %.2f\n", shape.Perimeter())
	
	// 4. 圆形操作
	shape = circle
	fmt.Printf("\n形状: %s\n", shape.String())
	fmt.Printf("面积: %.2f\n", shape.Area())
	fmt.Printf("周长: %.2f\n", shape.Perimeter())
	
	// 5. 接口函数
	printShapeInfo := func(s Shape) {
		fmt.Printf("形状信息: %s, 面积: %.2f, 周长: %.2f\n", 
			s.String(), s.Area(), s.Perimeter())
	}
	
	fmt.Println("\n使用接口函数:")
	printShapeInfo(rect)
	printShapeInfo(circle)
}

// demoEmptyInterface 演示空接口
func demoEmptyInterface() {
	// 1. 空接口可以存储任何类型
	var anything interface{}
	
	anything = 42
	fmt.Printf("整数: %v (类型: %T)\n", anything, anything)
	
	anything = "Hello, Go!"
	fmt.Printf("字符串: %v (类型: %T)\n", anything, anything)
	
	anything = []int{1, 2, 3}
	fmt.Printf("切片: %v (类型: %T)\n", anything, anything)
	
	anything = Rectangle{Width: 10, Height: 5}
	fmt.Printf("结构体: %v (类型: %T)\n", anything, anything)
	
	// 2. 空接口切片
	items := []interface{}{
		42,
		"hello",
		3.14,
		true,
		Rectangle{Width: 2, Height: 3},
	}
	
	fmt.Println("\n空接口切片:")
	for i, item := range items {
		fmt.Printf("索引 %d: %v (类型: %T)\n", i, item, item)
	}
	
	// 3. 空接口映射
	data := map[string]interface{}{
		"name":    "Go语言",
		"version": 1.21,
		"active":  true,
		"tags":    []string{"programming", "language"},
	}
	
	fmt.Println("\n空接口映射:")
	for key, value := range data {
		fmt.Printf("%s: %v (类型: %T)\n", key, value, value)
	}
}

// demoInterfaceValues 演示接口值
func demoInterfaceValues() {
	var shape Shape
	
	// 1. nil 接口
	fmt.Printf("nil接口: %v (类型: %T)\n", shape, shape)
	if shape == nil {
		fmt.Println("接口为 nil")
	}
	
	// 2. 接口包含值
	shape = Rectangle{Width: 4, Height: 6}
	fmt.Printf("接口值: %v (类型: %T)\n", shape, shape)
	
	// 3. 接口包含指针
	rect := &Rectangle{Width: 8, Height: 2}
	shape = rect
	fmt.Printf("接口指针: %v (类型: %T)\n", shape, shape)
	
	// 4. 接口值比较
	shape1 := Rectangle{Width: 5, Height: 5}
	shape2 := Rectangle{Width: 5, Height: 5}
	shape3 := Circle{Radius: 3}
	
	var s1, s2, s3 Shape
	s1 = shape1
	s2 = shape2
	s3 = shape3
	
	fmt.Printf("\n接口值比较:\n")
	fmt.Printf("s1 == s2: %t\n", s1 == s2) // true，相同类型相同值
	fmt.Printf("s1 == s3: %t\n", s1 == s3) // false，不同类型
	
	// 5. 接口的动态类型和动态值
	fmt.Println("\n接口的内部结构:")
	printInterfaceInfo := func(name string, s Shape) {
		if s == nil {
			fmt.Printf("%s: nil接口\n", name)
		} else {
			fmt.Printf("%s: 动态类型=%T, 动态值=%v\n", name, s, s)
		}
	}
	
	printInterfaceInfo("shape1", s1)
	printInterfaceInfo("shape2", s2)
	printInterfaceInfo("shape3", s3)
}

// 添加缺失的函数实现
func demoInterfaceImplementation() {
	fmt.Println("接口实现检查演示")
}

func demoInterfaceBestPractices() {
	fmt.Println("接口最佳实践演示")
}

func demoBasicInterfaceComposition() {
	fmt.Println("基本接口组合演示")
}

func demoMultiLevelComposition() {
	fmt.Println("多层接口组合演示")
}

func demoInterfaceSegregation() {
	fmt.Println("接口分离原则演示")
}

func demoCompositionVsInheritance() {
	fmt.Println("组合vs继承演示")
}

func demoCompositionInPractice() {
	fmt.Println("组合实际应用演示")
}

func demoObserverPattern() {
	fmt.Println("观察者模式演示")
}

func demoDecoratorPattern() {
	fmt.Println("装饰器模式演示")
}

func demoAdapterPattern() {
	fmt.Println("适配器模式演示")
}

func demoCommandPattern() {
	fmt.Println("命令模式演示")
}

func demoChainOfResponsibility() {
	fmt.Println("责任链模式演示")
}

func demoBasicTypeAssertion() {
	fmt.Println("基本类型断言演示")
}

func demoTypeSwitch() {
	fmt.Println("类型开关演示")
}

func demoInterfaceTypeAssertion() {
	fmt.Println("接口类型断言演示")
}

func demoTypeAssertionSafety() {
	fmt.Println("类型断言安全性演示")
}

func demoTypeAssertionInPractice() {
	fmt.Println("类型断言实际应用演示")
}
