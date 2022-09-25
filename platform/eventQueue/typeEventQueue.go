package eventQueue

import (
	"log"
	"sync"
)

type EventQueue struct {
	queue          []Event
	buffer         []Event
	found          []int
	list           [][]Event
	listName       []Event
	listSize       int
	oppositeEvents [][2]Event

	activeEventsEmpty bool
	activeEvents      []Event
	activeEventsNow   bool

	sync.Mutex
}

// GetStatus
//
// English:
//
// Returns the last status, defined by the AddEvent() function.
//
//	Output:
//	  newEvent: new event since this function was last read;
//	  empty: empty event list
//	  activeEvents: list of events, with the last event being the first in the list.
//
// Português:
//
// Retorna o ultimo status, definido pela função AddEvent()
//
//	Saída:
//	  newEvent: novo evento desde a última leitura desta função;
//	  empty: lista de eventos limpa;
//	  activeEvents: lista de eventos, sendo o último evento, o primeiro da lista.
func (e *EventQueue) GetStatus() (newEvent, empty bool, activeEvents []Event) {
	e.Lock()
	defer e.Unlock()

	newEvent, empty, activeEvents = e.activeEventsNow, e.activeEventsEmpty, e.activeEvents
	e.activeEventsNow = false
	return
}

// AddSpecialEvent
//
// English:
//
// Adds a list of event combinations to generate a special event.
//
// Português:
//
// Adiciona uma lista de combinação de teclas para gerar um evento especial.
func (e *EventQueue) AddSpecialEvent(name string, list []Event) {
	e.Lock()
	defer e.Unlock()

	sizeList := len(list)
	if sizeList == 0 {
		log.Printf("bug: length of list is 0")
		return
	}

	newEvent := Event{}
	newEvent.Set(name, true, 0)

	if e.listSize < sizeList {
		e.listSize = sizeList
	}

	if len(e.list) == 0 {
		e.list = append(e.list, list)
		e.listName = append(e.listName, newEvent)
		return
	}

	var pass bool
	var index int
	var prevList []Event
	for index, prevList = range e.list {
		sizePrev := len(prevList)
		if sizePrev < sizeList {
			pass = true
			break
		}
	}

	if pass == false {
		index = len(e.list)
	}

	if len(e.list) == index {
		e.list = append(e.list, list)
		e.listName = append(e.listName, newEvent)
		return
	}

	e.list = append(e.list[:index+1], e.list[index:]...)
	e.list[index] = list

	e.listName = append(e.listName[:index+1], e.listName[index:]...)
	e.listName[index] = newEvent
}

// Init
//
// English:
//
// initializes the object
//
// Português:
//
// Inicializa o objeto
func (e *EventQueue) Init() {
	e.queue = make([]Event, 0)
	e.buffer = make([]Event, 0)
	e.list = make([][]Event, 0)
	e.listName = make([]Event, 0)
	e.oppositeEvents = make([][2]Event, 0)
	e.activeEvents = make([]Event, 0)
}

// filterOppositeEvents
//
// English:
//
// Filters opposite events such as right and left or up and down.
//
// Português:
//
// Filtra eventos opostos, como por exemplo, direita e esquerda ou sobe e desce.
func (e *EventQueue) filterOppositeEvents(eventList *[]Event, a, b int) {
	for k := range e.oppositeEvents {
		for kl := range *eventList {
			pass := false
			if (*eventList)[kl].label == e.oppositeEvents[k][a].label {
				for i := kl + 1; i != len(*eventList); i += 1 {
					if (*eventList)[i].label == e.oppositeEvents[k][b].label {
						e.deleteFromQueue(eventList, i)
						pass = true
						break
					}
				}
			}

			if pass == true {
				break
			}
		}
	}
}

// AddOppositeEvent
//
// English:
//
// Adds opposite events to be filtered, such as right and left.
//
//	Notes:
//	  * If this filter is used and the user activates the left and right events at the same time, the list will contain
//	    the last fired event.
//
// Português:
//
// Adiciona eventos opostos para serem filtradas, como por exemplo, direita e esquerda.
//
//	Notas:
//	  * Caso este filtro seja usado e o usuário ativar os eventos direita e esquerda ao mesmo tempo, a lista vai conter
//	    o último evento disparado.
func (e *EventQueue) AddOppositeEvent(eventA, eventB string) {
	e.Lock()
	defer e.Unlock()

	var newEventA Event
	newEventA.Set(eventA, true, 0)
	var newEventB Event
	newEventB.Set(eventB, true, 0)

	e.oppositeEvents = append(e.oppositeEvents, [2]Event{newEventA, newEventB})
}

