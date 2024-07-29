package factoryBrowser

import (
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/mathUtil"
)

// NewTagFigure
//
// English:
//
// The <figure> HTML element represents self-contained content, potentially with an optional caption,
// which is specified using the <figcaption> element. The figure, its caption, and its contents are
// referenced as a single unit.
//
//	Notes:
//	  * Usually a <figure> is an image, illustration, diagram, code snippet, etc., that is referenced
//	    in the main flow of a document, but that can be moved to another part of the document or to
//	    an appendix without affecting the main flow.
//	  * Being a sectioning root, the outline of the content of the <figure> element is excluded from
//	    the main outline of the document.
//	  * A caption can be associated with the <figure> element by inserting a <figcaption> inside it
//	    (as the first or the last child). The first <figcaption> element found in the figure is
//	    presented as the figure's caption.
//
// Português:
//
// O elemento HTML <figure> representa conteúdo autocontido, potencialmente com uma legenda opcional,
// que é especificada usando o elemento <figcaption>. A figura, sua legenda e seu conteúdo são
// referenciados como uma única unidade.
//
//	Notas:
//	  * Normalmente uma <figura> é uma imagem, ilustração, diagrama, trecho de código, etc., que é
//	    referenciado no fluxo principal de um documento, mas que pode ser movido para outra parte do
//	    documento ou para um apêndice sem afetar o fluxo principal.
//	  * Sendo a seção principal, o contorno do conteúdo do elemento <figure> é excluído do contorno
//	    principal do documento.
//	  * Uma legenda pode ser associada ao elemento <figure> inserindo um <figcaption> dentro dele
//	    (como o primeiro ou o último filho). O primeiro elemento <figcaption> encontrado na figura é
//	    apresentado como legenda da figura.
func NewTagFigure() (ref *html.TagFigure) {
	ref = &html.TagFigure{}
	ref.Init()
	ref.Id(mathUtil.GetUID())

	return ref
}
