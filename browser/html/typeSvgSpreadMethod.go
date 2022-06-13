package html

type SvgSpreadMethod string

func (e SvgSpreadMethod) String() string {
	return string(e)
}

const (
	// KSvgSpreadMethodPad
	//
	// English:
	//
	// This value indicates that the final color of the gradient fills the shape beyond the gradient's edges.
	//
	// Português:
	//
	// Esse valor indica que a cor final do gradiente preenche a forma além das bordas do gradiente.
	KSvgSpreadMethodPad SvgSpreadMethod = "pad"

	// KSvgSpreadMethodReflect
	//
	// English:
	//
	// This value indicates that the gradient repeats in reverse beyond its edges.
	//
	// Português:
	//
	// Este valor indica que o gradiente se repete em sentido inverso além de suas bordas.
	KSvgSpreadMethodReflect SvgSpreadMethod = "reflect"

	// KSvgSpreadMethodRepeat
	//
	// English:
	//
	// This value indicates that the gradient repeats in reverse beyond its edges.
	//
	// Português:
	//
	// Este valor indica que o gradiente se repete em sentido inverso além de suas bordas.
	KSvgSpreadMethodRepeat SvgSpreadMethod = "repeat"
)
