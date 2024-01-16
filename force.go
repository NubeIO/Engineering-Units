package units

// Force is a force object that can store a force value and convert between units of force.
type Force struct {
	EngUnit
}

// NewForce creates a new Force object.
func NewForce(value float64, unit string) *Force {
	return &Force{
		EngUnit: EngUnit{
			value:       value,
			unit:        unit,
			conversions: forceConversions,
			symbols:     forceSymbols,
		},
	}
}

var forceSymbols = map[string]string{
	"N":         "N",
	"kN":        "kN",
	"MN":        "MN",
	"GN":        "GN",
	"gf":        "gf",
	"kgf":       "kgf",
	"dyn":       "dyn",
	"J/m":       "J/m",
	"J/cm":      "J/cm",
	"shortTonF": "shortTonF",
	"longTonF":  "longTonF",
	"kipf":      "kipf",
	"lbf":       "lbf",
	"ozf":       "ozf",
	"pdl":       "pdl",
}

var forceConversions = map[string]float64{
	"N":         1.0,
	"kN":        1.0 / 1000.0,
	"MN":        1.0 / 1000000.0,
	"GN":        1.0 / 1000000000.0,
	"gf":        1.019716213e+2,
	"kgf":       1.019716213e-1,
	"dyn":       1e+5,
	"J/m":       1.0,
	"J/cm":      100.0,
	"shortTonF": 1.124045e-4,
	"longTonF":  1.003611e-4,
	"kipf":      2.248089e-4,
	"lbf":       2.248089431e-1,
	"ozf":       3.5969430896,
	"pdl":       7.2330138512,
}
