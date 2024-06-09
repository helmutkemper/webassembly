package generic

import (
	"syscall/js"
)

type Data struct {
	// This
	//
	// English:
	//
	// This is the equivalent property of JavaScript's 'this'.
	//
	// The way to use it is This.Get(property string name). E.g. chan.This.Get("id")
	//
	// Português:
	//
	// Esta é a propriedade equivalente ao 'this' do JavaScript.
	//
	// A forma de usar é This.Get(property string name). Ex. chan.This.Get("id")
	This js.Value

	// AltKey
	//
	// English:
	//
	// Returns a boolean value that is true if the Alt (Option or ⌥ on macOS) key was active when the key event was
	// generated.
	//
	// Português:
	//
	// Retorna um valor booleano que é verdadeiro se a tecla Alt (Option ou ⌥ no macOS) estiver ativa quando o evento de
	// chave foi gerado.
	AltKey bool

	// Code
	//
	// English:
	//
	// Returns a string with the code value of the physical key represented by the event.
	//
	//  Warning:
	//    This ignores the user's keyboard layout, so that if the user presses the key at the "Y" position in a QWERTY
	//    keyboard layout (near the middle of the row above the home row), this will always return "KeyY", even if the
	//    user has a QWERTZ keyboard (which would mean the user expects a "Z" and all the other properties would indicate
	//    a "Z") or a Dvorak keyboard layout (where the user would expect an "F"). If you want to display the correct
	//    keystrokes to the user, you can use Keyboard.getLayoutMap().
	//
	// Português:
	//
	// Retorna uma string com o valor do código da chave física representada pelo evento.
	//
	//  Cuidado:
	//    Ele ignora o layout do teclado do usuário, de modo que, se o usuário pressionar a tecla na posição "Y" em um
	//    layout de teclado QWERTY (perto do meio da linha acima da linha inicial), isso sempre retornará "KeyY", mesmo
	//    se o usuário tiver um teclado QWERTZ (o que significa que o usuário espera um "Z" e todas as outras propriedades
	//    indicariam um "Z") ou um layout de teclado Dvorak (onde o usuário espera um "F"). Se você deseja exibir as
	//    teclas corretas para o usuário, você pode usar Keyboard.getLayoutMap().
	Code string

	// CtrlKey
	//
	// English:
	//
	// Returns a boolean value that is true if the Ctrl key was active when the key event was generated.
	//
	// Português:
	//
	// Retorna um valor booleano que é true se a tecla Ctrl estava ativa quando o evento de chave foi gerado.
	CtrlKey bool

	// IsComposing
	//
	// English:
	//
	// Returns a boolean value that is true if the event is fired between after compositionstart and before
	// compositionend.
	//
	// Português:
	//
	// Retorna um valor booleano que é verdadeiro se o evento for acionado entre após o início da composição e antes da
	// conclusão da composição.
	IsComposing bool

	// Key
	//
	// English:
	//
	// Returns a string representing the key value of the key represented by the event.
	//
	// Português:
	//
	// Retorna uma string representando o valor da chave representada pelo evento.
	Key string

	// Location
	//
	// English:
	//
	// Returns a number representing the location of the key on the keyboard or other input device.
	//
	// Português:
	//
	// Retorna um número que representa a localização da tecla no teclado ou outro dispositivo de entrada.
	Location Location

	// MetaKey
	//
	// English:
	//
	// Returns a boolean value that is true if the Meta key (on Mac keyboards, the ⌘ Command key; on Windows keyboards,
	// the Windows key (⊞)) was active when the key event was generated.
	//
	// Português:
	//
	// Returns a boolean value that is true if the Meta key (on Mac keyboards, the ⌘ Command key; on Windows keyboards,
	// the Windows key (⊞)) was active when the key event was generated.
	MetaKey bool

	// Repeat
	//
	// English:
	//
	// Returns a boolean value that is true if the key is being held down such that it is automatically repeating.
	//
	// Português:
	//
	// Retorna um valor booleano que é true se a chave estiver sendo mantida pressionada de forma que ela se repita
	// automaticamente.
	Repeat bool

	// ShiftKey
	//
	// English:
	//
	// Returns a boolean value that is true if the Shift key was active when the key event was generated.
	//
	// Português:
	//
	// Retorna um valor booleano que é true se a tecla Shift estava ativa quando o evento de chave foi gerado.
	ShiftKey bool

	Object js.Value

	// EventName
	//
	// English:
	//
	// Name of event
	//
	// Português:
	//
	// Nome do evento
	EventName EventName

	Accesskey       string
	Autocapitalize  string
	Autofocus       string
	Class           string
	Contenteditable string
	Dir             string
	Draggable       int
	Enterkeyhint    int
	Exportparts     string
	Hidden          string
	Id              string
	Inert           bool
	InputMode       string
	Is              string
	ItemId          string
	Itemprop        string
	ItemRef         string
	ItemScope       string
	ItemType        string
	Lang            string
	Nonce           string
	Popover         string
	Role            string
	Spellcheck      bool
	Style           string
	TabIndex        int
	Title           string
	Translate       string
}

