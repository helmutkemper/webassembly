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
	// todo: fazer depois
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

// mountWhiteBalanceMode
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
func (e *Image) mountWhiteBalanceMode(image *map[string]interface{}) {
	if e.WhiteBalanceMode != "" {
		(*image)["whiteBalanceMode"] = e.WhiteBalanceMode.String()
	}
}

// mountExposureMode
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
func (e *Image) mountExposureMode(image *map[string]interface{}) {
	if e.ExposureMode != "" {
		(*image)["exposureMode"] = e.ExposureMode.String()
	}
}

// mountFocusMode
//
// English:
//
// A const specifying one of KFocusModeNone, KFocusModeManual, KFocusModeSingleShot, or KFocusModeContinuous.
//
// Português:
//
// Uma constante especificando um de KFocusModeNone, KFocusModeManual, KFocusModeSingleShot ou KFocusModeContinuous.
func (e *Image) mountFocusMode(image *map[string]interface{}) {
	if e.FocusMode != "" {
		(*image)["focusMode"] = e.FocusMode.String()
	}
}

// mountExposureCompensation
//
// English:
//
// A ConstrainDouble specifying f-stop adjustment by up to ±3.
//
// Português:
//
// A ConstrainDouble especificando o ajuste f-stop em até ±3.
func (e *Image) mountExposureCompensation(image *map[string]interface{}) {
	if e.ExposureCompensation.Exact != nil {
		(*image)["exposureCompensation"] = map[string]interface{}{"exact": e.ExposureCompensation.Exact}
		return
	}

	if e.ExposureCompensation.Max != nil || e.ExposureCompensation.Min != nil || e.ExposureCompensation.Ideal != nil {
		(*image)["exposureCompensation"] = make(map[string]interface{})

		if e.ExposureCompensation.Max != nil {
			(*image)["exposureCompensation"].(map[string]interface{})["max"] = e.ExposureCompensation.Max
		}

		if e.ExposureCompensation.Min != nil {
			(*image)["exposureCompensation"].(map[string]interface{})["min"] = e.ExposureCompensation.Min
		}

		if e.ExposureCompensation.Ideal != nil {
			(*image)["exposureCompensation"].(map[string]interface{})["ideal"] = e.ExposureCompensation.Ideal
		}

		return
	}

	if e.ExposureCompensation.Value != nil {
		(*image)["exposureCompensation"] = e.ExposureCompensation.Value
		return
	}
}

// mountColorTemperature
//
// English:
//
// A ConstrainDouble specifying a desired color temperature in degrees kelvin.
//
// Português:
//
// Um ConstrainDouble especificando uma temperatura de cor desejada em graus kelvin.
func (e *Image) mountColorTemperature(image *map[string]interface{}) {
	if e.ColorTemperature.Exact != nil {
		(*image)["colorTemperature"] = map[string]interface{}{"exact": e.ColorTemperature.Exact}
		return
	}

	if e.ColorTemperature.Max != nil || e.ColorTemperature.Min != nil || e.ColorTemperature.Ideal != nil {
		(*image)["colorTemperature"] = make(map[string]interface{})

		if e.ColorTemperature.Max != nil {
			(*image)["colorTemperature"].(map[string]interface{})["max"] = e.ColorTemperature.Max
		}

		if e.ColorTemperature.Min != nil {
			(*image)["colorTemperature"].(map[string]interface{})["min"] = e.ColorTemperature.Min
		}

		if e.ColorTemperature.Ideal != nil {
			(*image)["colorTemperature"].(map[string]interface{})["ideal"] = e.ColorTemperature.Ideal
		}

		return
	}

	if e.ColorTemperature.Value != nil {
		(*image)["colorTemperature"] = e.ColorTemperature.Value
		return
	}
}

