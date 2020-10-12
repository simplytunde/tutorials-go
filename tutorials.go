package main

import (
	"fmt"
)

func f1() {
	fmt.Println("This is f1")
}
func f2() {
	fmt.Println("This is f2")
}
func main() {
	go f1()
	go f2()
	// time.Sleep(2 * time.Second)
	fmt.Println("I called f1 and f2 Go routine")
}
