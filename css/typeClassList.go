package css

import (
	"log"
	"strings"
	"syscall/js"
	"time"
)

const (
	KCssListNameNotFound = "html.CssToggle().error: css list name not found:"
)

// Class
//
// English:
//
//  Allows you to easily change the cc of an HTML element
//
//   Exemplo:
//     var class = new(css.Class)
//     // Create a css list named "red" with value "user red"
//     class.SetList("red", "user", "red").
//       // Create a css list named "yellow" with value "user yellow"
//       SetList("yellow", "user", "yellow").
//       // Create a css list named "user" with value "user"
//       SetList("user", "user").
//       // Defines that the "red" and "yellow" lists will change every second
//       ToggleTime(time.Second, "red", "yellow").
//       // Limit trades to 10 interactions
//       ToggleLoop(10).
//       // Defines the list named "norm" as the active list at the end of interactions
//       OnLoopEnd("user").
//       // Start interactions. Caution: they only work after being added to the element
//       ToggleStart()
//
//     var a html.Div
//     // Create a div with id "example_A";
//     a.NewDiv("example_A").
//       // Sets css to be "name_a name_b name_N";
//       Css("name_a", "name_b", "name_N").
//       // Adds the div to the element id "stage".
//       AppendById("stage")
//
//     var b html.Div
//     // Create a div with id "example_B";
//     b.NewDiv("example_B").
//       // Sets css to be "name_a name_b name_N";
//       Css("name_a", "name_b", "name_N").
//       // css.Class cannot work properly before being added, due to lack of reference to the parent
//       // object.
//       SetCssController(class).
//       // Adds the div to the element id "stage".
//       AppendById("stage")
//
// Português:
//
//  Permite alterar o cc de um elemento HTML de forma fácil
//
//   Exemplo:
//     var class = new(css.Class)
//     // Crie uma lista css de nome "red" com o valor "user red"
//     class.SetList("red", "user", "red").
//       // Crie uma lista css de nome "yellow" com o valor "user yellow"
//       SetList("yellow", "user", "yellow").
//       // Crie uma lista css de nome "user" com o valor "user"
//       SetList("user", "user").
//       // Define que as listas "red" e "yellow" vão trocar a cada segundo
//       ToggleTime(time.Second, "red", "yellow").
//       // Limita as trocas em 10 interações
//       ToggleLoop(10).
//       // Define  alista de nome "normal" como sendo a lista ativa ao final das interações
//       OnLoopEnd("user").
//       // Inicia as interações. Cuidado: elas só funcionam após serem adicionadas ao elemento
//       ToggleStart()
//
//     var a html.Div
//     // Cria uma div de id "example_A";
//     a.NewDiv("example_A").
//       // Define css como sendo "name_a name_b name_N";
//       Css("name_a", "name_b", "name_N").
//       // Adds the div to the element id "stage".
//       // Adiciona a div ao elemento de id "stage".
//       AppendById("stage")
//
//     var b html.Div
//     // Cria uma div de id "example_B";
//     b.NewDiv("example_B").
//       // Define css como sendo "name_a name_b name_N";
//       Css("name_a", "name_b", "name_N").
//       // css.Class não consegue funcionar corretamente antes de ser adicionada, por falta de
//       // referência do objeto pai.
//       SetCssController(class).
//       // Adiciona a div ao elemento de id "stage".
//       AppendById("stage")
type Class struct {
	refElement *js.Value
	id         string

	list          map[string][]string
	activeName    string
	counter       int
	max           int
	toggleList    []string
	timeout       *time.Ticker
	interval      time.Duration
	done          chan struct{}
	loop          int
	loopFlag      bool
	onLoopEnd     string
	onLoopEndFunc func(name string)
	onToggleFunc  func(name string)
}

// Remove
//
// English:
//
//  This function must be called when the object is removed.
//
// Português:
//
//  Esta função deve ser chamada quando o objeto for removido.
func (e *Class) Remove() {
	e.done <- struct{}{}
}

// String
//
// Português:
//
//  Converte a lista ativa em string.
func (e Class) String() (list string) {
	if e.list == nil {
		return
	}

	return strings.Join(e.list[e.activeName], " ")
}

