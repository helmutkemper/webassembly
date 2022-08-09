package battery

// https://developer.mozilla.org/en-US/docs/Web/API/BatteryManager

import (
	"syscall/js"
)

var objectBactery js.Value

// Battery
//
// English:
//
// The BatteryManager interface of the Battery Status API provides information about the system's battery charge level.
//
// Português:
//
// A interface BatteryManager da API de status da bateria fornece informações sobre o nível de carga da bateria
// do sistema.
type Battery struct {
	// fnChargingChange
	//
	// English:
	//
	// Fired when the battery charging state (the charging property) is updated.
	//
	// Português:
	//
	// Acionado quando o estado de carregamento da bateria (a propriedade de carregamento) é atualizado.
	fnChargingChange *js.Func

	// fnChargingTimeChange
	//
	// English:
	//
	// Fired when the battery charging time (the chargingTime property) is updated.
	//
	// Português:
	//
	// Acionado quando o tempo de carregamento da bateria (propriedade loadingTime) é atualizado.
	fnChargingTimeChange *js.Func

	// fnDischargingTimeChange
	//
	// English:
	//
	// Fired when the battery discharging time (the dischargingTime property) is updated.
	//
	// Português:
	//
	// Acionado quando o tempo de descarga da bateria (propriedade dischargingTime) é atualizado.
	fnDischargingTimeChange *js.Func

	// fnLevelChange
	//
	// English:
	//
	// Fired when the battery level (the level property) is updated.
	//
	// Português:
	//
	// Acionado quando o nível da bateria (a propriedade de nível) é atualizado.
	fnLevelChange *js.Func
}

// Now
//
// English:
//
// Returns the device's current battery status.
//
// Português:
//
// Retorna o status atual da bateria do dispositivo.
func (e *Battery) Now() (data Data) {
	var event = Event{}
	event.Object = objectBactery
	data.EventName = "battery now"
	data.This = js.Value{}

	data.Charging = event.GetIsCharging()
	data.ChargingTime = event.GetChargingTime()
	data.DischargingTime = event.GetDischargingTime()
	data.Level = event.GetLevel()

	return
}

// Init
//
// English:
//
// Correctly initialize the battery object.
//
// Português:
//
// Inicializa o objeto battery de forma correta.
func (e *Battery) Init() {
	var wait = make(chan struct{})
	var success js.Func

	success = js.FuncOf(func(this js.Value, args []js.Value) any {
		defer success.Release()

		objectBactery = args[0]
		wait <- struct{}{}

		return nil
	})

	js.Global().Get("navigator").Call("getBattery").Call("then", success)

	// impede eventos de bateria antes do objeto inicializado
	<-wait
}

// AddListenerChargingChange
//
// English:
//
// Adds a battery charging change event listener equivalent to the JavaScript command addEventListener('chargingchange',fn).
//
//	Input:
//	  mouseEvet: pointer to channel battery.Data
//
// Fired when the battery charging state (the charging property) is updated.
//
// Português:
//
// Adiciona um ouvinte de evento de carga da bateria mudou, equivalente ao comando JavaScript addEventListener('chargingchange',fn).
//
//	Entrada:
//	  mouseEvet: ponteiro para o channel battery.Data
//
// Acionado quando o estado de carregamento da bateria (a propriedade de carregamento) é atualizado.
func (e *Battery) AddListenerChargingChange(batteryData *chan Data) (ref *Battery) {
	var fn js.Func

	if e.fnChargingChange == nil {
		fn = js.FuncOf(func(this js.Value, args []js.Value) any {
			if len(args) == 0 {
				return nil
			}

			*batteryData <- EventManager(KEventChargingChange, this, args)

			return nil
		})
		e.fnChargingChange = &fn
	}

	objectBactery.Call(
		"addEventListener",
		"chargingchange",
		*e.fnChargingChange,
	)

	return e
}

// RemoveListenerChargingChange
//
// English:
//
// Remove a battery charging change event listener equivalent to the JavaScript command removeEventListener('chargingchange',fn).
//
// Fired when the battery charging state (the charging property) is updated.
//
// Português:
//
// Remove um ouvinte de evento de carga da bateria mudou, equivalente ao comando JavaScript removeEventListener('chargingchange',fn).
//
// Acionado quando o estado de carregamento da bateria (a propriedade de carregamento) é atualizado.
func (e *Battery) RemoveListenerChargingChange() (ref *Battery) {
	if e.fnChargingChange == nil {
		return e
	}

	objectBactery.Call(
		"removeEventListener",
		"chargingchange",
		*e.fnChargingChange,
	)

	e.fnChargingChange = nil
	return e
}

