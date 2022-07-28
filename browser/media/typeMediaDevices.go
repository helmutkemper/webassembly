package media

import (
	"errors"
	"syscall/js"
)

type Devices struct{}

// DevicesData
//
// English:
//
// The MediaDeviceInfo interface contains information that describes a single media input or output device.
//
// Português:
//
// A interface MediaDeviceInfo contém informações que descrevem um único dispositivo de entrada ou saída de mídia.
type DevicesData struct {

	// DeviceId
	//
	// English:
	//
	// Identifier for the represented device that is persisted across sessions.
	//
	// It is un-guessable by other applications and unique to the origin of the calling application.
	// It is reset when the user clears cookies (for Private Browsing, a different identifier is used that is not
	// persisted across sessions).
	//
	// Português:
	//
	// Identificador para o dispositivo representado que persiste nas sessões.
	//
	// Ele não pode ser adivinhado por outros aplicativos e exclusivo para a origem do aplicativo de chamada.
	// Ele é redefinido quando o usuário limpa os cookies (para Navegação Privada, é usado um identificador diferente que
	// não persiste nas sessões).
	DeviceId string

	// GroupId
	//
	// English:
	//
	// Group identifier. Two devices have the same group identifier if they belong to the same physical device — for
	// example a monitor with both a built-in camera and a microphone.
	//
	// Português:
	//
	// Retorna uma string que é um identificador de grupo. Dois dispositivos têm o mesmo identificador de grupo se
	// pertencerem ao mesmo dispositivo físico — por exemplo, um monitor com uma câmera integrada e um microfone.
	GroupId string

	// Kind
	//
	// English:
	//
	// Returns an enumerated value that is either "videoinput", "audioinput" or "audiooutput".
	//
	// Português:
	//
	// Retorna um valor enumerado que é "videoinput", "audioinput" ou "audiooutput".
	Kind Kind

	// Label
	//
	// English:
	//
	// Returns a string describing this device (for example "External USB Webcam").
	//
	//   Notes:
	//     * For security reasons, the label field is always blank unless an active media stream exists or the user has
	//       granted persistent permission for media device access. The set of device labels could otherwise be used as
	//       part of a fingerprinting mechanism to identify a user.
	//
	// Português:
	//
	// Retorna uma string descrevendo este dispositivo (por exemplo, "External USB Webcam").
	//
	//   Notas:
	//     * Por motivos de segurança, o campo de rótulo está sempre em branco, a menos que exista um fluxo de mídia ativo
	//       ou que o usuário tenha concedido permissão persistente para acesso ao dispositivo de mídia. O conjunto de
	//       etiquetas do dispositivo poderia ser usado como parte de um mecanismo de impressão digital para identificar
	//       um usuário.
	Label string
}

// EnumerateDevices
//
// English:
//
// Return a list of the available media input and output devices, such as microphones, cameras, headsets, and so forth.
//
// Access to particular devices is gated by the Permissions API. The list of returned devices will omit any devices for
// which the corresponding permission has not been granted, including: microphone, camera, speaker-selection (for output
// devices), and so on.
//
//	Notes:
//	  * For security reasons, the label field is always blank unless an active media stream exists or the user has
//	    granted persistent permission for media device access. The set of device labels could otherwise be used as
//	    part of a fingerprinting mechanism to identify a user.
//
// Português:
//
// Retorne uma lista dos dispositivos de entrada e saída de mídia disponíveis, como microfones, câmeras, fones de ouvido
// e assim por diante.
//
// O acesso a dispositivos específicos é controlado pela API de permissões. A lista de dispositivos devolvidos omitirá
// quaisquer dispositivos para os quais a permissão correspondente não tenha sido concedida, incluindo: microfone,
// câmera, seleção de alto-falante (para dispositivos de saída) e assim por diante.
//
//	Notas:
//	  * Por motivos de segurança, o campo de rótulo está sempre em branco, a menos que exista um fluxo de mídia ativo
//	    ou que o usuário tenha concedido permissão persistente para acesso ao dispositivo de mídia. O conjunto de
//	    etiquetas do dispositivo poderia ser usado como parte de um mecanismo de impressão digital para identificar
//	    um usuário.
func (e *Devices) EnumerateDevices() (list []DevicesData, err error) {
	list = make([]DevicesData, 0)
	end := make(chan struct{})

	// golang has a bug:
	// enumerateDevices() returns an array, but, go returns an object.
	forEach := js.FuncOf(func(_ js.Value, args []js.Value) any {
		data := DevicesData{
			DeviceId: args[0].Get("deviceId").String(),
			GroupId:  args[0].Get("groupId").String(),
			Kind:     Kind(args[0].Get("kind").String()),
			Label:    args[0].Get("label").String(),
		}
		list = append(list, data)

		// aways return nil
		return nil
	})

	// promise success function
	var success = js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		args[0].Call("forEach", forEach)
		end <- struct{}{}

		// aways return nil
		return nil
	})

	var failure = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		err = errors.New(args[0].Get("message").String())

		// aways return nil
		return nil
	})

	js.Global().Get("navigator").Get("mediaDevices").Call("enumerateDevices").Call("then", success, failure)

	// wait async call
	<-end

	return
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
