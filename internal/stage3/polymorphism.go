package stage3

import "fmt"

// ===== 多态演示 =====

// Animal 动物接口
type Animal interface {
	Speak() string
	Move() string
}

// Dog 狗
type Dog struct {
	Name string
}

// Speak 狗说话
func (d Dog) Speak() string {
	return fmt.Sprintf("%s says: Woof!", d.Name)
}

// Move 狗移动
func (d Dog) Move() string {
	return fmt.Sprintf("%s runs on four legs", d.Name)
}

// Cat 猫
type Cat struct {
	Name string
}

// Speak 猫说话
func (c Cat) Speak() string {
	return fmt.Sprintf("%s says: Meow!", c.Name)
}

// Move 猫移动
func (c Cat) Move() string {
	return fmt.Sprintf("%s walks silently", c.Name)
}

// Bird 鸟
type Bird struct {
	Name string
}

// Speak 鸟说话
func (b Bird) Speak() string {
	return fmt.Sprintf("%s says: Tweet!", b.Name)
}

// Move 鸟移动
func (b Bird) Move() string {
	return fmt.Sprintf("%s flies in the sky", b.Name)
}

// demoBasicPolymorphism 演示基本多态
func demoBasicPolymorphism() {
	// 1. 创建不同的动物
	animals := []Animal{
		Dog{Name: "Buddy"},
		Cat{Name: "Whiskers"},
		Bird{Name: "Tweety"},
	}
	
	// 2. 多态调用
	fmt.Println("动物们的行为:")
	for i, animal := range animals {
		fmt.Printf("动物 %d:\n", i+1)
		fmt.Printf("  %s\n", animal.Speak())
		fmt.Printf("  %s\n", animal.Move())
	}
	
	// 3. 多态函数
	makeAnimalPerform := func(a Animal) {
		fmt.Printf("表演: %s, %s\n", a.Speak(), a.Move())
	}
	
	fmt.Println("\n动物表演:")
	for _, animal := range animals {
		makeAnimalPerform(animal)
	}
}

// demoPolymorphicSlice 演示接口切片多态
func demoPolymorphicSlice() {
	// 1. 混合类型的切片
	shapes := []Shape{
		Rectangle{Width: 5, Height: 3},
		Circle{Radius: 4},
		Rectangle{Width: 2, Height: 8},
		Circle{Radius: 2.5},
	}
	
	// 2. 统一处理
	fmt.Println("形状统计:")
	totalArea := 0.0
	totalPerimeter := 0.0
	
	for i, shape := range shapes {
		area := shape.Area()
		perimeter := shape.Perimeter()
		
		fmt.Printf("形状 %d: %s\n", i+1, shape.String())
		fmt.Printf("  面积: %.2f, 周长: %.2f\n", area, perimeter)
		
		totalArea += area
		totalPerimeter += perimeter
	}
	
	fmt.Printf("\n总面积: %.2f\n", totalArea)
	fmt.Printf("总周长: %.2f\n", totalPerimeter)
	
	// 3. 过滤和分类
	fmt.Println("\n大面积形状 (面积 > 20):")
	for _, shape := range shapes {
		if shape.Area() > 20 {
			fmt.Printf("  %s, 面积: %.2f\n", shape.String(), shape.Area())
		}
	}
}

// ShapeType 形状类型
type ShapeType int

const (
	RectangleType ShapeType = iota
	CircleType
	TriangleType
)

// Triangle 三角形
type Triangle struct {
	Base   float64
	Height float64
}

// Area 计算三角形面积
func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

// Perimeter 计算三角形周长（简化为等腰三角形）
func (t Triangle) Perimeter() float64 {
	side := (t.Base*t.Base + t.Height*t.Height) / (2 * t.Base)
	return t.Base + 2*side
}

// String 三角形字符串表示
func (t Triangle) String() string {
	return fmt.Sprintf("Triangle{Base: %.2f, Height: %.2f}", t.Base, t.Height)
}

// ShapeFactory 形状工厂
type ShapeFactory struct{}

