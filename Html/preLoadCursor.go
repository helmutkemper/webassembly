package Html

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
	Img Image
	id  string
}

var PreLoadMouseList []PreLoadImage

func PreLoadCursor(parent interface{}, path string, list map[string]string) {

	if len(PreLoadMouseList) == 0 {
		PreLoadMouseList = make([]PreLoadImage, len(list))
	}

	key := 0
	for id, img := range list {

		PreLoadMouseList[key] = PreLoadImage{
			Img: Image{
				SetProperty: map[string]interface{}{
					"id":  "visibleMousePointer",
					"src": path + "/" + img,
				},
				WaitLoad: true,
			},
			id: id,
		}
		PreLoadMouseList[key].Img.SetParent(parent)
		PreLoadMouseList[key].Img.Create()

		key += 1
	}
}
