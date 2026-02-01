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

// Part 4 function
func ExploreProcess() {
	pid := os.Getpid()
	fmt.Println(pid)
	ppid := os.Getppid()
	fmt.Println(ppid)

	data := []int{1, 2, 3, 4, 5}
	fmt.Printf("Memory address of data slice is: %p\n", &data)
	fmt.Printf("Memory address of first element in slice is: %p\n", &data[0])
	fmt.Println("Note: Other processes cannot access these memory addresses directly.")
	//What a process ID is: process id is a number of each individual isntance running on a computer's operating systems allowing for you monitor different statistics
	//Process Isolation is important because it allows for processes not to interrupt one another and also allows you to use multiple processes at a time
	//Difference between the slice header address and element address is the slice header address gives you the memory address of the slice along with the pointer, length and capacity (usually stored in the stack) while the element address is the memory of the first element in the array and is usually stored in the heap

}

// Part 5 functions:
func DoubleValue(x int) {
	x = x * x
	fmt.Println(x)
	//I don't believe this will modify the orginal variable but instead create a copy for use
}

func DoublePointer(x *int) {
	*x = (*x) * (*x)
}

func CreateOnStack() int {
	localVar := 0
	return localVar
	//This variable stays on the stack
}

func CreateOnHeap() *int {
	localVar := 0
	var p *int = &localVar
	return p
}

func SwapValues(a, b int) (int, int) {
	temp := a
	a = b
	b = temp
	return a, b
}

func SwapPointers(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func AnalyzeEscape() {
	CreateOnStack()
	CreateOnHeap()
	/*
	   	In a comment in your code, explain:

	   Which variables escaped to the heap? 'x' variable escaped to the heap
	   Why did they escape? the x variable escaped to the heap because it uses fmt.Println which forces it to the heap
	   What does "escapes to heap" mean? Escape to the heap means that data is kept after the function ends usually in function returns so when a function ends it returns variables which escape to the heap for longer use
	*/
}

func main() {
	ExploreProcess()
}
