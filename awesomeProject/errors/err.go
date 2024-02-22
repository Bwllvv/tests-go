package main

import "fmt"

type error interface {
	Error() string
}

type apperror struct {
	err    error
	custom string
	Field  int
}

func (e *apperror) Error() string {
	return e.err.Error()
}

func main() {
	err := m()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func m() error {
	return &apperror{
		err:    fmt.Errorf("my error"),
		custom: "value here",
		Field:  89,
	}
}
