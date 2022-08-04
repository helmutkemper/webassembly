package media

type Video struct {

	// AspectRatio
	//
	// English:
	//
	// A ConstrainDouble specifying the video aspect ratio or range of aspect ratios which are acceptable and/or required.
	//
	// Português:
	//
	// Um ConstrainDouble especificando a proporção do vídeo ou o intervalo de proporções do vídeo que são aceitáveis
	// e/ou obrigatórios.
	AspectRatio ConstrainDouble

	// FacingMode
	//
	// English:
	//
	// A ConstrainDOMString object specifying a facing or an array of facings which are acceptable and/or required.
	//
	// Português:
	//
	// Um objeto ConstrainDOMString especificando um revestimento ou uma matriz de revestimentos que são aceitáveis
	// e/ou obrigatórios.
	FacingMode FacingMode //`js:"facingMode"`

	// FrameRate
	//
	// English:
	//
	// A ConstrainDouble specifying the frame rate or range of frame rates which are acceptable and/or required.
	//
	// Português:
	//
	// Um ConstrainDouble especificando a taxa de quadros ou intervalo de taxas de quadros que são aceitáveis
	// e/ou necessários.
	FrameRate ConstrainDouble //`js:"frameRate"`

	// Height
	//
	// English:
	//
	// A ConstrainULong specifying the video height or range of heights which are acceptable and/or required.
	//
	// Português:
	//
	// Um ConstrainULong especificando a altura do vídeo ou o intervalo de alturas que são aceitáveis e/ou obrigatórios.
	Height ConstrainULong //`js:"height"`

	// Width
	//
	// English:
	//
	// A ConstrainULong specifying the video width or range of widths which are acceptable and/or required.
	//
	// Português:
	//
	// Um ConstrainULong especificando a largura do vídeo ou o intervalo de larguras que são aceitáveis e/ou obrigatórios.
	Width ConstrainULong //`js:"width"`

	// ResizeMode
	//
	// English:
	//
	// A ConstrainDOMString object specifying a mode or an array of modes the UA can use to derive the resolution of a
	// video track. Allowed values are none and crop-and-scale. none means that the user agent uses the resolution
	// provided by the camera, its driver or the OS. crop-and-scale means that the user agent can use cropping and
	// downscaling on the camera output in order to satisfy other constraints that affect the resolution.
	//
	// Português:
	//
	// Um objeto ConstrainDOMString especificando um modo ou uma matriz de modos que o UA pode usar para derivar a
	// resolução de uma trilha de vídeo. Os valores permitidos são none e crop-and-scale. none significa que o agente do
	// usuário usa a resolução fornecida pela câmera, seu driver ou o sistema operacional. recortar e dimensionar
	// significa que o agente do usuário pode usar recorte e redução de escala na saída da câmera para atender a outras
	// restrições que afetam a resolução.
	ResizeMode interface{} //`js:"resizeMode"`
}

// js.Global().Call("Object", map[string]interface{}{"a": 1, "b": 2})

// mountAspectRatio
//
// English:
//
// A ConstrainDouble specifying the video aspect ratio or range of aspect ratios which are acceptable and/or required.
//
// Português:
//
// Um ConstrainDouble especificando a proporção do vídeo ou o intervalo de proporções do vídeo que são aceitáveis
// e/ou obrigatórios.
func (e *Video) mountAspectRatio(aspectRatio *map[string]interface{}) {
	if e.AspectRatio.Exact != nil {
		(*aspectRatio)["aspectRatio"] = map[string]interface{}{"exact": e.AspectRatio.Exact}
		return
	}

	if e.AspectRatio.Max != nil || e.AspectRatio.Min != nil || e.AspectRatio.Ideal != nil {
		(*aspectRatio)["aspectRatio"] = make(map[string]interface{})

		if e.AspectRatio.Max != nil {
			(*aspectRatio)["aspectRatio"].(map[string]interface{})["max"] = e.AspectRatio.Max
		}

		if e.AspectRatio.Min != nil {
			(*aspectRatio)["aspectRatio"].(map[string]interface{})["min"] = e.AspectRatio.Min
		}

		if e.AspectRatio.Ideal != nil {
			(*aspectRatio)["aspectRatio"].(map[string]interface{})["ideal"] = e.AspectRatio.Ideal
		}

		return
	}

	if e.AspectRatio.Value != nil {
		(*aspectRatio)["aspectRatio"] = e.AspectRatio.Value
		return
	}
}

// mountFacingMode
//
// English:
//
// A ConstrainDOMString object specifying a facing or an array of facings which are acceptable and/or required.
//
// Português:
//
// Um objeto ConstrainDOMString especificando um revestimento ou uma matriz de revestimentos que são aceitáveis
// e/ou obrigatórios.
func (e *Video) mountFacingMode(video *map[string]interface{}) {
	e.FacingMode.mount(video)
}

