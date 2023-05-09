//“Exercise 2.2:
//Write a general-purpose unit-conversion program analogous to cf that reads
//numbers from its command-line arguments or from the standard input
//if there are no arguments, and converts each number into
//units like temperature “in Celsius and Fahrenheit,
//length in feet and meters, weight in pounds and kilograms, and the like.”

package main

import (
	"fmt"
	"os"
	"strconv"
)

type Celsius float64
type Fahrenheit float64
type Kelvin float64
type Metre float64
type Feet float64
type Kilogram float64
type Pound float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
	AbsoluteZeroK Kelvin  = 0
	FreezingK     Kelvin  = 273.15
	BoilingK      Kelvin  = 373.15
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%g°K", k) }
func (m Metre) String() string      { return fmt.Sprintf("%g m", m) }
func (l Feet) String() string       { return fmt.Sprintf("%g ft", l) }
func (k Kilogram) String() string   { return fmt.Sprintf("%g kg", k) }
func (p Pound) String() string      { return fmt.Sprintf("%g lb", p) }

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// KToC converts a Kelvin temperature to Celsius.
func KToC(k Kelvin) Celsius { return Celsius(k - 279.15) }

// CToK converts a Celsius temperature to Kelvin.
func CToK(c Celsius) Kelvin { return Kelvin(c + 279.15) }

// FToK converts a Fahrenheit temperature to Kelvin.
func FToK(f Fahrenheit) Kelvin { return Kelvin(((f - 32) * 5 / 9) + 279.15) }

// KToF converts a Kelvin temperature to Fahrenheit.
func KToF(k Kelvin) Fahrenheit { return Fahrenheit((k-279.15)*9/5 + 32) }

// FToM converts a Feets length to Meters.
func FToM(f Feet) Metre { return Metre(f * 0.3048) }

// MToF converts a Metre length to Feets.
func MToF(m Metre) Feet { return Feet(m / 0.3048) }

// PToK converts a Pound weight to Kilograms.
func PToK(p Pound) Kilogram { return Kilogram(p * 0.45359237) }

// KToP converts a Kilogram weight to Pounds.
func KToP(k Kilogram) Pound { return Pound(k / 0.45359237) }

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}

		c := Celsius(t)
		f := Fahrenheit(t)
		l := Feet(t)
		m := Metre(t)
		p := Pound(t)
		k := Kilogram(t)

		fmt.Println("Temperature:")
		fmt.Printf("If it is temperature of %s degrees Celsius then it's %s degrees Fahrenheit\n", c, CToF(c))
		fmt.Printf("If it is temperature of %s degrees Fahrenheit then it's %s degrees Celsius\n", f, FToC(f))

		fmt.Println("Length:")
		fmt.Printf("Test first %v and second %v \n", m, MToF(m))
		fmt.Printf("If it is length of %v metres then it's %v feets\n", m, MToF(m))
		fmt.Printf("If it is length of %v feets then it's %v metres \n", l, FToM(l))

		fmt.Println("Weight:")
		fmt.Printf("If it is weight of %v kilograms then it's %v pounds\n", k, KToP(k))
		fmt.Printf("If it is weight of %v pounds then it's %v kilograms\n", p, PToK(p))

	}

}
