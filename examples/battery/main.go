//go:build js

package main

import (
	"fmt"
	"github.com/helmutkemper/webassembly/browser/event/battery"
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
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
				text += fmt.Sprintf("event name: %v<br>", data.EventName)
				text += fmt.Sprintf("level: %v<br>", data.Level)
				text += fmt.Sprintf("charging: %v<br>", data.Charging)
				text += fmt.Sprintf("charging time: %v<br>", data.ChargingTime)
				text += fmt.Sprintf("discharging time: %v<br>", data.DischargingTime)
				div.Html(text)
			}
		}
	}()

	bat := factoryBrowser.NewBattery()
	bat.AddListenerChargingChange(batteryEvent)
	bat.AddListenerDischargingTimeChange(batteryEvent)
	bat.AddListenerChargingTimeChange(batteryEvent)
	bat.AddListenerLevelChange(batteryEvent)

	batteryEvent <- bat.Now()

	done := make(chan struct{})
	<-done
}
