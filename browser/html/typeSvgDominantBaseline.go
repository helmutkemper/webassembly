package html

type SvgDominantBaseline string

func (e SvgDominantBaseline) String() string {
	return string(e)
}

const (
	// KSvgDominantBaselineAuto
	//
	// English:
	//
	//  If this property occurs on a <text> element, then the computed value depends on the value of the writing-mode
	//  attribute.
	//
	// If the writing-mode is horizontal, then the value of the dominant-baseline component is alphabetic. Otherwise, if
	// the writing-mode is vertical, then the value of the dominant-baseline component is central.
	//
	// If this property occurs on a <tspan>, <tref>, <altGlyph>, or <textPath> element, then the dominant-baseline and
	// the baseline-table components remain the same as those of the parent text content element.
	//
	// If the computed baseline-shift value actually shifts the baseline, then the baseline-table font-size component is
	// set to the value of the font-size attribute on the element on which the dominant-baseline attribute occurs,
	// otherwise the baseline-table font-size remains the same as that of the element.
	//
	// If there is no parent text content element, the scaled-baseline-table value is constructed as above for <text>
	// elements.
	//
	// Português:
	//
	// Se esta propriedade ocorrer em um elemento <text>, então o valor calculado depende do valor do atributo write-mode.
	//
	//If the writing-mode is horizontal, then the value of the dominant-baseline component is alphabetic. Otherwise, if the writing-mode is vertical, then the value of the dominant-baseline component is central.
	//
	//If this property occurs on a <tspan>, <tref>, <altGlyph>, or <textPath> element, then the dominant-baseline and the baseline-table components remain the same as those of the parent text content element.
	//
	//If the computed baseline-shift value actually shifts the baseline, then the baseline-table font-size component is set to the value of the font-size attribute on the element on which the dominant-baseline attribute occurs, otherwise the baseline-table font-size remains the same as that of the element.
	//
	//If there is no parent text content element, the scaled-baseline-table value is constructed as above for <text> elements.
	//
	//
	//
	//
	//
	//
	KSvgDominantBaselineAuto       SvgDominantBaseline = "auto"
	KSvgDominantBaselineTextBottom SvgDominantBaseline = "text-bottom"
	KSvgDominantBaseline           SvgDominantBaseline = ""
	KSvgDominantBaseline           SvgDominantBaseline = ""
	KSvgDominantBaseline           SvgDominantBaseline = ""
	KSvgDominantBaseline           SvgDominantBaseline = ""
	KSvgDominantBaseline           SvgDominantBaseline = ""
	KSvgDominantBaseline           SvgDominantBaseline = ""
	KSvgDominantBaseline           SvgDominantBaseline = ""
)
