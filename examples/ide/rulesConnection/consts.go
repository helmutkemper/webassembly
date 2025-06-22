package rulesConnection

import (
	"fmt"
	"github.com/helmutkemper/webassembly/examples/ide/rulesDensity"
)

const (
	// KWidth Connection width
	KWidth = 6

	// KHeight Connection height
	KHeight = 4

	// KWidthArea Connection width, mouse área.
	//
	//  Creates a larger area to facilitate the click by the user
	KWidthArea = KWidth + 6

	// KHeightArea Connection height, mouse área.
	//
	//  Creates a larger area to facilitate the click by the user
	KHeightArea = KHeight + 6

	KConnectionPrefix = "connection"
)

func GetPathDraw(x, y rulesDensity.Density) (path []string) {
	return []string{
		fmt.Sprintf("M %v %v", x, y),
		fmt.Sprintf("l %v 0", rulesDensity.Density(KWidth)),
		fmt.Sprintf("l 0 %v", rulesDensity.Density(KHeight)),
		fmt.Sprintf("l -%v 0", rulesDensity.Density(KWidth)),
		fmt.Sprintf("l 0 -%v", rulesDensity.Density(KHeight)),
	}
}

func GetPathAreaDraw(x, y rulesDensity.Density) (path []string) {
	return []string{
		fmt.Sprintf("M %v %v", x-rulesDensity.Density(KWidthArea-KWidth)/2, y-rulesDensity.Density(KHeightArea-KHeight)/2),
		fmt.Sprintf("l %v 0", rulesDensity.Density(KWidthArea)),
		fmt.Sprintf("l 0 %v", rulesDensity.Density(KHeightArea)),
		fmt.Sprintf("l -%v 0", rulesDensity.Density(KWidthArea)),
		fmt.Sprintf("l 0 -%v", rulesDensity.Density(KHeightArea)),
	}
}