// CreateShape 创建形状
func (sf ShapeFactory) CreateShape(shapeType ShapeType, params ...float64) Shape {
	switch shapeType {
	case RectangleType:
		if len(params) >= 2 {
			return Rectangle{Width: params[0], Height: params[1]}
		}
		return Rectangle{Width: 1, Height: 1}
	case CircleType:
		if len(params) >= 1 {
			return Circle{Radius: params[0]}
		}
		return Circle{Radius: 1}
	case TriangleType:
		if len(params) >= 2 {
			return Triangle{Base: params[0], Height: params[1]}
		}
		return Triangle{Base: 1, Height: 1}
	default:
		return Rectangle{Width: 1, Height: 1}
	}
}

// demoPolymorphicFactory 演示多态工厂模式
func demoPolymorphicFactory() {
	factory := ShapeFactory{}
	
	// 1. 使用工厂创建不同形状
	shapes := []Shape{
		factory.CreateShape(RectangleType, 4, 6),
		factory.CreateShape(CircleType, 3),
		factory.CreateShape(TriangleType, 5, 4),
		factory.CreateShape(RectangleType, 2, 2),
	}
	
	fmt.Println("工厂创建的形状:")
	for i, shape := range shapes {
		fmt.Printf("形状 %d: %s\n", i+1, shape.String())
		fmt.Printf("  面积: %.2f, 周长: %.2f\n", shape.Area(), shape.Perimeter())
	}
	
	// 2. 批量创建
	shapeConfigs := []struct {
		shapeType ShapeType
		params    []float64
	}{
		{RectangleType, []float64{3, 5}},
		{CircleType, []float64{2.5}},
		{TriangleType, []float64{6, 3}},
	}
	
	fmt.Println("\n批量创建形状:")
	for i, config := range shapeConfigs {
		shape := factory.CreateShape(config.shapeType, config.params...)
		fmt.Printf("批量形状 %d: %s, 面积: %.2f\n", 
			i+1, shape.String(), shape.Area())
	}
}

// SortStrategy 排序策略接口
type SortStrategy interface {
	Sort(shapes []Shape) []Shape
	Name() string
}

// AreaSortStrategy 按面积排序
type AreaSortStrategy struct{}

// Sort 按面积排序
func (ass AreaSortStrategy) Sort(shapes []Shape) []Shape {
	// 简单的冒泡排序
	result := make([]Shape, len(shapes))
	copy(result, shapes)
	
	for i := 0; i < len(result); i++ {
		for j := 0; j < len(result)-1-i; j++ {
			if result[j].Area() > result[j+1].Area() {
				result[j], result[j+1] = result[j+1], result[j]
			}
		}
	}
	
	return result
}

// Name 策略名称
func (ass AreaSortStrategy) Name() string {
	return "按面积排序"
}

// PerimeterSortStrategy 按周长排序
type PerimeterSortStrategy struct{}

// Sort 按周长排序
func (pss PerimeterSortStrategy) Sort(shapes []Shape) []Shape {
	result := make([]Shape, len(shapes))
	copy(result, shapes)
	
	for i := 0; i < len(result); i++ {
		for j := 0; j < len(result)-1-i; j++ {
			if result[j].Perimeter() > result[j+1].Perimeter() {
				result[j], result[j+1] = result[j+1], result[j]
			}
		}
	}
	
	return result
}

// Name 策略名称
func (pss PerimeterSortStrategy) Name() string {
	return "按周长排序"
}

// ShapeSorter 形状排序器
type ShapeSorter struct {
	strategy SortStrategy
}

// SetStrategy 设置排序策略
func (ss *ShapeSorter) SetStrategy(strategy SortStrategy) {
	ss.strategy = strategy
}

// Sort 排序
func (ss *ShapeSorter) Sort(shapes []Shape) []Shape {
	if ss.strategy == nil {
		return shapes
	}
	return ss.strategy.Sort(shapes)
}

