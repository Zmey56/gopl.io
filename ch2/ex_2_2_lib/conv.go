package tempconv

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// FToM converts a Feets length to Meters.
func FToM(f Feet) Metre { return Metre(f * 0.3048) }

// MToF converts a Metre length to Feets.
func MToF(m Metre) Feet { return Feet(m / 0.3048) }

// PToK converts a Pound weight to Kilograms.
func PToK(p Pound) Kilogram { return Kilogram(p * 0.45359237) }

// KToP converts a Kilogram weight to Pounds.
func KToP(k Kilogram) Pound { return Pound(k / 0.45359237) }

//!-
