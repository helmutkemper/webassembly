package geolocation

import "syscall/js"

// Geolocation
//
// English:
//
// The Geolocation interface represents an object able to obtain the position of the device programmatically. It gives
// Web content access to the location of the device. This allows a website or app to offer customized results based on
// the user's location.
//
// An object with this interface is obtained using the navigator.geolocation property implemented by the Navigator
// object.
//
//   Notes:
//    * For security reasons, when a web page tries to access location information, the user is notified and asked to
//      grant permission. Be aware that each browser has its own policies and methods for requesting this permission.
//    * This feature is available only in secure contexts (HTTPS), in some or all supporting browsers.
//
// Português:
//
// A interface Geolocation representa um objeto capaz de obter a posição do dispositivo programaticamente.
// Dá acesso ao conteúdo da Web à localização do dispositivo. Isso permite que um site ou aplicativo ofereça resultados
// personalizados com base na localização do usuário.
//
// An object with this interface is obtained using the navigator.geolocation property implemented by the Navigator
// object.
//
//   Notes:
//    * Por motivos de segurança, quando uma página da Web tenta acessar informações de localização, o usuário é
//      notificado e solicitado a conceder permissão. Esteja ciente de que cada navegador tem suas próprias políticas e
//      métodos para solicitar essa permissão.
//    * Este recurso está disponível apenas em contextos seguros (HTTPS), em alguns ou em todos os navegadores compatíveis.
type Geolocation struct {

	// enableHighAccuracy
	//
	// English:
	//
	// A boolean value that indicates the application would like to receive the best possible results.
	// If true and if the device is able to provide a more accurate position, it will do so. Note that this can result in
	// slower response times or increased power consumption (with a GPS chip on a mobile device for example). On the other
	// hand, if false, the device can take the liberty to save resources by responding more quickly and/or using less
	// power.
	//
	// Default: false.
	//
	// Português:
	//
	// Um valor booleano que indica que o aplicativo gostaria de receber os melhores resultados possíveis. Se for
	// verdadeiro e se o dispositivo for capaz de fornecer uma posição mais precisa, ele o fará. Observe que isso pode
	// resultar em tempos de resposta mais lentos ou maior consumo de energia (com um chip GPS em um dispositivo móvel,
	// por exemplo). Por outro lado, se falso, o dispositivo pode ter a liberdade de economizar recursos respondendo mais
	// rapidamente e ou usando menos energia.
	//
	// Padrão: falso.
	enableHighAccuracy bool

	// timeout
	//
	// English:
	//
	// A positive long value representing the maximum length of time (in milliseconds) the device is allowed to take in
	// order to return a position.
	// The default value is Infinity, meaning that getCurrentPosition() won't return until the position is available.
	//
	// Português:
	//
	// Um valor longo positivo que representa o tempo máximo (em milissegundos) que o dispositivo pode levar para retornar
	// uma posição.
	// O valor padrão é Infinity, o que significa que getCurrentPosition() não retornará até que a posição esteja
	// disponível.
	timeout int

	// maximumAge
	//
	// English:
	//
	// A positive long value indicating the maximum age in milliseconds of a possible cached position that is acceptable
	// to return.
	// If set to 0, it means that the device cannot use a cached position and must attempt to retrieve the real current
	// position.
	// If set to -1 the device must return a cached position regardless of its age.
	// Default: 0.
	//
	// Português:
	//
	// Um valor longo positivo que indica a idade máxima em milissegundos de uma possível posição em cache que é aceitável
	// retornar.
	// Se definido como 0, significa que o dispositivo não pode usar uma posição em cache e deve tentar recuperar a
	// posição atual real.
	// Se definido como -1, o dispositivo deve retornar uma posição em cache, independentemente de sua idade.
	// Padrão: 0.
	maximumAge int
}

// MaximumAge
//
// English:
//
// A positive long value indicating the maximum age in milliseconds of a possible cached position that is acceptable
// to return.
//
//   Input:
//     maximumAge: value indicating the maximum age in milliseconds.
//
// If set to 0, it means that the device cannot use a cached position and must attempt to retrieve the real current
// position.
// If set to -1 the device must return a cached position regardless of its age.
// Default: 0.
//
// Português:
//
// Um valor longo positivo que indica a idade máxima em milissegundos de uma possível posição em cache que é aceitável
// retornar.
//
//   Entrada:
//     maximumAge: valor que indica a idade máxima em milissegundos.
//
// Se definido como 0, significa que o dispositivo não pode usar uma posição em cache e deve tentar recuperar a
// posição atual real.
// Se definido como -1, o dispositivo deve retornar uma posição em cache, independentemente de sua idade.
// Padrão: 0.
func (e *Geolocation) MaximumAge(maximumAge int) {
	e.maximumAge = maximumAge
}

// Timeout
//
// English:
//
// A positive long value representing the maximum length of time (in milliseconds) the device is allowed to take in
// order to return a position.
//
//   Input:
//     timeout: value representing the maximum length of time (in milliseconds)
//
// The default value is Infinity, meaning that getCurrentPosition() won't return until the position is available.
//
// Português:
//
// Um valor longo positivo que representa o tempo máximo (em milissegundos) que o dispositivo pode levar para retornar
// uma posição.
//
//   Entrada:
//     timeout: valor que representa a duração máxima de tempo (em milissegundos)
//
// O valor padrão é Infinity, o que significa que getCurrentPosition() não retornará até que a posição esteja
// disponível.
func (e *Geolocation) Timeout(timeout int) {
	e.timeout = timeout
}