// demoStrategyPattern 演示策略模式
func demoStrategyPattern() {
	// 1. 创建形状数据
	shapes := []Shape{
		Rectangle{Width: 4, Height: 3},
		Circle{Radius: 2},
		Rectangle{Width: 2, Height: 6},
		Circle{Radius: 3},
		Triangle{Base: 4, Height: 5},
	}
	
	fmt.Println("原始形状:")
	for i, shape := range shapes {
		fmt.Printf("%d. %s (面积: %.2f, 周长: %.2f)\n", 
			i+1, shape.String(), shape.Area(), shape.Perimeter())
	}
	
	// 2. 使用不同的排序策略
	sorter := &ShapeSorter{}
	
	strategies := []SortStrategy{
		AreaSortStrategy{},
		PerimeterSortStrategy{},
	}
	
	for _, strategy := range strategies {
		sorter.SetStrategy(strategy)
		sorted := sorter.Sort(shapes)
		
		fmt.Printf("\n%s:\n", strategy.Name())
		for i, shape := range sorted {
			fmt.Printf("%d. %s (面积: %.2f, 周长: %.2f)\n", 
				i+1, shape.String(), shape.Area(), shape.Perimeter())
		}
	}
}

// Processor 处理器接口
type Processor interface {
	Process(data interface{}) interface{}
	Name() string
}

// StringProcessor 字符串处理器
type StringProcessor struct{}

// Process 处理字符串
func (sp StringProcessor) Process(data interface{}) interface{} {
	if str, ok := data.(string); ok {
		return fmt.Sprintf("处理后的字符串: %s", str)
	}
	return "无法处理非字符串数据"
}

// Name 处理器名称
func (sp StringProcessor) Name() string {
	return "字符串处理器"
}

// NumberProcessor 数字处理器
type NumberProcessor struct{}

// Process 处理数字
func (np NumberProcessor) Process(data interface{}) interface{} {
	switch v := data.(type) {
	case int:
		return v * 2
	case float64:
		return v * 2.0
	default:
		return "无法处理非数字数据"
	}
}

// Name 处理器名称
func (np NumberProcessor) Name() string {
	return "数字处理器"
}

// ShapeProcessor 形状处理器
type ShapeProcessor struct{}

// Process 处理形状
func (sp ShapeProcessor) Process(data interface{}) interface{} {
	if shape, ok := data.(Shape); ok {
		return fmt.Sprintf("形状信息: %s, 面积: %.2f", 
			shape.String(), shape.Area())
	}
	return "无法处理非形状数据"
}

// Name 处理器名称
func (sp ShapeProcessor) Name() string {
	return "形状处理器"
}

// demoPolymorphismInPractice 演示多态的实际应用
func demoPolymorphismInPractice() {
	// 1. 创建处理器
	processors := []Processor{
		StringProcessor{},
		NumberProcessor{},
		ShapeProcessor{},
	}
	
	// 2. 创建不同类型的数据
	data := []interface{}{
		"Hello, World!",
		42,
		3.14,
		Rectangle{Width: 5, Height: 3},
		"Go语言",
		Circle{Radius: 2},
	}
	
	fmt.Println("多态数据处理:")
	
	// 3. 使用多态处理数据
	for i, item := range data {
		fmt.Printf("\n数据 %d: %v (类型: %T)\n", i+1, item, item)
		
		for _, processor := range processors {
			result := processor.Process(item)
			fmt.Printf("  %s: %v\n", processor.Name(), result)
		}
	}
	
	// 4. 智能处理器选择
	fmt.Println("\n智能处理器选择:")
	
	smartProcess := func(data interface{}) {
		fmt.Printf("处理数据: %v (类型: %T)\n", data, data)
		
		for _, processor := range processors {
			result := processor.Process(data)
			if resultStr, ok := result.(string); ok {
				if resultStr != "无法处理非字符串数据" && 
				   resultStr != "无法处理非数字数据" && 
				   resultStr != "无法处理非形状数据" {
					fmt.Printf("  使用 %s: %v\n", processor.Name(), result)
					return
				}
			} else {
				fmt.Printf("  使用 %s: %v\n", processor.Name(), result)
				return
			}
		}
		fmt.Println("  没有合适的处理器")
	}
	
	for _, item := range data {
		smartProcess(item)
	}
}