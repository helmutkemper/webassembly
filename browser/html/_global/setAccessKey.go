package _global

// SetAccessKey
//
// English:
//
//  Specifies a shortcut key to activate/focus an element.
//
//   Input:
//     character: A single character that specifies the shortcut key to activate/focus the element.
//
//   Note:
//     * The accessKey attribute value must be a single character (a letter or a digit).
//     * Adapting accessKeys to all international languages are difficult.
//     * The accessKey value may not be present on all keyboards.
//
//
//   Warning:
//     Using accessKeys is difficult because they may conflict with other key standards in the
//     browser;
//     To avoid this problem, most browsers will use accessKeys only if pressed together with the Alt
//     key.
//
// Português:
//
//  Especifica uma tecla de atalho para ativar o foco de um elemento.
//
//   Entrada:
//     character: Um único caractere que especifica a tecla de atalho para ativar o foco do elemento.
//
//   Nota:
//     * O valor do atributo accessKey deve ser um único caractere (uma letra ou um dígito).
//     * Adaptar as teclas de acesso a todos os idiomas internacionais é difícil.
//     * O valor accessKey pode não estar presente em todos os teclados.
//
//   Aviso:
//     O uso de accessKeys é difícil porque eles podem entrar em conflito com outros padrões
//     importantes no navegador;
//     Para evitar esse problema, a maioria dos navegadores usará as teclas de acesso somente se
//     pressionadas junto com a tecla Alt.
func (e *GlobalAttributes) SetAccessKey(key string) (ref *GlobalAttributes) {
	e.selfElement.Set("accesskey", key)
	return e
}
