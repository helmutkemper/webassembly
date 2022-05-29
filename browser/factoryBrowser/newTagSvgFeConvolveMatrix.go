package factoryBrowser

import "github.com/helmutkemper/iotmaker.webassembly/browser/html"

func NewTagSvgFeConvolveMatrix(id string) (ref *html.TagSvgFeConvolveMatrix) {
	ref = &html.TagSvgFeConvolveMatrix{}
	ref.Init(id)

	return ref
}
