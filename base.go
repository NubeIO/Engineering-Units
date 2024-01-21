package units

import (
	"fmt"
)

type Unit interface {
	ChangeUnit(unit string) float64
	ChangeUnitAsSymbol(unit string, decimalPlace int) string
	AsSymbol() string
	AsSymbolWithDecimal(decimalPlace int) string
	Convert(toUnit string) float64
	CheckUnit(unit string) error
	GetSymbols() map[string]string
	GetSymbol(key string) string
}

type EngineeringUnits interface {
	Mass(value float64, unit string) Unit
	Length(value float64, unit string) Unit
	Temperature(value float64, unit string) Unit
	Current(value float64, unit string) Unit
	Time(value float64, unit string) Unit
	Pressure(value float64, unit string) Unit
	Force(value float64, unit string) Unit
	Power(value float64, unit string) Unit
	Flow(value float64, unit string) Unit
}

type Units struct{}

func New() Units {
	return Units{}
}

func (u Units) Conversion(category, unit string, value float64) (Unit, error) {
	switch category {
	case "mass":
		return NewMass(value, unit), nil
	case "length":
		return NewLength(value, unit), nil
	case "temperature":
		return NewTemperature(value, unit), nil
	case "current":
		return NewCurrent(value, unit), nil
	case "time":
		return NewEngTime(value, unit), nil
	case "pressure":
		return NewPressure(value, unit), nil
	case "force":
		return NewForce(value, unit), nil
	case "power":
		return NewPower(value, unit), nil
	case "flow":
		return NewFlow(value, unit), nil
	}
	return nil, fmt.Errorf("category: %s", category)
}
func (u Units) Mass(value float64, unit string) Unit {

	return NewMass(value, unit)
}

func (u Units) Length(value float64, unit string) Unit {
	return NewLength(value, unit)
}

func (u Units) Temperature(value float64, unit string) Unit {
	return NewTemperature(value, unit)
}

func (u Units) Current(value float64, unit string) Unit {
	return NewCurrent(value, unit)
}

func (u Units) Time(value float64, unit string) Unit {
	return NewEngTime(value, unit)
}

func (u Units) Pressure(value float64, unit string) Unit {
	return NewPressure(value, unit)
}

func (u Units) Force(value float64, unit string) Unit {
	return NewForce(value, unit)
}

func (u Units) Power(value float64, unit string) Unit {
	return NewPower(value, unit)
}

func (u Units) Flow(value float64, unit string) Unit {
	return NewFlow(value, unit)
}
