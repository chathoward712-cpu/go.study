package stage2

import (
	"fmt"
)

// demoStringConversion 演示字符串转换
func demoStringConversion() {
	// 1. 字符串和字节切片转换
	str := "Hello, 世界"
	bytes := []byte(str)
	backToStr := string(bytes)

	fmt.Printf("字符串转换:\n")
	fmt.Printf("原字符串: %s\n", str)
	fmt.Printf("字节切片: %v\n", bytes)
	fmt.Printf("转回字符串: %s\n", backToStr)

	// 2. 字符串和rune切片转换
	runes := []rune(str)
	backToStr2 := string(runes)

	fmt.Printf("rune切片: %v\n", runes)
	fmt.Printf("转回字符串: %s\n", backToStr2)

	// 3. 数字和字符串转换
	fmt.Printf("数字转换:\n")

	// 整数转字符串
	intVal := 123
	intStr := intToString(intVal)
	fmt.Printf("整数 %d 转字符串: %s\n", intVal, intStr)

	// 字符串转整数
	strVal := "456"
	intVal2, err := stringToInt(strVal)
	if err == nil {
		fmt.Printf("字符串 %s 转整数: %d\n", strVal, intVal2)
	}

	// 浮点数转字符串
	floatVal := 3.14159
	floatStr := floatToString(floatVal, 2)
	fmt.Printf("浮点数 %.5f 转字符串: %s\n", floatVal, floatStr)

	// 4. 布尔值转换
	boolVal := true
	boolStr := boolToString(boolVal)
	fmt.Printf("布尔值 %t 转字符串: %s\n", boolVal, boolStr)

	// 5. 进制转换
	num := 255
	fmt.Printf("进制转换:\n")
	fmt.Printf("十进制 %d 转二进制: %s\n", num, intToBase(num, 2))
	fmt.Printf("十进制 %d 转八进制: %s\n", num, intToBase(num, 8))
	fmt.Printf("十进制 %d 转十六进制: %s\n", num, intToBase(num, 16))
}

// intToString 整数转字符串
func intToString(n int) string {
	if n == 0 {
		return "0"
	}

	negative := n < 0
	if negative {
		n = -n
	}

	var digits []byte
	for n > 0 {
		digits = append([]byte{byte('0' + n%10)}, digits...)
		n /= 10
	}

	if negative {
		digits = append([]byte{'-'}, digits...)
	}

	return string(digits)
}

// stringToInt 字符串转整数
func stringToInt(s string) (int, error) {
	if len(s) == 0 {
		return 0, fmt.Errorf("empty string")
	}

	negative := false
	start := 0

	if s[0] == '-' {
		negative = true
		start = 1
	} else if s[0] == '+' {
		start = 1
	}

	if start >= len(s) {
		return 0, fmt.Errorf("invalid number")
	}

	result := 0
	for i := start; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0, fmt.Errorf("invalid character: %c", s[i])
		}
		result = result*10 + int(s[i]-'0')
	}

	if negative {
		result = -result
	}

	return result, nil
}

// floatToString 浮点数转字符串
func floatToString(f float64, precision int) string {
	return fmt.Sprintf("%."+intToString(precision)+"f", f)
}

// boolToString 布尔值转字符串
func boolToString(b bool) string {
	if b {
		return "true"
	}
	return "false"
}

// intToBase 整数转指定进制字符串
func intToBase(n, base int) string {
	if n == 0 {
		return "0"
	}

	digits := "0123456789ABCDEF"
	if base < 2 || base > 16 {
		return "invalid base"
	}

	negative := n < 0
	if negative {
		n = -n
	}

	var result []byte
	for n > 0 {
		result = append([]byte{digits[n%base]}, result...)
		n /= base
	}

	if negative {
		result = append([]byte{'-'}, result...)
	}

	return string(result)
}

// demoStringValidation 演示字符串验证
func demoStringValidation() {
	// 1. 检查字符串类型
	testStrings := []string{
		"123",
		"12.34",
		"hello",
		"Hello123",
		"HELLO",
		"hello world",
		"",
		"   ",
		"user@example.com",
		"192.168.1.1",
	}

	fmt.Printf("字符串验证:\n")
	for _, s := range testStrings {
		fmt.Printf("'%s':\n", s)
		fmt.Printf("  是否为数字: %t\n", isNumeric(s))
		fmt.Printf("  是否为字母: %t\n", isAlpha(s))
		fmt.Printf("  是否为字母数字: %t\n", isAlphaNumeric(s))
		fmt.Printf("  是否为大写: %t\n", isUpper(s))
		fmt.Printf("  是否为小写: %t\n", isLower(s))
		fmt.Printf("  是否为空白: %t\n", isBlank(s))
		fmt.Printf("  是否为邮箱: %t\n", isEmail(s))
		fmt.Printf("  是否为IP: %t\n", isIP(s))
		fmt.Println()
	}

	// 2. 字符串清理
	dirtyString := "  Hello, World!  \n\t"
	fmt.Printf("字符串清理:\n")
	fmt.Printf("原字符串: '%s'\n", dirtyString)
	fmt.Printf("去除空白: '%s'\n", trimSpace(dirtyString))
	fmt.Printf("去除左空白: '%s'\n", trimLeft(dirtyString))
	fmt.Printf("去除右空白: '%s'\n", trimRight(dirtyString))

	// 3. 字符串长度限制
	longString := "This is a very long string that needs to be truncated"
	fmt.Printf("字符串截断:\n")
	fmt.Printf("原字符串: %s\n", longString)
	fmt.Printf("截断到20字符: %s\n", truncate(longString, 20))
	fmt.Printf("截断到20字符(带省略号): %s\n", truncateWithEllipsis(longString, 20))
}

