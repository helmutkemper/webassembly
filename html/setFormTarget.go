package html

import "log"

// SetFormTarget
//
// English:
//
//  If the button is a submit button, this attribute is an author-defined name or standardized, underscore-prefixed keyword indicating where to display the response from submitting the form. This is the name of, or keyword for, a browsing context (a tab, window, or <iframe>). If this attribute is specified, it overrides the target attribute of the button's form owner. The following keywords have special meanings:
//
//   KTargetSelf: the current browsing context; (Default)
//   KTargetBlank: usually a new tab, but users can configure browsers to open a new window instead;
//   KTargetParent: the parent browsing context of the current one. If no parent, behaves as _self;
//   KTargetTop: the topmost browsing context (the "highest" context that's an ancestor of the current
//     one). If no ancestors, behaves as _self.
//
// Português:
//
//   KTargetSelf: o contexto de navegação atual; (padrão)
//   KTargetBlank: geralmente uma nova guia, mas os usuários podem configurar os navegadores para
//     abrir uma nova janela;
//   KTargetParent: o contexto de navegação pai do atual. Se nenhum pai, se comporta como _self;
//   KTargetTop: o contexto de navegação mais alto (o contexto "mais alto" que é um ancestral do
//     atual). Se não houver ancestrais, se comporta como _self.
func (e *GlobalAttributes) SetFormTarget(formtarget Target) (ref *GlobalAttributes) {
	switch e.tag {
	case KTagButton:
	default:
		log.Printf("tag " + e.tag.String() + " does not support formtarget property")
	}

	e.selfElement.Set("formtarget", formtarget.String())
	return e
}
