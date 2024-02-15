package main

import (
	"testing"

	"github.com/harshakp06/test/process"
)

func TestCheckEven(t *testing.T) { // test func should start with capital "T"
	i := 10
	expected := "YESs"

	res := process.CheckEven(i)

	if expected != res {
		t.Errorf("expected: %v, got: %v", expected, res)
	}
}
