package main

import (
	"testing"
)

func TestFactorial(t *testing.T) {
	tests := []struct {
		name      string
		input     int
		want      int
		wantError bool
	}{
		{name: "factorial of 0", input: 0, want: 1, wantError: false},
		// Add at least 5 more test cases
		{name: "factorial of 1", input: 1, want: 1, wantError: false},
		{name: "factorial of negative", input: -53, want: 0, wantError: true},
		{name: "factorial of 5", input: 5, want: 120, wantError: false},
		{name: "factorial of 10", input: 10, want: 3628800, wantError: false},
		{name: "factorial of negative", input: -1, want: 0, wantError: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Factorial(tt.input)
			if (err != nil) != tt.wantError {
				t.Errorf("Factorial() error = %v, wantError %v", err, tt.wantError)
				return
			}
			if got != tt.want {
				t.Errorf("Factorial() = %v, want %v", got, tt.want)
			}

		})
	}
}

// Test cases for IsPrime
func TestIsPrime(t *testing.T) {
	tests := []struct {
		name      string
		input     int
		want      bool
		wantError bool
	}{
		{name: "prime 2", input: 2, want: true, wantError: false},
		{name: "prime 3", input: 3, want: true, wantError: false},
		{name: "composite 4", input: 4, want: false, wantError: false},
		{name: "prime 13", input: 13, want: true, wantError: false},
		{name: "composite 20", input: 20, want: false, wantError: false},
		{name: "negative number", input: -7, want: false, wantError: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsPrime(tt.input)
			if (err != nil) != tt.wantError {
				t.Errorf("IsPrime() error = %v, wantError %v", err, tt.wantError)
				return
			}
			if got != tt.want {
				t.Errorf("IsPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Test cases for Power
func TestPower(t *testing.T) {
	tests := []struct {
		name      string
		base      int
		exponent  int
		want      int
		wantError bool
	}{
		{name: "base^0", base: 5, exponent: 0, want: 1, wantError: false},
		{name: "0^exponent", base: 0, exponent: 5, want: 0, wantError: false},
		{name: "positive exponent", base: 2, exponent: 3, want: 8, wantError: false},
		{name: "negative exponent", base: 2, exponent: -2, want: 0, wantError: true},
		{name: "1^exponent", base: 1, exponent: 100, want: 1, wantError: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Power(tt.base, tt.exponent)
			if (err != nil) != tt.wantError {
				t.Errorf("Power() error = %v, wantError %v", err, tt.wantError)
				return
			}
			if got != tt.want {
				t.Errorf("Power() = %v, want %v", got, tt.want)
			}
		})
	}
}
