// Package globalEngine
//
// English:
//
// Engine is an object responsible for making the calculations to support graphics and animations.
//
// Due to a series of limitations, and compatibility issues between browsers, the use of non-native calculation
// functions to the browser was chosen to animate the graphic elements, allowing greater freedom of use.
//
// Note that changing the engine to 24pfs or 240fps does not affect the movement of native animations and does not
// always affect the quality of the animations controlled by the engine, due to a limitation presented by browsers
// (07/2022).
//
// The algorithm considers the response time to adjust fps between the maximum and minimum values for better processing,
// however, if the fps rate is too high, the browser crashes (07/2022).
//
// Português:
//
// Engine é um objeto responssável por fazer os cálculos de suporte aos elementos gráficos e animações.
//
// Devido à uma série de limitações, e problemas de compatibilidade entre navegadores, foi escolhido o uso de funções de
// cálculos não nativos ao navegador para animar os elementos gráficos, possibilitando maior liberdade de uso.
//
// Perceba que mudar engine para 24pfs ou 240fps não afeta o movimento de animações nativas e nem sempre afeta a
// qualidade das animações controladas pela engine, devido a uma limitação apresentada pelos navegadores (07/2022).
//
// O algorítimo considera o tempo de resposta para ajusta fps entre os valores máximos e mínimos visando um melhor
// processamento, porém, se a taxa de fps for muito elevada, o navegador trava (07/2022).
package globalEngine
