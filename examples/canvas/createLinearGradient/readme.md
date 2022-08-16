# canvas: ArcTo

### English:

This example shows how to use createLinearGradient on cavas element.

### Português:

Este exemplo mostra como usar createLinearGradient no elemnto canvas.

### Makefile

```shell
make help         ## This help command
make buildandrun  ## build this example and run local server
make build        ## build main.wasm file to run this example
make server       ## run local server
```

### Local server

[https://localhost/examples/canvas/createLinearGradient/](https://localhost/examples/canvas/createLinearGradient/)

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
	"github.com/helmutkemper/iotmaker.webassembly/browser/html"
	"github.com/helmutkemper/iotmaker.webassembly/platform/factoryColor"
)

var canvas *html.TagCanvas

func main() {

	canvas = factoryBrowser.NewTagCanvas(800, 600).
		CreateLinearGradient(0, 0, 170, 0).
		AddColorStopPosition(0.0, factoryColor.NewBlack()).
		AddColorStopPosition(0.5, factoryColor.NewOrangered()).
		AddColorStopPosition(1.0, factoryColor.NewWhite()).
		FillStyleGradient().
		FillRect(20, 20, 150, 100)

	var stage = factoryBrowser.NewStage()
	stage.Append(canvas)

	done := make(chan struct{}, 0)
	<-done
}
```