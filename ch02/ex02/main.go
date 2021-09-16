package main

import (
	"basic-go/ch02/ex02/unitconv"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args[1:]) == 0 {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			t, err := strconv.ParseFloat(input.Text(), 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "cf: %v\n", err)
				os.Exit(1)
			}
			view_result(t)
		}
	}

	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		view_result(t)
	}
}

func view_result(t float64) {
	fa := unitconv.Fahrenheit(t)
	c := unitconv.Celsius(t)
	fe := unitconv.Feet(t)
	m := unitconv.Meters(t)
	p := unitconv.Pound(t)
	k := unitconv.KiloGram(t)
	fmt.Printf("%s = %s, %s = %s\n%s = %s, %s = %s\n%s = %s, %s = %s\n", fa, unitconv.FToC(fa), c, unitconv.CToF(c), fe, unitconv.FToM(fe), m, unitconv.MToF(m), p, unitconv.PToK(p), k, unitconv.KToP(k))
}
