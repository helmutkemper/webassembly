package splashScreen

import (
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/platform/easingTween"
	"github.com/helmutkemper/webassembly/utilsText"
	"github.com/helmutkemper/webassembly/utilsWindow"
	"image/color"
	"reflect"
	"sync"
	"syscall/js"
	"time"
)

type textBox struct {
	x, y, width, height float64
}

type Control struct {
	textLine      []string
	lineHeight    int
	svgGroup      *html.TagSvgG
	svgBlur       *html.TagSvgRect
	svgImage      *html.TagSvgImage
	svgText       *html.TagSvgText
	border        int
	textColor     any
	fontFamily    string
	fontSize      int
	textBoxRatio  textBox
	textBoxPixels textBox
	textPadding   int
	fontWeight    html.FontWeightRule
	fontStyle     html.FontStyleRule
	path          string
	stage         *html.TagSvg
}

func (e *Control) Show() {
	// todo: controle de z-index
	e.stage.Get().Call("appendChild", e.svgGroup.Get())
}
func (e *Control) Hide() {

	wg := new(sync.WaitGroup)
	wg.Add(1)
	new(easingTween.Tween).
		SetDuration(1 * time.Second).
		SetOnStepFunc(func(value, percentToComplete float64, arguments interface{}) {
			e.svgGroup.AddStyle("opacity", 1.0-percentToComplete)
		}).
		SetLoops(0).
		SetTweenFunc(easingTween.KLinear).
		SetDoNotReverseMotion().
		//todo: criar uma função onTermination
		SetOnEndFunc(func(_ float64, _ interface{}) {
			e.Clear()
			e.svgGroup.AddStyle("opacity", 1)
			e.stage.Get().Call("removeChild", e.svgGroup.Get())
			wg.Done()
		}).
		Start()
	wg.Wait()
}

func (e *Control) Clear() {
	e.svgText.Html("")
	e.textLine = make([]string, 0)
}

func (e *Control) AddText(text string) {
	textToBox := e.splitter(text)

	e.svgText.X(e.textBoxPixels.x).
		Y(e.textBoxPixels.y).
		FontFamily(e.fontFamily).
		FontSize(e.fontSize).
		TextAnchor(html.KSvgTextAnchorStart).
		Html("")

	for _, tspan := range textToBox {
		e.svgText.Append(tspan)
	}

	time.Sleep(300 * time.Millisecond)
}

