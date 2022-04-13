package css

import (
	"strings"
)

// Class
//
// Português:
//
//  Arquiva listas de classes css.
//
// Uma list de classes css é basicamente a propriedade css, tipo, <div class="a b...">...</div> e o
// comando Toggle(name) muda a lista ativa conforme o nome passado.
type Class struct {
	list       map[string][]string
	activeName string
}

// String
//
// Português:
//
//  Converte a lista ativa em string.
func (el Class) String() (list string) {
	if el.list == nil {
		return
	}

	return strings.Join(el.list[el.activeName], " ")
}

// SetList
//
// Português:
//
//  Carrega ou substitui uma lista inteira.
//
//   Entrada:
//     name: nome da lista;
//     classes: lista de classes css a serem adicionadas.
func (el *Class) SetList(name string, classes ...string) {
	if el.list == nil {
		el.list = make(map[string][]string)
		el.activeName = name
	}

	if el.list[name] == nil {
		el.list[name] = make([]string, 0)
	}

	el.list[name] = append(el.list[name], classes...)
}

// AddToList
//
// Português:
//
//  Adiciona uma classe css a lista.
//
//   Entrada:
//     name: nome da list;
//     class: nome da classe css.
func (el *Class) AddToList(name, class string) {
	if el.list == nil {
		el.list = make(map[string][]string)
		el.activeName = name
	}

	if el.list[name] == nil {
		el.list[name] = make([]string, 0)
	}

	el.list[name] = append(el.list[name], class)
}

// RemoveFromList
//
// Português:
//
//  Remove uma classe css da lista de classes.
//
//   Entrada:
//     name: nome da list;
//     class: nome da classe css.
func (el *Class) RemoveFromList(name, class string) {
	if el.list == nil {
		el.list = make(map[string][]string)
		el.list[name] = make([]string, 0)
		return
	}

	if el.list[name] == nil {
		el.list[name] = make([]string, 0)
		return
	}

	for k, classInList := range el.list[name] {
		if classInList == class {
			el.list[name] = append(el.list[name][:k], el.list[name][k+1:]...)
		}
	}
}

// DeleteList
//
// Português:
//
// Apaga uma lista de classes css caso a lista apagada não seja a lista ativa.
//
//   Entrada:
//     name: nome da lista css.
//
//   Saída:
//     ok: true para comando executado com sucesso.
func (el *Class) DeleteList(name string) (ok bool) {
	if el.list == nil || el.activeName == name {
		return
	}

	ok = true
	delete(el.list, name)

	return
}

// Toggle
//
// Português:
//
// Muda o nome da lista ativa.
//
//   Entrada:
//     name: nome da lista css.
//
//   Saída:
//     ok: true para comando executado com sucesso.
func (el *Class) Toggle(name string) (ok bool) {
	if el.list == nil {
		return
	}

	el.activeName = name
	_, ok = el.list[name]
	return
}
