package factoryBrowser

import (
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/utilsMath"
)

// NewTagSvgImage
//
// English:
//
// The <image> SVG element includes images inside SVG documents. It can display raster image files or other SVG files.
//
// The only image formats SVG software must support are JPEG, PNG, and other SVG files. Animated GIF behavior is
// undefined.
//
// SVG files displayed with <image> are treated as an image: external resources aren't loaded, :visited styles aren't
// applied, and they cannot be interactive. To include dynamic SVG elements, try <use> with an external URL. To include
// SVG files and run scripts inside them, try <object> inside of <foreignObject>.
//
//	Notes:
//	  * The HTML spec defines <image> as a synonym for <img> while parsing HTML. This specific element and its
//	    behavior only apply inside SVG documents or inline SVG.
//
// Português:
//
// O elemento SVG <image> inclui imagens dentro de documentos SVG. Ele pode exibir arquivos de imagem raster ou outros
// arquivos SVG.
//
// Os únicos formatos de imagem que o software SVG deve suportar são JPEG, PNG e outros arquivos SVG. O comportamento
// do GIF animado é indefinido.
//
// Arquivos SVG exibidos com <image> são tratados como uma imagem: recursos externos não são carregados, estilos
// :visited não são aplicados e não podem ser interativos. Para incluir elementos SVG dinâmicos, tente <use> com uma
// URL externa. Para incluir arquivos SVG e executar scripts dentro deles, tente <object> dentro de <foreignObject>.
//
//	Notes:
//	  * The HTML spec defines <image> as a synonym for <img> while parsing HTML. This specific element and its
//	    behavior only apply inside SVG documents or inline SVG.
func NewTagSvgImage() (ref *html.TagSvgImage) {
	ref = &html.TagSvgImage{}
	ref.Init()
	ref.Id(utilsMath.GetUID())

	return ref
}
