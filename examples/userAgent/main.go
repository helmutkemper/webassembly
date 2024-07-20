//go:build js

package main

import (
	"fmt"
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/browser/userAgent"
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
