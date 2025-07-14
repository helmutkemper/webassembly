package canvas

type Control interface{}

type register struct {
	element []Control
}

func (e *register) Init() {
	e.element = make([]Control, 0)
}
