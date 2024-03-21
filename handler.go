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

type StringInputReader struct {
	Expr string
}

func NewStringInputReader(expr string) *StringInputReader {
	return &StringInputReader{Expr: expr}
}

func (ir *StringInputReader) Read() (string, error) {
	return ir.Expr, nil
}

type FileInputReader struct {
	File io.Reader
}

func NewFileInputReader(file io.Reader) *FileInputReader {
	return &FileInputReader{File: file}
}

func (ir *FileInputReader) Read() (string, error) {
	data, err := io.ReadAll(ir.File)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

type StdOutputWriter struct{}

func NewStdOutputWriter() *StdOutputWriter {
	return &StdOutputWriter{}
}

func (ow *StdOutputWriter) Write(result string) error {
	fmt.Println(result)
	return nil
}

type FileOutputWriter struct {
	File io.Writer
}

func NewFileOutputWriter(file io.Writer) *FileOutputWriter {
	return &FileOutputWriter{File: file}
}

func (ow *FileOutputWriter) Write(result string) error {
	_, err := ow.File.Write([]byte(result))
	return err
}