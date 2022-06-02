package main

import "fmt"

func main() {
	sentense := "съешь ещё этих мягких французских булок, да выпей чаю"
	m := make(map[rune]int)
	for _, val := range sentense {
		m[val]++
	}
	for key, val := range m {
		fmt.Printf("%q - %d\n", key, val)
	}
}