// SetOnLoopEndFunc
//
// English:
//
//  Defines the function to be executed at the end of interactions
//
//   Input:
//     f: function to be executed at the end of interactions.
//       listName: current list name.
//
//   Note:
//     * This function is equivalent to html.SetList("current", classes...)
//     * Css has a feature that allows you to easily change the list of css classes of an html tag,
//       with the functions SetList(), CssToggle() and CssToggleTime();
//     * Is the equivalent of <... css="name1 name2 nameN">
//
// Português:
//
//  Define a função a ser executada ao fim das interações
//
//   Entrada:
//     f: função a ser executada ao fim das interações.
//       listName: nome da lista atual.
//
//   Nota:
//     * Esta função equivale a SetList("current", classes...);
//     * Css tem uma funcionalidade que permite trocar a lista de classes css de uma tag html de forma
//       fácil, com as funções SetList(), CssToggle() e CssToggleTime();
//     * Equivale a <... css="name1 name2 nameN">
func (e *Class) SetOnLoopEndFunc(f func(listName string)) (ref *Class) {
	e.onLoopEndFunc = f

	return e
}

// SetOnToggleFunc
//
// English:
//
//  Defines the function to be performed at each iteration.
//
//   Input:
//     f: function to be executed at each iteration.
//       listName: current list name.
//
//   Note:
//     * This function is equivalent to html.SetList("current", classes...)
//     * Css has a feature that allows you to easily change the list of css classes of an html tag,
//       with the functions SetList(), CssToggle() and CssToggleTime();
//     * Is the equivalent of <... css="name1 name2 nameN">
//
// Português:
//
//  Define a função a ser executada a cada interação.
//
//   Entrada:
//     f: função a ser executada a cada interação.
//       listName: nome da lista atual.
//
//   Nota:
//     * Esta função equivale a SetList("current", classes...);
//     * Css tem uma funcionalidade que permite trocar a lista de classes css de uma tag html de forma
//       fácil, com as funções SetList(), CssToggle() e CssToggleTime();
//     * Equivale a <... css="name1 name2 nameN">
func (e *Class) SetOnToggleFunc(f func(listName string)) (ref *Class) {
	e.onToggleFunc = f

	return e
}

// SetList
//
// English:
//
//  Adds a new list of css classes.
//
//   Input:
//     name: css list name;
//     classes: css class list.
//
//   Note:
//     * This function is equivalent to html.SetList("current", classes...)
//     * Css has a feature that allows you to easily change the list of css classes of an html tag,
//       with the functions SetList(), CssToggle() and CssToggleTime();
//     * Is the equivalent of <... css="name1 name2 nameN">
//
// Português:
//
//  Adiciona uma nova lista de classes css.
//
//   Entrada:
//     name: nome da lista css;
//     classes: lista de classes css.
//
//   Nota:
//     * Esta função equivale a SetList("current", classes...);
//     * Css tem uma funcionalidade que permite trocar a lista de classes css de uma tag html de forma
//       fácil, com as funções SetList(), CssToggle() e CssToggleTime();
//     * Equivale a <... css="name1 name2 nameN">
func (e *Class) SetList(name string, classes ...string) (ref *Class) {
	if e.list == nil {
		e.list = make(map[string][]string)
		e.activeName = name
	}

	if e.list[name] == nil {
		e.list[name] = make([]string, 0)
	}

	e.list[name] = append(e.list[name], classes...)

	if e.refElement != nil {
		e.refElement.Set("classList", e.String())
	}

	return e
}

// SetRef
//
// English:
//
//  References the parent object.
//
//   Input:
//     id: parent object id.
//     refElement: object reference.
//
//   Note:
//     * This function is equivalent to html.SetList("current", classes...)
//     * Css has a feature that allows you to easily change the list of css classes of an html tag,
//       with the functions SetList(), CssToggle() and CssToggleTime();
//     * Is the equivalent of <... css="name1 name2 nameN">
//
// Português:
//
//  Referencia o objeto pai.
//
//   Entrada:
//     id: id do objeto pai.
//     refElement: referência do objeto.
//
//   Nota:
//     * Esta função equivale a SetList("current", classes...);
//     * Css tem uma funcionalidade que permite trocar a lista de classes css de uma tag html de forma
//       fácil, com as funções SetList(), CssToggle() e CssToggleTime();
//     * Equivale a <... css="name1 name2 nameN">
func (e *Class) SetRef(id string, refElement *js.Value) (ref *Class) {
	e.id = id
	e.refElement = refElement

	return e
}

