package main

import (
	"fmt"
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/browser/stage"
	"strconv"
	"syscall/js"
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
	menu  *html.TagDiv
	setup map[string]string
	stage *stage.Stage
	fixed bool
}

func (e *ContextMenu) SetAsFixed() {
	e.fixed = true
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
		e.setup["gridGridTemplateColumns"] = "repeat(2, 1fr)"
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

	if _, found := e.setup["cellCursor"]; !found {
		e.setup["cellCursor"] = "pointer"
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

	if _, found := e.setup["submenuDisplay"]; !found {
		e.setup["submenuDisplay"] = "none"
	}

	if _, found := e.setup["submenuWhiteSpace"]; !found {
		e.setup["submenuWhiteSpace"] = "nowrap"
	}

	if _, found := e.setup["submenuZIndex"]; !found {
		e.setup["submenuZIndex"] = "1001"
	}
}

func (e *ContextMenu) Init() {
	e.setupInit()

	e.menu = factoryBrowser.NewTagDiv()
	e.menu.AddStyle("position", "absolute")
	e.menu.AddStyle("background", e.setup["backgroundColor"])
	e.menu.AddStyle("border", e.setup["border"])
	e.menu.AddStyle("boxShadow", e.setup["shadow"])
	e.menu.AddStyle("padding", "5px")
	e.menu.AddStyle("zIndex", "1000")

	//e.hide()

	//pai := factoryBrowser.NewTagDiv().Class("panel open")
	//header := factoryBrowser.NewTagDiv().Class("panel open").Append(
	//	factoryBrowser.NewTagDiv().Class("panelHeader").Append(
	//		factoryBrowser.NewTagDiv().Class("headerText").Html("Control panel"),
	//		factoryBrowser.NewTagDiv().AddStyle("cursor", "move").Html("◇"),
	//		factoryBrowser.NewTagDiv().Html("&nbsp;"),
	//		factoryBrowser.NewTagDiv().AddStyle("cursor", "pointer").Html("▾"),
	//		factoryBrowser.NewTagDiv().Html("&nbsp;"),
	//		factoryBrowser.NewTagDiv().AddStyle("cursor", "pointer").Html("⊗"),
	//	),

	//factoryBrowser.NewTagDiv().Class("panelBody").Append(
	//	factoryBrowser.NewTagDiv().Class("panelCel").Append(
	//		factoryBrowser.NewTagDiv().Class("labelCel").Append(
	//			factoryBrowser.NewTagDiv().Class("labelText").Html("Easing tween time"),
	//			factoryBrowser.NewTagDiv().Class("closeIcon").Html("ˇ"),
	//		),
	//		factoryBrowser.NewTagDiv().Class("compCel").Append(
	//			factoryBrowser.NewTagDiv().Class("component"),
	//			factoryBrowser.NewTagDiv().Class("component").Append(e.menu),
	//		),
	//	),
	//),
	//)

	e.stage.Append(e.menu)

	if e.fixed {
		e.menu.AddStyle("top", "50px")
		e.menu.AddStyle("left", "10px")
		e.hide()
		js.Global().Get("document").Call("addEventListener", "click", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			e.hide()
			return nil
		}))
	}
}

func (e *ContextMenu) Menu(options []Options) {
	e.menu.Html("")
	e.mountMenu(options, e.menu)
}

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
				cell.AddStyle("cursor", e.setup["cellCursor"])
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
						cell.AddStyle("border", e.setup["border"])
						subMenu.AddStyle("display", "block")

						submenuRect := subMenu.Get().Call("getBoundingClientRect")
						screenWidth := js.Global().Get("window").Get("innerWidth").Int()
						screenHeight := js.Global().Get("window").Get("innerHeight").Int()

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
				subMenu.AddStyle("display", "block")

				window := js.Global().Get("window")
				itemClientRect := item.Get().Call("getBoundingClientRect")

				submenuRect := subMenu.Get().Call("getBoundingClientRect")
				screenWidth := window.Get("innerWidth").Int()
				//screenHeight := window.Get("innerHeight").Int()

				itemRight := itemClientRect.Get("right").Int()

				if itemRight+submenuRect.Get("width").Int() > screenWidth {
					subMenu.AddStyle("left", "auto")
					subMenu.AddStyle("right", "100%")
				} else {
					subMenu.AddStyle("left", "100%")
					subMenu.AddStyle("right", "auto")
				}

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

func (e *ContextMenu) show(x, y int) {
	e.menu.AddStyle("display", "block")
	e.menu.AddStyle("left", "0px")
	e.menu.AddStyle("top", "0px")

	bbox := e.menu.Get().Call("getBoundingClientRect")
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

	e.menu.AddStyle("left", strconv.FormatInt(int64(adjustedX), 10)+"px")
	e.menu.AddStyle("top", strconv.FormatInt(int64(adjustedY), 10)+"px")
}

func (e *ContextMenu) hide() {
	e.menu.AddStyle("display", "none")
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
	contextMenu.Init()
	contextMenu.AttachMenu(js.Global().Get("document"))
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
