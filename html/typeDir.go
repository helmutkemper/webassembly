package html

// Dir
//
// English:
//
//  Specifies the text direction for the content in an element.
//
// Português:
//
//  Especifica a direção do texto para o conteúdo em um elemento.
type Dir int

func (e Dir) String() (element string) {
	if e == 0 {
		return
	}

	return dirString[e]
}

const (
	// KDirLeftToRight
	//
	// English:
	//
	//  Default. Left-to-right text direction.
	//
	// Português:
	//
	//  Padrão. Direção do texto da esquerda para a direita.
	KDirLeftToRight Dir = iota + 1

	// KDirRightToLeft
	//
	// English:
	//
	//  Right-to-left text direction.
	//
	// Português:
	//
	//  Direção do texto da direita para a esquerda.
	KDirRightToLeft

	// KDirAuto
	//
	// English:
	//
	//  Let the browser figure out the text direction, based on the content (only recommended if the
	//  text direction is unknown)
	//
	// Português:
	//
	//  Deixe o navegador descobrir a direção do texto, com base no conteúdo (recomendado apenas se a
	//  direção do texto for desconhecida)
	KDirAuto
)

var dirString = [...]string{
	"",
	"ltr",
	"rtl",
	"auto",
}
