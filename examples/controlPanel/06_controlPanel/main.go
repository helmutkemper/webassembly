package main

import (
	"encoding/json"
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/platform/algorithm"
	"github.com/helmutkemper/webassembly/platform/components"
	"github.com/helmutkemper/webassembly/platform/easingTween"
	"github.com/helmutkemper/webassembly/platform/factoryAlgorithm"
	"github.com/helmutkemper/webassembly/platform/factoryColor"
	"log"
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
	SimpleForm    *SimpleForm `wasmPanel:"type:component;label:simple form"`
	Share         *ShareForm  `wasmPanel:"type:component;label:Share"`
}

type DraggingEffect struct {
	components.Range

	TagRange     *html.TagInputRange  `wasmPanel:"type:inputTagRange"`
	TagNumber    *html.TagInputNumber `wasmPanel:"type:inputTagNumber"`
	Dragging     float64              `wasmPanel:"type:value;min:2;max:50;step:1;default:42"`
	NumberChange *OnChangeEvent       `wasmPanel:"type:listener;event:change;func:OnChange"`
	RangeChange  *OnChangeEvent       `wasmPanel:"type:listener;event:input;func:OnInputEvent"`
}

type OnChangeEvent struct {
	IsTrusted bool    `wasmGet:"isTrusted"`
	Value     float64 `wasmGet:"value"`
	Min       float64 `wasmGet:"min"`
	Max       float64 `wasmGet:"max"`
	Type      string  `wasmGet:"type"`
}

func (e *OnChangeEvent) OnChange(event OnChangeEvent, reference *ControlPanel) {
	//ref := reference.BoatAnimation.Dragging

	//log.Printf("reference.TagNumber.GetValue(): %v", ref.TagNumber.GetValue())
	//log.Printf("reference.TagRange.GetValue(): %v", ref.TagRange.GetValue())
}

func (e *OnChangeEvent) OnInputEvent(event OnChangeEvent, reference *ControlPanel) {
	ref := reference.Body.BoatAnimation.Dragging
	switch event.Type {
	case "range":
		ref.TagNumber.Value(ref.RangeCalcFormula(event.Min, event.Max, ref.TagRange.GetValue()))
	case "number":
		ref.TagRange.Value(ref.RangeCalcFormula(event.Min, event.Max, ref.TagNumber.GetValue()))
	}
}

func (e *DraggingEffect) RangeCalcFormula(min, max, value float64) (result float64) {
	return (max - value) + min
}

func (e *DraggingEffect) Init() {
	e.TagNumber.Value(e.RangeCalcFormula(2, 50, e.TagRange.GetValue()))
	e.TagRange.Value(e.RangeCalcFormula(2, 50, e.TagNumber.GetValue()))
}

type SimpleForm struct {
	Text     *TextForm     `wasmPanel:"type:text;label:Text"`
	Url      *UrlForm      `wasmPanel:"type:url;label:Url"`
	Tel      *TelForm      `wasmPanel:"type:tel;label:Telephone"`
	Time     *TimeForm     `wasmPanel:"type:time;label:Time"`
	Month    *MonthForm    `wasmPanel:"type:month;label:Month"`
	Week     *WeekForm     `wasmPanel:"type:week;label:Week"`
	Date     *DateForm     `wasmPanel:"type:date;label:Date"`
	Color    *ColorForm    `wasmPanel:"type:color;label:Color"`
	Password *PasswordForm `wasmPanel:"type:password;label:Password"`
	Mail     *MailForm     `wasmPanel:"type:mail;label:E-Mail"`
	TextArea *TextAreaForm `wasmPanel:"type:textArea;label:Text"`
	Radio    *ListRadio    `wasmPanel:"type:radio;label:Select one"`
	Checkbox *ListCheckbox `wasmPanel:"type:checkbox;label:Select all"`
}

type ShareForm struct {
	QRCode *QRCodeForm `wasmPanel:"type:qrcode;label:QR Code"`
}

type BoatAdjust struct {
	Dragging *DraggingEffect   `wasmPanel:"type:range;label:effect"`
	Tween    *TweenSelect      `wasmPanel:"type:select;label:Tween function"`
	Start    *EasingTweenStart `wasmPanel:"type:button;label:start easing tween"`
}

type TextAreaForm struct {
	components.TextArea

	TextTag *html.TagTextArea `wasmPanel:"type:inputTagTextArea"`
	Value   string            `wasmPanel:"type:value;placeHolder:Digite um texto;default:default text"`
	Change  *OnTextAreaEvent  `wasmPanel:"type:listener;event:input;func:OnChangeEvent"`
}

