package units

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func PrintJOSN(x interface{}) {
	ioWriter := os.Stdout
	w := json.NewEncoder(ioWriter)
	w.SetIndent("", "    ")
	w.Encode(x)
}

func Test_main(t *testing.T) {
	eu := Units{}

	mass := eu.Mass(1000, "g")
	fmt.Printf("mass %s \n", mass.AsSymbol())

	mass = eu.Mass(1000, "g")
	s := mass.GetSymbols()

	PrintJOSN(s)

	err := mass.CheckUnit("kg")
	fmt.Println(err)
	if err != nil {
		return
	}
	m := mass.ChangeUnit("KG")
	fmt.Printf("mass %f \n", m)
	fmt.Printf("Converted mass: %v\n", mass.AsSymbol())

	temperature := eu.Temperature(75.0, "F")
	f := temperature.ChangeUnit("C")
	fmt.Printf("Converted Temperature: %f \n", f)
	fmt.Printf("Converted Temperature: %v\n", temperature.AsSymbol())

	PrintJOSN(MergeAllUnits())

}

func TestGetUnitByKey(t *testing.T) {
	eu := New()
	conversion, err := eu.Conversion("temperature", "F", 75.5)
	if err != nil {
		return
	}
	fmt.Printf("Unit: %v\n", conversion.AsSymbolWithDecimal(0))
	fmt.Printf("Unit: %v\n", conversion.ChangeUnitAsSymbol("C", 2))

}
