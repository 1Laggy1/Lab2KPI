package lab2

import (
    "fmt"
    "io"
)

type ComputeHandler struct {
    InputReader  io.Reader
    OutputWriter io.Writer
}

func (ch *ComputeHandler) Compute() error {
    var input string
    , err := fmt.Fscan(ch.InputReader, &input)
    if err != nil {
        return err
    }

    result, err := PostfixToInfix(input)
    if err != nil {
        return err
    }

    , err = fmt.Fprintln(ch.OutputWriter, result)
    if err != nil {
        return err
    }

    return nil
}
