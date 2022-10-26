package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Metre float64
type Feet float64
type Pound float64
type Kilogram float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%.2f°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%.2f°F", f) }
func (m Metre) String() string      { return fmt.Sprintf("%.2f", m) }
func (f Feet) String() string       { return fmt.Sprintf("%.2f", f) }
func (p Pound) String() string      { return fmt.Sprintf("%.2f", p) }
func (k Kilogram) String() string   { return fmt.Sprintf("%.2f", k) }

//!-
