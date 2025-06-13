package rulesConnection

import (
	"github.com/helmutkemper/webassembly/platform/factoryColor"
	"image/color"
	"reflect"
)

// TypeToColor For each type of data, numerical, string, boolean ... a visual identity color is attributed
func TypeToColor(dataType reflect.Kind) (color color.RGBA) {
	switch dataType {
	case reflect.Bool:
		color = factoryColor.NewGreen()
	case reflect.Int:
		color = factoryColor.NewBlue()
	case reflect.Int8:
		color = factoryColor.NewBlue()
	case reflect.Int16:
		color = factoryColor.NewBlue()
	case reflect.Int32:
		color = factoryColor.NewBlue()
	case reflect.Int64:
		color = factoryColor.NewBlue()
	case reflect.Uint:
		color = factoryColor.NewBlueViolet()
	case reflect.Uint8:
		color = factoryColor.NewBlueViolet()
	case reflect.Uint16:
		color = factoryColor.NewBlueViolet()
	case reflect.Uint32:
		color = factoryColor.NewBlueViolet()
	case reflect.Uint64:
		color = factoryColor.NewBlueViolet()
	case reflect.Uintptr:
		color = factoryColor.NewBlueViolet()
	case reflect.Float32:
		color = factoryColor.NewYellowGreen()
	case reflect.Float64:
		color = factoryColor.NewYellowGreen()
	case reflect.Complex64:
		color = factoryColor.NewAntiqueWhite()
	case reflect.Complex128:
		color = factoryColor.NewAntiqueWhite()
	case reflect.Array:
		color = factoryColor.NewDarkMagenta()
	case reflect.Chan:
	case reflect.Func:
	case reflect.Interface:
	case reflect.Map:
	case reflect.Pointer:
	case reflect.Slice:
		color = factoryColor.NewDarkMagenta()
	case reflect.String:
		color = factoryColor.NewMediumTurquoise()
	case reflect.Struct:
		color = factoryColor.NewGainsboro()
	case reflect.UnsafePointer:
	}

	return
}
