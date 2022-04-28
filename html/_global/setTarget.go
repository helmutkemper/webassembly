package _global

import "log"

// SetTarget
//
// English:
//
// Where to display the linked URL, as the name for a browsing context (a tab, window, or <iframe>).
// The following keywords have special meanings for where to load the URL:
//
//   KTargetSelf: the current browsing context; (Default)
//   KTargetBlank: usually a new tab, but users can configure browsers to open a new window instead;
//   KTargetParent: the parent browsing context of the current one. If no parent, behaves as _self;
//   KTargetTop: the topmost browsing context (the "highest" context that's an ancestor of the current
//     one). If no ancestors, behaves as _self.
//
//   Note:
//     * Setting KTargetBlank on <a> elements implicitly provides the same rel behavior as setting
//       rel="noopener" which does not set window.opener. See browser compatibility for support
//       status.
//
// Português:
//
// Onde exibir a URL vinculada, como o nome de um contexto de navegação (uma guia, janela ou
// <iframe>). As seguintes palavras-chave têm significados especiais para onde carregar o URL:
//
//   KTargetSelf: o contexto de navegação atual; (padrão)
//   KTargetBlank: geralmente uma nova guia, mas os usuários podem configurar os navegadores para
//     abrir uma nova janela;
//   KTargetParent: o contexto de navegação pai do atual. Se nenhum pai, se comporta como _self;
//   KTargetTop: o contexto de navegação mais alto (o contexto "mais alto" que é um ancestral do
//     atual). Se não houver ancestrais, se comporta como _self.
//
//   Nota:
//     * Definir KTargetBlank em elementos <a> fornece implicitamente o mesmo comportamento rel
//       que definir rel="noopener" que não define window.opener. Consulte a compatibilidade do
//       navegador para obter o status do suporte.
func (e *GlobalAttributes) SetTarget(target Target) (ref *GlobalAttributes) {
	switch e.tag {
	case KTagA:
	case KTagForm:
	default:
		log.Printf("tag " + e.tag.String() + " does not support target property")
	}

	e.selfElement.Set("target", target.String())
	return e
}
