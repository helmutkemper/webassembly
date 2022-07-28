//go:build js

package main

import (
	"fmt"
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/browser/html"
	"github.com/helmutkemper/iotmaker.webassembly/browser/media"
)

func main() {

	stage := factoryBrowser.NewStage()

	tagText := &html.TagSvgText{}
	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 400, 200}).Append(
		factoryBrowser.NewTagSvgText().X(5).Y(20).FontSize(10).Reference(&tagText).Html("Loading..."),
	)

	stage.Append(s1)

	text := ""
	textY := 20
	devicces := media.Devices{}
	listOfDevices, err := devicces.EnumerateDevices()
	if err != nil {
		text += fmt.Sprintf("<tspan x='5' y='%v'>Error: %v</tspan>", textY, err.Error())
	} else {
		for k, v := range listOfDevices {
			textY += k * 40
			text += fmt.Sprintf("<tspan x='5' y='%v'>Kind: %v</tspan>", textY+10, v.Kind)
			text += fmt.Sprintf("<tspan x='5' y='%v'>group id: %v</tspan>", textY+20, v.GroupId)
			text += fmt.Sprintf("<tspan x='5' y='%v'>label: %v</tspan>", textY+30, v.Label)
			text += fmt.Sprintf("<tspan x='5' y='%v'>device id: %v</tspan>", textY+40, v.DeviceId)
		}
	}
	tagText.Html(text)

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
