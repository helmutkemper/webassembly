//go:build js

package main

import (
	"fmt"
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/browser/geolocation"
)

func main() {

	div1 := factoryBrowser.NewTagDiv().Html("loading ...")

	stage := factoryBrowser.NewStage()
	stage.Append(div1)

	var coordinate = make(chan geolocation.Coordinate)

	go func() {
		select {
		case converted := <-coordinate:
			text := fmt.Sprintf("Latitude: %v<br>", converted.Latitude)
			text += fmt.Sprintf("Longitude: %v<br>", converted.Longitude)
			text += fmt.Sprintf("Accuracy: %v meters<br>", converted.Accuracy)
			text += fmt.Sprintf("Error: %v<br>", converted.ErrorMessage)
			div1.Html(text)
		}
	}()

	var geo = factoryBrowser.NewGeoLocation()
	geo.GetPosition(&coordinate)

	done := make(chan struct{}, 0)
	<-done
}
