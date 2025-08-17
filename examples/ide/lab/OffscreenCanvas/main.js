// main.js — sem setar canvas.width/height após transferir para OffscreenCanvas

const canvas = document.getElementById('c');
const worker = new Worker('worker.js', { type: 'module' });

let dpr = window.devicePixelRatio || 1;

function sendResize() {
    const rect = canvas.getBoundingClientRect();
    const cssW = Math.max(1, Math.floor(rect.width));
    const cssH = Math.max(1, Math.floor(rect.height));
    const pxW  = Math.floor(cssW * dpr);
    const pxH  = Math.floor(cssH * dpr);

    // ❌ NÃO fazer:
    // canvas.width  = cssW;
    // canvas.height = cssH;

    // ✅ Deixa o worker ajustar o OffscreenCanvas:
    worker.postMessage({ type: 'resize', cssW, cssH, pxW, pxH, dpr });
}

function start() {
    const offscreen = canvas.transferControlToOffscreen();
    worker.postMessage({ type: 'init', canvas: offscreen, dpr }, [offscreen]);
    sendResize();
    requestAnimationFrame(tick);
}

function tick(now) {
    worker.postMessage({ type: 'tick', now });
    requestAnimationFrame(tick);
}

window.addEventListener('resize', () => {
    dpr = window.devicePixelRatio || 1;
    sendResize();
});

start();
