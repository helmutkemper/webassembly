package components

import (
	"fmt"
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/browser/html"
	"strconv"
	"syscall/js"
	"time"
)

var GlobalMenuList []*menu

type options struct {
	Label     string
	Icon      string
	IconLeft  string
	IconRight string
	Type      string
	Items     []options
	Action    js.Func
	Submenu   []options
}

type MainMenu struct {
	menu
}

type ContextMenu struct {
	menu
}

type menu struct {
	__zIndexList     []*html.TagDiv
	__body           *html.TagDiv
	__header         *html.TagDiv
	__content        *html.TagDiv
	__menu           *html.TagDiv
	__subMenuToClose []*html.TagDiv
	__setup          map[string]string
	__options        []options
	__fixed          bool
	__bodyX          int
	__bodyY          int
	__isDragging     bool
	__offsetX        int
	__offsetY        int
	__buttonDrag     bool
	__buttonMinimize bool
	__buttonClose    bool
	__escapeFunction js.Func
}

func (e *menu) HideButtons(drag, minimize, close bool) {
	e.__buttonDrag = drag
	e.__buttonMinimize = minimize
	e.__buttonClose = close
}

func (e *menu) Menu(options []options) {
	e.__options = options
}

func (e *menu) AttachMenu(element html.Compatible) {
	if e.__fixed {
		return
	}

	element.Get().Call("addEventListener", "contextmenu", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		e.hide()
		args[0].Call("preventDefault")
		args[0].Call("stopPropagation")
		e.show(args[0].Get("clientX").Int(), args[0].Get("clientY").Int())
		return nil
	}))
}

func (e *menu) FixedMenu(x, y int) {
	e.__fixed = true
	e.__bodyX = x
	e.__bodyY = y
}

func (e *menu) Css(key, value string) {
	if e.__setup == nil {
		e.__setup = make(map[string]string)
	}

	e.__setup[key] = value
}

