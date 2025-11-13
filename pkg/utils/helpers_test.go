package utils

import "testing"

// TestMax 测试 Max 函数
func TestMax(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive numbers", 3, 5, 5},
		{"negative numbers", -10, -5, -5},
		{"equal numbers", 7, 7, 7},
		{"zero and positive", 0, 10, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Max(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, result)
			}
		})
	}
}

// TestMin 测试 Min 函数
func TestMin(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive numbers", 3, 5, 3},
		{"negative numbers", -10, -5, -10},
		{"equal numbers", 7, 7, 7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Min(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, result)
			}
		})
	}
}

// TestReverse 测试 Reverse 函数
func TestReverse(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"simple string", "hello", "olleh"},
		{"with spaces", "hello world", "dlrow olleh"},
		{"chinese characters", "你好Go", "oG好你"},
		{"empty string", "", ""},
		{"single character", "a", "a"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Reverse(tt.input)
			if result != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, result)
			}
		})
	}
}
