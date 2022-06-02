package main

import (
	"fmt"
)

func main() {
	var b bool = true
	slice := make([]string, 5, 5)
	slice[0] = "aaae"
	slice[1] = "aaad"
	slice[2] = "aaac"
	slice[3] = "aaab"
	slice[4] = "aaaa"

	if slice[0] < slice[1] { // проверка  на возрастание
		for i := 0; i < len(slice)-1; i++ {
			if slice[i] > slice[i+1] {
				b = false
				break
			}
		}
	} else { // проверка на убывание
		for i := 0; i < len(slice)-1; i++ {
			if slice[i] < slice[i+1] {
				b = false
				break
			}
		}
	}

	fmt.Println(b)
}
