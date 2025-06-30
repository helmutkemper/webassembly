package block

import (
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/browser/stage"
	"github.com/helmutkemper/webassembly/examples/ide/manager"
	"github.com/helmutkemper/webassembly/examples/ide/managerCollision"
	"github.com/helmutkemper/webassembly/examples/ide/ornament"
	"github.com/helmutkemper/webassembly/examples/ide/rulesDensity"
	"github.com/helmutkemper/webassembly/examples/ide/rulesSequentialId"
	"github.com/helmutkemper/webassembly/examples/ide/rulesStage"
	"github.com/helmutkemper/webassembly/examples/ide/utils"
	"github.com/helmutkemper/webassembly/platform/components"
	"github.com/helmutkemper/webassembly/platform/factoryColor"
	"image/color"
	"math"
	"strconv"
	"syscall/js"
)

type Block struct {
	id     string
	autoId string
	name   string

	x      rulesDensity.Density
	y      rulesDensity.Density
	width  rulesDensity.Density
	height rulesDensity.Density

	dragDeltaTop  rulesDensity.Density
	dragDeltaLeft rulesDensity.Density

	resizeLimitTop    rulesDensity.Density
	resizeLimitBottom rulesDensity.Density
	resizeLimitLeft   rulesDensity.Density
	resizeLimitRight  rulesDensity.Density

	initialized bool

	blockMinimumWidth  rulesDensity.Density
	blockMinimumHeight rulesDensity.Density

	classListName string

	isResizing    bool
	resizeEnabled bool
	resizeLocked  bool
	selectEnable  bool
	selectLocked  bool
	dragEnabled   bool
	dragLocked    bool

	resizerColor     color.RGBA
	resizerLine      []int
	resizerLineWidth int

	ideStage          *html.TagSvg
	block             *html.TagSvg
	selectDivAppended bool
	selectDiv         *html.TagSvgRect
	main              *html.TagSvg

	resizerButton ResizeButton
	draggerButton ResizeButton

	resizerTopLeft     ResizeButton
	resizerTopRight    ResizeButton
	resizerBottomLeft  ResizeButton
	resizerBottomRight ResizeButton

	resizerTopMiddle    ResizeButton
	resizerBottomMiddle ResizeButton
	resizerLeftMiddle   ResizeButton
	resizerRightMiddle  ResizeButton

	resizerMoveBorderLimit rulesDensity.Density

	draggerTopMiddle    ResizeButton
	draggerBottomMiddle ResizeButton
	draggerLeftMiddle   ResizeButton
	draggerRightMiddle  ResizeButton

	ornament ornament.Draw

	warningMarkAppended bool
	warningMark         ornament.WarningMark
	warningMarkEnabled  bool

	onResizeFunc func(args []js.Value, width, height rulesDensity.Density)

	gridAdjust rulesStage.GridAdjust

	ruleBook         map[string]func()
	ruleAdjustToGrid bool
}

// initRuleBook
//
// English:
//
//	Organises complex rules, mainly business rules and visual rules.
//
//	All rules must be straightforward and respect the single responsibility of the function, and the function must be
//	self-contained, that is, enabling something already enabled does not have an adverse effect on the
//	functioning of the code.
//
//	All functions must be straightforward.
//
// Português:
//
//	Organiza as regras complexas, principalmente regras de negócios e regras visuais.
//
//	Todas as regras devem ser simples e respeitar a responsabilidade única da função, e a função deve ser
//	auto contida, ou seja, habilitar algo que já está habilitado não gera efeito adverso ao funcionamento do código.
//
//	Todas as funções devem ser simples
func (e *Block) initRuleBook() {
	e.ruleBook = make(map[string]func())

	e.ruleBook["onInit"] = func() {
		// English:   Within the rules block, do not delete the commented functions, just comment or uncomment
		//            the desired function.
		// Português: Dentro do bloco de regras, não apague as funções comentadas, apenas comente ou descomente a
		//            função desejada.

		// Register the device in the general record of the system
		e.Register()
	}

	// Rule: adjustToGrid
	//
	// English:
	//
	//  Snaps the top-left corner and bottom-right corner to the stage positioning grid.
	//
	// Português:
	//
	//  Ajusta a ponta top-left e a ponta bottom-right ao grid de posicionamento do palco.
	e.ruleBook["adjustToGrid"] = func() {
		// English:   Within the rules block, do not delete the commented functions, just comment or uncomment
		//            the desired function.
		// Português: Dentro do bloco de regras, não apague as funções comentadas, apenas comente ou descomente a
		//            função desejada.

		// Adjusts coordinates (x,y) and (width,height) to the center of the grid
		e.adjustToGridRuleOn()
	}

	// Rule: setWarningOn
	//
	// English:
	//
	//  Activates the warning mark that something is wrong.
	//
	// Português:
	//
	//  Ativa a marca de advertência de que alguma coisa está errada.
	e.ruleBook["setWarningOn"] = func() {
		// English:   Within the rules block, do not delete the commented functions, just comment or uncomment
		//            the desired function.
		// Português: Dentro do bloco de regras, não apague as funções comentadas, apenas comente ou descomente a
		//            função desejada.

		e.setWarningOn()
		e.setWarningFlashOn()
	}

	// Rule: setWarningOff
	//
	// English:
	//
	//  Deactivates the warning mark that something is wrong.
	//
	// Português:
	//
	//  Desativa a marca de advertência de que alguma coisa está errada.
	e.ruleBook["setWarningOff"] = func() {
		// English:   Within the rules block, do not delete the commented functions, just comment or uncomment
		//            the desired function.
		// Português: Dentro do bloco de regras, não apague as funções comentadas, apenas comente ou descomente a
		//            função desejada.

		// This rule may be called by another rule, so update the status here
		//
		// Pode ser que esta regra seja chamada por outra regra, por isto, atualizar o status aqui
		e.warningMarkEnabled = false

		e.setWarningOff()
		e.setWarningFlashOff()
	}

	// Rule: setDragOn
	//
	// English:
	//
	//  Enables the drag repositioning tool
	//
	// Português:
	//
	//  Habilita a ferramenta de reposicionamento por arrasto
	e.ruleBook["setDragEnableOn"] = func() {
		// English:   Within the rules block, do not delete the commented functions, just comment or uncomment
		//            the desired function.
		// Português: Dentro do bloco de regras, não apague as funções comentadas, apenas comente ou descomente a
		//            função desejada.

		e.ruleBook["setResizeOff"]()
		e.ruleBook["setSelectOff"]()

		e.setDragOrnamentOn()
		e.setSelectOrnamentAttentionColorOn()
		e.setDragCursorOn()
		e.selectForDragOn()
	}

	e.ruleBook["setOnDraggingEvent"] = func() {
		// English:   Within the rules block, do not delete the commented functions, just comment or uncomment
		//            the desired function.
		// Português: Dentro do bloco de regras, não apague as funções comentadas, apenas comente ou descomente a
		//            função desejada.

		//e.draggingMoveDraggingSelectedOn()
		e.draggingMoveSelectedOnStage()
	}

	// Rule: setDragOff
	//
	// English:
	//
	//  Disables the drag repositioning tool
	//
	// Português:
	//
	//  Desabilita a ferramenta de reposicionamento por arrasto
	e.ruleBook["setDragEnableOff"] = func() {
		// English:   Within the rules block, do not delete the commented functions, just comment or uncomment
		//            the desired function.
		// Português: Dentro do bloco de regras, não apague as funções comentadas, apenas comente ou descomente a
		//            função desejada.

		// This rule may be called by another rule, so update the status here
		//
		// Pode ser que esta regra seja chamada por outra regra, por isto, atualizar o status aqui
		e.dragEnabled = false

		e.setDragOrnamentOff()
		e.setSelectOrnamentAttentionColorOff()
		e.setDragCursorOff()
		e.selectForDragOff()
	}

	// Rule: setResizeOn
	//
	// English:
	//
	//  Enables the resizing tool
	//
	// Português:
	//
	//  Habilita a ferramenta de redimensionamento
	e.ruleBook["setResizeOn"] = func() {
		// English:   Within the rules block, do not delete the commented functions, just comment or uncomment
		//            the desired function.
		// Português: Dentro do bloco de regras, não apague as funções comentadas, apenas comente ou descomente a
		//            função desejada.

		e.ruleBook["setDragEnableOff"]()
		e.ruleBook["setSelectOff"]()

		e.setResizeOrnamentVisibleOn()
	}

	// Rule: setResizingOn
	//
	// English:
	//
	//  Event occurs during resizing
	//
	// Português:
	//
	//  Evento ocorre durante o redimensionamento
	e.ruleBook["setResizingOn"] = func() {
		// English:   Within the rules block, do not delete the commented functions, just comment or uncomment
		//            the desired function.
		// Português: Dentro do bloco de regras, não apague as funções comentadas, apenas comente ou descomente a
		//            função desejada.

		e.resizerMoveBorderLimit = rulesDensity.Convert(30)
		e.calculateLimitForResizeOn()
	}

	// Rule: setResizeOff
	//
	// English:
	//
	//  Disables the resizing tool
	//
	// Português:
	//
	//  Desabilita a ferramenta de redimensionamento
	e.ruleBook["setResizeOff"] = func() {
		// English:   Within the rules block, do not delete the commented functions, just comment or uncomment
		//            the desired function.
		// Português: Dentro do bloco de regras, não apague as funções comentadas, apenas comente ou descomente a
		//            função desejada.

		// This rule may be called by another rule, so update the status here
		//
		// Pode ser que esta regra seja chamada por outra regra, por isto, atualizar o status aqui
		e.resizeEnabled = false

		e.setResizeOrnamentVisibleOff()
	}

	// Rule: setSelectOn
	//
	// English:
	//
	//  Enables the selection tool
	//
	// Português:
	//
	//  Habilita a ferramenta de seleção
	e.ruleBook["setSelectOn"] = func() {
		// English:   Within the rules block, do not delete the commented functions, just comment or uncomment
		//            the desired function.
		// Português: Dentro do bloco de regras, não apague as funções comentadas, apenas comente ou descomente a
		//            função desejada.

		e.ruleBook["setDragEnableOff"]()
		e.ruleBook["setResizeOff"]()

		e.setSelectRectangleOrnamentOn()
		e.setSelectOrnamentAttentionColorOn()
	}

	// Rule: setSelectOff
	//
	// English:
	//
	//  Disables the selection tool
	//
	// Português:
	//
	//  Desabilita a ferramenta de seleção
	e.ruleBook["setSelectOff"] = func() {
		// English:   Within the rules block, do not delete the commented functions, just comment or uncomment
		//            the desired function.
		// Português: Dentro do bloco de regras, não apague as funções comentadas, apenas comente ou descomente a
		//            função desejada.

		// This rule may be called by another rule, so update the status here
		//
		// Pode ser que esta regra seja chamada por outra regra, por isto, atualizar o status aqui
		e.selectEnable = false

		e.setSelectRectangleOrnamentOff()
		e.setSelectOrnamentAttentionColorOff()
	}

	e.ruleBook["adjustToGrid"]()
}

