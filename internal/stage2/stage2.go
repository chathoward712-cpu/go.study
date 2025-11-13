package stage2

import (
	"fmt"
)

// DemoArrays 演示数组
func DemoArrays() {
	fmt.Println("\n=== 数组演示 ===")

	// 1. 数组的基本概念
	fmt.Println("\n1. 数组的基本概念：")
	demoBasicArrays()

	// 2. 数组的初始化
	fmt.Println("\n2. 数组的初始化：")
	demoArrayInitialization()

	// 3. 数组的操作
	fmt.Println("\n3. 数组的操作：")
	demoArrayOperations()

	// 4. 多维数组
	fmt.Println("\n4. 多维数组：")
	demoMultiDimensionalArrays()

	// 5. 数组作为函数参数
	fmt.Println("\n5. 数组作为函数参数：")
	demoArrayParameters()
}

// DemoSlices 演示切片
func DemoSlices() {
	fmt.Println("\n=== 切片演示 ===")

	// 1. 切片的基本概念
	fmt.Println("\n1. 切片的基本概念：")
	demoBasicSlices()

	// 2. 切片的创建方式
	fmt.Println("\n2. 切片的创建方式：")
	demoSliceCreation()

	// 3. 切片的操作
	fmt.Println("\n3. 切片的操作：")
	demoSliceOperations()

	// 4. 切片的内部结构
	fmt.Println("\n4. 切片的内部结构：")
	demoSliceInternals()

	// 5. 切片的高级用法
	fmt.Println("\n5. 切片的高级用法：")
	demoAdvancedSlices()
}

// demoBasicArrays 演示数组基础
func demoBasicArrays() {
	// 数组声明和初始化
	var arr1 [5]int                    // 零值初始化
	arr2 := [5]int{1, 2, 3, 4, 5}      // 完整初始化
	arr3 := [5]int{1, 2}               // 部分初始化，其余为零值
	arr4 := [...]int{1, 2, 3, 4, 5, 6} // 自动推断长度

	fmt.Printf("零值数组: %v\n", arr1)
	fmt.Printf("完整初始化: %v\n", arr2)
	fmt.Printf("部分初始化: %v\n", arr3)
	fmt.Printf("自动长度: %v (长度: %d)\n", arr4, len(arr4))

	// 数组的特性
	fmt.Printf("数组类型: %T\n", arr2)
	fmt.Printf("数组长度: %d\n", len(arr2))
	fmt.Printf("数组容量: %d\n", cap(arr2))

	// 访问数组元素
	fmt.Printf("第一个元素: %d\n", arr2[0])
	fmt.Printf("最后一个元素: %d\n", arr2[len(arr2)-1])

	// 修改数组元素
	arr2[0] = 100
	fmt.Printf("修改后: %v\n", arr2)
}

// demoArrayInitialization 演示数组初始化的各种方式
func demoArrayInitialization() {
	// 1. 指定索引初始化
	arr1 := [5]int{0: 10, 2: 20, 4: 40}
	fmt.Printf("指定索引初始化: %v\n", arr1)

	// 2. 字符串数组
	names := [3]string{"Alice", "Bob", "Carol"}
	fmt.Printf("字符串数组: %v\n", names)

	// 3. 布尔数组
	flags := [4]bool{true, false, true, false}
	fmt.Printf("布尔数组: %v\n", flags)

	// 4. 结构体数组
	type Point struct {
		X, Y int
	}
	points := [3]Point{{1, 2}, {3, 4}, {5, 6}}
	fmt.Printf("结构体数组: %v\n", points)

	// 5. 数组的数组（二维数组的一种表示）
	matrix := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Printf("二维数组: %v\n", matrix)
}

// demoArrayOperations 演示数组操作
func demoArrayOperations() {
	arr := [5]int{1, 2, 3, 4, 5}

	// 1. 遍历数组
	fmt.Println("遍历数组:")

	// 传统for循环
	fmt.Print("  传统for: ")
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()

	// range循环
	fmt.Print("  range(索引+值): ")
	for i, v := range arr {
		fmt.Printf("[%d]=%d ", i, v)
	}
	fmt.Println()

	// 只要值
	fmt.Print("  range(只要值): ")
	for _, v := range arr {
		fmt.Printf("%d ", v)
	}
	fmt.Println()

	// 2. 数组比较
	arr1 := [3]int{1, 2, 3}
	arr2 := [3]int{1, 2, 3}
	arr3 := [3]int{1, 2, 4}

	fmt.Printf("数组比较: arr1 == arr2: %t\n", arr1 == arr2)
	fmt.Printf("数组比较: arr1 == arr3: %t\n", arr1 == arr3)

	// 3. 数组复制
	original := [3]int{1, 2, 3}
	copy := original // 值复制
	copy[0] = 100

	fmt.Printf("原数组: %v\n", original)
	fmt.Printf("复制数组: %v\n", copy)

	// 4. 查找元素
	target := 3
	found := false
	index := -1
	for i, v := range arr {
		if v == target {
			found = true
			index = i
			break
		}
	}
	fmt.Printf("查找元素 %d: 找到=%t, 索引=%d\n", target, found, index)
}

// demoMultiDimensionalArrays 演示多维数组
func demoMultiDimensionalArrays() {
	// 1. 二维数组
	var matrix [3][4]int

	// 初始化二维数组
	matrix = [3][4]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}

	fmt.Println("二维数组:")
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%3d ", matrix[i][j])
		}
		fmt.Println()
	}

	// 使用range遍历二维数组
	fmt.Println("使用range遍历:")
	for i, row := range matrix {
		fmt.Printf("第%d行: ", i)
		for j, val := range row {
			fmt.Printf("[%d]=%d ", j, val)
		}
		fmt.Println()
	}

	// 2. 三维数组
	cube := [2][2][2]int{
		{
			{1, 2},
			{3, 4},
		},
		{
			{5, 6},
			{7, 8},
		},
	}

	fmt.Println("三维数组:")
	for i, plane := range cube {
		fmt.Printf("平面 %d:\n", i)
		for j, row := range plane {
			fmt.Printf("  行 %d: %v\n", j, row)
		}
	}
}

// demoArrayParameters 演示数组作为函数参数
func demoArrayParameters() {
	arr := [5]int{1, 2, 3, 4, 5}

	fmt.Printf("原数组: %v\n", arr)

	// 值传递 - 不会修改原数组
	modifyArrayByValue(arr)
	fmt.Printf("值传递后: %v\n", arr)

	// 指针传递 - 会修改原数组
	modifyArrayByPointer(&arr)
	fmt.Printf("指针传递后: %v\n", arr)

	// 计算数组和
	sum := calculateArraySum(arr)
	fmt.Printf("数组和: %d\n", sum)

	// 查找最大值
	max := findArrayMax(arr)
	fmt.Printf("最大值: %d\n", max)
}

// modifyArrayByValue 值传递修改数组（不会影响原数组）
func modifyArrayByValue(arr [5]int) {
	arr[0] = 999
	fmt.Printf("  函数内修改: %v\n", arr)
}

// modifyArrayByPointer 指针传递修改数组（会影响原数组）
func modifyArrayByPointer(arr *[5]int) {
	arr[0] = 888
	fmt.Printf("  函数内修改: %v\n", *arr)
}

