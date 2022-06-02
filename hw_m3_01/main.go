package main

import (
	"fmt"
)

type HealthCheck struct {
	ServiceID int
	status    MyType
}

type MyType string

const (
	PassStatus MyType = "pass"
	FailStatus MyType = "fail"
)

func main() {
	fmt.Println("Пробуем торрент курс!")
	slice := GenerateCheck()
	for _, val := range slice {
		if val.status == PassStatus {
			fmt.Println(val.ServiceID)
		}
	}
}

func GenerateCheck() []HealthCheck {
	slice := make([]HealthCheck, 0, 6)
	var stat MyType
	for i := 0; i <= 5; i++ {
		if i%2 == 0 {
			stat = PassStatus
		} else {
			stat = FailStatus
		}
		slice = append(slice, HealthCheck{i, stat})

	}

	return slice
}
