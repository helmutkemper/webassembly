package ornament

import (
	"fmt"
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/examples/ide/easterEgg"
	"github.com/helmutkemper/webassembly/examples/ide/rulesConnection"
	"github.com/helmutkemper/webassembly/examples/ide/rulesDensity"
	"github.com/helmutkemper/webassembly/utilsDraw"
	"image/color"
	"time"
)

type WarningMark interface {
	// GetSvg returns the SVG element of the warning mark
	GetSvg() (svg *html.TagSvg)

	// Update Draw the image
	Update(x, y, width, height rulesDensity.Density) (err error)

	// Flash Makes the warning indication blinking
	Flash(warning bool)
}

// WarningMarkExclamation Responsible for drawing the alert plate symbol, with an exclamation in the middle, warning
// in case of error
type WarningMarkExclamation struct {
	easterEgg.MorseCode

	width                   rulesDensity.Density
	height                  rulesDensity.Density
	warningBackgroundColor  color.RGBA
	warningBorderColor      color.RGBA
	warningExclamationColor color.RGBA
	warningEnabled          bool
	warningMarkMargin       rulesDensity.Density
	warningOpacity          float64
	flashTicker             *time.Ticker
	stopTicker              *time.Ticker
	svg                     *html.TagSvg
	hexagonRed              *html.TagSvgPath
	hexagonWhite            *html.TagSvgPath
	exclamation             *html.TagSvgPath
}

// SetMargin sets the margin, in pixels, of the warning mark
func (e *WarningMarkExclamation) SetMargin(margin rulesDensity.Density) {
	e.warningMarkMargin = margin
}

// GetMargin returns the margin, in pixels, of the warning mark
func (e *WarningMarkExclamation) GetMargin() rulesDensity.Density {
	return e.warningMarkMargin
}

// GetSvg returns the SVG element of the warning mark
func (e *WarningMarkExclamation) GetSvg() (svg *html.TagSvg) {
	return e.svg
}

// GetWarning returns the visibility of the warning mark
func (e *WarningMarkExclamation) GetWarning() bool {
	return e.warningEnabled
}

// Init Initializes the instance
func (e *WarningMarkExclamation) Init() (err error) {
	e.MorseCode.Init()

	e.warningBackgroundColor = rulesConnection.KTrafficSignBackgroundColor
	e.warningBorderColor = rulesConnection.KTrafficSignBorderColor
	e.warningExclamationColor = rulesConnection.KTrafficSignWarningExclamationColor
	e.warningOpacity = 0.5

	e.warningEnabled = false

	e.svg = factoryBrowser.NewTagSvg().
		Opacity(e.warningOpacity).
		Id("uniqueid") //.
	//Visibility("hidden")

	e.hexagonRed = factoryBrowser.NewTagSvgPath().
		Fill(e.warningBorderColor).
		FillRule("evenodd").
		Stroke("none")
	e.svg.Append(e.hexagonRed)

	e.hexagonWhite = factoryBrowser.NewTagSvgPath().
		Fill(e.warningBackgroundColor).
		FillRule("evenodd").
		Stroke("none")
	e.svg.Append(e.hexagonWhite)

	e.exclamation = factoryBrowser.NewTagSvgPath().
		Fill(e.warningExclamationColor).
		FillRule("evenodd")
	e.svg.Append(e.exclamation)

	return
}

// min Returns the minimum value
func (e *WarningMarkExclamation) min(a, b rulesDensity.Density) rulesDensity.Density {
	if a < b {
		return a
	}
	return b
}

func (e *WarningMarkExclamation) GetWidth() rulesDensity.Density {
	return e.width
}

func (e *WarningMarkExclamation) GetHeight() rulesDensity.Density {
	return e.height
}

// Update Draw the image
func (e *WarningMarkExclamation) Update(_, _, width, height rulesDensity.Density) (err error) {
	e.width = width
	e.height = height

	//e.svg.ViewBox([]int{0, 0, width, height})
	marginInternal := rulesDensity.Density(0)
	r := e.min(width-marginInternal-2.0*e.warningMarkMargin, height-marginInternal-2.0*e.warningMarkMargin) / 2.0
	rotation := 0.0 // -math.Pi / 2;
	hexagonExternalPath := utilsDraw.PolygonPath(6, r, width/2, height/2, rotation)
	hexagonInternalPath := utilsDraw.PolygonPath(6, r-5, width/2, height/2, rotation)

	e.hexagonRed.D(append(hexagonExternalPath, hexagonInternalPath...))
	e.hexagonWhite.D(hexagonInternalPath)

	// draw exclamation mark (!)
	// The points are based on the path of the image below, which makes an exclamation
	// M 185 120 L 190 240 L 210 240 L 215 120 L 185 120 z
	// M 190 260 L 190 280 L 210 280 L 210 260 L 190 260 z
	originalPoints := [][]rulesDensity.Density{
		{185, 120}, {190, 240}, {210, 240}, {215, 120}, {185, 120},
		{190, 260}, {190, 280}, {210, 280}, {210, 260}, {190, 260},
	}
	// The points are calculated in a box whit 400 x 400 and converted into vector
	pl := utilsDraw.PointsInTheBox(originalPoints, r, width, height, 0)

	exclamationMark := []string{
		fmt.Sprintf("M %v %v", pl[0][0], pl[0][1]), // M 185 120
		fmt.Sprintf("L %v %v", pl[1][0], pl[1][1]), // L 190 240
		fmt.Sprintf("L %v %v", pl[2][0], pl[2][1]), // L 210 240
		fmt.Sprintf("L %v %v", pl[3][0], pl[3][1]), // L 215 120
		fmt.Sprintf("L %v %v", pl[4][0], pl[4][1]), // L 185 120
		"z", // z

		fmt.Sprintf("M %v %v", pl[5][0], pl[5][1]), // M 190 260
		fmt.Sprintf("L %v %v", pl[6][0], pl[6][1]), // L 190 280
		fmt.Sprintf("L %v %v", pl[7][0], pl[7][1]), // L 210 280
		fmt.Sprintf("L %v %v", pl[8][0], pl[8][1]), // L 210 260
		fmt.Sprintf("L %v %v", pl[9][0], pl[9][1]), // L 190 260
		"z", // z
	}
	e.exclamation.D(exclamationMark)

	return
}

// Flash Makes the warning indication blinking
func (e *WarningMarkExclamation) Flash(warning bool) {
	if warning {
		e.MorseCode.FlashErrorMsg(e.svg)
		return
	}

	e.MorseCode.FlashEnd()
}
