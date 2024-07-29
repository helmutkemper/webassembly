onmessage = function(e) {
    if (e.data === 'start') {
        runLoop();
    } else if (e.data === 'stop') {
        stopLoop();
    }
};

let stop = false;

function runLoop() {
    stop = false;
    function loop() {
        if (stop) {
            return;
        }
        myFunc(); // Chama a função myFunc exportada do WebAssembly
        setTimeout(loop, 0); // Evita bloqueio da thread principal
    }
    loop();
}

function stopLoop() {
    stop = true;
}

function myFunc() {
    easingTween.Worker();
}

