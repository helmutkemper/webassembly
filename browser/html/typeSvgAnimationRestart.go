package html

type SvgAnimationRestart string

func (e SvgAnimationRestart) String() string {
	return string(e)
}

const (
	// KSvgAnimationRestartAlways
	//
	// English:
	//
	//  This value indicates that the animation can be restarted at any time.
	//
	// Português:
	//
	//  Este valor indica que a animação pode ser reiniciada a qualquer momento.
	KSvgAnimationRestartAlways SvgAnimationRestart = "always"

	// KSvgAnimationRestartWhenNotActive
	//
	// English:
	//
	//  This value indicates that the animation can only be restarted when it is not active (i.e. after the active end).
	//
	// Attempts to restart the animation during its active duration are ignored.
	//
	// Português:
	//
	//  Este valor indica que a animação só pode ser reiniciada quando não estiver ativa (ou seja, após o término ativo).
	//
	// As tentativas de reiniciar a animação durante sua duração ativa são ignoradas.
	KSvgAnimationRestartWhenNotActive SvgAnimationRestart = "whenNotActive"

	// KSvgAnimationRestartNever
	//
	// English:
	//
	//  This value indicates that the animation cannot be restarted for the time the document is loaded.
	//
	// Português:
	//
	//  Esse valor indica que a animação não pode ser reiniciada durante o carregamento do documento.
	KSvgAnimationRestartNever SvgAnimationRestart = "never"
)
