//go:build js
// +build js

//
package main

import (
	global "github.com/helmutkemper/iotmaker.santa_isabel_theater.globalConfig"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/browserMouse"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/css"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/factoryBrowserImage"
	doc "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/globalDocument"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/html"
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

	var class = new(css.Class)
	// Create a css list named "red" with value "user red"
	// Crie uma lista css de nome "red" com o valor "user red"
	class.SetList("red", "user", "red").
		// Create a css list named "yellow" with value "user yellow"
		// Crie uma lista css de nome "yellow" com o valor "user yellow"
		SetList("yellow", "user", "yellow").
		// Create a css list named "user" with value "user"
		// Crie uma lista css de nome "user" com o valor "user"
		SetList("user", "user").
		// Defines that the "red" and "yellow" lists will change every second
		// Define que as listas "red" e "yellow" vão trocar a cada segundo
		ToggleTime(time.Second, "red", "yellow").
		// Limit trades to 10 interactions
		// Limita as trocas em 10 interações
		ToggleLoop(10).
		// Defines the list named "norm" as the active list at the end of interactions
		// Define  alista de nome "normal" como sendo a lista ativa ao final das interações
		OnLoopEnd("user").
		// Start interactions. Caution: they only work after being added to the element
		// Inicia as interações. Cuidado: elas só funcionam após serem adicionadas ao elemento
		ToggleStart()

	var a html.Div

	// Create a div with id "example";
	// Cria uma div de id "example_A";
	a.NewDiv("example_A").
		// Sets css to be "name_a name_b name_N";
		// Define css como sendo "name_a name_b name_N";
		Css("animate").
		SetMousePointer(browserMouse.KCursorMove).
		// Adds the div to the element id "stage".
		// Adiciona a div ao elemento de id "stage".
		AppendById("stage")

	var b html.Div
	// Create a div with id "example";
	// Cria uma div de id "example_A";
	b.NewDiv("example_B").
		// Sets css to be "name_a name_b name_N";
		// Define css como sendo "name_a name_b name_N";
		Css("name_a", "name_b", "name_N").
		// css.Class cannot work properly before being added, due to lack of reference to the parent
		// object.
		// css.Class não consegue funcionar corretamente antes de ser adicionada, por falta de referência
		// do objeto pai.
		SetCssController(class).
		// Adds the div to the element id "stage".
		// Adiciona a div ao elemento de id "stage".
		AppendById("stage")

	var err error
	//document.GetElementById(document, "stage")
	for a := 0; a != 300; a += 1 {

		id := "div_" + strconv.FormatInt(int64(a), 10)
		_, err = document.CreateElementAndAppend(
			"stage",
			"div",
			[]string{"animate"},
			doc.P{P: "id", V: id},
		)
		if err != nil {
			log.Printf("document.CreateElement().error: %v", err.Error())
		}
		var e = document.GetElementById(document, id)
		var border = 0
		factoryTween.NewLinear(
			time.Duration(mathUtil.Int(1000, 3000))*time.Millisecond,
			mathUtil.Float64FomInt(border, global.Global.Document.GetDocumentWidth()-29-border),
			mathUtil.Float64FomInt(border, global.Global.Document.GetDocumentWidth()-29-border),
			func(x, p float64, ars ...interface{}) {
				px := strconv.FormatFloat(x, 'E', 10, 32) + "px"
				document.SetElementStyle(e, "left", px)
			},
			-1,
		)

		factoryTween.NewLinear(
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