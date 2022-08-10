# battery event

### English:

This example shows how to get data about the browser and operating system.

### Português:

Este exemplo mostra como pegar dados sobre o navegador e o sistema operacional.

### Makefile

```shell
make help         ## This help command
make buildandrun  ## build this example and run local server
make build        ## build main.wasm file to run this example
make server       ## run local server
```

### Local server

[https://localhost/examples/userAgent/](https://localhost/examples/userAgent/)

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
  "fmt"
  "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
  "github.com/helmutkemper/iotmaker.webassembly/browser/userAgent"
)

func main() {

  div := factoryBrowser.NewTagDiv()
  stage := factoryBrowser.NewStage()
  stage.Append(div)

  data := userAgent.GetHighEntropyValues(
    userAgent.KHintsArchitecture,
    userAgent.KHintsBitness,
    userAgent.KHintsModel,
    userAgent.KHintsPlatformVersion,
    userAgent.KHintsFullVersionList,
  )

  text := fmt.Sprintf("Brands: %v<br>", data.Brands)
  text += fmt.Sprintf("Mobile: %v<br>", data.Mobile)
  text += fmt.Sprintf("Platform: %v<br>", data.Platform)
  text += fmt.Sprintf("Architecture: %v<br>", data.Architecture)
  text += fmt.Sprintf("Bitness: %v<br>", data.Bitness)
  text += fmt.Sprintf("Model: %v<br>", data.Model)
  text += fmt.Sprintf("PlatformVersion: %v<br>", data.PlatformVersion)
  text += fmt.Sprintf("FullVersionList: %v<br>", data.FullVersionList)
  div.Html(text)

  done := make(chan struct{})
  <-done
}
```
