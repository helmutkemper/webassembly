package splashScreen

import (
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/utilsText"
	"github.com/helmutkemper/webassembly/utilsWindow"
	"syscall/js"
)

func Load(stage *html.TagSvg, path string) {
	border := 50
	screenWidth, screenHeight := utilsWindow.GetScreenSize()
	screenWidth -= border
	screenHeight -= border

	svgG := factoryBrowser.NewTagSvgG()
	img := factoryBrowser.NewTagSvgImage().HRef(path)
	text := factoryBrowser.NewTagSvgText().Fill("white")
	rect := factoryBrowser.NewTagSvgRect().Stroke("red").StrokeWidth(1)
	svgG.Append(img)
	svgG.Append(text)
	svgG.Append(rect)
	stage.Append(svgG)

	img.Get().Call("addEventListener", "load", js.FuncOf(func(this js.Value, p []js.Value) interface{} {
		var widthImg, heightImg int

		bbox := img.Get().Call("getBBox")
		widthImg = bbox.Get("width").Int()
		heightImg = bbox.Get("height").Int()

		widthRatio := float64(screenWidth) / float64(widthImg)
		heightRatio := float64(screenHeight) / float64(heightImg)

		if widthRatio < heightRatio {
			widthImg = screenWidth
			heightImg = int(float64(heightImg) * widthRatio)
		} else {
			widthImg = int(float64(widthImg) * heightRatio)
			heightImg = screenHeight
		}

		x := screenWidth/2 - widthImg/2
		y := screenHeight/2 - heightImg/2

		x = x + border/2
		y = y + border/2
		img.X(x).
			Y(y).
			Width(widthImg).
			Height(heightImg)

		x = x + int(float64(widthImg)*0.2)
		y = y + int(float64(heightImg)*0.1)

		t := "Mussum Ipsum, cacilds vidis litro abertis. Copo furadis é disculpa de bebadis, arcu quam euismod magna. Quem num gosta di mim que vai caçá sua turmis! Per aumento de cachacis, eu reclamis. Suco de cevadiss deixa as pessoas mais interessantis.\n\nDiuretics paradis num copo é motivis de denguis. Interagi no mé, cursus quis, vehicula ac nisi. Segunda-feiris nun dá, eu vô me pirulitá! Paisis, filhis, espiritis santis.\n\nPraesent vel viverra nisi. Mauris aliquet nunc non turpis scelerisque, eget. Interagi no mé, cursus quis, vehicula ac nisi. Pra lá, depois divoltis porris, paradis. Si u mundo tá muito paradis? Toma um mé que o mundo vai girarzis!"

		textToBox := splitter(t, x, y, int(float64(widthImg)*0.6))

		text.X(x).
			Y(y).
			FontFamily("Verdana").
			FontSize(10).
			Width(widthImg).
			Height(heightImg).
			//TextLength(widthImg).
			TextAnchor(html.KSvgTextAnchorStart)

		for _, tspan := range textToBox {
			text.Append(tspan)
		}

		rect.X(x).Y(y).Width(int(float64(widthImg) * 0.6)).Height(int(float64(heightImg) * 0.15))

		return nil
	}))

}

func splitter(text string, x, y, widthMax int) (tSpan []*html.TagSvgTSpan) {
	tSpan = make([]*html.TagSvgTSpan, 0)
	textList := make([]string, 0)

	l := len(text)
	c := 0
	width, height := 0, 0
	for i := 1; i != l; i += 1 {
		width, height = utilsText.GetTextSize(
			text[c:i],
			"verdana",
			"normal",
			"normal",
			10,
		)
		if width >= widthMax {
			textList = append(textList, text[c:i-1])
			c = i - 1
		}
	}

	for k, splittedText := range textList {
		tSpan = append(
			tSpan,
			factoryBrowser.NewTagSvgTSpan().
				X(x).
				Y(y+(height+4)*k).
				Text(splittedText))
	}

	return
}