type OnTextAreaEvent struct {
	Value string `wasmGet:"value"`
}

func (e *OnTextAreaEvent) OnChangeEvent(event OnTextAreaEvent, reference *ControlPanel) {
	log.Printf("text: %v", event.Value)
}

type TextForm struct {
	components.Text

	TextTag *html.TagInputText `wasmPanel:"type:inputTagText"`
	Value   string             `wasmPanel:"type:value;placeHolder:Digite um texto;default:default text"`
	Change  *OnTextEvent       `wasmPanel:"type:listener;event:change;func:OnChangeEvent"`
}

type OnTextEvent struct {
	Value string `wasmGet:"value"`
}

func (e *OnTextEvent) OnChangeEvent(event OnTextEvent, reference *ControlPanel) {
	log.Printf("text: %v", event.Value)
}

type QRCodeForm struct {
	components.QRCode

	TextTag       *html.TagCanvas `wasmPanel:"type:TagCanvas"`
	QRCodeValue   string          `wasmPanel:"type:value;size:309;level:4;color:#000000;background:#00ff00;default:Hello Word!"`
	RecoveryLevel int             `wasmPanel:"type:level"`
	Color         string          `wasmPanel:"type:color"`
	Background    string          `wasmPanel:"type:background"`
	DisableBorder bool            `wasmPanel:"type:disableBorder"`
	Change        *OnQRCodeEvent  `wasmPanel:"type:listener;event:click;func:OnChangeEvent"`
}

type OnQRCodeEvent struct {
	//Value string `wasmGet:"value"`
}

func (e *OnQRCodeEvent) OnChangeEvent(event OnQRCodeEvent, reference *ControlPanel) {
	ref := reference.Body.Share

	ref.QRCode.SetColor("#ff00ff")
	ref.QRCode.SetBackground("#ffff00")
	ref.QRCode.SetValue(time.Now().Format(time.TimeOnly))
}

type UrlForm struct {
	components.Url

	TextTag  *html.TagInputUrl `wasmPanel:"type:inputTagUrl"`
	UrlValue string            `wasmPanel:"type:value;placeHolder:Digite a URL;default:'https://www.google.com'"`
	Change   *OnUrlEvent       `wasmPanel:"type:listener;event:change;func:OnChangeEvent"`
}

type OnUrlEvent struct {
	Value string `wasmGet:"value"`
}

func (e *OnUrlEvent) OnChangeEvent(event OnUrlEvent, reference *ControlPanel) {
	log.Printf("Url: %v", event.Value)
}

type TelForm struct {
	components.Tel

	TelTag   *html.TagInputTel `wasmPanel:"type:inputTagTel"`
	TelValue string            `wasmPanel:"type:value;placeHolder:Digite o telefone;default:(99)9999999999"`
	Change   *OnTelEvent       `wasmPanel:"type:listener;event:change;func:OnChangeEvent"`
}

func (e *TelForm) Init() {
	e.Value("(88)888888888")
}

type OnTelEvent struct {
	Value string `wasmGet:"value"`
}

func (e *OnTelEvent) OnChangeEvent(event OnTelEvent, reference *ControlPanel) {
	log.Printf("Tel: %v", event.Value)
}

type MailForm struct {
	components.Mail

	TextTag   *html.TagInputMail `wasmPanel:"type:inputTagMail"`
	MailValue string             `wasmPanel:"type:value;placeHolder:Digite um e-mail;default:eu@eu.eu"`
	Change    *OnMailEvent       `wasmPanel:"type:listener;event:change;func:OnChangeEvent"`
}

func (e *MailForm) Init() {
	e.Value("tu@eu")
}

type OnMailEvent struct {
	Value string `wasmGet:"value"`
}

func (e *OnMailEvent) OnChangeEvent(event OnMailEvent, reference *ControlPanel) {
	log.Printf("mail: %v", event.Value)
}

type TimeForm struct {
	components.Time

	TextTag   *html.TagInputTime `wasmPanel:"type:inputTagTime"`
	TimeValue string             `wasmPanel:"type:value;default:'12:30'"`
	Change    *OnTimeEvent       `wasmPanel:"type:listener;event:change;func:OnChangeEvent"`
}

func (e *TimeForm) Init() {
	e.Value("22:59")
}

