//Exercise 7.6:
//Add support for Kelvin temperatures to tempflag.

package main

import (
	"flag"
	"fmt"
)

type Celsius float64
type Fahrenheit float64
type Kelvin float64

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9.0/5.0 + 32.0) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32.0) * 5.0 / 9.0) }
func KToC(k Kelvin) Celsius     { return Celsius(k - 273.15) }
func CToK(c Celsius) Kelvin     { return Kelvin(c + 273.15) }

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

type Value interface {
	String() string
	Set(string) error
}

type temperatureFlag struct {
	value float64
	unit  string
}

func (f *temperatureFlag) Set(s string) error {
	var value float64
	var unit string
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C", "°C":
		f.value = float64(Celsius(value))
		f.unit = "°C"
		return nil
	case "F", "°F":
		f.value = float64(FToC(Fahrenheit(value)))
		f.unit = "°C"
		return nil
	case "K", "°K":
		f.value = float64(KToC(Kelvin(value)))
		f.unit = "°C"
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

func (f *temperatureFlag) String() string {
	return fmt.Sprintf("%g%s", f.value, f.unit)
}

func TemperatureFlag(name string, value float64, usage string) *float64 {
	f := temperatureFlag{value: value, unit: "°C"}
	flag.CommandLine.Var(&f, name, usage)
	return &f.value
}

func main() {
	temp := TemperatureFlag("temp", 20.0, "the temperature")

	flag.Parse()
	fmt.Println(*temp)
}
