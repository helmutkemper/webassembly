package math

import (
	"fmt"
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/examples/ide/connection"
	"github.com/helmutkemper/webassembly/examples/ide/connection/factoryConnection"
	"github.com/helmutkemper/webassembly/examples/ide/rulesConnection"
	"github.com/helmutkemper/webassembly/examples/ide/rulesDensity"
	"image/color"
	"syscall/js"
)

// OrnamentOpAmpSymbol Responsible for drawing the operational amplifier symbol used in analog electronics for mathematical
// operations
type OrnamentOpAmpSymbol struct {
	//ornament.WarningMarkExclamation

	deviceBorderNormalColor     color.RGBA
	deviceBackgroundNormalColor color.RGBA
	deviceSymbolNormalColor     color.RGBA

	deviceBorderSelectedColor     color.RGBA
	deviceBackgroundSelectedColor color.RGBA
	deviceSymbolSelectedColor     color.RGBA

	width         rulesDensity.Density
	height        rulesDensity.Density
	deviceAdjustX rulesDensity.Density
	deviceAdjustY rulesDensity.Density

	deviceSymbolText       string
	deviceSymbolFontSize   rulesDensity.Density
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

func (e *OrnamentOpAmpSymbol) GetWidth() rulesDensity.Density {
	return e.width
}

func (e *OrnamentOpAmpSymbol) GetHeight() rulesDensity.Density {
	return e.height
}

func (e *OrnamentOpAmpSymbol) ToPngResized(width, height float64) (pngData js.Value) {
	return e.svg.ToPngResized(width, height)
}

func (e *OrnamentOpAmpSymbol) SetSelected(selected bool) {
	if selected {
		e.deviceBorder.Fill(e.deviceBackgroundSelectedColor)
		e.deviceBorder.Stroke(e.deviceBorderSelectedColor)
		e.deviceSymbol.Fill(e.deviceSymbolSelectedColor)
		return
	}

	e.deviceBorder.Fill(e.deviceBackgroundNormalColor)
	e.deviceBorder.Stroke(e.deviceBorderNormalColor)
	e.deviceSymbol.Fill(e.deviceSymbolNormalColor)
}

// SetWarning sets the visibility of the warning mark
func (e *OrnamentOpAmpSymbol) SetWarning(warning bool) {
	//e.WarningMarkExclamation.SetWarning(warning) // todo: fazer
}

// SetAdjustX defines the X adjustment of the symbol
func (e *OrnamentOpAmpSymbol) SetAdjustX(adjustX rulesDensity.Density) {
	e.deviceAdjustX = adjustX
}

// GetAdjustX returns the X adjustment of the symbol
func (e *OrnamentOpAmpSymbol) GetAdjustX() rulesDensity.Density {
	return e.deviceAdjustX
}

// SetAdjustY defines the Y adjustment of the symbol
func (e *OrnamentOpAmpSymbol) SetAdjustY(adjustY rulesDensity.Density) {
	e.deviceAdjustY = adjustY
}

// GetAdjustY returns the Y adjustment of the symbol
func (e *OrnamentOpAmpSymbol) GetAdjustY() rulesDensity.Density {
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
func (e *OrnamentOpAmpSymbol) SetSymbolFontSize(fontSize rulesDensity.Density) {
	e.deviceSymbolFontSize = fontSize
	e.deviceSymbol.FontSize(fontSize.Pixel())
}

// GetSymbolFontSize returns the font size of the symbol
func (e *OrnamentOpAmpSymbol) GetSymbolFontSize() rulesDensity.Density {
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

// SetBorderNormalColor defines the color of the border
func (e *OrnamentOpAmpSymbol) SetBorderNormalColor(color color.RGBA) {
	e.deviceBorderNormalColor = color
	e.deviceBorder.Stroke(color)
}

// SetBorderSelectedColor defines the color of the border
func (e *OrnamentOpAmpSymbol) SetBorderSelectedColor(color color.RGBA) {
	e.deviceBorderSelectedColor = color
	e.deviceBorder.Stroke(color)
}

// GetBorderNormalColor returns the color of the border
func (e *OrnamentOpAmpSymbol) GetBorderNormalColor() color.RGBA {
	return e.deviceBorderNormalColor
}

// GetBorderSelectedColor returns the color of the border
func (e *OrnamentOpAmpSymbol) GetBorderSelectedColor() color.RGBA {
	return e.deviceBorderSelectedColor
}

// SetBackgroundNormalColor defines the color of the device background
func (e *OrnamentOpAmpSymbol) SetBackgroundNormalColor(color color.RGBA) {
	e.deviceBackgroundNormalColor = color
	e.deviceBorder.Fill(color)
}

// SetBackgroundSelectedColor defines the color of the device background
func (e *OrnamentOpAmpSymbol) SetBackgroundSelectedColor(color color.RGBA) {
	e.deviceBackgroundSelectedColor = color
	e.deviceBorder.Fill(color)
}

// GetBackgroundNormalColor returns the color of the device background
func (e *OrnamentOpAmpSymbol) GetBackgroundNormalColor() color.RGBA {
	return e.deviceBackgroundNormalColor
}

// GetBackgroundSelectedColor returns the color of the device background
func (e *OrnamentOpAmpSymbol) GetBackgroundSelectedColor() color.RGBA {
	return e.deviceBackgroundSelectedColor
}

// SetSymbolNormalColor defines the color of the symbol
func (e *OrnamentOpAmpSymbol) SetSymbolNormalColor(color color.RGBA) {
	e.deviceSymbolNormalColor = color
	e.deviceSymbol.Fill(color)
}

// SetSymbolSelectedColor defines the color of the symbol
func (e *OrnamentOpAmpSymbol) SetSymbolSelectedColor(color color.RGBA) {
	e.deviceSymbolSelectedColor = color
	e.deviceSymbol.Fill(color)
}

// GetSymbolNormalColor returns the color of the symbol
func (e *OrnamentOpAmpSymbol) GetSymbolNormalColor() color.RGBA {
	return e.deviceSymbolNormalColor
}

// GetSymbolSelectedColor returns the color of the symbol
func (e *OrnamentOpAmpSymbol) GetSymbolSelectedColor() color.RGBA {
	return e.deviceSymbolSelectedColor
}

// Init Initializes the SVG element and its content
func (e *OrnamentOpAmpSymbol) Init() (err error) {
	//_ = e.WarningMarkExclamation.Init()

	e.deviceBorderNormalColor = color.RGBA{R: 15, G: 48, B: 216, A: 255}
	e.deviceBackgroundNormalColor = color.RGBA{R: 253, G: 255, B: 23, A: 255}
	e.deviceSymbolNormalColor = color.RGBA{R: 83, G: 83, B: 81, A: 255}

	e.deviceBorderSelectedColor = color.RGBA{R: 65, G: 48, B: 216, A: 255}
	e.deviceBackgroundSelectedColor = color.RGBA{R: 253, G: 205, B: 0, A: 255}
	e.deviceSymbolSelectedColor = color.RGBA{R: 133, G: 83, B: 81, A: 255}

	e.deviceSymbolText = "?"
	e.deviceSymbolFontSize = rulesDensity.Density(35)
	e.deviceSymbolFontFamily = "Arial"
	e.deviceSymbolFontWeight = "bold"

	e.svg = factoryBrowser.NewTagSvg()

	e.deviceBorder = factoryBrowser.NewTagSvgPath().
		Fill(e.deviceBackgroundNormalColor).
		Stroke(e.deviceBorderNormalColor).
		StrokeWidth(rulesDensity.Density(1).GetInt()).
		MarkerEnd("url(#deviceBorder)")
	e.svg.Append(e.deviceBorder)

	e.deviceSymbol = factoryBrowser.NewTagSvgText().
		Fill(e.deviceSymbolNormalColor).
		Stroke("none").
		MarkerEnd("url(#deviceSymbol)").
		TextAnchor("middle").
		DominantBaseline("middle").
		FontSize(e.deviceSymbolFontSize.Pixel()).
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

	//e.svg.Append(e.WarningMarkExclamation.GetSvg())
	e.SetWarning(false)
	return
}

// GetSvg Returns the SVG used as a base in the ornament
func (e *OrnamentOpAmpSymbol) GetSvg() (svg *html.TagSvg) {
	return e.svg
}

// Update Desenha o ornamento
func (e *OrnamentOpAmpSymbol) Update(x, y, width, height rulesDensity.Density) (err error) {
	e.width = width
	e.height = height

	//_ = e.WarningMarkExclamation.Update(x, y, width, height)

	//e.svg.ViewBox([]int{0.0, 0.0, width, height})

	// draw the triangle
	border := rulesDensity.Density(8)
	device := []string{
		fmt.Sprintf("M %v %v", 0+border, 0+border),
		fmt.Sprintf("L %v %v", width-border, height/2),
		fmt.Sprintf("L %v %v", 0+border, height-border),
		fmt.Sprintf("L %v %v", 0+border, 0+border),
		"z",
	}
	e.deviceBorder.D(device)

	// calculate the center of the triangle
	a := [2]int{0 + border.GetInt(), 0 + border.GetInt()}
	b := [2]int{width.GetInt() - border.GetInt(), height.GetInt() / 2}
	c := [2]int{0 + border.GetInt(), height.GetInt() - border.GetInt()}

	// center of the triangle
	xc := (a[0] + b[0] + c[0]) / 3
	yc := (a[1] + b[1] + c[1]) / 3

	// update deviceSymbol position
	e.deviceSymbol.X(rulesDensity.Convert(xc).GetInt() + e.deviceAdjustX.GetInt())
	e.deviceSymbol.Y(rulesDensity.Convert(yc).GetInt() + e.deviceAdjustY.GetInt())

	e.inputXConnection.D(rulesConnection.GetPathDraw(rulesDensity.Density(2), rulesDensity.Density(15)))
	e.inputXConnectionArea.GetSvgPath().D(rulesConnection.GetPathAreaDraw(rulesDensity.Density(2), rulesDensity.Density(15)))
	e.inputXConnectionArea.SetXY(x+rulesDensity.Density(2), y+rulesDensity.Density(15))

	e.inputYConnection.D(rulesConnection.GetPathDraw(rulesDensity.Density(2), e.height-rulesDensity.Density(18)))
	e.inputYConnectionArea.GetSvgPath().D(rulesConnection.GetPathAreaDraw(rulesDensity.Density(2), e.height-rulesDensity.Density(18)))
	e.inputYConnectionArea.SetXY(x+rulesDensity.Density(2), y+e.height-rulesDensity.Density(18))

	e.outputConnection.D(rulesConnection.GetPathDraw(e.width-rulesDensity.Density(12), e.height/2-rulesDensity.Density(2)))
	e.outputConnectionArea.GetSvgPath().D(rulesConnection.GetPathAreaDraw(e.width-rulesDensity.Density(12), e.height/2-rulesDensity.Density(2)))
	e.outputConnectionArea.SetXY(x+e.width-rulesDensity.Density(12), y+e.height/2-rulesDensity.Density(2))

	return
}
