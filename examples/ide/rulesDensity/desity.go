package rulesDensity

import (
	"github.com/helmutkemper/webassembly/utilsBrowser"
	"log"
	"strconv"
)

var density float64 = 1.0

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

func GetDensity() float64 {
	return density
}

func NewInt(value int) (p *Density) {
	p = new(Density)
	*p = Density(value)
	return
}

func Convert(value int) (p Density) {
	return Density(float64(value) / float64(density))
}

type Density float64

func (e Density) GetInt() int {
	return int(float64(e) * density)
}

func (e Density) GetOriginalInt() int {
	return int(e)
}

func (e Density) GetFloat() float64 {
	return float64(e) * density
}

func (e Density) String() string {
	return strconv.FormatFloat(float64(e)*density, 'g', -1, 32)
}

func (e Density) Pixel() string {
	return strconv.FormatFloat(float64(e)*density, 'g', -1, 32) + "px"
}