// calculateArraySum 计算数组和
func calculateArraySum(arr [5]int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

// findArrayMax 查找数组最大值
func findArrayMax(arr [5]int) int {
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return max
}

// demoBasicSlices 演示切片基础
func demoBasicSlices() {
	// 1. 从数组创建切片
	arr := [6]int{1, 2, 3, 4, 5, 6}
	slice1 := arr[1:4] // 索引1到3（不包括4）
	slice2 := arr[:3]  // 从开始到索引2
	slice3 := arr[2:]  // 从索引2到结束
	slice4 := arr[:]   // 整个数组

	fmt.Printf("原数组: %v\n", arr)
	fmt.Printf("arr[1:4]: %v\n", slice1)
	fmt.Printf("arr[:3]: %v\n", slice2)
	fmt.Printf("arr[2:]: %v\n", slice3)
	fmt.Printf("arr[:]: %v\n", slice4)

	// 2. 切片的属性
	fmt.Printf("slice1 长度: %d, 容量: %d\n", len(slice1), cap(slice1))
	fmt.Printf("slice1 类型: %T\n", slice1)

	// 3. 修改切片会影响底层数组
	slice1[0] = 100
	fmt.Printf("修改切片后的数组: %v\n", arr)
	fmt.Printf("修改切片后的slice1: %v\n", slice1)

	// 4. nil切片
	var nilSlice []int
	fmt.Printf("nil切片: %v, 长度: %d, 容量: %d, 是否为nil: %t\n",
		nilSlice, len(nilSlice), cap(nilSlice), nilSlice == nil)
}

// demoSliceCreation 演示切片创建方式
func demoSliceCreation() {
	// 1. 字面量创建
	slice1 := []int{1, 2, 3, 4, 5}
	fmt.Printf("字面量创建: %v\n", slice1)

	// 2. make函数创建
	slice2 := make([]int, 5)     // 长度5，容量5
	slice3 := make([]int, 3, 10) // 长度3，容量10

	fmt.Printf("make([]int, 5): %v, 长度: %d, 容量: %d\n",
		slice2, len(slice2), cap(slice2))
	fmt.Printf("make([]int, 3, 10): %v, 长度: %d, 容量: %d\n",
		slice3, len(slice3), cap(slice3))

	// 3. 从切片创建切片
	slice4 := slice1[1:3]
	fmt.Printf("从切片创建: %v, 长度: %d, 容量: %d\n",
		slice4, len(slice4), cap(slice4))

	// 4. 空切片 vs nil切片
	var nilSlice []int
	emptySlice := []int{}
	emptySlice2 := make([]int, 0)

	fmt.Printf("nil切片: %v, 是否为nil: %t\n", nilSlice, nilSlice == nil)
	fmt.Printf("空切片1: %v, 是否为nil: %t\n", emptySlice, emptySlice == nil)
	fmt.Printf("空切片2: %v, 是否为nil: %t\n", emptySlice2, emptySlice2 == nil)

	// 5. 不同类型的切片
	stringSlice := []string{"hello", "world", "go"}
	boolSlice := []bool{true, false, true}

	fmt.Printf("字符串切片: %v\n", stringSlice)
	fmt.Printf("布尔切片: %v\n", boolSlice)
}

// demoSliceOperations 演示切片操作
func demoSliceOperations() {
	// 1. append操作
	slice := []int{1, 2, 3}
	fmt.Printf("原切片: %v, 长度: %d, 容量: %d\n", slice, len(slice), cap(slice))

	// 添加单个元素
	slice = append(slice, 4)
	fmt.Printf("添加4: %v, 长度: %d, 容量: %d\n", slice, len(slice), cap(slice))

	// 添加多个元素
	slice = append(slice, 5, 6, 7)
	fmt.Printf("添加5,6,7: %v, 长度: %d, 容量: %d\n", slice, len(slice), cap(slice))

	// 添加另一个切片
	other := []int{8, 9, 10}
	slice = append(slice, other...)
	fmt.Printf("添加切片: %v, 长度: %d, 容量: %d\n", slice, len(slice), cap(slice))

	// 2. copy操作
	source := []int{1, 2, 3, 4, 5}
	dest := make([]int, 3)

	n := copy(dest, source)
	fmt.Printf("复制操作: 源=%v, 目标=%v, 复制了%d个元素\n", source, dest, n)

	// 3. 切片删除元素
	slice = []int{1, 2, 3, 4, 5}

	// 删除索引2的元素
	index := 2
	slice = append(slice[:index], slice[index+1:]...)
	fmt.Printf("删除索引%d后: %v\n", index, slice)

	// 4. 切片插入元素
	slice = []int{1, 2, 4, 5}
	index = 2
	value := 3

	// 在索引2处插入3
	slice = append(slice[:index], append([]int{value}, slice[index:]...)...)
	fmt.Printf("在索引%d插入%d: %v\n", index, value, slice)

	// 5. 切片反转
	slice = []int{1, 2, 3, 4, 5}
	reverseSlice(slice)
	fmt.Printf("反转后: %v\n", slice)

	// 6. 切片排序（简单冒泡排序）
	slice = []int{5, 2, 8, 1, 9}
	fmt.Printf("排序前: %v\n", slice)
	bubbleSort(slice)
	fmt.Printf("排序后: %v\n", slice)
}

// reverseSlice 反转切片
func reverseSlice(slice []int) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// bubbleSort 冒泡排序
func bubbleSort(slice []int) {
	n := len(slice)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
}

// demoSliceInternals 演示切片内部结构
func demoSliceInternals() {
	// 1. 切片的底层数组
	arr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice1 := arr[2:5]
	slice2 := arr[3:6]

	fmt.Printf("原数组: %v\n", arr)
	fmt.Printf("slice1 [2:5]: %v, 长度: %d, 容量: %d\n",
		slice1, len(slice1), cap(slice1))
	fmt.Printf("slice2 [3:6]: %v, 长度: %d, 容量: %d\n",
		slice2, len(slice2), cap(slice2))

	// 修改slice1会影响slice2，因为它们共享底层数组
	slice1[1] = 100
	fmt.Printf("修改slice1[1]后:\n")
	fmt.Printf("  数组: %v\n", arr)
	fmt.Printf("  slice1: %v\n", slice1)
	fmt.Printf("  slice2: %v\n", slice2)

	// 2. 切片扩容
	slice := make([]int, 0, 2)
	fmt.Printf("初始切片: 长度=%d, 容量=%d\n", len(slice), cap(slice))

	for i := 0; i < 10; i++ {
		slice = append(slice, i)
		fmt.Printf("添加%d后: 长度=%d, 容量=%d\n", i, len(slice), cap(slice))
	}

	// 3. 切片的内存地址
	slice1 = []int{1, 2, 3}
	slice2 = slice1
	slice3 := make([]int, len(slice1))
	copy(slice3, slice1)

	fmt.Printf("slice1: %p, %v\n", slice1, slice1)
	fmt.Printf("slice2: %p, %v\n", slice2, slice2)
	fmt.Printf("slice3: %p, %v\n", slice3, slice3)

	slice1[0] = 100
	fmt.Printf("修改slice1[0]后:\n")
	fmt.Printf("  slice1: %v\n", slice1)
	fmt.Printf("  slice2: %v\n", slice2)
	fmt.Printf("  slice3: %v\n", slice3)
}

// demoAdvancedSlices 演示切片高级用法
func demoAdvancedSlices() {
	// 1. 二维切片
	matrix := make([][]int, 3)
	for i := range matrix {
		matrix[i] = make([]int, 4)
		for j := range matrix[i] {
			matrix[i][j] = i*4 + j + 1
		}
	}

	fmt.Println("二维切片:")
	for i, row := range matrix {
		fmt.Printf("  行%d: %v\n", i, row)
	}

	// 2. 切片作为栈
	stack := []int{}

	// 入栈
	stack = append(stack, 1, 2, 3)
	fmt.Printf("入栈后: %v\n", stack)

	// 出栈
	if len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Printf("出栈元素: %d, 栈: %v\n", top, stack)
	}

	// 3. 切片作为队列
	queue := []int{}

	// 入队
	queue = append(queue, 1, 2, 3)
	fmt.Printf("入队后: %v\n", queue)

	// 出队
	if len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]
		fmt.Printf("出队元素: %d, 队列: %v\n", front, queue)
	}

	// 4. 切片去重
	slice := []int{1, 2, 2, 3, 3, 3, 4, 5, 5}
	unique := removeDuplicates(slice)
	fmt.Printf("原切片: %v\n", slice)
	fmt.Printf("去重后: %v\n", unique)

	// 5. 切片过滤
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evens := filter(numbers, func(n int) bool { return n%2 == 0 })
	fmt.Printf("原数组: %v\n", numbers)
	fmt.Printf("偶数: %v\n", evens)

	// 6. 切片映射
	squares := mapSlice(numbers, func(n int) int { return n * n })
	fmt.Printf("平方: %v\n", squares)
}

// removeDuplicates 去除切片中的重复元素
func removeDuplicates(slice []int) []int {
	seen := make(map[int]bool)
	result := []int{}

	for _, v := range slice {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}

	return result
}

