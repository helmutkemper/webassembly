package iotmaker_platform_webbrowser

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
	el.selfContext.Set("font", font.String())
}

// en: Sets or returns the current alignment for text content
//
//     The textAlign property sets or returns the current alignment for text content, according to the anchor point.
//     Normally, the text will START in the position specified, however, if you set textAlign="right" and place the text in position 150, it means that the text should END in position 150.
//     Tip: Use the fillText() or the strokeText() method to actually draw and position the text on the canvas.
//     Default value: start
//     JavaScript syntax: context.textAlign = "center | end | left | right | start";
//
//     Example:
//     var c = document.getElementById("myCanvas");
//     var ctx = c.getContext("2d");
//     // Create a red line in position 150
//     ctx.strokeStyle = "red";
//     ctx.moveTo(150, 20);
//     ctx.lineTo(150, 170);
//     ctx.stroke();
//     ctx.font = "15px Arial";
//     // Show the different textAlign values
//     ctx.textAlign = "start";
//     ctx.fillText("textAlign = start", 150, 60);
//     ctx.textAlign = "end";
//     ctx.fillText("textAlign = end", 150, 80);
//     ctx.textAlign = "left";
//     ctx.fillText("textAlign = left", 150, 100);
//     ctx.textAlign = "center";
//     ctx.fillText("textAlign = center", 150, 120);
//     ctx.textAlign = "right";
//     ctx.fillText("textAlign = right", 150, 140);
func (el *Canvas) TextAlign(value CanvasFontAlignRule) {
	el.selfContext.Set("textAlign", value.String())
}

// en: Sets or returns the current text baseline used when drawing text
//     Value:
//          alphabetic:  Default. The text baseline is the normal alphabetic baseline
//          top:         The text baseline is the top of the em square
//          hanging:     The text baseline is the hanging baseline
//          middle:      The text baseline is the middle of the em square
//          ideographic: The text baseline is the ideographic baseline
//          bottom:      The text baseline is the bottom of the bounding box
//
//     The textBaseline property sets or returns the current text baseline used when drawing text.
//     Note: The fillText() and strokeText() methods will use the specified textBaseline value when positioning the text
//     on the canvas.
//     Default value: alphabetic
//     JavaScript syntax: context.textBaseline = "alphabetic|top|hanging|middle|ideographic|bottom";
//
//     Example:
//     var c = document.getElementById("myCanvas");
//     var ctx = c.getContext("2d");
//     //Draw a red line at y=100
//     ctx.strokeStyle = "red";
//     ctx.moveTo(5, 100);
//     ctx.lineTo(395, 100);
//     ctx.stroke();
//     ctx.font = "20px Arial"
//     //Place each word at y=100 with different textBaseline values
//     ctx.textBaseline = "top";
//     ctx.fillText("Top", 5, 100);
//     ctx.textBaseline = "bottom";
//     ctx.fillText("Bottom", 50, 100);
//     ctx.textBaseline = "middle";
//     ctx.fillText("Middle", 120, 100);
//     ctx.textBaseline = "alphabetic";
//     ctx.fillText("Alphabetic", 190, 100);
//     ctx.textBaseline = "hanging";
//     ctx.fillText("Hanging", 290, 100);
func (el *Canvas) TextBaseline(value CanvasTextBaseLineRule) {
	el.selfContext.Set("textBaseline", value.String())
}

// en: Draws "filled" text on the canvas
//     text:     Specifies the text that will be written on the canvas
//     x:        The x coordinate where to start painting the text (relative to the canvas)
//     y:        The y coordinate where to start painting the text (relative to the canvas)
//     maxWidth: Optional. The maximum allowed width of the text, in pixels
//
//     The fillText() method draws filled text on the canvas. The default color of the text is black.
//     Tip: Use the font property to specify font and font size, and use the fillStyle property to render the text in another color/gradient.
//     JavaScript syntax: context.fillText(text, x, y, maxWidth);
//
//     Example:
//     var c = document.getElementById("myCanvas");
//     var ctx = c.getContext("2d");
//     ctx.font = "20px Georgia";
//     ctx.fillText("Hello World!", 10, 50);
//     ctx.font = "30px Verdana";
//     // Create gradient
//     var gradient = ctx.createLinearGradient(0, 0, c.width, 0);
//     gradient.addColorStop("0", "magenta");
//     gradient.addColorStop("0.5", "blue");
//     gradient.addColorStop("1.0", "red");
//     // Fill with gradient
//     ctx.fillStyle = gradient;
//     ctx.fillText("Big smile!", 10, 90);
func (el *Canvas) FillText(text string, x, y, maxWidth float64) {
	el.selfContext.Call("fillText", text, x, y, maxWidth)
}

// en: Draws text on the canvas (no fill)
//     text:     Specifies the text that will be written on the canvas
//     x:        The x coordinate where to start painting the text (relative to the canvas)
//     y:        The y coordinate where to start painting the text (relative to the canvas)
//     maxWidth: Optional. The maximum allowed width of the text, in pixels
//
//     The strokeText() method draws text (with no fill) on the canvas. The default color of the text is black.
//     Tip: Use the font property to specify font and font size, and use the strokeStyle property to render the text in another color/gradient.
//     JavaScript syntax: context.strokeText(text, x, y, maxWidth);
//
//     var c = document.getElementById("myCanvas");
//     var ctx = c.getContext("2d");
//     ctx.font = "20px Georgia";
//     ctx.strokeText("Hello World!", 10, 50);
//     ctx.font = "30px Verdana";
//     // Create gradient
//     var gradient = ctx.createLinearGradient(0, 0, c.width, 0);
//     gradient.addColorStop("0", "magenta");
//     gradient.addColorStop("0.5", "blue");
//     gradient.addColorStop("1.0", "red");
//     // Fill with gradient
//     ctx.strokeStyle = gradient;
//     ctx.strokeText("Big smile!", 10, 90);
func (el *Canvas) StrokeText(text string, x, y, maxWidth float64) {
	el.selfContext.Call("strokeText", text, x, y, maxWidth)
}

// en: Returns an object that contains the width of the specified text
//     text: The text to be measured
//
//     The measureText() method returns an object that contains the width of the specified text, in pixels.
//     Tip: Use this method if you need to know the width of a text, before writing it on the canvas.
//     JavaScript syntax: context.measureText(text).width;
//
//     Example:
//     var c = document.getElementById("myCanvas");
//     var ctx = c.getContext("2d");
//     ctx.font = "30px Arial";
//     var txt = "Hello World"
//     ctx.fillText("width:" + ctx.measureText(txt).width, 10, 50)
//     ctx.fillText(txt, 10, 100);
func (el *Canvas) MeasureText(text string) {
	el.selfContext.Call("measureText", text)
}