// mountFrameRate
//
// English:
//
// A ConstrainDouble specifying the frame rate or range of frame rates which are acceptable and/or required.
//
// Português:
//
// Um ConstrainDouble especificando a taxa de quadros ou intervalo de taxas de quadros que são aceitáveis
// e/ou necessários.
func (e *Video) mountFrameRate(video *map[string]interface{}) {
	if e.FrameRate.Exact != nil {
		(*video)["frameRate"] = map[string]interface{}{"exact": e.FrameRate.Exact}
		return
	}

	if e.FrameRate.Max != nil || e.FrameRate.Min != nil || e.FrameRate.Ideal != nil {
		(*video)["frameRate"] = make(map[string]interface{})

		if e.FrameRate.Max != nil {
			(*video)["frameRate"].(map[string]interface{})["max"] = e.FrameRate.Max
		}

		if e.FrameRate.Min != nil {
			(*video)["frameRate"].(map[string]interface{})["min"] = e.FrameRate.Min
		}

		if e.FrameRate.Ideal != nil {
			(*video)["frameRate"].(map[string]interface{})["ideal"] = e.FrameRate.Ideal
		}

		return
	}

	if e.FrameRate.Value != nil {
		(*video)["frameRate"] = e.FrameRate.Value
		return
	}
}

// mountHeight
//
// English:
//
// A ConstrainULong specifying the video height or range of heights which are acceptable and/or required.
//
// Português:
//
// Um ConstrainULong especificando a altura do vídeo ou o intervalo de alturas que são aceitáveis e/ou obrigatórios.
func (e *Video) mountHeight(video *map[string]interface{}) {
	if e.Height.Exact != nil {
		(*video)["height"] = map[string]interface{}{"exact": e.Height.Exact}
		return
	}

	if e.Height.Max != nil || e.Height.Min != nil || e.Height.Ideal != nil {
		(*video)["height"] = make(map[string]interface{})

		if e.Height.Max != nil {
			(*video)["height"].(map[string]interface{})["max"] = e.Height.Max
		}

		if e.Height.Min != nil {
			(*video)["height"].(map[string]interface{})["min"] = e.Height.Min
		}

		if e.Height.Ideal != nil {
			(*video)["height"].(map[string]interface{})["ideal"] = e.Height.Ideal
		}

		return
	}

	if e.Height.Value != nil {
		(*video)["height"] = e.Height.Value
		return
	}
}

// Width
//
// English:
//
// A ConstrainULong specifying the video width or range of widths which are acceptable and/or required.
//
// Português:
//
// Um ConstrainULong especificando a largura do vídeo ou o intervalo de larguras que são aceitáveis e/ou obrigatórios.
func (e *Video) mountWidth(video *map[string]interface{}) {
	if e.Width.Exact != nil {
		(*video)["width"] = map[string]interface{}{"exact": e.Width.Exact}
		return
	}

	if e.Width.Max != nil || e.Width.Min != nil || e.Width.Ideal != nil {
		(*video)["width"] = make(map[string]interface{})

		if e.Width.Max != nil {
			(*video)["width"].(map[string]interface{})["max"] = e.Width.Max
		}

		if e.Width.Min != nil {
			(*video)["width"].(map[string]interface{})["min"] = e.Width.Min
		}

		if e.Width.Ideal != nil {
			(*video)["width"].(map[string]interface{})["ideal"] = e.Width.Ideal
		}

		return
	}

	if e.Width.Value != nil {
		(*video)["width"] = e.Width.Value
		return
	}
}

// mountResizeMode
//
// English:
//
// A ConstrainDOMString object specifying a mode or an array of modes the UA can use to derive the resolution of a
// video track. Allowed values are none and crop-and-scale. none means that the user agent uses the resolution
// provided by the camera, its driver or the OS. crop-and-scale means that the user agent can use cropping and
// downscaling on the camera output in order to satisfy other constraints that affect the resolution.
//
// Português:
//
// Um objeto ConstrainDOMString especificando um modo ou uma matriz de modos que o UA pode usar para derivar a
// resolução de uma trilha de vídeo. Os valores permitidos são none e crop-and-scale. none significa que o agente do
// usuário usa a resolução fornecida pela câmera, seu driver ou o sistema operacional. recortar e dimensionar
// significa que o agente do usuário pode usar recorte e redução de escala na saída da câmera para atender a outras
// restrições que afetam a resolução.
func (e *Video) mountResizeMode(video *map[string]interface{}) {
	if e.ResizeMode != nil {
		(*video)["resizeMode"] = e.ResizeMode
	}
}

func (e *Video) mount(video *map[string]interface{}) {
	if *video == nil {
		*video = make(map[string]interface{})
	}

	e.mountAspectRatio(video)
	e.mountFacingMode(video)
	e.mountFrameRate(video)
	e.mountHeight(video)
	e.mountWidth(video)
	e.mountResizeMode(video)
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
