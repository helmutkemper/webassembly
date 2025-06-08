package devices

import (
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/platform/components"
)

type Menu struct {
	node    html.Compatible
	content []components.MenuOptions
	title   string
}

func (e *Menu) SetTitle(title string) {
	e.title = title
}

func (e *Menu) SetNode(node html.Compatible) {
	e.node = node
}

func (e *Menu) SetContent(content []components.MenuOptions) {
	e.content = content
}

func (e *Menu) Init() {

	contextMenu := new(components.ContextMenu)
	contextMenu.Title(e.title)
	contextMenu.Menu(e.content)
	contextMenu.AttachMenu(e.node)
	contextMenu.Columns(3)
	contextMenu.Init()
}
