# Bezier Curve

### English:

This example shows how to use the Bézier curve.

### Português:

Este exemplo mostra como usar a curva de Bézier.

### Makefile

```shell
make help         ## This help command
make buildandrun  ## build this example and run local server
make build        ## build main.wasm file to run this example
make server       ## run local server
```

### Local server

[https://localhost/examples/BezierCurve/1_addCurve/](https://localhost/examples/BezierCurve/1_addCurve/)

### Code:

##### HTML

```html
<html>
<head>
  <meta charset="utf-8"/>
  <style>
      body {
          margin: 0 !important;
          padding: 0 !important;
      }
  </style>
  <script src="../support/wasm_exec.js"></script>
  <script>
    const go = new Go();
    WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then((result) => {
      go.run(result.instance);
    });
  </script>
</head>
<body>
</body>
</html>
```

##### Golang

```go
//go:build js

package main

import (
  "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
  "github.com/helmutkemper/iotmaker.webassembly/browser/factoryFontFamily"
  "github.com/helmutkemper/iotmaker.webassembly/browser/html"
  "github.com/helmutkemper/iotmaker.webassembly/platform/algorithm"
  "github.com/helmutkemper/iotmaker.webassembly/platform/factoryAlgorithm"
  "github.com/helmutkemper/iotmaker.webassembly/platform/factoryColor"
  "math"
  "strconv"
)

var canvas *html.TagCanvas

func main() {

  var stage = factoryBrowser.NewStage()

  canvas = factoryBrowser.NewTagCanvas(stage.GetWidth(), stage.GetHeight())
  stage.Append(canvas)

  var bezier = factoryAlgorithm.NewBezierCurve()

  border := 50.0
  wight := 400.0
  height := 400.0

  // E.g.: P0 (1,0) = (1*wight,0*height)
  // E.g.: P1 (2,0) = (2*wight,0*height)
  // E.g.: P2 (2,1) = (2*wight,1*height)
  //
  //     (0,0)            (1,0)            (2,0)
  //       +----------------+----------------+
  //       | P7            P0             P1 |
  //       |                                 |
  //       |                                 |
  //       |                                 |
  // (0,1) + P6                           P2 + (2,1)
  //       |                                 |
  //       |                                 |
  //       |                                 |
  //       | P5            P4             P3 |
  //       +----------------+----------------+
  //     (0,2)            (1,2)            (2,2)

  bezier.Add(algorithm.Point{X: 1*wight + border, Y: 0*height + border})
  bezier.Add(algorithm.Point{X: 2*wight + border, Y: 0*height + border})
  bezier.Add(algorithm.Point{X: 2*wight + border, Y: 1*height + border})
  bezier.Add(algorithm.Point{X: 2*wight + border, Y: 2*height + border})
  bezier.Add(algorithm.Point{X: 1*wight + border, Y: 2*height + border})
  bezier.Add(algorithm.Point{X: 0*wight + border, Y: 2*height + border})
  bezier.Add(algorithm.Point{X: 0*wight + border, Y: 1*height + border})
  bezier.Add(algorithm.Point{X: 0*wight + border, Y: 0*height + border})
  bezier.Add(algorithm.Point{X: 1*wight + border, Y: 0*height + border})
  bezier.Process(1000)

  for v, point := range *bezier.GetOriginal() {
    AddRedPointer(int(point.X), int(point.Y))
    AddIndex(int(point.X), int(point.Y), v)
  }

  for _, point := range *bezier.GetProcessed() {
    AddDotBlue(int(point.X), int(point.Y))
  }

  done := make(chan struct{}, 0)
  <-done
}

func AddDotBlue(x, y int) {
  canvas.BeginPath().
    FillStyle(factoryColor.NewBlueHalfTransparent()).
    Arc(x, y, 0.5, 0, 2*math.Pi, false).
    Fill()
}

func AddRedPointer(x, y int) {
  canvas.BeginPath().
    FillStyle(factoryColor.NewRedHalfTransparent()).
    Arc(x, y, 3, 0, 2*math.Pi, false).
    Fill()
}

func AddIndex(x, y, i int) {
  xStr := strconv.FormatInt(int64(x), 10)
  yStr := strconv.FormatInt(int64(y), 10)
  iStr := strconv.FormatInt(int64(i), 10)

  if i == 8 {
    y += 16
  }

  x += 5
  y += 20
  var font html.Font
  font.Family = factoryFontFamily.NewArial()
  font.Size = 17

  canvas.BeginPath().
    Font(font).
    FillStyle(factoryColor.NewRed()).
    FillText(
      "#"+iStr,
      x,
      y,
      300,
    )

  font.Size = 12
  canvas.BeginPath().
    Font(font).
    FillStyle(factoryColor.NewRed()).
    FillText(
      "("+xStr+", "+yStr+")",
      x+20,
      y,
      300,
    )
}
```
