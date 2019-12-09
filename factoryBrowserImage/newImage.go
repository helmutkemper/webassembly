package factoryBrowserImage

import (
	iotmaker_platform_coordinate "github.com/helmutkemper/iotmaker.platform.coordinate"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/document"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/element"
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

	imageElement := element.Element{}
	imageElement.SelfElement = parent.SelfDocument.Call("createElement", "img")
	imageElement.SelfElement.Set("id", id)
	imageElement.SelfElement.Set("width", width)
	imageElement.SelfElement.Set("height", height)
	imageElement.SelfElement.Set("src", source)

	if waitOnLoad == true {
		wg.Add(1)
		imageElement.SelfElement.Call("addEventListener", "load", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			wg.Done()
			return nil
		}))

		wg.Wait()
	}

	if appendToParent == true {
		parent.AppendChildToDocumentBody(imageElement.SelfElement)
	}

	return imageElement.SelfElement
}
