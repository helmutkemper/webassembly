package canvas

// en: Sets or returns the current font properties for text content
//     font-style:            Specifies the font style. Possible values:
//          normal | italic | oblique
//
//     font-variant:          Specifies the font variant. Possible values:
//          normal | small-caps
//
//     font-weight:           Specifies the font weight. Possible values:
//          normal | bold | bolder | lighter | 100 | 200 | 300 | 400 | 500 | 600 | 700 | 800 | 900
//
//     font-size/line-height: Specifies the font size and the line-height, in pixels
//     font-family:           Specifies the font family
//     caption:               Use the font captioned controls (like buttons, drop-downs, etc.)
//     icon:                  Use the font used to label icons
//     menu:                  Use the font used in menus (drop-down menus and menu lists)
//     message-box:           Use the font used in dialog boxes
//     small-caption:         Use the font used for labeling small controls
//     status-bar:            Use the fonts used in window status bar
//
//     The font property sets or returns the current font properties for text content on the canvas.
//     The font property uses the same syntax as the CSS font property.
//     Default value: 10px sans-serif
//     JavaScript syntax: context.font = "italic small-caps bold 12px arial";
//
//     Example:
//     var c = document.getElementById("myCanvas");
//     var ctx = c.getContext("2d");
//     ctx.font = "30px Arial";
//     ctx.fillText("Hello World", 10, 50);
func (el *Canvas) Font(font Font) {
	el.SelfContext.Set("font", font.String())
}
