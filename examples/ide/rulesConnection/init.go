package rulesConnection

import "image/color"

// TypeOfDataCurrentlyInEffect Defines which language the system recognizes
var TypeOfDataCurrentlyInEffect = KGoLang

// It initializes the pointers of the type of languages
func init() {
	// Add the new language constructor inside the switch
	switch TypeOfDataCurrentlyInEffect {
	case KGoLang:
		conversion = new(golangManagerTypeConversion)

		// case KCPlusPlus:
		//   ...
	}

	// ATTENTION: Do not change the pointer indications below

	// Simplifies the object use of the object
	TypeToColor = conversion.TypeToColor
	TypeVerify = conversion.Verify
	GetError = conversion.GetError
}

// ATTENTION: Do not change the pointers functions below

// Explanation:
//
// The `Conversion RulesDataType` object has been initialized in the `init()` function and features of functions such as
// `TypeToColor = Conversion.TypeToColor` facilitate syntax during use

// conversion Defines the Type Converter pointer
var conversion RulesDataType

// TypeToColor For each type of data, numerical, string, boolean ... a visual identity color is attributed
var TypeToColor func(dataType string) (color color.RGBA)

// TypeVerify Check the syntax looking for errors
var TypeVerify func(dataType string)
var GetError func() (err error)
