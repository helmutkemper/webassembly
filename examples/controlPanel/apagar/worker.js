// worker.js
self.importScripts('../../support/wasm_exec.js');

const go = new Go();
WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then((result) => {
    // run golang wasm code
    go.run(result.instance);

    // call function runInWorker() inside go wasm
    // result.instance.exports.runInWorker();

    self.onmessage = function(e) {
        console.log("onmessage ok")
        const name = e.data;
        const greeting = self.hello(name);
        self.postMessage(greeting);
    };

    // send message when go wasm is ready
    self.postMessage('runtimeReady');
}).catch(err => {
    console.error(err);
});
