package canvas

import (
	iotmaker_platform_textMetrics "github.com/helmutkemper/iotmaker.platform.textMetrics"
	"syscall/js"
)

// en: Returns a struct TextMetrics that contains the width of the specified text
//     text: The text to be measured
//
// pt_br: Retorna o struct TextMetrics com os dados de comprimento do texto
//     text: Texto a ser medido
func (el *Canvas) MeasureText(text string) iotmaker_platform_textMetrics.TextMetrics {

	var jsMetric js.Value
	jsMetric = el.SelfContext.Call("measureText", text)

	ret := iotmaker_platform_textMetrics.TextMetrics{}
	ret.Width = jsMetric.Get("width").Float()

	if jsMetric.Get("actualBoundingBoxLeft") != js.Undefined() {
		ret.ActualBoundingBoxLeft = jsMetric.Get("actualBoundingBoxLeft").Float()
	} else {
		ret.ActualBoundingBoxLeft = -1
	}

	if jsMetric.Get("actualBoundingBoxRight") != js.Undefined() {
		ret.ActualBoundingBoxRight = jsMetric.Get("actualBoundingBoxRight").Float()
	} else {
		ret.ActualBoundingBoxRight = -1
	}

	if jsMetric.Get("fontBoundingBoxAscent") != js.Undefined() {
		ret.FontBoundingBoxAscent = jsMetric.Get("fontBoundingBoxAscent").Float()
	} else {
		ret.FontBoundingBoxAscent = -1
	}

	if jsMetric.Get("fontBoundingBoxDescent") != js.Undefined() {
		ret.FontBoundingBoxDescent = jsMetric.Get("fontBoundingBoxDescent").Float()
	} else {
		ret.FontBoundingBoxDescent = -1
	}

	if jsMetric.Get("actualBoundingBoxAscent") != js.Undefined() {
		ret.ActualBoundingBoxAscent = jsMetric.Get("actualBoundingBoxAscent").Float()
	} else {
		ret.ActualBoundingBoxAscent = -1
	}

	if jsMetric.Get("actualBoundingBoxDescent") != js.Undefined() {
		ret.ActualBoundingBoxDescent = jsMetric.Get("actualBoundingBoxDescent").Float()
	} else {
		ret.ActualBoundingBoxDescent = -1
	}

	if jsMetric.Get("emHeightAscent") != js.Undefined() {
		ret.EmHeightAscent = jsMetric.Get("emHeightAscent").Float()
	} else {
		ret.EmHeightAscent = -1
	}

	if jsMetric.Get("emHeightDescent") != js.Undefined() {
		ret.EmHeightDescent = jsMetric.Get("emHeightDescent").Float()
	} else {
		ret.EmHeightDescent = -1
	}

	if jsMetric.Get("hangingBaseline") != js.Undefined() {
		ret.HangingBaseline = jsMetric.Get("hangingBaseline").Float()
	} else {
		ret.HangingBaseline = -1
	}

	if jsMetric.Get("alphabeticBaseline") != js.Undefined() {
		ret.AlphabeticBaseline = jsMetric.Get("alphabeticBaseline").Float()
	} else {
		ret.AlphabeticBaseline = -1
	}

	if jsMetric.Get("ideographicBaseline") != js.Undefined() {
		ret.IdeographicBaseline = jsMetric.Get("ideographicBaseline").Float()
	} else {
		ret.IdeographicBaseline = -1
	}

	return ret
}
