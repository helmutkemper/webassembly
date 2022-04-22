package html

type InputType string

func (e InputType) String() string {
	return string(e)
}

const (
	// KInputTypeButton
	//
	// English:
	//
	//  A push button with no default behavior displaying the value of the value attribute, empty by
	//  default.
	//
	// Português:
	//
	//  Um botão de ação sem comportamento padrão exibindo o valor do atributo value, vazio por padrão.
	KInputTypeButton InputType = "button"

	// KInputTypeCheckbox
	//
	// English:
	//
	//  A check box allowing single values to be selected/deselected.
	//
	// Português:
	//
	//  Uma caixa de seleção que permite que valores únicos sejam selecionados desmarcados.
	KInputTypeCheckbox InputType = "checkbox"

	// KInputTypeColor
	//
	// English:
	//
	//  A control for specifying a color; opening a color picker when active in supporting browsers.
	//
	// Português:
	//
	//  Um controle para especificar uma cor; abrindo um seletor de cores quando ativo em navegadores
	//  compatíveis.
	KInputTypeColor InputType = "color"

	// KInputTypeDate
	//
	// English:
	//
	//  A control for entering a date (year, month, and day, with no time). Opens a date picker or
	//  numeric wheels for year, month, day when active in supporting browsers.
	//
	// Português:
	//
	//  Um controle para inserir uma data (ano, mês e dia, sem hora). Abre um seletor de data ou rodas
	//  numéricas para ano, mês, dia quando ativo em navegadores compatíveis.
	KInputTypeDate InputType = "date"

	// KInputTypeDatetimeLocal
	//
	// English:
	//
	//  A control for entering a date and time, with no time zone. Opens a date picker or numeric
	//  wheels for date- and time-components when active in supporting browsers.
	//
	// Português:
	//
	//  Um controle para inserir uma data e hora, sem fuso horário. Abre um seletor de data ou rodas
	//  numéricas para componentes de data e hora quando ativo em navegadores compatíveis.
	KInputTypeDatetimeLocal InputType = "datetime-local"

	// KInputTypeEmail
	//
	// English:
	//
	//  A field for editing an email address. Looks like a text input, but has validation parameters and
	//  relevant keyboard in supporting browsers and devices with dynamic keyboards.
	//
	// Português:
	//
	//  Um campo para editar um endereço de e-mail. Parece uma entrada de texto, mas possui parâmetros
	//  de validação e teclado relevante no suporte a navegadores e dispositivos com teclados dinâmicos.
	KInputTypeEmail InputType = "email"

	// KInputTypeFile
	//
	// English:
	//
	//  A control that lets the user select a file. Use the accept attribute to define the types of
	//  files that the control can select.
	//
	// Português:
	//
	//  Um controle que permite ao usuário selecionar um arquivo. Use o atributo accept para definir os
	//  tipos de arquivos que o controle pode selecionar.
	KInputTypeFile InputType = "file"

	// KInputTypeHidden
	//
	// English:
	//
	//  A control that is not displayed but whose value is submitted to the server. There is an example
	//  in the next column, but it's hidden!
	//
	// Português:
	//
	//  Um controle que não é exibido, mas cujo valor é enviado ao servidor. Há um exemplo na próxima
	//  coluna, mas está oculto!
	KInputTypeHidden InputType = "hidden"

	// KInputTypeImage
	//
	// English:
	//
	//  A graphical submit button. Displays an image defined by the src attribute. The alt attribute
	//  displays if the image src is missing.
	//
	// Português:
	//
	//  Um botão de envio gráfico. Exibe uma imagem definida pelo atributo src. O atributo alt é
	//  exibido se o src da imagem estiver ausente.
	KInputTypeImage InputType = "image"

	// KInputTypeMonth
	//
	// English:
	//
	//  A control for entering a month and year, with no time zone.
	//
	// Português:
	//
	//  Um controle para inserir um mês e ano, sem fuso horário.
	KInputTypeMonth InputType = "month"

	// KInputTypeNumber
	//
	// English:
	//
	//  A control for entering a number. Displays a spinner and adds default validation when supported.
	//  Displays a numeric keypad in some devices with dynamic keypads.
	//
	// Português:
	//
	//  Um controle para inserir um número. Exibe um spinner e adiciona validação padrão quando
	//  suportado. Exibe um teclado numérico em alguns dispositivos com teclados dinâmicos.
	KInputTypeNumber InputType = "number"

	// KInputTypePassword
	//
	// English:
	//
	//  A single-line text field whose value is obscured. Will alert user if site is not secure.
	//
	// Português:
	//
	//  Um campo de texto de linha única cujo valor está obscurecido. Alertará o usuário se o site não
	//  for seguro.
	KInputTypePassword InputType = "password"

	// KInputTypeRadio
	//
	// English:
	//
	//  A radio button, allowing a single value to be selected out of multiple choices with the same
	//  name value.
	//
	// Português:
	//
	//  Um botão de opção, permitindo que um único valor seja selecionado entre várias opções com o
	//  mesmo valor de nome.
	KInputTypeRadio InputType = "radio"

	// KInputTypeRange
	//
	// English:
	//
	//  A control for entering a number whose exact value is not important. Displays as a range widget
	//  defaulting to the middle value. Used in conjunction min and max to define the range of
	//  acceptable values.
	//
	// Português:
	//
	//  Um controle para inserir um número cujo valor exato não é importante. Exibe como um widget de
	//  intervalo padronizado para o valor médio. Usado em conjunto min e max para definir a faixa de
	//  valores aceitáveis.
	KInputTypeRange InputType = "range"

	// KInputTypeSearch
	//
	// English:
	//
	//  A single-line text field for entering search strings. Line-breaks are automatically removed
	//  from the input value. May include a delete icon in supporting browsers that can be used to
	//  clear the field. Displays a search icon instead of enter key on some devices with dynamic
	//  keypads.
	//
	// Português:
	//
	//  Um campo de texto de linha única para inserir strings de pesquisa. As quebras de linha são
	//  removidas automaticamente do valor de entrada. Pode incluir um ícone de exclusão em navegadores
	//  de suporte que podem ser usados para limpar o campo. Exibe um ícone de pesquisa em vez da tecla
	//  Enter em alguns dispositivos com teclados dinâmicos.
	KInputTypeSearch InputType = "search"

	// KInputTypeSubmit
	//
	// English:
	//
	//  A button that submits the form.
	//
	// Português:
	//
	//  Um botão que envia o formulário.
	KInputTypeSubmit InputType = "submit"

	// KInputTypeTel
	//
	// English:
	//
	//  A control for entering a telephone number. Displays a telephone keypad in some devices with
	//  dynamic keypads.
	//
	// Português:
	//
	//  Um controle para inserir um número de telefone. Exibe um teclado de telefone em alguns
	//  dispositivos com teclados dinâmicos.
	KInputTypeTel InputType = "tel"

	// KInputTypeText
	//
	// English:
	//
	//  A single-line text field. Line-breaks are automatically removed from the input value. (Default)
	//
	// Português:
	//
	//  Mampo de texto de linha única. As quebras de linha são removidas automaticamente do valor de
	//  entrada. (Padrão)
	KInputTypeText InputType = "text"

	// KInputTypeTime
	//
	// English:
	//
	//  A control for entering a time value with no time zone.
	//
	// Português:
	//
	//  Um controle para inserir um valor de tempo sem fuso horário.
	KInputTypeTime InputType = "time"

	// KInputTypeUrl
	//
	// English:
	//
	//  A field for entering a URL. Looks like a text input, but has validation parameters and relevant
	//  keyboard in supporting browsers and devices with dynamic keyboards.
	//
	// Português:
	//
	//  Um campo para inserir um URL. Parece uma entrada de texto, mas possui parâmetros de validação
	//  e teclado relevante no suporte a navegadores e dispositivos com teclados dinâmicos.
	KInputTypeUrl InputType = "url"

	// KInputTypeWeek
	//
	// English:
	//
	//  A control for entering a date consisting of a week-year number and a week number with no time
	//  zone.
	//
	// Português:
	//
	//  Um controle para inserir uma data que consiste em um número de ano-semana e um número de semana
	//  sem fuso horário.
	KInputTypeWeek InputType = "week"
)
