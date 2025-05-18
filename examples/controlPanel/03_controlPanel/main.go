package main

import (
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/platform/algorithm"
	"github.com/helmutkemper/webassembly/platform/components"
	"github.com/helmutkemper/webassembly/platform/factoryAlgorithm"
	"github.com/helmutkemper/webassembly/platform/factoryColor"
	"github.com/helmutkemper/webassembly/platform/factoryEasingTween"
	"log"
	"math"
	"syscall/js"
	"time"
)

type ComponentControlPanel struct {
	components.Components
	components.Menu

	Panel    *ControlPanel  `wasmPanel:"type:panel"`
	ContMenu *[]MenuOptions `wasmPanel:"type:contextMenu;func:InitMenu;attach:components"`
}

func (e *ComponentControlPanel) InitMenu() {
	defer log.Printf("InitMenu(): entrou!")
	e.ContMenu = &[]MenuOptions{
		{
			Type: "grid",
			Items: []MenuOptions{
				{
					Label:  "Cat 1",
					Icon:   "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
					Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} { log.Printf("cat 1"); return nil }),
				},
				{
					Label:  "Cat 2",
					Icon:   "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
					Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} { log.Printf("cat 2"); return nil }),
				},
				{
					Label:  "Cat 3",
					Icon:   "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
					Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} { log.Printf("cat 3"); return nil }),
				},
				{
					Label:  "Cat 4",
					Icon:   "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
					Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} { log.Printf("cat 4"); return nil }),
				},
				{
					Label:  "Cat 5",
					Icon:   "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
					Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} { log.Printf("cat 5"); return nil }),
				},
				{
					Label:  "Cat 6",
					Icon:   "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
					Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} { log.Printf("cat 6"); return nil }),
				},
			},
		},
		{
			Label: "-",
		},
		{
			Label: "Label 1",
			//Icon:      "icon 1",
			//IconLeft:  "icon left 1",
			//IconRight: "icon right 1",
			Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} { log.Printf("action 1 ok"); return nil }),
		},
		{
			Label: "Label 2",
			//Icon:      "icon 1",
			//IconLeft:  "icon left 1",
			//IconRight: "icon right 1",
			Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} { log.Printf("action 2 ok"); return nil }),
			Submenu: []MenuOptions{
				{
					Label: "Label 1",
					//Icon:      "icon 1",
					//IconLeft:  "icon left 1",
					//IconRight: "icon right 1",
					Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} { log.Printf("action 1 ok"); return nil }),
				},
				{
					Label: "Label 2",
					//Icon:      "icon 1",
					//IconLeft:  "icon left 1",
					//IconRight: "icon right 1",
					Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} { log.Printf("action 2 ok"); return nil }),
					Submenu: []MenuOptions{
						{
							Type: "grid",
							Items: []MenuOptions{
								{
									Label:  "Cat 1",
									Icon:   "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
									Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} { log.Printf("cat 1"); return nil }),
								},
								{
									Label:  "Cat 2",
									Icon:   "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
									Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} { log.Printf("cat 2"); return nil }),
								},
								{
									Label:  "Cat 3",
									Icon:   "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
									Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} { log.Printf("cat 3"); return nil }),
								},
								{
									Label:  "Cat 4",
									Icon:   "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
									Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} { log.Printf("cat 4"); return nil }),
								},
								{
									Label:  "Cat 5",
									Icon:   "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
									Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} { log.Printf("cat 5"); return nil }),
								},
								{
									Label:  "Cat 6",
									Icon:   "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
									Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} { log.Printf("cat 6"); return nil }),
									Submenu: []MenuOptions{
										{
											Type: "grid",
											Items: []MenuOptions{
												{
													Label:  "Cat 1",
													Icon:   "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
													Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} { log.Printf("cat 1"); return nil }),
												},
												{
													Label:  "Cat 2",
													Icon:   "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
													Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} { log.Printf("cat 2"); return nil }),
												},
												{
													Label:  "Cat 3",
													Icon:   "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
													Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} { log.Printf("cat 3"); return nil }),
												},
												{
													Label:  "Cat 4",
													Icon:   "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
													Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} { log.Printf("cat 4"); return nil }),
												},
												{
													Label:  "Cat 5",
													Icon:   "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
													Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} { log.Printf("cat 5"); return nil }),
												},
												{
													Label:  "Cat 6",
													Icon:   "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
													Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} { log.Printf("cat 6"); return nil }),
												},
											},
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
			Type: "grid",
			Items: []MenuOptions{
				{
					Label:  "Cat 1",
					Icon:   "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
					Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} { log.Printf("cat 1"); return nil }),
				},
				{
					Label:  "Cat 2",
					Icon:   "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
					Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} { log.Printf("cat 2"); return nil }),
				},
				{
					Label:  "Cat 3",
					Icon:   "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
					Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} { log.Printf("cat 3"); return nil }),
				},
				{
					Label:  "Cat 4",
					Icon:   "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
					Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} { log.Printf("cat 4"); return nil }),
				},
				{
					Label:  "Cat 5",
					Icon:   "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
					Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} { log.Printf("cat 5"); return nil }),
				},
				{
					Label:  "Cat 6",
					Icon:   "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/81_INF_DIV_SSI.jpg/50px-81_INF_DIV_SSI.jpg",
					Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} { log.Printf("cat 6"); return nil }),
				},
			},
		},
	}
}

