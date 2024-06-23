package components

import "github.com/helmutkemper/iotmaker.webassembly/browser/html"

type Button struct {
	__value     any
	__label     any
	__buttonTag *html.TagInputButton `wasmPanel:"type:inputTagButton"`
}

func (e *Button) init() {
	if e.__value != nil {
		e.value(e.__value)
		e.__value = nil
	}
}

func (e *Button) value(value any) {
	e.__buttonTag.Value(value)
}

// Value Sets the value of the component.
func (e *Button) Value(value any) {
	if e.__buttonTag == nil {
		e.__value = value
		return
	}

	e.value(value)
}
