package ornament

import (
	"fmt"
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/examples/ide/easterEgg"
	"github.com/helmutkemper/webassembly/examples/ide/rulesConnection"
	"github.com/helmutkemper/webassembly/utilsDraw"
	"image/color"
	"time"
)

// WarningMarkExclamation Responsible for drawing the alert plate symbol, with an exclamation in the middle, warning
// in case of error
type WarningMarkExclamation struct {
	easterEgg.MorseCode

	width                   int
	height                  int
	warningBackgroundColor  color.RGBA
	warningBorderColor      color.RGBA
	warningExclamationColor color.RGBA
	warningEnabled          bool
	warningMarkMargin       int
	warningOpacity          float64
	flashTicker             *time.Ticker
	stopTicker              *time.Ticker
	svgWarning              *html.TagSvg
	hexagonRed              *html.TagSvgPath
	hexagonWhite            *html.TagSvgPath
	exclamation             *html.TagSvgPath
}

// SetWarningMarkMargin sets the margin, in pixels, of the warning mark
func (e *WarningMarkExclamation) SetWarningMarkMargin(margin int) {
	e.warningMarkMargin = margin
}

// GetWarningMarkMargin returns the margin, in pixels, of the warning mark
func (e *WarningMarkExclamation) GetWarningMarkMargin() int {
	return e.warningMarkMargin
}

// GetWarningMark returns the SVG element of the warning mark
func (e *WarningMarkExclamation) GetWarningMark() *html.TagSvg {
	return e.svgWarning
}

// SetWarning sets the visibility of the warning mark
func (e *WarningMarkExclamation) SetWarning(warning bool) {
	if warning == e.warningEnabled {
		e.svgWarning.AddStyle("visibility", "hidden")
		return
	}

	// Update the state and visibility
	e.warningEnabled = warning
	visibility := "hidden"
	if warning {
		visibility = "visible"
	}
	e.svgWarning.AddStyle("visibility", visibility)

	// Trigger the flash mark functionality
	e.flashMark(warning)
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

	e.svgWarning = factoryBrowser.NewTagSvg().
		Opacity(e.warningOpacity).
		Id("uniqueid") //.
	//Visibility("hidden")

	e.hexagonRed = factoryBrowser.NewTagSvgPath().
		Fill(e.warningBorderColor).
		FillRule("evenodd").
		Stroke("none")
	e.svgWarning.Append(e.hexagonRed)

	e.hexagonWhite = factoryBrowser.NewTagSvgPath().
		Fill(e.warningBackgroundColor).
		FillRule("evenodd").
		Stroke("none")
	e.svgWarning.Append(e.hexagonWhite)

	e.exclamation = factoryBrowser.NewTagSvgPath().
		Fill(e.warningExclamationColor).
		FillRule("evenodd")
	e.svgWarning.Append(e.exclamation)

	return
}

// min Returns the minimum value
func (e *WarningMarkExclamation) min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (e *WarningMarkExclamation) GetWidth() int {
	return e.width
}

func (e *WarningMarkExclamation) GetHeight() int {
	return e.height
}

// Update Draw the image
func (e *WarningMarkExclamation) Update(_, _, width, height int) (err error) {
	e.width = width
	e.height = height

	//e.svgWarning.ViewBox([]int{0, 0, width, height})
	marginInternal := 0
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
	originalPoints := [][]int{
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

// flashMark Makes the warning indication blinking
func (e *WarningMarkExclamation) flashMark(warning bool) {
	if warning {
		e.MorseCode.FlashMarkErrorMsg(e.svgWarning)
		return
	}

	e.MorseCode.FlashEnd()
}
