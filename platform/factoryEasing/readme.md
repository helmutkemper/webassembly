# Example

Webassembly and Golang factoryTween functions example

This example can move 200 divs simultaneously on my Sansung Galaxy A51 phone (a basic cell phone model, with android) and 700 divs on my Mac Book.

Environment variables:
```shell
GOARCH=wasm
GOOS=js
```

Go tool arguments:
```shell
-o main.wasm
```

Code Golang:
```go
//go:build js
// +build js

//
package main

import (
	global "github.com/helmutkemper/iotmaker.santa_isabel_theater.globalConfig"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/css"
	document2 "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/document"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/factoryBrowserImage"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/factoryTween"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/mathUtil"
	"log"
	"strconv"
	"time"
)

func main() {
	
	done := make(chan struct{}, 0)
	document := global.Global.Document
	
	// Carrega a imagem
	factoryBrowserImage.NewImage(
		29,
		50,
		map[string]interface{}{
			"id":  "spacecraft",
			"src": "./small.png",
		},
		true,
		false,
	)
	
	var err error
	document.GetElementById(document, "palco")
	for a := 0; a != 10; a += 1 {
		
		id := "div_" + strconv.FormatInt(int64(a), 10)
		var cssClass = css.Class{}
		cssClass.SetList("current", "animate")
		err = document.CreateElement(document, "palco", "div", document2.Property{Property: "id", Value: id}, cssClass)
		if err != nil {
			log.Printf("document.CreateElement().error: %v", err.Error())
		}
		var e = document.GetElementById(document, id)
		var border = 200
		factoryTween.NewSelectRandom(
			time.Duration(mathUtil.Int(1000, 3000))*time.Millisecond,
			mathUtil.Float64FomInt(border, global.Global.Document.GetDocumentWidth()-29-border),
			mathUtil.Float64FomInt(border, global.Global.Document.GetDocumentWidth()-29-border),
			func(x, p float64, ars ...interface{}) {
				px := strconv.FormatFloat(x, 'E', 10, 32) + "px"
				document.SetElementStyle(e, "left", px)
			},
			-1,
		)
		
		factoryTween.NewSelectRandom(
			time.Duration(mathUtil.Int(1000, 3000))*time.Millisecond,
			mathUtil.Float64FomInt(border, global.Global.Document.GetDocumentHeight()-50-border),
			mathUtil.Float64FomInt(border, global.Global.Document.GetDocumentHeight()-50-border),
			func(y, p float64, ars ...interface{}) {
				py := strconv.FormatFloat(y, 'E', 10, 32) + "px"
				document.SetElementStyle(e, "top", py)
			},
			-1,
		)
	
	}
	
	<-done
}
```

Html code:
```html
<html>
<head>
  <meta charset="utf-8"/>
  <style>
      body {
          margin: 0 !important;
          padding: 0 !important;
      }

      .animate {
        width: 29px;
        height: 50px;
        position: absolute;
        background-image: url("./small.png");
      }
  </style>
  <script src="wasm_exec.js"></script>
  <script>
    const go = new Go();
    WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then((result) => {
      go.run(result.instance);
    });
  </script>
</head>
<body>
  <div id="palco"></div>
</body>
</html>
```

Webassembly needs a server to run. The example below is a simple static server that prints the IP address to standard output.
```go
package main

import (
  "log"
  "net"
  "net/http"
)

func main() {
  var err error
  var addrs []net.Addr
  
  var ifaces []net.Interface
  
  ifaces, err = net.Interfaces()
  // handle err
  for _, i := range ifaces {
    addrs, err = i.Addrs()
    // handle err
    for _, addr := range addrs {
      var ip net.IP
      switch v := addr.(type) {
      case *net.IPNet:
        ip = v.IP
      case *net.IPAddr:
        ip = v.IP
      }
      log.Printf("addr: %v", ip)
    }
  }
  
  fs := http.FileServer(http.Dir("./"))
  http.Handle("/", fs)
  
  log.Println("Listening on :3000..")
  err = http.ListenAndServe(":3000", nil)
  if err != nil {
    log.Fatal(err)
  }
}
```

<!-- https://github.com/ai/easings.net/blob/master/src/math/math.pug -->
<!-- https://easings.net/pt-br -->
<!-- https://gist.github.com/cjddmut/d789b9eb78216998e95c -->
<!-- https://gist.github.com/cjddmut -->