func (e *ComponentControlPanel) Init() (panel *html.TagDiv, err error) {
	panel, err = e.Components.Init(e)
	return
}

type ControlPanel struct {
	Header string `wasmPanel:"type:headerText;label:Control panel"`
	Body   *Body  `wasmPanel:"type:panelBody"`
}

type MenuOptions struct {
	Label     string        `wasmPanel:"type:label"`
	Icon      string        `wasmPanel:"type:icon"`
	IconLeft  string        `wasmPanel:"type:iconLeft"`
	IconRight string        `wasmPanel:"type:iconRight"`
	Type      string        `wasmPanel:"type:type"`
	Items     []MenuOptions `wasmPanel:"type:options"`
	Action    js.Func       `wasmPanel:"type:action"`
	Submenu   []MenuOptions `wasmPanel:"type:subMenu"`
}

type Body struct {
	BoatAnimation *BoatAdjust `wasmPanel:"type:component;label:Easing tween time"`
}

type BoatAdjust struct {
	components.Board
	components.Menu

	Dragging *DraggingEffect `wasmPanel:"type:range;label:time (s)"`
	ContMenu *[]MenuOptions  `wasmPanel:"type:contextMenu;func:InitMenu;attach:board"`
}

func (e *BoatAdjust) InitMenu() {
	defer log.Printf("InitMenu(): entrou!")
	e.ContMenu = &[]MenuOptions{
		{
			Label: "Label 1",
			//Icon:      "icon 1",
			//IconLeft:  "icon left 1",
			//IconRight: "icon right 1",
			Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} { log.Printf("action 1 ok"); return nil }),
		},
		{
			Label: "Label 2",
			//Icon:      "icon 1",
			//IconLeft:  "icon left 1",
			//IconRight: "icon right 1",
			Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} { log.Printf("action 2 ok"); return nil }),
		},
	}
}

type DraggingEffect struct {
	components.Range

	TagRange   *html.TagInputRange  `wasmPanel:"type:inputTagRange"`
	TagNumber  *html.TagInputNumber `wasmPanel:"type:inputTagNumber"`
	Time       float64              `wasmPanel:"type:value;min:2;max:50;step:1;default:15"`
	TimeChange *OnChangeEvent       `wasmPanel:"type:listener;event:change;func:OnChangeEvent"`
}

type OnChangeEvent struct {
	IsTrusted bool    `wasmGet:"isTrusted"`
	Value     float64 `wasmGet:"value"`
	Min       float64 `wasmGet:"min"`
	Max       float64 `wasmGet:"max"`
	Type      string  `wasmGet:"type"`
}

func (e *OnChangeEvent) OnChangeEvent(event OnChangeEvent, controlPanel *ControlPanel) {
	var value = event.Value

	factoryEasingTween.NewRandom(
		time.Duration(value)*time.Second,
		0,
		1000000,
		tagDivRocket.EasingTweenWalkingAndRotateIntoPoints,
		0,
	).
		SetArgumentsFunc(any(tagDivRocket)).
		SetDoNotReverseMotion()
}

var canvas *html.TagCanvas
var tagDivRocket *html.TagDiv

func main() {
	var err error
	var panel *html.TagDiv

	stage := factoryBrowser.NewStage()

	controlPanel := new(ComponentControlPanel)
	//controlPanel.Panel = new(ControlPanel)
	//controlPanel.Panel.Body = new(Body)
	//controlPanel.Panel.Body.ContMenu = &[]MenuOptions{
	//	{
	//		Label:     "Label 1",
	//		Icon:      "icon 1",
	//		IconLeft:  "icon left 1",
	//		IconRight: "icon right 1",
	//		Type:      "type 1",
	//	},
	//}
	if panel, err = controlPanel.Init(); err != nil {
		panic(err)
	}

	//log.Printf("controlPanel.Panel: %v", controlPanel.Panel)
	//log.Printf("controlPanel.Panel.Body: %v", controlPanel.Panel.Body)
	//log.Printf("controlPanel.Panel.Body.ContMenu: %v", controlPanel.Panel.Body.ContMenu)
	//log.Printf("controlPanel.Panel.Header: %v", controlPanel.Panel.Header)

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