// setupInit
//
// English:
//
//	Configures the menu css
//
// Português:
//
//	Configura o css do menu
func (e *menu) setupInit() {
	if e.__setup == nil {
		e.__setup = make(map[string]string)
	}

	if _, found := e.__setup["border-radius"]; !found {
		e.__setup["border-radius"] = "10px"
	}

	if _, found := e.__setup["bodyShadow"]; !found {
		e.__setup["bodyShadow"] = "0 8px 16px rgba(0, 0, 0, 0.2)"
	}

	if _, found := e.__setup["shadow"]; !found {
		e.__setup["shadow"] = "-1px -1px 10px rgba(0,0,0,0.2)"
	}

	if _, found := e.__setup["border"]; !found {
		e.__setup["border"] = "1px solid #ccc"
	}

	if _, found := e.__setup["backgroundColor"]; !found {
		e.__setup["backgroundColor"] = "#fff"
	}

	if _, found := e.__setup["dividerMargin"]; !found {
		e.__setup["dividerMargin"] = "5px 0"
	}

	if _, found := e.__setup["gridDisplay"]; !found {
		e.__setup["gridDisplay"] = "grid"
	}

	if _, found := e.__setup["gridGridTemplateColumns"]; !found {
		e.__setup["gridGridTemplateColumns"] = "repeat(3, 1fr)"
	}

	if _, found := e.__setup["gridGap"]; !found {
		e.__setup["gridGap"] = "8px"
	}

	if _, found := e.__setup["gridPadding"]; !found {
		e.__setup["gridPadding"] = "2px"
	}

	if _, found := e.__setup["cellTextAlign"]; !found {
		e.__setup["cellTextAlign"] = "center"
	}

	if _, found := e.__setup["cursor"]; !found {
		e.__setup["cursor"] = "pointer"
	}

	if _, found := e.__setup["cellBorder"]; !found {
		e.__setup["cellBorder"] = "0px solid #ccc"
	}

	if _, found := e.__setup["cellBorderRadius"]; !found {
		e.__setup["cellBorderRadius"] = "4px"
	}

	if _, found := e.__setup["cellPadding"]; !found {
		e.__setup["cellPadding"] = "5px"
	}

	if _, found := e.__setup["imgWidth"]; !found {
		e.__setup["imgWidth"] = "32px"
	}

	if _, found := e.__setup["imgHeight"]; !found {
		e.__setup["imgHeight"] = "32px"
	}

	if _, found := e.__setup["imgDisplay"]; !found {
		e.__setup["imgDisplay"] = "block"
	}

	if _, found := e.__setup["imgMargin"]; !found {
		e.__setup["imgMargin"] = "0 auto 5px"
	}

	if _, found := e.__setup["imgPadding"]; !found {
		e.__setup["imgPadding"] = "5px"
	}

	if _, found := e.__setup["textFontSize"]; !found {
		e.__setup["textFontSize"] = "12px"
	}

	if _, found := e.__setup["fontFamily"]; !found {
		e.__setup["fontFamily"] = "Arial, sans-serif"
	}

	if _, found := e.__setup["itemPadding"]; !found {
		e.__setup["itemPadding"] = "5px 10px"
	}

	if _, found := e.__setup["itemCursor"]; !found {
		e.__setup["itemCursor"] = "pointer"
	}

	if _, found := e.__setup["itemPosition"]; !found {
		e.__setup["itemPosition"] = "relative"
	}

	if _, found := e.__setup["itemTextContent"]; !found {
		e.__setup["itemTextContent"] = "&nbsp;&nbsp;▶"
	}

	if _, found := e.__setup["itemDisplay"]; !found {
		e.__setup["itemDisplay"] = "flex"
	}

	if _, found := e.__setup["itemAlignItems"]; !found {
		e.__setup["itemAlignItems"] = "center"
	}

	if _, found := e.__setup["submenuLeft"]; !found {
		e.__setup["submenuLeft"] = "100%"
	}

	if _, found := e.__setup["submenuTop"]; !found {
		e.__setup["submenuTop"] = "0"
	}

	if _, found := e.__setup["submenuBackground"]; !found {
		e.__setup["submenuBackground"] = "#ccc"
	}

	if _, found := e.__setup["submenuBorder"]; !found {
		e.__setup["submenuBorder"] = e.__setup["border"]
	}

	if _, found := e.__setup["submenuBoxShadow"]; !found {
		e.__setup["submenuBoxShadow"] = e.__setup["shadow"]
	}

	if _, found := e.__setup["submenuPadding"]; !found {
		e.__setup["submenuPadding"] = "5px"
	}

	if _, found := e.__setup["menuPadding"]; !found {
		e.__setup["menuPadding"] = "5px"
	}

	if _, found := e.__setup["submenuWhiteSpace"]; !found {
		e.__setup["submenuWhiteSpace"] = "nowrap"
	}

	if _, found := e.__setup["menuTitle"]; !found {
		e.__setup["menuTitle"] = "Menu"
	}

	if _, found := e.__setup["menuMoveIcon"]; !found {
		e.__setup["menuMoveIcon"] = "◇"
	}

	if _, found := e.__setup["menuMoveLabel"]; !found {
		e.__setup["menuMoveLabel"] = "Move"
	}

	if _, found := e.__setup["menuMinimizeIcon"]; !found {
		e.__setup["menuMinimizeIcon"] = "▾"
	}

	if _, found := e.__setup["menuMinimizeLabel"]; !found {
		e.__setup["menuMinimizeLabel"] = "Minimize"
	}

	if _, found := e.__setup["menuCloseIcon"]; !found {
		e.__setup["menuCloseIcon"] = "⊗"
	}

	if _, found := e.__setup["menuCloseLabel"]; !found {
		e.__setup["menuCloseLabel"] = "Close"
	}

	if _, found := e.__setup["headerBackground"]; !found {
		e.__setup["headerBackground"] = "#fff"
	}

	if _, found := e.__setup["headerPadding"]; !found {
		e.__setup["headerPadding"] = "7px 4px"
	}

	if _, found := e.__setup["headerMargin"]; !found {
		e.__setup["headerMargin"] = "-5px -5px 0px -5px"
	}

	if _, found := e.__setup["contentGap"]; !found {
		e.__setup["contentGap"] = "8px"
	}

	if _, found := e.__setup["contentPadding"]; !found {
		e.__setup["contentPadding"] = "2px"
	}
}

