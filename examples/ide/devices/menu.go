package devices

import (
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/platform/components"
)

type Menu struct {
	contextMenu *components.ContextMenu
	node        html.Compatible
	contentFunc func() (options []components.MenuOptions)
	title       string
}

func (e *Menu) ReInit() {
	if e.contextMenu == nil {
		return
	}

	e.contextMenu.ReInit()
}

func (e *Menu) SetTitle(title string) {
	e.title = title
}

func (e *Menu) SetNode(node html.Compatible) {
	e.node = node
}

func (e *Menu) SetContentFunc(f func() (content []components.MenuOptions)) {
	e.contentFunc = f
}

func (e *Menu) Init() {

	e.contextMenu = new(components.ContextMenu)
	e.contextMenu.Title(e.title)
	e.contextMenu.MenuFunc(e.contentFunc)
	e.contextMenu.AttachMenu(e.node)
	e.contextMenu.Columns(3)
	e.contextMenu.Init()
}
