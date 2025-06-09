// Package block
//
// English:
//
// block is the main instance of the device and is responsible for all generic functions, such as selecting, resizing
// or drawing the device.
//
// I am aware of the fact that Golang does not support much injection of dependence and I understand the reasons, but
// this makes the code easier for a single person gives maintenance in the initial phase of the project.
//
//	Tips:
//	  * The IDE stage used for the following properties in CSS:
//	    position: relative; width: 100vw; height: 100vh
//	    This was defined by the SetFatherId() function.
//
// Português:
//
// block é à instância principal do dispositivo e é responsável por todas as funções genéricas, como selecionar,
// redimensionar ou desenhar o dispositivo.
//
// Eu estou ciente do fato de golang não apoia muito injeção de dependência e eu entendo os motivos, mas, isto deixa a
// criação do código mais fácil para uma única pessoa dá manutenção na fase inicial do projeto
//
//	Dicas:
//	  * A div usada para palco da IDE passa a ter no CSS as seguintes propriedades:
//	    position: relative; width: 100vw; height: 100vh
//	    isto foi definido pela função SetFatherId().
package block
