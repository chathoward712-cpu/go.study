package stage1

import (
	"fmt"
	"math"
	"strconv"
	"unsafe"
)

// DemoNumericTypes 演示数值类型
func DemoNumericTypes() {
	fmt.Println("\n=== 数值类型演示 ===")

	// 1. 整数类型
	fmt.Println("\n1. 整数类型：")
	var int8Val int8 = 127
	var int16Val int16 = 32767
	var int32Val int32 = 2147483647
	var int64Val int64 = 9223372036854775807

	fmt.Printf("int8: %d (大小: %d字节, 范围: %d ~ %d)\n",
		int8Val, unsafe.Sizeof(int8Val), math.MinInt8, math.MaxInt8)
	fmt.Printf("int16: %d (大小: %d字节, 范围: %d ~ %d)\n",
		int16Val, unsafe.Sizeof(int16Val), math.MinInt16, math.MaxInt16)
	fmt.Printf("int32: %d (大小: %d字节, 范围: %d ~ %d)\n",
		int32Val, unsafe.Sizeof(int32Val), math.MinInt32, math.MaxInt32)
	fmt.Printf("int64: %d (大小: %d字节)\n", int64Val, unsafe.Sizeof(int64Val))

	// 2. 无符号整数类型
	fmt.Println("\n2. 无符号整数类型：")
	var uint8Val uint8 = 255
	var uint16Val uint16 = 65535
	var uint32Val uint32 = 4294967295
	var uint64Val uint64 = 18446744073709551615

	fmt.Printf("uint8: %d (大小: %d字节, 范围: 0 ~ %d)\n",
		uint8Val, unsafe.Sizeof(uint8Val), math.MaxUint8)
	fmt.Printf("uint16: %d (大小: %d字节, 范围: 0 ~ %d)\n",
		uint16Val, unsafe.Sizeof(uint16Val), math.MaxUint16)
	fmt.Printf("uint32: %d (大小: %d字节, 范围: 0 ~ %d)\n",
		uint32Val, unsafe.Sizeof(uint32Val), math.MaxUint32)
	fmt.Printf("uint64: %d (大小: %d字节)\n", uint64Val, unsafe.Sizeof(uint64Val))

	// 3. 平台相关类型
	fmt.Println("\n3. 平台相关类型：")
	var intVal int = 42
	var uintVal uint = 42
	var uintptrVal uintptr = 0x12345678

	fmt.Printf("int: %d (大小: %d字节)\n", intVal, unsafe.Sizeof(intVal))
	fmt.Printf("uint: %d (大小: %d字节)\n", uintVal, unsafe.Sizeof(uintVal))
	fmt.Printf("uintptr: 0x%x (大小: %d字节)\n", uintptrVal, unsafe.Sizeof(uintptrVal))

	// 4. 浮点数类型
	fmt.Println("\n4. 浮点数类型：")
	var float32Val float32 = 3.14159
	var float64Val float64 = 3.141592653589793

	fmt.Printf("float32: %.7f (大小: %d字节, 精度: ~7位)\n",
		float32Val, unsafe.Sizeof(float32Val))
	fmt.Printf("float64: %.15f (大小: %d字节, 精度: ~15位)\n",
		float64Val, unsafe.Sizeof(float64Val))

	// 5. 复数类型
	fmt.Println("\n5. 复数类型：")
	var complex64Val complex64 = 3 + 4i
	var complex128Val complex128 = 5 + 12i

	fmt.Printf("complex64: %v (大小: %d字节)\n", complex64Val, unsafe.Sizeof(complex64Val))
	fmt.Printf("complex128: %v (大小: %d字节)\n", complex128Val, unsafe.Sizeof(complex128Val))
	fmt.Printf("复数运算: |%v| = %.2f\n", complex128Val,
		math.Sqrt(real(complex128Val)*real(complex128Val)+imag(complex128Val)*imag(complex128Val)))

	// 6. 类型转换
	fmt.Println("\n6. 类型转换：")
	var a int = 42
	var b float64 = float64(a)
	var c int32 = int32(a)

	fmt.Printf("int转float64: %d -> %.1f\n", a, b)
	fmt.Printf("int转int32: %d -> %d\n", a, c)

	// 注意：不同类型之间不能直接运算
	// fmt.Println(a + b) // 编译错误
	fmt.Printf("类型转换后运算: %d + %.1f = %.1f\n", a, b, float64(a)+b)

	// 7. 数值字面量
	fmt.Println("\n7. 数值字面量：")
	decimal := 42
	binary := 0b101010  // 二进制
	octal := 0o52       // 八进制
	hexadecimal := 0x2A // 十六进制

	fmt.Printf("十进制: %d\n", decimal)
	fmt.Printf("二进制: 0b101010 = %d\n", binary)
	fmt.Printf("八进制: 0o52 = %d\n", octal)
	fmt.Printf("十六进制: 0x2A = %d\n", hexadecimal)

	// 8. 科学计数法
	fmt.Println("\n8. 科学计数法：")
	scientific1 := 1.23e4  // 12300
	scientific2 := 1.23e-4 // 0.000123

	fmt.Printf("1.23e4 = %.1f\n", scientific1)
	fmt.Printf("1.23e-4 = %.6f\n", scientific2)
}

