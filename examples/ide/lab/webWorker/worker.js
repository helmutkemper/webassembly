/* worker.js */
self.importScripts('wasm_exec.js'); // clássico, não-ESM
const go = new Go();

let wasmStarted = false;
let canvas = null;

self.onmessage = (e) => {
    const msg = e.data;

    if (msg.type === 'init') {
        canvas = msg.canvas;
        // Expõe o OffscreenCanvas no escopo global do worker para o Go acessar
        self.canvas = canvas;

        if (!wasmStarted) {
            wasmStarted = true;
            WebAssembly.instantiateStreaming(fetch('main.wasm'), go.importObject)
                .then((res) => go.run(res.instance)); // inicia o runtime Go dentro do worker
        }
    }

    if (msg.type === 'resize' && canvas) {
        // Redimensione o OffscreenCanvas aqui (no worker), não o <canvas> original
        canvas.width = msg.width;
        canvas.height = msg.height;
    }
};
