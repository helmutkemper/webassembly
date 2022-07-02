// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Element/metadata
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Element/metadata
//
//  <svg width="400" viewBox="0 0 400 300"
//    xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink">
//    <metadata>
//      <rdf:RDF xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#"
//               xmlns:connect="http://www.w3.org/1999/08/29-svg-connections-in-RDF#">
//        <rdf:Description about="#CableA">
//          <connect:ends rdf:resource="#socket1"/>
//          <connect:ends rdf:resource="#ComputerA"/>
//        </rdf:Description>
//        <rdf:Description about="#CableB">
//          <connect:ends rdf:resource="#socket2"/>
//          <connect:ends rdf:resource="#ComputerB"/>
//        </rdf:Description>
//        <rdf:Description about="#CableN">
//          <connect:ends rdf:resource="#socket5"/>
//          <connect:ends>Everything</connect:ends>
//        </rdf:Description>
//        <rdf:Description about="#Hub">
//          <connect:ends rdf:resource="#socket1"/>
//          <connect:ends rdf:resource="#socket2"/>
//          <connect:ends rdf:resource="#socket3"/>
//          <connect:ends rdf:resource="#socket4"/>
//          <connect:ends rdf:resource="#socket5"/>
//        </rdf:Description>
//      </rdf:RDF>
//    </metadata>
//    <title>Network</title>
//    <desc>An example of a computer network based on a hub.</desc>
//
//    <style>
//      svg {
//        /* Default styles to be inherited */
//        fill: white;
//        stroke: black;
//      }
//      text { fill: black; stroke: none; }
//      path { fill: none; }
//    </style>
//
//    <!-- Define symbols used in the SVG -->
//    <defs>
//
//      <!-- hubPlug symbol. Used by hub symbol -->
//      <symbol id="hubPlug">
//        <desc>A 10BaseT/100baseTX socket</desc>
//        <path d="M0,10 h5 v-9 h12 v9 h5 v16 h-22 z"/>
//      </symbol>
//
//      <!-- hub symbol -->
//      <symbol id="hub">
//        <desc>A typical 10BaseT/100BaseTX network hub</desc>
//        <text x="0" y="15">Hub</text>
//        <g transform="translate(0 20)">
//          <rect width="253" height="84"/>
//          <rect width="229" height="44" x="12" y="10"/>
//          <circle fill="red" cx="227" cy="71" r="7" />
//          <!-- five groups each using the defined socket -->
//          <g id="sock1et" transform="translate(25 20)">
//            <title>Socket 1</title>
//            <use xlink:href="#hubPlug"/>
//          </g>
//          <g id="socket2" transform="translate(70 20)">
//            <title>Socket 2</title>
//            <use xlink:href="#hubPlug"/>
//          </g>
//          <g id="socket3" transform="translate(115 20)">
//            <title>Socket 3</title>
//            <use xlink:href="#hubPlug"/>
//          </g>
//          <g id="socket4" transform="translate(160 20)">
//            <title>Socket 4</title>
//            <use xlink:href="#hubPlug"/>
//          </g>
//          <g id="socket5" transform="translate(205 20)">
//            <title>Socket 5</title>
//            <use xlink:href="#hubPlug"/>
//          </g>
//        </g>
//      </symbol>
//
//      <!-- computer symbol -->
//      <symbol id="computer">
//        <desc>A common desktop PC</desc>
//        <g id="monitorStand" transform="translate(40 121)">
//          <title>Monitor stand</title>
//          <desc>One of those cool swivelling monitor stands that sit under the monitor</desc>
//          <path d="m0,0 S 10 10 40 12"/>
//          <path d="m80,0 S 70 10 40 12"/>
//          <path d="m0,20 L 10 10 S 40 12 70 10 L 80 20z"/>
//        </g>
//        <g id="monitor">
//          <title>Monitor</title>
//          <desc>A very fancy monitor</desc>
//          <rect width="160" height="120"/>
//          <rect fill="lightgrey" width="138" height="95" x="11" y="12"/>
//        </g>
//        <g id="processor" transform="translate(0 142)">
//          <title>The computer</title>
//          <desc>A desktop computer - broad flat box style</desc>
//          <rect width="160" height="60"/>
//          <g id="discDrive" transform="translate(70 8)">
//            <title>disc drive</title>
//            <desc>A built-in disc drive</desc>
//            <rect width="58" height="3" x="12" y="8"/>
//            <rect width="8" height="2" x="12" y="15"/>
//          </g>
//          <circle cx="135" cy="40" r="5"/>
//        </g>
//       </symbol>
//    </defs>
//
//    <text x="0" y="15">Network</text>
//
//    <!-- Use the hub symbol. -->
//    <g id="Hub" transform="translate(80 45)">
//      <title>Hub</title>
//      <use xlink:href="#hub" transform="scale(0.75)"/>
//    </g>
//
//    <!-- Use the computer symbol. -->
//    <g id="ComputerA" transform="translate(20 170)">
//      <title>Computer A</title>
//      <use xlink:href="#computer" transform="scale(0.5)"/>
//    </g>
//
//    <!-- Use the same computer symbol. -->
//    <g id="ComputerB" transform="translate(300 170)">
//      <title>Computer B</title>
//      <use xlink:href="#computer" transform="scale(0.5)"/>
//    </g>
//
//    <!-- Draw Cable A. -->
//    <g id="CableA" transform="translate(107 88)">
//      <title>Cable A</title>
//      <desc>10BaseT twisted pair cable</desc>
//      <path d="M0,0c100,140 50,140 -8,160"/>
//    </g>
//
//    <!-- Draw Cable B. -->
//    <g id="CableB" transform="translate(142 88)">
//      <title>Cable B</title>
//      <desc>10BaseT twisted pair cable</desc>
//      <path d="M0,0c100,180 110,160 159,160"/>
//    </g>
//
//    <!-- Draw Cable N. -->
//    <g id="CableN" transform="translate(242 88)">
//       <title>Cable N</title>
//       <desc>10BaseT twisted pair cable</desc>
//       <path d="M0,0c0,-70 20,-50 60,-50"/>
//    </g>
//  </svg>

