package html

import "syscall/js"

type Compatible interface {
	Get() js.Value
}
