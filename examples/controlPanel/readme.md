# SVG Event

### English:

This example shows how to use event on SVG elements.

### Português:

Este exemplo mostra como usar evento em elementos SVG.

### Makefile

```shell
make help         ## This help command
make buildandrun  ## build this example and run local server
make build        ## build main.wasm file to run this example
make server       ## run local server
```

### Local server

[https://localhost/examples/event/use/](https://localhost/examples/event/use/)

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
	"github.com/helmutkemper/iotmaker.webassembly/browser/event/mouse"
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/browser/html"
	"github.com/helmutkemper/iotmaker.webassembly/platform/factoryColor"
)

func main() {

	done := make(chan struct{}, 0)

	tagDiv := &html.TagDiv{}
	tagUse := &html.TagSvgUse{}
	mouseEvent := make(chan mouse.Data)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 10}).Append(
		factoryBrowser.NewTagSvgCircle().AddListenerClick(&mouseEvent).Id("myCircle").Cx(5).Cy(5).R(4).Stroke(factoryColor.NewBlue()),
		factoryBrowser.NewTagSvgUse().Reference(&tagUse).HRef("#myCircle").X(10).Fill(factoryColor.NewBlue()),
		factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(20).Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewRed()),
	)
	div1 := factoryBrowser.NewTagDiv().Reference(&tagDiv)

	go func() {
		text := ""
		for {
			select {
			case <-mouseEvent:
				text += "click<br>"
				tagDiv.Html(text)
				// English: addEventListener('click') was created on the <circle> element, so the reference is invalid and the command does not work.
				// Português: addEventListener('click') foi criado no elemento <circle>, por isto, a refereência é inválida e o comando não funciona.
				tagUse.RemoveListenerClick()
			}
		}
	}()

	stage.Append(s1, div1)

	<-done
}
```
