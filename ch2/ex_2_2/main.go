//“Exercise 2.2:
//Write a general-purpose unit-conversion program analogous to cf that reads
//numbers from its command-line arguments or from the standard input
//if there are no arguments, and converts each number into
//units like temperature in Celsius and Fahrenheit,
//length in feet and meters, weight in pounds and kilograms, and the like.”

package main

import (
	"fmt"
	tempconv "gopl.io/ch2/ex_2_2_lib"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		c := tempconv.Celsius(t)
		f := tempconv.Fahrenheit(t)
		l := tempconv.Feet(t)
		m := tempconv.Metre(t)
		p := tempconv.Pound(t)
		k := tempconv.Kilogram(t)

		fmt.Println("Temperature:")
		fmt.Printf("If it is temperature of %s degrees Celsius"+
			" then it's %s degrees Fahrenheit\n", c, tempconv.CToF(c))
		fmt.Printf("If it is temperature of %s degrees Fahrenheit"+
			" then it's %s degrees Celsius\n", f, tempconv.FToC(f))

		fmt.Println("Length:")
		fmt.Printf("If it is length of %s metres"+
			" then it's %s feets\n", m, tempconv.MToF(m))
		fmt.Printf("If it is length of %s feets"+
			" then it's %s metres \n", l, tempconv.FToM(l))

		fmt.Println("Weight:")
		fmt.Printf("If it is weight of %s kilograms"+
			" then it's %s pounds\n", k, tempconv.KToP(k))
		fmt.Printf("If it is weight of %s pounds"+
			" then it's %s kilograms\n", p, tempconv.PToK(p))

	}
}
