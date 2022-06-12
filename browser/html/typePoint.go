package html

import (
	"fmt"
)

type Point struct {
	X float64
	Y float64
}

func (e Point) String() string {
	return fmt.Sprintf("%v,%v", e.X, e.Y)
}
