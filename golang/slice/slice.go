package main

import "fmt"

func addrStore() {
	a := []int{1, 2, 3, 4, 5}
	b := []*int{}

	for k, v := range a {
		fmt.Println(k, &v)
		b = append(b, &v)
	}
	
	fmt.Println(b)
}

func intSliceMerge(s1 []int, s2 []int) []int {
	slice := make([]int, len(s1) + len(s2))

	copy(slice, s1)
	copy(slice[len(s1):], s2)

	return slice
}