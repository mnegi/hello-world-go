// Package main : Executable commands must always use package main.
package main

// import standard fmt package  for formatted I/O functions
// also import the newmath package that we created in example 2
import (
	"fmt"
	"newmath"
	)

// main function : the starting point of an executable go program
func main(){

	// print hello world using fmt packages' Printf function
	fmt.Printf("Hello world!\n")

	// use custom package and call Sqrt on it
	fmt.Printf("Sqrt(2) = %v\n", newmath.Sqrt(2))
}
