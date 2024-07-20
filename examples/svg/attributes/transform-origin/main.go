// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/transform-origin
//
//   Notes:
//     * The CSS is inside the example HTML file.
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/transform-origin
//
//   Notas:
//     * O CSS está dentro do arquivo HTML de exemplo.
//
// CSS:
//
//  h4 {
//    font-family: sans-serif;
//  }
//
//  figure {
//    border: thin #c0c0c0 solid;
//    display: inline-flex;
//    flex-flow: column;
//    padding: 5px;
//    max-width: 200px;
//    margin: auto;
//  }
//
//  figcaption {
//    margin-top: 5px;
//    background-color: #222;
//    color: #fff;
//    font: smaller sans-serif;
//    padding: 3px;
//    text-align: center;
//  }
//
// HTML:
//
//  <h4>Reference image</h4>
//
//  <div>
//    <figure>
//      <img src="reference.png" alt="PNG reference image"/>
//      <figcaption>Figure 1. PNG reference image. The images following this should look exactly the same as this.</figcaption>
//    </figure>
//  </div>
//
//  <div>
//    <figure>
//      <svg xmlns="http://www.w3.org/2000/svg" width="200" height="200" viewBox="0 0 200 200">
//        <circle cx="100" cy="100" r="100" stroke="none" fill="black"/>
//        <line x1="100" y1="0" x2="100" y2="200" stroke="rebeccapurple" stroke-width="2"/>
//        <line x1="0" y1="100" x2="200" y2="100" stroke="rebeccapurple" stroke-width="2"/>
//
//        <circle cx="100" cy="100" r="75" stroke="none" fill="blue"/>
//        <line x1="100" y1="25" x2="100" y2="175" stroke="rebeccapurple" stroke-width="1.5"/>
//        <line x1="25" y1="100" x2="175" y2="100" stroke="rebeccapurple" stroke-width="1.5"/>
//
//        <circle cx="100" cy="100" r="50" stroke="none" fill="red"/>
//        <line x1="100" y1="50" x2="100" y2="150" stroke="rebeccapurple" stroke-width="1"/>
//        <line x1="50" y1="100" x2="150" y2="100" stroke="rebeccapurple" stroke-width="1"/>
//
//        <circle cx="100" cy="100" r="25" stroke="none" fill="yellow"/>
//        <line x1="100" y1="75" x2="100" y2="125" stroke="rebeccapurple" stroke-width="0.5"/>
//        <line x1="75" y1="100" x2="125" y2="100" stroke="rebeccapurple" stroke-width="0.5"/>
//      </svg>
//      <figcaption>Figure 2. SVG reference image. The images following this should look exactly the same as this.</figcaption>
//    </figure>
//  </div>
//
//  <h4>Transformation with transform-origin</h4>
//
//  <div>
//    <figure>
//      <svg xmlns="http://www.w3.org/2000/svg" width="200" height="200" viewBox="0 0 200 200">
//        <defs>
//          <g id="target-g-1">
//            <circle cx="100" cy="100" r="100" stroke="none"/>
//            <line x1="100" y1="0" x2="100" y2="200" stroke="rebeccapurple" stroke-width="2"/>
//            <line x1="0" y1="100" x2="200" y2="100" stroke="rebeccapurple" stroke-width="2"/>
//          </g>
//        </defs>
//
//        <use href="#target-g-1" fill="black"/>
//        <use href="#target-g-1" fill="blue"
//            transform="scale(0.75 0.75)"
//            transform-origin="100 100"/>
//
//        <svg xmlns="http://www.w3.org/2000/svg" x="0" y="0" width="200" height="200" viewBox="0 0 200 200">
//          <use href="#target-g-1" fill="red"
//            transform="scale(0.5 0.5)"
//            transform-origin="100 100"/>
//          <use href="#target-g-1" fill="yellow"
//            transform="scale(0.25 0.25)"
//            transform-origin="100 100"/>
//        </svg>
//      </svg>
//
//      <figcaption>Figure 3. transform-origin used.
//        This image should look exactly the same as the reference image in Figure 2.</figcaption>
//    </figure>
//  </div>
//
//  <h4>Transformation without transform-origin</h4>
//
//  <div>
//    <figure>
//      <svg xmlns="http://www.w3.org/2000/svg" width="200" height="200" viewBox="0 0 200 200">
//        <defs>
//          <g id="target-g-1">
//            <circle cx="100" cy="100" r="100" stroke="none"/>
//            <line x1="100" y1="0" x2="100" y2="200" stroke="rebeccapurple" stroke-width="2"/>
//            <line x1="0" y1="100" x2="200" y2="100" stroke="rebeccapurple" stroke-width="2"/>
//          </g>
//        </defs>
//
//        <use href="#target-g-1" fill="black"/>
//        <use href="#target-g-1" fill="blue"
//            transform="translate(100 100) scale(0.75 0.75) translate(-100 -100)"/>
//
//        <svg xmlns="http://www.w3.org/2000/svg" x="0" y="0" width="200" height="200" viewBox="0 0 200 200">
//          <use href="#target-g-1" fill="red"
//              transform="translate(100 100) scale(0.5 0.5) translate(-100 -100)"/>
//          <use href="#target-g-1" fill="yellow"
//              transform="translate(100 100) scale(0.25 0.25) translate(-100 -100)"/>
//        </svg>
//      </svg>
//
//      <figcaption>Figure 4. transform-origin not used.
//        This image should look exactly the same as the reference image in Figure 2.</figcaption>
//    </figure>
//  </div>

