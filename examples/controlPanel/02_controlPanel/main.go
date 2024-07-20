package main

import (
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/platform/components"
	"log"
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
	BoatAnimation *BoatAdjust `wasmPanel:"type:component;label:Boat dragging effect"`
}

type BoatAdjust struct {
	Dragging *DraggingEffect `wasmPanel:"type:range;label:effect"`
}

type DraggingEffect struct {
	components.Range

	TagRange    *html.TagInputRange  `wasmPanel:"type:inputTagRange"`
	TagNumber   *html.TagInputNumber `wasmPanel:"type:inputTagNumber"`
	Color       float64              `wasmPanel:"type:value;min:2;max:50;step:1;default:15"`
	ColorChange *OnChangeEvent       `wasmPanel:"type:listener;event:change;func:OnChangeEvent"`
}

func (e *DraggingEffect) MathematicalFormula(min, max, value float64) (result float64) {
	return (max - value) + min
}

func (e *DraggingEffect) Init() {
	e.TagNumber.Value(e.MathematicalFormula(2, 50, e.TagRange.GetValue()))
	e.TagRange.Value(e.MathematicalFormula(2, 50, e.TagNumber.GetValue()))
}

type OnChangeEvent struct {
	IsTrusted bool    `wasmGet:"isTrusted"`
	Value     float64 `wasmGet:"value"`
	Min       float64 `wasmGet:"min"`
	Max       float64 `wasmGet:"max"`
	Type      string  `wasmGet:"type"`
}

func (e *OnChangeEvent) OnChangeEvent(event OnChangeEvent, reference DraggingEffect) {
	log.Printf("isTrusted: %v", event.IsTrusted)
	log.Printf("type: %v", event.Type)
	log.Printf("min: %v", event.Min)
	log.Printf("max: %v", event.Max)
	log.Printf("value: %v", event.Value)

	log.Printf("reference.TagNumber.GetValue(): %v", reference.TagNumber.GetValue())
	log.Printf("reference.TagRange.GetValue(): %v", reference.TagRange.GetValue())
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
