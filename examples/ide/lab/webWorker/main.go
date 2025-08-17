//go:build js && wasm

package main

import (
	"math"
	"syscall/js"
	"time"
)

// initWorkerBindings
//
// English:
//
//	Initializes bindings used inside the Web Worker (OffscreenCanvas, timer loop, and simple 2D draw).
//
//	All functions are straightforward and single-responsibility. Enabling already-enabled features has no adverse effect.
//
// Português:
//
//	Inicializa os bindings usados dentro do Web Worker (OffscreenCanvas, laço por timer e desenho 2D simples).
//
//	Todas as funções são simples e de responsabilidade única. Habilitar algo já habilitado não causa efeito adverso.
func initWorkerBindings() {
	global := js.Global()

	// Recupera o OffscreenCanvas exposto pelo worker (self.canvas)
	canvas := global.Get("canvas")
	if canvas.IsUndefined() || canvas.IsNull() {
		// Worker ainda não enviou o canvas
		return
	}

	ctx := canvas.Call("getContext", "2d")
	start := time.Now()

	// Função de desenho: limpa e desenha um quadrado animado
	draw := js.FuncOf(func(this js.Value, args []js.Value) any {
		w := canvas.Get("width").Float()
		h := canvas.Get("height").Float()

		// clear
		ctx.Call("clearRect", 0, 0, w, h)

		// animação simples: movimento senoidal no eixo X
		t := time.Since(start).Seconds()
		size := 60.0
		x := (w-size)/2 + math.Sin(t*2.0)*((w-size)/3)
		y := (h - size) / 2

		ctx.Set("fillStyle", "rgba(255,40,255,0.9)")
		ctx.Call("fillRect", x, y, size, size)
		return nil
	})

	// Timer no worker (workers não possuem requestAnimationFrame clássico)
	// 120 FPS alvo (ajuste conforme necessário)
	ticker := time.NewTicker(time.Second / 120)
	go func() {
		for range ticker.C {
			draw.Invoke()
		}
	}()
}

// main
//
// English:
//
//	Entry point for the Web Worker runtime. It only initializes worker bindings and blocks forever.
//
// Português:
//
//	Ponto de entrada no runtime do Web Worker. Apenas inicializa os bindings do worker e bloqueia indefinidamente.
func main() {
	initWorkerBindings()
	select {}
}
