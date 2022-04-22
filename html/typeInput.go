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
	//  A push button with no default behavior displaying the value of the value attribute, empty by default.
	//
	// Português:
	//
	//  A push button with no default behavior displaying the value of the value attribute, empty by default.
	KInputTypeButton InputType = "button"

	// KInputTypeCheckbox
	//
	// English:
	//
	//  A check box allowing single values to be selected/deselected.
	//
	// Português:
	//
	//  A check box allowing single values to be selected/deselected.
	KInputTypeCheckbox InputType = "checkbox"

	// KInputTypeColor
	//
	// English:
	//
	//  A control for specifying a color; opening a color picker when active in supporting browsers.
	//
	// Português:
	//
	//  A control for specifying a color; opening a color picker when active in supporting browsers.
	KInputTypeColor InputType = "color"

	// KInputTypeDate
	//
	// English:
	//
	//  A control for entering a date (year, month, and day, with no time). Opens a date picker or numeric wheels for year, month, day when active in supporting browsers.
	//
	// Português:
	//
	//  A control for entering a date (year, month, and day, with no time). Opens a date picker or numeric wheels for year, month, day when active in supporting browsers.
	KInputTypeDate InputType = "date"

	// KInputTypeDatetimeLocal
	//
	// English:
	//
	//  A control for entering a date and time, with no time zone. Opens a date picker or numeric wheels for date- and time-components when active in supporting browsers.
	//
	// Português:
	//
	//  A control for entering a date and time, with no time zone. Opens a date picker or numeric wheels for date- and time-components when active in supporting browsers.
	KInputTypeDatetimeLocal InputType = "datetime-local"

	// KInputTypeEmail
	//
	// English:
	//
	//  A field for editing an email address. Looks like a text input, but has validation parameters and relevant keyboard in supporting browsers and devices with dynamic keyboards.
	//
	// Português:
	//
	//  A field for editing an email address. Looks like a text input, but has validation parameters and relevant keyboard in supporting browsers and devices with dynamic keyboards.
	KInputTypeEmail InputType = "email"

	// KInputTypeFile
	//
	// English:
	//
	//  A control that lets the user select a file. Use the accept attribute to define the types of files that the control can select.
	//
	// Português:
	//
	//  A control that lets the user select a file. Use the accept attribute to define the types of files that the control can select.
	KInputTypeFile InputType = "file"

	// KInputTypeHidden
	//
	// English:
	//
	//  A control that is not displayed but whose value is submitted to the server. There is an example in the next column, but it's hidden!
	//
	// Português:
	//
	//  A control that is not displayed but whose value is submitted to the server. There is an example in the next column, but it's hidden!
	KInputTypeHidden InputType = "hidden"

	// KInputTypeImage
	//
	// English:
	//
	//  A graphical submit button. Displays an image defined by the src attribute. The alt attribute displays if the image src is missing.
	//
	// Português:
	//
	//  A graphical submit button. Displays an image defined by the src attribute. The alt attribute displays if the image src is missing.
	KInputTypeImage InputType = "image"

	// KInputTypeMonth
	//
	// English:
	//
	//  A control for entering a month and year, with no time zone.
	//
	// Português:
	//
	//  A control for entering a month and year, with no time zone.
	KInputTypeMonth InputType = "month"

	// KInputTypeNumber
	//
	// English:
	//
	//  A control for entering a number. Displays a spinner and adds default validation when supported. Displays a numeric keypad in some devices with dynamic keypads.
	//
	// Português:
	//
	//  A control for entering a number. Displays a spinner and adds default validation when supported. Displays a numeric keypad in some devices with dynamic keypads.
	KInputTypeNumber InputType = "number"

	// KInputTypePassword
	//
	// English:
	//
	//  A single-line text field whose value is obscured. Will alert user if site is not secure.
	//
	// Português:
	//
	//  A single-line text field whose value is obscured. Will alert user if site is not secure.
	KInputTypePassword InputType = "password"

	// KInputTypeRadio
	//
	// English:
	//
	//  A radio button, allowing a single value to be selected out of multiple choices with the same name value.
	//
	// Português:
	//
	//  A radio button, allowing a single value to be selected out of multiple choices with the same name value.
	KInputTypeRadio InputType = "radio"

	// KInputTypeRange
	//
	// English:
	//
	//  A control for entering a number whose exact value is not important. Displays as a range widget defaulting to the middle value. Used in conjunction min and max to define the range of acceptable values.
	//
	// Português:
	//
	//  A control for entering a number whose exact value is not important. Displays as a range widget defaulting to the middle value. Used in conjunction min and max to define the range of acceptable values.
	KInputTypeRange InputType = "range"

	// KInputTypeReset
	//
	// English:
	//
	//  A button that resets the contents of the form to default values. Not recommended.
	//
	// Português:
	//
	//  A button that resets the contents of the form to default values. Not recommended.
	KInputTypeReset InputType = "reset"

	// KInputTypeSearch
	//
	// English:
	//
	//  A single-line text field for entering search strings. Line-breaks are automatically removed from the input value. May include a delete icon in supporting browsers that can be used to clear the field. Displays a search icon instead of enter key on some devices with dynamic keypads.
	//
	// Português:
	//
	//  A single-line text field for entering search strings. Line-breaks are automatically removed from the input value. May include a delete icon in supporting browsers that can be used to clear the field. Displays a search icon instead of enter key on some devices with dynamic keypads.
	KInputTypeSearch InputType = "search"

	// KInputTypeSubmit
	//
	// English:
	//
	//  A button that submits the form.
	//
	// Português:
	//
	//  A button that submits the form.
	KInputTypeSubmit InputType = "submit"

	// KInputTypeTel
	//
	// English:
	//
	//  A control for entering a telephone number. Displays a telephone keypad in some devices with dynamic keypads.
	//
	// Português:
	//
	//  A control for entering a telephone number. Displays a telephone keypad in some devices with dynamic keypads.
	KInputTypeTel InputType = "tel"

	// KInputTypeText
	//
	// English:
	//
	//  The default value. A single-line text field. Line-breaks are automatically removed from the input value.
	//
	// Português:
	//
	//  The default value. A single-line text field. Line-breaks are automatically removed from the input value.
	KInputTypeText InputType = "text"

	// KInputTypeTime
	//
	// English:
	//
	//  A control for entering a time value with no time zone.
	//
	// Português:
	//
	//  A control for entering a time value with no time zone.
	KInputTypeTime InputType = "time"

	// KInputTypeUrl
	//
	// English:
	//
	//  A field for entering a URL. Looks like a text input, but has validation parameters and relevant keyboard in supporting browsers and devices with dynamic keyboards.
	//
	// Português:
	//
	//  A field for entering a URL. Looks like a text input, but has validation parameters and relevant keyboard in supporting browsers and devices with dynamic keyboards.
	KInputTypeUrl InputType = "url"

	// KInputTypeWeek
	//
	// English:
	//
	//  A control for entering a date consisting of a week-year number and a week number with no time zone.
	//
	// Português:
	//
	//  A control for entering a date consisting of a week-year number and a week number with no time zone.
	KInputTypeWeek InputType = "week"
)
