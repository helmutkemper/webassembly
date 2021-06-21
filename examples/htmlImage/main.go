// +build js

package main

import (
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/factoryBrowserDocument"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/factoryBrowserHtml"
)

func main() {

	done := make(chan struct{}, 0)

	browserDocument := factoryBrowserDocument.NewDocument()

	factoryBrowserHtml.NewImage(
		browserDocument.SelfDocument,
		map[string]interface{}{
			"id":  "player",
			"src": "./player_big.png",
		},
		true,
		true,
	)

	<-done
}
