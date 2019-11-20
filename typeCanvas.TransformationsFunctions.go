package iotmaker_platform_webbrowser

// en: Scales the current drawing bigger or smaller
//     scaleWidth:  Scales the width of the current drawing (1=100%, 0.5=50%, 2=200%, etc.)
//     scaleHeight: Scales the height of the current drawing (1=100%, 0.5=50%, 2=200%, etc.)
//
//     The scale() method scales the current drawing, bigger or smaller.
//     Note: If you scale a drawing, all future drawings will also be scaled. The positioning will also be scaled. If
//     you scale(2,2); drawings will be positioned twice as far from the left and top of the canvas as you specify.
//     JavaScript syntax: context.scale(scalewidth, scaleheight);
//
//     Example:
//     var c = document.getElementById("myCanvas");
//     var ctx = c.getContext("2d");
//     ctx.strokeRect(5, 5, 25, 15);
//     ctx.scale(2, 2);
//     ctx.strokeRect(5, 5, 25, 15);
func (el *Canvas) Scale(scaleWidth, scaleHeight float64) {
	el.SelfContext.Call("scale", scaleWidth, scaleHeight)
}

// en: Rotates the current drawing
//     angle: The rotation angle, in radians.
//            To calculate from degrees to radians: degrees*Math.PI/180.
//            Example: to rotate 5 degrees, specify the following: 5*Math.PI/180
//
//     The rotate() method rotates the current drawing.
//     Note: The rotation will only affect drawings made AFTER the rotation is done.
//     JavaScript syntax: context.rotate(angle);
//
//     Example:
//     var c = document.getElementById("my Canvas");
//     var ctx = c.getContext("2d");
//     ctx.rotate(20 * Math.PI / 180);
//     ctx.fillRect(50, 20, 100, 50);
func (el *Canvas) Rotate(angle float64) {
	el.SelfContext.Call("rotate", angle)
}

// en: Remaps the (0,0) position on the canvas
//     x: The value to add to horizontal (x) coordinates
//     y: The value to add to vertical (y) coordinates
//
//     The translate() method remaps the (0,0) position on the canvas.
//     Note: When you call a method such as fillRect() after translate(), the value is added to the x- and y-coordinate
//     values.
//     JavaScript syntax: context.translate(x, y);
//
//     Example:
//     var c = document.getElementById("myCanvas");
//     var ctx = c.getContext("2d");
//     ctx.fillRect(10, 10, 100, 50);
//     ctx.translate(70, 70);
//     ctx.fillRect(10, 10, 100, 50);
func (el *Canvas) Translate(x, y float64) {
	el.SelfContext.Call("translate", x, y)
}

// en: Replaces the current transformation matrix for the drawing
//     a: Scales the drawing horizontally
//     b: Skew the the drawing horizontally
//     c: Skew the the drawing vertically
//     d: Scales the drawing vertically
//     e: Moves the the drawing horizontally
//     f: Moves the the drawing vertically
//
//     Each object on the canvas has a current transformation matrix.
//     The transform() method replaces the current transformation matrix. It multiplies the current transformation
//     matrix with the matrix described by:
//
//     a | c | e
//    -----------
//     b | d | f
//    -----------
//     0 | 0 | 1
//
//     In other words, the transform() method lets you scale, rotate, move, and skew the current context.
//     Note: The transformation will only affect drawings made after the transform() method is called.
//     Note: The transform() method behaves relatively to other transformations made by rotate(), scale(), translate(),
//     or transform(). Example: If you already have set your drawing to scale by two, and the transform() method scales
//     your drawings by two, your drawings will now scale by four.
//     Tip: Check out the setTransform() method, which does not behave relatively to other transformations.
//     JavaScript syntax: context.transform(a, b, c, d, e, f);
//
//     Example:
//     var c = document.getElementById("myCanvas");
//     var ctx = c.getContext("2d");
//     ctx.fillStyle = "yellow";
//     ctx.fillRect(0, 0, 250, 100)
//     ctx.transform(1, 0.5, -0.5, 1, 30, 10);
//     ctx.fillStyle = "red";
//     ctx.fillRect(0, 0, 250, 100);
//     ctx.transform(1, 0.5, -0.5, 1, 30, 10);
//     ctx.fillStyle = "blue";
//     ctx.fillRect(0, 0, 250, 100);
func (el *Canvas) Transform(a, b, c, d, e, f float64) {
	el.SelfContext.Call("transform", a, b, c, d, e, f)
}

// en: Resets the current transform to the identity matrix. Then runs transform()
//     a: Scales the drawings horizontally
//     b: Skews the drawings horizontally
//     c: Skews the drawings vertically
//     d: Scales the drawings vertically
//     e: Moves the the drawings horizontally
//     f: Moves the the drawings vertically
//
//     Each object on the canvas has a current transformation matrix.
//     The setTransform() method resets the current transform to the identity matrix, and then runs transform() with the
//     same arguments.
//     In other words, the setTransform() method lets you scale, rotate, move, and skew the current context.
//     Note: The transformation will only affect drawings made after the setTransform method is called.
//     JavaScript syntax: context.setTransform(a, b, c, d, e, f);
//
//     Example:
//     var c = document.getElementById("myCanvas");
//     var ctx = c.getContext("2d");
//     ctx.fillStyle = "yellow";
//     ctx.fillRect(0, 0, 250, 100)
//     ctx.setTransform(1, 0.5, -0.5, 1, 30, 10);
//     ctx.fillStyle = "red";
//     ctx.fillRect(0, 0, 250, 100);
//     ctx.setTransform(1, 0.5, -0.5, 1, 30, 10);
//     ctx.fillStyle = "blue";
//     ctx.fillRect(0, 0, 250, 100);
func (el *Canvas) SetTransform(a, b, c, d, e, f float64) {
	el.SelfContext.Call("setTransform", a, b, c, d, e, f)
}
