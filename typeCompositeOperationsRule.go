package iotmaker_platform_webbrowser

type CanvasCompositeOperationsRule int

var CanvasCompositeOperationsRules = [...]string{
	"",
	"source-over",
	"source-atop",
	"source-in",
	"source-out",
	"destination-over",
	"destination-atop",
	"destination-in",
	"destination-out",
	"lighter",
	"copy",
	"xor",
}

func (el CanvasCompositeOperationsRule) String() string {
	return CanvasCompositeOperationsRules[el]
}

const (
	// en: Default. Displays the source image over the destination image
	KCompositeOperationsRuleSourceOver CanvasCompositeOperationsRule = iota + 1

	// en: Displays the source image on top of the destination image. The part of the source image that is outside the
	// destination image is not shown
	KCompositeOperationsRuleSourceAtop

	// en: Displays the source image in to the destination image. Only the part of the source image that is INSIDE the
	// destination image is shown, and the destination image is transparent
	KCompositeOperationsRuleSourceIn

	// en: Displays the source image out of the destination image. Only the part of the source image that is OUTSIDE the
	// destination image is shown, and the destination image is transparent
	KCompositeOperationsRuleSourceOut

	// en: Displays the destination image over the source image
	KCompositeOperationsRuleDestinationOver

	// en: Displays the destination image on top of the source image. The part of the destination image that is outside
	// the source image is not shown
	KCompositeOperationsRuleDestinationAtop

	// en: Displays the destination image in to the source image. Only the part of the destination image that is INSIDE
	// the source image is shown, and the source image is transparent
	KCompositeOperationsRuleDestinationIn

	// en: Displays the destination image out of the source image. Only the part of the destination image that is OUTSIDE
	// the source image is shown, and the source image is transparent
	KCompositeOperationsRuleDestinationOut

	// en: Displays the source image + the destination image
	KCompositeOperationsRuleLighter

	// en: Displays the source image. The destination image is ignored
	KCompositeOperationsRuleCopy

	// en: The source image is combined by using an exclusive OR with the destination image
	KCompositeOperationsRuleXor
)
