// //go:build js && wasm

package main

import (
	"syscall/js"
)

// RAF (RequestAnimationFrame) wrapper
//
// English:
//
//	Lightweight wrapper around window.requestAnimationFrame.
//	- Start(): starts the loop
//	- Stop(): cancels the loop and releases JS resources
//	- tick(dt): user callback gets delta-time in seconds
//
// Português:
//
//	Envoltório leve sobre window.requestAnimationFrame.
//	- Start(): inicia o loop
//	- Stop(): cancela o loop e libera recursos JS
//	- tick(dt): callback do usuário recebe delta-time em segundos
type RAF struct {
	win     js.Value
	cb      js.Func
	reqID   int
	running bool
	lastTS  float64
	tick    func(dt float64)
}

// NewRAF
//
// English:
//
//	Creates a RAF instance. The tick function receives dt (seconds).
//
// Português:
//
//	Cria uma instância de RAF. A função tick recebe dt (segundos).
func NewRAF(tick func(dt float64)) *RAF {
	return &RAF{
		win:  js.Global().Get("window"),
		tick: tick,
	}
}

// Start
//
// English:
//
//	Starts the animation loop. If already running, it does nothing.
//
// Português:
//
//	Inicia o loop de animação. Se já estiver rodando, não faz nada.
func (r *RAF) Start() {
	if r.running {
		return
	}
	r.running = true
	r.lastTS = 0

	r.cb = js.FuncOf(func(this js.Value, args []js.Value) any {
		// DOMHighResTimeStamp (ms)
		ts := args[0].Float()
		var dt float64
		if r.lastTS > 0 {
			dt = (ts - r.lastTS) / 1000.0
		} else {
			dt = 0
		}
		r.lastTS = ts

		// user tick
		if r.tick != nil {
			r.tick(dt)
		}

		if r.running {
			r.reqID = r.win.Call("requestAnimationFrame", r.cb).Int()
		}
		return nil
	})

	r.reqID = r.win.Call("requestAnimationFrame", r.cb).Int()
}

// Stop
//
// English:
//
//	Stops the loop, cancels the pending frame and releases the JS callback.
//
// Português:
//
//	Para o loop, cancela o frame pendente e libera o callback JS.
func (r *RAF) Stop() {
	if !r.running {
		return
	}
	r.running = false
	r.win.Call("cancelAnimationFrame", r.reqID)
	if r.cb.Truthy() {
		r.cb.Release()
	}
}

// SetupCanvas
//
// English:
//
//	Creates a canvas, sets a CSS size (logical) and scales for devicePixelRatio.
//	Returns canvas and 2D context.
//
// Português:
//
//	Cria um canvas, define tamanho CSS (lógico) e escala para devicePixelRatio.
//	Retorna o canvas e o contexto 2D.
func SetupCanvas(cssW, cssH int) (js.Value, js.Value) {
	doc := js.Global().Get("document")
	canvas := doc.Call("createElement", "canvas")
	body := doc.Get("body")
	body.Call("appendChild", canvas)

	// CSS size (logical px)
	canvas.Get("style").Set("width", cssW)
	canvas.Get("style").Set("height", cssH)
	canvas.Get("style").Set("border", "1px solid #ccc")

	// DPR scale for crisp rendering
	dpr := js.Global().Get("devicePixelRatio")
	if !dpr.Truthy() {
		dpr = js.ValueOf(1)
	}
	scale := dpr.Float()
	canvas.Set("width", int(float64(cssW)*scale))
	canvas.Set("height", int(float64(cssH)*scale))

	ctx := canvas.Call("getContext", "2d")
	ctx.Call("scale", scale, scale)

	return canvas, ctx
}

// Demo: moving box
//
// English:
//
//	Simple demo: a box moves horizontally. Space toggles start/stop.
//
// Português:
//
//	Demo simples: um quadrado se move na horizontal. Espaço alterna iniciar/parar.
func main() {
	_, ctx := SetupCanvas(480, 200)

	// state
	var (
		x     float64 = 10
		y     float64 = 80
		speed float64 = 100 // px/s
		w, h  float64 = 50, 50
	)

	// draw function
	draw := func() {
		ctx.Call("clearRect", 0, 0, 480, 200)
		ctx.Set("fillStyle", "#2b6cb0")
		ctx.Call("fillRect", x, y, w, h)
	}

	// tick updates position with dt
	raf := NewRAF(func(dt float64) {
		x += speed * dt
		if x > 480 {
			x = -w
		}
		draw()
	})

	// Start immediately
	raf.Start()

	// Toggle with Space
	doc := js.Global().Get("document")
	keyHandler := js.FuncOf(func(this js.Value, args []js.Value) any {
		ev := args[0]
		if ev.Get("code").String() == "Space" {
			if raf.running {
				raf.Stop()
			} else {
				raf.Start()
			}
			ev.Call("preventDefault")
		}
		return nil
	})
	doc.Call("addEventListener", "keydown", keyHandler)

	// Keep WASM alive
	select {}
}