// SetWarningMark
//
// English:
//
//	Defines the instance responsible for warning ornamentation, in case of error.
//
// Português:
//
//	Define a instância responsável pela ornamentação de atenção, em caso de erro.
func (e *Block) SetWarningMark(warningMark ornament.WarningMark) {
	e.warningMark = warningMark
}

// GetWarning
//
// English:
//
//	Returns the visibility of the warning mark
//
// Português:
//
//	Retorna a visibilidade da marca de aviso
func (e *Block) GetWarning() (warning bool) {
	return e.warningMarkEnabled
}

// SetWarning
//
// English:
//
//	Sets the visibility of the warning mark
//
// Português:
//
//	Define a visibilidade da marca de aviso
func (e *Block) SetWarning(warning bool) {
	e.warningMarkEnabled = warning
	if warning {
		e.ruleBook["setWarningOn"]()
		return
	}

	e.ruleBook["setWarningOff"]()
}

// setWarningOn
//
// English:
//
//	Makes the warning mark visible and adjusts the dimensions and coordinates
//
// Português:
//
//	Faz a marca de aviso ficar visível e ajusta as dimensões e coordenadas
func (e *Block) setWarningOn() {
	e.warningMarkAppended = true
	e.block.Append(e.warningMark.GetSvg())
	_ = e.warningMark.Update(e.x, e.y, e.width, e.height)
}

// setWarningOff
//
// English:
//
//	Hides the warning mark
//
// Português:
//
//	Esconde a marca de aviso
func (e *Block) setWarningOff() {
	if !e.warningMarkAppended {
		return
	}

	e.warningMarkAppended = false
	e.block.Get().Call("removeChild", e.warningMark.GetSvg().Get())
}

// setWarningFlashOn
//
// English:
//
//	Makes the warning mark blink to get the user's attention, usually an easter egg
//
// Português:
//
//	Faz a marca de aviso piscar para chamar atenção do usuário, geralmente um easter egg
func (e *Block) setWarningFlashOn() {
	e.warningMark.Flash(true)
}

// setWarningFlashOff
//
// English:
//
//	Stops the warning mark from flashing
//
// Português:
//
//	Faz a marca de aviso parar de piscar
func (e *Block) setWarningFlashOff() {
	e.warningMark.Flash(false)
}

// SetMainSvg
//
// English:
//
//	Defines the Main svg, where the stage is drawn
//
// Português:
//
//	Define o svg Principal, onde o palco é desenhado
func (e *Block) SetMainSvg(svg *html.TagSvg) {
	e.main = svg
}

// SetResizerButton
//
// English:
//
//	Defines the instance object responsible for drawing the resize buttons
//
// Português:
//
//	Define o objeto da instância responsável por desenhar os botões de redimensionamento
func (e *Block) SetResizerButton(resizerButton ResizeButton) {
	e.resizerButton = resizerButton
}

// SetDraggerButton
//
// English:
//
//	Defines the instance object responsible for drawing the drag indicator buttons
//
// Português:
//
//	Define o objeto da instância responsável por desenhar os botões indicadores de arrasto
func (e *Block) SetDraggerButton(draggerButton ResizeButton) {
	e.draggerButton = draggerButton
}

// SetGridAdjust
//
// English:
//
//	Makes dragged content snap to grid
//
// Português:
//
//	Faz o conteúdo arrastado ser ajustado ao grid
func (e *Block) SetGridAdjust(gridAdjust rulesStage.GridAdjust) {
	e.gridAdjust = gridAdjust
}

// adjustToGridRuleOn
//
// English:
//
//	Adjusts coordinates (x, y) and (width, height) to the center of the grid
//
// Português:
//
//	Ajusta as coordenadas (x, y) e (width, height) ao centro do grid
func (e *Block) adjustToGridRuleOn() {
	e.ruleAdjustToGrid = true
}

