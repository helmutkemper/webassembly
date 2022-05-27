package html

type SvgDirection string

func (e SvgDirection) String() string {
	return string(e)
}

const (
	// KSvgDirectionLtr
	//
	// English:
	//
	//  Default. Left-to-right text direction.
	//
	// Português:
	//
	//  Padrão. Direção do texto da esquerda para a direita.
	KSvgDirectionLtr SvgDirection = "ltr"

	// KSvgDirectionRtl
	//
	// English:
	//
	//  Right-to-left text direction.
	//
	// Português:
	//
	//  Direção do texto da direita para a esquerda.
	KSvgDirectionRtl SvgDirection = "rtl"
)
