package gimmock

import (
	"fmt"
	"strings"
)

type UnexpectedCall struct {
	Interface, Method string
	Args              []interface{}
}

func (e *UnexpectedCall) String() string {
	args := []string{}
	for _, arg := range e.Args {
		args = append(args, fmt.Sprintf("%#v", arg))
	}
	return fmt.Sprintf("Unexpected call %s.%s(%s)",
		e.Interface, e.Method, strings.Join(args, ", "))
}

type WrongArgument struct {
	Interface, Method string
	Idx               int
	Got, Want         interface{}
}

func (e *WrongArgument) String() string {
	return fmt.Sprintf("Wrong argument %d to %s.%s: got %#v, want %#v",
		e.Idx, e.Interface, e.Method, e.Got, e.Want)
}
