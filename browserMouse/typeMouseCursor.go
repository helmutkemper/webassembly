package browserMouse

type CursorType int

func (el CursorType) String() string {
	return cursorTypesWebBrowser[el]
}

var cursorTypesWebBrowser = [...]string{
	"cursor:auto",
	"cursor:alias",
	"cursor:all-scroll",
	"cursor:cell",
	"cursor:context-menu",
	"cursor:col-resize",
	"cursor:copy",
	"cursor:crosshair",
	"cursor:default",
	"cursor:e-resize",
	"cursor:ew-resize",
	"cursor:help",
	"cursor:move",
	"cursor:n-resize",
	"cursor:ne-resize",
	"cursor:nesw-resize",
	"cursor:ns-resize",
	"cursor:nw-resize",
	"cursor:nwse-resize",
	"cursor:no-drop",
	"cursor:none",
	"cursor:not-allowed",
	"cursor:pointer",
	"cursor:progress",
	"cursor:row-resize",
	"cursor:s-resize",
	"cursor:se-resize",
	"cursor:sw-resize",
	"cursor:text",
	//"cursor:URL",
	"cursor:vertical-text",
	"cursor:w-resize",
	"cursor:wait",
	"cursor:zoom-in",
	"cursor:zoom-out",
	//"cursor:initial",
	//"cursor:inherit",
}

