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

// ObjectTransformToLog
//
// English:
//
//	Logs the current transformation matrix of the given SVG element, including
//	inherited transforms, by reading its CTM (Current Transformation Matrix) and
//	printing each component with descriptive labels.
//
// Português:
//
//	Exibe no log a matriz de transformação atual do elemento SVG fornecido,
//	incluindo transformações herdadas, obtendo sua CTM (Current Transformation Matrix)
//	e imprimindo cada componente com rótulos descritivos.
func ObjectTransformToLog(label string, value js.Value) {
	// Retrieve the current transformation matrix (CTM)
	// Recupera a matriz de transformação atual (CTM)
	matrix := value.Call("getCTM")

	// If no matrix is found, log and exit
	// Se nenhuma matriz for encontrada, exibe log e retorna
	if matrix.IsNull() || matrix.IsUndefined() {
		log.Printf("%v No transformation matrix found", label)
		return
	}

	// Log each component of the transformation matrix
	// Exibe cada componente da matriz de transformação
	log.Printf("%v Transform Matrix:", label)
	log.Printf("  a (scaleX) = %f", matrix.Get("a").Float()) // horizontal scale / escala horizontal
	log.Printf("  b (skewY)  = %f", matrix.Get("b").Float()) // vertical skew / cisalhamento vertical
	log.Printf("  c (skewX)  = %f", matrix.Get("c").Float()) // horizontal skew / cisalhamento horizontal
	log.Printf("  d (scaleY) = %f", matrix.Get("d").Float()) // vertical scale / escala vertical
	log.Printf("  e (transX) = %f", matrix.Get("e").Float()) // horizontal translation / translação horizontal
	log.Printf("  f (transY) = %f", matrix.Get("f").Float()) // vertical translation / translação vertical
}
