package factoryBrowser

import "github.com/helmutkemper/iotmaker.webassembly/browser/event/battery"

// NewBattery
//
// English:
//
// The BatteryManager interface of the Battery Status API provides information about the system's battery charge level.
//
// Português:
//
// A interface BatteryManager da API de status da bateria fornece informações sobre o nível de carga da bateria
// do sistema.
func NewBattery() (ref *battery.Battery) {
	ref = new(battery.Battery)
	ref.Init()

	return
}
