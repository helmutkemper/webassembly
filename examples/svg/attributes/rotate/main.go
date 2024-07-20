// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/rotate
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/rotate
//
//  <svg width="400" height="120" viewBox="0 0 480 120"
//      xmlns="http://www.w3.org/2000/svg">
//
//    <!-- Draw the outline of the motion path in grey -->
//    <path d="M10,110 A120,120 -45 0,1 110 10 A120,120 -45 0,1 10,110"
//        stroke="lightgrey" stroke-width="2"
//        fill="none" id="theMotionPath"/>
//
//    <!-- Red arrow which will not rotate -->
//    <path fill="red" d="M-5,-5 L10,0 -5,5 0,0 Z">
//      <!-- Define the motion path animation -->
//      <animateMotion dur="6s" repeatCount="indefinite" rotate="0">
//        <mpath href="#theMotionPath"/>
//      </animateMotion>
//    </path>
//
//    <g transform="translate(100, 0)">
//      <use href="#theMotionPath"/>
//    <!-- Green arrow which will rotate along the motion path -->
//    <path fill="green" d="M-5,-5 L10,0 -5,5 0,0 Z">
//      <!-- Define the motion path animation -->
//      <animateMotion dur="6s" repeatCount="indefinite" rotate="auto">
//        <mpath href="#theMotionPath"/>
//      </animateMotion>
//    </path>
//    </g>
//
//    <g transform="translate(200, 0)">
//      <use href="#theMotionPath"/>
//    <!-- Blue arrow which will rotate backwards along the motion path -->
//    <path fill="blue" d="M-5,-5 L10,0 -5,5 0,0 Z">
//      <!-- Define the motion path animation -->
//      <animateMotion dur="6s" repeatCount="indefinite" rotate="auto-reverse">
//        <mpath href="#theMotionPath"/>
//      </animateMotion>
//    </path>
//    </g>
//
//    <g transform="translate(300, 0)">
//      <use href="#theMotionPath"/>
//    <!-- Purple arrow which will have a static rotation of 210 degrees -->
//    <path fill="purple" d="M-5,-5 L10,0 -5,5 0,0 Z">
//      <!-- Define the motion path animation -->
//      <animateMotion dur="6s" repeatCount="indefinite" rotate="210">
//        <mpath href="#theMotionPath"/>
//      </animateMotion>
//    </path>
//    </g>
//  </svg>

//go:build js

package main

import (
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/platform/factoryColor"
	"time"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()
	s1 := factoryBrowser.NewTagSvg().Width(400).Height(120).ViewBox([]float64{0, 0, 480, 120}).Append(
		// Draw the outline of the motion path in grey
		factoryBrowser.NewTagSvgPath().D(factoryBrowser.NewPath().M(10, 110).A(120, 120, -45, 0, 1, 110, 10).A(120, 120, -45, 0, 1, 10, 110)).Stroke(factoryColor.NewLightgray()).StrokeWidth(2).Fill(nil).Id("theMotionPath"),

		// Red arrow which will not rotate
		factoryBrowser.NewTagSvgPath().Fill(factoryColor.NewRed()).D(factoryBrowser.NewPath().M(-5, -5).L(10, 0).L(-5, 5).L(0, 0).Z()).Append(
			// Define the motion path animation
			factoryBrowser.NewTagSvgAnimateMotion().Dur(6*time.Second).RepeatCount(html.KSvgDurIndefinite).Rotate(0).Append(
				factoryBrowser.NewTagSvgMPath().HRef("#theMotionPath"),
			),
		),

		factoryBrowser.NewTagSvgG().Transform(factoryBrowser.NewTransform().Translate(100, 0)).Append(
			factoryBrowser.NewTagSvgUse().HRef("#theMotionPath"),

			// Green arrow which will rotate along the motion path
			factoryBrowser.NewTagSvgPath().Fill(factoryColor.NewGreen()).D(factoryBrowser.NewPath().M(-5, -5).L(10, 0).L(-5, 5).L(0, 0).Z()).Append(
				// Define the motion path animation
				factoryBrowser.NewTagSvgAnimateMotion().Dur(6*time.Second).RepeatCount(html.KSvgDurIndefinite).Rotate(html.KSvgRotateAuto).Append(
					factoryBrowser.NewTagSvgMPath().HRef("#theMotionPath"),
				),
			),
		),

		factoryBrowser.NewTagSvgG().Transform(factoryBrowser.NewTransform().Translate(200, 0)).Append(
			factoryBrowser.NewTagSvgUse().HRef("#theMotionPath"),
			// Blue arrow which will rotate backwards along the motion path
			factoryBrowser.NewTagSvgPath().Fill(factoryColor.NewBlue()).D(factoryBrowser.NewPath().M(-5, -5).L(10, 0).L(-5, 5).L(0, 0).Z()).Append(
				// Define the motion path animation
				factoryBrowser.NewTagSvgAnimateMotion().Dur(6*time.Second).RepeatCount(html.KSvgDurIndefinite).Rotate(html.KSvgRotateAutoReverse).Append(
					factoryBrowser.NewTagSvgMPath().HRef("#theMotionPath"),
				),
			),
		),

		factoryBrowser.NewTagSvgG().Transform(factoryBrowser.NewTransform().Translate(300, 0)).Append(
			factoryBrowser.NewTagSvgUse().HRef("#theMotionPath"),
			// Purple arrow which will have a static rotation of 210 degrees
			factoryBrowser.NewTagSvgPath().Fill(factoryColor.NewPurple()).D(factoryBrowser.NewPath().M(-5, -5).L(10, 0).L(-5, 5).L(0, 0).Z()).Append(
				// Define the motion path animation
				factoryBrowser.NewTagSvgAnimateMotion().Dur(6*time.Second).RepeatCount(html.KSvgDurIndefinite).Rotate(210).Append(
					factoryBrowser.NewTagSvgMPath().HRef("#theMotionPath"),
				),
			),
		),
	)

	stage.Append(s1)

	<-done
}

//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
