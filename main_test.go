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

func TestMakeCounter(t *testing.T) {
	tests := []struct {
		name  string
		start int
		calls int
		want  []int
	}{
		{name: "start at 0 three calls", start: 0, calls: 3, want: []int{1, 2, 3}},
		{name: "start at 5 three calls", start: 5, calls: 3, want: []int{1, 2, 3}},
		{name: "start negative", start: -4, calls: 2, want: []int{1, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			counter := MakeCounter(tt.start)

			for i := 0; i < tt.calls; i++ {
				got := counter()
				if got != tt.want[i] {
					t.Errorf("call %d = %d, want %d", i+1, got, tt.want[i])
				}
			}
		})
	}
}

func TestMakeMultiplier(t *testing.T) {
	tests := []struct {
		name   string
		factor int
		input  int
		want   int
	}{
		{name: "double", factor: 2, input: 3, want: 6},
		{name: "times five", factor: 5, input: 4, want: 20},
		{name: "zero input", factor: 7, input: 0, want: 0},
		{name: "negative input", factor: 3, input: -2, want: -6},
		{name: "negative factor", factor: -2, input: 3, want: -6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MakeMultiplier(tt.factor)
			got := m(tt.input)
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestMakeAccumulator(t *testing.T) {
	tests := []struct {
		name    string
		initial int
		adds    []int
		subs    []int
		want    int
	}{
		{name: "add only", initial: 10, adds: []int{5, 5}, subs: []int{}, want: 20},
		{name: "add and subtract", initial: 10, adds: []int{5}, subs: []int{3}, want: 12},
		{name: "subtract only", initial: 10, adds: []int{}, subs: []int{4}, want: 6},
		{name: "mixed ops", initial: 0, adds: []int{10, 2}, subs: []int{5}, want: 7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			add, subtract, get := MakeAccumulator(tt.initial)

			for _, v := range tt.adds {
				add(v)
			}
			for _, v := range tt.subs {
				subtract(v)
			}

			got := get()
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}
