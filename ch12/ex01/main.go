package main

import (
	"basic-go/ch12/ex01/display"
)

func main() {
	keygahairetusmap := make(map[[2]string]string)
	_key1 := [2]string{"1", "1"}
	keygahairetusmap[_key1] = "11"
	_key2 := [2]string{"2", "2"}
	keygahairetusmap[_key2] = "22"

	type Key struct {
		element1 string
		element2 string
	}

	_skey10 := Key{
		element1: "10-1",
		element2: "10-2",
	}

	_skey20 := Key{
		element1: "20-1",
		element2: "20-2",
	}

	keygakouzoutai := make(map[Key]string)
	keygakouzoutai[_skey10] = "10"
	keygakouzoutai[_skey20] = "20"

	type Movie struct {
		Title, Subtitle  string
		Year             int
		Color            bool
		Actor            map[string]string
		Oscars           []string
		Sequel           *string
		KeyGaHairetsuMap map[[2]string]string
		KeyGaKouzoutai   map[Key]string
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
		KeyGaHairetsuMap: keygahairetusmap,
		KeyGaKouzoutai: keygakouzoutai,
	}
	display.Display("strangelove", strangelove)
	// fmt.Printf("%v\n", reflect.ValueOf(strangelove).Kind())
	// fmt.Println(reflect.ValueOf(strangelove).Field(1))
	// fmt.Println(reflect.ValueOf(strangelove).Field(1).Kind())
	// fmt.Println(reflect.ValueOf(strangelove).Type().Field(1).Name)
	// fmt.Println(reflect.ValueOf(strangelove).Field(1).Kind(), "s")
	// fmt.Println(reflect.ValueOf(strangelove).Type().Field(4).Name)
	// fmt.Println(reflect.ValueOf(strangelove).Field(4))
	// fmt.Println(reflect.ValueOf(strangelove).Type().Field(5).Name)
	// fmt.Println(reflect.ValueOf(strangelove).Field(5))
}
