package media

type Image struct {

	// WhiteBalanceMode
	//
	// English:
	//
	// A const specifying one of KWhiteBalanceModeNone, KWhiteBalanceModeManual, KWhiteBalanceModeSingleShot, or
	// KWhiteBalanceModeContinuous.
	//
	// Português:
	//
	// Uma constante especificando KWhiteBalanceModeNone, KWhiteBalanceModeManual, KWhiteBalanceModeSingleShot ou
	// KWhiteBalanceModeContinuous.
	WhiteBalanceMode WhiteBalanceMode `js:"whiteBalanceMode"`

	// ExposureMode
	//
	// English:
	//
	// A const specifying one of KExposureModeNone, KExposureModeManual, KExposureModeSingleShot, or
	// KExposureModeContinuous.
	//
	// Português:
	//
	// Uma constante especificando um de KExposureModeNone, KExposureModeManual, KExposureModeSingleShot ou
	// KExposureModeContinuous.
	ExposureMode ExposureMode `js:"exposureMode"`

	// FocusMode
	//
	// English:
	//
	// A const specifying one of KFocusModeNone, KFocusModeManual, KFocusModeSingleShot, or KFocusModeContinuous.
	//
	// Português:
	//
	// Uma constante especificando um de KFocusModeNone, KFocusModeManual, KFocusModeSingleShot ou KFocusModeContinuous.
	FocusMode FocusMode `js:"focusMode"`

	// PointsOfInterest
	//
	// English:
	//
	// The pixel coordinates on the sensor of one or more points of interest.
	// This is either an object in the form { x:value, y:value } or an array of such objects, where value is a
	// double-precision integer.
	//
	// Português:
	//
	// O pixel coordena no sensor de um ou mais pontos de interesse.
	// Este é um objeto no formato { x:value, y:value } ou uma matriz de tais objetos, onde value é um inteiro de precisão
	// dupla.
	PointsOfInterest interface{} `js:"pointsOfInterest"`

	// ExposureCompensation
	//
	// English:
	//
	// A ConstrainDouble specifying f-stop adjustment by up to ±3.
	//
	// Português:
	//
	// A ConstrainDouble especificando o ajuste f-stop em até ±3.
	ExposureCompensation ConstrainDouble `js:"exposureCompensation"`

	// ColorTemperature
	//
	// English:
	//
	// A ConstrainDouble specifying a desired color temperature in degrees kelvin.
	//
	// Português:
	//
	// Um ConstrainDouble especificando uma temperatura de cor desejada em graus kelvin.
	ColorTemperature ConstrainDouble `js:"colorTemperature"`

	// Iso
	//
	// English:
	//
	// A ConstrainDouble specifying a desired iso setting.
	//
	// Português:
	//
	// Um ConstrainDouble especificando uma configuração iso desejada.
	Iso ConstrainDouble `js:"iso"`

	// Brightness
	//
	// English:
	//
	// A ConstrainDouble specifying a desired brightness setting.
	//
	// Português:
	//
	// Um ConstrainDouble especificando uma configuração de brilho desejada.
	Brightness ConstrainDouble `js:"brightness"`

	// Contrast
	//
	// English:
	//
	// A ConstrainDouble specifying the degree of difference between light and dark.
	//
	// Português:
	//
	// Um ConstrainDouble especificando o grau de diferença entre claro e escuro.
	Contrast ConstrainDouble `js:"contrast"`

	// Saturation
	//
	// English:
	//
	// A ConstrainDouble specifying the degree of color intensity.
	//
	// Português:
	//
	// Um ConstrainDouble especificando o grau de intensidade da cor.
	Saturation ConstrainDouble `js:"saturation"`

	// Sharpness
	//
	// English:
	//
	// A ConstrainDouble specifying the intensity of edges.
	//
	// Português:
	//
	// Um ConstrainDouble especificando à intensidade das arestas.
	Sharpness ConstrainDouble `js:"sharpness"`

	// FocusDistance
	//
	// English:
	//
	// A ConstrainDouble specifying distance to a focused object.
	//
	// Português:
	//
	// Um ConstrainDouble especificando a distância para um objeto focado.
	FocusDistance ConstrainDouble `js:"focusDistance"`

	// Zoom
	//
	// English:
	//
	// A ConstrainDouble specifying the desired focal length.
	//
	// Português:
	//
	// A ConstrainDouble especificando a distância focal desejada.
	Zoom ConstrainDouble `js:"zoom"`

	// Torch
	//
	// English:
	//
	// A boolean value defining whether the fill light is continuously connected, meaning it stays on as long as the
	// track is active.
	//
	// Português:
	//
	// Um valor booleano que define se a luz de preenchimento está continuamente conectada, o que significa que permanece
	// acesa enquanto a trilha estiver ativa.
	Torch BOOLEAN `js:"torch"`
}
