package doubleLoopArrow

import (
	"fmt"
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/examples/ide/connection"
	"github.com/helmutkemper/webassembly/examples/ide/ornament"
	"github.com/helmutkemper/webassembly/examples/ide/rulesConnection"
	"image/color"
	"syscall/js"
)

// DoubleLoopArrow Responsible for drawing the ornament used in the loop function, a box with two rounded arrows
type DoubleLoopArrow struct {
	ornament.WarningMarkExclamation

	arrowColor      color.RGBA
	backgroundColor color.RGBA

	svg                      *html.TagSvg
	backgroundContent        *html.TagSvgPath
	borderArrow              *html.TagSvgPath
	stopButtonCircle         *html.TagSvgPath
	stopButtonBorder         *html.TagSvgPath
	stopButtonConnection     *html.TagSvgPath
	stopButtonConnectionArea connection.Connection
}

func (e *DoubleLoopArrow) GetConnectionError() (err error) {
	return rulesConnection.GetError()
}

func (e *DoubleLoopArrow) SetClickFunc(f js.Func) {
	e.stopButtonConnectionArea.SetClickFunc(f)
}

func (e *DoubleLoopArrow) SetStopButtonSetAsDataInput(dataInput bool) {
	e.stopButtonConnectionArea.SetAsDataInput(dataInput)
}

func (e *DoubleLoopArrow) SetStopButtonConnectionLockedUp(lockedUp bool) {
	e.stopButtonConnectionArea.SetConnectionLockedUp(lockedUp)
}

func (e *DoubleLoopArrow) SetStopButtonAcceptNotConnected(accept bool) {
	e.stopButtonConnectionArea.SetAcceptNotConnected(accept)
}

func (e *DoubleLoopArrow) SetStopButtonName(name string) {
	e.stopButtonConnectionArea.SetName(name)
}

func (e *DoubleLoopArrow) SetStopButtonDataType(connType string) {
	e.stopButtonConnectionArea.SetDataType(connType)
}

func (e *DoubleLoopArrow) SetStopButtonFatherId(fatherId string) {
	e.stopButtonConnectionArea.SetFatherId(fatherId)
}

func (e *DoubleLoopArrow) ToPngResized(width, height float64) (pngData js.Value) {
	return e.svg.ToPngResized(width, height)
}

// SetWarning sets the visibility of the warning mark
func (e *DoubleLoopArrow) SetWarning(warning bool) {
	e.WarningMarkExclamation.SetWarning(warning)
}

// SetArrowColor defines the color of the arrow used as a border.
func (e *DoubleLoopArrow) SetArrowColor(color color.RGBA) {
	e.arrowColor = color
	e.borderArrow.Stroke(e.arrowColor)
}

// SetBackgroundColor defines the color of the background.
func (e *DoubleLoopArrow) SetBackgroundColor(color color.RGBA) {
	e.backgroundColor = color
	e.backgroundContent.Fill(e.backgroundColor)
}

// GetSvg Returns the SVG tag with the element design
func (e *DoubleLoopArrow) GetSvg() (svg *html.TagSvg) {
	return e.svg
}

// Init Initializes the element design
func (e *DoubleLoopArrow) Init() (err error) {
	_ = e.WarningMarkExclamation.Init()

	e.arrowColor = color.RGBA{R: 255, G: 120, B: 0, A: 255}
	e.backgroundColor = color.RGBA{R: 255, G: 240, B: 240, A: 255}

	e.svg = factoryBrowser.NewTagSvg()

	e.backgroundContent = factoryBrowser.NewTagSvgPath().
		Fill(e.backgroundColor).
		Stroke("none").
		MarkerEnd("url(#backgroundContent)")
	e.svg.Append(e.backgroundContent)

	e.borderArrow = factoryBrowser.NewTagSvgPath().
		Fill("none").
		Stroke(e.arrowColor).
		StrokeWidth(5).
		StrokeLineCap(html.KSvgStrokeLinecapRound).
		StrokeLineJoin(html.KSvgStrokeLinejoinRound).
		MarkerEnd("url(#borderArrow)")
	e.svg.Append(e.borderArrow)

	e.stopButtonCircle = factoryBrowser.NewTagSvgPath().
		Fill("red").
		Stroke("red").
		StrokeWidth(2).
		MarkerEnd("url(#stopButtonCircle)")
	e.svg.Append(e.stopButtonCircle)

	e.stopButtonBorder = factoryBrowser.NewTagSvgPath().
		Fill("none").
		Stroke("blue").
		StrokeWidth(2).
		MarkerEnd("url(#stopButtonBorder)")
	e.svg.Append(e.stopButtonBorder)

	e.stopButtonConnection = factoryBrowser.NewTagSvgPath().
		Fill(rulesConnection.TypeToColor("bool")).
		Stroke("none").
		MarkerEnd("url(#stopButtonConnection)")
	e.svg.Append(e.stopButtonConnection)

	e.stopButtonConnectionArea.Init()
	e.stopButtonConnectionArea.GetSvgPath().
		Fill("transparent").
		Stroke("none").
		AddStyle("cursor", "crosshair").
		MarkerEnd("url(#stopButtonConnection)")
	e.svg.Append(e.stopButtonConnectionArea.GetSvgPath())

	e.svg.Append(e.GetWarningMark())
	e.SetWarning(false)

	return
}

