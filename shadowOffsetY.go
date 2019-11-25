package iotmaker_platform_webbrowser

// en: Sets or returns the vertical distance of the shadow from the shape
//     The shadowOffsetY property sets or returns the vertical distance of the shadow from the shape.
//     shadowOffsety = 0 indicates that the shadow is right behind the shape.
//     shadowOffsetY = 20 indicates that the shadow starts 20 pixels below the shape's top position.
//     shadowOffsetY = -20 indicates that the shadow starts 20 pixels above the shape's top position.
//     Tip: To adjust the horizontal distance of the shadow from the shape, use the shadowOffsetX property.
//     Default value: 0
//     JavaScript syntax: context.shadowOffsetY = number;
func (el *Canvas) ShadowOffsetY(value int) {
	el.SelfContext.Set("shadowOffsetY", value)
}
