package stage1

import (
	"fmt"
	"unsafe"
)

// DemoPointers 演示指针基础
func DemoPointers() {
	fmt.Println("\n=== 指针基础演示 ===")

	// 1. 指针的基本概念
	fmt.Println("\n1. 指针的基本概念：")
	demoBasicPointers()

	// 2. 指针的零值
	fmt.Println("\n2. 指针的零值：")
	demoNilPointers()

	// 3. 指针运算
	fmt.Println("\n3. 指针操作：")
	demoPointerOperations()

	// 4. 指针作为函数参数
	fmt.Println("\n4. 指针作为函数参数：")
	demoPointerParameters()

	// 5. 指针与数组
	fmt.Println("\n5. 指针与数组：")
	demoPointersAndArrays()
}

// DemoPointersAdvanced 演示指针高级用法
func DemoPointersAdvanced() {
	fmt.Println("\n=== 指针高级用法演示 ===")

	// 1. 指针的指针
	fmt.Println("\n1. 指针的指针：")
	demoPointerToPointer()

	// 2. 函数指针
	fmt.Println("\n2. 函数指针：")
	demoFunctionPointers()

	// 3. 结构体指针
	fmt.Println("\n3. 结构体指针：")
	demoStructPointers()

	// 4. 指针与内存管理
	fmt.Println("\n4. 指针与内存管理：")
	demoMemoryManagement()

	// 5. unsafe 包的使用
	fmt.Println("\n5. unsafe 包的使用：")
	demoUnsafePointers()
}

// demoBasicPointers 演示指针基础
func demoBasicPointers() {
	// 声明变量
	var num int = 42
	var str string = "Hello"

	// 获取变量的地址（指针）
	var numPtr *int = &num
	var strPtr *string = &str

	fmt.Printf("变量 num 的值: %d\n", num)
	fmt.Printf("变量 num 的地址: %p\n", &num)
	fmt.Printf("指针 numPtr 的值: %p\n", numPtr)
	fmt.Printf("指针 numPtr 指向的值: %d\n", *numPtr)

	fmt.Printf("变量 str 的值: %s\n", str)
	fmt.Printf("变量 str 的地址: %p\n", &str)
	fmt.Printf("指针 strPtr 的值: %p\n", strPtr)
	fmt.Printf("指针 strPtr 指向的值: %s\n", *strPtr)

	// 通过指针修改值
	*numPtr = 100
	*strPtr = "World"

	fmt.Printf("通过指针修改后 num: %d\n", num)
	fmt.Printf("通过指针修改后 str: %s\n", str)

	// 短变量声明
	value := 123
	ptr := &value
	fmt.Printf("短声明 - 值: %d, 地址: %p, 指针: %p, 解引用: %d\n",
		value, &value, ptr, *ptr)
}

// demoNilPointers 演示nil指针
func demoNilPointers() {
	var ptr *int
	fmt.Printf("未初始化的指针: %v\n", ptr)
	fmt.Printf("指针是否为nil: %t\n", ptr == nil)

	// 检查指针是否为nil再使用
	if ptr != nil {
		fmt.Printf("指针指向的值: %d\n", *ptr)
	} else {
		fmt.Println("指针为nil，不能解引用")
	}

	// 初始化指针
	num := 42
	ptr = &num
	fmt.Printf("初始化后的指针: %p\n", ptr)
	fmt.Printf("指针是否为nil: %t\n", ptr == nil)
	fmt.Printf("指针指向的值: %d\n", *ptr)

	// 将指针设为nil
	ptr = nil
	fmt.Printf("设为nil后的指针: %v\n", ptr)
}

// demoPointerOperations 演示指针操作
func demoPointerOperations() {
	// 指针比较
	a := 10
	b := 20
	c := 10

	ptrA := &a
	ptrB := &b
	ptrC := &c
	ptrA2 := &a

	fmt.Printf("ptrA == ptrB: %t (不同变量的地址)\n", ptrA == ptrB)
	fmt.Printf("ptrA == ptrA2: %t (同一变量的地址)\n", ptrA == ptrA2)
	fmt.Printf("ptrA == ptrC: %t (不同变量，相同值)\n", ptrA == ptrC)

	// 指针的值比较
	fmt.Printf("*ptrA == *ptrC: %t (指向的值相同)\n", *ptrA == *ptrC)

	// 指针类型
	var intPtr *int
	var floatPtr *float64
	// intPtr = floatPtr // 编译错误：类型不匹配

	fmt.Printf("intPtr 类型: %T\n", intPtr)
	fmt.Printf("floatPtr 类型: %T\n", floatPtr)
}

