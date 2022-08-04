package media

//todo: fazer

type Media struct {

	// DeviceId
	//
	// English:
	//
	// A ConstrainDOMString object specifying a device ID or an array of device IDs which are acceptable and/or required.
	//
	// Português:
	//
	// Um objeto ConstrainDOMString que especifica um ID de dispositivo ou uma matriz de IDs de dispositivo que são
	// aceitáveis e/ou obrigatórios.
	// https://developer.mozilla.org/en-US/docs/Web/API/MediaTrackConstraints/deviceId
	DeviceId interface{} `js:"deviceId"`

	// GroupId
	//
	// English:
	//
	// A ConstrainDOMString object specifying a group ID or an array of group IDs which are acceptable and/or required.
	//
	// Português:
	//
	// Um objeto ConstrainDOMString especificando um ID de grupo ou uma matriz de IDs de grupo que são aceitáveis e/ou
	// obrigatórios.
	// https://developer.mozilla.org/en-US/docs/Web/API/MediaTrackConstraints/groupId
	GroupId interface{} `js:"groupId"`
}
