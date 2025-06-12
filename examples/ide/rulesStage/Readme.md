# Stage

To run the application you need to create a html div of ID `rulesStage.KStageId` and then run the code.
By default, `rulesStage.KStageId` is defined as `graphicGopherIde`.

An example of simple application would be:

```html
<html>
<head>
    <meta charset="utf-8"/>
    <title>Graphic Gopher IDE</title>
    <style>
        body {
            margin: 0 !important;
            padding: 0 !important;
        }
    </style>
    <script src="../support/wasm_exec.js"></script>
    <script>
        document.addEventListener("DOMContentLoaded", () => {
            const go = new Go();
            WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then((result) => {
                go.run(result.instance);
            });
        });
    </script>
</head>
<body>
<div id="graphicGopherIde"></div>
</body>
</html>
```