// AddEvent
//
// English:
//
// Adds a event represented by a string.
//
//	Notes:
//	  * If AddOppositeEvent() filter is used and the user activates the left and right events at the same time, the list
//	    will contain the last fired event.
//
// Português:
//
// Adiciona um evento representado por uma string.
//
//	Notas:
//	  * Caso AddOppositeEvent() seja usado e o usuário ativar os eventos direita e esquerda ao mesmo tempo, a lista
//	    vai conter o último evento disparado.
func (e *EventQueue) AddEvent(event string, active bool) (empty bool, activeEvents []Event) {
	e.Lock()
	defer e.Unlock()

	var newEvent Event
	newEvent.Set(event, active, 0)
	e.queue = append(e.queue, newEvent)

	if e.listSize > 0 {
		e.buffer = append(e.buffer, newEvent)
		l := len(e.buffer)
		if l > e.listSize {
			e.buffer = e.buffer[l-e.listSize:]
		}
	}

	e.activeEventsNow = true

	empty, activeEvents = e.getActivesEvent()
	if len(activeEvents) > 1 {
		e.filterOppositeEvents(&activeEvents, 0, 1)
		e.filterOppositeEvents(&activeEvents, 1, 0)
	}

	if active == false {
		e.sanitizeDeactivateEvent(&e.queue, newEvent)
	}

	lb := len(e.buffer)
	for k := range e.list {
		// English: If the list is longer than buffer, it makes no sense to check
		// Português: Se a lista é maior do que buffer, não faz sentido verificar
		if lb < len(e.list[k]) {
			continue
		}

		// English: If it finds a special event combination, it only returns the name of the combination.
		// Português: Caso encontre uma combinação de teclas especiais, retorna apenas o nome da combinação.
		validTimeout, found := e.search(e.list[k])
		if validTimeout && found {
			//activeEvents = append([]Event{e.listName[k]}, activeEvents...)
			activeEvents = []Event{e.listName[k]}
			empty = false
			e.activeEvents = activeEvents
			e.activeEventsEmpty = empty
			return
		}
	}

	empty = len(activeEvents) == 0
	e.activeEvents = activeEvents
	e.activeEventsEmpty = empty
	return
}

func (e *EventQueue) Debug() {
	e.Lock()
	defer e.Unlock()

	log.Print(">>>")
	for _, v := range e.activeEvents {
		log.Printf("%v: %v", v.GetLabel(), v.GetActive())
	}
	log.Print("<<<")
}

// clear
//
// English:
//
// # Removes the list of found events
//
// Português:
//
// Remove a lista de eventos encontrados.
func (e *EventQueue) clearAfterSearch() {
	for len(e.found) > 0 {
		k := e.foundPop()
		e.buffer = append(e.buffer[:k], e.buffer[k+1:]...)
	}
}

// Clear
//
// English:
//
// # Clear all queue events
//
// Português:
//
// Limpa todos os eventos da fila
func (e *EventQueue) Clear() {
	e.Lock()
	defer e.Unlock()

	e.Init()
}

// foundPop
//
// English:
//
// pops the last found index
//
// Português:
//
// pops o último índice encontrado
func (e *EventQueue) foundPop() (index int) {
	k := len(e.found)
	if k > 0 {
		k -= 1
	}

	index, e.found = e.found[k], e.found[:k]
	return
}

// searchEvent
//
// English:
//
// Search the event and returns the index or -1.
//
// Português:
//
// Procura um evento e retorna o índice ou -1.
func (e *EventQueue) searchEvent(queue *[]Event, start int, event Event, active bool) (index int) {
	for i := start - 1; i != -1; i -= 1 {
		eventQ := (*queue)[i]
		if eventQ.label == event.label && eventQ.active == active {
			return i
		}
	}

	return -1
}

// SanitizeDeactivateEvent
//
// English:
//
// Searches for the on/off event pair, starting at the end of the queue, and removes it.
//
// Português:
//
// Procura pelo par de eventos ativo/desativo, começando pelo fim da fila, e a remove.
func (e *EventQueue) SanitizeDeactivateEvent(event Event) {
	e.Lock()
	defer e.Unlock()

	e.sanitizeDeactivateEvent(&e.queue, event)
}

// sanitizeDeactivateEvent
//
// English:
//
// Searches for the on/off event pair, starting at the end of the queue, and removes it.
//
// Português:
//
// Procura pelo par de eventos ativo/desativo, começando pelo fim da fila, e a remove.
func (e *EventQueue) sanitizeDeactivateEvent(queue *[]Event, event Event) {
	for i := len(*queue) - 1; i != -1; i -= 1 {
		eventQ := (*queue)[i]
		if !(eventQ.label == event.label && eventQ.active == false) {
			continue
		}

		index := e.searchEvent(queue, i, eventQ, !eventQ.active)
		if index == -1 {
			return
		}

		e.deleteFromQueue(queue, i)
		e.deleteFromQueue(queue, index)

		return
	}
}

