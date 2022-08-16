# canvas: AddColorStopPosition

### English:

This example shows how to use addColorStopPosition on cavas element.

### Português:

Este exemplo mostra como usar addColorStopPosition no elemnto canvas.

### Makefile

```shell
make help         ## This help command
make buildandrun  ## build this example and run local server
make build        ## build main.wasm file to run this example
make server       ## run local server
```

### Local server

[https://localhost/examples/canvas/addColorStopPosition/](https://localhost/examples/canvas/addColorStopPosition/)

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
    <script src="../../support/wasm_exec.js"></script>
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
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryFontStyle"
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryFontVariant"
	"github.com/helmutkemper/iotmaker.webassembly/browser/html"
	"github.com/helmutkemper/iotmaker.webassembly/platform/factoryColor"
)

var canvas *html.TagCanvas

func main() {

	var fontA html.Font
	fontA.Family = factoryFontFamily.NewArial()
	fontA.Variant = factoryFontVariant.NewSmallCaps()
	fontA.Style = factoryFontStyle.NewItalic()
	fontA.Size = 20

	var fontB html.Font
	fontB.Family = factoryFontFamily.NewVerdana()
	fontB.Size = 35

	canvas = factoryBrowser.NewTagCanvas(800, 600).
		Font(fontA).
		FillText("Hello World!", 10, 50, 300).
		CreateLinearGradient(0, 0, 160, 0).
		AddColorStopPosition(0.0, factoryColor.NewMagenta()).
		AddColorStopPosition(0.5, factoryColor.NewBlue()).
		AddColorStopPosition(1.0, factoryColor.NewRed()).
		FillStyleGradient().
		Font(fontB).
		FillText("Big smile!", 10, 90, 300)

	var stage = factoryBrowser.NewStage()
	stage.Append(canvas)

	done := make(chan struct{}, 0)
	<-done
}
```
