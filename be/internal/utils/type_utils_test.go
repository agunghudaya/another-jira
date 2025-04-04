package utils

import (
	"testing"
)

func TestSafeFloat64(t *testing.T) {
	tests := []struct {
		name         string
		input        interface{}
		defaultValue float64
		expected     float64
	}{
		{
			name:         "Valid float64 value",
			input:        42.5,
			defaultValue: 0.0,
			expected:     42.5,
		},
		{
			name:         "Invalid type (string)",
			input:        "not a float",
			defaultValue: 0.0,
			expected:     0.0,
		},
		{
			name:         "Invalid type (int)",
			input:        42,
			defaultValue: 1.1,
			expected:     1.1,
		},
		{
			name:         "Nil value",
			input:        nil,
			defaultValue: 2.2,
			expected:     2.2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SafeFloat64(tt.input, tt.defaultValue)
			if result != tt.expected {
				t.Errorf("SafeFloat64(%v, %v) = %v; want %v", tt.input, tt.defaultValue, result, tt.expected)
			}
		})
	}
}
