package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/png"
	"log"
	"net/http"
	"sync"
	"syscall/js"
	"time"
)

func GetOsmTile(x, y, z int) (jsImg js.Value) {
	var err error
	var req *http.Request
	var resp *http.Response
	var pngImg image.Image
	var url = "https://tile.openstreetmap.org/%v/%v/%v.png"

	url = fmt.Sprintf(url, z, x, y)
	req, err = http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Printf("client: could not create request: %s\n", err)
		return
	}
	//req.Header.Set("Content-Type", "application/json")
	//req.Header.Set("Access-Control-Allow-Origin", "*")
	//req.Header.Set("Access-Control-Allow-Methods", "GET")
	//req.Header.Set("Access-Control-Allow-Headers", "*")

	client := http.Client{
		Timeout: 15 * time.Second, // todo: permitir setup
	}

	resp, err = client.Do(req)
	if err != nil {
		log.Printf("client: error making http request: %s\n", err)
		return
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	pngImg, err = png.Decode(resp.Body)
	if err != nil {
		log.Fatalf("Erro ao decodificar a imagem: %v", err)
	}

	jsImg = js.Global().Get("Image").New()
	buf := new(bytes.Buffer)
	err = png.Encode(buf, pngImg)
	if err != nil {
		log.Printf("Erro ao bufferizar a imagem: %v", err)
		return
	}

	wg := new(sync.WaitGroup)
	wg.Add(1)
	jsImg.Set("src", "data:image/png;base64,"+base64.StdEncoding.EncodeToString(buf.Bytes()))
	jsImg.Call("addEventListener", "load", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		wg.Done()
		return nil
	}))
	wg.Wait()

	return
}

func main() {
	self := js.Global().Get("self")
	self.Call("addEventListener", "message", js.FuncOf(func(this js.Value, args []js.Value) any {
		var x, y, z int
		x = args[0].Int()
		y = args[1].Int()
		z = args[2].Int()

		jsImg := GetOsmTile(x, y, z)

		blobOptions := js.Global().Get("Object").New()
		//blobOptions.Set("type", "text/javascript")
		blobOptions.Set("transfer", jsImg)
		blob := js.Global().Get("Blob").New(blobOptions)
		array := js.Global().Get("Array").New(blob)

		return nil
	}))
}
