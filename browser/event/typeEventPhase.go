package event

type Phase int

const (
	// KPhaseNone
	//
	// English:
	//
	// The event is not being processed at this time.
	//
	// Português:
	//
	// O evento não está sendo processado neste momento.
	KPhaseNone Phase = 0

	// KPhaseCapturingPhase
	//
	// English:
	//
	// The event is being propagated through the target's ancestor objects.
	//
	// This process starts with the Window, then Document, then the HTMLHtmlElement, and so on through the elements until
	// the target's parent is reached. Event listeners registered for capture mode when EventTarget.addEventListener()
	// was called are triggered during this phase.
	//
	// Português:
	//
	// O evento está sendo propagado por meio dos objetos ancestrais do destino.
	//
	// Esse processo começa com Window, depois Document, depois HTMLHtmlElement e assim por diante pelos elementos até
	// que o pai do destino seja alcançado. Os ouvintes de eventos registrados para o modo de captura quando
	// EventTarget.addEventListener() foi chamado são acionados durante esta fase.
	KPhaseCapturingPhase Phase = 1

	// KPhaseAtTarget
	//
	// English:
	//
	// The event has arrived at the event's target.
	//
	// Event listeners registered for this phase are called at this time. If Event.bubbles is false, processing the event
	// is finished after this phase is complete.
	//
	// Português:
	//
	// O evento chegou ao destino do evento.
	//
	// Os ouvintes de eventos registrados para esta fase são chamados neste momento. Se Event.bubbles for false, o
	// processamento do evento será concluído após a conclusão desta fase.
	KPhaseAtTarget Phase = 2

	// KPhaseBubblingPhase
	//
	// English:
	//
	// The event is propagating back up through the target's ancestors in reverse order, starting with the parent, and
	// eventually reaching the containing Window. This is known as bubbling, and occurs only if Event.bubbles is true.
	// Event listeners registered for this phase are triggered during this process.
	//
	// Português:
	//
	// O evento está se propagando de volta pelos ancestrais do destino em ordem inversa, começando com o pai e,
	// eventualmente, atingindo a janela que o contém. Isso é conhecido como borbulhamento e ocorre somente se
	// Event.bubbles for verdadeiro. Os ouvintes de eventos registrados para esta fase são acionados durante esse
	// processo.
	KPhaseBubblingPhase Phase = 3
)

//
//
//
//
//
//
//
//
//
//
//
//
//
//
