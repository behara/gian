package main

import "testing"

func Test_mytest(t *testing.T) {
	tests := []struct {
		name string
		expected string
	}{
		{"test",
			"test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}