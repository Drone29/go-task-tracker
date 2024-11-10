package parser

import (
	"os"
	"testing"
)

func TestParser(t *testing.T) {
	var first_func_called bool
	var second_func_called bool

	AddCmd("call-first", func(args []string) {
		first_func_called = true
	})
	AddCmd("call-second", func(args []string) {
		second_func_called = true
	})
	os.Args = []string{os.Args[0], "call-first"}
	Parse()
	if !first_func_called {
		t.Error("Function was not called")
	}
	os.Args = []string{os.Args[0], "call-second"}
	Parse()
	if !second_func_called {
		t.Error("Function was not called")
	}
}
