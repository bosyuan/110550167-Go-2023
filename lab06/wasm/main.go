package main

import (
	"fmt"
	"math/big"
	"syscall/js"
)

func CheckPrime(this js.Value, args []js.Value) interface{} {
	// Extract the input value from the function arguments
	str := js.Global().Get("value").Get("value").String()
	inputValue, success := new(big.Int).SetString(str, 10)
	if !success {
		return js.ValueOf(nil)
	}

	// Check if the number is prime
	isPrime := isPrime(inputValue)

	// Convert the result to a string and return it
	result := "It's not prime"
	if isPrime {
		result = "It's prime"
	}

	// Set the result in the answer element
	js.Global().Get("answer").Set("innerText", result)

	return nil
}

func isPrime(n *big.Int) bool {
	if n.Cmp(big.NewInt(1)) <= 0 {
		return false
	}

	// Use the ProbablyPrime method to check primality
	return n.ProbablyPrime(20)
}

func registerCallbacks() {
	// Register the CheckPrime function
	js.Global().Set("CheckPrime", js.FuncOf(CheckPrime))
}

func main() {
	fmt.Println("Golang main function executed")
	registerCallbacks()

	// Block the main thread forever
	select {}
}
