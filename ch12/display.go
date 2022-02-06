package main

import (
	"basic-go/ch12/display"
	"fmt"
)

func main() {
	fmt.Println("sss")
	type Movie struct {
		Title, Subtitle string
		Year            int
		Color           bool
		Actor           map[string]string
		Oscars          []string
		Sequel          *string
	}
	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned",
		Year:     1964,
		Color:    false,
		Actor: map[string]string{
			"aaa": "AAA",
			"bbb": "BBB",
			"ccc": "CCC",
		},

		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
		},
	}
	display.Display("strangelove", strangelove)

}
