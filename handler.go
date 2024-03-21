// lab2/handler.go
package lab2

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strings"
)

// ComputeHandler відповідає за обробку виразів та запис результатів
type ComputeHandler struct {
	input  io.Reader
	output io.Writer
}

// NewComputeHandler створює новий екземпляр ComputeHandler
func NewComputeHandler(input io.Reader, output io.Writer) *ComputeHandler {
	return &ComputeHandler{input: input, output: output}
}

// Compute обробляє вираз та записує результат
// Compute обробляє вираз та записує результат
func (ch *ComputeHandler) Compute() error {
	// Зчитування виразу з введеного джерела
	scanner := bufio.NewScanner(ch.input)
	if !scanner.Scan() {
		return errors.New("відсутні дані для обробки")
	}
	expression := scanner.Text()

	// Логіка обчислення виразу з використанням функцій з implementation.go
	result, err := PostfixToInfix(expression)
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
func ParseInput(expressionFlag, fileFlag string) (io.Reader, error) {
	if expressionFlag != "" {
		return strings.NewReader(expressionFlag), nil
	} else if fileFlag != "" {
		file, err := os.Open(fileFlag)
		if err != nil {
			return nil, err
		}
		defer file.Close()
		return file, nil
	} else {
		return nil, errors.New("не вказаний вираз або файл з виразом")
	}
}

// ParseOutput розпізнає тип виведення та повертає відповідне місце виведення
func ParseOutput(outputFlag string) (io.Writer, error) {
	if outputFlag != "" {
		file, err := os.Create(outputFlag)
		if err != nil {
			return nil, err
		}
		return file, nil
	} else {
		return os.Stdout, nil
	}
}