//go:build js
// +build js

package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/platform/factoryColor"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().
		ViewBox([]float64{0, 0, 400, 300}).
		XmlnsXLink("http://www.w3.org/1999/xlink").
		Append(

			factoryBrowser.NewTagSvgMetadata().
				Html(
					"<rdf:RDF xmlns:rdf=\"http://www.w3.org/1999/02/22-rdf-syntax-ns#\"\n"+
						"xmlns:connect=\"http://www.w3.org/1999/08/29-svg-connections-in-RDF#\">\n"+
						"<rdf:Description about=\"#CableA\">\n"+
						"<connect:ends rdf:resource=\"#socket1\"/>\n"+
						"<connect:ends rdf:resource=\"#ComputerA\"/>\n"+
						"</rdf:Description>\n"+
						"<rdf:Description about=\"#CableB\">\n"+
						"<connect:ends rdf:resource=\"#socket2\"/>\n"+
						"<connect:ends rdf:resource=\"#ComputerB\"/>\n"+
						"</rdf:Description>\n"+
						"<rdf:Description about=\"#CableN\">\n"+
						"<connect:ends rdf:resource=\"#socket5\"/>\n"+
						"<connect:ends>Everything</connect:ends>\n"+
						"</rdf:Description>\n"+
						"<rdf:Description about=\"#Hub\">\n"+
						"<connect:ends rdf:resource=\"#socket1\"/>\n"+
						"<connect:ends rdf:resource=\"#socket2\"/>\n"+
						"<connect:ends rdf:resource=\"#socket3\"/>\n"+
						"<connect:ends rdf:resource=\"#socket4\"/>\n"+
						"<connect:ends rdf:resource=\"#socket5\"/>\n"+
						"</rdf:Description>\n"+
						"</rdf:RDF>",
				),
			factoryBrowser.NewTagSvgTitle().
				Text("Network"),

			factoryBrowser.NewTagSvgDesc().
				Text("An example of a computer network based on a hub."),

			factoryBrowser.NewTagSvgStyle().
				Style("svg {\n"+
					"/* Default styles to be inherited */\n"+
					"fill: white;\n"+
					"stroke: black;\n"+
					"}\n"+
					"text { fill: black; stroke: none; }\n"+
					"path { fill: none; }",
				),

			// Define symbols used in the SVG
			factoryBrowser.NewTagSvgDefs().Append(

				// hubPlug symbol. Used by hub symbol
				factoryBrowser.NewTagSvgSymbol().Id("hubPlug").Append(

					factoryBrowser.NewTagSvgDesc().Text("A 10BaseT/100baseTX socket"),

					//factoryBrowser.NewTagSvgPath().D("M0,10 h5 v-9 h12 v9 h5 v16 h-22 z"),
					factoryBrowser.NewTagSvgPath().D(factoryBrowser.NewPath().M(0, 10).Hd(5).Vd(-9).Hd(12).Vd(9).Hd(5).Vd(16).Hd(-22).Z()),
				),

				// hub symbol
				factoryBrowser.NewTagSvgSymbol().Id("hub").Append(

					factoryBrowser.NewTagSvgDesc().Text("A typical 10BaseT/100BaseTX network hub"),
					factoryBrowser.NewTagSvgText().X(0).Y(15).Text("Hub"),
					factoryBrowser.NewTagSvgG().Transform(factoryBrowser.NewTransform().Translate(0, 20)).Append(

						factoryBrowser.NewTagSvgRect().Width(253).Height(84),
						factoryBrowser.NewTagSvgRect().Width(229).Height(44).X(12).Y(10),
						factoryBrowser.NewTagSvgCircle().Fill(factoryColor.NewRed()).Cx(227).Cy(71).R(7),

						//five groups each using the defined socket
						factoryBrowser.NewTagSvgG().Id("sock1et").Transform(factoryBrowser.NewTransform().Translate(25, 20)).Append(

							factoryBrowser.NewTagSvgTitle().Text("Socket 1"),
							factoryBrowser.NewTagSvgUse().HRef("#hubPlug"),
						),

						factoryBrowser.NewTagSvgG().Id("socket2").Transform(factoryBrowser.NewTransform().Translate(70, 20)).Append(

							factoryBrowser.NewTagSvgTitle().Text("Socket 2"),
							factoryBrowser.NewTagSvgUse().HRef("#hubPlug"),
						),

						factoryBrowser.NewTagSvgG().Id("socket3").Transform(factoryBrowser.NewTransform().Translate(115, 20)).Append(

							factoryBrowser.NewTagSvgTitle().Text("Socket 3"),
							factoryBrowser.NewTagSvgUse().HRef("#hubPlug"),
						),

						factoryBrowser.NewTagSvgG().Id("socket4").Transform(factoryBrowser.NewTransform().Translate(160, 20)).Append(

							factoryBrowser.NewTagSvgTitle().Text("Socket 4"),
							factoryBrowser.NewTagSvgUse().HRef("#hubPlug"),
						),

						factoryBrowser.NewTagSvgG().Id("socket5").Transform(factoryBrowser.NewTransform().Translate(205, 20)).Append(

							factoryBrowser.NewTagSvgTitle().Text("Socket 5"),
							factoryBrowser.NewTagSvgUse().HRef("#hubPlug"),
						),
					),
				),

				// computer symbol
				factoryBrowser.NewTagSvgSymbol().Id("computer").Append(

					factoryBrowser.NewTagSvgDesc().Text("A common desktop PC"),
					factoryBrowser.NewTagSvgG().Id("monitorStand").Transform(factoryBrowser.NewTransform().Translate(40, 121)).Append(

						factoryBrowser.NewTagSvgTitle().Text("Monitor stand"),
						factoryBrowser.NewTagSvgDesc().Text("One of those cool swivelling monitor stands that sit under the monitor"),
						//factoryBrowser.NewTagSvgPath().D("m0,0 S 10 10 40 12"),
						factoryBrowser.NewTagSvgPath().D(factoryBrowser.NewPath().Md(0, 0).S(10, 10, 40, 12)),
						//factoryBrowser.NewTagSvgPath().D("m80,0 S 70 10 40 12"),
						factoryBrowser.NewTagSvgPath().D(factoryBrowser.NewPath().Md(80, 0).S(70, 10, 40, 12)),
						//factoryBrowser.NewTagSvgPath().D("m0,20 L 10 10 S 40 12 70 10 L 80 20z"),
						factoryBrowser.NewTagSvgPath().D(factoryBrowser.NewPath().Md(0, 20).L(10, 10).S(40, 12, 70, 10).L(80, 20).Z()),
					),
					factoryBrowser.NewTagSvgG().Id("monitor").Append(

						factoryBrowser.NewTagSvgTitle().Text("Monitor"),
						factoryBrowser.NewTagSvgDesc().Text("A very fancy monitor"),
						factoryBrowser.NewTagSvgRect().Width(160).Height(120),
						factoryBrowser.NewTagSvgRect().Fill(factoryColor.NewLightgrey()).Width(138).Height(95).X(11).Y(12),
					),

					factoryBrowser.NewTagSvgG().Id("processor").Transform(factoryBrowser.NewTransform().Translate(0, 142)).Append(

						factoryBrowser.NewTagSvgTitle().Text("The computer"),
						factoryBrowser.NewTagSvgDesc().Text("A desktop computer - broad flat box style"),
						factoryBrowser.NewTagSvgRect().Width(160).Height(60),
						factoryBrowser.NewTagSvgG().Id("discDrive").Transform(factoryBrowser.NewTransform().Translate(70, 8)).Append(

							factoryBrowser.NewTagSvgTitle().Text("disc drive"),
							factoryBrowser.NewTagSvgDesc().Text("A built-in disc drive"),
							factoryBrowser.NewTagSvgRect().Width(58).Height(3).X(12).Y(8),
							factoryBrowser.NewTagSvgRect().Width(8).Height(2).X(12).Y(15),
						),

						factoryBrowser.NewTagSvgCircle().Cx(135).Cy(40).R(5),
					),
				),
			),

			factoryBrowser.NewTagSvgText().X(0).Y(15).Text("Network"),

			// Use the hub symbol.
			factoryBrowser.NewTagSvgG().Id("Hub").Transform(factoryBrowser.NewTransform().Translate(80, 45)).Append(
				factoryBrowser.NewTagSvgTitle().Text("Hub"),
				//factoryBrowser.NewTagSvgUse().HRef("#Hub").Transform("scale(0.75)"),
				factoryBrowser.NewTagSvgUse().HRef("#hub").Transform(factoryBrowser.NewTransform().Scale(0.75, 0.75)),
			),

			// Use the computer symbol.
			//factoryBrowser.NewTagSvgG().Id("ComputerA").Transform("translate(20 170)"),
			factoryBrowser.NewTagSvgG().Id("ComputerA").Transform(factoryBrowser.NewTransform().Translate(20, 170)).Append(

				factoryBrowser.NewTagSvgTitle().Text("Computer A"),
				//factoryBrowser.NewTagSvgUse().HRef("#computer").Transform("scale(0.5)"),
				factoryBrowser.NewTagSvgUse().HRef("#computer").Transform(factoryBrowser.NewTransform().Scale(0.5, 0.5)),
			),

			// Use the same computer symbol.
			//factoryBrowser.NewTagSvgG().Id("ComputerB").Transform("translate(300 170)"),
			factoryBrowser.NewTagSvgG().Id("ComputerB").Transform(factoryBrowser.NewTransform().Translate(300, 170)).Append(

				factoryBrowser.NewTagSvgTitle().Text("Computer B"),
				factoryBrowser.NewTagSvgUse().HRef("#computer").Transform(factoryBrowser.NewTransform().Scale(0.5, 0.5)),
				factoryBrowser.NewTagSvgUse().HRef("#computer").Transform("scale(0.5)"),
			),

			// Draw Cable A.
			//factoryBrowser.NewTagSvgG().Id("CableA").Transform("translate(107 88)"),
			factoryBrowser.NewTagSvgG().Id("CableA").Transform(factoryBrowser.NewTransform().Translate(107, 88)).Append(

				factoryBrowser.NewTagSvgTitle().Text("Cable A"),
				//factoryBrowser.NewTagSvgPath().D("M0,0c100,140 50,140 -8,160"),
				factoryBrowser.NewTagSvgPath().D(factoryBrowser.NewPath().M(0, 0).Cd(100, 140, 50, 140, -8, 160)),
			),

			// Draw Cable B.
			//factoryBrowser.NewTagSvgG().Id("CableB").Transform("translate(142 88)"),
			factoryBrowser.NewTagSvgG().Id("CableB").Transform(factoryBrowser.NewTransform().Translate(142, 88)).Append(

				factoryBrowser.NewTagSvgTitle().Text("Cable B"),
				factoryBrowser.NewTagSvgDesc().Text("10BaseT twisted pair cable"),
				//factoryBrowser.NewTagSvgPath().D("M0,0c100,180 110,160 159,160"),
				factoryBrowser.NewTagSvgPath().D(factoryBrowser.NewPath().M(0, 0).Cd(100, 180, 110, 160, 159, 160)),
			),

			// Draw Cable N.
			factoryBrowser.NewTagSvgG().Id("CableN").Transform(factoryBrowser.NewTransform().Translate(242, 88)).Append(

				factoryBrowser.NewTagSvgTitle().Text("Cable N"),
				factoryBrowser.NewTagSvgDesc().Text("10BaseT twisted pair cable"),
				//factoryBrowser.NewTagSvgPath().D("M0,0c0,-70 20,-50 60,-50"),
				factoryBrowser.NewTagSvgPath().D(factoryBrowser.NewPath().M(0, 0).Cd(0, -70, 20, -50, 60, -50)),
			),
		)

	stage.Append(s1)

	<-done
}
