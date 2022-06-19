package html

type SvgChannelSelector string

func (e SvgChannelSelector) String() string {
	return string(e)
}

const (
	// KSvgChannelSelectorR
	//
	// English:
	//
	// This keyword specifies that the red color channel of the input image defined in in2 will be used to displace the
	// pixels of the input image defined in in along the x-axis.
	//
	// Português:
	//
	// Esta palavra-chave especifica que o canal de cor vermelha da imagem de entrada definida em in2 será usado para
	// deslocar os pixels da imagem de entrada definida em ao longo do eixo x.
	KSvgChannelSelectorR SvgChannelSelector = "R"

	// KSvgChannelSelectorG
	//
	// English:
	//
	// This keyword specifies that the green color channel of the input image defined in in2 will be used to displace the
	// pixels of the input image defined in in along the x-axis.
	//
	// Português:
	//
	// Esta palavra-chave especifica que o canal de cor verde da imagem de entrada definida em in2 será usado para
	// deslocar os pixels da imagem de entrada definida em ao longo do eixo x.
	KSvgChannelSelectorG SvgChannelSelector = "G"

	// KSvgChannelSelectorB
	//
	// English:
	//
	// This keyword specifies that the blue color channel of the input image defined in in2 will be used to displace the
	// pixels of the input image defined in in along the x-axis.
	//
	// Português:
	//
	// Essa palavra-chave especifica que o canal de cor azul da imagem de entrada definida em in2 será usado para deslocar
	// os pixels da imagem de entrada definida em ao longo do eixo x.
	KSvgChannelSelectorB SvgChannelSelector = "B"

	// KSvgChannelSelectorA
	//
	// English:
	//
	// This keyword specifies that the alpha channel of the input image defined in in2 will be used to displace the pixels
	// of the input image defined in in along the x-axis.
	//
	// Português:
	//
	// Essa palavra-chave especifica que o canal alfa da imagem de entrada definida em in2 será usado para deslocar os
	// pixels da imagem de entrada definida em ao longo do eixo x.
	KSvgChannelSelectorA SvgChannelSelector = "A"
)