// changeZIndex
//
// English:
//
//	Updates the zIndex property of all elements in the menu to ensure proper stacking order.
//
// Português:
//
//	Atualiza a propriedade zIndex de todos os elementos no menu para garantir que sejam empilhados de forma correta
func (e *menu) changeZIndex() {
	nextIndex := e.getNextZIndex()

	for k := range e.__zIndexList {
		e.__zIndexList[k].AddStyle("zIndex", nextIndex+k)
	}
}

// recordsTheMenuGlobally
//
// English:
//
//	Register globally all menus created so that they are closed when another menu contextual is opened
//
// Português:
//
//	Registra globalmente todos os menus criados para que eles sejam fechados quando um outro contextual menu é aberto
func (e *menu) recordsTheMenuGlobally() {
	if GlobalMenuList == nil {
		GlobalMenuList = make([]*menu, 0)
	}
	GlobalMenuList = append(GlobalMenuList, e)
}

// hidesAllRegisteredGloballyMenus
//
// English:
//
//	Closes all open contextual menus
//
// Português:
//
//	Fecha todos os menus contextuais abertos
func (e *menu) hidesAllRegisteredGloballyMenus() {
	for k := range GlobalMenuList {
		if GlobalMenuList[k].__fixed {
			continue
		}

		GlobalMenuList[k].hide()
	}
}

func (e *menu) Init() {
	e.recordsTheMenuGlobally()

	e.__subMenuToClose = make([]*html.TagDiv, 0)
	e.__zIndexList = make([]*html.TagDiv, 0)

	e.__escapeFunction = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if args[0].Get("key").String() == "Escape" {
			e.hide()
		}
		return nil
	})

	e.setupInit()

	e.__body = factoryBrowser.NewTagDiv()
	e.__body.AddStyle("position", "absolute")
	e.__body.AddStyle("border-radius", e.__setup["border-radius"])
	e.__body.AddStyle("background", e.__setup["backgroundColor"])
	e.__body.AddStyle("border", e.__setup["border"])
	e.__body.AddStyle("boxShadow", e.__setup["bodyShadow"])
	e.__body.AddStyle("padding", e.__setup["menuPadding"])
	e.__body.AddStyle("user-select", "none")
	e.__zIndexList = append(e.__zIndexList, e.__body)

	dragIcon := factoryBrowser.NewTagSpan()
	if !e.__buttonDrag {
		dragIcon.AddStyle("cursor", "move").
			Title(e.__setup["menuMoveLabel"]).
			Html(e.__setup["menuMoveIcon"])
		e.headerAddDragListener(dragIcon)
	}

	minimizeIcon := factoryBrowser.NewTagSpan()
	if !e.__buttonMinimize {
		minimizeIcon.AddStyle("cursor", e.__setup["cursor"]).
			Title(e.__setup["menuMinimizeLabel"]).
			Html(e.__setup["menuMinimizeIcon"])
		e.headerAddMinimizeListener(minimizeIcon)
	}

	closeIcon := factoryBrowser.NewTagSpan()
	if !e.__buttonClose {
		closeIcon.AddStyle("cursor", e.__setup["cursor"]).
			Title(e.__setup["menuCloseLabel"]).
			Html(e.__setup["menuCloseIcon"])
		e.headerAddCloseListener(closeIcon)
	}

	e.__header = factoryBrowser.NewTagDiv().Append(
		factoryBrowser.NewTagSpan().Html(e.__setup["menuTitle"]),
	)
	e.__header.AddStyle("display", "flex")
	e.__header.AddStyle("justify-content", "space-between")
	e.__header.AddStyle("align-items", "center")
	e.__header.AddStyle("background", e.__setup["headerBackground"])
	e.__header.AddStyle("padding", e.__setup["headerPadding"])
	e.__header.AddStyle("font-family", e.__setup["fontFamily"])
	e.__header.AddStyle("font-weight", "bold")

	if e.__fixed {
		e.__header.Append(
			factoryBrowser.NewTagSpan().
				AddStyle("display", "flex").
				AddStyle("gap", "8px").
				Append(
					factoryBrowser.NewTagSpan().Html("&nbsp;"),
					dragIcon,
					minimizeIcon,
					closeIcon,
				),
		)
	}

	e.__content = factoryBrowser.NewTagDiv()
	e.__content.AddStyle("display", "grid")
	e.__content.AddStyle("gap", e.__setup["contentGap"])
	e.__content.AddStyle("padding", e.__setup["contentPadding"])
	e.__content.FadeFunc(e.contentFadeProgress)

	e.__body.Append(
		e.__header,
		e.__content,
	)

	e.__menu = factoryBrowser.NewTagDiv()
	e.__content.Append(e.__menu)

	e.__body.HideForFade()
	e.__body.FadeFunc(e.bodyFadeProgress)

	stage := factoryBrowser.NewStage()
	stage.Append(e.__body)
	e.__body.Fade(300 * time.Millisecond)

	if !e.__fixed {
		e.hide()
		js.Global().Get("document").Call("addEventListener", "click", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			e.hide()
			return nil
		}))
	} else {
		e.__body.AddStyle("left", fmt.Sprintf("%vpx", e.__bodyX))
		e.__body.AddStyle("top", fmt.Sprintf("%vpx", e.__bodyY))
	}

	e.ReInit()
}

