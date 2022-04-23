package html

type Decoding string

func (e Decoding) String() string {
	return string(e)
}

const (
	// KDecodingSync
	//
	// English:
	//
	//  Decode the image synchronously, for atomic presentation with other content.
	//
	// Português:
	//
	//  Decodifique a imagem de forma síncrona, para apresentação atômica com outro conteúdo.
	KDecodingSync Decoding = "sync"

	// KDecodingAsync
	//
	// English:
	//
	//  Decode the image asynchronously, to reduce delay in presenting other content.
	//
	// Português:
	//
	//  Decodifique a imagem de forma assíncrona, para reduzir o atraso na apresentação de outros
	//  conteúdos.
	KDecodingAsync Decoding = "async"

	// KDecodingAuto
	//
	// English:
	//
	//  No preference for the decoding mode. The browser decides what is best for the user. (Default)
	//
	// Português:
	//
	//  Sem preferência para o modo de decodificação. O navegador decide o que é melhor para o usuário.
	//  (Padrão)
	KDecodingAuto Decoding = "auto"
)
