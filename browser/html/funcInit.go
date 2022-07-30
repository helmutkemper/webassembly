package html

import "sync"

// htmlGlobalAllElementsList
//
// English:
//
// Saves the reference of all elements with ID for later use.
//
// Português:
//
// Salva a referência de todos os elementos com ID para uso posterior.
var htmlGlobalAllElementsList = new(sync.Map)
