package lab2

import (
	"fmt"
	"io"
)

// Handler містить методи для обробки вхідних даних та запису результатів
type Handler struct {
	Reader InputReader
	Writer OutputWriter
}

// Compute обробляє вхідні дані та записує результат
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