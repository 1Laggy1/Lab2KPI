// main.go
package main

import (
	"flag"
	"fmt"
	"os"

	lab2 "github.com/1Laggy1/Lab2KPI"
)

func main() {
	// Оголошення флагів
	expressionFlag := flag.String("e", "", "вираз")
	fileFlag := flag.String("f", "", "файл із виразом")
	outputFlag := flag.String("o", "", "файл результату")

	// Парсинг флагів командного рядка
	flag.Parse()

	// Розпізнавання типу вводу
	input, err := lab2.ParseInput(*expressionFlag, *fileFlag)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Помилка визначення вводу: %v\n", err)
		os.Exit(1)
	}

	// Розпізнавання типу виведення
	output, err := lab2.ParseOutput(*outputFlag)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Помилка визначення виводу: %v\n", err)
		os.Exit(1)
	}

	// Створення екземпляру ComputeHandler та обробка виразу
	handler := lab2.NewComputeHandler(input, output)
	err = handler.Compute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Помилка обробки виразу: %v\n", err)
		os.Exit(1)
	}
}
