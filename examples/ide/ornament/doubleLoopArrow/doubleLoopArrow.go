package doubleLoopArrow

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

// DoubleLoopArrow Responsible for drawing the ornament used in the loop function, a box with two rounded arrows
type DoubleLoopArrow struct {
	//ornament.WarningMarkExclamation

	arrowNormalColor            color.RGBA
	backgroundNormalColor       color.RGBA
	stopButtonCircleNormalColor color.RGBA
	stopButtonBorderNormalColor color.RGBA

	arrowSelectedColor            color.RGBA
	backgroundSelectedColor       color.RGBA
	stopButtonCircleSelectedColor color.RGBA
	stopButtonBorderSelectedColor color.RGBA

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

func (e *DoubleLoopArrow) StopButtonSetup(setup connection.Setup) {
	e.stopButtonConnectionArea.Setup(setup)
}

func (e *DoubleLoopArrow) ToPngResized(width, height float64) (pngData js.Value) {
	return e.svg.ToPngResized(width, height)
}

func (e *DoubleLoopArrow) SetSelected(selected bool) {
	if selected {
		e.backgroundContent.Fill(e.backgroundSelectedColor)
		e.borderArrow.Stroke(e.arrowSelectedColor)
		e.stopButtonCircle.Fill(e.stopButtonCircleSelectedColor)
		e.stopButtonCircle.Stroke(e.stopButtonCircleSelectedColor)
		e.stopButtonBorder.Stroke(e.stopButtonBorderSelectedColor)
		return
	}

	e.backgroundContent.Fill(e.backgroundNormalColor)
	e.borderArrow.Stroke(e.arrowNormalColor)
	e.stopButtonCircle.Fill(e.stopButtonCircleNormalColor)
	e.stopButtonCircle.Stroke(e.stopButtonCircleNormalColor)
	e.stopButtonBorder.Stroke(e.stopButtonBorderNormalColor)
}

func (e *DoubleLoopArrow) SetStopButtonColor(border, circle color.RGBA) {
	e.stopButtonBorderNormalColor = border
	e.stopButtonCircleNormalColor = circle

	e.stopButtonBorder.Stroke(e.stopButtonBorderNormalColor)
	e.stopButtonCircle.Fill(e.stopButtonCircleNormalColor)
	e.stopButtonCircle.Stroke(e.stopButtonCircleNormalColor)
}

// SetArrowColor defines the color of the arrow used as a border.
func (e *DoubleLoopArrow) SetArrowColor(color color.RGBA) {
	e.arrowNormalColor = color
	e.borderArrow.Stroke(e.arrowNormalColor)
}

// SetBackgroundColor defines the color of the background.
func (e *DoubleLoopArrow) SetBackgroundColor(color color.RGBA) {
	e.backgroundNormalColor = color
	e.backgroundContent.Fill(e.backgroundNormalColor)
}

// GetSvg Returns the SVG tag with the element design
func (e *DoubleLoopArrow) GetSvg() (svg *html.TagSvg) {
	return e.svg
}

// Init Initializes the element design
func (e *DoubleLoopArrow) Init() (err error) {

	e.arrowSelectedColor = color.RGBA{R: 255, G: 80, B: 0, A: 255}
	e.backgroundSelectedColor = color.RGBA{R: 255, G: 200, B: 200, A: 255}
	e.stopButtonCircleSelectedColor = color.RGBA{R: 255, G: 0, B: 0, A: 255}
	e.stopButtonBorderSelectedColor = color.RGBA{R: 0, G: 0, B: 255, A: 255}

	e.arrowNormalColor = color.RGBA{R: 255, G: 100, B: 0, A: 255}
	e.backgroundNormalColor = color.RGBA{R: 255, G: 240, B: 240, A: 255}
	e.stopButtonCircleNormalColor = color.RGBA{R: 255, G: 0, B: 0, A: 255}
	e.stopButtonBorderNormalColor = color.RGBA{R: 0, G: 0, B: 255, A: 255}

	e.svg = factoryBrowser.NewTagSvg()

	e.backgroundContent = factoryBrowser.NewTagSvgPath().
		Fill(e.backgroundNormalColor).
		Stroke("none").
		MarkerEnd("url(#backgroundContent)")
	e.svg.Append(e.backgroundContent)

	e.borderArrow = factoryBrowser.NewTagSvgPath().
		Fill("none").
		Stroke(e.arrowNormalColor).
		StrokeWidth(rulesDensity.NewInt(5).GetInt()).
		StrokeLineCap(html.KSvgStrokeLinecapRound).
		StrokeLineJoin(html.KSvgStrokeLinejoinRound).
		MarkerEnd("url(#borderArrow)")
	e.svg.Append(e.borderArrow)

	e.stopButtonCircle = factoryBrowser.NewTagSvgPath().
		Fill(e.stopButtonCircleNormalColor).
		Stroke(e.stopButtonCircleNormalColor).
		StrokeWidth(rulesDensity.NewInt(2).GetInt()).
		MarkerEnd("url(#stopButtonCircle)")
	e.svg.Append(e.stopButtonCircle)

	e.stopButtonBorder = factoryBrowser.NewTagSvgPath().
		Fill("none").
		Stroke(e.stopButtonBorderNormalColor).
		StrokeWidth(rulesDensity.NewInt(2).GetInt()).
		MarkerEnd("url(#stopButtonBorder)")
	e.svg.Append(e.stopButtonBorder)

	e.stopButtonConnection = factoryConnection.NewConnection("bool", "url(#stopButtonConnection)")
	e.svg.Append(e.stopButtonConnection)

	e.stopButtonConnectionArea.Init("url(#stopButtonConnectionArea)")
	e.svg.Append(e.stopButtonConnectionArea.GetSvgPath())

	return
}

// Update Draw the element design
func (e *DoubleLoopArrow) Update(x, y, width, height rulesDensity.Density) (err error) {
	//e.svg.ViewBox([]int{0, 0, width, height})

	margin := rulesDensity.Density(10)
	r := rulesDensity.Density(20)
	s := rulesDensity.Density(40)

	// Define the double loop arrow path data
	arrow := []string{
		// Draw the top-right arrow
		// Base part of the arrow
		fmt.Sprintf("M %v %v", margin+s, margin),
		fmt.Sprintf("l %v %v", rulesDensity.Density(15), rulesDensity.Density(7)),

		// Arrowhead
		fmt.Sprintf("M %v %v", margin+s, margin),
		fmt.Sprintf("l %v %v", rulesDensity.Density(15), rulesDensity.Density(-7)),

		// Curved body of the arrow
		fmt.Sprintf("M %v %v", margin+s, margin),
		fmt.Sprintf("H %v", width-margin-r),
		fmt.Sprintf("Q %v %v, %v %v", width-margin, margin, width-margin, margin+r),
		fmt.Sprintf("V %v", height-margin-s),

		// Draw the bottom-left arrow
		// Base part of the arrow
		fmt.Sprintf("M %v %v", width-margin-s, height-margin),
		fmt.Sprintf("l %v %v", rulesDensity.Density(-15), rulesDensity.Density(7)),

		// Arrowhead
		fmt.Sprintf("M %v %v", width-margin-s, height-margin),
		fmt.Sprintf("l %v %v", rulesDensity.Density(-15), rulesDensity.Density(-7)),

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
	cr := rulesDensity.Density(5.0)
	cx := rulesDensity.Density(20.0)
	cy := rulesDensity.Density(20.0)
	xp := width - margin - 2.0*cr - 1.5*cx
	yp := height - margin - 2.0*cr - 1.5*cy
	L := 2*cr + rulesDensity.Density(10)

	// Define the path data for the stop button circle
	stopButtonCirclePath := []string{
		fmt.Sprintf("M %v %v", width-margin-2.0*cr-cx, height-margin-2.0*cr-cy),
		fmt.Sprintf("m -%v, 0", cr),
		fmt.Sprintf("a %v, %v 0 1, 1 %v, 0", cr, cr, 2*cr),  //--------------
		fmt.Sprintf("a %v, %v 0 1, 1 -%v, 0", cr, cr, 2*cr), //--------------
		"z",
	}
	e.stopButtonCircle.D(stopButtonCirclePath)

	// Define the path data for the stop button border
	stopButtonBorderPath := []string{
		fmt.Sprintf("M %v %v", xp-cr-rulesDensity.Density(5), yp-cr-rulesDensity.Density(5)),                                                     //-----------------------
		fmt.Sprintf("M %v %v", xp+rulesDensity.Density(5), yp),                                                                                   //-----------------------
		fmt.Sprintf("h %v", L-rulesDensity.Density(10)),                                                                                          //-----------------------
		fmt.Sprintf("a %v,%v 0 0 1 %v,%v", rulesDensity.Density(5), rulesDensity.Density(5), rulesDensity.Density(5), rulesDensity.Density(5)),   //---------------------
		fmt.Sprintf("v %v", L-rulesDensity.Density(10)),                                                                                          //-----------------------
		fmt.Sprintf("a %v,%v 0 0 1 -%v,%v", rulesDensity.Density(5), rulesDensity.Density(5), rulesDensity.Density(5), rulesDensity.Density(5)),  //---------------------
		fmt.Sprintf("h -%v", L-rulesDensity.Density(10)),                                                                                         //-----------------------
		fmt.Sprintf("a %v,%v 0 0 1 -%v,-%v", rulesDensity.Density(5), rulesDensity.Density(5), rulesDensity.Density(5), rulesDensity.Density(5)), //---------------------
		fmt.Sprintf("v -%v", L-rulesDensity.Density(10)),                                                                                         //-----------------------
		fmt.Sprintf("a %v,%v 0 0 1 %v,-%v", rulesDensity.Density(5), rulesDensity.Density(5), rulesDensity.Density(5), rulesDensity.Density(5)),  //---------------------
		"z",
	}
	e.stopButtonBorder.D(stopButtonBorderPath)

	e.stopButtonConnection.D(rulesConnection.GetPathDraw(width-rulesDensity.Density(57), height-rulesDensity.Density(42)))
	e.stopButtonConnectionArea.GetSvgPath().D(rulesConnection.GetPathAreaDraw(width-rulesDensity.Density(57), height-rulesDensity.Density(42)))
	e.stopButtonConnectionArea.SetXY(x+width-rulesDensity.Density(57), y+height-rulesDensity.Density(42))

	return
}
