package rulesDensity

import (
	"github.com/helmutkemper/webassembly/utilsBrowser"
	"log"
	"strconv"
)

// density
//
// English:
//
//	Global scale factor based on screen density (default is 1.0).
//
// Português:
//
//	Fator global de escala com base na densidade da tela (padrão é 1.0).
var density float64 = 1.0

// init
//
// English:
//
//	Initializes the density value using the "scale" query string parameter,
//	if available.
//
// Português:
//
//	Inicializa o valor da densidade usando o parâmetro "scale" da query string,
//	se estiver presente.
func init() {
	if scaleStr := utilsBrowser.GetQueryStringParam("scale"); scaleStr != "" {
		scale, err := strconv.ParseInt(scaleStr, 10, 32)
		if err != nil {
			return
		}
		density = float64(scale)
		log.Printf("density: %v", density)
	}
}

// GetDensity
//
// English:
//
//	Returns the current density value.
//
// Português:
//
//	Retorna o valor atual da densidade.
func GetDensity() float64 {
	return density
}

// NewInt
//
// English:
//
//	Creates a new Density object from an integer value (unscaled).
//
// Português:
//
//	Cria um novo objeto Density a partir de um valor inteiro (não escalado).
func NewInt(value int) (p *Density) {
	p = new(Density)
	*p = Density(value)
	return
}

// Convert
//
// English:
//
//	Converts an integer value to a Density value by removing scale effect.
//
// Português:
//
//	Converte um valor inteiro para um valor Density removendo o efeito da escala.
func Convert(value int) (p Density) {
	return Density(float64(value) / float64(density))
}

// Density type
//
// English:
//
//	Custom type to handle values affected by display density.
//
// Português:
//
//	Tipo personalizado para lidar com valores afetados pela densidade da tela.
type Density float64

// GetInt
//
// English:
//
//	Returns the scaled integer value.
//
// Português:
//
//	Retorna o valor inteiro escalado.
func (e Density) GetInt() int {
	return int(float64(e) * density)
}

// GetOriginalInt
//
// English:
//
//	Returns the unscaled original integer value.
//
// Português:
//
//	Retorna o valor original inteiro (sem escala).
func (e Density) GetOriginalInt() int {
	return int(e)
}

// GetFloat
//
// English:
//
//	Returns the scaled float64 value.
//
// Português:
//
//	Retorna o valor float64 escalado.
func (e Density) GetFloat() float64 {
	return float64(e) * density
}

// String
//
// English:
//
//	Returns the scaled value as a string.
//
// Português:
//
//	Retorna o valor escalado como uma string.
func (e Density) String() string {
	return strconv.FormatFloat(float64(e)*density, 'g', -1, 32)
}

// Pixel
//
// English:
//
//	Returns the scaled value formatted as a CSS pixel string (e.g., "32px").
//
// Português:
//
//	Retorna o valor escalado formatado como uma string de pixel CSS (ex: "32px").
func (e Density) Pixel() string {
	return strconv.FormatFloat(float64(e)*density, 'g', -1, 32) + "px"
}