// adjustXYToGrid
//
// English:
//
//	Receives (x, y) and returns the center coordinate of the grid element
//
// Português:
//
//	Recebe (x, y) e devolve a coordenada de centro do elemento do grid
func (e *Block) adjustXYToGrid(x, y int) (cx, cy int) {
	if e.ruleAdjustToGrid {
		return e.gridAdjust.AdjustCenter(x, y)
	}
	return x, y
}

// GetInitialized
//
// English:
//
//	Returns if the instance is ready for use
//
// Português:
//
//	Retorna se a instância está pronta para uso
func (e *Block) GetInitialized() (initialized bool) {
	return e.initialized
}

// SetDragBlocked
//
// English:
//
//	Prevents the user from enabling the drag tool
//
// Português:
//
//	Impede o usuário de habilitar a ferramenta de arrasto
func (e *Block) SetDragBlocked(blocked bool) {
	e.dragLocked = blocked
}

// GetDragBlocked Return the drag tool enable status

// GetDragBlocked
//
// English:
//
//	Returns whether the user is blocked from using the drag tool
//
// Português:
//
//	Retorna se o usuário está impedido de usar a ferramenta de arrasto
func (e *Block) GetDragBlocked() (blocked bool) {
	return e.dragLocked
}

// GetDragEnable
//
// English:
//
//	Return the drag tool status
//
// Português:
//
//	Retorne o status da ferramenta de arrasto
func (e *Block) GetDragEnable() (enabled bool) {
	return e.dragEnabled
}

// setDragOrnamentOn
//
// English:
//
// Português:
func (e *Block) setDragOrnamentOn() {
	if !e.initialized {
		return
	}

	e.draggerTopMiddle.SetVisible(true, e.ideStage)
	e.draggerRightMiddle.SetVisible(true, e.ideStage)
	e.draggerBottomMiddle.SetVisible(true, e.ideStage)
	e.draggerLeftMiddle.SetVisible(true, e.ideStage)
}

// English:
//
// Português:
func (e *Block) setDragOrnamentOff() {
	if !e.initialized {
		return
	}

	e.draggerTopMiddle.SetVisible(false, e.ideStage)
	e.draggerRightMiddle.SetVisible(false, e.ideStage)
	e.draggerBottomMiddle.SetVisible(false, e.ideStage)
	e.draggerLeftMiddle.SetVisible(false, e.ideStage)
}

// English:
//
// Português:
func (e *Block) setDragLockedOn() {
	e.dragLocked = true
}

// English:
//
// Português:
func (e *Block) setDragLockedOff() {
	e.dragLocked = false
}

// SetDrag Enables the device's drag tool

// English:
//
// Português:
func (e *Block) SetDragEnable(enabled bool) {
	if e.dragLocked {
		e.dragEnabled = false // todo: fazer a regra
		return
	}

	e.dragEnabled = enabled

	if e.dragEnabled {
		e.ruleBook["setDragEnableOn"]()
		return
	}

	e.ruleBook["setDragEnableOff"]()
}

// English:
//
// Português:
func (e *Block) setDragCursorOn() {
	if !e.initialized {
		return
	}

	e.block.AddStyleConditional(true, "cursor", "grab", "")
}

// English:
//
// Português:
func (e *Block) setDragCursorOff() {
	if !e.initialized {
		return
	}

	e.block.AddStyleConditional(false, "cursor", "grab", "")
}

// ResizeInverter Invert the resize tool status

// English:
//
// Português:
func (e *Block) ResizeInverter() {
	e.resizeEnabled = !e.resizeEnabled
}

// GetResizeEnable Return the resize tool status

// English:
//
// Português:
func (e *Block) GetResizeEnable() (enabled bool) {
	return e.resizeEnabled
}

// SetResizeEnable Defines the resize tool status

// English:
//
// Português:
func (e *Block) SetResizeEnable(enabled bool) {
	if e.resizeLocked {
		e.resizeEnabled = false // todo: fazer a regra
		return
	}

	e.resizeEnabled = enabled

	if e.resizeEnabled {
		e.ruleBook["setResizeOn"]()
		return
	}

	e.ruleBook["setResizeOff"]()

	//if enabled && e.selectEnable {
	//	e.SetSelected(false) // todo: fazer a regra
	//}
}

// ResizeBlockedInvert Invert the status from disables resize tool. Note: Used in the menu

// English:
//
// Português:
func (e *Block) ResizeBlockedInvert() {
	e.resizeLocked = !e.resizeLocked
}

// GetResizeBlocked Return the status from disables resize tool

// English:
//
// Português:
func (e *Block) GetResizeBlocked() (blocked bool) {
	return e.resizeLocked
}

// SetResizeBlocked Disables the resize tool

// English:
//
// Português:
func (e *Block) SetResizeBlocked(blocked bool) {
	e.resizeLocked = blocked
}

// SelectBlockedInvert Invert the status of the selection tool lock. Note: Used in the menu

// English:
//
// Português:
func (e *Block) SelectBlockedInvert() {
	e.selectLocked = !e.selectLocked
}

// GetSelectBlocked Returns the status of the selection tool lock

// English:
//
// Português:
func (e *Block) GetSelectBlocked() (blocked bool) {
	return e.selectLocked
}

// SetSelectBlocked Lock the selection tool

// English:
//
// Português:
func (e *Block) SetSelectBlocked(blocked bool) {
	e.selectLocked = blocked
}

// SelectedInvert Invert the status of the selection tool. Note: Used in the menu

// English:
//
// Português:
func (e *Block) SelectedInvert() {
	e.SetSelected(!e.selectEnable)
}

// English:
//
// Português:
func (e *Block) setSelectRectangleOrnamentOn() {
	if !e.initialized {
		return
	}

	e.selectDivAppended = true
	e.ideStage.Append(e.selectDiv)

	e.selectDiv.SetZIndex(stage.GetNextZIndex())
}

// English:
//
// Português:
func (e *Block) setSelectRectangleOrnamentOff() {
	if !e.initialized {
		return
	}

	if !e.selectDivAppended {
		return
	}

	e.selectDivAppended = false
	e.selectDiv.RemoveZIndex()
	e.ideStage.Get().Call("removeChild", e.selectDiv.Get()) //todo: colocar tag
}

// English:
//
// Português:
func (e *Block) setSelectOrnamentAttentionColorOn() {
	if !e.initialized {
		return
	}

	e.ornament.SetSelected(true)
}

// English:
//
// Português:
func (e *Block) setSelectOrnamentAttentionColorOff() {
	if !e.initialized {
		return
	}

	e.ornament.SetSelected(false)
}

// SetSelected Defines if the device selection tool is active

// English:
//
// Português:
func (e *Block) SetSelected(selected bool) {
	e.selectEnable = selected

	if e.selectLocked {
		e.selectEnable = false // todo: fazer a regra
		return
	}

	if e.selectEnable {
		e.ruleBook["setSelectOn"]()
		return
	}

	e.ruleBook["setSelectOff"]()

	//e.SetResize(false) // todo: fazer a regra
}

// GetSelected Return the select tool status

// English:
//
// Português:
func (e *Block) GetSelected() (selected bool) {
	return e.selectEnable
}

