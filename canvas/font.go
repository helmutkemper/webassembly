package canvas

import "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/font"

// en: Sets the current font properties for text content
//
// pt_br: Define as propriedades da fonte atual
func (el *Canvas) Font(font font.Font) {
	el.SelfContext.Set("font", font.String())
}
