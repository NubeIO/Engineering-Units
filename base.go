package units

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

type Unit interface {
	ChangeUnit(unit string) float64
	AsSymbol() string
	Convert(toUnit string) float64
	CheckUnit(unit string) error
	GetSymbols() map[string]string
	GetSymbol(key string) string
}
