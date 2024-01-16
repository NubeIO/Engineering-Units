package units

// Flow is a flow object that can store a flow value and convert between units of flow.
type Flow struct {
	EngUnit
}

// NewFlow creates a new Flow object.
func NewFlow(value float64, unit string) *Flow {
	return &Flow{
		EngUnit: EngUnit{
			value:       value,
			unit:        unit,
			conversions: flowConversions,
			symbols:     flowSymbols,
		},
	}
}

var flowSymbols = map[string]string{
	"BBL/hr":   "BBL/hr",
	"BBL/min":  "BBL/min",
	"BBL/sec":  "BBL/sec",
	"GAL/hr":   "GAL/hr",
	"GAL/min":  "GAL/min",
	"GAL/sec":  "GAL/sec",
	"ft/min":   "ft/min",
	"LFM":      "LFM",
	"m/s":      "m/s",
	"miles/hr": "MPH",
	"MPH":      "MPH",
	"ft3/min":  "ft3/min",
	"CFM":      "CFM",
	"m3/hr":    "m3/hr",
	"L/s":      "L/s",
}

var flowConversions = map[string]float64{
	"BBL/hr":  3600.0,
	"BBL/min": 60.0,
	"BBL/sec": 1.0,
	"GAL/hr":  151200.0,
	"GAL/min": 2520.0,
	"GAL/sec": 42.0,
	"ft/min":  0.00508, // 1 ft/min = 0.00508 m/s
	"LFM":     0.00508, // 1 LFM = 0.00508 m/s
	"m/s":     1.0,
	"MPH":     0.44704, // 1 MPH = 0.44704 m/s
	"ft3/min": 0.471947,
	"CFM":     0.471947,
	"m3/hr":   3600.0,
	"L/s":     1000.0,
}
