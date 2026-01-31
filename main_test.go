package main

import (
	"testing"
)

// Part 1 tests
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

// Part 2 Tests
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

// Part 3 tests
func TestApply(t *testing.T) {
	tests := []struct {
		name      string
		nums      []int
		operation func(int) int
		want      []int
	}{
		{name: "square", nums: []int{1, 2, 3}, operation: func(x int) int { return x * x }, want: []int{1, 4, 9}},
		{name: "double", nums: []int{1, 2, 3}, operation: func(x int) int { return x * 2 }, want: []int{2, 4, 6}},
		{name: "negate", nums: []int{1, -2, 3}, operation: func(x int) int { return -x }, want: []int{-1, 2, -3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Apply(tt.nums, tt.operation)
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("Apply()[%d] = %d, want %d", i, got[i], tt.want[i])
				}
			}
		})
	}
}

func TestFilter(t *testing.T) {
	tests := []struct {
		name      string
		nums      []int
		predicate func(int) bool
		want      []int
	}{
		{name: "even numbers", nums: []int{1, 2, 3, 4}, predicate: func(x int) bool { return x%2 == 0 }, want: []int{2, 4}},
		{name: "positive numbers", nums: []int{-1, 2, -3, 4}, predicate: func(x int) bool { return x > 0 }, want: []int{2, 4}},
		{name: "greater than 10", nums: []int{5, 11, 20, 7}, predicate: func(x int) bool { return x > 10 }, want: []int{11, 20}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Filter(tt.nums, tt.predicate)
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("Filter()[%d] = %d, want %d", i, got[i], tt.want[i])
				}
			}
		})
	}
}
func TestReduce(t *testing.T) {
	tests := []struct {
		name      string
		nums      []int
		initial   int
		operation func(int, int) int
		want      int
	}{
		{name: "sum", nums: []int{1, 2, 3, 4}, initial: 0, operation: func(acc, cur int) int { return acc + cur }, want: 10},
		{name: "product", nums: []int{1, 2, 3, 4}, initial: 1, operation: func(acc, cur int) int { return acc * cur }, want: 24},
		{name: "max", nums: []int{1, 5, 3, 4}, initial: -999, operation: func(acc, cur int) int {
			if cur > acc {
				return cur
			}
			return acc
		}, want: 5},
		{name: "min", nums: []int{1, 5, 3, 4}, initial: 999, operation: func(acc, cur int) int {
			if cur < acc {
				return cur
			}
			return acc
		}, want: 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Reduce(tt.nums, tt.initial, tt.operation)
			if got != tt.want {
				t.Errorf("Reduce() = %d, want %d", got, tt.want)
			}
		})
	}
}
func TestCompose(t *testing.T) {
	tests := []struct {
		name  string
		f     func(int) int
		g     func(int) int
		input int
		want  int
	}{
		{name: "square after double", f: func(x int) int { return x * x }, g: func(x int) int { return x * 2 }, input: 3, want: 36}, // (3*2)^2
		{name: "double after square", f: func(x int) int { return x * 2 }, g: func(x int) int { return x * x }, input: 3, want: 18}, // 2*(3^2)
		{name: "negate after add 5", f: func(x int) int { return -x }, g: func(x int) int { return x + 5 }, input: 2, want: -7},     // -(2+5)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Compose(tt.f, tt.g)
			got := h(tt.input)
			if got != tt.want {
				t.Errorf("Compose() = %d, want %d", got, tt.want)
			}
		})
	}
}