//go:build js

package main

import (
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/platform/factoryColor"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	dv1 := factoryBrowser.NewTagDiv().Append(
		factoryBrowser.NewTagH4().Html("Reference image"),

		factoryBrowser.NewTagDiv().Append(
			factoryBrowser.NewTagFigure().Append(
				factoryBrowser.NewTagImg().Src("https://yari-demos.prod.mdn.mozit.cloud/en-US/docs/Web/SVG/Attribute/transform-origin/reference.png", false).Alt("PNG reference image"),
				factoryBrowser.NewTagFigCaption().Html("Figure 1. PNG reference image. The images following this should look exactly the same as this."),
			),
		),

		factoryBrowser.NewTagDiv().Append(
			factoryBrowser.NewTagFigure().Append(
				factoryBrowser.NewTagSvg().Width(200).Height(200).ViewBox([]float64{0, 0, 200, 200}).Append(
					factoryBrowser.NewTagSvgCircle().Cx(100).Cy(100).R(100).Stroke(nil).Fill(factoryColor.NewBlack()),
					factoryBrowser.NewTagSvgLine().X1(100).Y1(0).X2(100).Y2(200).Stroke("rebeccapurple").StrokeWidth(2),
					factoryBrowser.NewTagSvgLine().X1(0).Y1(100).X2(100).Y2(100).Stroke("rebeccapurple").StrokeWidth(2),

					factoryBrowser.NewTagSvgCircle().Cx(100).Cy(100).R(75).Stroke(nil).Fill(factoryColor.NewBlue()),
					factoryBrowser.NewTagSvgLine().X1(100).Y1(25).X2(100).Y2(175).Stroke("rebeccapurple").StrokeWidth(1.5),
					factoryBrowser.NewTagSvgLine().X1(100).Y1(100).X2(175).Y2(100).Stroke("rebeccapurple").StrokeWidth(1.5),

					factoryBrowser.NewTagSvgCircle().Cx(100).Cy(100).R(50).Stroke(nil).Fill(factoryColor.NewRed()),
					factoryBrowser.NewTagSvgLine().X1(100).Y1(50).X2(100).Y2(150).Stroke("rebeccapurple").StrokeWidth(1),
					factoryBrowser.NewTagSvgLine().X1(50).Y1(100).X2(150).Y2(100).Stroke("rebeccapurple").StrokeWidth(1),

					factoryBrowser.NewTagSvgCircle().Cx(100).Cy(100).R(25).Stroke(nil).Fill(factoryColor.NewYellow()),
					factoryBrowser.NewTagSvgLine().X1(100).Y1(75).X2(100).Y2(125).Stroke("rebeccapurple").StrokeWidth(0.5),
					factoryBrowser.NewTagSvgLine().X1(75).Y1(100).X2(125).Y2(100).Stroke("rebeccapurple").StrokeWidth(0.5),
				),

				factoryBrowser.NewTagFigCaption().Html("Figure 2. SVG reference image. The images following this should look exactly the same as this."),
			),
		),

		factoryBrowser.NewTagH4().Html("Transformation with transform-origin"),

		factoryBrowser.NewTagDiv().Append(
			factoryBrowser.NewTagFigure().Append(
				factoryBrowser.NewTagSvg().Width(200).Height(200).ViewBox([]float64{0, 0, 200, 200}).Append(
					factoryBrowser.NewTagSvgDefs().Append(
						factoryBrowser.NewTagSvgG().Id("target-g-1").Append(
							factoryBrowser.NewTagSvgCircle().Cx(100).Cy(100).R(100).Stroke(nil),
							factoryBrowser.NewTagSvgLine().X1(100).Y1(0).X2(100).Y2(200).Stroke("rebeccapurple").StrokeWidth(2),
							factoryBrowser.NewTagSvgLine().X1(0).Y1(100).X2(200).Y2(100).Stroke("rebeccapurple").StrokeWidth(2),
						),
					),

					factoryBrowser.NewTagSvgUse().HRef("#target-g-1").Fill(factoryColor.NewBlack()),
					factoryBrowser.NewTagSvgUse().HRef("#target-g-1").Fill(factoryColor.NewBlue()).Transform(factoryBrowser.NewTransform().Scale(0.75, 0.75)).TransformOrigin([]float64{100, 100}),

					factoryBrowser.NewTagSvg().X(0).Y(0).Width(200).Height(200).ViewBox([]float64{0, 0, 200, 200}).Append(
						factoryBrowser.NewTagSvgUse().HRef("#target-g-1").Fill(factoryColor.NewRed()).Transform(factoryBrowser.NewTransform().Scale(0.5, 0.5)).TransformOrigin([]float64{100, 100}),
						factoryBrowser.NewTagSvgUse().HRef("#target-g-1").Fill(factoryColor.NewYellow()).Transform(factoryBrowser.NewTransform().Scale(0.25, 0.25)).TransformOrigin([]float64{100, 100}),
					),
				),

				factoryBrowser.NewTagFigCaption().Html("Figure 3. transform-origin used.<br>This image should look exactly the same as the reference image in Figure 2."),
			),
		),

		factoryBrowser.NewTagH4().Html("Transformation without transform-origin"),

		factoryBrowser.NewTagDiv().Append(
			factoryBrowser.NewTagFigure().Append(
				factoryBrowser.NewTagSvg().Width(200).Height(200).ViewBox([]float64{0, 0, 200, 200}).Append(
					factoryBrowser.NewTagSvgDefs().Append(
						factoryBrowser.NewTagSvgG().Id("target-g-1").Append(
							factoryBrowser.NewTagSvgCircle().Cx(100).Cy(100).R(100).Stroke(nil),
							factoryBrowser.NewTagSvgLine().X1(100).Y1(0).X2(100).Y2(200).Stroke("rebeccapurple").StrokeWidth(2),
							factoryBrowser.NewTagSvgLine().X1(0).Y1(100).X2(200).Y2(100).Stroke("rebeccapurple").StrokeWidth(2),
						),
					),

					factoryBrowser.NewTagSvgUse().HRef("#target-g-1").Fill(factoryColor.NewBlack()),
					factoryBrowser.NewTagSvgUse().HRef("#target-g-1").Fill(factoryColor.NewBlue()).Transform(factoryBrowser.NewTransform().Translate(100, 100).Scale(0.75, 0.75).Translate(-100, -100)),

					factoryBrowser.NewTagSvg().X(0).Y(0).Width(200).Height(200).ViewBox([]float64{0, 0, 200, 200}).Append(
						factoryBrowser.NewTagSvgUse().HRef("#target-g-1").Fill(factoryColor.NewRed()).Transform(factoryBrowser.NewTransform().Translate(100, 100).Scale(0.5, 0.5).Translate(-100, -100)),
						factoryBrowser.NewTagSvgUse().HRef("#target-g-1").Fill(factoryColor.NewYellow()).Transform(factoryBrowser.NewTransform().Translate(100, 100).Scale(0.25, 0.25).Translate(-100, -100)),
					),
				),

				factoryBrowser.NewTagFigCaption().Html("Figure 4. transform-origin not used.\nThis image should look exactly the same as the reference image in Figure 2."),
			),
		),
	)

	stage.Append(dv1)

	<-done
}