// AddListenerChargingTimeChange
//
// English:
//
// Adds a battery charging time change event listener equivalent to the JavaScript command addEventListener('chargingtimechange',fn).
//
//	Input:
//	  mouseEvet: pointer to channel battery.Data
//
// Fired when the battery charging time (the chargingTime property) is updated.
//
// Português:
//
// Adiciona um ouvinte de evento de tempo de carga da bateria mudou, equivalente ao comando JavaScript addEventListener('chargingtimechange',fn).
//
//	Entrada:
//	  mouseEvet: ponteiro para o channel battery.Data
//
// Acionado quando o tempo de carregamento da bateria (propriedade loadingTime) é atualizado.
func (e *Battery) AddListenerChargingTimeChange(batteryData *chan Data) (ref *Battery) {
	var fn js.Func

	if e.fnChargingTimeChange == nil {
		fn = js.FuncOf(func(this js.Value, args []js.Value) any {
			if len(args) == 0 {
				return nil
			}

			*batteryData <- EventManager(KEventChargingTimeChange, this, args)
			return nil
		})
		e.fnChargingTimeChange = &fn
	}

	objectBactery.Call(
		"addEventListener",
		"chargingtimechange",
		*e.fnChargingTimeChange,
	)
	return e
}

func (e *Battery) AddListenerDischargingTimeChange(batteryData *chan Data) (ref *Battery) {
	var fn js.Func

	if e.fnDischargingTimeChange == nil {
		fn = js.FuncOf(func(this js.Value, args []js.Value) any {
			if len(args) == 0 {
				return nil
			}

			*batteryData <- EventManager(KEventDischargingTimeChange, this, args)
			return nil
		})
		e.fnDischargingTimeChange = &fn
	}

	objectBactery.Call(
		"addEventListener",
		"dischargingtimechange",
		*e.fnDischargingTimeChange,
	)
	return e
}

func (e *Battery) AddListenerLevelChange(batteryData *chan Data) (ref *Battery) {
	var fn js.Func

	if e.fnLevelChange == nil {
		fn = js.FuncOf(func(this js.Value, args []js.Value) any {
			if len(args) == 0 {
				return nil
			}

			*batteryData <- EventManager(KEventLevelChange, this, args)
			return nil
		})
		e.fnLevelChange = &fn
	}

	objectBactery.Call(
		"addEventListener",
		"levelchange",
		*e.fnLevelChange,
	)
	return e
}

type EventName string

func (e EventName) String() string {
	return string(e)
}

const (
	// KEventChargingChange
	//
	// English:
	//
	// Fired when the battery charging state (the charging property) is updated.
	//
	// Português:
	//
	// Acionado quando o estado de carregamento da bateria (a propriedade de carregamento) é atualizado.
	KEventChargingChange EventName = "chargingchange"

	// KEventChargingTimeChange
	//
	// English:
	//
	// Fired when the battery charging time (the chargingTime property) is updated.
	//
	// Português:
	//
	// Acionado quando o tempo de carregamento da bateria (propriedade loadingTime) é atualizado.
	KEventChargingTimeChange EventName = "chargingtimechange"

	// KEventDischargingTimeChange
	//
	// English:
	//
	// Fired when the battery discharging time (the dischargingTime property) is updated.
	//
	// Português:
	//
	// Acionado quando o tempo de descarga da bateria (propriedade dischargingTime) é atualizado.
	KEventDischargingTimeChange EventName = "dischargingtimechange"

	// KEventLevelChange
	//
	// English:
	//
	// Fired when the battery level (the level property) is updated.
	//
	// Português:
	//
	// Acionado quando o nível da bateria (a propriedade de nível) é atualizado.
	KEventLevelChange EventName = "levelchange"
)

