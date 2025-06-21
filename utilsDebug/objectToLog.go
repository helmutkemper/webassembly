package utilsDebug

import (
	"log"
	"syscall/js"
)

// ObjectToLog
//
// English:
//
//	Logs a JavaScript object as a stringified JSON for debugging.
//
// Português:
//
//	Exibe um objeto JavaScript convertido para JSON, útil para depuração.
func ObjectToLog(label string, value js.Value) {
	json := js.Global().Get("JSON")
	str := json.Call("stringify", value)
	log.Printf("%s: %s", label, str.String())
}