// AddToList
//
// English:
//
//  Add a class name to given list of classes.
//
//   Input:
//     name: name of the class list;
//     class: name of the class to be added.
//
//   Note:
//     * This function is equivalent to html.SetList("current", classes...)
//     * Css has a feature that allows you to easily change the list of css classes of an html tag,
//       with the functions SetList(), CssToggle() and CssToggleTime();
//     * Is the equivalent of <... css="name1 name2 nameN">
//
// Português:
//
//  Adiciona o nome de uma classe a uma determinada lista de classes.
//
//   Entrada:
//     name: nome da lista de classes;
//     class: nome da classe a ser adicionada.
//
//   Nota:
//     * Esta função equivale a SetList("current", classes...);
//     * Css tem uma funcionalidade que permite trocar a lista de classes css de uma tag html de forma
//       fácil, com as funções SetList(), CssToggle() e CssToggleTime();
//     * Equivale a <... css="name1 name2 nameN">
func (e *Class) AddToList(name, class string) (ref *Class) {
	if e.list == nil {
		e.list = make(map[string][]string)
		e.activeName = name
	}

	if e.list[name] == nil {
		e.list[name] = make([]string, 0)
	}

	e.list[name] = append(e.list[name], class)
	return e
}

// RemoveFromList
//
// English:
//
//  Removes a class name from a given list of classes.
//
//   Input:
//     name: name of the class list;
//     class: name of the class to be removed.
//
//   Note:
//     * This function is equivalent to html.SetList("current", classes...)
//     * Css has a feature that allows you to easily change the list of css classes of an html tag,
//       with the functions SetList(), CssToggle() and CssToggleTime();
//     * Is the equivalent of <... css="name1 name2 nameN">
//
// Português:
//
//  Remove o nome de uma classe de uma determinada lista de classes.
//
//   Entrada:
//     name: nome da lista de classes;
//     class: nome da classe a ser removida.
//
//   Nota:
//     * Esta função equivale a SetList("current", classes...);
//     * Css tem uma funcionalidade que permite trocar a lista de classes css de uma tag html de forma
//       fácil, com as funções SetList(), CssToggle() e CssToggleTime();
//     * Equivale a <... css="name1 name2 nameN">
func (e *Class) RemoveFromList(name, class string) (ref *Class) {
	if e.list == nil {
		e.list = make(map[string][]string)
		e.list[name] = make([]string, 0)
		return e
	}

	if e.list[name] == nil {
		e.list[name] = make([]string, 0)
		return e
	}

	for k, classInList := range e.list[name] {
		if classInList == class {
			e.list[name] = append(e.list[name][:k], e.list[name][k+1:]...)
		}
	}

	return e
}

// DeleteList
//
// English:
//
//  Removes a list of classes.
//
//   Input:
//     name: name of the list of classes to be removed.
//
//   Note:
//     * This function is equivalent to html.SetList("current", classes...)
//     * Css has a feature that allows you to easily change the list of css classes of an html tag,
//       with the functions SetList(), CssToggle() and CssToggleTime();
//     * Is the equivalent of <... css="name1 name2 nameN">
//
// Português:
//
//  Remove uma lista de classes.
//
//   Entrada:
//     name: nome da lista de classes a ser removida.
//
//   Nota:
//     * Esta função equivale a SetList("current", classes...);
//     * Css tem uma funcionalidade que permite trocar a lista de classes css de uma tag html de forma
//       fácil, com as funções SetList(), CssToggle() e CssToggleTime();
//     * Equivale a <... css="name1 name2 nameN">
func (e *Class) DeleteList(name string) (ref *Class) {
	if e.list == nil || e.activeName == name {
		log.Print(KCssListNameNotFound, name)
		return e
	}

	delete(e.list, name)

	return e
}

// Toggle
//
// English:
//
//  Swap the element's css list of classes.
//
//   Entrada:
//     name: name of the list of classes to use.
//
//   Note:
//     * This function is equivalent to html.SetList("current", classes...)
//     * Css has a feature that allows you to easily change the list of css classes of an html tag,
//       with the functions SetList(), CssToggle() and CssToggleTime();
//     * Is the equivalent of:
//         <... css="name1 name2 nameN">
//         var element = document.getElementById("myDIV");
//         element.classList.toggle("myStyle");
//
// Português:
//
//  Troca a lista de classes css do elemento.
//
//   Entrada:
//     name: nome da lista de classes a ser usada.
//
//   Nota:
//     * Esta função equivale a SetList("current", classes...);
//     * Css tem uma funcionalidade que permite trocar a lista de classes css de uma tag html de forma
//       fácil, com as funções SetList(), CssToggle() e CssToggleTime();
//     * Equivale a:
//         <... css="name1 name2 nameN">
//         var element = document.getElementById("myDIV");
//         element.classList.toggle("myStyle");
func (e *Class) Toggle(name string) (ref *Class) {
	if e.list == nil {
		return
	}

	e.activeName = name
	_, ok := e.list[name]

	if ok == false {
		log.Print(KCssListNameNotFound, name)
	}

	js.Global().Get("document").Call("getElementById", e.id).Set("classList", e.String())

	if e.onToggleFunc != nil {
		e.onToggleFunc(name)
	}
	return e
}

