package units

var lengthSymbols = map[string]string{
	"fm":        "fm",
	"pm":        "pm",
	"nm":        "nm",
	"um":        "um",
	"mm":        "mm",
	"cm":        "cm",
	"m":         "m",
	"dam":       "dam",
	"hm":        "hm",
	"km":        "km",
	"Mm":        "Mm",
	"Gm":        "Gm",
	"Tm":        "Tm",
	"Pm":        "Pm",
	"inch":      "inch",
	"ft":        "ft",
	"yd":        "yd",
	"mi":        "mi",
	"nautMi":    "nautMi",
	"lightYear": "lightYear",
}

var lengthConversions = map[string]float64{
	"fm":        1e15,
	"pm":        1e12,
	"nm":        1e9,
	"um":        1e6,
	"mm":        1e3,
	"cm":        1e2,
	"m":         1.0,
	"dam":       0.1,
	"hm":        0.01,
	"km":        0.001,
	"Mm":        1e-6,
	"Gm":        1e-9,
	"Tm":        1e-12,
	"Pm":        1e-15,
	"inch":      39.3701,
	"ft":        3.28084,
	"yd":        1.09361,
	"mi":        0.000621371,
	"nautMi":    1.0 / 1852.0,
	"lightYear": 1.0 / (9.4607304725808e15),
}

// Length is a length object that can store a length value and convert between units of length.
type Length struct {
	EngUnit
}

// NewLength creates a new Length object.
func NewLength(value float64, unit string) *Length {
	return &Length{
		EngUnit: EngUnit{
			value:       value,
			unit:        unit,
			conversions: lengthConversions,
			symbols:     lengthSymbols,
		},
	}
}
