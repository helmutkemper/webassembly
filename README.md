# iotmaker webassembly

### In development phase. 

> Use at your own risk.

Why do? / Por quê fazer?

> WebAssembly is a new type of code that can be run in modern web browsers — it is a low-level assembly-like language with a compact binary format that runs with near-native performance and provides languages such as C/C++, C# and Rust with a compilation target so that they can run on the web. It is also designed to run alongside JavaScript, allowing both to work together.

#### English:

This project is porting documentation from developer.mozilla.org into Golang and many of the examples are taken from it.

Now, there are more than [140 examples](examples/svg/attributes) of how to use svg tags and more than [50 examples](examples/svg/tags) of how to use their properties.

#### Português:

Este projeto está portando a documentação do site developer.mozilla.org para dentro do Golang e muitos dos exemplos foram tirados dele.

No momento, são mais de [140 exemplos](examples/svg/attributes) de como usar as tags svg e mais de [50 exemplos](examples/svg/tags) de como usar as suas propriedades.

### Instructions/Instruções:

How to turn on the local server: / Como ligar o servidor local:
```shell
  cd examples/server
  make build
```

To see what the documentation will be: / Para ver o que será a documentação:
```
http://localhost:3000
```

To see the stable examples for use: / Para ver os exemplos estáveis para uso:
```
http://localhost:3000/examples/svg/
```

### Example of use / Exemplo de uso

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
// +build js

package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/browser/html"
	"github.com/helmutkemper/iotmaker.webassembly/platform/factoryColor"
	"time"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 200, 100}).Append(
		factoryBrowser.NewTagSvgPath().Fill("none").Stroke(factoryColor.NewLightgrey()).D(factoryBrowser.NewPath().M(20, 50).C(20, -50, 180, 150, 180, 50).C(180, -50, 20, 150, 20, 50).Z()),
		factoryBrowser.NewTagSvgCircle().R(5).Fill(factoryColor.NewRed()).Append(
			factoryBrowser.NewTagSvgAnimateMotion().Dur(10*time.Second).RepeatCount(html.KSvgDurIndefinite).Path(factoryBrowser.NewPath().M(20, 50).C(20, -50, 180, 150, 180, 50).C(180, -50, 20, 150, 20, 50).Z()),
		),
	)

	stage.Append(s1)

	<-done
}
```

Browser: / Navegador:

![screen example](documentation/image/screen2.png)

How to generate binary file: / Como gerar o arquivo binário

```shell
  cd examples/svg/tags/animateMotion
  make build
```

### Documentation/Documentação


#### English:
While the current documentation site is not ready, use the in-code documentation. It was written in English and Portuguese.

#### Português:
Enquanto o site de documentação atual não fica pronta, use a documentação embutida no código. Ela foi escrita em inglês e em português.

![documentation](documentation/image/screen.png)