type OnTimeEvent struct {
	Value string `wasmGet:"value"`
}

func (e *OnTimeEvent) OnChangeEvent(event OnTimeEvent, reference *ControlPanel) {
	log.Printf("text: %v", event.Value)
}

type MonthForm struct {
	components.Month

	TextTag    *html.TagInputMonth `wasmPanel:"type:inputTagMonth"`
	MonthValue string              `wasmPanel:"type:value;default:2024-09"`
	Change     *OnMonthEvent       `wasmPanel:"type:listener;event:change;func:OnChangeEvent"`
}

func (e *MonthForm) Init() {
	date, _ := time.Parse(time.DateOnly, "1973-10-19")
	e.Value(date)
}

type OnMonthEvent struct {
	Value string `wasmGet:"value"`
}

func (e *OnMonthEvent) OnChangeEvent(event OnMonthEvent, reference *ControlPanel) {
	log.Printf("text: %v", event.Value)
}

type DateForm struct {
	components.Date

	DateTag   *html.TagInputDate `wasmPanel:"type:inputTagDate"`
	DateValue string             `wasmPanel:"type:value;default:2017-06-01"`
	Change    *OnDateEvent       `wasmPanel:"type:listener;event:change;func:OnChangeEvent"`
}

func (e *DateForm) Init() {
	e.Value("1973-10-19")
}

type OnDateEvent struct {
	Value string `wasmGet:"value"`
}

func (e *OnDateEvent) OnChangeEvent(event OnDateEvent, reference *ControlPanel) {
	log.Printf("date: %v", event.Value)
}

type WeekForm struct {
	components.Week

	DateTag   *html.TagInputWeek `wasmPanel:"type:inputTagWeek"`
	WeekValue string             `wasmPanel:"type:value;default:2017-W45"`
	Change    *OnWeekEvent       `wasmPanel:"type:listener;event:change;func:OnChangeEvent"`
}

func (e *WeekForm) Init() {
	e.Value(time.Date(1973, time.October, 19, 0, 0, 0, 0, time.UTC))
}

type OnWeekEvent struct {
	Value string `wasmGet:"value"`
}

func (e *OnWeekEvent) OnChangeEvent(event OnWeekEvent, reference *ControlPanel) {
	log.Printf("week: %v", event.Value)
}

type ColorForm struct {
	components.Color

	ColorTag   *html.TagInputColor `wasmPanel:"type:inputTagColor"`
	ColorValue string              `wasmPanel:"type:value;default:#ff0000"`
	Change     *OnColorEvent       `wasmPanel:"type:listener;event:change;func:OnChangeEvent"`
}

func (e *ColorForm) Init() {

	e.Value("#ffffff")
}

type OnColorEvent struct {
	Value string `wasmGet:"value"`
}

func (e *OnColorEvent) OnChangeEvent(event OnColorEvent, reference *ControlPanel) {
	log.Printf("text: %v", event.Value)
}

type PasswordForm struct {
	components.Password

	PassTag *html.TagInputPassword `wasmPanel:"type:inputTagPassword"`
	Value   string                 `wasmPanel:"type:value;placeHolder:Digite uma senha;default:senha"`
	Change  *OnPasswordEvent       `wasmPanel:"type:listener;event:change;func:OnChangeEvent"`
}

type OnPasswordEvent struct {
	Value string `wasmGet:"value"`
}

func (e *OnPasswordEvent) OnChangeEvent(event OnPasswordEvent, reference *ControlPanel) {
	log.Printf("text: %v", event.Value)
}

type ListRadio struct {
	components.Radio

	List *[]RadioType `wasmPanel:"type:value;name:radio;default:label 1,value 1,>label 2,value 2,label 3,value 3"` //;default:label 1,value 1,>label 2,value 2,label 3,value 3
}

func (e *ListRadio) Init() {
	data, _ := json.Marshal(&e.List)
	log.Printf("%s", data)

	//(*e.List)[0].TagRadio.Checked(true)
	//(*e.List)[0].TagLabel.Text("Vivo! >> ").Append((*e.List)[0].TagRadio)
}

type RadioType struct {
	TagRadio *html.TagInputRadio `wasmPanel:"type:inputTagRadio"`
	TagLabel *html.TagLabel      `wasmPanel:"type:inputTagLabel"`
	Label    string              `wasmPanel:"type:label"`
	Value    string              `wasmPanel:"type:value"`
	Disabled bool                `wasmPanel:"type:disabled"`
	Selected bool                `wasmPanel:"type:selected"`
	Change   *RadioChange        `wasmPanel:"type:listener;event:change;func:OnChangeEvent"`
}