// English:
//
// Português:
func (e *Block) GetZIndex() (zIndex int) {
	z := e.block.Get().Call("getAttribute", "zIndex").String()
	zStr, _ := strconv.Atoi(z)
	return zStr
}

// createBlock Prepare all divs and CSS

// English:
//
// Português:
func (e *Block) createBlock(x, y, width, height rulesDensity.Density) {
	e.block = factoryBrowser.NewTagSvg().
		Id(e.id).
		X(x.GetInt()).
		Y(y.GetInt()).
		Width(width.GetInt()).
		Height(height.GetInt()).
		SetZIndex(stage.GetNextZIndex())

	// Append, js appendChild, it should be used only in the necessary elements on the stage.
	// Any other visual element should be attached only when necessary.
	e.ideStage.Append(e.block)

	e.selectDiv = factoryBrowser.NewTagSvgRect().
		X(x.GetInt()).
		Y(y.GetInt()).
		Width(width.GetInt()).
		Height(height.GetInt()).
		Fill("none").Stroke(e.resizerColor).
		StrokeDasharray(e.resizerLine).
		StrokeWidth(rulesDensity.Density(e.resizerLineWidth).GetInt()).
		SetZIndex(stage.GetNextZIndex())

	e.resizerTopLeft = e.resizerButton.GetNew()
	e.resizerTopLeft.SetName("top-left")
	e.resizerTopLeft.SetCursor("nwse-resize")
	e.resizerTopLeft.SetCX(x - e.resizerTopLeft.GetSpace())
	e.resizerTopLeft.SetCY(y - e.resizerTopLeft.GetSpace())

	e.resizerTopRight = e.resizerButton.GetNew()
	e.resizerTopRight.SetName("top-right")
	e.resizerTopRight.SetCursor("nesw-resize")
	e.resizerTopRight.SetCX(x + width + e.resizerTopRight.GetSpace())
	e.resizerTopRight.SetCY(y - e.resizerTopRight.GetSpace())

	e.resizerBottomLeft = e.resizerButton.GetNew()
	e.resizerBottomLeft.SetName("bottom-left")
	e.resizerBottomLeft.SetCursor("nesw-resize")
	e.resizerBottomLeft.SetCX(x - e.resizerBottomLeft.GetSpace())
	e.resizerBottomLeft.SetCY(y + height + e.resizerBottomLeft.GetSpace())

	e.resizerBottomRight = e.resizerButton.GetNew()
	e.resizerBottomRight.SetName("bottom-right")
	e.resizerBottomRight.SetCursor("nwse-resize")
	e.resizerBottomRight.SetCX(x + width + e.resizerBottomRight.GetSpace())
	e.resizerBottomRight.SetCY(y + height + e.resizerBottomRight.GetSpace())

	//----------------------------------------------------

	e.resizerTopMiddle = e.resizerButton.GetNew()
	e.resizerTopMiddle.SetName("top-middle")
	e.resizerTopMiddle.SetCursor("ns-resize")
	e.resizerTopMiddle.SetCX(x + width/2)
	e.resizerTopMiddle.SetCY(y - e.resizerTopMiddle.GetSpace())

	e.resizerBottomMiddle = e.resizerButton.GetNew()
	e.resizerBottomMiddle.SetName("bottom-middle")
	e.resizerBottomMiddle.SetCursor("ns-resize")
	e.resizerBottomMiddle.SetCX(x + width/2)
	e.resizerBottomMiddle.SetCY(y + height + e.resizerBottomMiddle.GetSpace())

	e.resizerLeftMiddle = e.resizerButton.GetNew()
	e.resizerLeftMiddle.SetName("left-middle")
	e.resizerLeftMiddle.SetCursor("ew-resize")
	e.resizerLeftMiddle.SetCX(x - e.resizerLeftMiddle.GetSpace())
	e.resizerLeftMiddle.SetCY(y + height/2)

	e.resizerRightMiddle = e.resizerButton.GetNew()
	e.resizerRightMiddle.SetName("right-middle")
	e.resizerRightMiddle.SetCursor("ew-resize")
	e.resizerRightMiddle.SetCX(x + width + e.resizerRightMiddle.GetSpace())
	e.resizerRightMiddle.SetCY(y + height/2)

	//----------------------------------------------------------

	e.draggerTopMiddle = e.draggerButton.GetNew()
	e.draggerTopMiddle.SetName("top-middle")
	e.draggerTopMiddle.SetCX(x + width/2)
	e.draggerTopMiddle.SetCY(y - e.draggerTopMiddle.GetSpace())
	e.draggerTopMiddle.SetRotation(-90)

	e.draggerBottomMiddle = e.draggerButton.GetNew()
	e.draggerBottomMiddle.SetName("bottom-middle")
	e.draggerBottomMiddle.SetCX(x + width/2)
	e.draggerBottomMiddle.SetCY(y + height + e.draggerBottomMiddle.GetSpace())
	e.draggerBottomMiddle.SetRotation(90)

	e.draggerLeftMiddle = e.draggerButton.GetNew()
	e.draggerLeftMiddle.SetName("left-middle")
	e.draggerLeftMiddle.SetCX(x - e.draggerLeftMiddle.GetSpace())
	e.draggerLeftMiddle.SetCY(y + height/2)
	e.draggerLeftMiddle.SetRotation(180)

	e.draggerRightMiddle = e.draggerButton.GetNew()
	e.draggerRightMiddle.SetName("right-middle")
	e.draggerRightMiddle.SetCX(x + width + e.draggerRightMiddle.GetSpace())
	e.draggerRightMiddle.SetCY(y + height/2)
	e.draggerRightMiddle.SetRotation(0)
}

// GetDeviceDiv Returns the div from device

// English:
//
// Português:
func (e *Block) GetDeviceDiv() (element *html.TagSvg) {
	return e.block
}

// GetHeight returns the current height of the device.

// English:
//
// Português:
func (e *Block) GetHeight() (height rulesDensity.Density) {
	return e.height
}

// GetID Returns the device's div ID

// English:
//
// Português:
func (e *Block) GetID() (id string) {
	return e.id
}

// GetIdeStage Returns to Div where IDE is drawn

// English:
//
// Português:
func (e *Block) GetIdeStage() (ideStage *html.TagSvg) {
	return e.ideStage
}

// GetName Returns the single name of the device

// English:
//
// Português:
func (e *Block) GetName() (name string) {
	return e.name
}

// GetWidth returns the current width of the device.

// English:
//
// Português:
func (e *Block) GetWidth() (width rulesDensity.Density) {
	return e.width
}

// GetX Returns to coordinate X of the browser screen

// English:
//
// Português:
func (e *Block) GetX() (x rulesDensity.Density) {
	return e.x
}

// GetY Returns to coordinate Y of the browser screen

// English:
//
// Português:
func (e *Block) GetY() (y rulesDensity.Density) {
	return e.y
}

// Init Initializes the generic functions of the device

// English:
//
// Português:
func (e *Block) Init() (err error) {
	e.initRuleBook()
	e.resetLimitForResize()

	var id string
	id = rulesSequentialId.GetIdFromBase(e.name)
	if e.id, err = utils.VerifyUniqueId(id); err != nil {
		return
	}

	e.autoId = utils.GetRandomId()

	e.classListName = "block"

	e.resizerColor = factoryColor.NewRed() // todo: organizar - início
	e.resizerLine = []int{16, 4}
	e.resizerLineWidth = 3 // todo: organizar - fim

	e.createBlock(e.x, e.y, e.width, e.height)
	e.initEvents()

	e.initialized = true

	if e.ornament != nil {
		svg := e.ornament.GetSvg()
		e.block.Append(svg)

		_ = e.ornament.Update(e.x, e.y, e.width, e.height)
	}

	e.ruleBook["onInit"]()
	return
}

