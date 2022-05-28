package html

type SvgDur string

func (e SvgDur) String() string {
	return string(e)
}

const (

	// KSvgDurMedia
	//
	// English:
	//
	//  This value specifies the simple duration as the intrinsic media duration. This is only valid for elements that
	//  define media. (For animation elements the attribute will be ignored if media is specified.)
	//
	// Portuguese
	//
	//  Esse valor especifica a duração simples como a duração da mídia intrínseca. Isso só é válido para elementos que
	//  definem mídia. (Para elementos de animação, o atributo será ignorado se a mídia for especificada.)
	KSvgDurMedia SvgDur = "media"

	// KSvgDurIndefinite
	//
	// English:
	//
	//  This value specifies the simple duration as indefinite.
	//
	// Portuguese
	//
	//  Este valor especifica a duração simples como indefinida.
	//
	KSvgDurIndefinite SvgDur = "indefinite"
)