// ToggleTime
//
// English:
//
//  Swap the element's css list of classes.
//
//   Input:
//     interval: time interval between toggle;
//     name: names of the lists to be used in the toggle.
//
//   Note:
//     * This function is equivalent to html.SetList("current", classes...)
//     * Css has a feature that allows you to easily change the list of css classes of an html tag,
//       with the functions SetList(), CssToggle() and CssToggleTime();
//     * Is the equivalent of:
//         <... css="name1 name2 nameN">
//         var element = document.getElementById("myDIV");
//         element.classList.toggle("myStyle");
//
// Português:
//
//  Troca a lista de classes css do elemento.
//
//   Entrada:
//     interval: intervalo de tempo entre as trocas;
//     list: nome das listas a serem usadas na trocas.
//
//   Nota:
//     * Esta função equivale a SetList("current", classes...);
//     * Css tem uma funcionalidade que permite trocar a lista de classes css de uma tag html de forma
//       fácil, com as funções SetList(), CssToggle() e CssToggleTime();
//     * Equivale a:
//         <... css="name1 name2 nameN">
//         var element = document.getElementById("myDIV");
//         element.classList.toggle("myStyle");
func (e *Class) ToggleTime(interval time.Duration, list ...string) (ref *Class) {
	e.interval = interval
	e.toggleList = append(e.toggleList, list...)
	e.max = len(list) - 1
	e.done = make(chan struct{})

	return e
}

// ToggleList
//
// English:
//
//  Defines a new list.
//
//   Input:
//     name: names of the lists to be used in the toggle.
//
//   Note:
//     * This function is equivalent to html.SetList("current", classes...)
//     * Css has a feature that allows you to easily change the list of css classes of an html tag,
//       with the functions SetList(), CssToggle() and CssToggleTime();
//     * Is the equivalent of:
//         <... css="name1 name2 nameN">
//         var element = document.getElementById("myDIV");
//         element.classList.toggle("myStyle");
//
// Português:
//
//  Define uma nova lista.
//
//   Entrada:
//     list: nome das listas a serem usadas na trocas.
//
//   Nota:
//     * Esta função equivale a SetList("current", classes...);
//     * Css tem uma funcionalidade que permite trocar a lista de classes css de uma tag html de forma
//       fácil, com as funções SetList(), CssToggle() e CssToggleTime();
//     * Equivale a:
//         <... css="name1 name2 nameN">
//         var element = document.getElementById("myDIV");
//         element.classList.toggle("myStyle");
func (e *Class) ToggleList(list ...string) (ref *Class) {
	e.toggleList = append(e.toggleList, list...)
	e.max = len(list) - 1
	e.done = make(chan struct{})

	return e
}

// ToggleStop
//
// English:
//
//  Breaks the toggle loop between css classes.
//
//   Note:
//     * This function is equivalent to html.SetList("current", classes...)
//     * Css has a feature that allows you to easily change the list of css classes of an html tag,
//       with the functions SetList(), CssToggle() and CssToggleTime();
//     * Is the equivalent of <... css="name1 name2 nameN">
//
// Português:
//
//  Interrompe o laço de trocas entre classes css.
//
//   Nota:
//     * Esta função equivale a SetList("current", classes...);
//     * Css tem uma funcionalidade que permite trocar a lista de classes css de uma tag html de forma
//       fácil, com as funções SetList(), CssToggle() e CssToggleTime();
//     * Equivale a <... css="name1 name2 nameN">
func (e *Class) ToggleStop() (ref *Class) {
	e.done <- struct{}{}
	return e
}

