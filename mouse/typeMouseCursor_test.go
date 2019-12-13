package mouse

import "fmt"

func ExampleCursorType_String() {
	var c CursorType = KCursorZoomOut
	fmt.Println(c.String())

	// output:
	// cursor:zoom-out
}
