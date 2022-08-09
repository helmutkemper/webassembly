//go:build js

package main

import (
	"fmt"
	"github.com/helmutkemper/iotmaker.webassembly/browser/event/battery"
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
)

func main() {

	div := factoryBrowser.NewTagDiv()

	stage := factoryBrowser.NewStage()
	stage.Append(div)

	var batteryEvent = make(chan battery.Data)

	go func() {
		for {
			select {
			case data := <-batteryEvent:
				text := ""
				text += fmt.Sprintf("event name: %v", data.EventName)
				text += "<br>"
				text += fmt.Sprintf("level: %v", data.Level)
				text += "<br>"
				text += fmt.Sprintf("charging: %v", data.Charging)
				text += "<br>"
				text += fmt.Sprintf("charging time: %v", data.ChargingTime)
				text += "<br>"
				text += fmt.Sprintf("discharging time: %v", data.DischargingTime)
				div.Html(text)
			}
		}
	}()

	bat := factoryBrowser.NewBattery()
	bat.AddListenerChargingChange(&batteryEvent)
	bat.AddListenerDischargingTimeChange(&batteryEvent)
	bat.AddListenerChargingTimeChange(&batteryEvent)
	bat.AddListenerLevelChange(&batteryEvent)

	batteryEvent <- bat.Now()

	done := make(chan struct{})
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
