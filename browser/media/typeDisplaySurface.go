package media

type DisplaySurface string

func (e DisplaySurface) String() string {
	return string(e)
}

const (

	// KDisplaySurfaceApplication
	//
	// English:
	//
	// The stream contains all of the windows of the application chosen by the user rendered into the one video track.
	//
	// Português:
	//
	// O fluxo contém todas as janelas do aplicativo escolhido pelo usuário renderizado em uma faixa de vídeo.
	KDisplaySurfaceApplication DisplaySurface = "application"

	// KDisplaySurfaceBrowser
	//
	// English:
	//
	// The stream contains the contents of a single browser tab selected by the user.
	//
	// Português:
	//
	// O fluxo contém o conteúdo de uma única guia do navegador selecionada pelo usuário.
	KDisplaySurfaceBrowser DisplaySurface = "browser"

	// KDisplaySurfaceMonitor
	//
	// English:
	//
	// The stream's video track contains the entire contents of one or more of the user's screens.
	//
	// Português:
	//
	// A faixa de vídeo do stream contém o conteúdo completo de uma ou mais telas do usuário.
	KDisplaySurfaceMonitor DisplaySurface = "monitor"

	// KDisplaySurfaceWindow
	//
	// English:
	//
	// The stream contains a single window selected by the user for sharing.
	//
	// Português:
	//
	// O fluxo contém uma única janela selecionada pelo usuário para compartilhamento.
	KDisplaySurfaceWindow DisplaySurface = "window"
)
