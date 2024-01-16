package units

// Power is a power object that can store a power value and convert between units of power.
type Power struct {
	EngUnit
}

// NewPower creates a new Power object.
func NewPower(value float64, unit string) *Power {
	return &Power{
		EngUnit: EngUnit{
			value:       value,
			unit:        unit,
			conversions: powerConversions,
			symbols:     powerSymbols,
		},
	}
}

var powerSymbols = map[string]string{
	"kW":             "kW",
	"BTU/hr":         "BTU/hr",
	"BTU/min":        "BTU/min",
	"BTU/sec":        "BTU/sec",
	"cal/sec":        "cal/sec",
	"cal/min":        "cal/min",
	"cal/hr":         "cal/hr",
	"erg/sec":        "erg/sec",
	"erg/min":        "erg/min",
	"erg/hr":         "erg/hr",
	"ftlb/sec":       "ftlb/sec",
	"GW":             "GW",
	"MW":             "MW",
	"kCal/sec":       "kCal/sec",
	"kCal/min":       "kCal/min",
	"kCal/hr":        "kCal/hr",
	"mW":             "mW",
	"W":              "W",
	"VA":             "VA",
	"hp_mech":        "hp_mech",
	"hp_ele":         "hp_ele",
	"hp_metric":      "hp_metric",
	"metric_ton_ref": "metric_ton_ref",
	"US_ton_ref":     "US_ton_ref",
	"J/sec":          "J/sec",
	"J/min":          "J/min",
	"J/hr":           "J/hr",
	"kgf-m/sec":      "kgf-m/sec",
}

var powerConversions = map[string]float64{
	"kW":             1.0,
	"BTU/hr":         3412.14,
	"BTU/min":        56.869,
	"BTU/sec":        0.94781666666,
	"cal/sec":        238.85,
	"cal/min":        238.85 * 60,
	"cal/hr":         238.85 * 60 * 60,
	"erg/sec":        10e9,
	"erg/min":        10e9 * 60,
	"erg/hr":         10e9 * 60 * 60,
	"ftlb/sec":       737.56,
	"GW":             1e-6,
	"MW":             1e-3,
	"kCal/sec":       0.24,
	"kCal/min":       0.24 * 60,
	"kCal/hr":        0.24 * 60 * 60,
	"mW":             1e6,
	"W":              1e3,
	"VA":             1e3,
	"hp_mech":        1.3410220888,
	"hp_ele":         1.3404825737,
	"hp_metric":      1.359621617304,
	"metric_ton_ref": 0.259,
	"US_ton_ref":     0.2843451361,
	"J/sec":          1000.0,
	"J/min":          1000.0 * 60,
	"J/hr":           1000.0 * 60 * 60,
	"kgf-m/sec":      101.97162129779,
}
