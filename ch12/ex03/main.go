package main

import (
	"basic-go/ch12/ex03/sexpr"
	"fmt"
)

type Movie struct {
	Title, Subtitle string
	Year            int
	Color           bool
	Actor           map[string]string
	Oscars          []string
	Sequel          *string
	ComplexNumber   complex64
	FloatNumber     float32
	Interface       interface{}
}

func main() {
	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned",
		Year:     1964,
		Color:    true,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		},
		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
		ComplexNumber: 23 + 3i,
		FloatNumber:   1.3,
		Interface:     "AA",
	}
	// fmt.Println(sexpr.Marshal(strangelove))
	a, b := sexpr.Marshal(strangelove)
	fmt.Printf("%s\n", a)
	fmt.Println(b)

}
