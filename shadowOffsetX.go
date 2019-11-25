package iotmaker_platform_webbrowser

// en: Sets or returns the horizontal distance of the shadow from the shape
//     The shadowOffsetX property sets or returns the horizontal distance of the shadow from the shape.
//     shadowOffsetX = 0 indicates that the shadow is right behind the shape.
//     shadowOffsetX = 20 indicates that the shadow starts 20 pixels to the right (from the shape's left position).
//     shadowOffsetX = -20 indicates that the shadow starts 20 pixels to the left (from the shape's left position).
//     Tip: To adjust the vertical distance of the shadow from the shape, use the shadowOffsetY property.
//     Default value: 0
//     JavaScript syntax: context.shadowOffsetX = number;
func (el *Canvas) ShadowOffsetX(value int) {
	el.SelfContext.Set("shadowOffsetX", value)
}