// isNumeric 检查是否为数字
func isNumeric(s string) bool {
	if len(s) == 0 {
		return false
	}

	start := 0
	if s[0] == '+' || s[0] == '-' {
		start = 1
	}

	if start >= len(s) {
		return false
	}

	dotCount := 0
	for i := start; i < len(s); i++ {
		if s[i] == '.' {
			dotCount++
			if dotCount > 1 {
				return false
			}
		} else if s[i] < '0' || s[i] > '9' {
			return false
		}
	}

	return true
}

// isAlpha 检查是否只包含字母
func isAlpha(s string) bool {
	if len(s) == 0 {
		return false
	}

	for _, r := range s {
		if !((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')) {
			return false
		}
	}
	return true
}

// isAlphaNumeric 检查是否只包含字母和数字
func isAlphaNumeric(s string) bool {
	if len(s) == 0 {
		return false
	}

	for _, r := range s {
		if !((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9')) {
			return false
		}
	}
	return true
}

// isUpper 检查是否全为大写
func isUpper(s string) bool {
	if len(s) == 0 {
		return false
	}

	hasLetter := false
	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			return false
		}
		if r >= 'A' && r <= 'Z' {
			hasLetter = true
		}
	}
	return hasLetter
}

// isLower 检查是否全为小写
func isLower(s string) bool {
	if len(s) == 0 {
		return false
	}

	hasLetter := false
	for _, r := range s {
		if r >= 'A' && r <= 'Z' {
			return false
		}
		if r >= 'a' && r <= 'z' {
			hasLetter = true
		}
	}
	return hasLetter
}

// isBlank 检查是否为空白字符串
func isBlank(s string) bool {
	for _, r := range s {
		if r != ' ' && r != '\t' && r != '\n' && r != '\r' {
			return false
		}
	}
	return true
}

// isEmail 简单的邮箱验证
func isEmail(s string) bool {
	if len(s) == 0 {
		return false
	}

	// 简单检查是否包含@和.
	hasAt := false
	hasDot := false
	atPos := -1

	for i, r := range s {
		if r == '@' {
			if hasAt || i == 0 || i == len(s)-1 {
				return false
			}
			hasAt = true
			atPos = i
		} else if r == '.' && hasAt && i > atPos+1 && i < len(s)-1 {
			hasDot = true
		}
	}

	return hasAt && hasDot
}

// isIP 简单的IP地址验证
func isIP(s string) bool {
	// 简单的IP验证
	parts := make([]string, 0, 4)
	current := ""

	for _, r := range s {
		if r == '.' {
			if current == "" {
				return false
			}
			parts = append(parts, current)
			current = ""
		} else if r >= '0' && r <= '9' {
			current += string(r)
		} else {
			return false
		}
	}

	if current != "" {
		parts = append(parts, current)
	}

	if len(parts) != 4 {
		return false
	}

	for _, part := range parts {
		if !isNumeric(part) {
			return false
		}

		num, err := stringToInt(part)
		if err != nil || num < 0 || num > 255 {
			return false
		}
	}

	return true
}

// trimSpace 去除首尾空白字符
func trimSpace(s string) string {
	return trimRight(trimLeft(s))
}

// trimLeft 去除左侧空白字符
func trimLeft(s string) string {
	start := 0
	for start < len(s) && isWhitespace(rune(s[start])) {
		start++
	}
	return s[start:]
}

// trimRight 去除右侧空白字符
func trimRight(s string) string {
	end := len(s)
	for end > 0 && isWhitespace(rune(s[end-1])) {
		end--
	}
	return s[:end]
}

// isWhitespace 检查是否为空白字符
func isWhitespace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\n' || r == '\r'
}

// truncate 截断字符串
func truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen]
}

// truncateWithEllipsis 截断字符串并添加省略号
func truncateWithEllipsis(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	if maxLen <= 3 {
		return s[:maxLen]
	}
	return s[:maxLen-3] + "..."
}
