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
	s := mass.GetSymbols()

	PrintJOSN(s)

	err := mass.CheckUnit("aaaa")
	fmt.Println(err)
	if err != nil {
		return
	}
	m := mass.ChangeUnit("KG")
	fmt.Printf("mass %f \n", m)
	fmt.Printf("Converted mass: %v\n", mass.AsSymbol())

	temperature := eu.Temperature(75.0, "F")
	f := temperature.ChangeUnit("C")
	fmt.Printf("Converted Temperature: %f\n", f)
	fmt.Printf("Converted Temperature: %v\n", temperature.AsSymbol())

	//PrintJOSN(MergeAllUnits())

}
