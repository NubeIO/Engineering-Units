package units

import (
	"fmt"
)

// EngUnit is a generic engineering unit object.
type EngUnit struct {
	value       float64
	unit        string
	conversions map[string]float64
	symbols     map[string]string
}

func (e *EngUnit) GetSymbol(key string) string {
	symbol, ok := e.symbols[key]
	if !ok {
		return ""
	}
	return symbol
}

func (e *EngUnit) GetSymbols() map[string]string {
	return e.symbols
}

func (e *EngUnit) CheckUnit(toUnit string) error {
	_, ok := e.conversions[toUnit]
	if !ok {
		suggestion := "try: "
		count := 0
		for validUnit := range e.conversions {
			count++
			if count > 1 {
				break
			}
			suggestion += validUnit + ", "
		}
		return fmt.Errorf("unsupported  to-unit: %s, %s", toUnit, suggestion)
	}
	fromUnit := e.unit
	_, ok = e.conversions[fromUnit]
	if !ok {
		suggestion := "try: "
		count := 0
		for validUnit := range e.conversions {
			count++
			if count > 1 {
				break
			}
			suggestion += validUnit + ", "
		}
		return fmt.Errorf("unsupported from-unit: %s, %s", fromUnit, suggestion)
	}

	return nil
}

// Convert converts the object from one unit to another.
func (e *EngUnit) Convert(toUnit string) float64 {
	fromUnit := e.unit
	return e.value / e.conversions[fromUnit] * e.conversions[toUnit]
}

// ChangeUnit converts the current value of the object to a new unit.
func (e *EngUnit) ChangeUnit(unit string) float64 {
	e.value = e.Convert(unit)
	e.unit = unit
	return e.value
}

// SetValue sets the value and unit of the object.
func (e *EngUnit) SetValue(value float64, unit string) {
	e.value = value
	e.unit = unit
}

// GetValue returns a slice of the float value and unit of the object.
func (e *EngUnit) GetValue() []interface{} {
	return []interface{}{e.value, e.unit}
}

// AsSymbol returns a string representation of the value and unit using the unit symbol.
func (e *EngUnit) AsSymbol() string {
	symbol, exists := e.symbols[e.unit]
	if exists {
		return fmt.Sprintf("%.2f %s", e.value, symbol)
	}
	return fmt.Sprintf("%.2f %s", e.value, e.unit)
}

func MergeAllUnits() map[string]map[string]string {
	out := map[string]map[string]string{}

	// Mass units and symbols
	out["mass"] = massSymbols

	// Length units and symbols
	out["length"] = lengthSymbols

	// Temperature units and symbols
	out["temperature"] = temperatureSymbols

	// Current units and symbols
	out["current"] = currentSymbols

	// Time units and symbols
	out["time"] = timeSymbols

	// Pressure units and symbols
	out["pressure"] = pressureSymbols

	// Force units and symbols
	out["force"] = forceSymbols

	// Power units and symbols
	out["power"] = powerSymbols

	// Flow units and symbols
	out["flow"] = flowSymbols

	return out
}
