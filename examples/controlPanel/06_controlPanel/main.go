package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/browser/html"
	"github.com/helmutkemper/iotmaker.webassembly/platform/algorithm"
	"github.com/helmutkemper/iotmaker.webassembly/platform/components"
	"github.com/helmutkemper/iotmaker.webassembly/platform/easingTween"
	"github.com/helmutkemper/iotmaker.webassembly/platform/factoryAlgorithm"
	"github.com/helmutkemper/iotmaker.webassembly/platform/factoryColor"
	"math"
	"time"
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

type OnChangeEvent struct {
	IsTrusted bool    `wasmGet:"isTrusted"`
	Value     float64 `wasmGet:"value"`
	Min       float64 `wasmGet:"min"`
	Max       float64 `wasmGet:"max"`
	Type      string  `wasmGet:"type"`
}

func (e *OnChangeEvent) OnChange(event OnChangeEvent, reference *Body) {
	//ref := reference.BoatAnimation.Dragging

	//log.Printf("reference.TagNumber.GetValue(): %v", ref.TagNumber.GetValue())
	//log.Printf("reference.TagRange.GetValue(): %v", ref.TagRange.GetValue())
}

func (e *OnChangeEvent) OnInputEvent(event OnChangeEvent, reference *Body) {
	ref := reference.BoatAnimation.Dragging
	switch event.Type {
	case "range":
		ref.TagNumber.Value(ref.MathematicalFormula(event.Min, event.Max, ref.TagRange.GetValue()))
	case "number":
		ref.TagRange.Value(ref.MathematicalFormula(event.Min, event.Max, ref.TagNumber.GetValue()))
	}
}

type DraggingEffect struct {
	components.Range

	TagRange    *html.TagInputRange  `wasmPanel:"type:inputTagRange"`
	TagNumber   *html.TagInputNumber `wasmPanel:"type:inputTagNumber"`
	Dragging    float64              `wasmPanel:"type:value;min:2;max:50;step:1;default:15"`
	ColorChange *OnChangeEvent       `wasmPanel:"type:listener;event:change;func:OnChange"`
	RangeChange *OnChangeEvent       `wasmPanel:"type:listener;event:input;func:OnInputEvent"`
}

func (e *DraggingEffect) MathematicalFormula(min, max, value float64) (result float64) {
	return (max - value) + min
}

func (e *DraggingEffect) Init() {
	e.TagNumber.Value(e.MathematicalFormula(2, 50, e.TagRange.GetValue()))
	e.TagRange.Value(e.MathematicalFormula(2, 50, e.TagNumber.GetValue()))
}

type BoatAdjust struct {
	Dragging *DraggingEffect   `wasmPanel:"type:range;label:effect"`
	Tween    *TweenSelect      `wasmPanel:"type:select;label:Tween function"`
	Start    *EasingTweenStart `wasmPanel:"type:button;label:start easing tween"`
}

type TweenSelect struct {
	components.Select

	SelectTag *html.TagSelect `wasmPanel:"type:inputTagSelect"`
	List      *[]TweenType    `wasmPanel:"type:value"` //;default:label 1,value 1,>label 2,value 2,label 3,value 3
	Change    *OnSelectChange `wasmPanel:"type:listener;event:change;func:OnChangeEvent"`
	//todo: falta a captura de ref
}

func (e *TweenSelect) Init() {

	e.SelectTag.
		NewOptionGroup("linear", false).
		NewOption("linear", "Linear", false, true).
		NewOptionGroup("back", false).
		NewOption("ease in back", "EaseInBack", false, false).
		NewOption("ease out back", "EaseOutBack", false, false).
		NewOption("ease in out back", "EaseInOutBack", false, false).
		NewOptionGroup("bounce", false).
		NewOption("ease in bounce", "EaseInBounce", false, false).
		NewOption("ease out bounce", "EaseOutBounce", false, false).
		NewOption("ease in out bounce", "EaseInOutBounce", false, false).
		NewOptionGroup("circular", false).
		NewOption("ease in circular", "EaseInCircular", false, false).
		NewOption("ease out circular", "EaseOutCircular", false, false).
		NewOption("ease in out circular", "EaseInOutCircular", false, false).
		NewOptionGroup("cubic", false).
		NewOption("ease in cubic", "EaseInCubic", false, false).
		NewOption("ease out cubic", "EaseOutCubic", false, false).
		NewOption("ease in out cubic", "EaseInOutCubic", false, false).
		NewOptionGroup("elastic", false).
		NewOption("ease in elastic", "EaseInElastic", false, false).
		NewOption("ease out elastic", "EaseOutElastic", false, false).
		NewOption("ease in out elastic", "EaseInOutElastic", false, false).
		NewOptionGroup("exponential", false).
		NewOption("ease in exponential", "EaseInExponential", false, false).
		NewOption("ease out exponential", "EaseOutExponential", false, false).
		NewOption("ease in out exponential", "EaseInOutExponential", false, false).
		NewOptionGroup("quadratic", false).
		NewOption("ease in quadratic", "EaseInQuadratic", false, false).
		NewOption("ease out quadratic", "EaseOutQuadratic", false, false).
		NewOption("ease in out quadratic", "EaseInOutQuadratic", false, false).
		NewOptionGroup("quartic", false).
		NewOption("ease in quartic", "EaseInQuartic", false, false).
		NewOption("ease out quartic", "EaseOutQuartic", false, false).
		NewOption("ease in out quartic", "EaseInOutQuartic", false, false).
		NewOptionGroup("quintic", false).
		NewOption("ease in quintic", "EaseInQuintic", false, false).
		NewOption("ease out quintic", "EaseOutQuintic", false, false).
		NewOption("ease in out quintic", "EaseInOutQuintic", false, false).
		NewOptionGroup("sine", false).
		NewOption("ease in sine", "EaseInSine", false, false).
		NewOption("ease out sine", "EaseOutSine", false, false).
		NewOption("ease in out sine", "EaseInOutSine", false, false)

}

