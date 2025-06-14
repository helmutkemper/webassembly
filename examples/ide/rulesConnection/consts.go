package rulesConnection

import "fmt"

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

func GetPathDraw(x, y int) (path []string) {
	return []string{
		fmt.Sprintf("M %v %v", x, y),
		fmt.Sprintf("l %v 0", KWidth),
		fmt.Sprintf("l 0 %v", KHeight),
		fmt.Sprintf("l -%v 0", KWidth),
		fmt.Sprintf("l 0 -%v", KHeight),
	}
}

func GetPathAreaDraw(x, y int) (path []string) {
	return []string{
		fmt.Sprintf("M %v %v", x-(KWidthArea-KWidth)/2, y-(KHeightArea-KHeight)/2),
		fmt.Sprintf("l %v 0", KWidthArea),
		fmt.Sprintf("l 0 %v", KHeightArea),
		fmt.Sprintf("l -%v 0", KWidthArea),
		fmt.Sprintf("l 0 -%v", KHeightArea),
	}
}
