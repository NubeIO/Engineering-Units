package units

func mergeAllUnits() {
	out := map[string]map[string]string{}

	out["mass"] = massSymbols

}

var massSymbols = map[string]string{
	"kg":        "kg",
	"g":         "g",
	"mg":        "mg",
	"metricTon": "metric ton",
	"lb":        "lb",
	"oz":        "oz",
	"grain":     "grain",
	"shortTon":  "short ton",
	"longTon":   "long ton",
	"slug":      "slug",
}

var massConversions = map[string]float64{
	"kg":        1.0,
	"g":         1000.0,
	"mg":        1000000.0,
	"metricTon": 1.0 / 1000.0,
	"lb":        2.2046226218,
	"oz":        35.274,
	"grain":     2.2046226218 * 7000.0,
	"shortTon":  1.0 / 907.185,
	"longTon":   1.0 / 1016.047,
	"slug":      1.0 / 14.5939029,
}

// Mass is a mass object that can store a mass value and convert between units of mass.
type Mass struct {
	EngUnit
}

// NewMass creates a new Mass object.
func NewMass(value float64, unit string) *Mass {

	return &Mass{
		EngUnit: EngUnit{
			value:       value,
			unit:        unit,
			conversions: massConversions,
			symbols:     massSymbols,
		},
	}
}
