package mouse

// Button
//
// English:
//
// Archive the pressed mouse button, e.g. KMouseButtonMain to the main button.
//
// Português:
//
// Arquiva o botão do mouse pressionado, por exeplo, KMouseButtonMain para o botão principal.
type Button int

const (
	// KMouseButtonNoButton
	//
	// English:
	//
	//  No button or un-initialized.
	//
	// Português:
	//
	//  Sem botão ou não inicializado.
	KMouseButtonNoButton Button = 0

	// KMouseButtonMain
	//
	// English:
	//
	//  Main button pressed, usually the left button or the un-initialized state.
	//
	// Português:
	//
	//  Botão principal pressionado, geralmente o botão esquerdo ou o estado não inicializado.
	KMouseButtonMain Button = 1

	// KMouseButtonAuxiliary
	//
	// English:
	//
	//  Auxiliary button pressed, usually the wheel button or the middle button (if present).
	//
	// Português:
	//
	//  Botão auxiliar pressionado, geralmente o botão da roda ou o botão do meio (se houver).
	KMouseButtonAuxiliary Button = 2

	// KMouseButtonSecondary
	//
	// English:
	//
	//  Secondary button pressed, usually the right button.
	//
	// Português:
	//
	//  Botão secundário pressionado, geralmente o botão direito.
	KMouseButtonSecondary Button = 4

	// KMouseButtonFourth
	//
	// English:
	//
	//  Fourth button, typically the Browser Back button.
	//
	// Português:
	//
	//  Quarto botão, normalmente o botão Voltar do navegador.
	KMouseButtonFourth Button = 8

	// KMouseButtonFifth
	//
	// English:
	//
	//  Fifth button, typically the Browser Forward button.
	//
	// Português:
	//
	//  Quinto botão, normalmente o botão Browser Forward.
	KMouseButtonFifth Button = 16
)