// ReInit
//
// English:
//
//	Drawing the menu with new data
//
// Português:
//
//	Remonta o menu com novos dados
func (e *menu) ReInit() {
	e.__menu.Html("")
	e.mountMenu(e.__options, e.__menu)
	e.adjustContentWidth()
}

// bodyFadeProgress
//
// English:
//
//	Ajusta o conteúdo do menu quando a função div.Fade() funciona durante a função de fechar
//
// Português:
//
//	Adjusts the menu content when the Div.Fade() function works during the close function
func (e *menu) bodyFadeProgress(progress float64) {
	e.adjustContentWidth()

	if progress == 1.0 && !e.__content.FadeStatus() {
		e.__content.ShowForFade()
	}
}

// contentFadeProgress
//
// English:
//
//	Ajusta o conteúdo do menu quando a função div.Fade() funciona durante a função de minimizar
//
// Português:
//
//	Adjusts the menu content when the Div.Fade() function works during the minimize function
func (e *menu) contentFadeProgress(progress float64) {
	e.adjustContentWidth()

	if progress == 1.0 {
		if e.__content.FadeStatus() {
			e.fadeShowContent()
			return
		}

		e.fadeHideContent()
	}
}

// fadeShowContent
//
// English:
//
//	Show and adjusts the menu content when the div.Fade() function ends
//
// Português:
//
//	Mostra e ajusta o conteúdo do menu quando a função div.Fade() termina
func (e *menu) fadeShowContent() {
	e.__content.AddStyle("visibility", "visible")
	e.__menu.AddStyle("visibility", "visible")

	e.__body.AddStyle("padding", e.__setup["menuPadding"])
	e.__content.AddStyle("padding", e.__setup["contentPadding"])
}

// closeAllSubMenus
//
// English:
//
//	Closes all submenus after the click, as the fixed menu does not have the hide() function
//
// Português:
//
//	Fecha todos os submenus após o click, pois o menu fixo não tem a função hide()
func (e *menu) closeAllSubMenus() {
	for k := range e.__subMenuToClose {
		e.__subMenuToClose[k].AddStyle("display", "none")
	}
}

