package main

import (
	"fmt"
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/browser/stage"
	"strconv"
	"syscall/js"
	"time"
)

type Options struct {
	Label   string
	Icon    string
	Type    string
	Items   []Options
	Action  js.Func
	Submenu []Options
}

type ContextMenu struct {
	body    *html.TagDiv
	header  *html.TagDiv
	content *html.TagDiv
	menu    *html.TagDiv
	setup   map[string]string
	stage   *stage.Stage
	fixed   bool
	bodyX   int
	bodyY   int

	isDragging bool
	offsetX    int
	offsetY    int
}

func (e *ContextMenu) Menu(options []Options) {
	e.menu.Html("")
	e.mountMenu(options, e.menu)
	e.adjustContentWidth()
}

func (e *ContextMenu) AttachMenu(element js.Value) {
	if e.fixed {
		return
	}

	element.Call("addEventListener", "contextmenu", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		args[0].Call("preventDefault")
		e.show(args[0].Get("clientX").Int(), args[0].Get("clientY").Int())
		return nil
	}))
}

func (e *ContextMenu) FixedMenu(x, y int) {
	e.fixed = true
	e.bodyX = x
	e.bodyY = y
}

func (e *ContextMenu) Stage(stage *stage.Stage) {
	e.stage = stage
}

func (e *ContextMenu) Css(key, value string) {
	if e.setup == nil {
		e.setup = make(map[string]string)
	}

	e.setup[key] = value
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
func (e *ContextMenu) setupInit() {
	if e.setup == nil {
		e.setup = make(map[string]string)
	}

	if _, found := e.setup["shadow"]; !found {
		e.setup["shadow"] = "-1px -1px 10px rgba(0,0,0,0.2)"
	}

	if _, found := e.setup["border"]; !found {
		e.setup["border"] = "0px solid #ccc"
	}

	if _, found := e.setup["backgroundColor"]; !found {
		e.setup["backgroundColor"] = "#fff"
	}

	if _, found := e.setup["dividerMargin"]; !found {
		e.setup["dividerMargin"] = "5px 0"
	}

	if _, found := e.setup["gridDisplay"]; !found {
		e.setup["gridDisplay"] = "grid"
	}

	if _, found := e.setup["gridGridTemplateColumns"]; !found {
		e.setup["gridGridTemplateColumns"] = "repeat(3, 1fr)"
	}

	if _, found := e.setup["gridGap"]; !found {
		e.setup["gridGap"] = "8px"
	}

	if _, found := e.setup["gridPadding"]; !found {
		e.setup["gridPadding"] = "2px"
	}

	if _, found := e.setup["cellTextAlign"]; !found {
		e.setup["cellTextAlign"] = "center"
	}

	if _, found := e.setup["cursor"]; !found {
		e.setup["cursor"] = "pointer"
	}

	if _, found := e.setup["cellBorder"]; !found {
		e.setup["cellBorder"] = e.setup["border"]
	}

	if _, found := e.setup["cellBorderRadius"]; !found {
		e.setup["cellBorderRadius"] = "4px"
	}

	if _, found := e.setup["cellPadding"]; !found {
		e.setup["cellPadding"] = "5px"
	}

	if _, found := e.setup["imgWidth"]; !found {
		e.setup["imgWidth"] = "32px"
	}

	if _, found := e.setup["imgHeight"]; !found {
		e.setup["imgHeight"] = "32px"
	}

	if _, found := e.setup["imgDisplay"]; !found {
		e.setup["imgDisplay"] = "block"
	}

	if _, found := e.setup["imgMargin"]; !found {
		e.setup["imgMargin"] = "0 auto 5px"
	}

	if _, found := e.setup["imgPadding"]; !found {
		e.setup["imgPadding"] = "5px"
	}

	if _, found := e.setup["textFontSize"]; !found {
		e.setup["textFontSize"] = "12px"
	}

	if _, found := e.setup["fontFamily"]; !found {
		e.setup["fontFamily"] = "Arial, sans-serif"
	}

	if _, found := e.setup["itemPadding"]; !found {
		e.setup["itemPadding"] = "5px 10px"
	}

	if _, found := e.setup["itemCursor"]; !found {
		e.setup["itemCursor"] = "pointer"
	}

	if _, found := e.setup["itemPosition"]; !found {
		e.setup["itemPosition"] = "relative"
	}

	if _, found := e.setup["itemTextContent"]; !found {
		e.setup["itemTextContent"] = "&nbsp;&nbsp;▶"
	}

	if _, found := e.setup["itemDisplay"]; !found {
		e.setup["itemDisplay"] = "flex"
	}

	if _, found := e.setup["itemAlignItems"]; !found {
		e.setup["itemAlignItems"] = "center"
	}

	if _, found := e.setup["submenuPosition"]; !found {
		e.setup["submenuPosition"] = "absolute"
	}

	if _, found := e.setup["submenuLeft"]; !found {
		e.setup["submenuLeft"] = "100%"
	}

	if _, found := e.setup["submenuTop"]; !found {
		e.setup["submenuTop"] = "0"
	}

	if _, found := e.setup["submenuBackground"]; !found {
		e.setup["submenuBackground"] = "#ccc"
	}

	if _, found := e.setup["submenuBorder"]; !found {
		e.setup["submenuBorder"] = e.setup["border"]
	}

	if _, found := e.setup["submenuBoxShadow"]; !found {
		e.setup["submenuBoxShadow"] = e.setup["shadow"]
	}

	if _, found := e.setup["submenuPadding"]; !found {
		e.setup["submenuPadding"] = "5px"
	}

	if _, found := e.setup["menuPadding"]; !found {
		e.setup["menuPadding"] = "5px"
	}

	if _, found := e.setup["submenuDisplay"]; !found {
		e.setup["submenuDisplay"] = "none"
	}

	if _, found := e.setup["submenuWhiteSpace"]; !found {
		e.setup["submenuWhiteSpace"] = "nowrap"
	}

	if _, found := e.setup["submenuZIndex"]; !found {
		e.setup["submenuZIndex"] = "1001"
	}

	if _, found := e.setup["menuZIndex"]; !found {
		e.setup["menuZIndex"] = "1000"
	}

	if _, found := e.setup["menuTitle"]; !found {
		e.setup["menuTitle"] = "Menu"
	}

	if _, found := e.setup["menuMoveIcon"]; !found {
		e.setup["menuMoveIcon"] = "◇"
	}

	if _, found := e.setup["menuMoveLabel"]; !found {
		e.setup["menuMoveLabel"] = "Move"
	}

	if _, found := e.setup["menuMinimizeIcon"]; !found {
		e.setup["menuMinimizeIcon"] = "▾"
	}

	if _, found := e.setup["menuMinimizeLabel"]; !found {
		e.setup["menuMinimizeLabel"] = "Minimize"
	}

	if _, found := e.setup["menuCloseIcon"]; !found {
		e.setup["menuCloseIcon"] = "⊗"
	}

	if _, found := e.setup["menuCloseLabel"]; !found {
		e.setup["menuCloseLabel"] = "Close"
	}

	if _, found := e.setup["headerBackground"]; !found {
		e.setup["headerBackground"] = "#e0e0e0"
	}

	if _, found := e.setup["headerPadding"]; !found {
		e.setup["headerPadding"] = "4px 8px"
	}

	if _, found := e.setup["headerMargin"]; !found {
		e.setup["headerMargin"] = "-5px -5px 0px -5px"
	}

	if _, found := e.setup["contentGap"]; !found {
		e.setup["contentGap"] = "8px"
	}

	if _, found := e.setup["contentPadding"]; !found {
		e.setup["contentPadding"] = "2px"
	}
}

func (e *ContextMenu) Init() {
	e.setupInit()

	e.body = factoryBrowser.NewTagDiv()
	e.body.AddStyle("position", "absolute")
	e.body.AddStyle("background", e.setup["backgroundColor"])
	e.body.AddStyle("border", e.setup["border"])
	e.body.AddStyle("boxShadow", e.setup["shadow"])
	e.body.AddStyle("padding", e.setup["menuPadding"])
	e.body.AddStyle("zIndex", e.setup["menuZIndex"])

	dragIcon := factoryBrowser.NewTagSpan().
		AddStyle("cursor", "move").
		Title(e.setup["menuMoveLabel"]).
		Html(e.setup["menuMoveIcon"])
	e.headerAddDragListener(dragIcon)

	minimizeIcon := factoryBrowser.NewTagSpan().
		AddStyle("cursor", e.setup["cursor"]).
		Title(e.setup["menuMinimizeLabel"]).
		Html(e.setup["menuMinimizeIcon"])
	e.headerAddMinimizeListener(minimizeIcon)

	closeIcon := factoryBrowser.NewTagSpan().
		AddStyle("cursor", e.setup["cursor"]).
		Title(e.setup["menuCloseLabel"]).
		Html(e.setup["menuCloseIcon"])
	e.headerAddCloseListener(closeIcon)

	e.header = factoryBrowser.NewTagDiv().Append(
		factoryBrowser.NewTagSpan().Html(e.setup["menuTitle"]),
	)
	e.header.AddStyle("display", "flex")
	e.header.AddStyle("justify-content", "space-between")
	e.header.AddStyle("align-items", "center")
	e.header.AddStyle("background", e.setup["headerBackground"])
	e.header.AddStyle("padding", e.setup["headerPadding"])
	e.header.AddStyle("margin", e.setup["headerMargin"])
	e.header.AddStyle("font-family", e.setup["fontFamily"])
	e.header.AddStyle("font-weight", "bold")
	e.header.AddStyle("user-select", "none")

	if e.fixed {
		e.header.Append(
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

	e.content = factoryBrowser.NewTagDiv()
	e.content.AddStyle("display", "grid")
	e.content.AddStyle("gap", e.setup["contentGap"])
	e.content.AddStyle("padding", e.setup["contentPadding"])
	e.content.FadeFunc(e.contentFadeProgress)

	e.body.Append(
		e.header,
		e.content,
	)

	e.menu = factoryBrowser.NewTagDiv()
	e.content.Append(e.menu)

	e.body.HideForFade()
	e.body.FadeFunc(e.bodyFadeProgress)

	e.stage.Append(e.body)
	e.body.Fade(300 * time.Millisecond)

	if !e.fixed {
		e.hide()
		js.Global().Get("document").Call("addEventListener", "click", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			e.hide()
			return nil
		}))

		return
	}

	e.body.AddStyle("left", fmt.Sprintf("%vpx", e.bodyX))
	e.body.AddStyle("top", fmt.Sprintf("%vpx", e.bodyY))
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
func (e *ContextMenu) bodyFadeProgress(progress float64) {
	e.adjustContentWidth()

	if progress == 1.0 && !e.content.FadeStatus() {
		e.content.ShowForFade()
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
func (e *ContextMenu) contentFadeProgress(progress float64) {
	e.adjustContentWidth()

	if progress == 1.0 {
		if e.content.FadeStatus() {
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
func (e *ContextMenu) fadeShowContent() {
	e.content.AddStyle("visibility", "visible")
	e.menu.AddStyle("visibility", "visible")

	e.body.AddStyle("padding", e.setup["menuPadding"])
	e.content.AddStyle("padding", e.setup["contentPadding"])
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
func (e *ContextMenu) fadeHideContent() {
	e.content.AddStyle("visibility", "hidden")
	e.menu.AddStyle("visibility", "hidden")

	e.body.AddStyle("padding", "0")
	e.content.AddStyle("padding", "0")
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
func (e *ContextMenu) headerAddDragListener(dragIcon *html.TagSpan) {
	dragIcon.Get().Call("addEventListener", "mousedown", js.FuncOf(func(this js.Value, args []js.Value) any {
		e.isDragging = true
		e.offsetX = args[0].Get("clientX").Int() - e.body.Get().Call("getBoundingClientRect").Get("left").Int()
		e.offsetY = args[0].Get("clientY").Int() - e.body.Get().Call("getBoundingClientRect").Get("top").Int()
		return nil
	}))

	js.Global().Get("document").Call("addEventListener", "mousemove", js.FuncOf(func(this js.Value, args []js.Value) any {
		if e.isDragging {
			e.body.AddStyle("left", fmt.Sprintf("%vpx", args[0].Get("clientX").Int()-e.offsetX))
			e.body.AddStyle("top", fmt.Sprintf("%vpx", args[0].Get("clientY").Int()-e.offsetY))
		}
		return nil
	}))

	js.Global().Get("document").Call("addEventListener", "mouseup", js.FuncOf(func(this js.Value, args []js.Value) any {
		e.isDragging = false
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
func (e *ContextMenu) headerAddMinimizeListener(closeIcon *html.TagSpan) {
	closeIcon.Get().Call("addEventListener", "click", js.FuncOf(func(this js.Value, args []js.Value) any {
		args[0].Call("stopPropagation")
		e.content.Fade(300 * time.Millisecond)

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
func (e *ContextMenu) headerAddCloseListener(closeIcon *html.TagSpan) {
	closeIcon.Get().Call("addEventListener", "click", js.FuncOf(func(this js.Value, args []js.Value) any {
		args[0].Call("stopPropagation")

		e.body.Fade(200 * time.Millisecond)
		go func() {
			time.Sleep(1 * time.Second)
			e.body.Fade(200 * time.Millisecond)
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
func (e *ContextMenu) adjustContentWidth() {
	menuRect := e.menu.Get().Call("getBoundingClientRect")
	width := menuRect.Get("width").Int()
	height := menuRect.Get("height").Int()
	e.content.AddStyle("width", fmt.Sprintf("%vpx", width))
	e.content.AddStyle("height", fmt.Sprintf("%vpx", height))
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
func (e *ContextMenu) mountMenu(options []Options, container *html.TagDiv) {
	for _, option := range options {
		if option.Label == "-" {
			divider := factoryBrowser.NewTagHr()
			divider.AddStyle("margin", e.setup["dividerMargin"])
			container.Append(divider)
			continue

		} else if option.Type == "grid" && option.Items != nil && len(option.Items) > 0 {
			grid := factoryBrowser.NewTagDiv()
			grid.AddStyle("display", e.setup["gridDisplay"])
			grid.AddStyle("gridTemplateColumns", e.setup["gridGridTemplateColumns"])
			grid.AddStyle("gap", e.setup["gridGap"])
			grid.AddStyle("padding", e.setup["gridPadding"])

			for _, item := range option.Items {
				cell := factoryBrowser.NewTagDiv()
				cell.AddStyle("textAlign", e.setup["cellTextAlign"])
				cell.AddStyle("cursor", e.setup["cursor"])
				cell.AddStyle("border", e.setup["cellBorder"])
				cell.AddStyle("borderRadius", e.setup["cellBorderRadius"])
				cell.AddStyle("padding", e.setup["cellPadding"])
				cell.AddStyle("position", "relative")

				img := factoryBrowser.NewTagImg()
				img.Src(item.Icon, false)
				img.Alt(item.Label)
				img.AddStyle("width", e.setup["imgWidth"])
				img.AddStyle("height", e.setup["imgHeight"])
				img.AddStyle("display", e.setup["imgDisplay"])
				img.AddStyle("margin", e.setup["imgMargin"])
				img.AddStyle("padding", e.setup["imgPadding"])

				text := factoryBrowser.NewTagDiv()
				text.AddStyle("fontSize", e.setup["textFontSize"])
				text.AddStyle("fontFamily", e.setup["fontFamily"])
				if item.Submenu != nil && len(item.Submenu) > 0 {
					text.Html(fmt.Sprintf("<span style=\"flex:1; text-align:left;\">%v</span><span style=\"text-align:right;\">%v</span>", item.Label, e.setup["itemTextContent"]))
					text.AddStyle("display", e.setup["itemDisplay"])
					text.AddStyle("alignItems", e.setup["itemAlignItems"])
				} else {
					text.Text(item.Label)
				}

				cell.Append(img)
				cell.Append(text)

				if item.Submenu != nil && len(item.Submenu) > 0 {
					subMenu := factoryBrowser.NewTagDiv()
					subMenu.AddStyle("position", e.setup["submenuPosition"])
					subMenu.AddStyle("left", e.setup["submenuLeft"])
					subMenu.AddStyle("top", e.setup["submenuTop"])
					subMenu.AddStyle("background", e.setup["backgroundColor"])
					subMenu.AddStyle("border", e.setup["submenuBorder"])
					subMenu.AddStyle("boxShadow", e.setup["submenuBoxShadow"])
					subMenu.AddStyle("padding", e.setup["submenuPadding"])
					subMenu.AddStyle("display", e.setup["submenuDisplay"])
					subMenu.AddStyle("whiteSpace", e.setup["submenuWhiteSpace"])
					subMenu.AddStyle("zIndex", e.setup["submenuZIndex"])

					e.mountMenu(item.Submenu, subMenu)
					cell.Append(subMenu)

					cell.Get().Call("addEventListener", "mouseenter", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
						cell.AddStyle("background", e.setup["submenuBackground"])
						e.adjustSubMenuPosition(subMenu, cell)
						return nil
					}))

					cell.Get().Call("addEventListener", "mouseleave", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
						cell.AddStyle("background", "transparent")
						cell.AddStyle("border", e.setup["border"])
						subMenu.AddStyle("display", "none")

						return nil
					}))
				} else {
					cell.Get().Call("addEventListener", "click", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
						args[0].Call("stopPropagation")
						// gridItem.action(); // todo: fazer
						e.hide()
						return nil
					}))

					cell.Get().Call("addEventListener", "mouseenter", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
						cell.AddStyle("background", e.setup["submenuBackground"])
						cell.AddStyle("border", e.setup["border"])
						return nil
					}))

					cell.Get().Call("addEventListener", "mouseleave", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
						cell.AddStyle("background", "transparent")
						cell.AddStyle("border", e.setup["border"])
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
		item.AddStyle("fontSize", e.setup["textFontSize"])
		item.AddStyle("fontFamily", e.setup["fontFamily"])
		item.AddStyle("padding", e.setup["itemPadding"])
		item.AddStyle("cursor", e.setup["itemCursor"])
		item.AddStyle("position", e.setup["itemPosition"])

		item.Get().Call("addEventListener", "mouseenter", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			item.AddStyle("background", e.setup["submenuBackground"])
			return nil
		}))
		item.Get().Call("addEventListener", "mouseleave", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			item.AddStyle("background", "transparent")
			return nil
		}))

		// submenu em linha
		if option.Submenu != nil && len(option.Submenu) > 0 {
			item.Html(fmt.Sprintf("<span style=\"flex:1; text-align:left;\">%v</span><span style=\"text-align:right;\">%v</span>", option.Label, e.setup["itemTextContent"]))
			item.AddStyle("display", e.setup["itemDisplay"])
			item.AddStyle("alignItems", e.setup["itemAlignItems"])

			subMenu := factoryBrowser.NewTagDiv()
			subMenu.AddStyle("position", e.setup["submenuPosition"])
			subMenu.AddStyle("left", e.setup["submenuLeft"])
			subMenu.AddStyle("top", e.setup["submenuTop"])
			subMenu.AddStyle("background", e.setup["backgroundColor"])
			subMenu.AddStyle("border", e.setup["submenuBorder"])
			subMenu.AddStyle("boxShadow", e.setup["submenuBoxShadow"])
			subMenu.AddStyle("padding", e.setup["submenuPadding"])
			subMenu.AddStyle("display", e.setup["submenuDisplay"])
			subMenu.AddStyle("whiteSpace", e.setup["submenuWhiteSpace"])
			subMenu.AddStyle("zIndex", e.setup["submenuZIndex"])

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
				//opt.action(); // todo: fazer
				e.hide()
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
func (e *ContextMenu) adjustSubMenuPosition(subMenu, cell *html.TagDiv) {
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

// show
//
// English:
//
//	Show the menu when the contextual menu is triggered by the mouse
//
// Português:
//
//	Mostra o menu quando o menu contextual é acionado pelo mouse
func (e *ContextMenu) show(x, y int) {
	e.body.AddStyle("display", "block")
	e.body.AddStyle("left", "0px")
	e.body.AddStyle("top", "0px")

	e.adjustContentWidth()

	bbox := e.body.Get().Call("getBoundingClientRect")
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

	e.body.AddStyle("left", strconv.FormatInt(int64(adjustedX), 10)+"px")
	e.body.AddStyle("top", strconv.FormatInt(int64(adjustedY), 10)+"px")
}

// hide
//
// Esconde o menu quando este é configurado para ser um menu contextual
func (e *ContextMenu) hide() {
	e.body.AddStyle("display", "none")
}

func (e *ContextMenu) max(x, y int) (max int) {
	if x > y {
		return x
	}

	return y
}

func main() {

	stage := factoryBrowser.NewStage()

	contextMenu := new(ContextMenu)
	contextMenu.Stage(stage)
	contextMenu.FixedMenu(200, 200)
	contextMenu.AttachMenu(js.Global().Get("document"))
	contextMenu.Init()
	contextMenu.Menu([]Options{
		{
			Type: "grid",
			Items: []Options{
				{
					Label: "cat",
					Icon:  "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
				},
				{
					Label: "cat",
					Icon:  "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
				},
				{
					Label: "cat",
					Icon:  "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
					Type:  "grid",
					Submenu: []Options{
						{
							Type: "grid",
							Items: []Options{
								{
									Label: "cat",
									Icon:  "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
								},
								{
									Label: "cat",
									Icon:  "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
									Type:  "grid",
									Submenu: []Options{
										{
											Type: "grid",
											Items: []Options{
												{
													Label: "cat",
													Icon:  "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
												},
												{
													Label: "cat",
													Icon:  "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
												},
												{
													Label: "cat",
													Icon:  "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
												},
												{
													Label: "cat",
													Icon:  "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
												},
												{
													Label: "cat",
													Icon:  "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
												},
												{
													Label: "cat",
													Icon:  "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
													Type:  "grid",
													Submenu: []Options{
														{
															Type: "grid",
															Items: []Options{
																{
																	Label: "cat",
																	Icon:  "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
																},
																{
																	Label: "cat",
																	Icon:  "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
																},
																{
																	Label: "cat",
																	Icon:  "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
																},
																{
																	Label: "cat",
																	Icon:  "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
																},
																{
																	Label: "cat",
																	Icon:  "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
																	Type:  "grid",
																	Submenu: []Options{
																		{
																			Type: "grid",
																			Items: []Options{
																				{
																					Label: "cat",
																					Icon:  "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
																				},
																				{
																					Label: "cat",
																					Icon:  "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
																				},
																				{
																					Label: "cat",
																					Icon:  "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
																				},
																				{
																					Label: "cat",
																					Icon:  "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
																				},
																				{
																					Label: "cat",
																					Icon:  "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
																				},
																				{
																					Label: "cat",
																					Icon:  "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
																				},
																			},
																		},
																	},
																},
																{
																	Label: "cat",
																	Icon:  "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
																},
															},
														},
													},
												},
											},
										},
									},
								},
								{
									Label: "cat",
									Icon:  "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
								},
								{
									Label: "cat",
									Icon:  "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
								},
								{
									Label: "cat",
									Icon:  "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
								},
								{
									Label: "cat",
									Icon:  "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
								},
							},
						},
					},
				},
				{
					Label: "cat",
					Icon:  "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
				},
				{
					Label: "cat",
					Icon:  "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
				},
				{
					Label: "cat",
					Icon:  "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
				},
			},
		},
		{
			Label: "Option 1",
			Submenu: []Options{
				{
					Label: "Option 1",
				},
				{
					Label: "Option 2",
				},
			},
		},
		{
			Label: "Option 2",
		},
		{
			Label: "-",
		},
		{
			Label: "Option 3",
			Submenu: []Options{
				{
					Type: "grid",
					Items: []Options{
						{
							Label: "cat",
							Icon:  "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
						},
						{
							Label: "cat",
							Icon:  "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
						},
						{
							Label: "cat",
							Icon:  "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
						},
						{
							Label: "cat",
							Icon:  "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
						},
						{
							Label: "cat",
							Icon:  "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
						},
						{
							Label: "cat",
							Icon:  "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
						},
					},
				},
			},
		},
	})

	done := make(chan struct{})
	done <- struct{}{}
}
