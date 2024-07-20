package components

import "github.com/helmutkemper/webassembly/browser/html"

// __checkboxOnInputEvent faz a captura de dados do event input
type __checkboxOnInputEvent struct {
	Value string `wasmGet:"value"`
}

type Checkbox struct {
	__value  any
	__change *__checkboxOnInputEvent

	__checkboxTag *html.TagInputCheckBox `wasmPanel:"type:inputTagCheckbox"`
}

func (e *Checkbox) init() {
	if e.__value != nil {
		e.value(e.__value)
		e.__value = nil
	}
}

func (e *Checkbox) value(value any) {
	e.__checkboxTag.Value(value)
}

func (e *Checkbox) Value(value any) {
	if e.__checkboxTag == nil {
		e.__value = value
		return
	}

	e.value(value)
}
