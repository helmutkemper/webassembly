# battery event

### English:

This example shows how to monitor the device battery in real time.

### Português:

Este exemplo mostra como monitorar a bateria do dispositivo em tempo real.

### Makefile

```shell
make help         ## This help command
make buildandrun  ## build this example and run local server
make build        ## build main.wasm file to run this example
make server       ## run local server
```

### Local server

[https://localhost/examples/battery/](https://localhost/examples/battery/)

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
  "github.com/helmutkemper/iotmaker.webassembly/browser/event/battery"
  "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
)

func main() {

  div := factoryBrowser.NewTagDiv()

  stage := factoryBrowser.NewStage()
  stage.Append(div)

  var batteryEvent = make(chan battery.Data)

  bat := factoryBrowser.NewBattery()
  bat.AddListenerChargingChange(&batteryEvent)
  bat.AddListenerDischargingTimeChange(&batteryEvent)
  bat.AddListenerChargingTimeChange(&batteryEvent)
  bat.AddListenerLevelChange(&batteryEvent)

  go func() {
    for {
      select {
      case data := <-batteryEvent:
        text := ""
        text += fmt.Sprintf("event name: %v", data.EventName)
        text += "<br>"
        text += fmt.Sprintf("level: %v", data.Level)
        text += "<br>"
        text += fmt.Sprintf("charging: %v", data.Charging)
        text += "<br>"
        text += fmt.Sprintf("charging time: %v", data.ChargingTime)
        text += "<br>"
        text += fmt.Sprintf("discharging time: %v", data.DischargingTime)
        div.Html(text)
      }
    }
  }()

  batteryEvent <- bat.Now()

  done := make(chan struct{})
  <-done
}
```