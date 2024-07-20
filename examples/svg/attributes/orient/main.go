// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/orient
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/orient
//
//  <svg viewBox="0 0 100 100" xmlns="http://www.w3.org/2000/svg">
//    <defs>
//      <marker id="arrow" viewBox="0 0 10 10" refX="5" refY="5"
//          markerWidth="6" markerHeight="6"
//          orient="auto-start-reverse">
//        <path d="M 0 0 L 10 5 L 0 10 z" />
//      </marker>
//
//      <marker id="dataArrow" viewBox="0 0 10 10" refX="5" refY="5"
//          markerWidth="6" markerHeight="6"
//          orient="-65deg">
//        <path d="M 0 0 L 10 5 L 0 10 z" fill="red" />
//      </marker>
//    </defs>
//
//    <polyline points="10,10 10,90 90,90" fill="none" stroke="black"
//        marker-start="url(#arrow)" marker-end="url(#arrow)" />
//
//    <polyline points="15,80 29,50 43,60 57,30 71,40 85,15" fill="none" stroke="grey"
//        marker-start="url(#dataArrow)" marker-mid="url(#dataArrow)"
//        marker-end="url(#dataArrow)" />
//  </svg>

//go:build js

package main

import (
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/platform/factoryColor"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()
	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 100, 100}).Append(
		factoryBrowser.NewTagSvgDefs().Append(
			factoryBrowser.NewTagSvgMarker().Id("arrow").ViewBox([]float64{0, 0, 10, 10}).RefX(5).RefY(5).MarkerWidth(6).MarkerHeight(6).Orient(html.KSvgOrientAutoStartReverse).Append(
				factoryBrowser.NewTagSvgPath().D(factoryBrowser.NewPath().M(0, 0).L(10, 5).L(0, 10).Z()),
			),
			factoryBrowser.NewTagSvgMarker().Id("dataArrow").ViewBox([]float64{0, 0, 10, 10}).RefX(5).RefY(5).MarkerWidth(6).MarkerHeight(6).Orient(html.Degrees(-65)).Append(
				factoryBrowser.NewTagSvgPath().D(factoryBrowser.NewPath().M(0, 0).L(10, 5).L(0, 10).Z()).Fill(factoryColor.NewRed()),
			),
		),

		factoryBrowser.NewTagSvgPolyline().Points([][]float64{{10, 10}, {10, 90}, {90, 90}}).Fill(nil).Stroke(factoryColor.NewBlack()).MarkerStart("url(#arrow)").MarkerEnd("url(#arrow)"),
		factoryBrowser.NewTagSvgPolyline().Points([][]float64{{15, 80}, {29, 50}, {43, 60}, {57, 30}, {71, 40}, {85, 15}}).Fill(nil).Stroke(factoryColor.NewGrey()).MarkerStart("url(#dataArrow)").MarkerMid("url(#dataArrow)").MarkerEnd("url(#dataArrow)"),
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
