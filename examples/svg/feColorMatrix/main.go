// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feColorMatrix
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feColorMatrix
//
//  <svg width="100%" height="100%" viewBox="0 0 150 500"
//      preserveAspectRatio="xMidYMid meet"
//      xmlns="http://www.w3.org/2000/svg"
//      xmlns:xlink="http://www.w3.org/1999/xlink">
//
//    <!-- ref -->
//    <defs>
//      <g id="circles">
//        <circle cx="30" cy="30" r="20" fill="blue" fill-opacity="0.5" />
//        <circle cx="20" cy="50" r="20" fill="green" fill-opacity="0.5" />
//        <circle cx="40" cy="50" r="20" fill="red" fill-opacity="0.5" />
//      </g>
//    </defs>
//    <use href="#circles" />
//    <text x="70" y="50">Reference</text>
//
//
//
//
//
//
//
//
//
//    <!-- identity matrix -->
//    <filter id="colorMeTheSame">
//      <feColorMatrix in="SourceGraphic"
//          type="matrix"
//          values="1 0 0 0 0
//                  0 1 0 0 0
//                  0 0 1 0 0
//                  0 0 0 1 0" />
//     </filter>
//    <use href="#circles" transform="translate(0 70)" filter="url(#colorMeTheSame)" />
//    <text x="70" y="120">Identity matrix</text>
//
//
//
//
//
//
//
//
//    <!-- Combine RGB into green matrix -->
//    <filter id="colorMeGreen">
//      <feColorMatrix in="SourceGraphic"
//          type="matrix"
//          values="0 0 0 0 0
//                  1 1 1 1 0
//                  0 0 0 0 0
//                  0 0 0 1 0" />
//    </filter>
//    <use href="#circles" transform="translate(0 140)" filter="url(#colorMeGreen)" />
//    <text x="70" y="190">rgbToGreen</text>
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
//    <!-- saturate -->
//    <filter id="colorMeSaturate">
//      <feColorMatrix in="SourceGraphic"
//          type="saturate"
//          values="0.2" />
//    </filter>
//    <use href="#circles" transform="translate(0 210)" filter="url(#colorMeSaturate)" />
//    <text x="70" y="260">saturate</text>
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
//    <!-- hueRotate -->
//    <filter id="colorMeHueRotate">
//      <feColorMatrix in="SourceGraphic"
//          type="hueRotate"
//          values="180" />
//    </filter>
//    <use href="#circles" transform="translate(0 280)" filter="url(#colorMeHueRotate)" />
//    <text x="70" y="330">hueRotate</text>
//
//    <!-- luminanceToAlpha -->
//    <filter id="colorMeLTA">
//      <feColorMatrix in="SourceGraphic"
//          type="luminanceToAlpha" />
//    </filter>
//    <use href="#circles" transform="translate(0 350)" filter="url(#colorMeLTA)" />
//    <text x="70" y="400">luminanceToAlpha</text>
//  </svg>

//go:build js
// +build js

