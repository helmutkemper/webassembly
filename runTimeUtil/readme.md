# Trace

English:

Trace is a function designed to be used within error events, where it returns the name of the current function, the line and the name of the file where the trace function was called.

Português:

Trace é uma função idealizada para ser usada dentro dos eventos de erro, onde ela retorna o nome da função atual, a linha e o nome do arquivo onde a função trace foi chamada.

Example / Exemplo:
```go
if err = config.populate(tag); err != nil {
  file, line, funcName := runTimeUtil.Trace()
  err = errors.Join(fmt.Errorf("%v(line: %v).populate().error: %v", funcName, line, err))
  err = errors.Join(fmt.Errorf("file: %v", file), err)
  return
}
```

Output / Saída:
```
/Users/kemper/go/projetos/iotmaker.webassembly/examples/controlPanel/main.go
main.(*Admin).init(line: 409).makeControlCell().error: 
file: /Users/kemper/go/projetos/iotmaker.webassembly/examples/controlPanel/main.go
main.(*Admin).makeControlCell(line: 316).populate().error: 
file: /Users/kemper/go/projetos/iotmaker.webassembly/examples/controlPanel/main.go
wasm_exec.js:22 main.(*adminRangeInt).populate(line: 542).ParseInt(min).error: 
strconv.ParseInt: parsing "0a": invalid syntax
```
