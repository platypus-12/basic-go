package main

import "fmt"

func main() {
	fmt.Println(sum(1))
	fmt.Println(sum(1, 3, 4))
	fmt.Println(sum(1, -2, 5))
	fmt.Println(sum())
	vals := []int{1, 2, 3, 4}
	fmt.Println(sum(vals...))

}

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}