// Data
//
// English:
//
// The BatteryManager interface of the Battery Status API provides information about the system's battery charge level.
//
// Português:
//
// A interface BatteryManager da API de status da bateria fornece informações sobre o nível de carga da bateria do
// sistema.
type Data struct {

	// EventName
	//
	// English:
	//
	// Name of event
	//
	// Português:
	//
	// Nome do evento
	EventName EventName

	// This
	//
	// English:
	//
	// This is the equivalent property of JavaScript's 'this'.
	//
	// The way to use it is This.Get(property string name). E.g. chan.This.Get("id")
	//
	// Português:
	//
	// Esta é a propriedade equivalente ao 'this' do JavaScript.
	//
	// A forma de usar é This.Get(property string name). Ex. chan.This.Get("id")
	This js.Value

	// Charging
	//
	// English:
	//
	// Battery is currently being charged.
	//
	// Português:
	//
	// A bateria está sendo carregada no momento.
	Charging bool

	// ChargingTime
	//
	// English:
	//
	// A number representing the remaining time in seconds until the battery is fully charged, or 0 if the battery is
	// already fully charged.
	//
	//  Notes:
	//    * -1 is used to indicate an unknown value.
	//
	// Português:
	//
	// Um número que representa o tempo restante em segundos até que a bateria esteja totalmente carregada ou 0 se a
	// bateria já estiver totalmente carregada.
	//
	//  Notas:
	//    * -1 é usado para indicar um valor desconhecido.
	ChargingTime int

	// DischargingTime
	//
	// English:
	//
	// A number representing the remaining time in seconds until the battery is completely discharged and the system
	// suspends.
	//
	//  Notes:
	//    * -1 is used to indicate an unknown value.
	//
	// Português:
	//
	// Um número que representa o tempo restante em segundos até que a bateria seja completamente descarregada e o
	// sistema seja suspenso.
	//
	//  Notas:
	//    * -1 é usado para indicar um valor desconhecido.
	DischargingTime int

	// Level
	//
	// English:
	//
	// A number representing the system's battery charge level scaled to a value between 0.0 and 1.0.
	//
	// Português:
	//
	// Um número que representa o nível de carga da bateria do sistema dimensionado para um valor entre 0,0 e 1,0.
	Level float64
}

type Event struct {
	Object js.Value
}

// GetIsCharging
//
// English:
//
// Battery is currently being charged.
//
// Português:
//
// A bateria está sendo carregada no momento.
func (e Event) GetIsCharging() (isCharging bool) {
	return e.Object.Get("charging").Bool()
}

// GetChargingTime
//
// English:
//
// A number representing the remaining time in seconds until the battery is fully charged, or 0 if the battery is
// already fully charged.
//
//	Notes:
//	  * -1 is used to indicate an unknown value.
//
// Português:
//
// Um número que representa o tempo restante em segundos até que a bateria esteja totalmente carregada ou 0 se a
// bateria já estiver totalmente carregada.
//
//	Notas:
//	  * -1 é usado para indicar um valor desconhecido.
func (e Event) GetChargingTime() (chargingTime int) {
	chargingTime = e.Object.Get("chargingTime").Int()
	if chargingTime < 0 {
		return -1
	}

	return
}

// GetDischargingTime
//
// English:
//
// A number representing the remaining time in seconds until the battery is completely discharged and the system
// suspends.
//
//	Notes:
//	  * -1 is used to indicate an unknown value.
//
// Português:
//
// Um número que representa o tempo restante em segundos até que a bateria seja completamente descarregada e o
// sistema seja suspenso.
//
//	Notas:
//	  * -1 é usado para indicar um valor desconhecido.
func (e Event) GetDischargingTime() (dischargingTime int) {
	dischargingTime = e.Object.Get("dischargingTime").Int()
	if dischargingTime < 0 {
		return -1
	}

	return
}

// GetLevel
//
// English:
//
// A number representing the system's battery charge level scaled to a value between 0.0 and 1.0.
//
// Português:
//
// Um número que representa o nível de carga da bateria do sistema dimensionado para um valor entre 0,0 e 1,0.
func (e Event) GetLevel() (level float64) {
	return e.Object.Get("level").Float()
}

// EventManager
//
// English:
//
// Capture event information and format to Golang
//
//	Output:
//	  data: list with all the information provided by the browser.
//
// Português:
//
// Captura as informações do evento e formata para o Golang
//
//	Saída:
//	  data: lista com todas as informações fornecidas pelo navegador.
func EventManager(name EventName, this js.Value, _ []js.Value) (data Data) {
	var event = Event{}
	event.Object = objectBactery
	data.EventName = name
	data.This = this

	data.Charging = event.GetIsCharging()
	data.ChargingTime = event.GetChargingTime()
	data.DischargingTime = event.GetDischargingTime()
	data.Level = event.GetLevel()

	return
}
