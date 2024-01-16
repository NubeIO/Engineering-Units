package units

var currentSymbols = map[string]string{
	"A":  "A",
	"mA": "mA",
	"kA": "kA",
}

var currentConversions = map[string]float64{
	"A":  1.0,
	"mA": 1e3,
	"kA": 1e-3,
}

// Current is a current (amperage) object that can store a current (amperage) value and convert between units of current (amperage).
type Current struct {
	EngUnit
}

// NewCurrent creates a new Current object.
func NewCurrent(value float64, unit string) *Current {
	return &Current{
		EngUnit: EngUnit{
			value:       value,
			unit:        unit,
			conversions: currentConversions,
			symbols:     currentSymbols,
		},
	}
}
