package factoryBrowserStage

import (
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.interfaces/iStage"
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.IDraw"
	iotmaker_platform_coordinate "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.coordinate"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/abstractType/image"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/basic"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/factoryImage"
	"github.com/helmutkemper/iotmaker.webassembly/html"
)

var KTemplarianList = map[string]string{
	"default":                 "cursor-default-outline.svg",
	"arrowBottomLeft":         "arrow-bottom-left.svg",
	"arrowBottomRight":        "arrow-bottom-right.svg",
	"arrowDown":               "arrow-down.svg",
	"arrowCollapse":           "arrow-collapse.svg",
	"arrowCollapseAll":        "arrow-collapse-all.svg",
	"arrowCollapseDown":       "arrow-collapse-down.svg",
	"arrowCollapseHorizontal": "arrow-collapse-horizontal.svg",
	"arrowCollapseLeft":       "arrow-collapse-left.svg",
	"arrowCollapseRight":      "arrow-collapse-right.svg",
	"arrowCollapseUp":         "arrow-collapse-up.svg",
	"arrowCollapseVertical":   "arrow-collapse-vertical.svg",
	"arrowExpand":             "arrow-expand.svg",
	"arrowExpandAll":          "arrow-expand-all.svg",
	"arrowExpandDown":         "arrow-expand-down.svg",
	"arrowExpandHorizontal":   "arrow-expand-horizontal.svg",
	"arrowExpandLeft":         "arrow-expand-left.svg",
	"arrowExpandRight":        "arrow-expand-right.svg",
	"arrowExpandUp":           "arrow-expand-up.svg",
	"arrowExpandVertical":     "arrow-expand-vertical.svg",
	"arrowHorizontalLock":     "arrow-horizontal-lock.svg",
	"arrowLeft":               "arrow-left.svg",
	"arrowLeftRight":          "arrow-left-right.svg",
	"arrowRight":              "arrow-right.svg",
	"arrowTopLeft":            "arrow-top-left.svg",
	"arrowTopLeftBottomRight": "arrow-top-left-bottom-right.svg",
	"arrowTopRight":           "arrow-top-right.svg",
	"arrowTopRightBottomLeft": "arrow-top-right-bottom-left.svg",
	"arrowUp":                 "arrow-up.svg",
	"arrowUpDown":             "arrow-up-down.svg",
}

type PreLoadImage struct {
	Img html.Image
	id  string
}

var PreLoadMouseList []PreLoadImage
var cursor PreLoadImage

var cursorWidth = 24
var cursorHeight = 24

var imageCursor *image.Image

func PreLoadCursor(
	parent interface{},
	stage iStage.IStage,
	htmlPlatform iotmaker_platform_IDraw.IHtml,
	canvas,
	scratchPad iotmaker_platform_IDraw.IDraw,
	path string,
	list map[string]string,
	density interface{},
	iDensity iotmaker_platform_coordinate.IDensity) {

	if len(PreLoadMouseList) == 0 {
		PreLoadMouseList = make([]PreLoadImage, len(list))
	}

	key := 0
	for id, img := range list {

		PreLoadMouseList[key] = PreLoadImage{
			Img: htmlPlatform.NewImage(
				parent,
				map[string]interface{}{
					"id":  "visibleMousePointer",
					"src": path + "/" + img,
				},
				true,
			),
			id: id,
		}
		PreLoadMouseList[key].Img.SetParent(parent)
		PreLoadMouseList[key].Img.Create()

		if id == "default" {
			cursor = PreLoadMouseList[key]
		}

		key += 1
	}

	densityCalc := iDensity
	densityCalc.SetDensityFactor(density)

	densityCalc.SetInt(cursorWidth)
	cursorWidth = densityCalc.Int()

	densityCalc.SetInt(cursorHeight)
	cursorHeight = densityCalc.Int()

	imageCursor = factoryImage.NewImageWithDelta(
		"__mouse__cursor__",
		stage,
		canvas,
		scratchPad,
		cursor.Img.Get(),
		-1000,
		-1000,
		cursorWidth,
		cursorHeight,
		-8,
		-4,
		density,
		iDensity,
	)
	imageCursor.SetDragMode(basic.KDragModeAlways)
	imageCursor.DragStart()
}
