package html

// Append
//
// English:
//
// Adds a node to the end of the list of children of a specified parent node. If the node already exists in the document, it is removed from its current parent node before being added to the new parent.
//
//   Input:
//     append: element in js.Value format.
//
//   Note:
//     * The equivalent of:
//         var p = document.createElement("p");
//         document.body.appendChild(p);
//
// Português:
//
//  Adiciona um nó ao final da lista de filhos de um nó pai especificado. Se o nó já existir no documento, ele é removido de seu nó pai atual antes de ser adicionado ao novo pai.
//
//   Entrada:
//     appendId: elemento no formato js.Value.
//
//   Nota:
//     * Equivale a:
//         var p = document.createElement("p");
//         document.body.appendChild(p);
func (e *GlobalAttributes) Append(append interface{}) (ref *GlobalAttributes) {
	append.(GlobalAttributes).selfElement.Call("appendChild", e.selfElement)
	return e
}
