package main

import (
	"errors"
	"math"
)

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
	if exponent <= 0 {
		return 0, errors.New("negative exponents not supported")
	}

	result := 1
	for i := 0; i < exponent; i++ {
		result += result * base
	}
	return result, nil
}

func main() {

}
