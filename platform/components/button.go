package components

type Button struct {
	label     string
	click     func(label string)
	dbClick   func(label string)
	press     func(label string)
	release   func(label string)
	mouseOver func(label string)
	mouseOut  func(label string)
}

func (e *Button) SetLabel(label string) {
	e.label = label
}

func (e *Button) GetLabel() (label string) {
	return e.label
}

func (e *Button) SetClick(click func(label string)) {
	e.click = click
}

func (e *Button) SetDbClick(dbClick func(label string)) {
	e.dbClick = dbClick
}

func (e *Button) SetPress(press func(label string)) {
	e.press = press
}

func (e *Button) SetRelease(release func(label string)) {
	e.release = release
}

func (e *Button) SetMouseOver(over func(label string)) {
	e.mouseOver = over
}

func (e *Button) SetMouseOut(out func(label string)) {
	e.mouseOut = out
}
