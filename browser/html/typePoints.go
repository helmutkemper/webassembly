package html

type Points []Point

func (e *Points) Add(x, y float64) {
	if *e == nil {
		*e = make([]Point, 0)
	}

	*e = append(*e, Point{X: x, Y: y})
}

func (e Points) String() string {
	ret := ""
	for _, v := range e {
		ret += v.String() + " "
	}
	length := len(ret) - 1

	return ret[:length]
}
