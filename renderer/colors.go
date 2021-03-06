package renderer

import (
	"image/color"
)

// Background colours
var NormalBackground = color.RGBA{0xFF, 0xFF, 0xDC, 0xFF}
var InsertBackground = color.RGBA{0xEF, 0xFF, 0xEF, 0xFF}
var DeleteBackground = color.RGBA{0xFF, 0xDC, 0xDC, 0xFF}
var TaglineBackground = color.RGBA{0xEA, 0xFF, 0xFF, 0xFF}

// Text colours
var TextHighlight = color.RGBA{0xBC, 0xFC, 0xF9, 0xFF}
var TextColour = color.RGBA{0, 0, 0, 0xFF}
var CommentColour = color.RGBA{0, 0, 0xFF, 0xFF}
var KeywordColour = color.RGBA{0x6D, 0x6D, 0, 0xFF}
var BuiltinTypeColour = color.RGBA{0x00, 0x6D, 0, 0xFF}
var StringColour = color.RGBA{0xFF, 0, 0, 0xFF}

// Mostly for HTML-like languages
var TagDelimitorColour = color.RGBA{0x70, 0x94, 0xAE, 0xFF} // < and > characters in HTML
var AttributeColour = BuiltinTypeColour
var TagColour = color.RGBA{0x8A, 0x66, 0x00, 0xFF} //KeywordColour
var OperatorColour = color.RGBA{0x55, 0x55, 0x55, 0xFF}
