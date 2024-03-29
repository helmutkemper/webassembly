# iotmaker webassembly

### In development phase. 

> Use at your own risk.

Why do?

> WebAssembly is a new type of code that can be run in modern web browsers — it is a low-level assembly-like language with a compact binary format that runs with near-native performance and provides languages such as C/C++, C# and Rust with a compilation target so that they can run on the web. It is also designed to run alongside JavaScript, allowing both to work together.

#### English:

This project is porting documentation from [https://developer.mozilla.org](https://developer.mozilla.org) into Golang and many of the examples are taken from it.

### Instructions/Instruções:

How to turn on the local server:
```shell
  cd examples/server
  make build
```

To see examples:
```
http://localhost:3000/examples/svg/
```

### Example of use

**Reference/Referência:**

[https://developer.mozilla.org/en-US/docs/Web/SVG/Element/animateMotion](https://developer.mozilla.org/en-US/docs/Web/SVG/Element/animateMotion)

**HTML**
```html
<svg viewBox="0 0 200 100" xmlns="http://www.w3.org/2000/svg">
  <path fill="none" stroke="lightgrey"
    d="M20,50 C20,-50 180,150 180,50 C180-50 20,150 20,50 z" />

  <circle r="5" fill="red">
    <animateMotion dur="10s" repeatCount="indefinite"
      path="M20,50 C20,-50 180,150 180,50 C180-50 20,150 20,50 z" />
  </circle>
</svg>
```

**Golang**
```go
//go:build js
package main

import (
  "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
  "github.com/helmutkemper/iotmaker.webassembly/browser/html"
  "github.com/helmutkemper/iotmaker.webassembly/platform/factoryColor"
  "time"
)

func main() {
  
  s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 200, 100}).Append(
    factoryBrowser.NewTagSvgPath().Fill(nil).Stroke(factoryColor.NewLightgrey()).D(factoryBrowser.NewPath().M(20, 50).C(20, -50, 180, 150, 180, 50).C(180, -50, 20, 150, 20, 50).Z()),
    factoryBrowser.NewTagSvgCircle().R(5).Fill(factoryColor.NewRed()).Append(
      factoryBrowser.NewTagSvgAnimateMotion().Dur(10*time.Second).RepeatCount(html.KSvgDurIndefinite).Path(factoryBrowser.NewPath().M(20, 50).C(20, -50, 180, 150, 180, 50).C(180, -50, 20, 150, 20, 50).Z()),
    ),
  )

  stage := factoryBrowser.NewStage()
  stage.Append(s1)

  done := make(chan struct{}, 0)
  <-done
}
```

Browser:

![screen example](documentation/image/screen2.png)

How to generate binary file:

```shell
  cd examples/svg/tags/animateMotion
  make build
```

<!--
## Golnag JS Tips

### How to create a new `js.Value{}`:
```go
newObject := js.Global().Get("Object")
newArray  := js.Global().Get("Array")

test := js.Global().Get("Object")
test.Set("test", "I'm alive!")
log.Printf("test: %v", test.Get("test"))
```

### How to create a callback (of hell) function:

Javascript Example:
```javascript
const options = {
  enableHighAccuracy: true,
  timeout: 5000,
  maximumAge: 0
};

function success(pos) {
  const crd = pos.coords;

  console.log('Your current position is:');
  console.log(`Latitude : ${crd.latitude}`);
  console.log(`Longitude: ${crd.longitude}`);
  console.log(`More or less ${crd.accuracy} meters.`);
}

function error(err) {
  console.warn(`ERROR(${err.code}): ${err.message}`);
}

navigator.geolocation.getCurrentPosition(success, error, options);
```

Function success, javascript:
```javascript
function success(pos) {
  const crd = pos.coords;

  console.log('Your current position is:');
  console.log(`Latitude : ${crd.latitude}`);
  console.log(`Longitude: ${crd.longitude}`);
  console.log(`More or less ${crd.accuracy} meters.`);
}
```

Function success, golang:
```go
var success = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
  // javascript `pos` is golang `args[0]`
  var crd = args[0].Get("coords")
  
  log.Printf("Your current position is:")
  log.Printf("Latitude $v:", crd.Get("latitude"))
  log.Printf("Longitude $v:", crd.Get("longitude"))
  log.Printf("More or less $v meters", crd.Get("accuracy"))
  return nil
})
```

Function error, javascript:
```javascript
function error(err) {
  console.warn(`ERROR(${err.code}): ${err.message}`);
}
```

Function error, golang:
```go
var err = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
  // javascript `err` is golang `args[0]`
  log.Printf("ERROR(%v): %v", args[0].Get("code"), args[0].Get("message")) 
})
```

Javascript options:
```javascript
const options = {
  enableHighAccuracy: true,
  timeout: 5000,
  maximumAge: 0
};
```

Javascript options, golang:
```go
var options = js.Global().Get("Object")
options.Set("enableHighAccuracy", true)
options.Set("timeout", 5000)
options.Set("maximumAge", 0)
```

Complete golang function:
```go
var options = js.Global().Get("Object")
options.Set("enableHighAccuracy", true)
options.Set("timeout", 5000)
options.Set("maximumAge", 0)

var success = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
  // javascript `pos` is golang `args[0]`
  var crd = args[0].Get("coords")

  log.Printf("Your current position is:")
  log.Printf("Latitude $v:", crd.Get("latitude"))
  log.Printf("Longitude $v:", crd.Get("longitude"))
  log.Printf("More or less $v meters", crd.Get("accuracy"))
  return nil
})

var err = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
  // javascript `err` is golang `args[0]`
  log.Printf("ERROR(%v): %v", args[0].Get("code"), args[0].Get("message"))
})

js.Global().Get("navigator").Get("geolocation").Call("getCurrentPosition", success, err, options)
```

### Hoow to make a promise

```go

```

### How to get a promise, real example:
```go
type Data struct {
  DeviceId string
  GroupId  string
  Kind     string
  Label    string
}

list := make([]Data, 0)
end := make(chan struct{})

forEach := js.FuncOf(func(_ js.Value, args []js.Value) any {
  data := Data{
    DeviceId: args[0].Get("deviceId").String(),
    GroupId:  args[0].Get("groupId").String(),
    Kind:     args[0].Get("kind").String(),
    Label:    args[0].Get("label").String(),
  }
  list = append(list, data)
  return nil
})

var success = js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
  // enumerateDevices() returns an array, but, go returns an object (bug)
  // call a forEach() for correct this problem.
  args[0].Call("forEach", forEach)
  end <- struct{}{}
  return nil
})

var failure = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
  log.Printf("message: %v", args[0].Get("message"))
  return nil
})

// enumerateDevices() returns a promise
js.Global().Get("navigator").Get("mediaDevices").Call("enumerateDevices").Call("then", success, failure)
<-end
log.Printf("list: %+v", list)
```

### How to call javascript function:

Javascript:
```html
<script>
  window.test = function (a){
    console.log(a)
  }
</script>
```

Golang:
```go
js.Global().Call("test", value)
```

### How to use constructor

Native golang to javascript types:

| Go                     | JavaScript  |
|------------------------|-------------|
| js.Value               | [its value] |
| js.Func                | function    |
| nil                    | null        |
| bool                   | boolean     |
| integers and floats    | number      |
| string                 | string      |
| []interface{}          | new array   |
| map[string]interface{} | new object  |

Javascript code:
```javascript
  var aFileParts = ['<a id="a"><b id="b">hey!</b></a>'];
  var oMyBlob = new Blob(aFileParts, {type : 'text/html'}); // o blob
```

Golang code:
```go
  done := make(chan struct{}, 0)
  
  // use native golang to work!
  aFileParts := []interface{}{"<a id=\"a\"><b id=\"b\">hey!</b></a>"}
  fType := map[string]interface{}{"type": "text/html"}
  oMyBlob := js.Global().Get("Blob").New(aFileParts, fType)
  
  log.Printf("%v", oMyBlob.Get("size"))
  log.Printf("%v", oMyBlob.Get("type"))
  
  <-done
```
-->

### List of examples:

| Example                                                                                                                             | Date         |
|-------------------------------------------------------------------------------------------------------------------------------------|--------------|
| [battery](https://github.com/helmutkemper/webassembly/tree/master/examples/battery)                                                 | 07/august/22 |
| [event/animateMotion](https://github.com/helmutkemper/webassembly/tree/master/examples/event/animateMotion)                         | 07/august/22 |
| [event/feDisplacementMap](https://github.com/helmutkemper/webassembly/tree/master/examples/event/feDisplacementMap)                 | 07/august/22 |
| [event/use](https://github.com/helmutkemper/webassembly/tree/master/examples/event/use)                                             | 07/august/22 |
| [geolocation/getPosition](https://github.com/helmutkemper/webassembly/tree/master/examples/geolocation/getPosition)                 | 07/august/22 |
| [geolocation/watchPosition](https://github.com/helmutkemper/webassembly/tree/master/examples/geolocation/watchPosition)             | 07/august/22 |
| [BezierCurve/1_addCurve](https://github.com/helmutkemper/webassembly/tree/master/examples/BezierCurve/1_addCurve)                   | 08/august/22 |
| [BezierCurve/2_addEasingTween](https://github.com/helmutkemper/webassembly/tree/master/examples/BezierCurve/2_addEasingTween)       | 08/august/22 |
| [BezierCurve/3_addRipple](https://github.com/helmutkemper/webassembly/tree/master/examples/BezierCurve/3_addRipple)                 | 08/august/22 |
| [canvas/addColorStopPosition](https://github.com/helmutkemper/webassembly/tree/master/examples/canvas/addColorStopPosition)         | 16/august/22 |
| [canvas/arc](https://github.com/helmutkemper/webassembly/tree/master/examples/canvas/arc)                                           | 16/august/22 |
| [canvas/arcTo](https://github.com/helmutkemper/webassembly/tree/master/examples/canvas/arcTo)                                       | 16/august/22 |
| [canvas/beginPath](https://github.com/helmutkemper/webassembly/tree/master/examples/canvas/beginPath)                               | 16/august/22 |
| [canvas/bezierCurveTo](https://github.com/helmutkemper/webassembly/tree/master/examples/canvas/bezierCurveTo)                       | 16/august/22 |
| [canvas/clearRect](https://github.com/helmutkemper/webassembly/tree/master/examples/canvas/clearRect)                               | 16/august/22 |
| [canvas/createLinearGradient](https://github.com/helmutkemper/webassembly/tree/master/examples/canvas/createLinearGradient)         | 16/august/22 |
| [canvas/createPattern](https://github.com/helmutkemper/webassembly/tree/master/examples/canvas/createPattern)                       | 16/august/22 |
| [canvas/createRadialGradient](https://github.com/helmutkemper/webassembly/tree/master/examples/canvas/createRadialGradient)         | 16/august/22 |
| [canvas/fillRect](https://github.com/helmutkemper/webassembly/tree/master/examples/canvas/fillRect)                                 | 16/august/22 |
| [canvas/fillText](https://github.com/helmutkemper/webassembly/tree/master/examples/canvas/fillText)                                 | 16/august/22 |
| [canvas/font](https://github.com/helmutkemper/webassembly/tree/master/examples/canvas/font)                                         | 16/august/22 |
| [canvas/globalAlpha](https://github.com/helmutkemper/webassembly/tree/master/examples/canvas/globalAlpha)                           | 16/august/22 |
| [canvas/globalCompositeOperation](https://github.com/helmutkemper/webassembly/tree/master/examples/canvas/globalCompositeOperation) | 16/august/22 |
