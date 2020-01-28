package canvas

// en: Saves the state of the current context
//
// pt_br: Salva o estado atual do contexto atual
func (el *Canvas) Save() {
	el.SelfContext.Call("save")
}
