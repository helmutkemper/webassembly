package html

type SvgTypeFeColorMatrix string

func (e SvgTypeFeColorMatrix) String() string {
	return string(e)
}

const (
	KSvgTypeFeColorMatrixMatrix           SvgTypeFeColorMatrix = "matrix"
	KSvgTypeFeColorMatrixSaturate         SvgTypeFeColorMatrix = "saturate"
	KSvgTypeFeColorMatrixHueRotate        SvgTypeFeColorMatrix = "hueRotate"
	KSvgTypeFeColorMatrixLuminanceToAlpha SvgTypeFeColorMatrix = "luminanceToAlpha"
)
