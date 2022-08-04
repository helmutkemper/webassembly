package media

// SharedScreen
//
// English:
//
// These constraints apply to the video property of the object passed into getDisplayMedia() to obtain a stream for
// screen sharing.
//
// Português:
//
// Essas restrições se aplicam à propriedade video do objeto passado para getDisplayMedia() para obter um fluxo para
// compartilhamento de tela.
type SharedScreen struct {

	// DisplaySurface
	//
	// English:
	//
	// A ConstrainDOMString which specifies the types of display surface that may be selected by the user.
	//
	// Português:
	//
	// Um ConstrainDOMString que especifica os tipos de superfície de exibição que podem ser selecionados pelo usuário.
	//
	//https://developer.mozilla.org/en-US/docs/Web/API/MediaTrackConstraints/displaySurface
	DisplaySurface DisplaySurface //`js:"displaySurface"`

	// LogicalSurface
	//
	// English:
	//
	// A ConstrainBoolean value which may contain a single Boolean value or a set of them, indicating whether or not to
	// allow the user to choose source surfaces which do not directly correspond to display areas. These may include
	// backing buffers for windows to allow capture of window contents that are hidden by other windows in front of them,
	// or buffers containing larger documents that need to be scrolled through to see the entire contents in their
	// windows.
	//
	// Português:
	//
	// Um valor ConstrainBoolean que pode conter um único valor booleano ou um conjunto deles, indicando se deve ou não
	// permitir que o usuário escolha superfícies de origem que não correspondam diretamente às áreas de exibição.
	// Isso pode incluir buffers de backup para janelas para permitir a captura de conteúdo de janela oculto por outras
	// janelas à sua frente ou buffers contendo documentos maiores que precisam ser rolados para ver o conteúdo completo
	// em suas janelas.
	LogicalSurface BOOLEAN `js:"logicalSurface"`
}

// mountDisplaySurface
//
// English:
//
// A ConstrainDOMString which specifies the types of display surface that may be selected by the user.
//
// Português:
//
// Um ConstrainDOMString que especifica os tipos de superfície de exibição que podem ser selecionados pelo usuário.
//
// https://developer.mozilla.org/en-US/docs/Web/API/MediaTrackConstraints/displaySurface
func (e *SharedScreen) mountDisplaySurface(sharedScreen *map[string]interface{}) {
	if e.DisplaySurface != "" {
		(*sharedScreen)["displaySurface"] = e.DisplaySurface.String()
	}
}

// mountLogicalSurface
//
// English:
//
// A ConstrainBoolean value which may contain a single Boolean value or a set of them, indicating whether or not to
// allow the user to choose source surfaces which do not directly correspond to display areas. These may include
// backing buffers for windows to allow capture of window contents that are hidden by other windows in front of them,
// or buffers containing larger documents that need to be scrolled through to see the entire contents in their
// windows.
//
// Português:
//
// Um valor ConstrainBoolean que pode conter um único valor booleano ou um conjunto deles, indicando se deve ou não
// permitir que o usuário escolha superfícies de origem que não correspondam diretamente às áreas de exibição.
// Isso pode incluir buffers de backup para janelas para permitir a captura de conteúdo de janela oculto por outras
// janelas à sua frente ou buffers contendo documentos maiores que precisam ser rolados para ver o conteúdo completo
// em suas janelas.
func (e *SharedScreen) mountLogicalSurface(sharedScreen *map[string]interface{}) {
	if e.LogicalSurface.IsSet() {
		(*sharedScreen)["logicalSurface"] = e.LogicalSurface.Bool()
	}
}

func (e *SharedScreen) mount(sharedScreen *map[string]interface{}) {
	if *sharedScreen == nil {
		*sharedScreen = make(map[string]interface{})
	}

	e.mountDisplaySurface(sharedScreen)
	e.mountLogicalSurface(sharedScreen)
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
