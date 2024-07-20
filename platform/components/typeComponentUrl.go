package components

import "github.com/helmutkemper/webassembly/browser/html"

// __urlOnInputEvent faz a captura de dados do event input
type __urlOnInputEvent struct {
	Value string `wasmGet:"value"`
}

type Url struct {
	__value  any
	__change *__urlOnInputEvent

	__urlTag *html.TagInputUrl `wasmPanel:"type:inputTagInputUrl"`
}

func (e *Url) init() {
	if e.__value != nil {
		e.value(e.__value)
		e.__value = nil
	}
}

func (e *Url) value(value any) {
	e.__urlTag.Value(value)
}

func (e *Url) Value(value any) {
	if e.__urlTag == nil {
		e.__value = value
		return
	}

	e.value(value)
}
