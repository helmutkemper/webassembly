// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feComposite
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feComposite
//
//  <svg style="width:800px; height:400px; display: inline;" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink">
//    <defs>
//      <filter id="imageOver">
//        <feImage xlink:href="mdn_logo_only_color.png" x="10px" y="10px" width="160px" />
//        <feComposite in2="SourceGraphic" operator="over"/>
//      </filter>
//      <filter id="imageIn">
//        <feImage xlink:href="mdn_logo_only_color.png" x="10px" y="10px" width="160px" />
//        <feComposite in2="SourceGraphic" operator="in"/>
//      </filter>
//      <filter id="imageOut">
//        <feImage xlink:href="mdn_logo_only_color.png" x="10px" y="10px" width="160px" />
//        <feComposite in2="SourceGraphic" operator="out"/>
//      </filter>
//      <filter id="imageAtop">
//        <feImage xlink:href="mdn_logo_only_color.png" x="10px" y="10px" width="160px" />
//        <feComposite in2="SourceGraphic" operator="atop"/>
//      </filter>
//      <filter id="imageXor">
//        <feImage xlink:href="mdn_logo_only_color.png" x="10px" y="10px" width="160px" />
//        <feComposite in2="SourceGraphic" operator="xor"/>
//      </filter>
//      <filter id="imageArithmetic">
//        <feImage xlink:href="mdn_logo_only_color.png" x="10px" y="10px" width="160px" />
//        <feComposite in2="SourceGraphic" operator="arithmetic" k1="0.1" k2="0.2" k3="0.3" k4="0.4" />
//      </filter>
//      <filter id="imageLighter">
//        <feImage xlink:href="mdn_logo_only_color.png" x="10px" y="10px" width="160px" />
//        <feComposite in2="SourceGraphic" operator="lighter"/>
//      </filter>
//    </defs>
//    <g transform="translate(0,25)">
//      <circle cx="90px" cy="80px" r="70px" fill="#c00" style="filter:url(#imageOver)"/>
//      <text x="80" y="-5">over</text>
//    </g>
//    <g transform="translate(200,25)">
//      <circle cx="90px" cy="80px" r="70px" fill="#c00" style="filter:url(#imageIn)"/>
//      <text x="80" y="-5">in</text>
//    </g>
//    <g transform="translate(400,25)">
//      <circle cx="90px" cy="80px" r="70px" fill="#c00" style="filter:url(#imageOut)"/>
//      <text x="80" y="-5">out</text>
//    </g>
//    <g transform="translate(600,25)">
//      <circle cx="90px" cy="80px" r="70px" fill="#c00" style="filter:url(#imageAtop)"/>
//      <text x="80" y="-5">atop</text>
//    </g>
//    <g transform="translate(0,240)">
//      <circle cx="90px" cy="80px" r="70px" fill="#c00" style="filter:url(#imageXor)"/>
//      <text x="80" y="-5">xor</text>
//    </g>
//    <g transform="translate(200,240)">
//      <circle cx="90px" cy="80px" r="70px" fill="#c00" style="filter:url(#imageArithmetic)"/>
//      <text x="70" y="-5">arithmetic</text>
//    </g>
//    <g transform="translate(400,240)">
//      <circle cx="90px" cy="80px" r="70px" fill="#c00" style="filter:url(#imageLighter)"/>
//      <text x="80" y="-5">lighter</text>
//    </g>
//  </svg>

//go:build js
// +build js

