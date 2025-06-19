package textUtil

import (
	"sync"
	"syscall/js"
)

const (
	KFontAwesomeRegular = "FARegular"
	KFontAwesomeSolid   = "FASolid"
)

var onceInjectFontAwesome sync.Once

// InjectFontAwesomeCSS
//
// English:
//
//	Inject the style tag with CSS for Font Awesome automatically.
//
// Português:
//
//	Injeta a tag style com o css para a font awesome de forma automática.
//
//	  Example / Exemplo:
//
//	    func main() {
//	      textUtil.InjectFontAwesomeCSS()
//
//	      fontSize := 16
//	      fontStyle := html.KFontStyleRuleNormal
//	      fontWeight := html.KFontWeightRuleNormal
//	      fontFamily := textUtil.KFontAwesomeRegular
//	      width, height := textUtil.GetTextSize("\uf0ea", fontFamily, fontWeight, fontStyle, fontSize)
//
//	      icon := factoryBrowser.NewTagSvgText().
//	        Text("\uf0ea").
//	        FontSize(fontSize).
//	        FontFamily(fontFamily).
//	        FontWeight(fontWeight).
//	        X(cx - width/2).
//	        Y(cy + height/2 - height/5)
//	    }
func InjectFontAwesomeCSS() {
	onceInjectFontAwesome.Do(func() {
		document := js.Global().Get("document")

		css := `
			@font-face {
				font-family: "FASolid";
				src: url("https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.5.0/webfonts/fa-solid-900.woff2") format("woff2");
				font-style: normal;
			}
			@font-face {
				font-family: "FARegular";
				src: url("https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.5.0/webfonts/fa-regular-400.woff2") format("woff2");
				font-weight: 400;
				font-style: normal;
			}
			@font-face {
				font-family: "FALight";
				src: url("https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.5.0/webfonts/fa-light-400.woff2") format("woff2");
				font-weight: 400;
				font-style: normal;
			}
			@font-face {
				font-family: "FABrands";
				src: url("https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.5.0/webfonts/fa-brands-400.woff2") format("woff2");
				font-weight: 400;
				font-style: normal;
			}
			@font-face {
				font-family: "FADuotone";
				src: url("https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.5.0/webfonts/fa-duotone-900.woff2") format("woff2");
				font-weight: 900;
				font-style: normal;
			}
			@font-face {
				font-family: "FASharp";
				src: url("https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.5.0/webfonts/fa-sharp-solid-900.woff2") format("woff2");
				font-weight: 900;
				font-style: normal;
			}
			@font-face {
				font-family: "FASharpDuotone";
				src: url("https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.5.0/webfonts/fa-sharp-duotone-900.woff2") format("woff2");
				font-weight: 900;
				font-style: normal;
			}
		`

		styleEl := document.Call("createElement", "style")
		styleEl.Set("type", "text/css")
		styleEl.Set("textContent", css)

		document.Get("head").Call("appendChild", styleEl)
	})
}
