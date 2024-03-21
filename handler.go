package lab2

import (
	"errors"
	"io"
	"io/ioutil"
	"strings"
)

// ComputeHandler відповідає за обробку виразів та запис результатів
type ComputeHandler struct {
	input  string
	output io.Writer
}

// NewComputeHandler створює новий екземпляр ComputeHandler
func NewComputeHandler(input string, output io.Writer) *ComputeHandler {
	return &ComputeHandler{input: input, output: output}
}

// Compute обробляє вираз та записує результат
func (ch *ComputeHandler) Compute() error {
	// Розділити введений рядок на окремі елементи
	expression := strings.Fields(ch.input)
	if len(expression) == 0 {
		return errors.New("відсутні дані для обробки")
	}

	// Логіка обчислення виразу з використанням функцій з implementation.go
	result, err := PostfixToInfix(strings.Join(expression, " "))
	if err != nil {
		return err
	}

	// Запис результату у вказане місце
	_, err = ch.output.Write([]byte(result))
	if err != nil {
		return err
	}

	return nil
}

// ParseInput розпізнає тип введення та повертає відповідне джерело вводу
func ParseInput(expressionFlag, fileFlag string) (string, error) {
	if expressionFlag != "" {
		return expressionFlag, nil
	} else if fileFlag != "" {
		data, err := ioutil.ReadFile(fileFlag)
		if err != nil {
			return "", err
		}
		return string(data), nil
	} else {
		return "", errors.New("не вказаний вираз або файл з виразом")
	}
}
