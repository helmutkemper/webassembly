package main

import (
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/platform/components"
)

type ComponentControlPanel struct {
	components.Components

	Panel *ControlPanel `wasmPanel:"type:panel"`
}

func (e *ComponentControlPanel) Init() (panel *html.TagDiv, err error) {
	panel, err = e.Components.Init(e)
	return
}

type ControlPanel struct {
	Header string `wasmPanel:"type:headerText;label:Control panel"`
	Body   *Body  `wasmPanel:"type:panelBody"`
}

type Body struct {
}

func main() {
	var err error
	var panel *html.TagDiv

	controlPanel := ComponentControlPanel{}
	if panel, err = controlPanel.Init(); err != nil {
		panic(err)
	}

	stage := factoryBrowser.NewStage()
	stage.Append(panel)

	done := make(chan struct{})
	done <- struct{}{}

}
