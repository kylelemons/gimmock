package iface

import (
	"os"
)

type Iface interface {
	Boolean() bool
	String() string
	Convert(int) string
	Operation(string, int) ([]string, os.Error)
	AnotherOp([10]string) map[int]string
	Sink(float64)
}