type OnSelectChange struct {
	Value    string `wasmGet:"value"`
	function func(interactionCurrent, interactionTotal, currentPercentage, startValue, endValue, delta float64) float64
}

func (e *OnSelectChange) OnChangeEvent(event OnSelectChange, reference *Body) {
	// https://github.com/ai/easings.net/blob/master/src/easings/easingsFunctions.ts
	//log.Printf("value: %v", event.Value)
	switch event.Value {
	case "Linear":
		e.function = easingTween.KLinear
	case "EaseInBack":
		e.function = easingTween.KEaseInBack
	case "EaseInBounce":
		e.function = easingTween.KEaseInBounce
	case "EaseInCircular":
		e.function = easingTween.KEaseInCircular
	case "EaseInCubic":
		e.function = easingTween.KEaseInCubic
	case "EaseInElastic":
		e.function = easingTween.KEaseInElastic
	case "EaseInExponential":
		e.function = easingTween.KEaseInExponential
	case "EaseInOutBack":
		e.function = easingTween.KEaseInOutBack
	case "EaseInOutBounce":
		e.function = easingTween.KEaseInOutBounce
	case "EaseInOutCircular":
		e.function = easingTween.KEaseInOutCircular
	case "EaseInOutCubic":
		e.function = easingTween.KEaseInOutCubic
	case "EaseInOutElastic":
		e.function = easingTween.KEaseInOutElastic
	case "EaseInOutExponential":
		e.function = easingTween.KEaseInOutExponential
	case "EaseInOutQuadratic":
		e.function = easingTween.KEaseInOutQuadratic
	case "EaseInOutQuartic":
		e.function = easingTween.KEaseInOutQuartic
	case "EaseInOutQuintic":
		e.function = easingTween.KEaseInOutQuintic
	case "EaseInOutSine":
		e.function = easingTween.KEaseInOutSine
	case "EaseInQuadratic":
		e.function = easingTween.KEaseInQuadratic
	case "EaseInQuartic":
		e.function = easingTween.KEaseInQuartic
	case "EaseInQuintic":
		e.function = easingTween.KEaseInQuintic
	case "EaseInSine":
		e.function = easingTween.KEaseInSine
	case "EaseOutBack":
		e.function = easingTween.KEaseOutBack
	case "EaseOutBounce":
		e.function = easingTween.KEaseOutBounce
	case "EaseOutCircular":
		e.function = easingTween.KEaseOutCircular
	case "EaseOutCubic":
		e.function = easingTween.KEaseOutCubic
	case "EaseOutElastic":
		e.function = easingTween.KEaseOutElastic
	case "EaseOutExponential":
		e.function = easingTween.KEaseOutExponential
	case "EaseOutQuadratic":
		e.function = easingTween.KEaseOutQuadratic
	case "EaseOutQuartic":
		e.function = easingTween.KEaseOutQuartic
	case "EaseOutQuintic":
		e.function = easingTween.KEaseOutQuintic
	case "EaseOutSine":
		e.function = easingTween.KEaseOutSine
	}
}

type TweenType struct {
	Label    string `wasmPanel:"type:label"`
	Value    string `wasmPanel:"type:value"`
	Disabled bool   `wasmPanel:"type:disabled"`
	Selected bool   `wasmPanel:"type:selected"`
}

type OnClickEvent struct {
	IsTrusted bool   `wasmGet:"isTrusted"`
	Value     string `wasmGet:"value"`
}

