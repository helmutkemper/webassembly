package _math

import (
	"fmt"
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/examples/ide/connection"
	"github.com/helmutkemper/webassembly/examples/ide/connection/factoryConnection"
	"github.com/helmutkemper/webassembly/examples/ide/ornament"
	"github.com/helmutkemper/webassembly/examples/ide/rulesConnection"
	"image/color"
	"syscall/js"
)

// OrnamentOpAmpSymbol Responsible for drawing the operational amplifier symbol used in analog electronics for mathematical
// operations
type OrnamentOpAmpSymbol struct {
	ornament.WarningMarkExclamation

	deviceBorderColor     color.RGBA
	deviceBackgroundColor color.RGBA
	deviceSymbolColor     color.RGBA

	width         int
	height        int
	deviceAdjustX int
	deviceAdjustY int

	deviceSymbolText       string
	deviceSymbolFontSize   string
	deviceSymbolFontFamily string
	deviceSymbolFontWeight string

	svg                  *html.TagSvg
	deviceBorder         *html.TagSvgPath
	deviceSymbol         *html.TagSvgText
	inputXConnection     *html.TagSvgPath
	inputXConnectionArea connection.Connection
	inputYConnection     *html.TagSvgPath
	inputYConnectionArea connection.Connection
	outputConnection     *html.TagSvgPath
	outputConnectionArea connection.Connection
}

func (e *OrnamentOpAmpSymbol) InputXSetup(setup connection.Setup) {
	e.inputXConnectionArea.Setup(setup)
}

func (e *OrnamentOpAmpSymbol) InputYSetup(setup connection.Setup) {
	e.inputYConnectionArea.Setup(setup)
}

func (e *OrnamentOpAmpSymbol) OutputSetup(setup connection.Setup) {
	e.outputConnectionArea.Setup(setup)
}

func (e *OrnamentOpAmpSymbol) GetWidth() int {
	return e.width
}

func (e *OrnamentOpAmpSymbol) GetHeight() int {
	return e.height
}

func (e *OrnamentOpAmpSymbol) ToPngResized(width, height float64) (pngData js.Value) {
	return e.svg.ToPngResized(width, height)
}

// SetWarning sets the visibility of the warning mark
func (e *OrnamentOpAmpSymbol) SetWarning(warning bool) {
	e.WarningMarkExclamation.SetWarning(warning)
}

// SetAdjustX defines the X adjustment of the symbol
func (e *OrnamentOpAmpSymbol) SetAdjustX(adjustX int) {
	e.deviceAdjustX = adjustX
}

// GetAdjustX returns the X adjustment of the symbol
func (e *OrnamentOpAmpSymbol) GetAdjustX() int {
	return e.deviceAdjustX
}

// SetAdjustY defines the Y adjustment of the symbol
func (e *OrnamentOpAmpSymbol) SetAdjustY(adjustY int) {
	e.deviceAdjustY = adjustY
}

// GetAdjustY returns the Y adjustment of the symbol
func (e *OrnamentOpAmpSymbol) GetAdjustY() int {
	return e.deviceAdjustY
}

// SetSymbol defines the symbol of the device
func (e *OrnamentOpAmpSymbol) SetSymbol(text string) {
	e.deviceSymbolText = text
	e.deviceSymbol.Text(text)
}

// GetSymbol returns the symbol of the device
func (e *OrnamentOpAmpSymbol) GetSymbol() string {
	return e.deviceSymbolText
}

// SetSymbolFontSize defines the font size of the symbol
func (e *OrnamentOpAmpSymbol) SetSymbolFontSize(fontSize string) {
	e.deviceSymbolFontSize = fontSize
	e.deviceSymbol.FontSize(fontSize)
}

// GetSymbolFontSize returns the font size of the symbol
func (e *OrnamentOpAmpSymbol) GetSymbolFontSize() string {
	return e.deviceSymbolFontSize
}

// SetSymbolFontFamily defines the font family of the symbol
func (e *OrnamentOpAmpSymbol) SetSymbolFontFamily(fontFamily string) {
	e.deviceSymbolFontFamily = fontFamily
	e.deviceSymbol.FontFamily(fontFamily)
}

// GetSymbolFontFamily returns the font family of the symbol
func (e *OrnamentOpAmpSymbol) GetSymbolFontFamily() string {
	return e.deviceSymbolFontFamily
}

// SetSymbolFontWeight defines the font weight of the symbol
func (e *OrnamentOpAmpSymbol) SetSymbolFontWeight(fontWeight string) {
	e.deviceSymbolFontWeight = fontWeight
	e.deviceSymbol.FontWeight(fontWeight)
}

// GetSymbolFontWeight returns the font weight of the symbol
func (e *OrnamentOpAmpSymbol) GetSymbolFontWeight() string {
	return e.deviceSymbolFontWeight
}

// SetBorderColor defines the color of the border
func (e *OrnamentOpAmpSymbol) SetBorderColor(color color.RGBA) {
	e.deviceBorderColor = color
	e.deviceBorder.Stroke(color)
}

// GetBorderColor returns the color of the border
func (e *OrnamentOpAmpSymbol) GetBorderColor() color.RGBA {
	return e.deviceBorderColor
}

// SetBackgroundColor defines the color of the device background
func (e *OrnamentOpAmpSymbol) SetBackgroundColor(color color.RGBA) {
	e.deviceBackgroundColor = color
	e.deviceBorder.Fill(color)
}