// Register
//
// English:
//
//	Records the element in the global identification record
//
// Português:
//
//	Registra o elemento no registro global de identificação
func (e *Block) Register() {
	manager.Manager.Register(e)
}

// Unregister
//
// English:
//
//	Unrecords the element in the global identification record
//
// Português:
//
//	Remove o registro do elemento no registro global de identificação
func (e *Block) Unregister() {
	_ = manager.Manager.Unregister(e)
}

// English:
//
// Português:
func (e *Block) moveResizersAndDraggersX() {
	// todo: este bloco vai para setCoordinate e setSize
	e.selectDiv.X(e.x.GetInt())
	e.selectDiv.Width(e.width.GetInt())

	e.resizerTopLeft.SetCX(e.x - e.resizerTopLeft.GetSpace())
	e.resizerTopRight.SetCX(e.x + e.width + e.resizerTopRight.GetSpace())
	e.resizerBottomLeft.SetCX(e.x - e.resizerBottomLeft.GetSpace())
	e.resizerBottomRight.SetCX(e.x + e.width + e.resizerBottomRight.GetSpace())

	e.resizerTopMiddle.SetCX(e.x + e.width/2)
	e.resizerBottomMiddle.SetCX(e.x + e.width/2)
	e.resizerLeftMiddle.SetCX(e.x - e.resizerLeftMiddle.GetSpace())
	e.resizerRightMiddle.SetCX(e.x + e.width + e.resizerRightMiddle.GetSpace())

	e.draggerTopMiddle.SetCX(e.x + e.width/2)
	e.draggerBottomMiddle.SetCX(e.x + e.width/2)
	e.draggerLeftMiddle.SetCX(e.x - e.draggerLeftMiddle.GetSpace())
	e.draggerRightMiddle.SetCX(e.x + e.width + e.draggerRightMiddle.GetSpace())
}

// English:
//
// Português:
func (e *Block) moveResizersAndDraggersY() {
	// todo: este bloco vai para setCoordinate e setSize
	e.selectDiv.Y(e.y.GetInt())
	e.selectDiv.Height(e.height.GetInt())

	e.resizerTopLeft.SetCY(e.y - e.resizerTopLeft.GetSpace())
	e.resizerTopRight.SetCY(e.y - e.resizerTopRight.GetSpace())
	e.resizerBottomLeft.SetCY(e.y + e.height + e.resizerBottomLeft.GetSpace())
	e.resizerBottomRight.SetCY(e.y + e.height + e.resizerBottomRight.GetSpace())

	e.resizerTopMiddle.SetCY(e.y - e.resizerTopMiddle.GetSpace())
	e.resizerBottomMiddle.SetCY(e.y + e.height + e.resizerBottomMiddle.GetSpace())
	e.resizerLeftMiddle.SetCY(e.y + e.height/2)
	e.resizerRightMiddle.SetCY(e.y + e.height/2)

	e.draggerTopMiddle.SetCY(e.y - e.draggerTopMiddle.GetSpace())
	e.draggerBottomMiddle.SetCY(e.y + e.height + e.draggerBottomMiddle.GetSpace())
	e.draggerLeftMiddle.SetCY(e.y + e.height/2)
	e.draggerRightMiddle.SetCY(e.y + e.height/2)
}

// initEvents initialize mouse events