type Event struct {
	Object js.Value
}

// GetAccesskey
//
// English:
//
// The accesskey global attribute provides a hint for generating a keyboard shortcut for the current element.
// The attribute value must consist of a single printable character (which includes accented and other characters that can be generated by the keyboard).
func (e Event) GetAccesskey() (accesskey string) {
	obj := e.Object.Get("accesskey")

	if obj.IsNull() || obj.IsUndefined() || obj.IsNaN() {
		return ""
	}

	return obj.String()
}

// GetAutocapitalize
//
// English:
//
// The autocapitalize global attribute is an enumerated attribute that controls whether inputted text is automatically capitalized and, if so, in what manner. This is relevant to:
//
// <input> and <textarea> elements.
// Any element with contenteditable set on it.
// autocapitalize doesn't affect behavior when typing on a physical keyboard. It affects the behavior of other input mechanisms such as virtual keyboards on mobile devices and voice input. This can assist users by making data entry quicker and easier, for example by automatically capitalizing the first letter of each sentence.
//
// Value
// Possible values are:
//
// none or off
// Do not automatically capitalize any text.
//
// sentences or on
// Automatically capitalize the first character of each sentence.
//
// words
// Automatically capitalize the first character of each word.
//
// characters
// Automatically capitalize every character.
//
// Usage notes
// autocapitalize can be set on <input> and <textarea> elements, and on their containing <form> elements. When autocapitalize is set on a <form> element, it sets the autocapitalize behavior for all contained <input>s and <textarea>s, overriding any autocapitalize values set on contained elements.
// autocapitalize has no effect on the url, email, or password <input> types, where autocapitalization is never enabled.
// Where autocapitalize is not specified, the adopted default behavior varies between browsers. For example:
// Chrome and Safari default to on/sentences
// Firefox defaults to off/none.
func (e Event) GetAutocapitalize() (autocapitalize string) {
	obj := e.Object.Get("autocapitalize")

	if obj.IsNull() || obj.IsUndefined() || obj.IsNaN() {
		return ""
	}

	return obj.String()
}

// GetAutofocus
//
// English:
//
// The autofocus global attribute is a Boolean attribute indicating that an element should be focused on page load, or when the <dialog> that it is part of is displayed.
//
// HTML
// Copy to Clipboard
// <input name="q" autofocus />
// No more than one element in the document or dialog may have the autofocus attribute. If applied to multiple elements the first one will receive focus.
//
// Note: The autofocus attribute applies to all elements, not just form controls. For example, it might be used on a contenteditable area.
//
// Accessibility concerns
// Automatically focusing a form control can confuse visually-impaired people using screen-reading technology and people with cognitive impairments. When autofocus is assigned, screen-readers "teleport" their user to the form control without warning them beforehand.
//
// Use careful consideration for accessibility when applying the autofocus attribute. Automatically focusing on a control can cause the page to scroll on load. The focus can also cause dynamic keyboards to display on some touch devices. While a screen reader will announce the label of the form control receiving focus, the screen reader will not announce anything before the label, and the sighted user on a small device will equally miss the context created by the preceding content.
func (e Event) GetAutofocus() (autofocus string) {
	obj := e.Object.Get("autofocus")

	if obj.IsNull() || obj.IsUndefined() || obj.IsNaN() {
		return ""
	}

	return obj.String()
}

