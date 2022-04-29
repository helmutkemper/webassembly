package html

type FetchPriority string

func (e FetchPriority) String() string {
	return string(e)
}

const (
	// KFetchPriorityHigh
	//
	// English:
	//
	//  Signals a high-priority fetch relative to other images.
	//
	// Português:
	//
	//  Sinaliza uma busca de alta prioridade em relação a outras imagens.
	KFetchPriorityHigh FetchPriority = "high"

	// KFetchPriorityLow
	//
	// English:
	//
	//  Signals a low-priority fetch relative to other images.
	//
	// Português:
	//
	//  Sinaliza uma busca de baixa prioridade em relação a outras imagens.
	KFetchPriorityLow FetchPriority = "low"

	// KFetchPriorityAuto
	//
	// English:
	//
	//  Signals automatic determination of fetch priority relative to other images.
	//
	// Português:
	//
	//  Sinaliza a determinação automática da prioridade de busca em relação a outras imagens.
	KFetchPriorityAuto FetchPriority = "auto"
)