// EnableHighAccuracy
//
// English:
//
// A boolean value that indicates the application would like to receive the best possible results.
//
//   Input:
//     enableHighAccuracy: indicates the application would like to receive the best possible results.
//
// If true and if the device is able to provide a more accurate position, it will do so. Note that this can result in
// slower response times or increased power consumption (with a GPS chip on a mobile device for example). On the other
// hand, if false, the device can take the liberty to save resources by responding more quickly and/or using less
// power.
//
// Default: false.
//
// Português:
//
// Um valor booleano que indica que o aplicativo gostaria de receber os melhores resultados possíveis.
//
//   Entrada:
//     enableHighAccuracy: indica que o aplicativo gostaria de receber os melhores resultados possíveis.
//
// Se for verdadeiro e se o dispositivo for capaz de fornecer uma posição mais precisa, ele o fará. Observe que isso
// pode resultar em tempos de resposta mais lentos ou maior consumo de energia (com um chip GPS em um dispositivo móvel,
// por exemplo). Por outro lado, se falso, o dispositivo pode ter a liberdade de economizar recursos respondendo mais
// rapidamente e ou usando menos energia.
//
// Padrão: falso.
func (e *Geolocation) EnableHighAccuracy(enableHighAccuracy bool) {
	e.enableHighAccuracy = enableHighAccuracy
}

// GetPosition
//
// English:
//
// Get the current position of the device.
//
//   Output:
//     coordinate: coordinate object
//
//   Notes:
//     * Accuracy is in meters
//     * This feature is available only in secure contexts (HTTPS).
//
// Português:
//
// Retorna a posição atual do dispositivo.
//
//   Saída:
//     coordinate: objeto de coordenadas
//
//   Notas:
//     * Accuracy é em metros.
//     * Esse recurso está disponível apenas em contextos seguros (HTTPS).
func (e *Geolocation) GetPosition(chCoordinate *chan Coordinate) {
	var coordinate Coordinate
	var options = e.prepareOptions()

	onError := js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		coordinate.ErrorCode = args[0].Get("code").Int()
		coordinate.ErrorMessage = args[0].Get("message").String()
		*chCoordinate <- coordinate
		return nil
	})

	onSuccess := js.FuncOf(func(pos js.Value, args []js.Value) interface{} {
		if len(args) == 0 || args[0].IsUndefined() || args[0].IsNull() {
			return nil
		}

		coordinates := args[0].Get("coords")
		coordinate.Latitude = coordinates.Get("latitude").Float()
		coordinate.Longitude = coordinates.Get("longitude").Float()
		coordinate.Accuracy = coordinates.Get("accuracy").Float()
		*chCoordinate <- coordinate
		return nil
	})

	js.Global().Get("navigator").Get("geolocation").Call("getCurrentPosition", onSuccess, onError, options)
	return
}

// WatchPosition
//
// English:
//
// Register a handler function that will be called automatically each time the position of the device changes.
//
//   Notes:
//     * This feature is available only in secure contexts (HTTPS).
//
// Português:
//
// Registre uma função de manipulador que será chamada automaticamente toda vez que a posição do dispositivo for
// alterada.
//
//   Notas:
//     * Este recurso está disponível apenas em contextos seguros (HTTPS).
func (e *Geolocation) WatchPosition(chCoordinate *chan Coordinate) {
	var coordinate Coordinate
	var options = e.prepareOptions()

	onError := js.FuncOf(func(errJs js.Value, _ []js.Value) interface{} {
		coordinate.ErrorCode = errJs.Get("code").Int()
		coordinate.ErrorMessage = errJs.Get("message").String()

		*chCoordinate <- coordinate
		return nil
	})

	onSuccess := js.FuncOf(func(pos js.Value, _ []js.Value) interface{} {
		if pos.IsUndefined() || pos.IsNull() {
			return nil
		}

		coordinates := pos.Get("coords")
		coordinate.Latitude = coordinates.Get("latitude").Float()
		coordinate.Longitude = coordinates.Get("longitude").Float()
		coordinate.Accuracy = coordinates.Get("accuracy").Float()

		*chCoordinate <- coordinate
		return nil
	})

	js.Global().Get("navigator").Get("geolocation").Call("watchPosition", onSuccess, onError, options)
	return
}

// prepareOptions
//
// English:
//
// An optional object from Geolocation.getCurrentPosition()
//
// See https://developer.mozilla.org/en-US/docs/Web/API/Geolocation/getCurrentPosition
//
// Português:
//
// Um objeto opcional de Geolocation.getCurrentPosition()
//
// Veja https://developer.mozilla.org/en-US/docs/Web/API/Geolocation/getCurrentPosition
func (e *Geolocation) prepareOptions() (options js.Value) {
	options = js.Global().Get("Object")

	if e.maximumAge < 0 {
		options.Set("maximumAge", "Infinity")
	} else {
		options.Set("maximumAge", e.maximumAge)
	}

	if e.timeout > 0 {
		options.Set("timeout", e.timeout)
	} else {
		options.Set("timeout", "Infinity")
	}

	options.Set("enableHighAccuracy", e.enableHighAccuracy)
	return
}
