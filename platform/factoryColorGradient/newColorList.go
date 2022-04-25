package factoryColorGradient

import "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/abstractType/gradient"

func NewColorList(list ...gradient.ColorStop) []gradient.ColorStop {
	return list
}
