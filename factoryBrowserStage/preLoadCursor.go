package factoryBrowserStage

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.IDraw"
	iotmaker_platform_coordinate "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.coordinate"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/Html"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/abstractType/image"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/factoryImage"
)

const KTemplarianPath = "./fonts/Templarian/MaterialDesign/svg"

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
	Img Html.Image
	id  string
}

var PreLoadMouseList []PreLoadImage
var cursor PreLoadImage

var cursorWidth = 24.0
var cursorHeight = 24.0

var imageCursor *image.Image

func PreLoadCursor(
	parent interface{},
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

	densityCalc.Set(cursorWidth)
	cursorWidth = densityCalc.Float64()

	densityCalc.Set(cursorHeight)
	cursorHeight = densityCalc.Float64()

	imageCursor = factoryImage.NewImage(
		canvas,
		scratchPad,
		cursor.Img.Get(),
		-1000,
		-1000,
		cursorWidth,
		cursorHeight,
		density,
		iDensity,
	)
	imageCursor.DragStart()
}
