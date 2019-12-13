package mouse

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
	// en: Default. The browser sets a cursor
	//
	// pt_br: O navegador define o cursor
	KCursorAuto CursorType = iota

	// en: The cursor indicates an alias of something is to be created
	//
	// pt_br: O cursor indica um atalho ou que algo pode ser criado
	KCursorAlias

	// en: The cursor indicates that something can be scrolled in any direction
	//
	// pt_br: O cursor indica que algo pode ser corrido em qualquer direção
	KCursorAllScroll

	// en: The cursor indicates that a cell (or set of cells) may be selected
	//
	// pt_br: O cursor indica que uma célula ou conjunto de células pode ser
	// selecionada
	KCursorCell

	// en: The cursor indicates that a context-menu is available
	//
	// pt_br: O cursor indica que o menu de contexto está disponível
	KCursorContextMenu

	// en: The cursor indicates that the column can be resized horizontally
	//
	// pt_br: O cursor indica que a coluna pode ser redimensionada horizontalmente
	KCursorColResize

	// en: The cursor indicates something is to be copied
	//
	// pt_br: O cursor indica que algo pode ser copiado
	KCursorCopy

	// en: The cursor render as a crosshair
	//
	// pt_br: O cursor é transformado em uma cruz
	KCursorCrossHair

	// en: The default cursor
	//
	// pt_br: Cursor padrão
	KCursorDefault

	// ##
	// en: The cursor indicates that an edge of a box is to be moved right (east)
	//
	// pt_br: O cursor indica que uma borda pode ser movida a direita (oeste)
	KCursorEResize

	// ##
	// en: Indicates a bidirectional resize cursor horizontal
	//
	// pt_br: O cursor indica um redimensionamento horizontal bidirecional
	KCursorEWResize

	// en: The cursor indicates that help is available
	//
	// pt_br: O cursor indica que há uma ajuda disponível
	KCursorHelp

	// ##
	// en: The cursor indicates something is to be moved
	//
	// pt_br: O cursor indica que algo pode ser movido
	KCursorMove

	// ##
	// en: The cursor indicates that an edge of a box is to be moved up (north)
	//
	// pt_br: O cursor indica que uma borda pode ser movida para cima (norte)
	KCursorNResize

	// ##
	// en: The cursor indicates that an edge of a box is to be moved up and right
	// (north/east)
	//
	// pt_br: O cursor indica que uma borda pode ser movida a direita (norte/oeste)
	KCursorNeResize

	// ##
	// en: Indicates a bidirectional resize cursor 45º right
	//
	// pt_br: O cursor indica um redimensionamento bidirecional 45º a direita
	KCursorNESwResize

	// ##
	// en: Indicates a bidirectional resize cursor
	//
	// pt_br: O cursor indica um redimensionamento bidirecional
	KCursorNSResize

	// ##
	// en: The cursor indicates that an edge of a box is to be moved up and left
	// (north/west)
	//
	// pt_br: O cursor indica que um canto de uma caixa pode ser movida para cima e
	// para a esquerda (norte/oeste)
	KCursorNwResize

	// ##
	// en: Indicates a bidirectional resize cursor
	//
	// pt_br: O cursor indica um redimensionamento bidirecional
	KCursorNWSeResize

	// en: The cursor indicates that the dragged item cannot be dropped here
	//
	// pt_br: O cursor indica que um item arrastado não pode ser solto aqui
	KCursorNoDrop

	// en: No cursor is rendered for the element
	//
	// pt_br: Esconde o cursor
	KCursorNone

	// en: The cursor indicates that the requested action will not be executed
	//
	// pt_br: O cursor indica que a ação não pode ser executada
	KCursorNotAllowed

	// en: The cursor is a pointer and indicates a link
	//
	// pt_br: O cursor é um ponteiro e indica um link
	KCursorPointer

	// en: The cursor indicates that the program is busy (in progress)
	//
	// pt_br: O cursor indica que o programa está ocupado ou em progresso
	KCursorProgress

	// ##
	// en: The cursor indicates that the row can be resized vertically
	//
	// pt_br: O cursor indica que uma linha pode ser redimensionada verticalmente
	KCursorRowResize

	// ##
	// en: The cursor indicates that an edge of a box is to be moved down (south)
	//
	// pt_br: O cursor indica que uma borda pode ser movida para baixo (sul)
	KCursorSResize

	// ##
	// en: The cursor indicates that an edge of a box is to be moved down and right
	// (south/east)
	//
	// pt_br: O cursor indica que um canto pode ser movida para baixo e direita
	// (sul/oeste)
	KCursorSeResize

	// ##
	// en: The cursor indicates that an edge of a box is to be moved down and left
	// (south/west)
	//
	// pt_br: O cursor indica que um canto pode ser movido para baixo e esquerda
	// (sul/oeste)
	KCursorSwResize

	// en: The cursor indicates text that may be selected
	//
	// pt_br: O cursor indica que algo pode ser selecionado
	KCursorText

	// en: A comma separated list of URLs to custom cursors. Note: Always specify a
	// generic cursor at the end of the list, in case none of the URL-defined cursors
	// can be used
	//
	// pt_br: Uma lista de URLs separada por vírgula de cursores customizados. Nota:
	// Sempre especifique um cursor genérico no final da lista ou use um cursor
	// definido por URL
	// todo!
	// KCursorURL

	// en: The cursor indicates vertical-text that may be selected
	//
	// pt_br: O cursor indica que um texto vertical pode ser selecionado
	KCursorVerticalText

	// ##
	// en: The cursor indicates that an edge of a box is to be moved left (west)
	//
	// pt_br: O cursor indica que uma borda pode ser movida a esquerda (oeste)
	KCursorWResize

	// en: The cursor indicates that the program is busy
	//
	// pt_br: O cursor indica que o programa está ocupado
	KCursorWait

	// en: The cursor indicates that something can be zoomed in
	//
	// pt_br: O cursor indica que algo pode receber zoom in
	KCursorZoomIn

	// en: The cursor indicates that something can be zoomed out
	//
	// pt_br: O cursor indica que algo pode receber zoom out
	KCursorZoomOut

	// en: Sets this property to its default value. Read about initial
	//
	// pt_br: Defina esta propriedade para definir um valor padrão.
	// todo!
	// KCursorInitial

	// en: Inherits this property from its parent element. Read about inherit
	// todo!
	// KCursorInherit
)
