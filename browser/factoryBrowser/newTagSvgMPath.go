package factoryBrowser

import "github.com/helmutkemper/iotmaker.webassembly/browser/html"

// NewTagSvgMPath
//
// English:
//
// The <metadata> SVG element adds metadata to SVG content. Metadata is structured information about data.
// The contents of <metadata> should be elements from other XML namespaces such as RDF, FOAF, etc.
//
// Português:
//
// O elemento SVG <metadata> adiciona metadados ao conteúdo SVG. Metadados são informações estruturadas sobre dados.
// O conteúdo de <metadata> deve ser elementos de outros namespaces XML, como RDF, FOAF, etc.
func NewTagSvgMPath(id string) (ref *html.TagSvgMPath) {
	ref = &html.TagSvgMPath{}
	ref.Init(id)

	return ref
}
