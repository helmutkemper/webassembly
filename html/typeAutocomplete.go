package html

type Autocomplete string

func (e Autocomplete) String() string {
	return string(e)
}

const (
	// KAutocompleteOff
	//
	// English:
	//
	//  The browser is not permitted to automatically enter or select a value for this field. It is
	//  possible that the document or application provides its own autocomplete feature, or that
	//  security concerns require that the field's value not be automatically entered.
	//
	//   Note:
	//     * In most modern browsers, setting autocomplete to "off" will not prevent a password manager
	//       from asking the user if they would like to save username and password information, or from
	//       automatically filling in those values in a site's login form. See the autocomplete
	//       attribute and login fields.
	//
	// Português:
	//
	//  O navegador não tem permissão para inserir ou selecionar automaticamente um valor para este
	//  campo. É possível que o documento ou aplicativo forneça seu próprio recurso de preenchimento
	//  automático ou que questões de segurança exijam que o valor do campo não seja inserido
	//  automaticamente.
	//
	//   Nota:
	//     * Na maioria dos navegadores modernos, definir o preenchimento automático como "desativado"
	//       não impedirá que um gerenciador de senhas pergunte ao usuário se ele deseja salvar as
	//       informações de nome de usuário e senha ou de preencher automaticamente esses valores no
	//       formulário de login de um site. Consulte o atributo de preenchimento automático e os
	//       campos de login.
	KAutocompleteOff Autocomplete = "off"

	// KAutocompleteOn
	//
	// English:
	//
	//  The browser is allowed to automatically complete the input. No guidance is provided as to the
	//  type of data expected in the field, so the browser may use its own judgement.
	//
	// Português:
	//
	//  O navegador tem permissão para completar automaticamente a entrada. Nenhuma orientação é
	//  fornecida quanto ao tipo de dados esperados no campo, portanto, o navegador pode usar seu
	//  próprio julgamento.
	KAutocompleteOn Autocomplete = "on"

	// KAutocompleteName
	//
	// English:
	//
	//  The field expects the value to be a person's full name. Using "name" rather than breaking the
	//  name down into its components is generally preferred because it avoids dealing with the wide
	//  diversity of human names and how they are structured; however, you can use the following
	//  autocomplete values if you do need to break the name down into its components:
	//
	// Português:
	//
	//  The field expects the value to be a person's full name. Using "name" rather than breaking the
	//  name down into its components is generally preferred because it avoids dealing with the wide
	//  diversity of human names and how they are structured; however, you can use the following
	//  autocomplete values if you do need to break the name down into its components:
	KAutocompleteName Autocomplete = "name"

	// KAutocompleteHonorificPrefix
	//
	// English:
	//
	//  The prefix or title, such as "Mrs.", "Mr.", "Miss", "Ms.", "Dr.", or "Mlle.".
	//
	// Português:
	//
	//  O prefixo ou título, como "Mrs.", "Mr.", "Miss", "Ms.", "Dr." ou "Mlle.".
	KAutocompleteHonorificPrefix Autocomplete = "honorific-prefix"

	// KAutocompleteGivenName
	//
	// English:
	//
	//  The given (or "first") name.
	//
	// Português:
	//
	//  O primeiro nome.
	KAutocompleteGivenName Autocomplete = "given-name"

	// KAutocompleteAdditionalName
	//
	// English:
	//
	//  The middle name.
	//
	// Português:
	//
	//  O nome do meio.
	KAutocompleteAdditionalName Autocomplete = "additional-name"

	// KAutocompleteFamilyName
	//
	// English:
	//
	//  The family (or "last") name.
	//
	// Português:
	//
	//  Sobrenome
	KAutocompleteFamilyName Autocomplete = "family-name"

	// KAutocompleteHonorificSuffix
	//
	// English:
	//
	//  The suffix, such as "Jr.", "B.Sc.", "PhD.", "MBASW", or "IV".
	//
	// Português:
	//
	//  O sufixo, como "Jr.", "B.Sc.", "PhD.", "MBASW" ou "IV".
	KAutocompleteHonorificSuffix Autocomplete = "honorific-suffix"

	// KAutocompleteNickname
	//
	// English:
	//
	//  A nickname or handle.
	//
	// Português:
	//
	//  Um apelido ou identificador.
	KAutocompleteNickname Autocomplete = "nickname"

	// KAutocompleteEmail
	//
	// English:
	//
	//  An email address.
	//
	// Português:
	//
	//  Um endereço de e-mail.
	KAutocompleteEmail Autocomplete = "email"

	// KAutocompleteUsername
	//
	// English:
	//
	//  A username or account name.
	//
	// Português:
	//
	//  Um nome de usuário.
	KAutocompleteUsername Autocomplete = "username"

	// KAutocompleteNewPassword
	//
	// English:
	//
	//  A new password. When creating a new account or changing passwords, this should be used for an
	//  "Enter your new password" or "Confirm new password" field, as opposed to a general "Enter your
	//  current password" field that might be present. This may be used by the browser both to avoid
	//  accidentally filling in an existing password and to offer assistance in creating a secure
	//  password (see also Preventing autofilling with autocomplete="new-password").
	//
	// Português:
	//
	//  Uma nova senha. Ao criar uma nova conta ou alterar senhas, isso deve ser usado para um campo
	//  "Digite sua nova senha" ou "Confirme nova senha", em vez de um campo geral "Digite sua senha
	//  atual" que pode estar presente. Isso pode ser usado pelo navegador tanto para evitar o
	//  preenchimento acidental de uma senha existente quanto para oferecer assistência na criação de
	//  uma senha segura (consulte também Impedindo o preenchimento automático com
	//  autocomplete="new-password").
	KAutocompleteNewPassword Autocomplete = "new-password"

	// KAutocompleteCurrentPassword
	//
	// English:
	//
	//  The user's current password.
	//
	// Português:
	//
	//  A senha atual do usuário.
	KAutocompleteCurrentPassword Autocomplete = "current-password"

	// KAutocompleteOneTimeCode
	//
	// English:
	//
	//  A one-time code used for verifying user identity.
	//
	// Português:
	//
	//  Um código único usado para verificar a identidade do usuário.
	KAutocompleteOneTimeCode Autocomplete = "one-time-code"

	// KAutocompleteOrganizationTitle
	//
	// English:
	//
	//  A job title, or the title a person has within an organization, such as "Senior Technical
	//  Writer", "President", or "Assistant Troop Leader".
	//
	// Português:
	//
	//  Um cargo ou o título que uma pessoa tem dentro de uma organização, como "Escritor Técnico
	//  Sênior", "Presidente" ou "Líder de Tropa Assistente".
	KAutocompleteOrganizationTitle Autocomplete = "organization-title"

	// KAutocompleteOrganization
	//
	// English:
	//
	//  A company or organization name, such as "Acme Widget Company" or "Girl Scouts of America".
	//
	// Português:
	//
	//  Um nome de empresa ou organização, como "Acme Widget Company" ou "Girl Scouts of America".
	KAutocompleteOrganization Autocomplete = "organization"

	// KAutocompleteStreetAddress
	//
	// English:
	//
	//  A street address. This can be multiple lines of text, and should fully identify the location of
	//  the address within its second administrative level (typically a city or town), but should not
	//  include the city name, ZIP or postal code, or country name.
	//
	// Português:
	//
	//  Um endereço de rua. Isso pode ser várias linhas de texto e deve identificar totalmente o local
	//  do endereço em seu segundo nível administrativo (normalmente uma cidade ou vila), mas não deve
	//  incluir o nome da cidade, CEP ou código postal ou nome do país.
	KAutocompleteStreetAddress Autocomplete = "street-address"

	// KAutocompleteAddressLine1
	//
	// English:
	//
	//  Each individual line of the street address. These should only be present if the
	//  "street-address" is not present.
	//
	// Português:
	//
	//  Cada linha individual do endereço. Estes só devem estar presentes se o "endereço" não estiver
	//  presente.
	KAutocompleteAddressLine1 Autocomplete = "address-line1"

	// KAutocompleteAddressLine2
	//
	// English:
	//
	//  Each individual line of the street address. These should only be present if the
	//  "street-address" is not present.
	//
	// Português:
	//
	//  Cada linha individual do endereço. Estes só devem estar presentes se o "endereço" não estiver
	//  presente.
	KAutocompleteAddressLine2 Autocomplete = "address-line2"

	// KAutocompleteAddressLine3
	//
	// English:
	//
	//  Each individual line of the street address. These should only be present if the
	//  "street-address" is not present.
	//
	// Português:
	//
	//  Cada linha individual do endereço. Estes só devem estar presentes se o "endereço" não estiver
	//  presente.
	KAutocompleteAddressLine3 Autocomplete = "address-line3"

	// KAutocompleteAddressLevel4
	//
	// English:
	//
	//  The finest-grained administrative level, in addresses which have four levels.
	//
	// Português:
	//
	//  The finest-grained administrative level, in addresses which have four levels.
	KAutocompleteAddressLevel4 Autocomplete = "address-level4"

	// KAutocompleteAddressLevel3
	//
	// English:
	//
	//  The finest-grained administrative level, in addresses which have four levels.
	//
	// Português:
	//
	//  The finest-grained administrative level, in addresses which have four levels.
	KAutocompleteAddressLevel3 Autocomplete = "address-level3"

	// KAutocompleteAddressLevel2
	//
	// English:
	//
	//  The finest-grained administrative level, in addresses which have four levels.
	//
	// Português:
	//
	//  The finest-grained administrative level, in addresses which have four levels.
	KAutocompleteAddressLevel2 Autocomplete = "address-level2"

	// KAutocompleteAddressLevel1
	//
	// English:
	//
	//  The finest-grained administrative level, in addresses which have four levels.
	//
	// Português:
	//
	//  The finest-grained administrative level, in addresses which have four levels.
	KAutocompleteAddressLevel1 Autocomplete = "address-level1"

	// KAutocompleteCountry
	//
	// English:
	//
	//  A country or territory code.
	//
	// Português:
	//
	//  Um código de país ou território.
	KAutocompleteCountry Autocomplete = "country"

	// KAutocompleteCountryName
	//
	// English:
	//
	//  A country or territory name.
	//
	// Português:
	//
	//  Um nome de país ou território.
	KAutocompleteCountryName Autocomplete = "country-name"

	// KAutocompletePostalCode
	//
	// English:
	//
	//  A postal code (in the United States, this is the ZIP code).
	//
	// Português:
	//
	//  Um código postal (nos Estados Unidos, este é o CEP).
	KAutocompletePostalCode Autocomplete = "postal-code"

	// KAutocompleteCcName
	//
	// English:
	//
	//  The full name as printed on or associated with a payment instrument such as a credit card.
	//  Using a full name field is preferred, typically, over breaking the name into pieces.
	//
	// Português:
	//
	//  O nome completo impresso ou associado a um instrumento de pagamento, como um cartão de crédito.
	//  Normalmente, é preferível usar um campo de nome completo em vez de dividir o nome em partes.
	KAutocompleteCcName Autocomplete = "cc-name"

	// KAutocompleteCcGivenName
	//
	// English:
	//
	//  A given (first) name as given on a payment instrument like a credit card.
	//
	// Português:
	//
	//  Primeiro nome dado em um instrumento de pagamento, como um cartão de crédito.
	KAutocompleteCcGivenName Autocomplete = "cc-given-name"

	// KAutocompleteCcAdditionalName
	//
	// English:
	//
	//  A middle name as given on a payment instrument or credit card.
	//
	// Português:
	//
	//  Nome do meio fornecido em um instrumento de pagamento ou cartão de crédito.
	KAutocompleteCcAdditionalName Autocomplete = "cc-additional-name"

	// KAutocompleteCcFamilyName
	//
	// English:
	//
	//  A family name, as given on a credit card.
	//
	// Português:
	//
	//  Nome de família, conforme fornecido em um cartão de crédito.
	KAutocompleteCcFamilyName Autocomplete = "cc-family-name"

	// KAutocompleteCcNumber
	//
	// English:
	//
	//  A credit card number or other number identifying a payment method, such as an account number.
	//
	// Português:
	//
	//  Um número de cartão de crédito ou outro número que identifique um método de pagamento, como um
	//  número de conta.
	KAutocompleteCcNumber Autocomplete = "cc-number"

	// KAutocompleteCcExp
	//
	// English:
	//
	//  A payment method expiration date, typically in the form "MM/YY" or "MM/YYYY".
	//
	// Português:
	//
	//  Uma data de expiração do método de pagamento, normalmente no formato "MM/YY" ou "MM/YYYY".
	KAutocompleteCcExp Autocomplete = "cc-exp"

	// KAutocompleteCcExpMonth
	//
	// English:
	//
	//  The month in which the payment method expires.
	//
	// Português:
	//
	//  O mês em que a forma de pagamento expira.
	KAutocompleteCcExpMonth Autocomplete = "cc-exp-month"

	// KAutocompleteCcExpYear
	//
	// English:
	//
	//  The year in which the payment method expires.
	//
	// Português:
	//
	//  O ano em que a forma de pagamento expira.
	KAutocompleteCcExpYear Autocomplete = "cc-exp-year"

	// KAutocompleteCcCsc
	//
	// English:
	//
	//  The security code for the payment instrument; on credit cards, this is the 3-digit verification
	//  number on the back of the card.
	//
	// Português:
	//
	//  O código de segurança do instrumento de pagamento; em cartões de crédito, este é o número de
	//  verificação de 3 dígitos no verso do cartão.
	KAutocompleteCcCsc Autocomplete = "cc-csc"

	// KAutocompleteCcType
	//
	// English:
	//
	// The type of payment instrument (such as "Visa" or "Master Card").
	//
	// Português:
	//
	// O tipo de instrumento de pagamento (como "Visa" ou "Master Card").
	KAutocompleteCcType Autocomplete = "cc-type"

	// KAutocompleteTransactionCurrency
	//
	// English:
	//
	//  The currency in which the transaction is to take place.
	//
	// Português:
	//
	//  A moeda em que a transação deve ocorrer.
	KAutocompleteTransactionCurrency Autocomplete = "transaction-currency"

	// KAutocompleteTransactionAmount
	//
	// English:
	//
	//  The amount, given in the currency specified by "transaction-currency", of the transaction,
	//  for a payment form.
	//
	// Português:
	//
	//  O valor, fornecido na moeda especificada por "moeda da transação", da transação, para um
	//  formulário de pagamento.
	KAutocompleteTransactionAmount Autocomplete = "transaction-amount"

	// KAutocompleteLanguage
	//
	// English:
	//
	//  A preferred language, given as a valid BCP 47 language tag.
	//
	// Português:
	//
	//  Um idioma preferencial, fornecido como uma tag de idioma BCP 47 válida.
	KAutocompleteLanguage Autocomplete = "language"

	// KAutocompleteBbday
	//
	// English:
	//
	//  A birth date, as a full date.
	//
	// Português:
	//
	//  Uma data de nascimento, como uma data completa.
	KAutocompleteBbday Autocomplete = "bday"

	// KAutocompleteBdayDay
	//
	// English:
	//
	//  The day of the month of a birth date.
	//
	// Português:
	//
	//  O dia do mês de uma data de nascimento.
	KAutocompleteBdayDay Autocomplete = "bday-day"

	// KAutocompleteBdayMonth
	//
	// English:
	//
	//  The month of the year of a birth date.
	//
	// Português:
	//
	//  O mês do ano de uma data de nascimento.
	KAutocompleteBdayMonth Autocomplete = "bday-month"

	// KAutocompleteBdayYear
	//
	// English:
	//
	//  The year of a birth date.
	//
	// Português:
	//
	//  O ano de uma data de nascimento.
	KAutocompleteBdayYear Autocomplete = "bday-year"

	// KAutocompleteSex
	//
	// English:
	//
	//  A gender identity (such as "Female", "Fa'afafine", "Male"), as freeform text without newlines.
	//
	// Português:
	//
	//  Uma identidade de gênero (como "Feminino", "Fa'afafine", "Masculino"), como texto de forma
	//  livre sem novas linhas.
	KAutocompleteSex Autocomplete = "sex"

	// KAutocompleteTel
	//
	// English:
	//
	//  A full telephone number, including the country code. If you need to break the phone number
	//  up into its components, you can use these values for those fields:
	//
	// Português:
	//
	//  Um número de telefone completo, incluindo o código do país. Se você precisar dividir o número
	//  de telefone em seus componentes, poderá usar estes valores para esses campos:
	KAutocompleteTel Autocomplete = "tel"

	// KAutocompleteTelCountryCode
	//
	// English:
	//
	//  The country code, such as "1" for the United States, Canada, and other areas in North America
	//  and parts of the Caribbean.
	//
	// Português:
	//
	//  The country code, such as "1" for the United States, Canada, and other areas in North America
	//  and parts of the Caribbean.
	KAutocompleteTelCountryCode Autocomplete = "tel-country-code"

	// KAutocompleteTelNational
	//
	// English:
	//
	//  The entire phone number without the country code component, including a country-internal
	//  prefix. For the phone number "1-855-555-6502", this field's value would be "855-555-6502".
	//
	// Português:
	//
	//  O número de telefone completo sem o componente de código do país, incluindo um prefixo interno
	//  do país. Para o número de telefone "1-855-555-6502", o valor desse campo seria "855-555-6502".
	KAutocompleteTelNational Autocomplete = "tel-national"

	// KAutocompleteTelAreaCode
	//
	// English:
	//
	//  The area code, with any country-internal prefix applied if appropriate.
	//
	// Português:
	//
	//  O código de área, com qualquer prefixo interno do país aplicado, se apropriado.
	KAutocompleteTelAreaCode Autocomplete = "tel-area-code"

	// KAutocompleteTelLocal
	//
	// English:
	//
	//  The phone number without the country or area code. This can be split further into two parts,
	//  for phone numbers which have an exchange number and then a number within the exchange. For the
	//  phone number "555-6502", use "tel-local-prefix" for "555" and "tel-local-suffix" for "6502".
	//
	// Português:
	//
	//  O número de telefone sem o código do país ou área. Isso pode ser dividido em duas partes, para
	//  números de telefone que têm um número de troca e, em seguida, um número dentro da troca. Para
	//  o número de telefone "555-6502", use "tel-local-prefix" para "555" e "tel-local-suffix"
	//  para "6502".
	KAutocompleteTelLocal Autocomplete = "tel-local"

	// KAutocompleteTelExtension
	//
	// English:
	//
	//  A telephone extension code within the phone number, such as a room or suite number in a hotel
	//  or an office extension in a company.
	//
	// Português:
	//
	//  Um código de ramal de telefone dentro do número de telefone, como um número de quarto ou suíte
	//  em um hotel ou um ramal de escritório em uma empresa.
	KAutocompleteTelExtension Autocomplete = "tel-extension"

	// KAutocompleteImpp
	//
	// English:
	//
	//  A URL for an instant messaging protocol endpoint, such as "xmpp:username@example.net".
	//
	// Português:
	//
	//  Um URL para um terminal de protocolo de mensagens instantâneas, como
	//  "xmpp:username@example.net".
	KAutocompleteImpp Autocomplete = "impp"

	// KAutocompleteUrl
	//
	// English:
	//
	//  A URL, such as a home page or company web site address as appropriate given the context of
	//  the other fields in the form.
	//
	// Português:
	//
	//  A URL, such as a home page or company web site address as appropriate given the context of
	//  the other fields in the form.
	KAutocompleteUrl Autocomplete = "url"

	// KAutocompletePhoto
	//
	// English:
	//
	//  The URL of an image representing the person, company, or contact information given in the
	//  other fields in the form.
	//
	// Português:
	//
	//  O URL de uma imagem que representa a pessoa, empresa ou informações de contato fornecidas
	//  nos outros campos do formulário.
	KAutocompletePhoto Autocomplete = "photo"
)
