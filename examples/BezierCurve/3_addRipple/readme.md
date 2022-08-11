# Bezier Curve & Easing Tween

### English:

This example plays with the Bézier curve and easing tween.

### Português:

Este exemplo brinca com a curva de Bézier e easing tween.

### Makefile

```shell
make help         ## This help command
make buildandrun  ## build this example and run local server
make build        ## build main.wasm file to run this example
make server       ## run local server
```

### Local server

[https://localhost/examples/BezierCurve/3_addRipple/](https://localhost/examples/BezierCurve/3_addRipple/)

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
  "github.com/helmutkemper/iotmaker.webassembly/browser/html"
  "github.com/helmutkemper/iotmaker.webassembly/platform/algorithm"
  "github.com/helmutkemper/iotmaker.webassembly/platform/factoryAlgorithm"
  "github.com/helmutkemper/iotmaker.webassembly/platform/factoryColor"
  "github.com/helmutkemper/iotmaker.webassembly/platform/factoryEasingTween"
  "math"
  "time"
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
  bezier.Process(10000)
  bezier.SetNumberOfSegments(2000)

  bezier.GenerateRipple(20, 20)
  for _, point := range *bezier.GetProcessed() {
    AddDotBlue(int(point.X), int(point.Y))
  }

  div := factoryBrowser.NewTagDiv().
    Class("animate").
    AddPointsToEasingTween(bezier).
    SetDeltaX(-15).
    SetDeltaY(-25).
    RotateDelta(-math.Pi / 2)
  stage.Append(div)

  wasm := factoryBrowser.NewTagDiv().Style("position:absolute;font-size:40px;color:#555555").SetXY(200, 400).Html("golang = wasm = fast javascript")
  stage.Append(wasm)

  factoryEasingTween.NewLinear(
    20*time.Second,
    0,
    10000,
    div.EasingTweenWalkingAndRotateIntoPoints,
    -1,
  ).
    SetArgumentsFunc(any(div)).
    SetDoNotReverseMotion()

  done := make(chan struct{}, 0)
  <-done
}

func AddDotBlue(x, y int) {
  canvas.BeginPath().
    FillStyle(factoryColor.NewBlueHalfTransparent()).
    Arc(x, y, 0.5, 0, 2*math.Pi, false).
    Fill()
}
```
