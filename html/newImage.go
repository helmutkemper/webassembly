package html

import (
	iotmaker_platform_coordinate "github.com/helmutkemper/iotmaker.platform.coordinate"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/document"
	"sync"
	"syscall/js"
)

// Example: html.NewImage(browserDocument, "img", "./kemper_logo.png", 100, 100, false, true, density, densityManager)
func NewImage(parent document.Document, id, source string, width, height int, appendToParent, waitOnLoad bool, density interface{}, iDensity iotmaker_platform_coordinate.IDensity) js.Value {
	wg := sync.WaitGroup{}

	densityCalc := iDensity
	densityCalc.SetDensityFactor(density)

	densityCalc.Set(width)
	width = densityCalc.Int()

	densityCalc.Set(height)
	height = densityCalc.Int()

	element := document.Element{}
	element.SelfElement = parent.SelfDocument.Call("createElement", "img")
	element.SelfElement.Set("id", id)
	element.SelfElement.Set("width", width)
	element.SelfElement.Set("height", height)
	element.SelfElement.Set("src", source)

	if waitOnLoad == true {
		wg.Add(1)
		element.SelfElement.Call("addEventListener", "load", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			wg.Done()
			return nil
		}))

		wg.Wait()
	}

	if appendToParent == true {
		parent.AppendChildToDocumentBody(element.SelfElement)
	}

	return element.SelfElement
}
