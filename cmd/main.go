package main

import (
	"fmt"
	"log"

	"github.com/nikitakutrergin59/RPN/pkg/calculation/calc"
)

func main() {
	var input string
	fmt.Print("Введите выражение: ")
	fmt.Scanln(&input)
	result, err := calc.Calc(input)
	if err != nil {
		log.Fatalf("ошибка: %v", err)
	} else {
		fmt.Printf("Рузультат: %f", result)
	}
}
