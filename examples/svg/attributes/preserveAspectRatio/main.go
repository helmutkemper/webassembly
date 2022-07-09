// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/preserveAspectRatio
//
//   Notes:
//     * The CSS is inside the example HTML file.
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/preserveAspectRatio
//
//   Notas:
//     * O CSS está dentro do arquivo HTML de exemplo.
//
// CSS:
//
//  path {
//    fill: yellow;
//    stroke: black;
//    stroke-width: 8px;
//    stroke-linecap: round;
//    stroke-linejoin: round;
//    pointer-events: none;
//  }
//
//  rect:hover, rect:active {
//    outline: 1px solid red;
//  }
//
// HTML:
//
//  <svg viewBox="-1 -1 162 92" xmlns="http://www.w3.org/2000/svg">
//    <defs>
//       <path id="smiley" d="M50,10 A40,40,1,1,1,50,90 A40,40,1,1,1,50,10 M30,40 Q36,35,42,40 M58,40 Q64,35,70,40 M30,60 Q50,75,70,60 Q50,75,30,60" />
//    </defs>
//
//    <!-- (width>height) meet -->
//    <rect x="0" y="0" width="20" height="10">
//      <title>xMidYMid meet</title>
//    </rect>
//    <svg viewBox="0 0 100 100" width="20" height="10"
//         preserveAspectRatio="xMidYMid meet" x="0" y="0">
//      <use href="#smiley" />
//    </svg>
//
//    <rect x="25" y="0" width="20" height="10">
//      <title>xMinYMid meet</title>
//    </rect>
//    <svg viewBox="0 0 100 100" width="20" height="10"
//         preserveAspectRatio="xMinYMid meet" x="25" y="0">
//      <use href="#smiley" />
//    </svg>
//
//    <rect x="50" y="0" width="20" height="10">
//      <title>xMaxYMid meet</title>
//    </rect>
//    <svg viewBox="0 0 100 100" width="20" height="10"
//         preserveAspectRatio="xMaxYMid meet" x="50" y="0">
//      <use href="#smiley" />
//    </svg>
//
//    <!-- (width>height) slice -->
//    <rect x="0" y="15" width="20" height="10">
//      <title>xMidYMin slice</title>
//    </rect>
//    <svg viewBox="0 0 100 100" width="20" height="10"
//         preserveAspectRatio="xMidYMin slice" x="0" y="15">
//      <use href="#smiley" />
//    </svg>
//
//    <rect x="25" y="15" width="20" height="10">
//      <title>xMidYMid slice</title>
//    </rect>
//    <svg viewBox="0 0 100 100" width="20" height="10"
//         preserveAspectRatio="xMidYMid slice" x="25" y="15">
//      <use href="#smiley" />
//    </svg>
//
//    <rect x="50" y="15" width="20" height="10">
//      <title>xMidYMax slice</title>
//    </rect>
//    <svg viewBox="0 0 100 100" width="20" height="10"
//         preserveAspectRatio="xMidYMax slice" x="50" y="15">
//      <use href="#smiley" />
//    </svg>
//
//    <!-- (width<height) meet -->
//    <rect x="75" y="0" width="10" height="25">
//      <title>xMidYMin meet</title>
//    </rect>
//    <svg viewBox="0 0 100 100" width="10" height="25"
//         preserveAspectRatio="xMidYMin meet" x="75" y="0">
//      <use href="#smiley" />
//    </svg>
//
//    <rect x="90" y="0" width="10" height="25">
//      <title>xMidYMid meet</title>
//    </rect>
//    <svg viewBox="0 0 100 100" width="10" height="25"
//         preserveAspectRatio="xMidYMid meet" x="90" y="0">
//      <use href="#smiley" />
//    </svg>
//
//    <rect x="105" y="0" width="10" height="25">
//      <title>xMidYMax meet</title>
//    </rect>
//    <svg viewBox="0 0 100 100" width="10" height="25"
//         preserveAspectRatio="xMidYMax meet" x="105" y="0">
//      <use href="#smiley" />
//    </svg>
//
//    <!-- (width<height) slice -->
//    <rect x="120" y="0" width="10" height="25">
//      <title>xMinYMid slice</title>
//    </rect>
//    <svg viewBox="0 0 100 100" width="10" height="25"
//         preserveAspectRatio="xMinYMid slice" x="120" y="0">
//      <use href="#smiley" />
//    </svg>
//
//    <rect x="135" y="0" width="10" height="25">
//      <title>xMidYMid slice</title>
//    </rect>
//    <svg viewBox="0 0 100 100" width="10" height="25"
//         preserveAspectRatio="xMidYMid slice" x="135" y="0">
//      <use href="#smiley" />
//    </svg>
//
//    <rect x="150" y="0" width="10" height="25">
//      <title>xMaxYMid slice</title>
//    </rect>
//    <svg viewBox="0 0 100 100" width="10" height="25"
//         preserveAspectRatio="xMaxYMid slice" x="150" y="0">
//      <use href="#smiley" />
//    </svg>
//
//    <!-- none -->
//    <rect x="0" y="30" width="160" height="60">
//      <title>none</title>
//    </rect>
//    <svg viewBox="0 0 100 100" width="160" height="60"
//         preserveAspectRatio="none" x="0" y="30">
//      <use href="#smiley" />
//    </svg>
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
	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{-1, -1, 162, 92}).Append(
		factoryBrowser.NewTagSvgDefs().Append(
			factoryBrowser.NewTagSvgPath().Id("smiley").D(factoryBrowser.NewPath().M(50, 10).A(40, 40, 1, 1, 1, 50, 90).A(40, 40, 1, 1, 1, 50, 10).M(30, 40).Q(36, 35, 42, 40).M(58, 40).Q(64, 35, 70, 40).M(30, 60).Q(50, 75, 70, 60).Q(50, 75, 30, 60)),
		),

		// (width>height) meet
		factoryBrowser.NewTagSvgRect().X(0).Y(0).Width(20).Height(10).Append(
			factoryBrowser.NewTagSvgTitle().Title("xMidYMid meet"),
		),
		factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 100, 100}).Width(20).Height(10).
			PreserveAspectRatio(html.KRatioXMidYMid, html.KMeetOrSliceReferenceMeet).X(0).Y(0).Append(
			factoryBrowser.NewTagSvgUse().HRef("#smiley"),
		),

		factoryBrowser.NewTagSvgRect().X(25).Y(0).Width(20).Height(10).Append(
			factoryBrowser.NewTagSvgTitle().Title("xMinYMid meet"),
		),
		factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 100, 100}).Width(20).Height(10).
			PreserveAspectRatio(html.KRatioXMinYMid, html.KMeetOrSliceReferenceMeet).X(25).Y(0).Append(
			factoryBrowser.NewTagSvgUse().HRef("#smiley"),
		),

		factoryBrowser.NewTagSvgRect().X(50).Y(0).Width(20).Height(10).Append(
			factoryBrowser.NewTagSvgTitle().Title("xMaxYMid meet"),
		),
		factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 100, 100}).Width(20).Height(10).
			PreserveAspectRatio(html.KRatioXMaxYMid, html.KMeetOrSliceReferenceMeet).X(50).Y(0).Append(
			factoryBrowser.NewTagSvgUse().HRef("#smiley"),
		),

		// (width>height) slice
		factoryBrowser.NewTagSvgRect().X(0).Y(15).Width(20).Height(10).Append(
			factoryBrowser.NewTagSvgTitle().Title("xMidYMin slice"),
		),
		factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 100, 100}).Width(20).Height(10).
			PreserveAspectRatio(html.KRatioXMidYMin, html.KMeetOrSliceReferenceSlice).X(0).Y(15).Append(
			factoryBrowser.NewTagSvgUse().HRef("#smiley"),
		),

		factoryBrowser.NewTagSvgRect().X(25).Y(15).Width(20).Height(10).Append(
			factoryBrowser.NewTagSvgTitle().Title("xMidYMid slice"),
		),
		factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 100, 100}).Width(20).Height(10).
			PreserveAspectRatio(html.KRatioXMidYMid, html.KMeetOrSliceReferenceSlice).X(25).Y(15).Append(
			factoryBrowser.NewTagSvgUse().HRef("#smiley"),
		),

		factoryBrowser.NewTagSvgRect().X(50).Y(15).Width(20).Height(10).Append(
			factoryBrowser.NewTagSvgTitle().Title("xMidYMax slice"),
		),
		factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 100, 100}).Width(20).Height(10).
			PreserveAspectRatio(html.KRatioXMidYMax, html.KMeetOrSliceReferenceSlice).X(50).Y(15).Append(
			factoryBrowser.NewTagSvgUse().HRef("#smiley"),
		),

		// (width<height) meet
		factoryBrowser.NewTagSvgRect().X(75).Y(0).Width(10).Height(25).Append(
			factoryBrowser.NewTagSvgTitle().Title("xMidYMin meet"),
		),
		factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 100, 100}).Width(10).Height(25).
			PreserveAspectRatio(html.KRatioXMidYMin, html.KMeetOrSliceReferenceMeet).X(75).Y(0).Append(
			factoryBrowser.NewTagSvgUse().HRef("#smiley"),
		),

		factoryBrowser.NewTagSvgRect().X(90).Y(0).Width(10).Height(25).Append(
			factoryBrowser.NewTagSvgTitle().Title("xMidYMid meet"),
		),
		factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 100, 100}).Width(10).Height(25).
			PreserveAspectRatio(html.KRatioXMidYMid, html.KMeetOrSliceReferenceMeet).X(90).Y(0).Append(
			factoryBrowser.NewTagSvgUse().HRef("#smiley"),
		),

		factoryBrowser.NewTagSvgRect().X(105).Y(0).Width(10).Height(25).Append(
			factoryBrowser.NewTagSvgTitle().Title("xMidYMax meet"),
		),
		factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 100, 100}).Width(10).Height(25).
			PreserveAspectRatio(html.KRatioXMidYMax, html.KMeetOrSliceReferenceMeet).X(105).Y(0).Append(
			factoryBrowser.NewTagSvgUse().HRef("#smiley"),
		),

		// (width<height) slice
		factoryBrowser.NewTagSvgRect().X(120).Y(0).Width(10).Height(25).Append(
			factoryBrowser.NewTagSvgTitle().Title("xMinYMid slice"),
		),
		factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 100, 100}).Width(10).Height(25).
			PreserveAspectRatio(html.KRatioXMinYMid, html.KMeetOrSliceReferenceSlice).X(120).Y(0).Append(
			factoryBrowser.NewTagSvgUse().HRef("#smiley"),
		),

		factoryBrowser.NewTagSvgRect().X(135).Y(0).Width(10).Height(25).Append(
			factoryBrowser.NewTagSvgTitle().Title("xMidYMid slice"),
		),
		factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 100, 100}).Width(10).Height(25).
			PreserveAspectRatio(html.KRatioXMidYMid, html.KMeetOrSliceReferenceSlice).X(135).Y(0).Append(
			factoryBrowser.NewTagSvgUse().HRef("#smiley"),
		),

		factoryBrowser.NewTagSvgRect().X(150).Y(0).Width(10).Height(25).Append(
			factoryBrowser.NewTagSvgTitle().Title("xMaxYMid slice"),
		),
		factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 100, 100}).Width(10).Height(25).
			PreserveAspectRatio(html.KRatioXMaxYMid, html.KMeetOrSliceReferenceSlice).X(150).Y(0).Append(
			factoryBrowser.NewTagSvgUse().HRef("#smiley"),
		),

		// none
		factoryBrowser.NewTagSvgRect().X(0).Y(30).Width(160).Height(60).Append(
			factoryBrowser.NewTagSvgTitle().Title("noone"),
		),
		factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 100, 100}).Width(160).Height(60).
			PreserveAspectRatio(nil, nil).X(0).Y(30).Append(
			factoryBrowser.NewTagSvgUse().HRef("#smiley"),
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
