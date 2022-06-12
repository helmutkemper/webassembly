package html

import "fmt"

func ExampleDegrees_String() {
	a := Degrees(-65)
	fmt.Printf("%v", a)

	// output:
	// -65deg
}
