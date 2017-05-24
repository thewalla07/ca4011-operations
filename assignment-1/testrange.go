package main

import (
	"fmt"
)

func main() {
	var a []int
	a = append(a, 1, 2, 3, 4)

	for _, e := range a {
		fmt.Println(e)
	}
	for i, e := range a {
		a[i] += e
	}
	for _, e := range a {
		fmt.Println(e)
	}
}
