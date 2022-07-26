//go:build js

package main

import (
	"fmt"
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/browser/geolocation"
	"github.com/helmutkemper/iotmaker.webassembly/browser/html"
)

func main() {

	tagText := &html.TagSvgText{}
	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 400, 200}).Append(
		factoryBrowser.NewTagSvgText().X(5).Y(20).FontSize(12).Reference(&tagText).Html("Carregando..."),
	)

	stage.Append(s1)

	var g geolocation.Geolocation
	var coordinate = make(chan geolocation.Coordinate)
	g.GetPosition(&coordinate)
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