// GetClass
//
// English:
//
// The class global attribute is a space-separated list of the case-sensitive classes of the element. Classes allow CSS and JavaScript to select and access specific elements via the class selectors or functions like the DOM method document.getElementsByClassName.
func (e Event) GetClass() (class string) {
	obj := e.Object.Get("class")

	if obj.IsNull() || obj.IsUndefined() || obj.IsNaN() {
		return ""
	}

	return obj.String()
}

// GetContenteditable
//
// English:
//
// The contenteditable global attribute is an enumerated attribute indicating if the element should be editable by the user. If so, the browser modifies its widget to allow editing.
// Value
// The attribute must take one of the following values:
//
// true or an empty string, which indicates that the element is editable.
// false, which indicates that the element is not editable.
// plaintext-only, which indicates that the element's raw text is editable, but rich text formatting is disabled.
// If the attribute is given without a value, like <label contenteditable>Example Label</label>, its value is treated as an empty string.
//
// If this attribute is missing or its value is invalid, its value is inherited from its parent element: so the element is editable if its parent is editable.
//
// Note that although its allowed values include true and false, this attribute is an enumerated one and not a Boolean one.
//
// You can set the color used to draw the text insertion caret with the CSS caret-color property.
//
// Elements that are made editable, and therefore interactive, by using the contenteditable attribute can be focused. They participate in sequential keyboard navigation. However, elements with the contenteditable attribute nested within other contenteditable elements are not added to the tabbing sequence by default. You can add the nested contenteditable elements to the keyboard navigation sequence by specifying the tabindex value (tabindex="0").
func (e Event) GetContenteditable() (contenteditable string) {
	obj := e.Object.Get("contenteditable")
	if obj.IsNull() || obj.IsUndefined() || obj.IsNaN() {
		return ""
	}

	return obj.String()
}

// GetDir
//
// English:
//
// The dir global attribute is an enumerated attribute that indicates the directionality of the element's text.
// It can have the following values:
//
// ltr, which means left to right and is to be used for languages that are written from the left to the right (like English);
// rtl, which means right to left and is to be used for languages that are written from the right to the left (like Arabic);
// auto, which lets the user agent decide. It uses a basic algorithm as it parses the characters inside the element until it finds a character with a strong directionality, then applies that directionality to the whole element.
// Note: This attribute is mandatory for the <bdo> element where it has a different semantic meaning.
//
// This attribute is not inherited by the <bdi> element. If not set, its value is auto.
// This attribute can be overridden by the CSS properties direction and unicode-bidi, if a CSS page is active and the element supports these properties.
// As the directionality of the text is semantically related to its content and not to its presentation, it is recommended that web developers use this attribute instead of the related CSS properties when possible. That way, the text will display correctly even on a browser that doesn't support CSS or has the CSS deactivated.
// The auto value should be used for data with an unknown directionality, like data coming from user input, eventually stored in a database.
// Note: Browsers might allow users to change the directionality of <input> and <textarea>s in order to assist with authoring content. Chrome and Safari provide a directionality option in the contextual menu of input fields while Legacy Edge uses the key combinations Ctrl + Left Shift and Ctrl + Right Shift. Firefox uses Ctrl/Cmd + Shift + X but does NOT update the dir attribute value.
func (e Event) GetDir() (dir string) {
	obj := e.Object.Get("dir")

	if obj.IsNull() || obj.IsUndefined() || obj.IsNaN() {
		return ""
	}

	return obj.String()
}