// Update Draw the element design
func (e *DoubleLoopArrow) Update(width, height int) (err error) {
	_ = e.WarningMarkExclamation.Update(width, height)
	e.svg.ViewBox([]int{0.0, 0.0, width, height})

	margin := 10
	r := 30
	s := 40

	// Define the double loop arrow path data
	arrow := []string{
		// Draw the top-right arrow
		// Base part of the arrow
		fmt.Sprintf("M %v %v", margin+s, margin),
		"l 15 7",

		// Arrowhead
		fmt.Sprintf("M %v %v", margin+s, margin),
		"l 15 -7",

		// Curved body of the arrow
		fmt.Sprintf("M %v %v", margin+s, margin),
		fmt.Sprintf("H %v", width-margin-r),
		fmt.Sprintf("Q %v %v, %v %v", width-margin, margin, width-margin, margin+r),
		fmt.Sprintf("V %v", height-margin-s),

		// Draw the bottom-left arrow
		// Base part of the arrow
		fmt.Sprintf("M %v %v", width-margin-s, height-margin),
		"l -15 7",

		// Arrowhead
		fmt.Sprintf("M %v %v", width-margin-s, height-margin),
		"l -15 -7",

		// Curved body of the arrow
		fmt.Sprintf("M %v %v", width-margin-s, height-margin),
		fmt.Sprintf("H %v", margin+r),
		fmt.Sprintf("Q %v %v, %v %v", margin, height-margin, margin, height-margin-r),
		fmt.Sprintf("V %v", margin+s),
	}
	e.borderArrow.D(arrow)

	// Define the rounded background path data
	background := []string{
		// Draw the rounded background
		fmt.Sprintf("M %v %v", margin+r, margin),
		fmt.Sprintf("H %v", width-margin-r),
		fmt.Sprintf("Q %v %v, %v %v", width-margin, margin, width-margin, margin+r),
		fmt.Sprintf("V %v", height-margin-r),
		fmt.Sprintf("Q %v %v, %v %v", width-margin, height-margin, width-margin-r, height-margin),
		fmt.Sprintf("H %v", margin+r),
		fmt.Sprintf("Q %v %v, %v %v", margin, height-margin, margin, height-margin-r),
		fmt.Sprintf("V %v", margin+r),
		fmt.Sprintf("Q %v %v, %v %v", margin, margin, margin+r, margin),
	}
	e.backgroundContent.D(background)

	// draw the stop button
	cr := 5.0
	cx := 20.0
	cy := 20.0
	x := int(float64(width) - float64(margin) - 2.0*cr - 1.5*cx)
	y := int(float64(height) - float64(margin) - 2.0*cr - 1.5*cy)
	L := 2*cr + 10

	// Define the path data for the stop button circle
	stopButtonCirclePath := []string{
		fmt.Sprintf("M %v %v", int(float64(width)-float64(margin)-2.0*cr-cx), int(float64(height)-float64(margin)-2.0*cr-cy)),
		fmt.Sprintf("m -%v, 0", cr),
		fmt.Sprintf("a %v, %v 0 1, 1 %v, 0", cr, cr, 2*cr),
		fmt.Sprintf("a %v, %v 0 1, 1 -%v, 0", cr, cr, 2*cr),
		"z",
	}
	e.stopButtonCircle.D(stopButtonCirclePath)

	// Define the path data for the stop button border
	stopButtonBorderPath := []string{
		fmt.Sprintf("M %v %v", int(float64(x)-cr-5.0), int(float64(y)-cr-5)),
		fmt.Sprintf("M %v %v", x+5, y),
		fmt.Sprintf("h %v", L-10),
		"a 5,5 0 0 1 5,5",
		fmt.Sprintf("v %v", L-10),
		"a 5,5 0 0 1 -5,5",
		fmt.Sprintf("h -%v", L-10),
		"a 5,5 0 0 1 -5,-5",
		fmt.Sprintf("v -%v", L-10),
		"a 5,5 0 0 1 5,-5",
		"z",
	}
	e.stopButtonBorder.D(stopButtonBorderPath)

	stopButtonConnectionPath := []string{
		fmt.Sprintf("M %v %v", width-57, height-42),
		fmt.Sprintf("l %v 0", rulesConnection.KWidth),
		fmt.Sprintf("l 0 %v", rulesConnection.KHeight),
		fmt.Sprintf("l -%v 0", rulesConnection.KWidth),
		fmt.Sprintf("l 0 -%v", rulesConnection.KHeight),
	}
	e.stopButtonConnection.D(stopButtonConnectionPath)

	stopButtonConnectionAreaPath := []string{
		fmt.Sprintf("M %v %v", width-57-(rulesConnection.KWidthArea-rulesConnection.KWidth)/2, height-42-(rulesConnection.KHeightArea-rulesConnection.KHeight)/2),
		fmt.Sprintf("l %v 0", rulesConnection.KWidthArea),
		fmt.Sprintf("l 0 %v", rulesConnection.KHeightArea),
		fmt.Sprintf("l -%v 0", rulesConnection.KWidthArea),
		fmt.Sprintf("l 0 -%v", rulesConnection.KHeightArea),
	}
	e.stopButtonConnectionArea.GetSvgPath().D(stopButtonConnectionAreaPath)

	e.stopButtonConnectionArea.SetX(width - 57)
	e.stopButtonConnectionArea.SetY(height - 42)

	return
}
