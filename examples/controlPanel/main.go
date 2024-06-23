package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/html"
	"github.com/helmutkemper/iotmaker.webassembly/platform/components"
	"log"
)

type Control struct {
	components.Components
	Panel Panel `wasmPanel:"type:panel"`
}

type Panel struct {
	Header string `wasmPanel:"type:headerText"`
	Body   Body   `wasmPanel:"type:panelBody"`
}

type Body struct {
	Color        *ColorAdjust   `wasmPanel:"type:component;label:Color Adjust"`
	RunCommand   *RunCommand    `wasmPanel:"type:component;label:Run Command"`
	SelectFilter *SelectFilter  `wasmPanel:"type:component;label:File format"`
	Payment      *RadioGroup    `wasmPanel:"type:component;label:Payment"`
	Options      *CheckboxGroup `wasmPanel:"type:component;label:Options"`
	Colors       *ColorGroup    `wasmPanel:"type:component;label:Colors"`
}

type OnChangeEvent struct {
	IsTrusted bool    `wasmGet:"isTrusted"`
	Value     float64 `wasmGet:"value"`
}

func (e *OnChangeEvent) OnChange(event OnChangeEvent, reference ColorRange) {
	log.Printf("> Trusted: %+v", event.IsTrusted)
	log.Printf("> Value:   %+v", event.Value)
	reference.Max(500)
}

type ColorRange struct {
	components.Range

	TagRange    *html.TagInputRange  `wasmPanel:"type:inputTagRange"`
	TagNumber   *html.TagInputNumber `wasmPanel:"type:inputTagNumber"`
	Color       float64              `wasmPanel:"type:value;min:0;max:50;step:1;default:0"`
	ColorChange *OnChangeEvent       `wasmPanel:"type:listener;event:change;func:OnChange"`
}

func (e *ColorRange) Init() {
	e.Step(1)
	e.Max(10)
	e.Min(0)
	e.Value(5)
}

type ColorAdjust struct {
	Red   *ColorRange `wasmPanel:"type:range;label:Red"`
	Green *ColorRange `wasmPanel:"type:range;label:Green"`
	Blue  *ColorRange `wasmPanel:"type:range;label:Blue"`
}

type RunCommand struct {
	Button *ButtonEvent `wasmPanel:"type:button;label:Exec. command;value:Click me"`
	Undo   *ButtonEvent `wasmPanel:"type:button;label:Undo last exec.;value:Undo"`
}

type SelectFilter struct {
	Red        *ColorRange  `wasmPanel:"type:range;label:Red"`
	FileFormat []Select     `wasmPanel:"type:select;label:Select the file format"`
	Button     *ButtonEvent `wasmPanel:"type:button;label:Exec. command;value:Click me"`
}

type RadioGroup struct {
	Payments []Radio `wasmPanel:"type:radio;label:Payment method"`
}

type CheckboxGroup struct {
	Options []Checkbox `wasmPanel:"type:checkbox;label:Please, select all"`
}

type ColorGroup struct {
	Input  ColorEvent `wasmPanel:"type:color"`
	Output ColorEvent `wasmPanel:"type:color"`
}

type ColorEvent struct {
	Disabled  bool   `wasmPanel:"type:disabled"`
	Value     string `wasmPanel:"type:value"`
	Label     string `wasmPanel:"type:label"`
	OnPress   func() `wasmPanel:"type:onpress"`
	OnRelease func() `wasmPanel:"type:onRelease"`
	OnClick   func() `wasmPanel:"type:onclick"`
}

type Radio struct {
	Name     string `wasmPanel:"type:name"`
	Label    string `wasmPanel:"type:label"`
	Value    string `wasmPanel:"type:value"`
	Selected bool   `wasmPanel:"type:selected"`
	Disabled bool   `wasmPanel:"type:disabled"`
}

type Checkbox struct {
	Name     string `wasmPanel:"type:name"`
	Label    string `wasmPanel:"type:label"`
	Value    string `wasmPanel:"type:value"`
	Selected bool   `wasmPanel:"type:selected"`
	Disabled bool   `wasmPanel:"type:disabled"`
}