// GetDraggable
//
// English:
//
// The draggable global attribute is an enumerated attribute that indicates whether the element can be dragged, either with native browser behavior or the HTML Drag and Drop API.
//
// The draggable attribute may be applied to elements that strictly fall under the HTML namespace, which means that it cannot be applied to SVGs. For more information about what namespace declarations look like, and what they do, see Namespace crash course.
//
// draggable can have the following values:
//
// true: the element can be dragged.
// false: the element cannot be dragged.
// Warning: This attribute is enumerated and not Boolean. A value of true or false is mandatory, and shorthand like <img draggable> is forbidden. The correct usage is <img draggable="false">.
//
// If this attribute is not set, its default value is auto, which means drag behavior is the default browser behavior: only text selections, images, and links can be dragged. For other elements, the event ondragstart must be set for drag and drop to work, as shown in this comprehensive example.
func (e Event) GetDraggable() (draggable int) {
	obj := e.Object.Get("draggable")

	if obj.IsNull() || obj.IsUndefined() || obj.IsNaN() {
		return 0
	}

	return obj.Int()
}

// GetEnterkeyhint
//
// English:
//
// The enterkeyhint global attribute is an enumerated attribute defining what action label (or icon) to present for the enter key on virtual keyboards.
//
// enterkeyhint
// The enterkeyhint global attribute is an enumerated attribute defining what action label (or icon) to present for the enter key on virtual keyboards.
//
// # Try it
//
// Description
// Form controls (such as <textarea> or <input> elements) or elements using contenteditable can specify an inputmode attribute to control what kind of virtual keyboard will be used. To further improve the user's experience, the enter key can be customized specifically by providing an enterkeyhint attribute indicating how the enter key should be labeled (or which icon should be shown). The enter key usually represents what the user should do next; typical actions are: sending text, inserting a new line, or searching.
//
// If no enterkeyhint attribute is provided, the user agent might use contextual information from the inputmode, type, or pattern attributes to display a suitable enter key label (or icon).
//
// Values
// The enterkeyhint attribute is an enumerated attribute and only accepts the following values:
//
// Value                   Description                                                                                         Example label (depends on user agent and user language)
// enterkeyhint="enter"    Typically inserting a new line.                                                                     ↵
// enterkeyhint="done"     Typically meaning there is nothing more to input and the input method editor (IME) will be closed.  Done
// enterkeyhint="go"       Typically meaning to take the user to the target of the text they typed.                            Open
// enterkeyhint="next"     Typically taking the user to the next field that will accept text.                                  Next
// enterkeyhint="previous" Typically taking the user to the previous field that will accept text.                              Previous
// enterkeyhint="search"   Typically taking the user to the results of searching for the text they have typed.                 Search
// enterkeyhint="send"     Typically delivering the text to its target.                                                        Send
func (e Event) GetEnterkeyhint() (enterkeyhint int) {
	obj := e.Object.Get("enterkeyhint")

	if obj.IsNull() || obj.IsUndefined() || obj.IsNaN() {
		return 0
	}

	return obj.Int()
}

// GetExportparts
//
// English:
//
// Used to transitively export shadow parts from a nested shadow tree into a containing light tree.
func (e Event) GetExportparts() (exportparts string) {
	obj := e.Object.Get("exportparts")

	if obj.IsNull() || obj.IsUndefined() || obj.IsNaN() {
		return ""
	}

	return obj.String()
}

// GetHidden
//
// English:
//
// An enumerated attribute indicating that the element is not yet, or is no longer, relevant. For example, it can be used to hide elements of the page that can't be used until the login process has been completed. The browser won't render such elements. This attribute must not be used to hide content that could legitimately be shown.
func (e Event) GetHidden() (hidden string) {
	obj := e.Object.Get("hidden")

	if obj.IsNull() || obj.IsUndefined() || obj.IsNaN() {
		return ""
	}

	return obj.String()
}

// GetId
//
// English:
//
// Defines a unique identifier (ID) which must be unique in the whole document. Its purpose is to identify the element when linking (using a fragment identifier), scripting, or styling (with CSS).
func (e Event) GetId() (id string) {
	obj := e.Object.Get("id")

	if obj.IsNull() || obj.IsUndefined() || obj.IsNaN() {
		return ""
	}

	return obj.String()
}

// GetInert
//
// English:
//
// A boolean value that makes the browser disregard user input events for the element. Useful when click events are present.
func (e Event) GetInert() (inert bool) {
	obj := e.Object.Get("inert")

	if obj.IsNull() || obj.IsUndefined() || obj.IsNaN() {
		return false
	}

	return obj.Bool()
}

