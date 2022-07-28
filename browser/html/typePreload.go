package html

type Preload string

func (e Preload) String() string {
	return string(e)
}

const (

	// KPreloadNone
	//
	// English:
	//
	// Indicates that the video should not be preloaded.
	//
	// Português:
	//
	// Indica que o vídeo não deve ser precarregado.
	KPreloadNone Preload = "none"

	// KPreloadMetadata
	//
	// English:
	//
	// Indicates that only video metadata (e.g. length) is fetched.
	//
	// Português:
	//
	// Indica que apenas metadados de vídeo (por exemplo, duração) são buscados.
	KPreloadMetadata Preload = "metadata"

	// KPreloadAuto
	//
	// English:
	//
	// Indicates that the whole video file can be downloaded, even if the user is not expected to use it.
	//
	// Português:
	//
	// Indica que o conteúdo completo do arquivo de vídeo pode ser baixado, mesmo que não seja esperado que o usuário o
	// use.
	KPreloadAuto Preload = "auto"
)

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