// English:
//
// Português:
func (e *Block) initEvents() {
	var isDragging, isResizing bool
	var startX, startY, startWidth, startHeight, startLeft, startTop int

	// add / remove event listener requires pointers, so the variable should be initialized in this way
	var drag, stopDrag, resizeMouseMove, stopResize js.Func

	// Joins the calculations of X and Y of the drag
	drag = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if !isDragging {
			return nil
		}

		e.block.AddStyle("cursor", "grabbing") // todo: ruleBook

		dx, dy := e.block.GetPointerPosition(args, e.main)

		dx -= startX
		dy -= startY

		newTop := e.min(e.max(0, startTop+dy), e.ideStage.GetClientHeight()-e.block.GetOffsetHeight())
		newLeft := e.min(e.max(0, startLeft+dx), e.ideStage.GetClientWidth()-e.block.GetOffsetWidth())

		dx, newTop = e.adjustXYToGrid(dx, newTop)
		newLeft, dy = e.adjustXYToGrid(newLeft, dy)

		// Get the coordinate (x,y) before the dragging effect to calculate (dx,dy)
		preLeft := e.x
		preTop := e.y

		e.SetPosition(rulesDensity.Convert(newLeft), rulesDensity.Convert(newTop))

		// Calculate (dx,dy) before dragging effect
		e.dragDeltaLeft = e.x - preLeft
		e.dragDeltaTop = e.y - preTop

		e.ruleBook["setOnDraggingEvent"]()
		return nil
	})

	var pFunc func()
	// Removes events when the drag ends
	stopDrag = js.FuncOf(func(this js.Value, args []js.Value) interface{} { // feito
		pFunc()
		return nil
	})
	pFunc = func() {
		isDragging = false
		e.block.AddStyle("cursor", "grab") // todo: ruleBook

		js.Global().Call("removeEventListener", "mousemove", drag)
		js.Global().Call("removeEventListener", "touchmove", drag, false)

		js.Global().Call("removeEventListener", "mouseup", stopDrag)
		js.Global().Call("removeEventListener", "touchend", stopDrag, false)
		js.Global().Call("removeEventListener", "touchcancel", stopDrag, false)
	}

	// Adds the device drag event when the mouse pointer is pressed
	dragFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if !e.dragEnabled {
			return nil
		}

		startX, startY = e.block.GetPointerPosition(args, e.main)

		isDragging = true
		startLeft = e.x.GetInt()
		startTop = e.y.GetInt()

		// The movement of the mouse must be captured from the document and not the dragged element, or when the mouse moves
		// very fast, the drag to
		js.Global().Call("addEventListener", "mousemove", drag, map[string]any{"passive": true})
		js.Global().Call("addEventListener", "touchmove", drag, map[string]any{"passive": true})

		js.Global().Call("addEventListener", "mouseup", stopDrag, map[string]any{"passive": true})
		js.Global().Call("addEventListener", "touchend", stopDrag, map[string]any{"passive": true})
		js.Global().Call("addEventListener", "touchcancel", stopDrag, map[string]any{"passive": true})

		return nil
	})
	e.block.Get().Call("addEventListener", "mousedown", dragFunc, map[string]any{"passive": true})
	e.block.Get().Call("addEventListener", "touchstart", dragFunc, map[string]any{"passive": true})

	// When the resizing tool is activated, four rectangles are designed in the corners of the device.
	// These rectangles are called top-right e top-left, bottom-right, bottom-left.
	//
	// [tl]           [tr]
	//    +-----------+
	//    |           |
	//    |  device   |
	//    |           |
	//    +-----------+
	// [bl]           [br]

	resizeHorizontal := func(args []js.Value, name string) (newLeft, newWidth int) {
		dx, dy := e.block.GetPointerPosition(args, e.main)

		dx -= startX
		dy -= startY

		newLeft = startLeft
		newWidth = startWidth

		if name == "bottom-right" {
			newWidth = e.min(startWidth+dx, e.ideStage.GetClientWidth()-startLeft)
		} else if name == "bottom-left" {
			newWidth = e.min(startWidth-dx, startLeft+startWidth)
			newLeft = e.max(0, startLeft+dx)
		} else if name == "top-right" {
			newWidth = e.min(startWidth+dx, e.ideStage.GetClientWidth()-startLeft)
		} else if name == "top-left" {
			newWidth = e.min(startWidth-dx, startLeft+startWidth)
			newLeft = e.max(0, startLeft+dx)
		} else if name == "top-middle" {
			return
		} else if name == "bottom-middle" {
			return
		} else if name == "left-middle" {
			newWidth = e.min(startWidth-dx, startLeft+startWidth)
			newLeft = e.max(0, startLeft+dx)
		} else if name == "right-middle" {
			newWidth = e.min(startWidth+dx, e.ideStage.GetClientWidth()-startLeft)
		}

		// [tl]           [tr]
		//    +-----------+
		//    |           |
		//    |  device   |
		//    |           |
		//    +-----------+
		// [bl]           [br]
		//
		// Prevents the effect:
		//   When drag TR or BR left, and the size is below minimum, the block is dragged left.
		if newWidth < e.blockMinimumWidth.GetInt() {
			return
		}

		newWidth = e.max(e.blockMinimumWidth.GetInt(), newWidth)

		if e.resizeLimitLeft != math.MaxFloat32 && newLeft > e.resizeLimitLeft.GetInt() {
			return
		}

		if e.resizeLimitRight != math.MaxFloat32 && newLeft+newWidth < e.resizeLimitRight.GetInt() {
			return
		}

		e.SetX(rulesDensity.Convert(newLeft))
		e.SetWidth(rulesDensity.Convert(newWidth))

		return
	}

	resizeVertical := func(args []js.Value, name string) (newTop, newHeight int) {
		dx, dy := e.block.GetPointerPosition(args, e.main)

		dx -= startX
		dy -= startY

		newTop = startTop
		newHeight = startHeight

		if name == "bottom-right" {
			newHeight = e.min(startHeight+dy, e.ideStage.GetClientHeight()-startTop)
		} else if name == "bottom-left" {
			newHeight = e.min(startHeight+dy, e.ideStage.GetClientHeight()-newTop)
		} else if name == "top-right" {
			newHeight = e.min(startHeight-dy, startTop+startHeight)
			newTop = e.max(0, startTop+dy)
		} else if name == "top-left" {
			newHeight = e.min(startHeight-dy, startTop+startHeight)
			newTop = e.max(0, startTop+dy)
		} else if name == "top-middle" {
			newHeight = e.min(startHeight-dy, startTop+startHeight)
			newTop = e.max(0, startTop+dy)
		} else if name == "bottom-middle" {
			newHeight = e.min(startHeight+dy, e.ideStage.GetClientHeight()-newTop)
		} else if name == "left-middle" {
			return
		} else if name == "right-middle" {
			return
		}

		// [tl]           [tr]
		//    +-----------+
		//    |           |
		//    |  device   |
		//    |           |
		//    +-----------+
		// [bl]           [br]
		//
		// Prevents the effect:
		//   When drag TL or TR down, and the size is below minimum, the block is dragged down.
		if newHeight < e.blockMinimumHeight.GetInt() {
			return
		}

		newHeight = e.max(e.blockMinimumHeight.GetInt(), newHeight)

		if e.resizeLimitTop != math.MaxFloat32 && newTop > e.resizeLimitTop.GetInt() {
			return
		}

		if e.resizeLimitBottom != math.MaxFloat32 && newTop+newHeight < e.resizeLimitBottom.GetInt() {
			return
		}

		e.SetY(rulesDensity.Convert(newTop))
		e.SetHeight(rulesDensity.Convert(newHeight))
		return
	}

	resizeMouseMove = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if !isResizing {
			return nil
		}

		e.ruleBook["setResizingOn"]()

		resizerName := e.block.Get().Get("dataset").Get("resizeName").String()
		resizeHorizontal(args, resizerName)
		resizeVertical(args, resizerName)

		if e.warningMarkAppended {
			_ = e.warningMark.Update(e.x, e.y, e.width, e.height)
		}
		_ = e.ornament.Update(e.x, e.y, e.width, e.height)

		e.OnResize(args, e.width, e.height)

		return nil
	})

	stopResize = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		isResizing = false
		js.Global().Call("removeEventListener", "mousemove", resizeMouseMove)
		js.Global().Call("removeEventListener", "mouseup", stopResize)

		js.Global().Call("removeEventListener", "touchmove", resizeMouseMove, false)
		js.Global().Call("removeEventListener", "touchend", stopResize, false)
		js.Global().Call("removeEventListener", "touchcancel", stopResize, false)
		return nil
	})

	resizeFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if !e.resizeEnabled {
			return nil
		}

		resizerName := this.Get("dataset").Get("name").String()
		e.block.DataKey("resizeName", resizerName)

		isResizing = true

		startX, startY = e.block.GetPointerPosition(args, e.main)

		startWidth = e.width.GetInt()
		startHeight = e.height.GetInt()
		startLeft = e.x.GetInt()
		startTop = e.y.GetInt()

		js.Global().Call("addEventListener", "mousemove", resizeMouseMove, map[string]any{"passive": true})
		js.Global().Call("addEventListener", "mouseup", stopResize, map[string]any{"passive": true})

		js.Global().Call("addEventListener", "touchmove", resizeMouseMove, map[string]any{"passive": true})
		js.Global().Call("addEventListener", "touchend", stopResize, map[string]any{"passive": true})
		js.Global().Call("addEventListener", "touchcancel", stopResize, map[string]any{"passive": true})
		return nil
	})

	e.resizerTopLeft.GetSvg().Get().Call("addEventListener", "mousedown", resizeFunc, map[string]any{"passive": true})
	e.resizerTopRight.GetSvg().Get().Call("addEventListener", "mousedown", resizeFunc, map[string]any{"passive": true})
	e.resizerBottomLeft.GetSvg().Get().Call("addEventListener", "mousedown", resizeFunc, map[string]any{"passive": true})
	e.resizerBottomRight.GetSvg().Get().Call("addEventListener", "mousedown", resizeFunc, map[string]any{"passive": true})

	e.resizerTopMiddle.GetSvg().Get().Call("addEventListener", "mousedown", resizeFunc, map[string]any{"passive": true})
	e.resizerBottomMiddle.GetSvg().Get().Call("addEventListener", "mousedown", resizeFunc, map[string]any{"passive": true})
	e.resizerLeftMiddle.GetSvg().Get().Call("addEventListener", "mousedown", resizeFunc, map[string]any{"passive": true})
	e.resizerRightMiddle.GetSvg().Get().Call("addEventListener", "mousedown", resizeFunc, map[string]any{"passive": true})

	e.resizerTopLeft.GetSvg().Get().Call("addEventListener", "touchstart", resizeFunc, map[string]any{"passive": true})
	e.resizerTopRight.GetSvg().Get().Call("addEventListener", "touchstart", resizeFunc, map[string]any{"passive": true})
	e.resizerBottomLeft.GetSvg().Get().Call("addEventListener", "touchstart", resizeFunc, map[string]any{"passive": true})
	e.resizerBottomRight.GetSvg().Get().Call("addEventListener", "touchstart", resizeFunc, map[string]any{"passive": true})

	e.resizerTopMiddle.GetSvg().Get().Call("addEventListener", "touchstart", resizeFunc, map[string]any{"passive": true})
	e.resizerBottomMiddle.GetSvg().Get().Call("addEventListener", "touchstart", resizeFunc, map[string]any{"passive": true})
	e.resizerLeftMiddle.GetSvg().Get().Call("addEventListener", "touchstart", resizeFunc, map[string]any{"passive": true})
	e.resizerRightMiddle.GetSvg().Get().Call("addEventListener", "touchstart", resizeFunc, map[string]any{"passive": true})
}

