# canvas: GlobalCompositeOperation

### English:

This example shows how to use globalCompositeOperation on cavas element.

### Português:

Este exemplo mostra como usar globalCompositeOperation no elemnto canvas.

### Makefile

```shell
make help         ## This help command
make buildandrun  ## build this example and run local server
make build        ## build main.wasm file to run this example
make server       ## run local server
```

### Local server

[https://localhost/examples/canvas/globalCompositeOperation/](https://localhost/examples/canvas/globalCompositeOperation/)

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
		FillStyle(factoryColor.NewRed()).
		FillRect(20, 20, 75, 50).
		GlobalCompositeOperation(html.KCompositeOperationsRuleSourceOver).
		FillStyle(factoryColor.NewBlue()).
		FillRect(50, 50, 75, 50).
		FillStyle(factoryColor.NewRed()).
		FillRect(150, 20, 75, 50).
		GlobalCompositeOperation(html.KCompositeOperationsRuleDestinationOver).
		FillStyle(factoryColor.NewBlue()).
		FillRect(180, 50, 75, 50)

	var stage = factoryBrowser.NewStage()
	stage.Append(canvas)

	done := make(chan struct{}, 0)
	<-done
}
```