package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/browser/html"
	"github.com/helmutkemper/iotmaker.webassembly/platform/factoryColor"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().
		ViewBox([]float64{0, 0, 150, 500}).
		PreserveAspectRatio(
			html.KRatioXMidYMid,
			html.KMeetOrSliceReferenceMeet,
		).
		XmlnsXLink("http://www.w3.org/1999/xlink").
		Append(

			// ref
			factoryBrowser.NewTagSvgDefs().
				Append(

					factoryBrowser.NewTagSvgG().
						Id("circles").
						Append(

							factoryBrowser.NewTagSvgCircle().
								Cx(30).
								Cy(30).
								R(20).
								Fill(factoryColor.NewBlue()).
								FillOpacity(0.5),

							factoryBrowser.NewTagSvgCircle().
								Cx(20).
								Cy(50).
								R(20).
								Fill(factoryColor.NewGreen()).
								FillOpacity(0.5),

							factoryBrowser.NewTagSvgCircle().
								Cx(40).
								Cy(50).
								R(20).
								Fill(factoryColor.NewRed()).
								FillOpacity(0.5),
						),
				),

			factoryBrowser.NewTagSvgUse().
				HRef("#circles"),

			factoryBrowser.NewTagSvgText().
				X(70).
				Y(50).
				Text("Reference"),

			//
			//
			//
			//
			//
			//
			//
			//
			//

			// identity matrix
			factoryBrowser.NewTagSvgFilter().
				Id("colorMeTheSame").
				Append(

					factoryBrowser.NewTagSvgFeColorMatrix().
						In(html.KSvgInSourceGraphic).
						Type(html.KSvgTypeFeColorMatrixMatrix).
						Values(
							[]float64{
								1, 0, 0, 0, 0,
								0, 1, 0, 0, 0,
								0, 0, 1, 0, 0,
								0, 0, 0, 1, 0,
							},
						),
				),

			factoryBrowser.NewTagSvgUse().
				HRef("#circles").
				Transform(

					factoryBrowser.NewTransform().
						Translate(0, 70),
				).
				Filter("url(#colorMeTheSame)"),

			factoryBrowser.NewTagSvgText().
				X(70).
				Y(120).
				Text("Identity matrix"),

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
			//
			//
			//
			//
			//

			// Combine RGB into green matrix
			factoryBrowser.NewTagSvgFilter().
				Id("colorMeGreen").
				Append(

					factoryBrowser.NewTagSvgFeColorMatrix().
						In(html.KSvgInSourceGraphic).
						Type(html.KSvgTypeFeColorMatrixMatrix).
						Values(
							[]float64{
								0, 0, 0, 0, 0,
								1, 1, 1, 1, 0,
								0, 0, 0, 0, 0,
								0, 0, 0, 1, 0,
							},
						),
				),

			factoryBrowser.NewTagSvgUse().
				HRef("#circles").
				Transform(

					factoryBrowser.NewTransform().
						Translate(0, 140),
				).
				Filter("url(#colorMeGreen)"),

			factoryBrowser.NewTagSvgText().
				X(70).
				Y(190).
				Text("rgbToGreen"),

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

			//saturate
			factoryBrowser.NewTagSvgFilter().
				Id("colorMeSaturate").
				Append(

					factoryBrowser.NewTagSvgFeColorMatrix().
						In(html.KSvgInSourceGraphic).
						Type(html.KSvgTypeFeColorMatrixSaturate).
						Values(0.2),
				),

			factoryBrowser.NewTagSvgUse().
				HRef("#circles").
				Transform(
					factoryBrowser.NewTransform().Translate(0, 210),
				).
				Filter("url(#colorMeSaturate)"),

			factoryBrowser.NewTagSvgText().
				X(70).
				Y(260).
				Text("Saturate"),

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

			// hueRotate
			factoryBrowser.NewTagSvgFilter().
				Id("colorMeHueRotate").
				Append(

					factoryBrowser.NewTagSvgFeColorMatrix().
						In(html.KSvgInSourceGraphic).
						Type(html.KSvgTypeFeColorMatrixHueRotate).
						Values(180),
				),

			factoryBrowser.NewTagSvgUse().
				HRef("#circles").
				Transform(

					factoryBrowser.NewTransform().Translate(0, 280),
				).
				Filter("url(#colorMeHueRotate)"),

			factoryBrowser.NewTagSvgText().
				X(70).
				Y(330).
				Text("hueRotate"),

			// luminanceToAlpha
			factoryBrowser.NewTagSvgFilter().
				Id("colorMeLTA").
				Append(

					factoryBrowser.NewTagSvgFeColorMatrix().
						In(html.KSvgInSourceGraphic).
						Type(html.KSvgTypeFeColorMatrixLuminanceToAlpha),
				),

			factoryBrowser.NewTagSvgUse().
				HRef("#circles").
				Transform(

					factoryBrowser.NewTransform().Translate(0, 350),
				).
				Filter("url(#colorMeLTA)"),

			factoryBrowser.NewTagSvgText().
				X(70).
				Y(400).
				Text("luminanceToAlpha"),
		)

	stage.Append(s1)

	<-done
}
