// handler.go

package lab2

import (
	"fmt"
	"io"
)

type InputReader interface {
	Read() (string, error)
}

type OutputWriter interface {
	Write(string) error
}

type ComputeHandler struct {
	InputReader  InputReader
	OutputWriter OutputWriter
}

func (ch *ComputeHandler) Compute() error {
	expr, err := ch.InputReader.Read()
	if err != nil {
		return err
	}

	result, err := PostfixToInfix(expr)
	if err != nil {
		return err
	}

	err = ch.OutputWriter.Write(result)
	if err != nil {
		return err
	}

	return nil
}
