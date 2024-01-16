package units

var timeSymbols = map[string]string{
	"ms":     "ms",
	"s":      "s",
	"minute": "minute",
	"hr":     "hr",
	"day":    "day",
}

var timeConversions = map[string]float64{
	"ms":     1000.0,
	"s":      1.0,
	"minute": 1.0 / 60.0,
	"hr":     1.0 / 60.0 / 60.0,
	"day":    1.0 / 60.0 / 60.0 / 24.0,
}

// EngTime is a time object that can store a time value and convert between units of time.
type EngTime struct {
	EngUnit
}

// NewEngTime creates a new EngTime object.
func NewEngTime(value float64, unit string) *EngTime {
	return &EngTime{
		EngUnit: EngUnit{
			value:       value,
			unit:        unit,
			conversions: timeConversions,
			symbols:     timeSymbols,
		},
	}
}