func (e *Control) Init(stage *html.TagSvg) {
	// todo: controle de z-index
	e.stage = stage
	e.textLine = make([]string, 0)

	if e.border == 0 {
		e.border = 50
	}

	screenWidth, screenHeight := utilsWindow.GetScreenSize()
	screenWidth -= e.border
	screenHeight -= e.border

	if e.path == "" {
		e.path = "./splashScreen/splashScreen.png"
	}

	if reflect.DeepEqual(e.textBoxRatio, textBox{}) {
		e.textBoxRatio = textBox{
			x:      0.2,
			y:      0.1,
			width:  0.6,
			height: 0.15,
		}
	}

	if e.textColor == nil {
		e.textColor = "white"
	}

	if e.textPadding == 0 {
		e.textPadding = 0
	}

	if e.fontWeight == "" {
		e.fontWeight = html.KFontWeightRuleNormal
	}

	if e.fontStyle == "" {
		e.fontStyle = html.KFontStyleRuleNormal
	}

	if e.fontFamily == "" {
		e.fontFamily = "Verdana"
	}

	if e.fontSize == 0 {
		e.fontSize = 20
	}

	_, e.lineHeight = utilsText.GetTextSize(
		"aAbBcCçÇdDeEfFgGhHiIjJkKlLmMnNoOpPqQrRsStTuUvVwWxXyYzZ0123456789",
		e.fontFamily,
		e.fontWeight.String(),
		e.fontStyle.String(),
		e.fontSize,
	)

	e.svgGroup = factoryBrowser.NewTagSvgG().AddStyle("zIndex", "1001") // todo: controle de z=index
	e.svgBlur = factoryBrowser.NewTagSvgRect().AddStyle("zIndex", "1000").X(0).Y(0).Width(screenWidth + e.border).Height(screenHeight + e.border).Fill(color.RGBA{R: 0, G: 0, B: 0, A: 220})
	e.svgImage = factoryBrowser.NewTagSvgImage().HRef(e.path)
	e.svgText = factoryBrowser.NewTagSvgText().Fill(e.textColor)
	e.svgGroup.Append(e.svgBlur)
	e.svgGroup.Append(e.svgImage)
	e.svgGroup.Append(e.svgText)

	// CAUTION: For getBBox to work, the element must be visible in the DOM
	e.stage.Append(e.svgGroup)

	wg := new(sync.WaitGroup)
	wg.Add(1)
	e.svgImage.Get().Call("addEventListener", "load", js.FuncOf(func(this js.Value, p []js.Value) interface{} {
		var widthImg, heightImg int

		// CAUTION: For getBBox to work, the element must be visible in the DOM
		bbox := e.svgImage.Get().Call("getBBox")
		widthImg = bbox.Get("width").Int()
		heightImg = bbox.Get("height").Int()

		// Proportion of the screen size in relation to the size of the image
		widthRatio := float64(screenWidth) / float64(widthImg)
		heightRatio := float64(screenHeight) / float64(heightImg)

		// Choose the smallest size to adjust the image size
		if widthRatio < heightRatio {
			widthImg = screenWidth
			heightImg = int(float64(heightImg) * widthRatio)
		} else {
			widthImg = int(float64(widthImg) * heightRatio)
			heightImg = screenHeight
		}

		// Centralize the image
		x := screenWidth/2 - widthImg/2
		y := screenHeight/2 - heightImg/2

		// Adjusts the edge of the screen
		x = x + e.border/2
		y = y + e.border/2
		e.svgImage.X(x).
			Y(y).
			Width(widthImg).
			Height(heightImg)

		// Determines the size of the text box
		e.textBoxPixels.x = float64(x + int(float64(widthImg)*e.textBoxRatio.x))
		e.textBoxPixels.y = float64(y + int(float64(heightImg)*e.textBoxRatio.y))
		e.textBoxPixels.width = float64(widthImg) * e.textBoxRatio.width
		e.textBoxPixels.height = float64(heightImg) * e.textBoxRatio.height

		wg.Done()
		return nil
	}))

	wg.Wait()
}

func (e *Control) removesExcessLines() {
	heightActual := 0
	for range e.textLine {
		heightActual += e.lineHeight + e.textPadding
		if heightActual > int(e.textBoxPixels.height) {
			e.textLine = e.textLine[:len(e.textLine)-1]
			e.removesExcessLines()
			return
		}
	}
}

func (e *Control) splitter(text string) (tSpan []*html.TagSvgTSpan) {
	tSpan = make([]*html.TagSvgTSpan, 0)

	// breaks the text to be added in lines that fit in the particular space
	textLength := len(text)
	textStart := 0
	width := 0

	tmpLines := make([]string, 0)
	i := 0
	for i = 1; i != textLength; i += 1 {
		width, _ = utilsText.GetTextSize(
			text[textStart:i],
			e.fontFamily,
			e.fontWeight.String(),
			e.fontStyle.String(),
			e.fontSize,
		)
		if width >= int(e.textBoxPixels.width) {
			tmpLines = append(tmpLines, text[textStart:i-1])
			textStart = i - 1
		}
	}

	tmpLines = append(tmpLines, text[textStart:])

	// Add the new lines in front of the previous lines, in order
	e.textLine = append(tmpLines, e.textLine...)
	//log.Printf("adicionado: e.textLine: %+v", e.textLine)

	e.removesExcessLines()

	//log.Printf("removido: e.textLine: %+v", e.textLine)
	//log.Print("")
	//log.Print("")

	for textKey, splittedText := range e.textLine {
		y := e.textBoxPixels.y + float64(e.fontSize*textKey)

		tSpan = append(
			tSpan,
			factoryBrowser.NewTagSvgTSpan().
				X(e.textBoxPixels.x).
				Y(y).
				Text(splittedText))
	}

	return
}
