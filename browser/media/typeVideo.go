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
	AspectRatio ConstrainDouble `js:"aspectRatio"`

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
	FacingMode interface{} `js:"facingMode"`

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
	FrameRate ConstrainDouble `js:"frameRate"`

	// Height
	//
	// English:
	//
	// A ConstrainULong specifying the video height or range of heights which are acceptable and/or required.
	//
	// Português:
	//
	// Um ConstrainULong especificando a altura do vídeo ou o intervalo de alturas que são aceitáveis e/ou obrigatórios.
	Height ConstrainULong `js:"height"`

	// Width
	//
	// English:
	//
	// A ConstrainULong specifying the video width or range of widths which are acceptable and/or required.
	//
	// Português:
	//
	// Um ConstrainULong especificando a largura do vídeo ou o intervalo de larguras que são aceitáveis e/ou obrigatórios.
	Width ConstrainULong `js:"width"`

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
	ResizeMode interface{} `js:"resizeMode"`
}