// demoPointerParameters 演示指针作为函数参数
func demoPointerParameters() {
	// 值传递 vs 指针传递
	original := 100
	fmt.Printf("原始值: %d\n", original)

	// 值传递 - 不会修改原始值
	modifyByValue(original)
	fmt.Printf("值传递后: %d\n", original)

	// 指针传递 - 会修改原始值
	modifyByPointer(&original)
	fmt.Printf("指针传递后: %d\n", original)

	// 交换两个变量的值
	x, y := 10, 20
	fmt.Printf("交换前: x=%d, y=%d\n", x, y)
	swap(&x, &y)
	fmt.Printf("交换后: x=%d, y=%d\n", x, y)

	// 函数返回指针
	ptr := createInt(42)
	fmt.Printf("函数返回的指针: %p, 值: %d\n", ptr, *ptr)
}

// modifyByValue 值传递，不会修改原始值
func modifyByValue(num int) {
	num = 999
	fmt.Printf("  函数内修改为: %d\n", num)
}

// modifyByPointer 指针传递，会修改原始值
func modifyByPointer(ptr *int) {
	*ptr = 999
	fmt.Printf("  通过指针修改为: %d\n", *ptr)
}

// swap 交换两个整数的值
func swap(a, b *int) {
	*a, *b = *b, *a
}

// createInt 创建一个整数并返回其指针
func createInt(value int) *int {
	num := value
	return &num // 返回局部变量的地址是安全的，Go会自动处理
}

// demoPointersAndArrays 演示指针与数组
func demoPointersAndArrays() {
	// 数组指针
	arr := [5]int{1, 2, 3, 4, 5}
	arrPtr := &arr

	fmt.Printf("数组: %v\n", arr)
	fmt.Printf("数组指针: %p\n", arrPtr)
	fmt.Printf("通过指针访问数组: %v\n", *arrPtr)

	// 修改数组元素
	(*arrPtr)[0] = 100
	fmt.Printf("修改后的数组: %v\n", arr)

	// 指针数组
	a, b, c := 10, 20, 30
	ptrArray := [3]*int{&a, &b, &c}

	fmt.Printf("指针数组: %v\n", ptrArray)
	for i, ptr := range ptrArray {
		fmt.Printf("  索引 %d: 地址 %p, 值 %d\n", i, ptr, *ptr)
	}

	// 通过指针数组修改值
	*ptrArray[1] = 200
	fmt.Printf("修改后 b 的值: %d\n", b)

	// 切片与指针
	slice := []int{1, 2, 3, 4, 5}
	fmt.Printf("切片: %v\n", slice)

	// 获取切片元素的指针
	elemPtr := &slice[2]
	fmt.Printf("第3个元素的指针: %p, 值: %d\n", elemPtr, *elemPtr)

	*elemPtr = 300
	fmt.Printf("修改后的切片: %v\n", slice)
}

// demoPointerToPointer 演示指针的指针
func demoPointerToPointer() {
	value := 42
	ptr := &value
	ptrPtr := &ptr

	fmt.Printf("值: %d\n", value)
	fmt.Printf("指针: %p\n", ptr)
	fmt.Printf("指针的指针: %p\n", ptrPtr)

	fmt.Printf("通过指针访问值: %d\n", *ptr)
	fmt.Printf("通过指针的指针访问值: %d\n", **ptrPtr)

	// 修改值
	**ptrPtr = 100
	fmt.Printf("通过指针的指针修改后的值: %d\n", value)

	// 修改指针
	newValue := 200
	*ptrPtr = &newValue
	fmt.Printf("修改指针后，原值: %d, 新值: %d\n", value, *ptr)
}

// demoFunctionPointers 演示函数指针
func demoFunctionPointers() {
	// 函数变量
	var operation func(int, int) int

	// 赋值不同的函数
	operation = addFunc
	fmt.Printf("加法: %d\n", operation(10, 5))

	operation = multiplyFunc
	fmt.Printf("乘法: %d\n", operation(10, 5))

	// 函数指针数组
	operations := []func(int, int) int{addFunc, subtractFunc, multiplyFunc}
	operationNames := []string{"加法", "减法", "乘法"}

	for i, op := range operations {
		result := op(10, 5)
		fmt.Printf("%s: %d\n", operationNames[i], result)
	}

	// 将函数作为参数传递
	calculate(10, 5, addFunc)
	calculate(10, 5, multiplyFunc)
}

