package canvas

// en: Returns previously saved path state and attributes
//
// pt_br: Restaura o contexto e atributos previamente salvos
func (el *Canvas) Restore() {
	el.SelfContext.Call("restore")
}