// DemoStringTypes 演示字符串类型
func DemoStringTypes() {
	fmt.Println("\n=== 字符串类型演示 ===")

	// 1. 字符串基础
	fmt.Println("\n1. 字符串基础：")
	str1 := "Hello, 世界!"
	str2 := `这是一个
多行字符串
可以包含"引号"`

	fmt.Printf("普通字符串: %s (长度: %d字节)\n", str1, len(str1))
	fmt.Printf("原始字符串: %s\n", str2)

	// 2. 字符串是不可变的
	fmt.Println("\n2. 字符串不可变性：")
	original := "Hello"
	// original[0] = 'h' // 编译错误：字符串不可变
	modified := "h" + original[1:]
	fmt.Printf("原字符串: %s\n", original)
	fmt.Printf("修改后: %s\n", modified)

	// 3. 字符串索引和切片
	fmt.Println("\n3. 字符串索引和切片：")
	text := "Go语言"
	fmt.Printf("字符串: %s\n", text)
	fmt.Printf("第一个字节: %c (ASCII: %d)\n", text[0], text[0])
	fmt.Printf("前两个字节: %s\n", text[0:2])
	fmt.Printf("从第3个字节开始: %s\n", text[2:])

	// 4. rune 类型（Unicode字符）
	fmt.Println("\n4. rune 类型（Unicode字符）：")
	var r1 rune = 'A'
	var r2 rune = '中'
	var r3 rune = '🚀'

	fmt.Printf("rune 'A': %c (Unicode: %d, 0x%X)\n", r1, r1, r1)
	fmt.Printf("rune '中': %c (Unicode: %d, 0x%X)\n", r2, r2, r2)
	fmt.Printf("rune '🚀': %c (Unicode: %d, 0x%X)\n", r3, r3, r3)

	// 5. 字符串遍历
	fmt.Println("\n5. 字符串遍历：")
	chinese := "Go语言"

	fmt.Println("按字节遍历:")
	for i := 0; i < len(chinese); i++ {
		fmt.Printf("  索引%d: %c (0x%X)\n", i, chinese[i], chinese[i])
	}

	fmt.Println("按rune遍历:")
	for i, r := range chinese {
		fmt.Printf("  索引%d: %c (Unicode: %d)\n", i, r, r)
	}

	// 6. 字符串转换
	fmt.Println("\n6. 字符串转换：")

	// 字符串转数字
	numStr := "123"
	num, err := strconv.Atoi(numStr)
	if err == nil {
		fmt.Printf("字符串转整数: \"%s\" -> %d\n", numStr, num)
	}

	floatStr := "3.14"
	floatNum, err := strconv.ParseFloat(floatStr, 64)
	if err == nil {
		fmt.Printf("字符串转浮点数: \"%s\" -> %.2f\n", floatStr, floatNum)
	}

	// 数字转字符串
	intVal := 456
	intStr := strconv.Itoa(intVal)
	fmt.Printf("整数转字符串: %d -> \"%s\"\n", intVal, intStr)

	floatVal := 2.718
	floatStr2 := strconv.FormatFloat(floatVal, 'f', 3, 64)
	fmt.Printf("浮点数转字符串: %.3f -> \"%s\"\n", floatVal, floatStr2)

	// 7. 字符串和字节切片转换
	fmt.Println("\n7. 字符串和字节切片转换：")
	str := "Hello"
	bytes := []byte(str)
	backToStr := string(bytes)

	fmt.Printf("字符串: %s\n", str)
	fmt.Printf("字节切片: %v\n", bytes)
	fmt.Printf("转回字符串: %s\n", backToStr)

	// 8. 字符串和rune切片转换
	fmt.Println("\n8. 字符串和rune切片转换：")
	unicodeStr := "Go语言🚀"
	runes := []rune(unicodeStr)
	backToUnicodeStr := string(runes)

	fmt.Printf("Unicode字符串: %s (字节长度: %d)\n", unicodeStr, len(unicodeStr))
	fmt.Printf("rune切片: %v (rune个数: %d)\n", runes, len(runes))
	fmt.Printf("转回字符串: %s\n", backToUnicodeStr)
}

