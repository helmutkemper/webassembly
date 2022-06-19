package html

type SvgUnicodeBidi string

func (e SvgUnicodeBidi) String() string {
	return string(e)
}

const (
	KSvgUnicodeBidiNormal          SvgUnicodeBidi = "normal"
	KSvgUnicodeBidiEmbed           SvgUnicodeBidi = "embed"
	KSvgUnicodeBidiIsolate         SvgUnicodeBidi = "isolate"
	KSvgUnicodeBidiBidiOverride    SvgUnicodeBidi = "bidi-override"
	KSvgUnicodeBidiIsolateOverride SvgUnicodeBidi = "isolate-override"
	KSvgUnicodeBidiPlaintext       SvgUnicodeBidi = "plaintext"
)
