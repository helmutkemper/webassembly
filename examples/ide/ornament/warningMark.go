package ornament

import (
	"fmt"
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/drawUtils"
	"github.com/helmutkemper/webassembly/examples/ide/easterEgg"
	"github.com/helmutkemper/webassembly/platform/factoryColor"
	"image/color"
	"time"
)

type WarningMark struct {
	easterEgg.MorseCode

	warningBackgroundColor  color.RGBA
	warningBorderColor      color.RGBA
	warningExclamationColor color.RGBA
	warningFlashEnabled     bool
	warningEnabled          bool
	warningFlashDuration    float64
	warningFlashInterval    float64
	warningFlashUpdate      float64
	warningMarkMargin       int
	warningOpacity          float64
	flashTicker             *time.Ticker
	stopTicker              *time.Ticker
	svgWarning              *html.TagSvg
	hexagonRed              *html.TagSvgPath
	hexagonWhite            *html.TagSvgPath
	exclamation             *html.TagSvgPath
}

// SetWarningMarkFlash enables or disables the warning mark flash
// @param flashEnabled true to enable the flash, false to disable it
func (e *WarningMark) SetWarningMarkFlash(flashEnabled bool) {
	e.warningFlashEnabled = flashEnabled
}

// SetWarningMarkMargin sets the margin of the warning mark
// @param margin The margin value
func (e *WarningMark) SetWarningMarkMargin(margin int) {
	e.warningMarkMargin = margin
}

// GetWarningMarkMargin returns the margin of the warning mark
// @returns The margin value
func (e *WarningMark) GetWarningMarkMargin() int {
	return e.warningMarkMargin
}

// GetWarningMark returns the SVG element of the warning mark
// @returns The SVG element of the warning mark
func (e *WarningMark) GetWarningMark() *html.TagSvg {
	return e.svgWarning
}

// SetWarning sets the visibility of the warning mark
// @param warning true to show the warning mark, false to hide it
func (e *WarningMark) SetWarning(warning bool) {
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
	e.flashMark()
}

// GetWarning returns the visibility of the warning mark
// @returns true if the warning mark is visible, false otherwise
func (e *WarningMark) GetWarning() bool {
	return e.warningEnabled
}

func (e *WarningMark) Init() (err error) {
	e.MorseCode.Init()

	e.warningBackgroundColor = factoryColor.NewWhite()
	e.warningBorderColor = factoryColor.NewRed()
	e.warningExclamationColor = factoryColor.NewBlack()
	e.warningOpacity = 0.5

	e.warningFlashEnabled = false
	e.warningEnabled = false
	e.warningFlashDuration = 0.6
	e.warningFlashInterval = 0.15
	e.warningFlashUpdate = 1.0

	//e.warningMarkMargin = 0.0

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

func (e *WarningMark) min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (e *WarningMark) Update(width, height int) (err error) {
	e.svgWarning.ViewBox([]int{0, 0, width, height})
	marginInternal := 0
	r := e.min(width-marginInternal-2.0*e.warningMarkMargin, height-marginInternal-2.0*e.warningMarkMargin) / 2.0
	rotation := 0.0 // -math.Pi / 2;
	pointsExternal := drawUtils.Polygon(6, r, width, height, rotation)
	pointsInternal := drawUtils.Polygon(6, r-5, width, height, rotation)

	hexagonExternalPath := []string{
		fmt.Sprintf("M %v %v", pointsExternal[0][0], pointsExternal[0][1]), // Move to the first point
	}
	for i := 1; i < len(pointsExternal); i++ {
		hexagonExternalPath = append(hexagonExternalPath, fmt.Sprintf("L %v %v", pointsExternal[i][0], pointsExternal[i][1])) // Draw lines to remaining points
	}
	hexagonExternalPath = append(hexagonExternalPath, "z") // Close the path

	// Draw the hexagon internal
	hexagonInternalPath := []string{
		fmt.Sprintf("M %v %v", pointsInternal[0][0], pointsInternal[0][1]), // Move to the first point
	}
	for i := 1; i < len(pointsInternal); i++ {
		hexagonInternalPath = append(hexagonInternalPath, fmt.Sprintf("L %v %v", pointsInternal[i][0], pointsInternal[i][1])) // Draw lines to remaining points
	}
	hexagonInternalPath = append(hexagonInternalPath, "z") // Close the path

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
	pl := drawUtils.PointsInTheBox(originalPoints, r, width, height, 0)

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

func (e *WarningMark) appendManySlices(list ...[]bool) (slice []bool) {
	slice = make([]bool, 0)
	for i := 0; i < len(list); i++ {
		slice = append(slice, list[i]...)
	}

	return
}

func (e *WarningMark) flashMark() {
	e.MorseCode.FlashMarkSoS(e.svgWarning)
}
