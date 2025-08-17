package main

import (
	"math"
	"syscall/js"
)

// Engine representa o estado da simulação/desenho no Worker.
type Engine struct {
	ctx    js.Value
	canvas js.Value
	perf   js.Value
	// Dimensões em CSS pixels (lógicas)
	cssW, cssH float64
	// DPR atual
	dpr float64

	// Timestep fixo (lógica a 240 Hz)
	lastMS float64
	accMS  float64
	stepMS float64

	// Estado do demo (quadrado quicando)
	x, y   float64
	vx, vy float64
	size   float64

	// Mantém referências para evitar GC
	tickFunc   js.Func
	resizeFunc js.Func
}

// initEngine
//
// English:
//
//	Initializes the engine, wiring JS bridges, reading OffscreenCanvas from the
//	worker scope, and preparing the 2D context scaled by DPR.
//
// Português:
//
//	Inicializa a engine, conectando as pontes com o JS, lendo o OffscreenCanvas do
//	escopo do worker e preparando o contexto 2D escalado pelo DPR.
//
//	Todas as funções devem ser simples e de responsabilidade única.
func initEngine() *Engine {
	global := js.Global()
	wasmEnv := global.Get("wasmEnv") // definido no worker.js
	if wasmEnv.IsUndefined() || wasmEnv.IsNull() {
		panic("wasmEnv not found — worker.js must set globalThis.wasmEnv")
	}

	canvas := wasmEnv.Get("canvas")
	if canvas.IsUndefined() || canvas.IsNull() {
		panic("OffscreenCanvas not provided — main.js must postMessage({canvas})")
	}
	ctx := canvas.Call("getContext", "2d")
	perf := global.Get("performance")

	e := &Engine{
		ctx:    ctx,
		canvas: canvas,
		perf:   perf,
		dpr:    wasmEnv.Get("dpr").Float(),
		cssW:   wasmEnv.Get("cssW").Float(),
		cssH:   wasmEnv.Get("cssH").Float(),
		stepMS: 1000.0 / 240.0, // 240 Hz
		size:   40.0,           // em CSS px
		vx:     220.0,          // px/s
		vy:     160.0,          // px/s
	}

	e.applyTransform()
	e.clear()

	// Exporta callbacks para o worker.js acionar
	e.tickFunc = js.FuncOf(func(this js.Value, args []js.Value) any {
		var now float64
		if len(args) > 0 {
			now = args[0].Float()
		} else {
			now = e.perf.Call("now").Float()
		}
		e.onTick(now)
		return nil
	})
	js.Global().Set("onWasmTick", e.tickFunc)

	e.resizeFunc = js.FuncOf(func(this js.Value, args []js.Value) any {
		e.onResize()
		return nil
	})
	js.Global().Set("onWasmResize", e.resizeFunc)

	// Timestamp inicial
	e.lastMS = e.perf.Call("now").Float()
	return e
}

// onResize
//
// English:
//
//	Reads the latest size/DPR from wasmEnv and reapplies the 2D transform so
//	that we draw in CSS pixels while the backbuffer stays in device pixels.
//
// Português:
//
//	Lê o tamanho/DPR recente de wasmEnv e reaplica o transform 2D para desenhar
//	em pixels CSS enquanto o backbuffer permanece em pixels do dispositivo.
//
//	Todas as funções devem ser simples e de responsabilidade única.
func (e *Engine) onResize() {
	wasmEnv := js.Global().Get("wasmEnv")
	e.dpr = wasmEnv.Get("dpr").Float()
	e.cssW = wasmEnv.Get("cssW").Float()
	e.cssH = wasmEnv.Get("cssH").Float()
	e.applyTransform()
	e.clear()
}

// applyTransform
//
// English:
//
//	Sets the transform to map CSS pixels to device pixels using DPR.
//
// Português:
//
//	Define o transform para mapear pixels CSS para pixels do dispositivo usando DPR.
//
//	Todas as funções devem ser simples e de responsabilidade única.
func (e *Engine) applyTransform() {
	// Reseta a matriz e aplica escala DPR (desenharemos em coordenadas CSS)
	e.ctx.Call("setTransform", e.dpr, 0, 0, e.dpr, 0, 0)
}

// clear
//
// English:
//
//	Clears the canvas in CSS coordinates.
//
// Português:
//
//	Limpa o canvas em coordenadas CSS.
//
//	Todas as funções devem ser simples e de responsabilidade única.
func (e *Engine) clear() {
	e.ctx.Call("clearRect", 0, 0, e.cssW, e.cssH)
}

// onTick
//
// English:
//
//	Receives a vsync "tick" from main thread (requestAnimationFrame). Advances
//	the simulation with a fixed timestep (240 Hz) and renders once.
//
// Português:
//
//	Recebe um “tick” de vsync do thread principal (requestAnimationFrame).
//	Avança a simulação com passo fixo (240 Hz) e renderiza uma vez.
//
//	Todas as funções devem ser simples e de responsabilidade única.
func (e *Engine) onTick(nowMS float64) {
	frameMS := math.Min(nowMS-e.lastMS, 250.0) // trava picos longos
	e.lastMS = nowMS
	e.accMS += frameMS

	for e.accMS >= e.stepMS {
		e.update(e.stepMS / 1000.0) // dt em segundos
		e.accMS -= e.stepMS
	}
	alpha := e.accMS / e.stepMS
	e.render(alpha)
}

// update
//
// English:
//
//	Updates the demo physics using a fixed dt (seconds).
//
// Português:
//
//	Atualiza a “física” do demo usando dt fixo (segundos).
//
//	Todas as funções devem ser simples e de responsabilidade única.
func (e *Engine) update(dt float64) {
	e.x += e.vx * dt
	e.y += e.vy * dt

	half := e.size / 2.0
	if e.x-half < 0 {
		e.x = half
		e.vx = math.Abs(e.vx)
	} else if e.x+half > e.cssW {
		e.x = e.cssW - half
		e.vx = -math.Abs(e.vx)
	}
	if e.y-half < 0 {
		e.y = half
		e.vy = math.Abs(e.vy)
	} else if e.y+half > e.cssH {
		e.y = e.cssH - half
		e.vy = -math.Abs(e.vy)
	}
}

// render
//
// English:
//
//	Renders a simple bouncing square. Uses CSS pixel coordinates thanks to the
//	DPR transform set earlier.
//
// Português:
//
//	Desenha um quadrado quicando. Usa coordenadas em pixels CSS graças ao
//	transform de DPR aplicado anteriormente.
//
//	Todas as funções devem ser simples e de responsabilidade única.
func (e *Engine) render(alpha float64) {
	_ = alpha // mantido para técnicas de interpolação se desejar

	e.clear()
	e.ctx.Set("fillStyle", "#09f")
	e.ctx.Call("fillRect", e.x-e.size/2.0, e.y-e.size/2.0, e.size, e.size)
}

// main
//
// English:
//
//	Entry point: initializes the engine and parks the goroutine.
//
// Português:
//
//	Ponto de entrada: inicializa a engine e estaciona a goroutine.
//
//	Todas as funções devem ser simples e de responsabilidade única.
func main() {
	e := initEngine()

	// Posição inicial no centro
	e.x = e.cssW * 0.5
	e.y = e.cssH * 0.5

	// Impede o programa de sair
	select {}
}
