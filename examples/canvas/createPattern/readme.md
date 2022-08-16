# canvas: CreatePattern

### English:

This example shows how to use createPattern on cavas element.

### Português:

Este exemplo mostra como usar createPattern no elemnto canvas.

### Makefile

```shell
make help         ## This help command
make buildandrun  ## build this example and run local server
make build        ## build main.wasm file to run this example
make server       ## run local server
```

### Local server

[https://localhost/examples/canvas/createPattern/](https://localhost/examples/canvas/createPattern/)

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
)

var canvas *html.TagCanvas

func main() {

	var img = factoryBrowser.NewTagImg().
		Alt("spacecraft").
		Src("./small.png", true).
		Width(29).
		Height(50)

	canvas = factoryBrowser.NewTagCanvas(800, 600).
		CreatePattern(img, html.KRepeatRuleRepeat).
		Rect(0, 0, 300, 300).
		FillStylePattern().
		Fill()

	var stage = factoryBrowser.NewStage()
	stage.Append(canvas)

	done := make(chan struct{}, 0)
	<-done
}
```