// max Returns the maximum value

// English:
//
// Português:
func (e *Block) max(a, b int) (max int) {
	if a > b {
		return a
	}
	return b
}

// max Returns the maximum value

// English:
//
// Português:
func (e *Block) maxD(a, b rulesDensity.Density) (max rulesDensity.Density) {
	if a > b {
		return a
	}
	return b
}

// min Returns the minimum value

// English:
//
// Português:
func (e *Block) min(a, b int) (min int) {
	if a < b {
		return a
	}
	return b
}

// min Returns the minimum value

// English:
//
// Português:
func (e *Block) minD(a, b rulesDensity.Density) (min rulesDensity.Density) {
	if a < b {
		return a
	}
	return b
}

// OnResize cannot be shadowed by the main instance, so the function in SetOnResize

// English:
//
// Português:
func (e *Block) OnResize(args []js.Value, width, height rulesDensity.Density) {
	if e.onResizeFunc != nil { // todo: revisar isto
		e.onResizeFunc(args, width, height)
	}
}

// English:
//
// Português:
func (e *Block) setResizeOrnamentVisibleOn() {
	if !e.initialized {
		return
	}

	e.resizerTopLeft.SetVisible(true, e.ideStage)
	e.resizerTopRight.SetVisible(true, e.ideStage)
	e.resizerBottomLeft.SetVisible(true, e.ideStage)
	e.resizerBottomRight.SetVisible(true, e.ideStage)

	e.resizerTopMiddle.SetVisible(true, e.ideStage)
	e.resizerBottomMiddle.SetVisible(true, e.ideStage)
	e.resizerLeftMiddle.SetVisible(true, e.ideStage)
	e.resizerRightMiddle.SetVisible(true, e.ideStage)
}

// English:
//
// Português:
func (e *Block) setResizeOrnamentVisibleOff() {
	if !e.initialized {
		return
	}

	e.resizerTopLeft.SetVisible(false, e.ideStage)
	e.resizerTopRight.SetVisible(false, e.ideStage)
	e.resizerBottomLeft.SetVisible(false, e.ideStage)
	e.resizerBottomRight.SetVisible(false, e.ideStage)

	e.resizerTopMiddle.SetVisible(false, e.ideStage)
	e.resizerBottomMiddle.SetVisible(false, e.ideStage)
	e.resizerLeftMiddle.SetVisible(false, e.ideStage)
	e.resizerRightMiddle.SetVisible(false, e.ideStage)
}

// SetFatherId Receives the div ID used as a stage for the IDE and puts it to occupy the entire browser area

// English:
//
// Português:
func (e *Block) SetFatherId(fatherId string) {
	e.ideStage = factoryBrowser.NewTagSvg().
		Import(fatherId)
}

// SetID Define the device's div ID

// English:
//
// Português:
func (e *Block) SetID(id string) (err error) {
	e.id, err = utils.VerifyUniqueId(id)
	return
}

// SetMinimumHeight Defines the minimum height of the device

// English:
//
// Português:
func (e *Block) SetMinimumHeight(height rulesDensity.Density) {
	e.blockMinimumHeight = height
}

// SetMinimumWidth Defines the minimum width of the device

// English:
//
// Português:
func (e *Block) SetMinimumWidth(width rulesDensity.Density) {
	e.blockMinimumWidth = width
}

// SetName Defines a unique name for the device [compulsory]

// English:
//
// Português:
func (e *Block) SetName(name string) {
	e.name = rulesSequentialId.GetIdFromBase(name)
	return
}

// SetOnResize Receives the pointer to a function to be invoked during resizing
//
//	This function is because the main instance cannot shadow the OnResize function

// English:
//
// Português:
func (e *Block) SetOnResize(f func(args []js.Value, width, height rulesDensity.Density)) {
	e.onResizeFunc = f
}

// SetOrnament Sets the ornament draw object
//
//	ornament draw object is the instance in charge of making the SVG design of the device

// English:
//
// Português:
func (e *Block) SetOrnament(ornament ornament.Draw) {
	e.ornament = ornament
}

// English:
//
// Português:
func (e *Block) SetX(x rulesDensity.Density) {
	y := e.y
	xInt, _ := e.adjustXYToGrid(x.GetInt(), y.GetInt())
	x = rulesDensity.Convert(xInt)

	e.x = x

	if !e.initialized {
		return
	}

	e.block.X(x.GetInt())

	e.moveResizersAndDraggersX()
	e.moveResizersAndDraggersY()

	if e.warningMarkAppended {
		_ = e.warningMark.Update(e.x, e.y, e.width, e.height)
	}
	_ = e.ornament.Update(e.x, e.y, e.width, e.height)
}

// English:
//
// Português:
func (e *Block) SetY(y rulesDensity.Density) {
	x := e.x
	_, yInt := e.adjustXYToGrid(x.GetInt(), y.GetInt())
	y = rulesDensity.Convert(yInt)

	e.y = y

	if !e.initialized {
		return
	}

	e.block.Y(y.GetInt())

	e.moveResizersAndDraggersX()
	e.moveResizersAndDraggersY()

	if e.warningMarkAppended {
		_ = e.warningMark.Update(e.x, e.y, e.width, e.height)
	}
	_ = e.ornament.Update(e.x, e.y, e.width, e.height)
}

// SetPosition Defines the coordinates (x, y) of the device

// English:
//
// Português:
func (e *Block) SetPosition(x, y rulesDensity.Density) {
	xInt, yInt := e.adjustXYToGrid(x.GetInt(), y.GetInt())
	x, y = rulesDensity.Convert(xInt), rulesDensity.Convert(yInt)

	e.x = x
	e.y = y

	if !e.initialized {
		return
	}

	e.block.X(x.GetInt())
	e.block.Y(y.GetInt())

	e.moveResizersAndDraggersX()
	e.moveResizersAndDraggersY()

	if e.warningMarkAppended {
		_ = e.warningMark.Update(e.x, e.y, e.width, e.height)
	}
	_ = e.ornament.Update(e.x, e.y, e.width, e.height)
}

// English:
//
// Português:
func (e *Block) SetWidth(width rulesDensity.Density) {
	height := e.height

	widthInt, _ := e.adjustXYToGrid(width.GetInt(), height.GetInt())
	width = rulesDensity.Convert(widthInt)

	e.width = width

	if !e.initialized {
		return
	}

	e.block.Width(width.GetInt())

	e.moveResizersAndDraggersX()
	e.moveResizersAndDraggersY()

	if e.warningMarkAppended {
		_ = e.warningMark.Update(e.x, e.y, e.width, e.height)
	}
	_ = e.ornament.Update(e.x, e.y, e.width, e.height)
}