// deleteFromQueue
//
// English:
//
// Given an index, removes a event from the queue.
//
// Português:
//
// Dado um índice, remove um evento da fila.
func (e *EventQueue) deleteFromQueue(queue *[]Event, index int) {
	*queue = append((*queue)[:index], (*queue)[index+1:]...)
}

// search
//
// English:
//
// Searches for, and removes if found in the buffer, an event specified in the list
//
//	Notes:
//	  * Start from the largest list to the smallest;
//	  * In the case of keyboard, if the list only has `down` keys, the list accepts simultaneous keys, to avoid this,
//	    put a `down` key followed by an `up` key;
//	  * A found list is removed from the event queue, even if it times out.
//
// Português:
//
// Procura, e remove caso encontre no buffer, um evento especificado na lista
//
//	Notas:
//	  * Comece da maior lista para a menor;
//	  * No caso de teclas, se a lista só tiver as teclas `down`, a lista aceita teclas simultâneas, para evitar isto,
//	    coloque uma tecla `down` seguida de uma `up`;
//	  * Uma lista encontrada é removida da fila de eventos, mesmo que o tempo limite tenha excedido.
func (e *EventQueue) search(list []Event) (validTimeout bool, found bool) {
	validTimeout = true

	lenQueue := len(e.buffer)
	lenList := len(list)

	if lenQueue < lenList {
		return false, false
	}

	e.found = make([]int, 0)
	kQueue := 0
	for {
		for kList := range list {
			eventQ := e.buffer[kQueue]
			eventL := list[kList]
			kQueue += 1
			if eventQ.label == eventL.label && eventQ.active == eventL.active {
				found = true
				e.found = append(e.found, kQueue-1)

				if kQueue > 1 && eventL.timeout != 0 {
					eventPrevious := e.buffer[kQueue-2]
					if eventQ.date.Sub(eventPrevious.date) > eventL.timeout {
						validTimeout = false
					}
				}

				continue
			}

			found = false
			e.found = make([]int, 0)
			break
		}

		if found == true {
			e.clearAfterSearch()
			return
		}

		if kQueue >= lenQueue {
			return false, false
		}

		if kQueue+lenList > lenQueue {
			return false, false
		}
	}
}

// getActivesEvent
//
// English:
//
// Returns the list of the last active events.
//
//	Notes:
//	  * The zero index is the last event fired;
//	  * If the user fires opposite events like right and left at the same time, use the lowest index as the valid event;
//	  * The event list is not affected.
//
// Português:
//
// Retorna a lista das últimas teclas ativas.
//
//	Notas:
//	  * O índice zero é o último evento ocorrido;
//	  * Caso o usuário dispara eventos opostos, como direita e esquerda ao mesmo tempo, use a índice mais baixo como
//	    sendo o evento válido;
//	  * A lista de eventos não é afetada.
func (e *EventQueue) getActivesEvent() (empty bool, eventList []Event) {
	empty = len(e.queue) == 0
	if empty {
		return
	}

	eventList = make([]Event, 0)

	var event Event
	queue := make([]Event, len(e.queue))
	copy(queue, e.queue)

	for {
		event = e.lastEventQueue(&queue)
		if event.active == false {
			e.sanitizeDeactivateEvent(&queue, event)
		} else {
			eventList = append(eventList, event)
			e.popQueue(&queue)
		}

		if len(queue) == 0 {
			return
		}
	}
}

// lastEventQueue
//
// English:
//
// Returns the last event in the queue.
//
// Português:
//
// Retorna o último evento da fila.
func (e *EventQueue) lastEventQueue(queue *[]Event) (event Event) {
	l := len(*queue)
	if l != 0 {
		l -= 1
	}

	event = (*queue)[l]
	return
}

// popQueue
//
// English:
//
// Returns the last element in the queue, reducing its size.
//
// Português:
//
// Retorna o último elemento da fila, reduzindo o tamanho da mesma.
func (e *EventQueue) popQueue(queue *[]Event) (event Event) {
	l := len(*queue)
	if l != 0 {
		l -= 1
	}

	event, *queue = (*queue)[l], (*queue)[:l]
	return
}

// IsEmpty
//
// English:
//
// true when the event queue is empty.
//
// Português:
//
// true quando a fila de eventos estiver vazia.
func (e *EventQueue) IsEmpty() (empty bool) {
	e.Lock()
	defer e.Unlock()

	return e.activeEventsEmpty
}