// ToggleStartInterval
//
// English:
//
//  Start css class toggle functionality.
//
//   Entrada:
//     interval: intervalo de tempo entre as trocas.
//
//   Note:
//     * This function is equivalent to html.SetList("current", classes...)
//     * Css has a feature that allows you to easily change the list of css classes of an html tag,
//       with the functions SetList(), CssToggle() and CssToggleTime();
//     * Is the equivalent of <... css="name1 name2 nameN">
//
// Português:
//
//  Reinicializa a funcionalidade de troca de classes css após usar a função CssToggleStop().
//
//   Entrada:
//     interval: intervalo de tempo entre as trocas.
//
//   Nota:
//     * Esta função equivale a SetList("current", classes...);
//     * Css tem uma funcionalidade que permite trocar a lista de classes css de uma tag html de forma
//       fácil, com as funções SetList(), CssToggle() e CssToggleTime();
//     * Equivale a <... css="name1 name2 nameN">
func (e *Class) ToggleStartInterval(interval time.Duration) (ref *Class) {
	e.interval = interval
	e.ToggleStart()

	return e
}

// ToggleLoop
//
// English:
//
//  Defines a finite number of interactions.
//
//   Input:
//     loop: number of interactions.
//
//   Note:
//     * This function is equivalent to html.SetList("current", classes...)
//     * Css has a feature that allows you to easily change the list of css classes of an html tag,
//       with the functions SetList(), CssToggle() and CssToggleTime();
//     * Is the equivalent of <... css="name1 name2 nameN">
//
// Português:
//
//  Define um número finito de interações.
//
//   Entrada:
//     loop: número de interações.
//
//   Nota:
//     * Esta função equivale a SetList("current", classes...);
//     * Css tem uma funcionalidade que permite trocar a lista de classes css de uma tag html de forma
//       fácil, com as funções SetList(), CssToggle() e CssToggleTime();
//     * Equivale a <... css="name1 name2 nameN">
func (e *Class) ToggleLoop(loop int) (ref *Class) {
	e.loop = loop
	e.loopFlag = true

	return e
}

// OnLoopEnd
//
// English:
//
//  Defines the name of the css list to be used at the end of the loop.
//
//   Input:
//     name: css list name.
//
//   Note:
//     * This function is equivalent to html.SetList("current", classes...)
//     * Css has a feature that allows you to easily change the list of css classes of an html tag,
//       with the functions SetList(), CssToggle() and CssToggleTime();
//     * Is the equivalent of <... css="name1 name2 nameN">
//
// Português:
//
//  Define o nome da lista css a ser usada no final do loop.
//
//   Entrada:
//     name: nome da lista css.
//
//   Nota:
//     * Esta função equivale a SetList("current", classes...);
//     * Css tem uma funcionalidade que permite trocar a lista de classes css de uma tag html de forma
//       fácil, com as funções SetList(), CssToggle() e CssToggleTime();
//     * Equivale a <... css="name1 name2 nameN">
func (e *Class) OnLoopEnd(name string) (ref *Class) {
	e.onLoopEnd = name

	return e
}

// ToggleStart
//
// English:
//
//  Start css class toggle functionality.
//
//   Note:
//     * This function is equivalent to html.SetList("current", classes...)
//     * Css has a feature that allows you to easily change the list of css classes of an html tag,
//       with the functions SetList(), CssToggle() and CssToggleTime();
//     * Is the equivalent of <... css="name1 name2 nameN">
//
// Português:
//
//  Reinicializa a funcionalidade de troca de classes css após usar a função CssToggleStop().
//
//   Nota:
//     * Esta função equivale a SetList("current", classes...);
//     * Css tem uma funcionalidade que permite trocar a lista de classes css de uma tag html de forma
//       fácil, com as funções SetList(), CssToggle() e CssToggleTime();
//     * Equivale a <... css="name1 name2 nameN">
func (e *Class) ToggleStart() (ref *Class) {

	e.timeout = time.NewTicker(e.interval)
	go func(e *Class) {
		for {
			select {
			case <-e.done:
				e.timeout.Stop()
				return

			case <-e.timeout.C:

				name := e.toggleList[e.counter]
				e.Toggle(name)

				if e.counter == e.max {
					e.counter = 0
				} else {
					e.counter += 1
				}

				if e.loopFlag == true {
					if e.loop == 0 {

						e.loopFlag = false
						e.timeout.Stop()
						if e.onLoopEnd != "" {
							e.Toggle(e.onLoopEnd)
							if e.onLoopEndFunc != nil {
								e.onLoopEndFunc(e.onLoopEnd)
							}
						}
						return

					} else {
						e.loop -= 1
					}

				}
			}
		}
	}(e)

	return e
}