// English:
//
// Português:
func (e *Block) SetHeight(height rulesDensity.Density) {
	width := e.width

	_, heightInt := e.adjustXYToGrid(width.GetInt(), height.GetInt())
	height = rulesDensity.Convert(heightInt)

	e.height = height

	if !e.initialized {
		return
	}

	e.block.Height(height.GetInt())

	e.moveResizersAndDraggersX()
	e.moveResizersAndDraggersY()

	if e.warningMarkAppended {
		_ = e.warningMark.Update(e.x, e.y, e.width, e.height)
	}
	_ = e.ornament.Update(e.x, e.y, e.width, e.height)
}

// SetSize Defines the height and width of the device

// English:
//
// Português:
func (e *Block) SetSize(width, height rulesDensity.Density) {
	e.SetWidth(width)
	e.SetHeight(height)
}

// English:
//
// Português:
func (e *Block) getMenuLabel(condition bool, labelTrue, labelFalse string) (label string) {
	if condition {
		return labelTrue
	}

	return labelFalse
}

func (e *Block) GetMenuDebug() (options []components.MenuOptions) {
	// mover para o topo
	// mover para cima
	// mover para baixo
	// mover para o fim
	options = []components.MenuOptions{
		{
			Label: "Debug",
			Submenu: []components.MenuOptions{
				{
					Label: e.getMenuLabel(e.GetSelected(), "Unselect", "Select"),
					Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} {
						e.SetSelected(!e.GetSelected())
						return nil
					}),
				},
				{
					Label: e.getMenuLabel(e.GetSelectBlocked(), "Select lock disable", "Select lock enable"),
					Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} {
						e.SetSelectBlocked(!e.GetSelectBlocked())
						return nil
					}),
				},
				{
					Label: e.getMenuLabel(e.GetResizeEnable(), "Resize disable", "Resize enable"),
					Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} {
						e.SetResizeEnable(!e.GetResizeEnable())
						return nil
					}),
				},
				{
					Label: e.getMenuLabel(e.GetResizeBlocked(), "Resize lock disable", "Resize lock enable"),
					Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} {
						e.SetResizeBlocked(!e.GetResizeBlocked())
						return nil
					}),
				},
				{
					Label: e.getMenuLabel(e.GetDragEnable(), "Drag disable", "Drag enable"),
					Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} {
						e.SetDragEnable(!e.GetDragEnable())
						return nil
					}),
				},
				{
					Label: e.getMenuLabel(e.GetDragBlocked(), "Drag lock disable", "Drag lock enable"),
					Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} {
						e.SetDragBlocked(!e.GetDragBlocked())
						return nil
					}),
				},
			},
		},
	}

	return
}

// resetLimitForResize
//
// English:
//
//	Reset of the boundary box dimensions for resizing
//
// Português:
//
//	Reset das dimensões da caixa de limite para redimensionamento
func (e *Block) resetLimitForResize() {
	e.resizeLimitLeft = math.MaxFloat32
	e.resizeLimitRight = -math.MaxFloat32
	e.resizeLimitTop = math.MaxFloat32
	e.resizeLimitBottom = -math.MaxFloat32
}

// calculateLimitForResizeOn
//
// English:
//
//	Selects all blocks on top (collision and zIndex greater than) from the main block and calculates the box where
//	they are contained to prevent resizing inside the box.
//
// Português:
//
//	Seleciona todos os blocos em cima (colisão e zIndex maior que) do bloco principal e calcula a caixa onde eles
//	estão contidos para impedir o redimensionamento dentro da caixa.
func (e *Block) calculateLimitForResizeOn() {
	e.resetLimitForResize()

	_, total := managerCollision.Collision.Detect(e)
	zIndex := e.GetZIndex()
	for _, v := range total {
		if zIndex < v.GetZIndex() {
			x := v.GetX()
			y := v.GetY()
			width := v.GetWidth()
			height := v.GetHeight()

			e.resizeLimitLeft = e.minD(e.resizeLimitLeft, x)
			e.resizeLimitRight = e.maxD(e.resizeLimitRight, x+width)
			e.resizeLimitTop = e.minD(e.resizeLimitTop, y)
			e.resizeLimitBottom = e.maxD(e.resizeLimitBottom, y+height)
		}
	}

	e.resizeLimitLeft -= e.resizerMoveBorderLimit
	e.resizeLimitRight += e.resizerMoveBorderLimit
	e.resizeLimitTop -= e.resizerMoveBorderLimit
	e.resizeLimitBottom += e.resizerMoveBorderLimit
}

// selectForDragOn
//
// English:
//
//	Selects all blocks above (collision and zIndex greater than) from the main block and enable the mouse drag tool.
//
// Português:
//
//	Seleciona todos os blocos em cima (colisão e zIndex maior que) do bloco principal e habilita a ferramenta de
//	arrasto do mouse.
func (e *Block) selectForDragOn() {
	_, total := managerCollision.Collision.Detect(e)
	zIndex := e.GetZIndex()
	for _, v := range total {
		if !v.GetDragEnable() && zIndex < v.GetZIndex() {
			v.SetDragEnable(true)
		}
	}
}

// selectForDragOff
//
// English:
//
//	Selects all blocks above (collision and zIndex greater than) from the main block and disable the mouse drag tool.
//
// Português:
//
//	Seleciona todos os blocos em cima (colisão e zIndex maior que) do bloco principal e desabilita a ferramenta de
//	arrasto do mouse.
func (e *Block) selectForDragOff() {
	_, total := managerCollision.Collision.Detect(e)
	zIndex := e.GetZIndex()
	for _, v := range total {
		if v.GetDragEnable() && zIndex < v.GetZIndex() {
			v.SetDragEnable(false)
		}
	}
}

// draggingMoveDraggingSelectedOn
//
// English:
//
//	Get all blocks on top (collision and zIndex greater than) of the main block and move them (dx, dy).
//
// Português:
//
//	Pega todos os blocos em cima (colisão e zIndex maior que) do bloco principal e os move (dx,dy).
func (e *Block) draggingMoveDraggingSelectedOn() {
	_, total := managerCollision.Collision.Detect(e)
	zIndex := e.GetZIndex()
	for _, v := range total {
		if v.GetDragEnable() && zIndex < v.GetZIndex() {
			x := v.GetX()
			y := v.GetY()
			x += e.dragDeltaLeft
			y += e.dragDeltaTop
			v.SetX(x)
			v.SetY(y)
		}
	}
}

// draggingMoveSelectedOnStage
//
// English:
//
//	Gets all blocks marked for dragging and moves them
//
// Português:
//
//	Pega todos os blocos marcados para arrasto e os move
func (e *Block) draggingMoveSelectedOnStage() {
	all := manager.Manager.Get()
	for _, v := range all {
		if v.GetDragEnable() && e.id != v.GetID() {
			x := v.GetX()
			y := v.GetY()
			x += e.dragDeltaLeft
			y += e.dragDeltaTop
			v.SetX(x)
			v.SetY(y)
		}
	}
}

func (e *Block) stageSelectDisableAll() {
	all := manager.Manager.Get()
	for _, v := range all {
		if v.GetSelected() {
			v.SetSelected(false)
		}
	}
}

func (e *Block) stageDragDisableAll() {
	all := manager.Manager.Get()
	for _, v := range all {
		if v.GetDragEnable() {
			v.SetDragEnable(false)
		}
	}
}

func (e *Block) stageResizeDisableAll() {
	all := manager.Manager.Get()
	for _, v := range all {
		if v.GetResizeEnable() {
			v.SetResizeEnable(false)
		}
	}
}
