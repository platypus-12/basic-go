package main

import (
	"fmt"
	"sort"
)

func main() {
	ages := make(map[string]int)
	ages["alice"] = 31
	ages["charlie"] = 34
	fmt.Println(ages)
	ages2 := map[string]int{
		"alice":   31,
		"charlie": 34,
	}
	fmt.Println(ages2)
	kara := map[string]int{}
	fmt.Println(kara)
	fmt.Println(ages["alice"])
	delete(ages, "alice")
	fmt.Println(ages["alice"])
	fmt.Println(ages)

	ages["bob"] = ages["bob"] + 1

	ages["tom"] += 1

	ages["ken"]++

	fmt.Println(ages)

	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}

	//var names []string
	names := make([]string, 0, len(ages))
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	fmt.Println("----------")
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}
}