// GetInputMode
//
// English:
//
// Provides a hint to browsers about the type of virtual keyboard configuration to use when editing this element or its contents. Used primarily on <input> elements, but is usable on any element while in contenteditable mode.
func (e Event) GetInputMode() (inputMode string) {
	obj := e.Object.Get("inputmode")

	if obj.IsNull() || obj.IsUndefined() || obj.IsNaN() {
		return ""
	}

	return obj.String()
}

// GetIs
//
// English:
//
// Allows you to specify that a standard HTML element should behave like a registered custom built-in element (see Using custom elements for more details).
func (e Event) GetIs() (is string) {
	obj := e.Object.Get("is")

	if obj.IsNull() || obj.IsUndefined() || obj.IsNaN() {
		return ""
	}

	return obj.String()
}

// GetItemId
//
// English:
//
// The unique, global identifier of an item.
func (e Event) GetItemId() (itemId string) {
	obj := e.Object.Get("itemid")

	if obj.IsNull() || obj.IsUndefined() || obj.IsNaN() {
		return ""
	}

	return obj.String()
}

// GetItemprop
//
// English:
//
// Used to add properties to an item. Every HTML element may have an itemprop attribute specified, where an itemprop consists of a name and value pair.
func (e Event) GetItemprop() (itemprop string) {
	obj := e.Object.Get("itemprop")

	if obj.IsNull() || obj.IsUndefined() || obj.IsNaN() {
		return ""
	}

	return obj.String()
}

// GetItemRef
//
// English:
//
// Properties that are not descendants of an element with the itemscope attribute can be associated with the item using an itemref. It provides a list of element ids (not itemids) with additional properties elsewhere in the document.
func (e Event) GetItemRef() (itemRef string) {
	obj := e.Object.Get("itemref")

	if obj.IsNull() || obj.IsUndefined() || obj.IsNaN() {
		return ""
	}

	return obj.String()
}

// GetItemScope
//
// English:
//
// itemscope (usually) works along with itemtype to specify that the HTML contained in a block is about a particular item. itemscope creates the Item and defines the scope of the itemtype associated with it. itemtype is a valid URL of a vocabulary (such as schema.org) that describes the item and its properties context.
func (e Event) GetItemScope() (itemScope string) {
	obj := e.Object.Get("itemscope")

	if obj.IsNull() || obj.IsUndefined() || obj.IsNaN() {
		return ""
	}

	return obj.String()
}

// GetItemType
//
// English:
//
// Specifies the URL of the vocabulary that will be used to define itemprops (item properties) in the data structure. itemscope is used to set the scope of where in the data structure the vocabulary set by itemtype will be active.
func (e Event) GetItemType() (itemType string) {
	obj := e.Object.Get("itemtype")

	if obj.IsNull() || obj.IsUndefined() || obj.IsNaN() {
		return ""
	}

	return obj.String()
}

// GetLang
//
// English:
//
// Helps define the language of an element: the language that non-editable elements are in, or the language that editable elements should be written in by the user. The attribute contains one "language tag" (made of hyphen-separated "language subtags") in the format defined in RFC 5646: Tags for Identifying Languages (also known as BCP 47). xml:lang has priority over it.
func (e Event) GetLang() (lang string) {
	obj := e.Object.Get("lang")

	if obj.IsNull() || obj.IsUndefined() || obj.IsNaN() {
		return ""
	}

	return obj.String()
}

// GetNonce
//
// English:
//
// A cryptographic nonce ("number used once") which can be used by Content Security Policy to determine whether or not a given fetch will be allowed to proceed.
func (e Event) GetNonce() (nonce string) {
	obj := e.Object.Get("nonce")

	if obj.IsNull() || obj.IsUndefined() || obj.IsNaN() {
		return ""
	}

	return obj.String()
}

// GetPart
//
// English:
//
// A space-separated list of the part names of the element. Part names allows CSS to select and style specific elements in a shadow tree via the ::part pseudo-element.
func (e Event) GetPart() (part string) {
	obj := e.Object.Get("part")

	if obj.IsNull() || obj.IsUndefined() || obj.IsNaN() {
		return ""
	}

	return obj.String()
}

