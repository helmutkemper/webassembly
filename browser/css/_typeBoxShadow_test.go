package css

import (
	"fmt"
	"github.com/helmutkemper/webassembly/platform/factoryColor"
)

func ExampleBoxShadow_Get() {
	shadow := new(BoxShadow)
	shadow.SetColor(factoryColor.NewRed())
	shadow.SetXOffset(10)
	shadow.SetYOffset(20)
	shadow.SetBlur(30)
	shadow.SetSpreadRadius(40)
	shadow.SetInset(KInsetInherit)

	fmt.Printf("%v", shadow)

	// Output:
	// inherit 10px 20px 30px 40px rgba(255, 0, 0, 1.00)
}
