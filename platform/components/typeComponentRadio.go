package components

import "github.com/helmutkemper/webassembly/browser/html"

// __radioOnInputEvent faz a captura de dados do event input
type __radioOnInputEvent struct {
	Value string `wasmGet:"value"`
}

type Radio struct {
	__value  any
	__change *__radioOnInputEvent

	__radioTag *html.TagInputRadio `wasmPanel:"type:inputTagRadio"`
}

func (e *Radio) init() {
	if e.__value != nil {
		e.value(e.__value)
		e.__value = nil
	}
}

func (e *Radio) value(value any) {
	e.__radioTag.Value(value)
}

func (e *Radio) Value(value any) {
	if e.__radioTag == nil {
		e.__value = value
		return
	}

	e.value(value)
}
