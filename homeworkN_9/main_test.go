package main

import (
	"testing"
)

func TestDivide(t *testing.T) {
	tests := []struct {
		name        string
		a           float64
		b           float64
		expected    float64
		expectError bool
	}{
		{
			name:        "обычное деление",
			a:           10,
			b:           2,
			expected:    5,
			expectError: false,
		},
		{
			name:        "деление на ноль",
			a:           5,
			b:           0,
			expected:    0,
			expectError: true,
		},
		{
			name:        "деление отрицательных чисел",
			a:           -9,
			b:           3,
			expected:    -3,
			expectError: false,
		},
		{
			name:        "деление дробных чисел",
			a:           7.5,
			b:           2.5,
			expected:    3,
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := divide(tt.a, tt.b)

			if tt.expectError && err == nil {
				t.Fatalf("ожидалась ошибка, но её нет")
			}
			if !tt.expectError && err != nil {
				t.Fatalf("ошибка не ожидалась, но получена: %v", err)
			}
			if !tt.expectError && result != tt.expected {
				t.Errorf("divide(%v, %v) = %v, ожидалось %v",
					tt.a, tt.b, result, tt.expected)
			}
		})
	}
}