// GetBackgroundColor returns the color of the device background
func (e *OrnamentOpAmpSymbol) GetBackgroundColor() color.RGBA {
	return e.deviceBackgroundColor
}

// SetSymbolColor defines the color of the symbol
func (e *OrnamentOpAmpSymbol) SetSymbolColor(color color.RGBA) {
	e.deviceSymbolColor = color
	e.deviceSymbol.Fill(color)
}

// GetSymbolColor returns the color of the symbol
func (e *OrnamentOpAmpSymbol) GetSymbolColor() color.RGBA {
	return e.deviceSymbolColor
}

// Init initializes the SVG element and its content
func (e *OrnamentOpAmpSymbol) Init() (err error) {
	_ = e.WarningMarkExclamation.Init()

	e.deviceBorderColor = color.RGBA{R: 15, G: 48, B: 216, A: 255}
	e.deviceBackgroundColor = color.RGBA{R: 253, G: 255, B: 23, A: 255}
	e.deviceSymbolColor = color.RGBA{R: 83, G: 83, B: 81, A: 255}

	e.deviceSymbolText = "?"
	e.deviceSymbolFontSize = "35px"
	e.deviceSymbolFontFamily = "Arial"
	e.deviceSymbolFontWeight = "bold"

	e.svg = factoryBrowser.NewTagSvg()

	e.deviceBorder = factoryBrowser.NewTagSvgPath().
		Fill(e.deviceBackgroundColor).
		Stroke(e.deviceBorderColor).
		StrokeWidth(1).
		MarkerEnd("url(#deviceBorder)")
	e.svg.Append(e.deviceBorder)

	e.deviceSymbol = factoryBrowser.NewTagSvgText().
		Fill(e.deviceSymbolColor).
		Stroke("none").
		MarkerEnd("url(#deviceSymbol)").
		TextAnchor("middle").
		DominantBaseline("middle").
		FontSize(e.deviceSymbolFontSize).
		FontFamily(e.deviceSymbolFontFamily).
		FontWeight(e.deviceSymbolFontWeight).
		Text(e.deviceSymbolText).
		UserSelectNone()
	e.svg.Append(e.deviceSymbol)

	e.inputXConnection = factoryConnection.NewConnection("int", "url(#inputXConnection)")
	e.svg.Append(e.inputXConnection)

	e.inputXConnectionArea.Init("url(#inputXConnectionArea)")
	e.svg.Append(e.inputXConnectionArea.GetSvgPath())

	e.inputYConnection = factoryConnection.NewConnection("int", "url(#inputYConnection)")
	e.svg.Append(e.inputYConnection)

	e.inputYConnectionArea.Init("url(#inputYConnectionArea)")
	e.svg.Append(e.inputYConnectionArea.GetSvgPath())

	e.outputConnection = factoryConnection.NewConnection("int", "url(#outputConnection)")
	e.svg.Append(e.outputConnection)

	e.outputConnectionArea.Init("url(#stopButtonConnection)")
	e.svg.Append(e.outputConnectionArea.GetSvgPath())

	e.svg.Append(e.WarningMarkExclamation.GetWarningMark())
	e.SetWarning(false)
	return
}

// GetSvg Returns the SVG used as a base in the ornament
func (e *OrnamentOpAmpSymbol) GetSvg() (svg *html.TagSvg) {
	return e.svg
}

// Update Desenha o ornamento
func (e *OrnamentOpAmpSymbol) Update(x, y, width, height int) (err error) {
	e.width = width
	e.height = height

	_ = e.WarningMarkExclamation.Update(x, y, width, height)

	//e.svg.ViewBox([]int{0.0, 0.0, width, height})

	// draw the triangle
	border := 8
	device := []string{
		fmt.Sprintf("M %v %v", 0+border, 0+border),
		fmt.Sprintf("L %v %v", width-border, height/2),
		fmt.Sprintf("L %v %v", 0+border, height-border),
		fmt.Sprintf("L %v %v", 0+border, 0+border),
		"z",
	}
	e.deviceBorder.D(device)

	// calculate the center of the triangle
	a := [2]int{0 + border, 0 + border}
	b := [2]int{width - border, height / 2}
	c := [2]int{0 + border, height - border}

	// center of the triangle
	xc := (a[0] + b[0] + c[0]) / 3
	yc := (a[1] + b[1] + c[1]) / 3

	// update deviceSymbol position
	e.deviceSymbol.X(xc + e.deviceAdjustX)
	e.deviceSymbol.Y(yc + e.deviceAdjustY)

	e.inputXConnection.D(rulesConnection.GetPathDraw(2, 15))
	e.inputXConnectionArea.GetSvgPath().D(rulesConnection.GetPathAreaDraw(2, 15))
	e.inputXConnectionArea.SetXY(x+2, y+15)

	e.inputYConnection.D(rulesConnection.GetPathDraw(2, e.height-18))
	e.inputYConnectionArea.GetSvgPath().D(rulesConnection.GetPathAreaDraw(2, e.height-18))
	e.inputYConnectionArea.SetXY(x+2, y+e.height-18)

	e.outputConnection.D(rulesConnection.GetPathDraw(e.width-12, e.height/2-2))
	e.outputConnectionArea.GetSvgPath().D(rulesConnection.GetPathAreaDraw(e.width-12, e.height/2-2))
	e.outputConnectionArea.SetXY(x+e.width-12, y+e.height/2-2)

	return
}
