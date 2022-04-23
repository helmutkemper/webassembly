package html

type Warp string

func (e Warp) String() string {
	return string(e)
}

const (
	// KWarpHard
	//
	// English:
	//
	//  The browser automatically inserts line breaks (CR+LF) so that each line has no more than the
	//  width of the control;
	//
	//   Note:
	//     * The cols attribute must also be specified for this to take effect.
	//
	// Português:
	//
	//  O navegador insere automaticamente quebras de linha (CR+LF) para que cada linha não tenha mais
	//  que a largura do controle;
	//
	//   Nota:
	//     * O atributo cols também deve ser especificado para que tenha efeito.
	KWarpHard Warp = "hard"

	// KWarpSoft
	//
	// English:
	//
	//  The browser ensures that all line breaks in the value consist of a CR+LF pair, but does not
	//  insert any additional line breaks.
	//
	// Português:
	//
	//  O navegador garante que todas as quebras de linha no valor consistam em um par CR+LF, mas não
	//  insere nenhuma quebra de linha adicional.
	KWarpSoft Warp = "soft"

	// KWarpOff
	//
	// English:
	//
	//  Like soft but changes appearance to white-space: pre so line segments exceeding cols are not
	//  wrapped and the <textarea> becomes horizontally scrollable.
	//
	//   Note:
	//     * Non-Standard
	//
	// Português:
	//
	//  Como soft, mas altera a aparência para espaço em branco: antes disso, os segmentos de linha que
	//  excedem as colunas não são quebrados e a <textarea> torna-se rolável horizontalmente.
	//
	//   Nota:
	//     * Não-Padrão
	KWarpOff Warp = "off"
)
