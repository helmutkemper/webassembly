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
}

func (e *ContextMenu) Init(stage *stage.Stage) {
	e.setup = map[string]string{
		"shadow":                  "0 2px 5px rgba(0,0,0,0.2)",
		"border":                  "0px solid #ccc",
		"backgroundColor":         "#fff",
		"dividerMargin":           "5px 0",
		"gridDisplay":             "grid",
		"gridGridTemplateColumns": "repeat(2, 1fr)",
		"gridGap":                 "8px",
		"gridPadding":             "2px",
		"cellTextAlign":           "center",
		"cellCursor":              "pointer",
		"cellBorder":              "0px solid #ccc",
		"cellBorderRadius":        "4px",
		"cellPadding":             "5px",
		"imgWidth":                "32px",
		"imgHeight":               "32px",
		"imgDisplay":              "block",
		"imgMargin":               "0 auto 5px",
		"imgPadding":              "5px",
		"textFontSize":            "12px",
		"fontFamily":              "Arial, sans-serif",
		"itemPadding":             "5px 10px",
		"itemCursor":              "pointer",
		"itemPosition":            "relative",
		"itemTextContent":         "&nbsp;&nbsp;▶",
		"itemDisplay":             "flex",
		"itemAlignItems":          "center",
		"submenuPosition":         "absolute",
		"submenuLeft":             "100%",
		"submenuTop":              "0",
		"submenuBackground":       "#ccc",
		"submenuBorder":           "0px solid #ccc",
		"submenuBoxShadow":        "0 2px 5px rgba(0,0,0,0.2)",
		"submenuPadding":          "5px",
		"submenuDisplay":          "none",
		"submenuWhiteSpace":       "nowrap",
		"submenuZIndex":           "1001",
	}
	e.menu = factoryBrowser.NewTagDiv()
	e.menu.AddStyle("position", "absolute")
	e.menu.AddStyle("background", e.setup["backgroundColor"])
	e.menu.AddStyle("border", e.setup["border"])
	e.menu.AddStyle("boxShadow", e.setup["shadow"])
	e.menu.AddStyle("padding", "5px")
	e.menu.AddStyle("zIndex", "1000")
	e.hide()
	stage.Append(e.menu)

	js.Global().Get("document").Call("addEventListener", "click", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		e.hide()
		return nil
	}))
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
			container.Append(divider) //container
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
				img.Src(item.Icon, true)
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

						cellRight := cell.Get().Call("getBoundingClientRect").Get("right").Int()

						if cellRight+submenuRect.Get("width").Int() > screenWidth {
							subMenu.AddStyle("left", "auto")
							subMenu.AddStyle("right", "100%")
						} else {
							subMenu.AddStyle("left", "100%")
							subMenu.AddStyle("right", "auto")
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

				submenuRect := subMenu.Get().Call("getBoundingClientRect")
				screenWidth := js.Global().Get("window").Get("innerWidth").Int()

				itemRight := item.Get().Call("getBoundingClientRect").Get("right").Int()

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
	// return the max value
	if x > y {
		return x
	}

	return y
}

func main() {

	stage := factoryBrowser.NewStage()

	contextMenu := new(ContextMenu)
	contextMenu.Init(stage)
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