// mountIso
//
// English:
//
// A ConstrainDouble specifying a desired iso setting.
//
// Português:
//
// Um ConstrainDouble especificando uma configuração iso desejada.
func (e *Image) mountIso(image *map[string]interface{}) {
	if e.Iso.Exact != nil {
		(*image)["iso"] = map[string]interface{}{"exact": e.Iso.Exact}
		return
	}

	if e.Iso.Max != nil || e.Iso.Min != nil || e.Iso.Ideal != nil {
		(*image)["iso"] = make(map[string]interface{})

		if e.Iso.Max != nil {
			(*image)["iso"].(map[string]interface{})["max"] = e.Iso.Max
		}

		if e.Iso.Min != nil {
			(*image)["iso"].(map[string]interface{})["min"] = e.Iso.Min
		}

		if e.Iso.Ideal != nil {
			(*image)["iso"].(map[string]interface{})["ideal"] = e.Iso.Ideal
		}

		return
	}

	if e.Iso.Value != nil {
		(*image)["iso"] = e.Iso.Value
		return
	}
}

// mountBrightness
//
// English:
//
// A ConstrainDouble specifying a desired brightness setting.
//
// Português:
//
// Um ConstrainDouble especificando uma configuração de brilho desejada.
func (e *Image) mountBrightness(image *map[string]interface{}) {
	if e.Brightness.Exact != nil {
		(*image)["brightness"] = map[string]interface{}{"exact": e.Brightness.Exact}
		return
	}

	if e.Brightness.Max != nil || e.Brightness.Min != nil || e.Brightness.Ideal != nil {
		(*image)["brightness"] = make(map[string]interface{})

		if e.Brightness.Max != nil {
			(*image)["brightness"].(map[string]interface{})["max"] = e.Brightness.Max
		}

		if e.Brightness.Min != nil {
			(*image)["brightness"].(map[string]interface{})["min"] = e.Brightness.Min
		}

		if e.Brightness.Ideal != nil {
			(*image)["brightness"].(map[string]interface{})["ideal"] = e.Brightness.Ideal
		}

		return
	}

	if e.Brightness.Value != nil {
		(*image)["brightness"] = e.Brightness.Value
		return
	}
}

// mountContrast
//
// English:
//
// A ConstrainDouble specifying the degree of difference between light and dark.
//
// Português:
//
// Um ConstrainDouble especificando o grau de diferença entre claro e escuro.
func (e *Image) mountContrast(image *map[string]interface{}) {
	if e.Contrast.Exact != nil {
		(*image)["contrast"] = map[string]interface{}{"exact": e.Contrast.Exact}
		return
	}

	if e.Contrast.Max != nil || e.Contrast.Min != nil || e.Contrast.Ideal != nil {
		(*image)["contrast"] = make(map[string]interface{})

		if e.Contrast.Max != nil {
			(*image)["contrast"].(map[string]interface{})["max"] = e.Contrast.Max
		}

		if e.Contrast.Min != nil {
			(*image)["contrast"].(map[string]interface{})["min"] = e.Contrast.Min
		}

		if e.Contrast.Ideal != nil {
			(*image)["contrast"].(map[string]interface{})["ideal"] = e.Contrast.Ideal
		}

		return
	}

	if e.Contrast.Value != nil {
		(*image)["contrast"] = e.Contrast.Value
		return
	}
}

// mountSaturation
//
// English:
//
// A ConstrainDouble specifying the degree of color intensity.
//
// Português:
//
// Um ConstrainDouble especificando o grau de intensidade da cor.
func (e *Image) mountSaturation(image *map[string]interface{}) {
	if e.Saturation.Exact != nil {
		(*image)["saturation"] = map[string]interface{}{"exact": e.Saturation.Exact}
		return
	}

	if e.Saturation.Max != nil || e.Saturation.Min != nil || e.Saturation.Ideal != nil {
		(*image)["saturation"] = make(map[string]interface{})

		if e.Saturation.Max != nil {
			(*image)["saturation"].(map[string]interface{})["max"] = e.Saturation.Max
		}

		if e.Saturation.Min != nil {
			(*image)["saturation"].(map[string]interface{})["min"] = e.Saturation.Min
		}

		if e.Saturation.Ideal != nil {
			(*image)["saturation"].(map[string]interface{})["ideal"] = e.Saturation.Ideal
		}

		return
	}

	if e.Saturation.Value != nil {
		(*image)["saturation"] = e.Saturation.Value
		return
	}
}

