package css

import (
	"fmt"
	"github.com/helmutkemper/webassembly/platform/factoryColor"
)

func ExampleShadow_String() {
	shadow1 := new(BoxShadow)
	shadow1.SetColor(factoryColor.NewRed())
	shadow1.SetXOffset(10)
	shadow1.SetYOffset(20)
	shadow1.SetBlur(30)
	shadow1.SetSpreadRadius(40)
	shadow1.SetInset(KInsetInherit)

	shadow2 := new(BoxShadow)
	shadow2.SetColor(factoryColor.NewRed())
	shadow2.SetXOffset(10)
	shadow2.SetYOffset(20)
	shadow2.SetBlur(30)
	shadow2.SetSpreadRadius(40)
	shadow2.SetInset(KInsetInitial)

	shadow3 := new(BoxShadow)
	shadow3.SetColor(factoryColor.NewRed())
	shadow3.SetXOffset(10)
	shadow3.SetYOffset(20)
	shadow3.SetBlur(30)
	shadow3.SetSpreadRadius(40)
	shadow3.SetInset(KInsetRevertLayer)

	shadow := new(Shadow)
	shadow.Add(*shadow1)
	shadow.Add(*shadow2)
	shadow.Add(*shadow3)

	fmt.Printf("%v", shadow.String())

	// Output:
	// inherit 10px 20px 30px 40px rgba(255, 0, 0, 1.00) ,
	// initial 10px 20px 30px 40px rgba(255, 0, 0, 1.00) ,
	// revert-layer 10px 20px 30px 40px rgba(255, 0, 0, 1.00) ;
}
