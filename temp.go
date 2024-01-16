package units

// Temperature is a temperature object that can store a temperature value and convert between units of temperature.
type Temperature struct {
	EngUnit
}

var temperatureSymbols = map[string]string{
	"K": "K",
	"C": "°C",
	"R": "°R",
	"F": "°F",
}

var temperatureConversions = map[string]float64{
	"K": 1.0,
	"C": 1.0,
	"R": 1.8,
	"F": 1.0 / 1.8,
}

func (t *Temperature) Convert(toUnit string) float64 {
	fromUnit := t.unit
	if fromUnit == "F" && toUnit == "C" {
		// Convert from Fahrenheit to Celsius
		return (t.value - 32) * 5 / 9
	} else if fromUnit == "C" && toUnit == "F" {
		// Convert from Celsius to Fahrenheit
		return (t.value * 9 / 5) + 32
	} else if fromUnit == "F" && toUnit == "K" {
		// Convert from Fahrenheit to Kelvin
		return (t.value + 459.67) * 5 / 9
	} else if fromUnit == "K" && toUnit == "F" {
		// Convert from Kelvin to Fahrenheit
		return (t.value * 9 / 5) - 459.67
	} else if fromUnit == "C" && toUnit == "K" {
		// Convert from Celsius to Kelvin
		return t.value + 273.15
	} else if fromUnit == "K" && toUnit == "C" {
		// Convert from Kelvin to Celsius
		return t.value - 273.15
	} else if fromUnit == "R" && toUnit == "F" {
		// Convert from Rankine to Fahrenheit
		return t.value - 459.67
	} else if fromUnit == "F" && toUnit == "R" {
		// Convert from Fahrenheit to Rankine
		return t.value + 459.67
	}
	return t.EngUnit.Convert(toUnit)
}

func (t *Temperature) ChangeUnit(unit string) float64 {
	t.value = t.Convert(unit)
	t.unit = unit
	return t.value
}

// NewTemperature creates a new Temperature object.
func NewTemperature(value float64, unit string) *Temperature {
	return &Temperature{
		EngUnit: EngUnit{
			value:       value,
			unit:        unit,
			conversions: temperatureConversions,
			symbols:     temperatureSymbols,
		},
	}
}