func (e *OnClickEvent) OnClickEvent(event OnClickEvent, reference *Body) {
	ref := reference.BoatAnimation
	//log.Printf("Trusted: %v", event.IsTrusted)
	//log.Printf("Value:   %v", event.Value)

	var value = ref.Dragging.TagNumber.GetValue()

	t := new(easingTween.Tween)
	t.SetDuration(time.Duration(value)*time.Second).
		SetValues(0, 10000).
		SetOnStepFunc(tagDivRocket.EasingTweenWalkingAndRotateIntoPoints()).
		SetLoops(0).
		SetArgumentsFunc(any(tagDivRocket)).
		SetTweenFunc(reference.BoatAnimation.Tween.Change.function).
		SetDoNotReverseMotion().
		Start()
}

type EasingTweenStart struct {
	components.Button

	Label      string        `wasmPanel:"type:value;label:Start"`
	RunCommand *OnClickEvent `wasmPanel:"type:listener;event:click;func:OnClickEvent"`
}

func (e *EasingTweenStart) Init() {
	//e.Value("Initialized")
}

var canvas *html.TagCanvas
var tagDivRocket *html.TagDiv

func main() {
	var err error
	var panel *html.TagDiv

	stage := factoryBrowser.NewStage()

	controlPanel := ComponentControlPanel{
		//Panel: &ControlPanel{
		//	Body: &Body{
		//		BoatAnimation: &BoatAdjust{
		//			Tween: &TweenSelect{
		//				List: &[]TweenType{
		//					{
		//						Label: "label set 01",
		//						Value: "value set 01",
		//						//Disabled: true,
		//					},
		//					{
		//						Label: "label set 02",
		//						Value: "value set 02",
		//						//Selected: true,
		//					},
		//				},
		//			},
		//		},
		//	},
		//},
	}
	if panel, err = controlPanel.Init(); err != nil {
		panic(err)
	}

	canvas = factoryBrowser.NewTagCanvas(stage.GetWidth(), stage.GetHeight())
	stage.Append(canvas)

	border := 50.0
	wight := 400.0
	height := 400.0

	var bezier = BezierCurve(border, wight, height)
	for _, point := range *bezier.GetProcessed() {
		AddDotBlue(int(point.X), int(point.Y))
	}

	tagDivRocket = factoryBrowser.NewTagDiv().
		Class("animate").
		AddPointsToEasingTween(bezier).
		SetDeltaX(-25).
		SetDeltaY(-25).
		RotateDelta(-math.Pi).
		SetXY(int(1*wight+border), int(0*height+border)).
		Html("<img src=\"boat.png\" alt=\"Imagem\">")
	stage.Append(tagDivRocket)

	stage.Append(panel)

	done := make(chan struct{})
	done <- struct{}{}

}

func BezierCurve(border, wight, height float64) (bezier *algorithm.BezierCurve) {

	// E.g.: P0 (1,0) = (1*wight,0*height)
	// E.g.: P1 (2,0) = (2*wight,0*height)
	// E.g.: P2 (2,1) = (2*wight,1*height)
	//
	//     (0,0)            (1,0)            (2,0)
	//       +----------------+----------------+
	//       | P7            P0             P1 |
	//       |                                 |
	//       |                                 |
	//       |                                 |
	// (0,1) + P6                           P2 + (2,1)
	//       |                                 |
	//       |                                 |
	//       |                                 |
	//       | P5            P4             P3 |
	//       +----------------+----------------+
	//     (0,2)            (1,2)            (2,2)

	bezier = factoryAlgorithm.NewBezierCurve()
	bezier.Add(algorithm.Point{X: 1*wight + border, Y: 0*height + border})
	bezier.Add(algorithm.Point{X: 2*wight + border, Y: 0*height + border})
	bezier.Add(algorithm.Point{X: 2*wight + border, Y: 1*height + border})
	bezier.Add(algorithm.Point{X: 2*wight + border, Y: 2*height + border})
	bezier.Add(algorithm.Point{X: 1*wight + border, Y: 2*height + border})
	bezier.Add(algorithm.Point{X: 0*wight + border, Y: 2*height + border})
	bezier.Add(algorithm.Point{X: 0*wight + border, Y: 1*height + border})
	bezier.Add(algorithm.Point{X: 0*wight + border, Y: 0*height + border})
	bezier.Add(algorithm.Point{X: 1*wight + border, Y: 0*height + border})
	bezier.Process()

	return
}

func AddDotBlue(x, y int) {
	canvas.BeginPath().
		FillStyle(factoryColor.NewBlueHalfTransparent()).
		Arc(x, y, 0.5, 0, 2*math.Pi, false).
		Fill()
}
