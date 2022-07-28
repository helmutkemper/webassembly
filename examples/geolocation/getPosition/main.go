//go:build js

package main

import (
	"fmt"
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/browser/geolocation"
	"github.com/helmutkemper/iotmaker.webassembly/browser/html"
	"github.com/helmutkemper/iotmaker.webassembly/browser/media"
)

func main() {

	stage := factoryBrowser.NewStage()

	tagText := &html.TagSvgText{}
	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 400, 200}).Append(
		factoryBrowser.NewTagSvgText().X(5).Y(20).FontSize(12).Reference(&tagText).Html("Loading..."),
	)

	stage.Append(s1)

	var geo geolocation.Geolocation
	var coordinate = make(chan geolocation.Coordinate)
	geo.GetPosition(&coordinate)
	go func() {
		select {
		case converted := <-coordinate:
			text := ""
			text += fmt.Sprintf("<tspan x='5' y='%v'>Latitude: %v</tspan>", 25, converted.Latitude)
			text += fmt.Sprintf("<tspan x='5' y='%v'>Longitude: %v</tspan>", 35, converted.Longitude)
			text += fmt.Sprintf("<tspan x='5' y='%v'>Accuracy: %v meters</tspan>", 45, converted.Accuracy)
			text += fmt.Sprintf("<tspan x='5' y='%v'>Error: %v</tspan>", 55, converted.ErrorMessage)

			tagText.Html(text)
		}
	}()

	d := media.Devices{}
	d.EnumerateDevices()

	done := make(chan struct{}, 0)
	<-done
}

//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