type RadioChange struct {
	Value string `wasmGet:"value"`
}

func (e *RadioChange) OnChangeEvent(event RadioChange, reference *ControlPanel) {
	log.Printf("value: %v", event.Value)
}

type ListCheckbox struct {
	components.Checkbox

	List *[]CheckboxType `wasmPanel:"type:value;name:radio;default:label 1,value 1,>label 2,value 2,label 3,value 3"` //;default:label 1,value 1,>label 2,value 2,label 3,value 3
}

func (e *ListCheckbox) Init() {
	data, _ := json.Marshal(&e.List)
	log.Printf("%s", data)

	(*e.List)[0].TagRadio.Checked(true)
	(*e.List)[0].TagLabel.Text("Vivo! >> ").Append((*e.List)[0].TagRadio)
}

type CheckboxType struct {
	TagRadio *html.TagInputCheckBox `wasmPanel:"type:inputTagCheckbox"`
	TagLabel *html.TagLabel         `wasmPanel:"type:inputTagLabel"`
	Label    string                 `wasmPanel:"type:label"`
	Value    string                 `wasmPanel:"type:value"`
	Disabled bool                   `wasmPanel:"type:disabled"`
	Selected bool                   `wasmPanel:"type:selected"`
	Change   *CheckboxChange        `wasmPanel:"type:listener;event:change;func:OnChangeEvent"`
}

type CheckboxChange struct {
	Value   string `wasmGet:"value"`
	Checked bool   `wasmGet:"checked"`
}

func (e *CheckboxChange) OnChangeEvent(event CheckboxChange, reference *ControlPanel) {
	log.Printf("value: %v:%v", event.Value, event.Checked)
}

type TweenSelect struct {
	components.Select

	SelectTag *html.TagSelect `wasmPanel:"type:inputTagSelect"`
	List      *[]TweenType    `wasmPanel:"type:value"` //;default:label 1,value 1,>label 2,value 2,label 3,value 3
	Change    *OnSelectChange `wasmPanel:"type:listener;event:change;func:OnChangeEvent"`
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

func (e *OnSelectChange) OnChangeEvent(event OnSelectChange, reference *ControlPanel) {
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

	tween *easingTween.Tween
}

func (e *OnClickEvent) OnClickEvent(event OnClickEvent, reference *ControlPanel) {
	ref := reference.Body.BoatAnimation
	//log.Printf("Trusted: %v", event.IsTrusted)
	//log.Printf("Value:   %v", event.Value)

	var value = ref.Dragging.TagNumber.GetValue()

	ref.Start.Value("Stop")
	if e.tween != nil {
		e.tween.End()

		return
	}

	e.tween = new(easingTween.Tween)
	e.tween.SetDuration(time.Duration(value)*time.Second).
		SetValues(0, 1000000).
		SetOnStepFunc(tagDivRocket.EasingTweenWalkingAndRotateIntoPoints()).
		SetLoops(0).
		SetArgumentsFunc(any(tagDivRocket)).
		SetTweenFunc(ref.Tween.Change.function).
		SetDoNotReverseMotion().
		//todo: criar uma função onTermination
		SetOnEndFunc(func(_ float64, _ interface{}) {
			e.tween = nil
			ref.Start.Value("Restart")
		}).
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
		Panel: &ControlPanel{
			Body: &Body{
				SimpleForm: &SimpleForm{
					Checkbox: &ListCheckbox{
						List: &[]CheckboxType{
							{
								TagRadio: factoryBrowser.NewTagInputCheckBox(),
								TagLabel: factoryBrowser.NewTagLabel(),
								Label:    "label1",
								Value:    "Value_1",
								Disabled: false,
								Selected: false,
							},
							{
								TagRadio: factoryBrowser.NewTagInputCheckBox(),
								TagLabel: factoryBrowser.NewTagLabel(),
								Label:    "label2",
								Value:    "Value_2",
								Disabled: false,
								Selected: true,
							},
						},
					},
				},
				Share: &ShareForm{
					QRCode: &QRCodeForm{
						QRCodeValue:   "frankenstein",
						RecoveryLevel: 1,
						Color:         "#ffffff",
						Background:    "#000000",
						DisableBorder: true,
					},
				},
			},
		},
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
