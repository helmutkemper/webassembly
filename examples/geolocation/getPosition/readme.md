# get geolocation

### English:

This example shows how to get the device's position.

### Português:

Este exemplo mostra como pegar a posição do dispositivo.

### Makefile

```shell
make help         ## This help command
make buildandrun  ## build this example and run local server
make build        ## build main.wasm file to run this example
make server       ## run local server
```

### Local server

[https://localhost/examples/geolocation/getPosition/](https://localhost/examples/geolocation/getPosition/)

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

#### Golang

```go
//go:build js

package main

import (
	"fmt"
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/browser/geolocation"
)

func main() {

	div1 := factoryBrowser.NewTagDiv().Html("loading ...")

	stage := factoryBrowser.NewStage()
	stage.Append(div1)

	var coordinate = make(chan geolocation.Coordinate)

	go func() {
		select {
		case converted := <-coordinate:
			text := fmt.Sprintf("Latitude: %v<br>", converted.Latitude)
			text += fmt.Sprintf("Longitude: %v<br>", converted.Longitude)
			text += fmt.Sprintf("Accuracy: %v meters<br>", converted.Accuracy)
			text += fmt.Sprintf("Error: %v<br>", converted.ErrorMessage)
			div1.Html(text)
		}
	}()

	var geo = factoryBrowser.NewGeoLocation()
	geo.GetPosition(&coordinate)

	done := make(chan struct{}, 0)
	<-done
}
```