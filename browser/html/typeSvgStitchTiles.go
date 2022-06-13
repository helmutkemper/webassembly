package html

type SvgStitchTiles string

func (e SvgStitchTiles) String() string {
	return string(e)
}

const (
	// KSvgStitchTilesNoStitch
	//
	// English:
	//
	// This value indicates that no attempt is made to achieve smooth transitions at the border of tiles which contain a
	// turbulence function. Sometimes the result will show clear discontinuities at the tile borders.
	//
	// Português:
	//
	// Este valor indica que nenhuma tentativa é feita para obter transições suaves na borda dos ladrilhos que contêm uma
	// função de turbulência. Às vezes, o resultado mostrará descontinuidades claras nas bordas do ladrilho.
	KSvgStitchTilesNoStitch SvgStitchTiles = "noStitch"

	// KSvgStitchTilesStitch
	//
	// English:
	//
	// This value indicates that the user agent will automatically adjust the x and y values of the base frequency such
	// that the <feTurbulence> node's width and height (i.e., the width and height of the current subregion) contain an
	// integral number of the tile width and height for the first octave.
	//
	// Português:
	//
	// Esse valor indica que o agente do usuário ajustará automaticamente os valores x e y da frequência base de modo que
	// a largura e a altura do nó <feTurbulence> (ou seja, a largura e a altura da sub-região atual) contenham um número
	// inteiro da largura do bloco e altura para a primeira oitava.
	KSvgStitchTilesStitch SvgStitchTiles = "stitch"
)
