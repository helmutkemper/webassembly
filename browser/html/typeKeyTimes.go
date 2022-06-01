package html

import "strconv"

type KeyTimes struct {
	l []float64
}

func (e *KeyTimes) Add(k float64) {
	if e.l == nil {
		e.l = make([]float64, 0)
	}

	e.l = append(e.l, k)
}

func (e KeyTimes) String() string {
	ret := ""
	for _, keyTime := range e.l {
		ret += strconv.FormatFloat(keyTime, 'g', -1, 64)
		ret += ";"
	}

	l := len(ret) - 1
	return ret[:l]
}