// fadeHideContent
//
// English:
//
//	Hide and adjusts the menu content when the div.Fade() function ends
//
// Português:
//
//	Esconde e ajusta o conteúdo do menu quando a função div.Fade() termina
func (e *menu) fadeHideContent() {
	e.__content.AddStyle("visibility", "hidden")
	e.__menu.AddStyle("visibility", "hidden")

	e.__body.AddStyle("padding", "0")
	e.__content.AddStyle("padding", "0")
}

// headerAddDragListener
//
// English:
//
//	Add a listener to the moving button
//
// Português:
//
//	Adiciona o listener para o botão de mover
func (e *menu) headerAddDragListener(dragIcon *html.TagSpan) {
	dragIcon.Get().Call("addEventListener", "mousedown", js.FuncOf(func(this js.Value, args []js.Value) any {
		e.__isDragging = true
		e.__offsetX = args[0].Get("clientX").Int() - e.__body.Get().Call("getBoundingClientRect").Get("left").Int()
		e.__offsetY = args[0].Get("clientY").Int() - e.__body.Get().Call("getBoundingClientRect").Get("top").Int()
		return nil
	}))

	js.Global().Get("document").Call("addEventListener", "mousemove", js.FuncOf(func(this js.Value, args []js.Value) any {
		if e.__isDragging {
			e.__body.AddStyle("left", fmt.Sprintf("%vpx", args[0].Get("clientX").Int()-e.__offsetX))
			e.__body.AddStyle("top", fmt.Sprintf("%vpx", args[0].Get("clientY").Int()-e.__offsetY))
		}
		return nil
	}))

	js.Global().Get("document").Call("addEventListener", "mouseup", js.FuncOf(func(this js.Value, args []js.Value) any {
		e.__isDragging = false
		return nil
	}))
}

// headerAddMinimizeListener
//
// English:
//
//	Add a listener to the minimizing button
//
// Português:
//
//	Adiciona o listener para o botão de minimizar
func (e *menu) headerAddMinimizeListener(closeIcon *html.TagSpan) {
	closeIcon.Get().Call("addEventListener", "click", js.FuncOf(func(this js.Value, args []js.Value) any {
		args[0].Call("stopPropagation")
		e.__content.Fade(300 * time.Millisecond)

		return nil
	}))
}

// headerAddCloseListener
//
// English:
//
//	Add a listener to the closing button
//
// Português:
//
//	Adiciona o listener para o botão de fechar
func (e *menu) headerAddCloseListener(closeIcon *html.TagSpan) {
	closeIcon.Get().Call("addEventListener", "click", js.FuncOf(func(this js.Value, args []js.Value) any {
		args[0].Call("stopPropagation")

		e.__body.Fade(200 * time.Millisecond)
		go func() {
			time.Sleep(1 * time.Second)
			e.__body.Fade(200 * time.Millisecond)
		}()

		return nil
	}))
}

// adjustContentWidth
//
// English:
//
//	Adjusts the menu length when it opens by div.Fade()
//
// Português:
//
//	Ajusta o comprimento do menu quando ele abre por div.Fade()
func (e *menu) adjustContentWidth() {
	menuRect := e.__menu.Get().Call("getBoundingClientRect")
	width := menuRect.Get("width").Int()
	height := menuRect.Get("height").Int()
	e.__content.AddStyle("width", fmt.Sprintf("%vpx", width))
	e.__content.AddStyle("height", fmt.Sprintf("%vpx", height))
}

