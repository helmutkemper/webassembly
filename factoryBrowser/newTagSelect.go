package factoryBrowser

// NewTagSelect
//
// English:
//
//  Create a new element select.
//
// The <select> HTML element represents a control that provides a menu of options.
//
// Português
//
//  Cria um novo elemento select.
//
// O elemento HTML <select> representa um controle que fornece um menu de opções.
//
// Example
//
//   factoryBrowser.NewTagSelect("test_A").
//     SetNewOption("test_A_a", "label a", "value_a", true, false).
//     SetNewOption("test_A_b", "label b", "value_b", false, false).
//     SetNewOption("test_A_c", "label c", "value_c", false, false).
//     SetNewOption("test_A_d", "label d", "value_d", false, true).
//     AppendById("stage")