const (
	// KCursorAuto
	//
	// English:
	//
	//  Default. The browser sets a cursor
	//
	// Português:
	//
	//  O navegador define o cursor
	KCursorAuto CursorType = iota // este valor sempre deve ser zero!

	// KCursorAlias
	//
	// English:
	//
	//  The cursor indicates an alias of something is to be created
	//
	// Português:
	//
	//  O cursor indica um atalho ou que algo pode ser criado
	KCursorAlias

	// KCursorAllScroll
	//
	// English:
	//
	//  The cursor indicates that something can be scrolled in any direction
	//
	// Português:
	//
	//  O cursor indica que algo pode ser corrido em qualquer direção
	KCursorAllScroll

	// KCursorCell
	//
	// English:
	//
	//  The cursor indicates that a cell (or set of cells) may be selected
	//
	// Português:
	//
	//  O cursor indica que uma célula ou conjunto de células pode ser selecionada
	KCursorCell

	// KCursorContextMenu
	//
	// English:
	//
	//  The cursor indicates that a context-menu is available
	//
	// Português:
	//
	//  O cursor indica que o menu de contexto está disponível
	KCursorContextMenu

	// KCursorColResize
	//
	// English:
	//
	//  The cursor indicates that the column can be resized horizontally
	//
	// Português:
	//
	// O cursor indica que a coluna pode ser redimensionada horizontalmente
	KCursorColResize

	// KCursorCopy
	//
	// English:
	//
	//  The cursor indicates something is to be copied
	//
	// Português:
	//
	//  O cursor indica que algo pode ser copiado
	KCursorCopy

	// KCursorCrossHair
	//
	// English:
	//
	//  The cursor render as a crosshair
	//
	// Português:
	//
	//  O cursor é transformado em uma cruz
	KCursorCrossHair

	// KCursorDefault
	//
	// English:
	//
	//  The default cursor
	//
	// Português:
	//
	//  Cursor padrão
	KCursorDefault

	// KCursorEResize
	//
	// English:
	//
	//  The cursor indicates that an edge of a box is to be moved right (east)
	//
	// Português:
	//
	//  O cursor indica que uma borda pode ser movida a direita (oeste)
	KCursorEResize

	// KCursorEWResize
	//
	// English:
	//
	//  Indicates a bidirectional resize cursor horizontal
	//
	// Português:
	//
	//  O cursor indica um redimensionamento horizontal bidirecional
	KCursorEWResize

	// KCursorHelp
	//
	// English:
	//
	//  The cursor indicates that help is available
	//
	// Português:
	//
	//  O cursor indica que há uma ajuda disponível
	KCursorHelp

	// KCursorMove
	//
	// English:
	//
	//  The cursor indicates something is to be moved
	//
	// Português:
	//
	//  O cursor indica que algo pode ser movido
	KCursorMove

	// KCursorNResize
	//
	// English:
	//
	//  The cursor indicates that an edge of a box is to be moved up (north)
	//
	// Português:
	//
	//  O cursor indica que uma borda pode ser movida para cima (norte)
	KCursorNResize

	// KCursorNeResize
	//
	// English:
	//
	//  The cursor indicates that an edge of a box is to be moved up and right (north/east)
	//
	// Português:
	//
	//  O cursor indica que uma borda pode ser movida a direita (norte/oeste)
	KCursorNeResize

	// KCursorNESwResize
	//
	// English:
	//
	//  Indicates a bidirectional resize cursor 45º right
	//
	// Português:
	//
	//  O cursor indica um redimensionamento bidirecional 45º a direita
	KCursorNESwResize

	// KCursorNSResize
	//
	// English:
	//
	//  Indicates a bidirectional resize cursor
	//
	// Português:
	//
	//  O cursor indica um redimensionamento bidirecional
	KCursorNSResize

	// KCursorNwResize
	//
	// English:
	//
	//  The cursor indicates that an edge of a box is to be moved up and left (north/west)
	//
	// Português:
	//
	//  O cursor indica que um canto de uma caixa pode ser movida para cima e para a esquerda
	//  (norte/oeste)
	KCursorNwResize

	// KCursorNWSeResize
	//
	// English:
	//
	//  Indicates a bidirectional resize cursor
	//
	// Português:
	//
	//  O cursor indica um redimensionamento bidirecional
	KCursorNWSeResize

	// KCursorNoDrop
	//
	// English:
	//
	//  The cursor indicates that the dragged item cannot be dropped here
	//
	// Português:
	//
	//  O cursor indica que um item arrastado não pode ser solto aqui
	KCursorNoDrop

	// KCursorNone
	//
	// English:
	//
	//  No cursor is rendered for the element
	//
	// Português:
	//
	//  Esconde o cursor
	KCursorNone

	// KCursorNotAllowed
	//
	// English:
	//
	//  The cursor indicates that the requested action will not be executed
	//
	// Português:
	//
	//  O cursor indica que a ação não pode ser executada
	KCursorNotAllowed

	// KCursorPointer
	//
	// English:
	//
	//  The cursor is a pointer and indicates a link
	//
	// Português:
	//
	//  O cursor é um ponteiro e indica um link
	KCursorPointer

	// KCursorProgress
	//
	// English:
	//
	//  The cursor indicates that the program is busy (in progress)
	//
	// Português:
	//
	//  O cursor indica que o programa está ocupado ou em progresso
	KCursorProgress

	// KCursorRowResize
	//
	// English:
	//
	//  The cursor indicates that the row can be resized vertically
	//
	// Português:
	//
	//  O cursor indica que uma linha pode ser redimensionada verticalmente
	KCursorRowResize

	// KCursorSResize
	//
	// English:
	//
	//  The cursor indicates that an edge of a box is to be moved down (south)
	//
	// Português:
	//
	//  O cursor indica que uma borda pode ser movida para baixo (sul)
	KCursorSResize

	// KCursorSeResize
	//
	// English:
	//
	//  The cursor indicates that an edge of a box is to be moved down and right (south/east)
	//
	// Português:
	//
	//  O cursor indica que um canto pode ser movida para baixo e direita (sul/oeste)
	KCursorSeResize

	// KCursorSwResize
	//
	// English:
	//
	//  The cursor indicates that an edge of a box is to be moved down and left (south/west)
	//
	// Português:
	//
	//  O cursor indica que um canto pode ser movido para baixo e esquerda (sul/oeste)
	KCursorSwResize

	// KCursorText
	//
	// English:
	//
	//  The cursor indicates text that may be selected
	//
	// Português:
	//
	//  O cursor indica que algo pode ser selecionado
	KCursorText

	// KCursorVerticalText
	//
	// English:
	//
	//  The cursor indicates vertical-text that may be selected
	//
	// Português:
	//
	//  O cursor indica que um texto vertical pode ser selecionado
	KCursorVerticalText

	// KCursorWResize
	//
	// English:
	//
	//  The cursor indicates that an edge of a box is to be moved left (west)
	//
	// Português:
	//
	//  O cursor indica que uma borda pode ser movida a esquerda (oeste)
	KCursorWResize

	// KCursorWait
	//
	// English:
	//
	//  The cursor indicates that the program is busy
	//
	// Português:
	//
	//  O cursor indica que o programa está ocupado
	KCursorWait

	// KCursorZoomIn
	//
	// English:
	//
	//  The cursor indicates that something can be zoomed in
	//
	// Português:
	//
	//  O cursor indica que algo pode receber zoom in
	KCursorZoomIn

	// KCursorZoomOut
	//
	// English:
	//
	//  The cursor indicates that something can be zoomed out
	//
	// Português:
	//
	//  O cursor indica que algo pode receber zoom out
	KCursorZoomOut

	// KCursorInitial
	//
	// English:
	//
	//  Sets this property to its default value. Read about initial
	//
	// Português:
	//
	//  Defina esta propriedade para definir um valor padrão.
	//
	// todo!
	// KCursorInitial

	// KCursorInherit
	//
	// English:
	//
	//  Inherits this property from its parent element. Read about inherit
	//
	// Português:
	//
	//  Herda esta propriedade de seu elemento pai. Leia sobre herança
	//
	// todo!
	// KCursorInherit

	// KCursorURL
	//
	// English:
	//
	//  A comma separated list of URLs to custom cursors. Note: Always specify a generic cursor at the
	//  end of the list, in case none of the URL-defined cursors can be used
	//
	// Português:
	//
	//  Uma lista de URLs separada por vírgula de cursores customizados. Nota: Sempre especifique um
	//  cursor genérico no final da lista ou use um cursor definido por URL
	//
	// todo:
	//
	// KCursorURL
)
