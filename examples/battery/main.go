//go:build js

package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/event/battery"
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"log"
)

func main() {
	done := make(chan struct{})
	var batteryEvent = make(chan battery.Data, 1)

	bat := factoryBrowser.NewBattery()
	bat.AddListenerChargingChange(&batteryEvent)
	bat.AddListenerDischargingTimeChange(&batteryEvent)
	bat.AddListenerChargingTimeChange(&batteryEvent)
	bat.AddListenerLevelChange(&batteryEvent)

	go func() {
		for {
			select {
			case data := <-batteryEvent:
				log.Printf("event name: %v", data.EventName)
				log.Printf("level: %v", data.Level)
				log.Printf("charging: %v", data.Charging)
				log.Printf("charging time: %v", data.ChargingTime)
				log.Printf("discharging time: %v", data.DischargingTime)
			}
		}
	}()
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