func addFunc(a, b int) int      { return a + b }
func subtractFunc(a, b int) int { return a - b }
func multiplyFunc(a, b int) int { return a * b }

func calculate(a, b int, op func(int, int) int) {
	result := op(a, b)
	fmt.Printf("计算结果: %d\n", result)
}

// Person 结构体用于演示
type Person struct {
	Name string
	Age  int
}

// demoStructPointers 演示结构体指针
func demoStructPointers() {
	// 创建结构体
	person := Person{Name: "Alice", Age: 25}
	fmt.Printf("结构体: %+v\n", person)

	// 结构体指针
	personPtr := &person
	fmt.Printf("结构体指针: %p\n", personPtr)

	// 通过指针访问字段
	fmt.Printf("通过指针访问姓名: %s\n", (*personPtr).Name)
	fmt.Printf("通过指针访问年龄: %d\n", personPtr.Age) // Go的语法糖

	// 通过指针修改字段
	personPtr.Name = "Bob"
	personPtr.Age = 30
	fmt.Printf("修改后的结构体: %+v\n", person)

	// 使用new创建结构体指针
	personPtr2 := new(Person)
	personPtr2.Name = "Carol"
	personPtr2.Age = 35
	fmt.Printf("使用new创建: %+v\n", *personPtr2)

	// 结构体指针作为函数参数
	updatePerson(personPtr2, "David", 40)
	fmt.Printf("函数修改后: %+v\n", *personPtr2)
}

func updatePerson(p *Person, name string, age int) {
	p.Name = name
	p.Age = age
}

// demoMemoryManagement 演示内存管理
func demoMemoryManagement() {
	// 栈上分配
	stackVar := 42
	fmt.Printf("栈变量地址: %p\n", &stackVar)

	// 堆上分配
	heapVar := new(int)
	*heapVar = 42
	fmt.Printf("堆变量地址: %p, 值: %d\n", heapVar, *heapVar)

	// 返回局部变量的指针（逃逸到堆）
	ptr := createLocalVar()
	fmt.Printf("逃逸变量地址: %p, 值: %d\n", ptr, *ptr)

	// 大对象通常分配在堆上
	bigArray := make([]int, 1000000)
	fmt.Printf("大数组地址: %p\n", &bigArray[0])

	// 垃圾回收会自动处理内存释放
	// Go没有手动内存管理，不需要free()
}

func createLocalVar() *int {
	local := 123
	return &local // 这个变量会逃逸到堆上
}

// demoUnsafePointers 演示unsafe包的使用
func demoUnsafePointers() {
	fmt.Println("注意：unsafe包的使用需要谨慎，可能导致程序崩溃")

	// 基本unsafe操作
	num := int64(42)
	ptr := unsafe.Pointer(&num)

	fmt.Printf("原始值: %d\n", num)
	fmt.Printf("unsafe.Pointer: %p\n", ptr)

	// 类型转换（危险操作）
	intPtr := (*int64)(ptr)
	fmt.Printf("转换回int64指针的值: %d\n", *intPtr)

	// 获取变量大小
	fmt.Printf("int64大小: %d字节\n", unsafe.Sizeof(num))
	fmt.Printf("指针大小: %d字节\n", unsafe.Sizeof(ptr))

	// 结构体字段偏移
	p := Person{Name: "Alice", Age: 25}
	fmt.Printf("Person大小: %d字节\n", unsafe.Sizeof(p))
	fmt.Printf("Name字段偏移: %d字节\n", unsafe.Offsetof(p.Name))
	fmt.Printf("Age字段偏移: %d字节\n", unsafe.Offsetof(p.Age))

	// 通过偏移访问字段（危险操作）
	personPtr := unsafe.Pointer(&p)
	namePtr := (*string)(unsafe.Pointer(uintptr(personPtr) + unsafe.Offsetof(p.Name)))
	agePtr := (*int)(unsafe.Pointer(uintptr(personPtr) + unsafe.Offsetof(p.Age)))

	fmt.Printf("通过偏移访问Name: %s\n", *namePtr)
	fmt.Printf("通过偏移访问Age: %d\n", *agePtr)

	// 字符串和字节切片的零拷贝转换（高级用法）
	str := "Hello, World!"
	strPtr := unsafe.Pointer(&str)

	// 这是一个危险的操作示例，实际开发中应避免
	fmt.Printf("字符串长度: %d\n", len(str))
	fmt.Printf("字符串数据指针: %p\n", strPtr)
}
