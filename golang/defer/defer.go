package main

import "fmt"

func main() {
	n := 1

	defer func() {
		fmt.Println(n)
	}()

	defer fmt.Println(n)

	defer func(n int) {
		fmt.Println(n)
	}(n)

	n += 1
}