package lab2

import (
	"fmt"
	"io"
)

type Handler struct {
	Reader InputReader
	Writer OutputWriter
}

func (h *Handler) Compute() error {
	expr, err := h.Reader.Read()
	if err != nil {
		return err
	}

	infixExpr, err := PostfixToInfix(expr)
	if err != nil {
		return err
	}

	err = h.Writer.Write(infixExpr)
	if err != nil {
		return err
	}

	return nil
}