// mountMenu
//
// English:
//
//	Mounts the menu and the submenu
//
// Português:
//
//	Monta o menu e os submenu
func (e *menu) mountMenu(options []options, container *html.TagDiv) {
	for _, option := range options {
		if option.Label == "-" {
			divider := factoryBrowser.NewTagHr()
			divider.AddStyle("margin", e.__setup["dividerMargin"])
			container.Append(divider)
			continue

		} else if option.Type == "grid" && option.Items != nil && len(option.Items) > 0 {
			grid := factoryBrowser.NewTagDiv()
			grid.AddStyle("display", e.__setup["gridDisplay"])
			grid.AddStyle("gridTemplateColumns", e.__setup["gridGridTemplateColumns"])
			grid.AddStyle("gap", e.__setup["gridGap"])
			grid.AddStyle("padding", e.__setup["gridPadding"])

			for _, item := range option.Items {
				cell := factoryBrowser.NewTagDiv()
				cell.AddStyle("textAlign", e.__setup["cellTextAlign"])
				cell.AddStyle("cursor", e.__setup["cursor"])
				cell.AddStyle("border", e.__setup["cellBorder"])
				cell.AddStyle("borderRadius", e.__setup["cellBorderRadius"])
				cell.AddStyle("padding", e.__setup["cellPadding"])
				cell.AddStyle("position", "relative")

				img := factoryBrowser.NewTagImg()
				img.Src(item.Icon, false)
				img.Alt(item.Label)
				img.AddStyle("width", e.__setup["imgWidth"])
				img.AddStyle("height", e.__setup["imgHeight"])
				img.AddStyle("display", e.__setup["imgDisplay"])
				img.AddStyle("margin", e.__setup["imgMargin"])
				img.AddStyle("padding", e.__setup["imgPadding"])

				text := factoryBrowser.NewTagDiv()
				text.AddStyle("fontSize", e.__setup["textFontSize"])
				text.AddStyle("fontFamily", e.__setup["fontFamily"])
				text.AddStyle("white-space", "nowrap")
				if item.Submenu != nil && len(item.Submenu) > 0 {
					text.Html(fmt.Sprintf("<span style=\"flex:1; text-align:left;\">%v</span><span style=\"text-align:right;\">%v</span>", item.Label, e.__setup["itemTextContent"]))
					text.AddStyle("display", e.__setup["itemDisplay"])
					text.AddStyle("alignItems", e.__setup["itemAlignItems"])
				} else {
					text.Text(item.Label)
				}

				cell.Append(img)
				cell.Append(text)

				if item.Submenu != nil && len(item.Submenu) > 0 {
					subMenu := factoryBrowser.NewTagDiv()
					subMenu.AddStyle("position", "absolute")
					subMenu.AddStyle("border-radius", e.__setup["border-radius"])
					subMenu.AddStyle("left", e.__setup["submenuLeft"])
					subMenu.AddStyle("top", e.__setup["submenuTop"])
					subMenu.AddStyle("background", e.__setup["backgroundColor"])
					subMenu.AddStyle("border", e.__setup["submenuBorder"])
					subMenu.AddStyle("boxShadow", e.__setup["submenuBoxShadow"])
					subMenu.AddStyle("padding", e.__setup["submenuPadding"])
					subMenu.AddStyle("display", "none")
					subMenu.AddStyle("whiteSpace", e.__setup["submenuWhiteSpace"])
					subMenu.AddStyle("white-space", "nowrap")
					e.__zIndexList = append(e.__zIndexList, subMenu)

					e.__subMenuToClose = append(e.__subMenuToClose, subMenu)

					e.mountMenu(item.Submenu, subMenu)
					cell.Append(subMenu)

					cell.Get().Call("addEventListener", "mouseenter", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
						cell.AddStyle("background", e.__setup["submenuBackground"])
						e.adjustSubMenuPosition(subMenu, cell)
						return nil
					}))

					cell.Get().Call("addEventListener", "mouseleave", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
						cell.AddStyle("background", "transparent")
						cell.AddStyle("border", e.__setup["cellBorder"])
						subMenu.AddStyle("display", "none")

						return nil
					}))
				} else {
					cell.Get().Call("addEventListener", "click", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
						args[0].Call("stopPropagation")

						cell.AddStyle("background", "transparent")
						cell.AddStyle("border", e.__setup["cellBorder"])
						if !item.Action.IsUndefined() {
							js.ValueOf(item.Action).Invoke()
							e.closeAllSubMenus()
						}
						if !e.__fixed {
							e.hide()
						}
						return nil
					}))

					cell.Get().Call("addEventListener", "mouseenter", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
						cell.AddStyle("background", e.__setup["submenuBackground"])
						cell.AddStyle("border", e.__setup["cellBorder"])
						return nil
					}))

					cell.Get().Call("addEventListener", "mouseleave", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
						cell.AddStyle("background", "transparent")
						cell.AddStyle("border", e.__setup["cellBorder"])
						return nil
					}))
				}

				grid.Append(cell)
			}

			container.Append(grid)
			continue
		}

		item := factoryBrowser.NewTagDiv()
		item.Text(option.Label)
		item.AddStyle("textAlign", "left")
		item.AddStyle("fontSize", e.__setup["textFontSize"])
		item.AddStyle("fontFamily", e.__setup["fontFamily"])
		item.AddStyle("padding", e.__setup["itemPadding"])
		item.AddStyle("cursor", e.__setup["itemCursor"])
		item.AddStyle("position", e.__setup["itemPosition"])
		item.AddStyle("white-space", "nowrap")

		item.Get().Call("addEventListener", "mouseenter", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			item.AddStyle("background", e.__setup["submenuBackground"])
			return nil
		}))
		item.Get().Call("addEventListener", "mouseleave", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			item.AddStyle("background", "transparent")
			return nil
		}))

		// submenu em linha
		if option.Submenu != nil && len(option.Submenu) > 0 {
			item.Html(fmt.Sprintf("<span style=\"flex:1; text-align:left;\">%v</span><span style=\"text-align:right;\">%v</span>", option.Label, e.__setup["itemTextContent"]))
			item.AddStyle("display", e.__setup["itemDisplay"])
			item.AddStyle("alignItems", e.__setup["itemAlignItems"])

			subMenu := factoryBrowser.NewTagDiv()
			subMenu.AddStyle("position", "absolute")
			subMenu.AddStyle("border-radius", e.__setup["border-radius"])
			subMenu.AddStyle("left", e.__setup["submenuLeft"])
			subMenu.AddStyle("top", e.__setup["submenuTop"])
			subMenu.AddStyle("background", e.__setup["backgroundColor"])
			subMenu.AddStyle("border", e.__setup["submenuBorder"])
			subMenu.AddStyle("boxShadow", e.__setup["submenuBoxShadow"])
			subMenu.AddStyle("padding", e.__setup["submenuPadding"])
			subMenu.AddStyle("display", "none")
			subMenu.AddStyle("whiteSpace", e.__setup["submenuWhiteSpace"])
			subMenu.AddStyle("white-space", "nowrap")
			//subMenu.AddStyle("zIndex", e.__setup["submenuZIndex"])
			e.__zIndexList = append(e.__zIndexList, subMenu)

			e.__subMenuToClose = append(e.__subMenuToClose, subMenu)

			e.mountMenu(option.Submenu, subMenu)

			item.Append(subMenu)
			item.Get().Call("addEventListener", "mouseenter", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
				e.adjustSubMenuPosition(subMenu, item)
				return nil
			}))
			item.Get().Call("addEventListener", "mouseleave", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
				subMenu.AddStyle("display", "none")
				return nil
			}))
		} else {
			item.Text(option.Label)
			item.Get().Call("addEventListener", "click", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
				args[0].Call("stopPropagation")
				if !option.Action.IsUndefined() {
					js.ValueOf(option.Action).Invoke()
					e.closeAllSubMenus()
				}
				if !e.__fixed {
					e.hide()
				}
				return nil
			}))
		}

		container.Append(item)
	}
}

