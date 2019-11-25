package iotmaker_platform_webbrowser

import (
	"image/color"
	"strconv"
)

func RGBAToJs(color color.RGBA) string {
	return "rgba( " + strconv.Itoa(int(color.R)) + ", " + strconv.Itoa(int(color.G)) + ", " + strconv.Itoa(int(color.B)) + ", " + strconv.Itoa(int(color.A)) + " )"
}
