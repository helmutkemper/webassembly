package factoryBrowser

import "github.com/helmutkemper/webassembly/browser/html"

// NewTagSvgFeComponentTransfer
//
// English:
//
// The <feComponentTransfer> SVG filter primitive performs color-component-wise remapping of data for each pixel.
//
// It allows operations like brightness adjustment, contrast adjustment, color balance or thresholding.
//
// The calculations are performed on non-premultiplied color values. The colors are modified by changing each channel
// (R, G, B, and A) to the result of what the children <feFuncR>, <feFuncB>, <feFuncG>, and <feFuncA> return.
//
// Português:
//
// A primitiva de filtro SVG <feComponentTransfer> executa o remapeamento de dados por componente de cor para cada
// pixel.
//
// Permite operações como ajuste de brilho, ajuste de contraste, equilíbrio de cores ou limiar.
//
// Os cálculos são executados em valores de cores não pré-multiplicados. As cores são modificadas alterando cada canal
// (R, G, B e A) para o resultado do que os filhos <feFuncR>, <feFuncB>, <feFuncG> e <feFuncA> retornam.
func NewTagSvgFeComponentTransfer() (ref *html.TagSvgFeComponentTransfer) {
	ref = &html.TagSvgFeComponentTransfer{}
	ref.Init()

	return ref
}