// adjustSubMenuPosition
//
// English:
//
//	Adjusts the top and left position of the submenu that opens so that it remains visible on the screen
//
// Português:
//
//	Ajusta a posição top e left do submenu que abre para que o mesmo permaneça visível na tela
func (e *menu) adjustSubMenuPosition(subMenu, cell *html.TagDiv) {
	subMenu.AddStyle("display", "block")

	window := js.Global().Get("window")
	screenWidth := window.Get("innerWidth").Int()
	screenHeight := window.Get("innerHeight").Int()

	submenuRect := subMenu.Get().Call("getBoundingClientRect")

	cellRect := cell.Get().Call("getBoundingClientRect")
	cellRight := cellRect.Get("right").Int()
	cellTop := cellRect.Get("top").Int()

	if cellRight+submenuRect.Get("width").Int() > screenWidth {
		subMenu.AddStyle("left", "auto")
		subMenu.AddStyle("right", "100%")
	} else {
		subMenu.AddStyle("left", "100%")
		subMenu.AddStyle("right", "auto")
	}

	if cellTop+submenuRect.Get("height").Int() > screenHeight {
		subMenu.AddStyle("top", "auto")
		subMenu.AddStyle("bottom", "0")
	} else {
		subMenu.AddStyle("top", "0")
		subMenu.AddStyle("bottom", "auto")
	}
}