// GetPopover
//
// English:
//
// Used to designate an element as a popover element (see Popover API). Popover elements are hidden via display: none until opened via an invoking/control element (i.e. a <button> or <input type="button"> with a popovertarget attribute) or a HTMLElement.showPopover() call.
func (e Event) GetPopover() (popover string) {
	obj := e.Object.Get("popover")

	if obj.IsNull() || obj.IsUndefined() || obj.IsNaN() {
		return ""
	}

	return obj.String()
}

// GetRole
//
// English:
//
// Roles define the semantic meaning of content, allowing screen readers and other tools to present and support interaction with an object in a way that is consistent with user expectations of that type of object. roles are added to HTML elements using role="role_type", where role_type is the name of a role in the ARIA specification.
func (e Event) GetRole() (role string) {
	obj := e.Object.Get("role")

	if obj.IsNull() || obj.IsUndefined() || obj.IsNaN() {
		return ""
	}

	return obj.String()
}

// GetSlot
//
// English:
//
// Assigns a slot in a shadow DOM shadow tree to an element: An element with a slot attribute is assigned to the slot created by the <slot> element whose name attribute's value matches that slot attribute's value.
func (e Event) GetSlot() (slot string) {
	obj := e.Object.Get("slot")

	if obj.IsNull() || obj.IsUndefined() || obj.IsNaN() {
		return ""
	}

	return obj.String()
}

// GetSpellcheck
//
// English:
//
// An enumerated attribute defines whether the element may be checked for spelling errors. It may have the following values:
//
// empty string or true, which indicates that the element should be, if possible, checked for spelling errors;
// false, which indicates that the element should not be checked for spelling errors.
func (e Event) GetSpellcheck() (spellcheck bool) {
	obj := e.Object.Get("spellcheck")

	if obj.IsNull() || obj.IsUndefined() || obj.IsNaN() {
		return false
	}

	return obj.Bool()
}

// GetStyle
//
// English:
//
// Contains CSS styling declarations to be applied to the element. Note that it is recommended for styles to be defined in a separate file or files. This attribute and the <style> element have mainly the purpose of allowing for quick styling, for example for testing purposes.
func (e Event) GetStyle() (style string) {
	obj := e.Object.Get("style")

	if obj.IsNull() || obj.IsUndefined() || obj.IsNaN() {
		return ""
	}

	return obj.String()
}

// GetTabIndex
//
// English:
//
// An integer attribute indicating if the element can take input focus (is focusable), if it should participate to sequential keyboard navigation, and if so, at what position. It can take several values:
//
// a negative value means that the element should be focusable, but should not be reachable via sequential keyboard navigation;
// 0 means that the element should be focusable and reachable via sequential keyboard navigation, but its relative order is defined by the platform convention;
// a positive value means that the element should be focusable and reachable via sequential keyboard navigation; the order in which the elements are focused is the increasing value of the tabindex. If several elements share the same tabindex, their relative order follows their relative positions in the document.
func (e Event) GetTabIndex() (tabIndex int) {
	obj := e.Object.Get("tabindex")

	if obj.IsNull() || obj.IsUndefined() || obj.IsNaN() {
		return 0
	}

	return obj.Int()
}

// GetTitle
//
// English:
//
// Contains a text representing advisory information related to the element it belongs to. Such information can typically, but not necessarily, be presented to the user as a tooltip.
func (e Event) GetTitle() (title string) {
	obj := e.Object.Get("title")

	if obj.IsNull() || obj.IsUndefined() || obj.IsNaN() {
		return ""
	}

	return obj.String()
}

// GetTranslate
//
// English:
//
// An enumerated attribute that is used to specify whether an element's attribute values and the values of its Text node children are to be translated when the page is localized, or whether to leave them unchanged. It can have the following values:
//
// empty string or yes, which indicates that the element will be translated.
// no, which indicates that the element will not be translated.
func (e Event) GetTranslate() (translate string) {
	obj := e.Object.Get("translate")

	if obj.IsNull() || obj.IsUndefined() || obj.IsNaN() {
		return ""
	}

	return obj.String()
}
