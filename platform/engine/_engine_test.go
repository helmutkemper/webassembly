package engine

import (
	"fmt"
	"testing"
)

func TestSetZIndex(t *testing.T) {
	e := Engine{}
	e.Init()

	e.AddToSystem(func() { fmt.Println("one") })
	two := e.AddToSystem(func() { fmt.Println("two") })
	e.AddToSystem(func() { fmt.Println("three") })

	for _, runnerFunc := range e.funcListToSystem {
		if runnerFunc.f != nil {
			runnerFunc.f()
		}
	}

	e.toFrontGeneric(two, &e.funcListToSystem)

	for _, runnerFunc := range e.funcListToSystem {
		if runnerFunc.f != nil {
			runnerFunc.f()
		}
	}

	// output:
	// one
	// two
	// three
}
