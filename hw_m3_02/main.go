package main

import (
	"fmt"
)

func main() {
	nums := [8]int{8, 11, 7, 1, 5, 13, 1, 12}
	var b bool

	for i, val := range nums {
		for j, vall := range nums {
			if val == vall && i != j {
				b = true
				break
			}
			if b {
				break
			}
		}
	}

	fmt.Println(b)
}