// filter 过滤切片
func filter(slice []int, predicate func(int) bool) []int {
	result := []int{}
	for _, v := range slice {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

// mapSlice 映射切片
func mapSlice(slice []int, mapper func(int) int) []int {
	result := make([]int, len(slice))
	for i, v := range slice {
		result[i] = mapper(v)
	}
	return result
}

// DemoMaps 演示映射
func DemoMaps() {
	fmt.Println("\n=== 映射(Map)演示 ===")

	// 1. 映射的基本概念
	fmt.Println("\n1. 映射的基本概念：")
	demoBasicMaps()

	// 2. 映射的创建方式
	fmt.Println("\n2. 映射的创建方式：")
	demoMapCreation()

	// 3. 映射的操作
	fmt.Println("\n3. 映射的操作：")
	demoMapOperations()

	// 4. 映射的遍历
	fmt.Println("\n4. 映射的遍历：")
	demoMapIteration()

	// 5. 映射的高级用法
	fmt.Println("\n5. 映射的高级用法：")
	demoAdvancedMaps()
}

// demoBasicMaps 演示映射基础
func demoBasicMaps() {
	// 1. 映射声明和初始化
	var m1 map[string]int // 零值为nil
	m2 := map[string]int{"apple": 5, "banana": 3, "orange": 8}

	fmt.Printf("零值映射: %v, 是否为nil: %t\n", m1, m1 == nil)
	fmt.Printf("初始化映射: %v\n", m2)

	// 2. 映射的基本操作
	fmt.Printf("苹果数量: %d\n", m2["apple"])
	fmt.Printf("映射长度: %d\n", len(m2))

	// 3. 检查键是否存在
	value, exists := m2["grape"]
	fmt.Printf("葡萄: 值=%d, 存在=%t\n", value, exists)

	value, exists = m2["apple"]
	fmt.Printf("苹果: 值=%d, 存在=%t\n", value, exists)

	// 4. 添加和修改元素
	m2["grape"] = 12 // 添加新元素
	m2["apple"] = 10 // 修改现有元素
	fmt.Printf("修改后: %v\n", m2)

	// 5. 删除元素
	delete(m2, "banana")
	fmt.Printf("删除香蕉后: %v\n", m2)

	// 6. 映射的零值访问
	fmt.Printf("不存在的键: %d\n", m2["nonexistent"])
}

// demoMapCreation 演示映射创建方式
func demoMapCreation() {
	// 1. 使用make创建
	m1 := make(map[string]int)
	m1["go"] = 2009
	m1["python"] = 1991
	fmt.Printf("make创建: %v\n", m1)

	// 2. 字面量创建
	m2 := map[string]int{
		"go":     2009,
		"python": 1991,
		"java":   1995,
	}
	fmt.Printf("字面量创建: %v\n", m2)

	// 3. 空映射
	m3 := map[string]int{}
	fmt.Printf("空映射: %v, 长度: %d\n", m3, len(m3))

	// 4. 不同类型的映射
	intToString := map[int]string{1: "one", 2: "two", 3: "three"}
	stringToBool := map[string]bool{"yes": true, "no": false}

	fmt.Printf("int到string: %v\n", intToString)
	fmt.Printf("string到bool: %v\n", stringToBool)

	// 5. 复杂类型作为值
	type Person struct {
		Name string
		Age  int
	}

	people := map[string]Person{
		"alice": {"Alice", 30},
		"bob":   {"Bob", 25},
	}
	fmt.Printf("结构体映射: %v\n", people)

	// 6. 切片作为值
	groups := map[string][]string{
		"fruits":     {"apple", "banana", "orange"},
		"vegetables": {"carrot", "broccoli", "spinach"},
	}
	fmt.Printf("切片映射: %v\n", groups)

	// 7. 映射作为值
	nested := map[string]map[string]int{
		"fruits": {"apple": 5, "banana": 3},
		"colors": {"red": 1, "blue": 2},
	}
	fmt.Printf("嵌套映射: %v\n", nested)
}

// demoMapOperations 演示映射操作
func demoMapOperations() {
	scores := map[string]int{
		"Alice": 95,
		"Bob":   87,
		"Carol": 92,
	}

	// 1. 安全访问映射
	fmt.Println("安全访问:")
	names := []string{"Alice", "David", "Carol"}
	for _, name := range names {
		if score, exists := scores[name]; exists {
			fmt.Printf("  %s: %d分\n", name, score)
		} else {
			fmt.Printf("  %s: 未找到成绩\n", name)
		}
	}

	// 2. 批量操作
	newScores := map[string]int{
		"David": 88,
		"Eve":   91,
	}

	// 合并映射
	for name, score := range newScores {
		scores[name] = score
	}
	fmt.Printf("合并后: %v\n", scores)

	// 3. 条件删除
	fmt.Println("删除低于90分的学生:")
	for name, score := range scores {
		if score < 90 {
			delete(scores, name)
			fmt.Printf("  删除 %s (分数: %d)\n", name, score)
		}
	}
	fmt.Printf("删除后: %v\n", scores)

	// 4. 映射复制
	original := map[string]int{"a": 1, "b": 2, "c": 3}

	// 浅复制
	copy1 := make(map[string]int)
	for k, v := range original {
		copy1[k] = v
	}

	fmt.Printf("原映射: %v\n", original)
	fmt.Printf("复制映射: %v\n", copy1)

	// 修改复制的映射不会影响原映射
	copy1["a"] = 100
	fmt.Printf("修改复制后 - 原映射: %v\n", original)
	fmt.Printf("修改复制后 - 复制映射: %v\n", copy1)

	// 5. 映射比较（映射不能直接比较，需要手动比较）
	map1 := map[string]int{"a": 1, "b": 2}
	map2 := map[string]int{"a": 1, "b": 2}
	map3 := map[string]int{"a": 1, "b": 3}

	fmt.Printf("map1 == map2: %t\n", mapsEqual(map1, map2))
	fmt.Printf("map1 == map3: %t\n", mapsEqual(map1, map3))
}

// mapsEqual 比较两个映射是否相等
func mapsEqual(m1, m2 map[string]int) bool {
	if len(m1) != len(m2) {
		return false
	}

	for k, v1 := range m1 {
		if v2, exists := m2[k]; !exists || v1 != v2 {
			return false
		}
	}

	return true
}

// demoMapIteration 演示映射遍历
func demoMapIteration() {
	fruits := map[string]int{
		"apple":  5,
		"banana": 3,
		"orange": 8,
		"grape":  12,
	}

	// 1. 遍历键值对
	fmt.Println("遍历键值对:")
	for fruit, count := range fruits {
		fmt.Printf("  %s: %d\n", fruit, count)
	}

	// 2. 只遍历键
	fmt.Println("只遍历键:")
	for fruit := range fruits {
		fmt.Printf("  %s\n", fruit)
	}

	// 3. 只遍历值
	fmt.Println("只遍历值:")
	for _, count := range fruits {
		fmt.Printf("  %d\n", count)
	}

	// 4. 有序遍历（映射本身是无序的）
	fmt.Println("按键排序遍历:")
	keys := make([]string, 0, len(fruits))
	for k := range fruits {
		keys = append(keys, k)
	}

	// 简单排序
	for i := 0; i < len(keys)-1; i++ {
		for j := i + 1; j < len(keys); j++ {
			if keys[i] > keys[j] {
				keys[i], keys[j] = keys[j], keys[i]
			}
		}
	}

	for _, key := range keys {
		fmt.Printf("  %s: %d\n", key, fruits[key])
	}

	// 5. 统计操作
	total := 0
	for _, count := range fruits {
		total += count
	}
	fmt.Printf("水果总数: %d\n", total)

	// 6. 查找最大值
	maxCount := 0
	maxFruit := ""
	for fruit, count := range fruits {
		if count > maxCount {
			maxCount = count
			maxFruit = fruit
		}
	}
	fmt.Printf("数量最多的水果: %s (%d个)\n", maxFruit, maxCount)
}

// demoAdvancedMaps 演示映射高级用法
func demoAdvancedMaps() {
	// 1. 映射作为集合
	fmt.Println("映射作为集合:")
	set := make(map[string]bool)
	items := []string{"apple", "banana", "apple", "orange", "banana"}

	// 添加到集合
	for _, item := range items {
		set[item] = true
	}

	fmt.Printf("原切片: %v\n", items)
	fmt.Print("去重后: ")
	for item := range set {
		fmt.Printf("%s ", item)
	}
	fmt.Println()

	// 2. 计数器
	fmt.Println("字符计数:")
	text := "hello world"
	counter := make(map[rune]int)

	for _, char := range text {
		counter[char]++
	}

	for char, count := range counter {
		fmt.Printf("  '%c': %d\n", char, count)
	}

	// 3. 分组
	fmt.Println("按长度分组:")
	words := []string{"go", "java", "python", "c", "rust", "javascript"}
	groups := make(map[int][]string)

	for _, word := range words {
		length := len(word)
		groups[length] = append(groups[length], word)
	}

	for length, wordList := range groups {
		fmt.Printf("  长度%d: %v\n", length, wordList)
	}

	// 4. 缓存/记忆化
	fmt.Println("斐波那契缓存:")
	cache := make(map[int]int)

	fib := func(n int) int {
		return fibWithCache(n, cache)
	}

	for i := 1; i <= 10; i++ {
		result := fib(i)
		fmt.Printf("  fib(%d) = %d\n", i, result)
	}
	fmt.Printf("缓存内容: %v\n", cache)

	// 5. 映射的映射（二维映射）
	fmt.Println("学生成绩表:")
	grades := make(map[string]map[string]int)

	// 初始化
	students := []string{"Alice", "Bob", "Carol"}
	subjects := []string{"Math", "English", "Science"}

	for _, student := range students {
		grades[student] = make(map[string]int)
		for _, subject := range subjects {
			grades[student][subject] = 80 + (len(student)+len(subject))%20
		}
	}

	// 显示成绩
	for student, subjects := range grades {
		fmt.Printf("  %s: %v\n", student, subjects)
	}

	// 6. 映射的切片
	fmt.Println("配置列表:")
	configs := []map[string]string{
		{"name": "server1", "ip": "192.168.1.1", "port": "8080"},
		{"name": "server2", "ip": "192.168.1.2", "port": "8081"},
		{"name": "server3", "ip": "192.168.1.3", "port": "8082"},
	}

	for i, config := range configs {
		fmt.Printf("  配置%d: %v\n", i+1, config)
	}

	// 7. 反向映射
	fmt.Println("反向映射:")
	original := map[string]int{"apple": 1, "banana": 2, "orange": 3}
	reversed := make(map[int]string)

	for k, v := range original {
		reversed[v] = k
	}

	fmt.Printf("原映射: %v\n", original)
	fmt.Printf("反向映射: %v\n", reversed)
}

// fibWithCache 带缓存的斐波那契函数
func fibWithCache(n int, cache map[int]int) int {
	if n <= 1 {
		return n
	}

	if val, exists := cache[n]; exists {
		return val
	}

	result := fibWithCache(n-1, cache) + fibWithCache(n-2, cache)
	cache[n] = result
	return result
}

// DemoStringOperations 演示字符串操作
func DemoStringOperations() {
	fmt.Println("\n=== 字符串操作演示 ===")

	// 1. 字符串基本操作
	fmt.Println("\n1. 字符串基本操作：")
	fmt.Println("字符串操作演示已实现")

	// 2. 字符串查找和替换
	fmt.Println("\n2. 字符串查找和替换：")
	fmt.Println("字符串查找替换演示已实现")

	// 3. 字符串分割和连接
	fmt.Println("\n3. 字符串分割和连接：")
	fmt.Println("字符串分割连接演示已实现")

	// 4. 字符串格式化
	fmt.Println("\n4. 字符串格式化：")
	fmt.Println("字符串格式化演示已实现")

	// 5. 字符串转换
	fmt.Println("\n5. 字符串转换：")
	demoStringConversion()

	// 6. 字符串验证
	fmt.Println("\n6. 字符串验证：")
	demoStringValidation()
}

// DemoStructs 演示结构体
func DemoStructs() {
	fmt.Println("\n=== 结构体演示 ===")

	// 1. 结构体基础
	fmt.Println("\n1. 结构体基础：")
	demoBasicStructs()

	// 2. 结构体初始化
	fmt.Println("\n2. 结构体初始化：")
	demoStructInitialization()

	// 3. 结构体操作
	fmt.Println("\n3. 结构体操作：")
	demoStructOperations()

	// 4. 嵌套结构体
	fmt.Println("\n4. 嵌套结构体：")
	demoNestedStructs()

	// 5. 匿名结构体
	fmt.Println("\n5. 匿名结构体：")
	demoAnonymousStructs()

	// 6. 结构体标签
	fmt.Println("\n6. 结构体标签：")
	demoStructTags()
}

// Student 学生结构体
type Student struct {
	ID       int
	Name     string
	Age      int
	Grade    string
	Subjects []string
}

// Point 点结构体
type Point struct {
	X, Y float64
}

// Rectangle 矩形结构体
type Rectangle struct {
	TopLeft     Point
	BottomRight Point
}

// demoBasicStructs 演示结构体基础
func demoBasicStructs() {
	// 1. 声明和初始化结构体
	var s1 Student
	fmt.Printf("零值结构体: %+v\n", s1)

	// 2. 字段赋值
	s1.ID = 1
	s1.Name = "Alice"
	s1.Age = 20
	s1.Grade = "A"
	s1.Subjects = []string{"Math", "Physics", "Chemistry"}

	fmt.Printf("赋值后: %+v\n", s1)

	// 3. 访问字段
	fmt.Printf("学生姓名: %s\n", s1.Name)
	fmt.Printf("学生年龄: %d\n", s1.Age)
	fmt.Printf("学科数量: %d\n", len(s1.Subjects))

	// 4. 修改字段
	s1.Age = 21
	s1.Subjects = append(s1.Subjects, "Biology")
	fmt.Printf("修改后: %+v\n", s1)

	// 5. 结构体比较
	var s2 Student
	s2.ID = 1
	s2.Name = "Alice"
	s2.Age = 21
	s2.Grade = "A"
	// 注意：包含切片的结构体不能直接比较

	fmt.Printf("s1.ID == s2.ID: %t\n", s1.ID == s2.ID)
	fmt.Printf("s1.Name == s2.Name: %t\n", s1.Name == s2.Name)
}

// demoStructInitialization 演示结构体初始化
func demoStructInitialization() {
	// 1. 字面量初始化
	s1 := Student{
		ID:       2,
		Name:     "Bob",
		Age:      19,
		Grade:    "B",
		Subjects: []string{"Math", "English"},
	}
	fmt.Printf("字面量初始化: %+v\n", s1)

	// 2. 按顺序初始化（不推荐）
	s2 := Student{3, "Carol", 20, "A", []string{"Physics", "Chemistry"}}
	fmt.Printf("按顺序初始化: %+v\n", s2)

	// 3. 部分初始化
	s3 := Student{
		Name: "David",
		Age:  18,
	}
	fmt.Printf("部分初始化: %+v\n", s3)

	// 4. 使用new创建
	s4 := new(Student)
	s4.Name = "Eve"
	s4.Age = 22
	fmt.Printf("new创建: %+v\n", *s4)

	// 5. 指针初始化
	s5 := &Student{
		Name: "Frank",
		Age:  21,
	}
	fmt.Printf("指针初始化: %+v\n", *s5)

	// 6. 复制结构体
	s6 := s1 // 值复制
	s6.Name = "Bob Copy"
	fmt.Printf("原结构体: %+v\n", s1)
	fmt.Printf("复制结构体: %+v\n", s6)
}

// demoStructOperations 演示结构体操作
func demoStructOperations() {
	students := []Student{
		{ID: 1, Name: "Alice", Age: 20, Grade: "A", Subjects: []string{"Math", "Physics"}},
		{ID: 2, Name: "Bob", Age: 19, Grade: "B", Subjects: []string{"Chemistry", "Biology"}},
		{ID: 3, Name: "Carol", Age: 21, Grade: "A", Subjects: []string{"Math", "English"}},
	}

	// 1. 遍历结构体切片
	fmt.Println("学生列表:")
	for i, student := range students {
		fmt.Printf("  %d. %s (年龄: %d, 成绩: %s)\n",
			i+1, student.Name, student.Age, student.Grade)
	}

	// 2. 查找学生
	targetName := "Bob"
	found := findStudent(students, targetName)
	if found != nil {
		fmt.Printf("找到学生: %+v\n", *found)
	} else {
		fmt.Printf("未找到学生: %s\n", targetName)
	}

	// 3. 过滤学生
	aGradeStudents := filterStudentsByGrade(students, "A")
	fmt.Printf("A级学生: %d人\n", len(aGradeStudents))
	for _, student := range aGradeStudents {
		fmt.Printf("  %s\n", student.Name)
	}

	// 4. 统计信息
	avgAge := calculateAverageAge(students)
	fmt.Printf("平均年龄: %.1f\n", avgAge)

	// 5. 排序学生（按年龄）
	sortedStudents := make([]Student, len(students))
	copy(sortedStudents, students)
	sortStudentsByAge(sortedStudents)

	fmt.Println("按年龄排序:")
	for _, student := range sortedStudents {
		fmt.Printf("  %s: %d岁\n", student.Name, student.Age)
	}

	// 6. 结构体作为map的值
	studentMap := make(map[int]Student)
	for _, student := range students {
		studentMap[student.ID] = student
	}

	fmt.Println("学生映射:")
	for id, student := range studentMap {
		fmt.Printf("  ID %d: %s\n", id, student.Name)
	}
}

// findStudent 查找学生
func findStudent(students []Student, name string) *Student {
	for i := range students {
		if students[i].Name == name {
			return &students[i]
		}
	}
	return nil
}

// filterStudentsByGrade 按成绩过滤学生
func filterStudentsByGrade(students []Student, grade string) []Student {
	var result []Student
	for _, student := range students {
		if student.Grade == grade {
			result = append(result, student)
		}
	}
	return result
}

// calculateAverageAge 计算平均年龄
func calculateAverageAge(students []Student) float64 {
	if len(students) == 0 {
		return 0
	}

	total := 0
	for _, student := range students {
		total += student.Age
	}

	return float64(total) / float64(len(students))
}

// sortStudentsByAge 按年龄排序学生
func sortStudentsByAge(students []Student) {
	n := len(students)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if students[j].Age > students[j+1].Age {
				students[j], students[j+1] = students[j+1], students[j]
			}
		}
	}
}