package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/browser/html"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().
		Style("width:800px; height:400px; display: inline;").
		XmlnsXLink("http://www.w3.org/1999/xlink").
		Append(

			factoryBrowser.NewTagSvgDefs().Append(

				factoryBrowser.NewTagSvgFilter().
					Id("imageOver").Append(

					factoryBrowser.NewTagSvgFeImage().HRef("//developer.mozilla.org/files/6457/mdn_logo_only_color.png").X(10).Y(10).Width(160),
					factoryBrowser.NewTagSvgFeComposite().In2(html.KSvgIn2SourceGraphic).Operator(html.KSvgOperatorFeCompositeOver),
				),

				factoryBrowser.NewTagSvgFilter().
					Id("imageIn").Append(

					factoryBrowser.NewTagSvgFeImage().HRef("//developer.mozilla.org/files/6457/mdn_logo_only_color.png").X(10).Y(10).Width(160),
					factoryBrowser.NewTagSvgFeComposite().In2(html.KSvgIn2SourceGraphic).Operator(html.KSvgOperatorFeCompositeIn),
				),

				factoryBrowser.NewTagSvgFilter().
					Id("imageOut").Append(

					factoryBrowser.NewTagSvgFeImage().HRef("//developer.mozilla.org/files/6457/mdn_logo_only_color.png").X(10).Y(10).Width(160),
					factoryBrowser.NewTagSvgFeComposite().In2(html.KSvgIn2SourceGraphic).Operator(html.KSvgOperatorFeCompositeOut),
				),

				factoryBrowser.NewTagSvgFilter().
					Id("imageAtop").Append(

					factoryBrowser.NewTagSvgFeImage().HRef("//developer.mozilla.org/files/6457/mdn_logo_only_color.png").X(10).Y(10).Width(160),
					factoryBrowser.NewTagSvgFeComposite().In2(html.KSvgIn2SourceGraphic).Operator(html.KSvgOperatorFeCompositeAtop),
				),

				factoryBrowser.NewTagSvgFilter().
					Id("imageXor").Append(

					factoryBrowser.NewTagSvgFeImage().HRef("//developer.mozilla.org/files/6457/mdn_logo_only_color.png").X(10).Y(10).Width(160),
					factoryBrowser.NewTagSvgFeComposite().In2(html.KSvgIn2SourceGraphic).Operator(html.KSvgOperatorFeCompositeXor),
				),

				factoryBrowser.NewTagSvgFilter().
					Id("imageArithmetic").Append(

					factoryBrowser.NewTagSvgFeImage().HRef("//developer.mozilla.org/files/6457/mdn_logo_only_color.png").X(10).Y(10).Width(160),
					factoryBrowser.NewTagSvgFeComposite().In2(html.KSvgIn2SourceGraphic).Operator(html.KSvgOperatorFeCompositeArithmetic).K1(0.1).K2(0.2).K3(0.3).K4(0.4),
				),

				factoryBrowser.NewTagSvgFilter().
					Id("imageLighter").Append(

					factoryBrowser.NewTagSvgFeImage().HRef("//developer.mozilla.org/files/6457/mdn_logo_only_color.png").X(10).Y(10).Width(160),
					factoryBrowser.NewTagSvgFeComposite().In2(html.KSvgIn2SourceGraphic).Operator(html.KSvgOperatorFeCompositeLighter),
				),
			),

			factoryBrowser.NewTagSvgG().Transform(factoryBrowser.NewTransform().Translate(0, 25)).Append(
				factoryBrowser.NewTagSvgCircle().Cx(90).Cy(80).R(70).Fill("#c00").Style("filter:url(#imageOver)"),
				factoryBrowser.NewTagSvgText().X(80).Y(-5).Text("over"),
			),

			factoryBrowser.NewTagSvgG().Transform(factoryBrowser.NewTransform().Translate(200, 25)).Append(
				factoryBrowser.NewTagSvgCircle().Cx(90).Cy(80).R(70).Fill("#c00").Style("filter:url(#imageIn)"),
				factoryBrowser.NewTagSvgText().X(80).Y(-5).Text("in"),
			),

			factoryBrowser.NewTagSvgG().Transform(factoryBrowser.NewTransform().Translate(400, 25)).Append(
				factoryBrowser.NewTagSvgCircle().Cx(90).Cy(80).R(70).Fill("#c00").Style("filter:url(#imageOut)"),
				factoryBrowser.NewTagSvgText().X(80).Y(-5).Text("out"),
			),

			factoryBrowser.NewTagSvgG().Transform(factoryBrowser.NewTransform().Translate(600, 25)).Append(
				factoryBrowser.NewTagSvgCircle().Cx(90).Cy(80).R(70).Fill("#c00").Style("filter:url(#imageAtop)"),
				factoryBrowser.NewTagSvgText().X(80).Y(-5).Text("atop"),
			),

			factoryBrowser.NewTagSvgG().Transform(factoryBrowser.NewTransform().Translate(0, 240)).Append(
				factoryBrowser.NewTagSvgCircle().Cx(90).Cy(80).R(70).Fill("#c00").Style("filter:url(#imageXor)"),
				factoryBrowser.NewTagSvgText().X(80).Y(-5).Text("xor"),
			),

			factoryBrowser.NewTagSvgG().Transform(factoryBrowser.NewTransform().Translate(200, 240)).Append(
				factoryBrowser.NewTagSvgCircle().Cx(90).Cy(80).R(70).Fill("#c00").Style("filter:url(#imageArithmetic)"),
				factoryBrowser.NewTagSvgText().X(80).Y(-5).Text("arithmetic"),
			),

			factoryBrowser.NewTagSvgG().Transform(factoryBrowser.NewTransform().Translate(400, 240)).Append(
				factoryBrowser.NewTagSvgCircle().Cx(90).Cy(80).R(70).Fill("#c00").Style("filter:url(#imageLighter)"),
				factoryBrowser.NewTagSvgText().X(80).Y(-5).Text("lighter"),
			),
		)

	stage.Append(s1)

	<-done
}
