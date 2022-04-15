package html

// Mime
//
// Source:
//
//   https://www.iana.org/assignments/media-types/media-types.xhtml#audio
//
// English:
//
//  A MIME type (now properly called "media type", but also sometimes "content type") is a string sent
//  along with a file indicating the type of the file (describing the content format, for example, a
//  sound file might be labeled audio/ogg, or an image file image/png).
//
// It serves the same purpose as filename extensions traditionally do on Windows. The name originates
// from the MIME standard originally used in E-Mail.
//
// Português:
//
//  Um tipo MIME (agora chamado corretamente de "tipo de mídia", mas também às vezes "tipo de
//  conteúdo") é uma string enviada junto com um arquivo que indica o tipo do arquivo (descrevendo o
//  formato do conteúdo, por exemplo, um arquivo de som pode ser rotulado como audio/ogg , ou um
//  arquivo de imagem image/png).
//
// Ele serve ao mesmo propósito que as extensões de nome de arquivo tradicionalmente fazem no Windows.
// O nome se origina do padrão MIME originalmente usado em E-Mail.
type Mime string

func (e Mime) String() string {
	return string(e)
}