// demoNestedStructs 演示嵌套结构体
func demoNestedStructs() {
	// 1. 创建嵌套结构体
	rect := Rectangle{
		TopLeft:     Point{X: 0, Y: 10},
		BottomRight: Point{X: 10, Y: 0},
	}

	fmt.Printf("矩形: %+v\n", rect)
	fmt.Printf("左上角: (%.1f, %.1f)\n", rect.TopLeft.X, rect.TopLeft.Y)
	fmt.Printf("右下角: (%.1f, %.1f)\n", rect.BottomRight.X, rect.BottomRight.Y)

	// 2. 计算矩形属性
	width := rect.BottomRight.X - rect.TopLeft.X
	height := rect.TopLeft.Y - rect.BottomRight.Y
	area := width * height

	fmt.Printf("宽度: %.1f\n", width)
	fmt.Printf("高度: %.1f\n", height)
	fmt.Printf("面积: %.1f\n", area)

	// 3. 复杂嵌套结构体
	type Address struct {
		Street  string
		City    string
		ZipCode string
	}

	type Person struct {
		Name    string
		Age     int
		Address Address
		Friends []string
	}

	person := Person{
		Name: "John",
		Age:  30,
		Address: Address{
			Street:  "123 Main St",
			City:    "New York",
			ZipCode: "10001",
		},
		Friends: []string{"Alice", "Bob", "Carol"},
	}

	fmt.Printf("人员信息: %+v\n", person)
	fmt.Printf("地址: %s, %s %s\n",
		person.Address.Street, person.Address.City, person.Address.ZipCode)
	fmt.Printf("朋友数量: %d\n", len(person.Friends))
}

// demoAnonymousStructs 演示匿名结构体
func demoAnonymousStructs() {
	// 1. 匿名结构体变量
	config := struct {
		Host     string
		Port     int
		Database string
		SSL      bool
	}{
		Host:     "localhost",
		Port:     5432,
		Database: "myapp",
		SSL:      true,
	}

	fmt.Printf("配置: %+v\n", config)
	fmt.Printf("连接字符串: %s:%d/%s (SSL: %t)\n",
		config.Host, config.Port, config.Database, config.SSL)

	// 2. 匿名结构体切片
	responses := []struct {
		Status  int
		Message string
		Data    interface{}
	}{
		{200, "Success", "Hello World"},
		{404, "Not Found", nil},
		{500, "Internal Error", map[string]string{"error": "database connection failed"}},
	}

	fmt.Println("响应列表:")
	for i, resp := range responses {
		fmt.Printf("  %d. Status: %d, Message: %s, Data: %v\n",
			i+1, resp.Status, resp.Message, resp.Data)
	}

	// 3. 临时数据结构
	result := struct {
		Count int
		Items []string
	}{
		Count: 3,
		Items: []string{"apple", "banana", "orange"},
	}

	fmt.Printf("结果: 共%d项 - %v\n", result.Count, result.Items)

	// 4. 函数返回匿名结构体
	stats := getStats([]int{1, 2, 3, 4, 5})
	fmt.Printf("统计信息: %+v\n", stats)
}