// getNextZIndex
//
// English:
//
//	Looking for all graphic elements in the document and then calculates the next Zindex
//
// Português:
//
//	Procura todos os elementos gráficos no documento e em seguida calcula o próximo zIndex
func (e *menu) getNextZIndex() int {

	maxZIndex := 0
	elements := js.Global().Get("document").Call("getElementsByTagName", "*")
	length := elements.Length()

	for i := 0; i < length; i++ {
		element := elements.Index(i)
		style := js.Global().Get("window").Call("getComputedStyle", element)
		zIndex := style.Get("zIndex").String()
		if zIndex != "auto" {
			if parsedZIndex, err := strconv.Atoi(zIndex); err == nil {
				if parsedZIndex > maxZIndex {
					maxZIndex = parsedZIndex
				}
			}
		}
	}

	return maxZIndex + 1
}

// show
//
// English:
//
//	Show the menu when the contextual menu is triggered by the mouse
//
// Português:
//
//	Mostra o menu quando o menu contextual é acionado pelo mouse
func (e *menu) show(x, y int) {

	e.hidesAllRegisteredGloballyMenus()

	e.__body.AddStyle("display", "block")
	e.__body.AddStyle("left", "0px")
	e.__body.AddStyle("top", "0px")

	e.adjustContentWidth()

	bbox := e.__body.Get().Call("getBoundingClientRect")
	menuWidth := bbox.Get("width").Int()
	menuHeight := bbox.Get("height").Int()
	screenWidth := js.Global().Get("window").Get("innerWidth").Int()
	screenHeight := js.Global().Get("window").Get("innerHeight").Int()

	adjustedX := x
	adjustedY := y

	if x+menuWidth > screenWidth {
		adjustedX = screenWidth - menuWidth - 5
	}
	if y+menuHeight > screenHeight {
		adjustedY = screenHeight - menuHeight - 5
	}

	adjustedX = e.max(adjustedX, 0)
	adjustedY = e.max(adjustedY, 0)

	e.__body.AddStyle("left", strconv.FormatInt(int64(adjustedX), 10)+"px")
	e.__body.AddStyle("top", strconv.FormatInt(int64(adjustedY), 10)+"px")

	js.Global().Get("document").Call("addEventListener", "keydown", e.__escapeFunction)

	e.changeZIndex()
}

// hide
//
// Esconde o menu quando este é configurado para ser um menu contextual
func (e *menu) hide() {
	e.__body.AddStyle("display", "none")
	js.Global().Get("document").Call("removeEventListener", "keydown", e.__escapeFunction)
}

func (e *menu) max(x, y int) (max int) {
	if x > y {
		return x
	}

	return y
}