// DemoBoolType 演示布尔类型
func DemoBoolType() {
	fmt.Println("\n=== 布尔类型演示 ===")

	// 1. 布尔值基础
	fmt.Println("\n1. 布尔值基础：")
	var isTrue bool = true
	var isFalse bool = false
	var defaultBool bool // 零值为false

	fmt.Printf("true: %t\n", isTrue)
	fmt.Printf("false: %t\n", isFalse)
	fmt.Printf("零值: %t\n", defaultBool)

	// 2. 逻辑运算符
	fmt.Println("\n2. 逻辑运算符：")
	a, b := true, false

	fmt.Printf("a = %t, b = %t\n", a, b)
	fmt.Printf("a && b (与): %t\n", a && b)
	fmt.Printf("a || b (或): %t\n", a || b)
	fmt.Printf("!a (非): %t\n", !a)
	fmt.Printf("!b (非): %t\n", !b)

	// 3. 比较运算符
	fmt.Println("\n3. 比较运算符：")
	x, y := 10, 20

	fmt.Printf("x = %d, y = %d\n", x, y)
	fmt.Printf("x == y: %t\n", x == y)
	fmt.Printf("x != y: %t\n", x != y)
	fmt.Printf("x < y: %t\n", x < y)
	fmt.Printf("x > y: %t\n", x > y)
	fmt.Printf("x <= y: %t\n", x <= y)
	fmt.Printf("x >= y: %t\n", x >= y)

	// 4. 短路求值
	fmt.Println("\n4. 短路求值演示：")

	// && 短路：如果第一个为false，不会执行第二个
	fmt.Println("false && (会跳过的表达式)")
	result1 := false && printAndReturnTrue("这不会被打印")
	fmt.Printf("结果: %t\n", result1)

	// || 短路：如果第一个为true，不会执行第二个
	fmt.Println("true || (会跳过的表达式)")
	result2 := true || printAndReturnTrue("这也不会被打印")
	fmt.Printf("结果: %t\n", result2)

	// 5. 布尔值在条件语句中的使用
	fmt.Println("\n5. 布尔值在条件语句中：")
	isReady := true
	hasPermission := false

	if isReady && hasPermission {
		fmt.Println("可以执行操作")
	} else if isReady && !hasPermission {
		fmt.Println("准备就绪但没有权限")
	} else if !isReady && hasPermission {
		fmt.Println("有权限但未准备就绪")
	} else {
		fmt.Println("既没准备好也没权限")
	}

	// 6. 布尔值转换
	fmt.Println("\n6. 布尔值转换：")
	// Go中不支持隐式类型转换，包括布尔值
	// if 1 { } // 编译错误
	// if "hello" { } // 编译错误

	// 需要显式比较
	num := 0
	str := ""

	fmt.Printf("数字0的布尔判断: %t\n", num != 0)
	fmt.Printf("空字符串的布尔判断: %t\n", str != "")

	// 7. 三元运算符的替代
	fmt.Println("\n7. 条件赋值（Go没有三元运算符）：")
	score := 85
	var grade string

	if score >= 90 {
		grade = "A"
	} else if score >= 80 {
		grade = "B"
	} else if score >= 70 {
		grade = "C"
	} else {
		grade = "D"
	}

	fmt.Printf("分数 %d 对应等级: %s\n", score, grade)
}

// printAndReturnTrue 辅助函数，用于演示短路求值
func printAndReturnTrue(msg string) bool {
	fmt.Println(msg)
	return true
}
