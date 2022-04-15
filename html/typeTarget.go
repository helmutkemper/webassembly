package html

type Target string

func (e Target) String() string {
	return string(e)
}

const (
	// KTargetSelf
	//
	// English:
	//
	//  The current browsing context; (default)
	//
	// Português:
	//
	//  O contexto de navegação atual; (padrão)
	//
	KTargetSelf Target = "_self"

	// KTargetBlank
	//
	// English:
	//
	//  Usually a new tab, but users can configure browsers to open a new window instead;
	//
	// Português:
	//
	//  Normalmente, uma nova guia, mas os usuários podem configurar os navegadores para abrir uma nova
	//  janela;
	//
	KTargetBlank Target = "_blank"

	// KTargetParent
	//
	// English:
	//
	//  The parent browsing context of the current one. If no parent, behaves as _self;
	//
	// Português:
	//
	//  O contexto de navegação pai do atual. Se nenhum pai, se comporta como _self;
	//
	KTargetParent Target = "_parent"

	// KTargetTop
	//
	// English:
	//
	//  The topmost browsing context (the "highest" context that's an ancestor of the current one).
	//  If no ancestors, behaves as _self.
	//
	// Português:
	//
	//  O contexto de navegação mais alto (o contexto "mais alto" que é um ancestral do atual).
	//  Se não houver ancestrais, se comporta como _self.
	//
	KTargetTop Target = "_top"
)