// mountSharpness
//
// English:
//
// A ConstrainDouble specifying the intensity of edges.
//
// Português:
//
// Um ConstrainDouble especificando à intensidade das arestas.
func (e *Image) mountSharpness(image *map[string]interface{}) {
	if e.Sharpness.Exact != nil {
		(*image)["sharpness"] = map[string]interface{}{"exact": e.Sharpness.Exact}
		return
	}

	if e.Sharpness.Max != nil || e.Sharpness.Min != nil || e.Sharpness.Ideal != nil {
		(*image)["sharpness"] = make(map[string]interface{})

		if e.Sharpness.Max != nil {
			(*image)["sharpness"].(map[string]interface{})["max"] = e.Sharpness.Max
		}

		if e.Sharpness.Min != nil {
			(*image)["sharpness"].(map[string]interface{})["min"] = e.Sharpness.Min
		}

		if e.Sharpness.Ideal != nil {
			(*image)["sharpness"].(map[string]interface{})["ideal"] = e.Sharpness.Ideal
		}

		return
	}

	if e.Sharpness.Value != nil {
		(*image)["sharpness"] = e.Sharpness.Value
		return
	}
}

// mountFocusDistance
//
// English:
//
// A ConstrainDouble specifying distance to a focused object.
//
// Português:
//
// Um ConstrainDouble especificando a distância para um objeto focado.
func (e *Image) mountFocusDistance(image *map[string]interface{}) {
	if e.FocusDistance.Exact != nil {
		(*image)["focusDistance"] = map[string]interface{}{"exact": e.FocusDistance.Exact}
		return
	}

	if e.FocusDistance.Max != nil || e.FocusDistance.Min != nil || e.FocusDistance.Ideal != nil {
		(*image)["focusDistance"] = make(map[string]interface{})

		if e.FocusDistance.Max != nil {
			(*image)["focusDistance"].(map[string]interface{})["max"] = e.FocusDistance.Max
		}

		if e.FocusDistance.Min != nil {
			(*image)["focusDistance"].(map[string]interface{})["min"] = e.FocusDistance.Min
		}

		if e.FocusDistance.Ideal != nil {
			(*image)["focusDistance"].(map[string]interface{})["ideal"] = e.FocusDistance.Ideal
		}

		return
	}

	if e.FocusDistance.Value != nil {
		(*image)["focusDistance"] = e.FocusDistance.Value
		return
	}
}

// mountZoom
//
// English:
//
// A ConstrainDouble specifying the desired focal length.
//
// Português:
//
// A ConstrainDouble especificando a distância focal desejada.
func (e *Image) mountZoom(image *map[string]interface{}) {
	if e.Zoom.Exact != nil {
		(*image)["zoom"] = map[string]interface{}{"exact": e.Zoom.Exact}
		return
	}

	if e.Zoom.Max != nil || e.Zoom.Min != nil || e.Zoom.Ideal != nil {
		(*image)["zoom"] = make(map[string]interface{})

		if e.Zoom.Max != nil {
			(*image)["zoom"].(map[string]interface{})["max"] = e.Zoom.Max
		}

		if e.Zoom.Min != nil {
			(*image)["zoom"].(map[string]interface{})["min"] = e.Zoom.Min
		}

		if e.Zoom.Ideal != nil {
			(*image)["zoom"].(map[string]interface{})["ideal"] = e.Zoom.Ideal
		}

		return
	}

	if e.Zoom.Value != nil {
		(*image)["zoom"] = e.Zoom.Value
		return
	}
}

// mountTorch
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
func (e *Image) mountTorch(image *map[string]interface{}) {
	if e.Torch.IsSet() {
		(*image)["torch"] = map[string]interface{}{"exact": e.Torch.Bool()}
		return
	}
}

func (e *Image) mount(image *map[string]interface{}) {
	if *image == nil {
		*image = make(map[string]interface{})
	}

	e.mountWhiteBalanceMode(image)
	e.mountExposureMode(image)
	e.mountFocusMode(image)
	e.mountExposureCompensation(image)
	e.mountColorTemperature(image)
	e.mountIso(image)
	e.mountBrightness(image)
	e.mountContrast(image)
	e.mountSaturation(image)
	e.mountSharpness(image)
	e.mountFocusDistance(image)
	e.mountZoom(image)
	e.mountTorch(image)
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
