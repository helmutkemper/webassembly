package html

type CompositeOperationsRule string

func (e CompositeOperationsRule) String() string {
	return string(e)
}

const (
	// KCompositeOperationsRuleSourceOver
	//
	// English:
	//
	//  (Default) Displays the source image over the destination image.
	//
	// Português:
	//
	//  (Padrão) Exibe a imagem de origem sobre a imagem de destino.
	KCompositeOperationsRuleSourceOver CompositeOperationsRule = "source-over"

	// KCompositeOperationsRuleSourceAtop
	//
	// English:
	//
	//  Displays the source image on top of the destination image. The part of the source image that is
	//  outside the destination image is not shown.
	//
	// Português:
	//
	//  Exibe a imagem de origem sobre a imagem de destino. A parte da imagem de origem que está fora
	//  da imagem de destino não é mostrada.
	KCompositeOperationsRuleSourceAtop CompositeOperationsRule = "source-atop"

	// KCompositeOperationsRuleSourceIn
	//
	// English:
	//
	//  Displays the source image in to the destination image. Only the part of the source image that is
	//  INSIDE the destination image is shown, and the destination image is transparent.
	//
	// Português:
	//
	//  Exibe a imagem de origem na imagem de destino. Apenas a parte da imagem de origem que está
	//  DENTRO da imagem de destino é mostrada, e a imagem de destino é transparente.
	KCompositeOperationsRuleSourceIn CompositeOperationsRule = "source-in"

	// KCompositeOperationsRuleSourceOut
	//
	// English:
	//
	//  Displays the source image out of the destination image. Only the part of the source image that
	//  is OUTSIDE the destination image is shown, and the destination image is transparent.
	//
	// Português:
	//
	//  Exibe a imagem de origem fora da imagem de destino. Apenas a parte da imagem de origem que está
	//  FORA da imagem de destino é mostrada, e a imagem de destino é transparente.
	KCompositeOperationsRuleSourceOut CompositeOperationsRule = "source-out"

	// KCompositeOperationsRuleDestinationOver
	//
	// English:
	//
	//  Displays the destination image over the source image.
	//
	// Português:
	//
	//  Exibe a imagem de destino sobre a imagem de origem.
	KCompositeOperationsRuleDestinationOver CompositeOperationsRule = "destination-over"

	// KCompositeOperationsRuleDestinationAtop
	//
	// English:
	//
	//  Displays the destination image on top of the source image. The part of the destination image
	//  that is outside the source image is not shown.
	//
	// Português:
	//
	//  Exibe a imagem de destino sobre a imagem de origem. A parte da imagem de destino que está fora
	//  da imagem de origem não é mostrada.
	KCompositeOperationsRuleDestinationAtop CompositeOperationsRule = "destination-atop"

	// KCompositeOperationsRuleDestinationIn
	//
	// English:
	//
	//  Displays the destination image in to the source image. Only the part of the destination image
	//  that is INSIDE the source image is shown, and the source image is transparent.
	//
	// Português:
	//
	//  Exibe a imagem de destino na imagem de origem. Apenas a parte da imagem de destino que está
	//  DENTRO da imagem de origem é mostrada, e a imagem de origem é transparente.
	KCompositeOperationsRuleDestinationIn CompositeOperationsRule = "destination-in"

	// KCompositeOperationsRuleDestinationOut
	//
	// English:
	//
	//  Displays the destination image out of the source image. Only the part of the destination image
	//  that is OUTSIDE the source image is shown, and the source image is transparent.
	//
	// Português:
	//
	//  Exibe a imagem de destino fora da imagem de origem. Apenas a parte da imagem de destino que está
	//  FORA da imagem de origem é mostrada, e a imagem de origem é transparente.
	KCompositeOperationsRuleDestinationOut CompositeOperationsRule = "destination-out"

	// KCompositeOperationsRuleLighter
	//
	// English:
	//
	//  Displays the source image + the destination image.
	//
	// Português:
	//
	//  Exibe a imagem de origem + a imagem de destino.
	KCompositeOperationsRuleLighter CompositeOperationsRule = "lighter"

	// KCompositeOperationsRuleCopy
	//
	// English:
	//
	//  Displays the source image. The destination image is ignored.
	//
	// Português:
	//
	//  Exibe a imagem de origem. A imagem de destino é ignorada.
	KCompositeOperationsRuleCopy CompositeOperationsRule = "copy"

	// KCompositeOperationsRuleXor
	//
	// English:
	//
	//  The source image is combined by using an exclusive OR with the destination image.
	//
	// Português:
	//
	//  A imagem de origem é combinada usando um OR exclusivo com a imagem de destino.
	KCompositeOperationsRuleXor CompositeOperationsRule = "xor"
)
