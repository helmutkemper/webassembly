package html

type SvgVisibility string

func (e SvgVisibility) String() string {
	return string(e)
}

const (
	// KSvgVisibilityVisible
	//
	// English:
	//
	// This value indicates that the element will be painted.
	//
	// Português:
	//
	// Este valor indica que o elemento será pintado.
	KSvgVisibilityVisible SvgVisibility = "visible"

	// KSvgVisibilityHidden
	//
	// English:
	//
	// This value indicates that the element will not be painted. Though it is still part of the rendering tree, i.e. it
	// may receive pointer events depending on the pointer-events attribute, may receive focus depending on the tabindex
	// attribute, contributes to bounding box calculations and clipping paths, and does affect text layout.
	//
	// Português:
	//
	// Este valor indica que o elemento não será pintado. Embora ainda faça parte da árvore de renderização, ou seja, pode
	// receber eventos de ponteiro dependendo do atributo pointer-events, pode receber foco dependendo do atributo
	// tabindex, contribui para cálculos de caixa delimitadora e caminhos de recorte e afeta o layout do texto.
	KSvgVisibilityHidden SvgVisibility = "hidden"

	// KSvgVisibilityCollapse
	//
	// English:
	//
	// This value is equal to hidden.
	//
	// Português:
	//
	// Este valor é igual a oculto.
	KSvgVisibilityCollapse SvgVisibility = "collapse"
)
