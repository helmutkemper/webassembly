package canvas

import (
	iotmaker_platform_textMetrics "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.textMetrics"
	"syscall/js"
)

// MeasureText
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

	if jsMetric.Get("actualBoundingBoxLeft").IsUndefined() == false {
		ret.ActualBoundingBoxLeft = jsMetric.Get("actualBoundingBoxLeft").Float()
	} else {
		ret.ActualBoundingBoxLeft = -1
	}

	if jsMetric.Get("actualBoundingBoxRight").IsUndefined() == false {
		ret.ActualBoundingBoxRight = jsMetric.Get("actualBoundingBoxRight").Float()
	} else {
		ret.ActualBoundingBoxRight = -1
	}

	if jsMetric.Get("fontBoundingBoxAscent").IsUndefined() == false {
		ret.FontBoundingBoxAscent = jsMetric.Get("fontBoundingBoxAscent").Float()
	} else {
		ret.FontBoundingBoxAscent = -1
	}

	if jsMetric.Get("fontBoundingBoxDescent").IsUndefined() == false {
		ret.FontBoundingBoxDescent = jsMetric.Get("fontBoundingBoxDescent").Float()
	} else {
		ret.FontBoundingBoxDescent = -1
	}

	if jsMetric.Get("actualBoundingBoxAscent").IsUndefined() == false {
		ret.ActualBoundingBoxAscent = jsMetric.Get("actualBoundingBoxAscent").Float()
	} else {
		ret.ActualBoundingBoxAscent = -1
	}

	if jsMetric.Get("actualBoundingBoxDescent").IsUndefined() == false {
		ret.ActualBoundingBoxDescent = jsMetric.Get("actualBoundingBoxDescent").Float()
	} else {
		ret.ActualBoundingBoxDescent = -1
	}

	if jsMetric.Get("emHeightAscent").IsUndefined() == false {
		ret.EmHeightAscent = jsMetric.Get("emHeightAscent").Float()
	} else {
		ret.EmHeightAscent = -1
	}

	if jsMetric.Get("emHeightDescent").IsUndefined() == false {
		ret.EmHeightDescent = jsMetric.Get("emHeightDescent").Float()
	} else {
		ret.EmHeightDescent = -1
	}

	if jsMetric.Get("hangingBaseline").IsUndefined() == false {
		ret.HangingBaseline = jsMetric.Get("hangingBaseline").Float()
	} else {
		ret.HangingBaseline = -1
	}

	if jsMetric.Get("alphabeticBaseline").IsUndefined() == false {
		ret.AlphabeticBaseline = jsMetric.Get("alphabeticBaseline").Float()
	} else {
		ret.AlphabeticBaseline = -1
	}

	if jsMetric.Get("ideographicBaseline").IsUndefined() == false {
		ret.IdeographicBaseline = jsMetric.Get("ideographicBaseline").Float()
	} else {
		ret.IdeographicBaseline = -1
	}

	return ret
}
