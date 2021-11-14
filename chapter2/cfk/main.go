package main

import (
	"fmt"
	"os"
	"strconv"
	"tgopl/chapter2/tempconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cfk: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		k := tempconv.Kelvin(t)
		fmt.Printf("%s = %s = %s\n%s = %s = %s\n%s = %s = %s\n",
			f, tempconv.FToC(f), tempconv.FToK(f), c, tempconv.CToF(c),
			tempconv.CToK(c), k, tempconv.KToC(k), tempconv.KToF(k))
	}
}
