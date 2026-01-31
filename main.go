package main

import (
	"errors"
	"fmt"
	"math"
	"os"
)

// Part 1 Functions
func Factorial(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("Factorial is not defined for negative numbers")
	}

	if n == 0 {
		return 1, nil
	}

	result, err := Factorial(n - 1)
	if err != nil {
		return 0, err
	}

	return n * result, nil
}

// A prime number can only be divided (without a remainder) by itself and 1
func IsPrime(n int) (bool, error) {
	if n < 2 {
		return false, errors.New("prime check requires number >= 2")
	}

	if n == 2 {
		return true, nil
	}

	if n%2 == 0 {
		return false, nil
	}

	//For loop starting from 2 and does not go beyond √n
	for i := 3; float64(i) <= math.Sqrt(float64(n)); i = i + 2 {
		if n%i == 0 {
			return false, nil
		}
	}
	return true, nil
}

func Power(base, exponent int) (int, error) {
	if exponent < 0 {
		return 0, errors.New("negative exponents not supported")
	}
	//any number raised to zero exponent is equal 1
	if exponent == 0 {
		return 1, nil
	}

	if base == 0 {
		return 0, nil
	}

	result := 1
	for i := 0; i < exponent; i++ {
		result = result * base
	}
	return result, nil
}

// Part 2 Functions
func MakeCounter(start int) func() int {
	counter := 0

	return func() int {
		counter++
		return counter
	}
}

func MakeMultiplier(factor int) func(int) int {
	return func(num int) int {
		return factor * num
	}
}

func MakeAccumulator(initial int) (add func(int), subtract func(int), get func() int) {
	total := initial
	add = func(num int) {
		total += num
	}

	subtract = func(num int) {
		total -= num
	}

	get = func() int {
		return total
	}

	return add, subtract, get
}

// Part 3 functions
func Apply(nums []int, operation func(int) int) []int {
	newSlice := []int{}

	//For loop to loop through the slice and apply the function to each of the elements
	for _, v := range nums {
		newSlice = append(newSlice, operation(v))
	}

	return newSlice
}

func Filter(nums []int, predicate func(int) bool) []int {
	newSlice := []int{}

	for _, v := range nums {
		if predicate(v) {
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}

func Reduce(nums []int, initial int, operation func(accumulator, current int) int) int {
	accumulator := initial

	for i := 0; i < len(nums); i++ {
		current := nums[i]
		accumulator = operation(accumulator, current)
	}
	return accumulator
}

func Compose(f func(int) int, g func(int) int) func(int) int {
	return func(x int) int {
		//returns a composite function that gets the answer to g(x) first then f() is run and given the result of g(x)
		return f(g(x))
	}
}

func ExploreProcess() {
	pid := os.Getpid()
	fmt.Println(pid)
	ppid := os.Getppid()
	fmt.Println(ppid)

	data := []int{1, 2, 3, 4, 5}
	fmt.Printf("Memory address of data slice is: %p\n", &data)
	fmt.Printf("Memory address of first element in slice is: %p\n", &data[0])
	fmt.Println("Note: Other processes cannot access these memory addresses directly.")
}

func main() {
	ExploreProcess()
}