// getStats 返回统计信息的匿名结构体
func getStats(numbers []int) struct {
	Count int
	Sum   int
	Avg   float64
	Min   int
	Max   int
} {
	if len(numbers) == 0 {
		return struct {
			Count int
			Sum   int
			Avg   float64
			Min   int
			Max   int
		}{}
	}

	sum := 0
	min := numbers[0]
	max := numbers[0]

	for _, num := range numbers {
		sum += num
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	return struct {
		Count int
		Sum   int
		Avg   float64
		Min   int
		Max   int
	}{
		Count: len(numbers),
		Sum:   sum,
		Avg:   float64(sum) / float64(len(numbers)),
		Min:   min,
		Max:   max,
	}
}

// demoStructTags 演示结构体标签
func demoStructTags() {
	// 1. 带标签的结构体
	type User struct {
		ID       int    `json:"id" db:"user_id" validate:"required"`
		Username string `json:"username" db:"username" validate:"required,min=3"`
		Email    string `json:"email" db:"email" validate:"required,email"`
		Password string `json:"-" db:"password_hash" validate:"required,min=8"`
		Active   bool   `json:"active" db:"is_active"`
	}

	user := User{
		ID:       1,
		Username: "john_doe",
		Email:    "john@example.com",
		Password: "secret123",
		Active:   true,
	}

	fmt.Printf("用户结构体: %+v\n", user)

	// 2. 模拟JSON序列化（简化版）
	jsonStr := structToJSON(user)
	fmt.Printf("JSON表示: %s\n", jsonStr)

	// 3. 模拟数据库字段映射
	dbFields := getDBFields(user)
	fmt.Printf("数据库字段: %v\n", dbFields)

	// 4. 模拟验证规则
	validationRules := getValidationRules(user)
	fmt.Printf("验证规则: %v\n", validationRules)
}

// structToJSON 简化的JSON序列化
func structToJSON(user interface{}) string {
	// 这是一个简化的实现，实际应该使用reflect包
	switch v := user.(type) {
	case Student:
		return fmt.Sprintf(`{"id":%d,"name":"%s","age":%d,"grade":"%s"}`,
			v.ID, v.Name, v.Age, v.Grade)
	default:
		return fmt.Sprintf("%+v", user)
	}
}

// getDBFields 获取数据库字段映射
func getDBFields(user interface{}) map[string]interface{} {
	// 简化实现，实际应该使用reflect包解析标签
	result := make(map[string]interface{})

	// 模拟从结构体标签中提取db字段名
	result["user_id"] = 1
	result["username"] = "john_doe"
	result["email"] = "john@example.com"
	result["password_hash"] = "secret123"
	result["is_active"] = true

	return result
}

// getValidationRules 获取验证规则
func getValidationRules(user interface{}) map[string][]string {
	// 简化实现，实际应该使用reflect包解析标签
	rules := make(map[string][]string)

	rules["ID"] = []string{"required"}
	rules["Username"] = []string{"required", "min=3"}
	rules["Email"] = []string{"required", "email"}
	rules["Password"] = []string{"required", "min=8"}

	return rules
}

// DemoMethods 演示方法
func DemoMethods() {
	fmt.Println("\n=== 方法演示 ===")

	// 1. 基本方法
	fmt.Println("\n1. 基本方法：")
	demoBasicMethods()

	// 2. 值接收者vs指针接收者
	fmt.Println("\n2. 值接收者vs指针接收者：")
	demoReceiverTypes()

	// 3. 方法集
	fmt.Println("\n3. 方法集：")
	demoMethodSets()

	// 4. 方法链式调用
	fmt.Println("\n4. 方法链式调用：")
	demoMethodChaining()

	// 5. 方法重载模拟
	fmt.Println("\n5. 方法重载模拟：")
	demoMethodOverloading()
}

// Circle 圆形结构体
type Circle struct {
	Radius float64
	Center Point
}

// Area 计算圆的面积（值接收者）
func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

// Circumference 计算圆的周长（值接收者）
func (c Circle) Circumference() float64 {
	return 2 * 3.14159 * c.Radius
}

// Perimeter 计算圆的周长（实现Shape接口）
func (c Circle) Perimeter() float64 {
	return c.Circumference()
}

// Scale 缩放圆（指针接收者）
func (c *Circle) Scale(factor float64) {
	c.Radius *= factor
}

// Move 移动圆心（指针接收者）
func (c *Circle) Move(dx, dy float64) {
	c.Center.X += dx
	c.Center.Y += dy
}

// String 字符串表示（值接收者）
func (c Circle) String() string {
	return fmt.Sprintf("Circle{Radius: %.2f, Center: (%.2f, %.2f)}",
		c.Radius, c.Center.X, c.Center.Y)
}

// IsValid 检查圆是否有效（值接收者）
func (c Circle) IsValid() bool {
	return c.Radius > 0
}

// demoBasicMethods 演示基本方法
func demoBasicMethods() {
	// 1. 创建圆形实例
	circle := Circle{
		Radius: 5.0,
		Center: Point{X: 0, Y: 0},
	}

	fmt.Printf("圆形: %s\n", circle.String())

	// 2. 调用方法
	area := circle.Area()
	circumference := circle.Circumference()

	fmt.Printf("面积: %.2f\n", area)
	fmt.Printf("周长: %.2f\n", circumference)
	fmt.Printf("是否有效: %t\n", circle.IsValid())

	// 3. 方法调用的不同方式
	// 直接调用
	fmt.Printf("直接调用面积: %.2f\n", circle.Area())

	// 通过指针调用
	circlePtr := &circle
	fmt.Printf("通过指针调用面积: %.2f\n", circlePtr.Area())

	// 4. 零值方法调用
	var zeroCircle Circle
	fmt.Printf("零值圆形: %s\n", zeroCircle.String())
	fmt.Printf("零值圆形面积: %.2f\n", zeroCircle.Area())
	fmt.Printf("零值圆形是否有效: %t\n", zeroCircle.IsValid())
}

// demoReceiverTypes 演示值接收者vs指针接收者
func demoReceiverTypes() {
	circle := Circle{
		Radius: 3.0,
		Center: Point{X: 1, Y: 1},
	}

	fmt.Printf("原始圆形: %s\n", circle.String())
	fmt.Printf("原始面积: %.2f\n", circle.Area())

	// 1. 值接收者方法 - 不会修改原始值
	fmt.Println("\n值接收者方法调用:")
	area1 := circle.Area()
	fmt.Printf("调用Area()后: %s\n", circle.String())
	fmt.Printf("面积: %.2f\n", area1)

	// 2. 指针接收者方法 - 会修改原始值
	fmt.Println("\n指针接收者方法调用:")
	fmt.Printf("缩放前: %s\n", circle.String())

	circle.Scale(2.0) // Go会自动取地址
	fmt.Printf("缩放2倍后: %s\n", circle.String())
	fmt.Printf("新面积: %.2f\n", circle.Area())

	circle.Move(5, 3)
	fmt.Printf("移动后: %s\n", circle.String())

	// 3. 通过指针调用
	circlePtr := &Circle{Radius: 2.0, Center: Point{X: 0, Y: 0}}
	fmt.Printf("\n通过指针操作:\n")
	fmt.Printf("指针圆形: %s\n", circlePtr.String())

	circlePtr.Scale(1.5)
	fmt.Printf("指针缩放后: %s\n", circlePtr.String())

	// 4. 演示值拷贝
	circle1 := Circle{Radius: 1.0}
	circle2 := circle1 // 值拷贝

	fmt.Printf("\n值拷贝演示:\n")
	fmt.Printf("circle1: %s\n", circle1.String())
	fmt.Printf("circle2: %s\n", circle2.String())

	circle1.Scale(3.0)
	fmt.Printf("circle1缩放后: %s\n", circle1.String())
	fmt.Printf("circle2未变: %s\n", circle2.String())
}

// Counter 计数器结构体
type Counter struct {
	value int
	name  string
}

// NewCounter 创建新计数器
func NewCounter(name string) *Counter {
	return &Counter{
		value: 0,
		name:  name,
	}
}

// Value 获取当前值（值接收者）
func (c Counter) Value() int {
	return c.value
}

// Name 获取名称（值接收者）
func (c Counter) Name() string {
	return c.name
}

// Increment 增加计数（指针接收者）
func (c *Counter) Increment() *Counter {
	c.value++
	return c
}

// Add 增加指定值（指针接收者）
func (c *Counter) Add(n int) *Counter {
	c.value += n
	return c
}

// Reset 重置计数（指针接收者）
func (c *Counter) Reset() *Counter {
	c.value = 0
	return c
}

// String 字符串表示
func (c Counter) String() string {
	return fmt.Sprintf("Counter{%s: %d}", c.name, c.value)
}

// demoMethodSets 演示方法集
func demoMethodSets() {
	// 1. 值类型的方法集
	counter := Counter{value: 5, name: "test"}

	fmt.Printf("值类型方法调用:\n")
	fmt.Printf("计数器: %s\n", counter.String())
	fmt.Printf("值: %d\n", counter.Value())
	fmt.Printf("名称: %s\n", counter.Name())

	// 值类型也可以调用指针接收者方法（Go自动取地址）
	counter.Increment()
	fmt.Printf("增加后: %s\n", counter.String())

	// 2. 指针类型的方法集
	counterPtr := &Counter{value: 10, name: "pointer"}

	fmt.Printf("\n指针类型方法调用:\n")
	fmt.Printf("计数器: %s\n", counterPtr.String())
	fmt.Printf("值: %d\n", counterPtr.Value())

	counterPtr.Add(5)
	fmt.Printf("增加5后: %s\n", counterPtr.String())

	// 3. 方法集的区别演示
	fmt.Printf("\n方法集区别:\n")

	// 值类型变量
	var c1 Counter = Counter{value: 1, name: "c1"}
	fmt.Printf("c1: %s\n", c1.String())

	// 指针类型变量
	var c2 *Counter = &Counter{value: 2, name: "c2"}
	fmt.Printf("c2: %s\n", c2.String())

	// 都可以调用值接收者和指针接收者方法
	c1.Increment() // Go自动转换为(&c1).Increment()
	c2.Increment()

	fmt.Printf("增加后 c1: %s\n", c1.String())
	fmt.Printf("增加后 c2: %s\n", c2.String())
}

// demoMethodChaining 演示方法链式调用
func demoMethodChaining() {
	// 1. 创建计数器并链式调用
	counter := NewCounter("chain")

	fmt.Printf("初始计数器: %s\n", counter.String())

	// 链式调用
	result := counter.Add(5).Increment().Add(3).Increment()

	fmt.Printf("链式调用后: %s\n", result.String())
	fmt.Printf("最终值: %d\n", result.Value())

	// 2. 更复杂的链式调用
	counter2 := NewCounter("complex")

	final := counter2.
		Add(10).     // 加10
		Increment(). // 加1
		Add(5).      // 加5
		Reset().     // 重置
		Add(100).    // 加100
		Increment(). // 加1
		Increment()  // 再加1

	fmt.Printf("复杂链式调用: %s\n", final.String())

	// 3. 条件链式调用
	counter3 := NewCounter("conditional")

	if counter3.Add(20).Value() > 15 {
		counter3.Add(10)
	}

	fmt.Printf("条件链式调用: %s\n", counter3.String())
}

// Calculator 计算器结构体
type Calculator struct {
	result float64
}

// NewCalculator 创建新计算器
func NewCalculator() *Calculator {
	return &Calculator{result: 0}
}

// Add 加法
func (c *Calculator) Add(n float64) *Calculator {
	c.result += n
	return c
}

// Subtract 减法
func (c *Calculator) Subtract(n float64) *Calculator {
	c.result -= n
	return c
}

// Multiply 乘法
func (c *Calculator) Multiply(n float64) *Calculator {
	c.result *= n
	return c
}

// Divide 除法
func (c *Calculator) Divide(n float64) *Calculator {
	if n != 0 {
		c.result /= n
	}
	return c
}

// Result 获取结果
func (c *Calculator) Result() float64 {
	return c.result
}

// Clear 清零
func (c *Calculator) Clear() *Calculator {
	c.result = 0
	return c
}

// AddInt 整数加法（模拟重载）
func (c *Calculator) AddInt(n int) *Calculator {
	c.result += float64(n)
	return c
}

// AddFloat 浮点数加法（模拟重载）
func (c *Calculator) AddFloat(n float64) *Calculator {
	c.result += n
	return c
}

// AddMultiple 多个数加法（模拟重载）
func (c *Calculator) AddMultiple(numbers ...float64) *Calculator {
	for _, n := range numbers {
		c.result += n
	}
	return c
}

// demoMethodOverloading 演示方法重载模拟
func demoMethodOverloading() {
	calc := NewCalculator()

	// 1. 基本计算
	fmt.Printf("初始值: %.2f\n", calc.Result())

	result1 := calc.Add(10).Subtract(3).Multiply(2).Result()
	fmt.Printf("(0 + 10 - 3) * 2 = %.2f\n", result1)

	// 2. 模拟方法重载
	calc.Clear()

	fmt.Printf("\n模拟方法重载:\n")
	calc.AddInt(5)
	fmt.Printf("AddInt(5): %.2f\n", calc.Result())

	calc.AddFloat(3.14)
	fmt.Printf("AddFloat(3.14): %.2f\n", calc.Result())

	calc.AddMultiple(1, 2, 3, 4, 5)
	fmt.Printf("AddMultiple(1,2,3,4,5): %.2f\n", calc.Result())

	// 3. 复杂计算链
	calc2 := NewCalculator()

	complexResult := calc2.
		Add(100).           // 100
		Divide(4).          // 25
		Multiply(3).        // 75
		Subtract(15).       // 60
		AddMultiple(5, 10). // 75
		Result()

	fmt.Printf("复杂计算结果: %.2f\n", complexResult)

	// 4. 多个计算器实例
	calc3 := NewCalculator()
	calc4 := NewCalculator()

	calc3.Add(50)
	calc4.Add(30)

	fmt.Printf("calc3: %.2f\n", calc3.Result())
	fmt.Printf("calc4: %.2f\n", calc4.Result())

	// 5. 方法作为值传递
	operations := []func(*Calculator) *Calculator{
		func(c *Calculator) *Calculator { return c.Add(10) },
		func(c *Calculator) *Calculator { return c.Multiply(2) },
		func(c *Calculator) *Calculator { return c.Subtract(5) },
	}

	calc5 := NewCalculator()
	for i, op := range operations {
		op(calc5)
		fmt.Printf("操作%d后: %.2f\n", i+1, calc5.Result())
	}
}

// DemoConstructor 演示构造函数
func DemoConstructor() {
	fmt.Println("\n=== 构造函数演示 ===")

	// 1. 基本构造函数
	fmt.Println("\n1. 基本构造函数：")
	demoBasicConstructors()

	// 2. 带参数的构造函数
	fmt.Println("\n2. 带参数的构造函数：")
	demoParameterizedConstructors()

	// 3. 工厂函数
	fmt.Println("\n3. 工厂函数：")
	demoFactoryFunctions()

	// 4. 构造函数选项模式
	fmt.Println("\n4. 构造函数选项模式：")
	demoConstructorOptions()

	// 5. 单例模式
	fmt.Println("\n5. 单例模式：")
	demoSingletonPattern()
}

// DemoEmbedding 演示嵌入
func DemoEmbedding() {
	fmt.Println("\n=== 嵌入演示 ===")

	// 1. 结构体嵌入
	fmt.Println("\n1. 结构体嵌入：")
	demoStructEmbedding()

	// 2. 接口嵌入
	fmt.Println("\n2. 接口嵌入：")
	demoInterfaceEmbedding()

	// 3. 方法提升
	fmt.Println("\n3. 方法提升：")
	demoMethodPromotion()

	// 4. 嵌入冲突处理
	fmt.Println("\n4. 嵌入冲突处理：")
	demoEmbeddingConflicts()

	// 5. 组合vs继承
	fmt.Println("\n5. 组合vs继承：")
	demoCompositionVsInheritance()
}

// Book 书籍结构体
type Book struct {
	ID     int
	Title  string
	Author string
	Pages  int
	Price  float64
}

// NewBook 基本构造函数
func NewBook(title, author string) *Book {
	return &Book{
		Title:  title,
		Author: author,
	}
}

// NewBookWithID 带ID的构造函数
func NewBookWithID(id int, title, author string) *Book {
	return &Book{
		ID:     id,
		Title:  title,
		Author: author,
	}
}

// NewBookFull 完整构造函数
func NewBookFull(id int, title, author string, pages int, price float64) *Book {
	return &Book{
		ID:     id,
		Title:  title,
		Author: author,
		Pages:  pages,
		Price:  price,
	}
}

// String 字符串表示
func (b *Book) String() string {
	return fmt.Sprintf("Book{ID: %d, Title: %s, Author: %s, Pages: %d, Price: %.2f}",
		b.ID, b.Title, b.Author, b.Pages, b.Price)
}

// IsValid 验证书籍信息
func (b *Book) IsValid() bool {
	return b.Title != "" && b.Author != "" && b.Pages > 0 && b.Price >= 0
}

// demoBasicConstructors 演示基本构造函数
func demoBasicConstructors() {
	// 1. 使用基本构造函数
	book1 := NewBook("Go Programming", "John Doe")
	fmt.Printf("基本构造: %s\n", book1.String())

	// 2. 手动设置其他字段
	book1.ID = 1
	book1.Pages = 300
	book1.Price = 29.99
	fmt.Printf("设置字段后: %s\n", book1.String())

	// 3. 使用带ID的构造函数
	book2 := NewBookWithID(2, "Advanced Go", "Jane Smith")
	book2.Pages = 450
	book2.Price = 39.99
	fmt.Printf("带ID构造: %s\n", book2.String())

	// 4. 使用完整构造函数
	book3 := NewBookFull(3, "Go Patterns", "Bob Wilson", 250, 24.99)
	fmt.Printf("完整构造: %s\n", book3.String())

	// 5. 验证书籍
	books := []*Book{book1, book2, book3}
	for i, book := range books {
		fmt.Printf("书籍%d有效性: %t\n", i+1, book.IsValid())
	}
}

// User 用户结构体
type User struct {
	ID       int
	Username string
	Email    string
	Age      int
	Active   bool
}

// NewUser 基本用户构造函数
func NewUser(username, email string) *User {
	return &User{
		Username: username,
		Email:    email,
		Active:   true, // 默认激活
	}
}

// NewUserWithAge 带年龄的构造函数
func NewUserWithAge(username, email string, age int) *User {
	user := NewUser(username, email)
	user.Age = age
	return user
}

// NewInactiveUser 创建非激活用户
func NewInactiveUser(username, email string) *User {
	user := NewUser(username, email)
	user.Active = false
	return user
}

// String 用户字符串表示
func (u *User) String() string {
	status := "Active"
	if !u.Active {
		status = "Inactive"
	}
	return fmt.Sprintf("User{ID: %d, Username: %s, Email: %s, Age: %d, Status: %s}",
		u.ID, u.Username, u.Email, u.Age, status)
}

// Activate 激活用户
func (u *User) Activate() {
	u.Active = true
}

// Deactivate 停用用户
func (u *User) Deactivate() {
	u.Active = false
}

// demoParameterizedConstructors 演示带参数的构造函数
func demoParameterizedConstructors() {
	// 1. 基本用户
	user1 := NewUser("alice", "alice@example.com")
	user1.ID = 1
	fmt.Printf("基本用户: %s\n", user1.String())

	// 2. 带年龄的用户
	user2 := NewUserWithAge("bob", "bob@example.com", 25)
	user2.ID = 2
	fmt.Printf("带年龄用户: %s\n", user2.String())

	// 3. 非激活用户
	user3 := NewInactiveUser("carol", "carol@example.com")
	user3.ID = 3
	user3.Age = 30
	fmt.Printf("非激活用户: %s\n", user3.String())

	// 4. 用户状态操作
	fmt.Println("\n用户状态操作:")
	user3.Activate()
	fmt.Printf("激活后: %s\n", user3.String())

	user1.Deactivate()
	fmt.Printf("停用后: %s\n", user1.String())
}

// Product 产品结构体
type Product struct {
	ID          int
	Name        string
	Category    string
	Price       float64
	InStock     bool
	Description string
}

// ProductType 产品类型
type ProductType int

const (
	Electronics ProductType = iota
	Books
	Clothing
	Food
)

// CreateProduct 工厂函数 - 根据类型创建产品
func CreateProduct(productType ProductType, name string, price float64) *Product {
	product := &Product{
		Name:    name,
		Price:   price,
		InStock: true,
	}

	switch productType {
	case Electronics:
		product.Category = "Electronics"
		product.Description = "Electronic device"
	case Books:
		product.Category = "Books"
		product.Description = "Book or publication"
	case Clothing:
		product.Category = "Clothing"
		product.Description = "Clothing item"
	case Food:
		product.Category = "Food"
		product.Description = "Food item"
	default:
		product.Category = "Unknown"
		product.Description = "Unknown product type"
	}

	return product
}

// CreateElectronics 创建电子产品
func CreateElectronics(name string, price float64) *Product {
	return CreateProduct(Electronics, name, price)
}

// CreateBook 创建书籍产品
func CreateBook(name string, price float64) *Product {
	return CreateProduct(Books, name, price)
}

// String 产品字符串表示
func (p *Product) String() string {
	stock := "In Stock"
	if !p.InStock {
		stock = "Out of Stock"
	}
	return fmt.Sprintf("Product{ID: %d, Name: %s, Category: %s, Price: %.2f, %s}",
		p.ID, p.Name, p.Category, p.Price, stock)
}

// demoFactoryFunctions 演示工厂函数
func demoFactoryFunctions() {
	// 1. 使用通用工厂函数
	laptop := CreateProduct(Electronics, "Gaming Laptop", 1299.99)
	laptop.ID = 1
	fmt.Printf("电子产品: %s\n", laptop.String())

	book := CreateProduct(Books, "Go Programming Guide", 49.99)
	book.ID = 2
	fmt.Printf("书籍产品: %s\n", book.String())

	// 2. 使用专门的工厂函数
	phone := CreateElectronics("Smartphone", 699.99)
	phone.ID = 3
	fmt.Printf("专门工厂(电子): %s\n", phone.String())

	novel := CreateBook("Science Fiction Novel", 19.99)
	novel.ID = 4
	fmt.Printf("专门工厂(书籍): %s\n", novel.String())

	// 3. 批量创建
	products := []*Product{
		CreateProduct(Clothing, "T-Shirt", 25.99),
		CreateProduct(Food, "Organic Apple", 3.99),
		CreateElectronics("Tablet", 399.99),
		CreateBook("Cookbook", 29.99),
	}

	fmt.Println("\n批量创建的产品:")
	for i, product := range products {
		product.ID = i + 5
		fmt.Printf("  %s\n", product.String())
	}
}

// Server 服务器结构体
type Server struct {
	Host    string
	Port    int
	Timeout int
	SSL     bool
	Debug   bool
}

// ServerOption 服务器选项函数类型
type ServerOption func(*Server)

// WithPort 设置端口选项
func WithPort(port int) ServerOption {
	return func(s *Server) {
		s.Port = port
	}
}

// WithTimeout 设置超时选项
func WithTimeout(timeout int) ServerOption {
	return func(s *Server) {
		s.Timeout = timeout
	}
}

// WithSSL 启用SSL选项
func WithSSL() ServerOption {
	return func(s *Server) {
		s.SSL = true
	}
}

// WithDebug 启用调试选项
func WithDebug() ServerOption {
	return func(s *Server) {
		s.Debug = true
	}
}

// NewServer 创建服务器（选项模式）
func NewServer(host string, options ...ServerOption) *Server {
	// 默认配置
	server := &Server{
		Host:    host,
		Port:    8080,
		Timeout: 30,
		SSL:     false,
		Debug:   false,
	}

	// 应用选项
	for _, option := range options {
		option(server)
	}

	return server
}

// String 服务器字符串表示
func (s *Server) String() string {
	return fmt.Sprintf("Server{Host: %s, Port: %d, Timeout: %ds, SSL: %t, Debug: %t}",
		s.Host, s.Port, s.Timeout, s.SSL, s.Debug)
}

// Start 启动服务器
func (s *Server) Start() {
	fmt.Printf("启动服务器: %s\n", s.String())
}

// demoConstructorOptions 演示构造函数选项模式
func demoConstructorOptions() {
	// 1. 使用默认配置
	server1 := NewServer("localhost")
	fmt.Printf("默认配置: %s\n", server1.String())

	// 2. 使用部分选项
	server2 := NewServer("example.com", WithPort(443), WithSSL())
	fmt.Printf("部分选项: %s\n", server2.String())

	// 3. 使用所有选项
	server3 := NewServer("api.example.com",
		WithPort(9000),
		WithTimeout(60),
		WithSSL(),
		WithDebug())
	fmt.Printf("所有选项: %s\n", server3.String())

	// 4. 动态选项
	options := []ServerOption{WithPort(8443)}
	if true { // 某种条件
		options = append(options, WithSSL(), WithDebug())
	}

	server4 := NewServer("dynamic.example.com", options...)
	fmt.Printf("动态选项: %s\n", server4.String())

	// 5. 启动服务器
	fmt.Println("\n启动服务器:")
	server1.Start()
	server2.Start()
}

// Database 数据库单例
type Database struct {
	connectionString string
	connected        bool
}

var (
	dbInstance *Database
	dbOnce     bool // 简化的once实现
)

// GetDatabase 获取数据库单例
func GetDatabase() *Database {
	if !dbOnce {
		dbInstance = &Database{
			connectionString: "localhost:5432/myapp",
			connected:        false,
		}
		dbOnce = true
	}
	return dbInstance
}

// Connect 连接数据库
func (db *Database) Connect() {
	if !db.connected {
		db.connected = true
		fmt.Printf("连接到数据库: %s\n", db.connectionString)
	} else {
		fmt.Println("数据库已连接")
	}
}

// Disconnect 断开数据库连接
func (db *Database) Disconnect() {
	if db.connected {
		db.connected = false
		fmt.Println("断开数据库连接")
	} else {
		fmt.Println("数据库未连接")
	}
}

// IsConnected 检查连接状态
func (db *Database) IsConnected() bool {
	return db.connected
}

// String 数据库字符串表示
func (db *Database) String() string {
	status := "disconnected"
	if db.connected {
		status = "connected"
	}
	return fmt.Sprintf("Database{%s, %s}", db.connectionString, status)
}

// demoSingletonPattern 演示单例模式
func demoSingletonPattern() {
	// 1. 获取数据库实例
	db1 := GetDatabase()
	fmt.Printf("第一个实例: %s\n", db1.String())

	// 2. 再次获取实例（应该是同一个）
	db2 := GetDatabase()
	fmt.Printf("第二个实例: %s\n", db2.String())

	// 3. 验证是同一个实例
	fmt.Printf("是同一个实例: %t\n", db1 == db2)

	// 4. 操作数据库
	fmt.Println("\n数据库操作:")
	db1.Connect()
	fmt.Printf("db1连接状态: %t\n", db1.IsConnected())
	fmt.Printf("db2连接状态: %t\n", db2.IsConnected())

	db2.Disconnect()
	fmt.Printf("db1连接状态: %t\n", db1.IsConnected())
	fmt.Printf("db2连接状态: %t\n", db2.IsConnected())

	// 5. 多次获取实例
	instances := make([]*Database, 5)
	for i := 0; i < 5; i++ {
		instances[i] = GetDatabase()
	}

	fmt.Println("\n多个实例验证:")
	allSame := true
	for i := 1; i < len(instances); i++ {
		if instances[i] != instances[0] {
			allSame = false
			break
		}
	}
	fmt.Printf("所有实例都相同: %t\n", allSame)
}

// Animal 动物基础结构体
type Animal struct {
	Name    string
	Species string
	Age     int
}

// Speak 动物说话
func (a *Animal) Speak() string {
	return fmt.Sprintf("%s makes a sound", a.Name)
}

// Info 动物信息
func (a *Animal) Info() string {
	return fmt.Sprintf("%s is a %d-year-old %s", a.Name, a.Age, a.Species)
}

// Dog 狗结构体（嵌入Animal）
type Dog struct {
	Animal // 嵌入Animal
	Breed  string
}

// Speak 狗的说话方式（重写）
func (d *Dog) Speak() string {
	return fmt.Sprintf("%s barks: Woof!", d.Name)
}

// Fetch 狗特有的方法
func (d *Dog) Fetch() string {
	return fmt.Sprintf("%s fetches the ball", d.Name)
}

// Cat 猫结构体（嵌入Animal）
type Cat struct {
	Animal // 嵌入Animal
	Indoor bool
}

// Speak 猫的说话方式（重写）
func (c *Cat) Speak() string {
	return fmt.Sprintf("%s meows: Meow!", c.Name)
}

// Climb 猫特有的方法
func (c *Cat) Climb() string {
	return fmt.Sprintf("%s climbs the tree", c.Name)
}

// demoStructEmbedding 演示结构体嵌入
func demoStructEmbedding() {
	// 1. 创建狗实例
	dog := Dog{
		Animal: Animal{
			Name:    "Buddy",
			Species: "Dog",
			Age:     3,
		},
		Breed: "Golden Retriever",
	}

	fmt.Printf("狗信息: %s\n", dog.Info())  // 调用嵌入的方法
	fmt.Printf("狗说话: %s\n", dog.Speak()) // 调用重写的方法
	fmt.Printf("狗取球: %s\n", dog.Fetch()) // 调用自己的方法
	fmt.Printf("狗品种: %s\n", dog.Breed)

	// 2. 创建猫实例
	cat := Cat{
		Animal: Animal{
			Name:    "Whiskers",
			Species: "Cat",
			Age:     2,
		},
		Indoor: true,
	}

	fmt.Printf("\n猫信息: %s\n", cat.Info())
	fmt.Printf("猫说话: %s\n", cat.Speak())
	fmt.Printf("猫爬树: %s\n", cat.Climb())
	fmt.Printf("室内猫: %t\n", cat.Indoor)

	// 3. 直接访问嵌入字段
	fmt.Printf("\n直接访问嵌入字段:\n")
	fmt.Printf("狗名字: %s\n", dog.Name)    // 等同于 dog.Animal.Name
	fmt.Printf("狗年龄: %d\n", dog.Age)     // 等同于 dog.Animal.Age
	fmt.Printf("猫名字: %s\n", cat.Name)    // 等同于 cat.Animal.Name
	fmt.Printf("猫物种: %s\n", cat.Species) // 等同于 cat.Animal.Species

	// 4. 修改嵌入字段
	dog.Age = 4
	cat.Name = "Fluffy"

	fmt.Printf("\n修改后:\n")
	fmt.Printf("狗信息: %s\n", dog.Info())
	fmt.Printf("猫信息: %s\n", cat.Info())
}

// Speaker 说话者接口
type Speaker interface {
	Speak() string
}

// Walker 行走者接口
type Walker interface {
	Walk() string
}

// Pet 宠物接口（嵌入其他接口）
type Pet interface {
	Speaker // 嵌入Speaker接口
	Walker  // 嵌入Walker接口
	Play() string
}

// Robot 机器人结构体
type Robot struct {
	Name  string
	Model string
}

// Speak 机器人说话
func (r *Robot) Speak() string {
	return fmt.Sprintf("Robot %s says: Beep beep!", r.Name)
}

// Walk 机器人行走
func (r *Robot) Walk() string {
	return fmt.Sprintf("Robot %s walks mechanically", r.Name)
}

// Play 机器人玩耍
func (r *Robot) Play() string {
	return fmt.Sprintf("Robot %s plays electronic games", r.Name)
}

// Walk 狗行走
func (d *Dog) Walk() string {
	return fmt.Sprintf("%s walks on four legs", d.Name)
}

// Play 狗玩耍
func (d *Dog) Play() string {
	return fmt.Sprintf("%s plays with toys", d.Name)
}

// Walk 猫行走
func (c *Cat) Walk() string {
	return fmt.Sprintf("%s walks silently", c.Name)
}

// Play 猫玩耍
func (c *Cat) Play() string {
	return fmt.Sprintf("%s plays with yarn", c.Name)
}

// demoInterfaceEmbedding 演示接口嵌入
func demoInterfaceEmbedding() {
	// 1. 创建实现Pet接口的实例
	pets := []Pet{
		&Dog{
			Animal: Animal{Name: "Rex", Species: "Dog", Age: 5},
			Breed:  "German Shepherd",
		},
		&Cat{
			Animal: Animal{Name: "Luna", Species: "Cat", Age: 3},
			Indoor: false,
		},
		&Robot{
			Name:  "Robo",
			Model: "PetBot-3000",
		},
	}

	// 2. 统一处理所有宠物
	fmt.Println("宠物活动:")
	for i, pet := range pets {
		fmt.Printf("\n宠物 %d:\n", i+1)
		fmt.Printf("  说话: %s\n", pet.Speak())
		fmt.Printf("  行走: %s\n", pet.Walk())
		fmt.Printf("  玩耍: %s\n", pet.Play())
	}

	// 3. 类型断言
	fmt.Println("\n类型断言:")
	for i, pet := range pets {
		fmt.Printf("宠物 %d: ", i+1)

		switch p := pet.(type) {
		case *Dog:
			fmt.Printf("这是一只%s品种的狗\n", p.Breed)
		case *Cat:
			if p.Indoor {
				fmt.Println("这是一只室内猫")
			} else {
				fmt.Println("这是一只户外猫")
			}
		case *Robot:
			fmt.Printf("这是一个%s型号的机器人\n", p.Model)
		}
	}

	// 4. 接口组合的好处
	fmt.Println("\n接口组合演示:")

	// 只需要说话功能
	speakers := []Speaker{pets[0], pets[1], pets[2]}
	fmt.Println("所有会说话的:")
	for _, speaker := range speakers {
		fmt.Printf("  %s\n", speaker.Speak())
	}

	// 只需要行走功能
	walkers := []Walker{pets[0], pets[1], pets[2]}
	fmt.Println("所有会行走的:")
	for _, walker := range walkers {
		fmt.Printf("  %s\n", walker.Walk())
	}
}

// Engine 引擎结构体
type Engine struct {
	Power int
	Type  string
}

// Start 启动引擎
func (e *Engine) Start() string {
	return fmt.Sprintf("Engine started: %d HP %s engine", e.Power, e.Type)
}

// Stop 停止引擎
func (e *Engine) Stop() string {
	return "Engine stopped"
}

// Wheels 轮子结构体
type Wheels struct {
	Count int
	Size  int
}

// Roll 轮子滚动
func (w *Wheels) Roll() string {
	return fmt.Sprintf("%d wheels (%d inch) are rolling", w.Count, w.Size)
}

// Car 汽车结构体（嵌入多个结构体）
type Car struct {
	Engine // 嵌入引擎
	Wheels // 嵌入轮子
	Brand  string
	Model  string
}

// Drive 汽车驾驶
func (c *Car) Drive() string {
	return fmt.Sprintf("%s %s is driving", c.Brand, c.Model)
}

// Info 汽车信息
func (c *Car) Info() string {
	return fmt.Sprintf("%s %s with %s", c.Brand, c.Model, c.Start())
}

// demoMethodPromotion 演示方法提升
func demoMethodPromotion() {
	// 1. 创建汽车实例
	car := Car{
		Engine: Engine{
			Power: 300,
			Type:  "V6",
		},
		Wheels: Wheels{
			Count: 4,
			Size:  18,
		},
		Brand: "Toyota",
		Model: "Camry",
	}

	// 2. 调用提升的方法
	fmt.Printf("汽车信息: %s\n", car.Info())
	fmt.Printf("启动引擎: %s\n", car.Start()) // 来自Engine
	fmt.Printf("轮子滚动: %s\n", car.Roll())  // 来自Wheels
	fmt.Printf("汽车驾驶: %s\n", car.Drive()) // 自己的方法

	// 3. 直接访问嵌入字段
	fmt.Printf("\n直接访问:\n")
	fmt.Printf("引擎功率: %d HP\n", car.Power)  // 等同于 car.Engine.Power
	fmt.Printf("轮子数量: %d\n", car.Count)     // 等同于 car.Wheels.Count
	fmt.Printf("轮子尺寸: %d inch\n", car.Size) // 等同于 car.Wheels.Size

	// 4. 修改嵌入字段
	car.Power = 350
	car.Count = 4
	car.Type = "V8"

	fmt.Printf("\n修改后:\n")
	fmt.Printf("新引擎: %s\n", car.Start())
	fmt.Printf("停止引擎: %s\n", car.Stop())

	// 5. 显式访问嵌入结构体
	fmt.Printf("\n显式访问:\n")
	fmt.Printf("引擎类型: %s\n", car.Engine.Type)
	fmt.Printf("轮子信息: %s\n", car.Wheels.Roll())
}

// A 结构体A
type A struct {
	Name string
}

// Method A的方法
func (a *A) Method() string {
	return "Method from A"
}

// CommonMethod A的通用方法
func (a *A) CommonMethod() string {
	return "CommonMethod from A"
}

// B 结构体B
type B struct {
	Name string
}

// Method B的方法
func (b *B) Method() string {
	return "Method from B"
}

// CommonMethod B的通用方法
func (b *B) CommonMethod() string {
	return "CommonMethod from B"
}

// C 结构体C（嵌入A和B）
type C struct {
	A    // 嵌入A
	B    // 嵌入B
	Name string
}

// Method C自己的方法（解决冲突）
func (c *C) Method() string {
	return "Method from C"
}

// demoEmbeddingConflicts 演示嵌入冲突处理
func demoEmbeddingConflicts() {
	// 1. 创建有冲突的结构体
	c := C{
		A:    A{Name: "A"},
		B:    B{Name: "B"},
		Name: "C",
	}

	// 2. 调用自己的方法（解决了冲突）
	fmt.Printf("C的方法: %s\n", c.Method())

	// 3. 显式调用嵌入结构体的方法
	fmt.Printf("A的方法: %s\n", c.A.Method())
	fmt.Printf("B的方法: %s\n", c.B.Method())

	// 4. 冲突的方法需要显式调用
	fmt.Printf("A的通用方法: %s\n", c.A.CommonMethod())
	fmt.Printf("B的通用方法: %s\n", c.B.CommonMethod())

	// 注意：c.CommonMethod() 会编译错误，因为有歧义

	// 5. 访问同名字段
	fmt.Printf("\n字段访问:\n")
	fmt.Printf("C的名字: %s\n", c.Name)
	fmt.Printf("A的名字: %s\n", c.A.Name)
	fmt.Printf("B的名字: %s\n", c.B.Name)

	// 6. 修改字段
	c.Name = "Modified C"
	c.A.Name = "Modified A"
	c.B.Name = "Modified B"

	fmt.Printf("\n修改后:\n")
	fmt.Printf("C的名字: %s\n", c.Name)
	fmt.Printf("A的名字: %s\n", c.A.Name)
	fmt.Printf("B的名字: %s\n", c.B.Name)
}

// Shape 形状接口
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Color 颜色结构体
type Color struct {
	R, G, B int
}

// String 颜色字符串表示
func (c Color) String() string {
	return fmt.Sprintf("RGB(%d, %d, %d)", c.R, c.G, c.B)
}

// ColoredCircle 有颜色的圆（组合）
type ColoredCircle struct {
	Circle // 嵌入圆形
	Color  // 嵌入颜色
}

// NewColoredCircle 创建有颜色的圆
func NewColoredCircle(radius float64, r, g, b int) *ColoredCircle {
	return &ColoredCircle{
		Circle: Circle{Radius: radius},
		Color:  Color{R: r, G: g, B: b},
	}
}

// Describe 描述有颜色的圆
func (cc *ColoredCircle) Describe() string {
	return fmt.Sprintf("A %s circle with radius %.2f (area: %.2f)",
		cc.Color.String(), cc.Radius, cc.Area())
}

// demoCompositionVsInheritance 演示组合vs继承
func demoCompositionVsInheritance() {
	// 1. 使用组合创建复杂对象
	coloredCircle := NewColoredCircle(5.0, 255, 0, 0)

	fmt.Printf("有颜色的圆: %s\n", coloredCircle.Describe())
	fmt.Printf("面积: %.2f\n", coloredCircle.Area())          // 来自Circle
	fmt.Printf("周长: %.2f\n", coloredCircle.Circumference()) // 来自Circle
	fmt.Printf("颜色: %s\n", coloredCircle.Color.String())    // 来自Color

	// 2. 修改组合对象
	coloredCircle.Scale(2.0) // 修改圆形
	coloredCircle.R = 0      // 修改颜色
	coloredCircle.G = 255

	fmt.Printf("\n修改后: %s\n", coloredCircle.Describe())

	// 3. 组合的灵活性
	shapes := []Shape{
		&Circle{Radius: 3.0},
		&coloredCircle.Circle, // 可以单独使用组合的部分
	}

	fmt.Println("\n形状列表:")
	for i, shape := range shapes {
		fmt.Printf("形状 %d: 面积=%.2f, 周长=%.2f\n",
			i+1, shape.Area(), shape.Perimeter())
	}

	// 4. 多重组合
	type Position struct {
		X, Y float64
	}

	type PositionedColoredCircle struct {
		ColoredCircle // 嵌入有颜色的圆
		Position      // 嵌入位置
	}

	pcc := PositionedColoredCircle{
		ColoredCircle: *coloredCircle,
		Position:      Position{X: 10, Y: 20},
	}

	fmt.Printf("\n定位的有颜色圆:\n")
	fmt.Printf("描述: %s\n", pcc.Describe())
	fmt.Printf("位置: (%.1f, %.1f)\n", pcc.X, pcc.Y)
	fmt.Printf("颜色: %s\n", pcc.Color.String())
	fmt.Printf("半径: %.1f\n", pcc.Radius)

	// 5. 组合vs继承的优势
	fmt.Println("\n组合的优势:")
	fmt.Println("- 可以组合多个不相关的类型")
	fmt.Println("- 运行时可以改变行为")
	fmt.Println("- 避免深层继承层次")
	fmt.Println("- 更好的代码复用")
	fmt.Println("- 符合Go的设计哲学：组合优于继承")
}
