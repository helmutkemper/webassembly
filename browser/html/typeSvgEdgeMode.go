package html

type SvgEdgeMode string

func (e SvgEdgeMode) String() string {
	return string(e)
}

const (

	// KSvgEdgeModeDuplicate
	//
	// English:
	//
	//  This value indicates that the input image is extended along each of its borders as necessary by duplicating the
	//  color values at the given edge of the input image.
	//
	// Portuguese
	//
	//  Esse valor indica que a imagem de entrada é estendida ao longo de cada uma de suas bordas conforme necessário,
	//  duplicando os valores de cor na borda especificada da imagem de entrada.
	KSvgEdgeModeDuplicate SvgEdgeMode = "duplicate"

	// KSvgEdgeModeWrap
	//
	// English:
	//
	//  This value indicates that the input image is extended by taking the color values from the opposite edge of the
	//  image.
	//
	// Portuguese
	//
	//  Este valor indica que a imagem de entrada é estendida tomando os valores de cor da borda oposta da imagem.
	KSvgEdgeModeWrap SvgEdgeMode = "wrap"

	// KSvgEdgeModeNone
	//
	// English:
	//
	//  This value indicates that the input image is extended with pixel values of zero for R, G, B and A.
	//
	// Portuguese
	//
	//  Este valor indica que a imagem de entrada é estendida com valores de pixel de zero para R, G, B e A.
	KSvgEdgeModeNone SvgEdgeMode = "none"
)
