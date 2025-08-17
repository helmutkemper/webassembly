// Carrega wasm_exec.js do Go e inicializa o runtime
importScripts('wasm_exec.js');

const go = new Go();
let wasmReady = null;

// Estado simples para o Go consumir via js.Global()
self.wasmEnv = {
    canvas: null,
    dpr: 1,
    cssW: 1,
    cssH: 1,
    pxW: 1,
    pxH: 1,
};

// Encaminha mensagens do main thread
self.onmessage = async (e) => {
    const msg = e.data;

    if (msg.type === 'init') {
        self.wasmEnv.dpr = msg.dpr || 1;
        self.wasmEnv.canvas = msg.canvas;

        // Importante: inicialize o WASM após receber o canvas.
        if (!wasmReady) {
            wasmReady = WebAssembly.instantiateStreaming(fetch('main.wasm'), go.importObject)
                .then(async (result) => {
                    // Expõe wasmEnv globalmente para o Go ler
                    self.globalThis.wasmEnv = self.wasmEnv;
                    await go.run(result.instance);
                })
                .catch((err) => console.error('WASM init error:', err));
        }
    }

    else if (msg.type === 'resize') {
        if (self.wasmEnv.canvas) {
            self.wasmEnv.cssW = msg.cssW;
            self.wasmEnv.cssH = msg.cssH;
            self.wasmEnv.pxW = msg.pxW;
            self.wasmEnv.pxH = msg.pxH;
            self.wasmEnv.dpr = msg.dpr || self.wasmEnv.dpr;
            // Ajusta o backbuffer no OffscreenCanvas
            self.wasmEnv.canvas.width = self.wasmEnv.pxW;
            self.wasmEnv.canvas.height = self.wasmEnv.pxH;

            // Se o Go já registrou o hook de resize, avise-o
            if (self.onWasmResize) self.onWasmResize();
        }
    }

    else if (msg.type === 'tick') {
        // Encaminha "rAF" para o Go (vsync do main thread)
        if (self.onWasmTick) self.onWasmTick(msg.now);
    }
};
