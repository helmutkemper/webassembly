package html

type SvgTypeTurbulence string

func (e SvgTypeTurbulence) String() string {
	return string(e)
}

const (
	KSvgTypeTurbulenceFractalNoise SvgTypeTurbulence = "fractalNoise"
	KSvgTypeTurbulenceTurbulence   SvgTypeTurbulence = "turbulence"
)