type Select struct {
	Label    string `wasmPanel:"type:label"`
	Value    string `wasmPanel:"type:value"`
	Disabled bool   `wasmPanel:"type:disabled"`
	Selected bool   `wasmPanel:"type:selected"`
	OnSelect func() `wasmPanel:"type:onselect"`
}

type OnClickEvent struct {
	IsTrusted bool   `wasmGet:"isTrusted"`
	Value     string `wasmGet:"value"`
}

func (e *OnClickEvent) OnClick(event OnClickEvent, ref ButtonEvent) {
	log.Printf("Trusted: %v", event.IsTrusted)
	log.Printf("Value:   %v", event.Value)
	ref.Value("Clicked")
}

type ButtonEvent struct {
	components.Button

	Label      string        `wasmPanel:"type:value;default:Ok"`
	RunCommand *OnClickEvent `wasmPanel:"type:listener;event:click;func:OnClick"`
}

func (e *ButtonEvent) Init() {
	e.Value("Initialized")
}

func (e *Control) Init() (err error) {
	err = e.Components.Init(e)
	return
}

func main() {

	c := Control{
		Panel: Panel{
			Header: "Control Panel",
			Body: Body{
				RunCommand: &RunCommand{
					Button: &ButtonEvent{
						Label: "Label set",
					},
				},
				SelectFilter: &SelectFilter{
					FileFormat: []Select{
						{
							Label:    "Please select",
							Value:    "",
							Disabled: false,
							Selected: false,
							OnSelect: nil,
						},
						{
							Label:    "label 1",
							Value:    "value 1",
							Disabled: false,
							Selected: false,
							OnSelect: nil,
						},
						{
							Label:    "label 2",
							Value:    "value 2",
							Disabled: false,
							Selected: true,
							OnSelect: nil,
						},
					},
				},
				Payment: &RadioGroup{
					Payments: []Radio{
						{
							Name:     "payment",
							Label:    "Option 1",
							Value:    "vista",
							Selected: false,
							Disabled: false,
						},
						{
							Name:     "payment",
							Label:    "Opti 2",
							Value:    "credit",
							Selected: false,
							Disabled: false,
						},
						{
							Name:     "payment",
							Label:    "Opt 3",
							Value:    "Option 3",
							Selected: false,
							Disabled: false,
						},
					},
				},
				Options: &CheckboxGroup{
					Options: []Checkbox{
						{
							Name:     "payment",
							Label:    "Option 1",
							Value:    "vista",
							Selected: false,
							Disabled: false,
						},
						{
							Name:     "payment",
							Label:    "Opti 2",
							Value:    "credit",
							Selected: false,
							Disabled: false,
						},
						{
							Name:     "payment",
							Label:    "Opt 3",
							Value:    "Option 3",
							Selected: false,
							Disabled: false,
						},
					},
				},
				Colors: &ColorGroup{
					Input: ColorEvent{
						Disabled:  false,
						Label:     "Input color",
						Value:     "#11aa66",
						OnPress:   nil,
						OnRelease: nil,
						OnClick:   nil,
					},
					Output: ColorEvent{
						Disabled:  true,
						Label:     "Output color",
						Value:     "#cc3300",
						OnPress:   nil,
						OnRelease: nil,
						OnClick:   nil,
					},
				},
			},
		},
	}
	if err := c.Init(); err != nil {
		log.Printf("%v", err)
		panic(err)
	}

	//red.TagNumber.Min(-20).Max(0).Value(-10)
	//red.TagRange.Min(-20).Max(0).Value(-10)
	//red.SetMax(39)

	//red.SetMin(0)
	//red.SetMax(5)
	//red.SetStep(1)
	//red.TagRange.Min(0).Max(3).Step(1)
	//red.TagNumber.Min(0).Max(3).Step(1)

	done := make(chan struct{})
	done <- struct{}{}

}
