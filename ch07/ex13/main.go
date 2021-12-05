package main

import (
	"basic-go/ch07/ex13/eval"
	"fmt"
)

func main() {
	expr1, err1 := eval.Parse("-2")
	fmt.Println(expr1, err1)
	expr2, err2 := eval.Parse("!true")
	fmt.Println(expr2, err2)
	expr3, err3 := eval.Parse("-34 + 12 + pow(x, 3)")
	fmt.Println(expr3, err3)
	expr4, err4 := eval.Parse(expr3.String())
	fmt.Println(expr4, err4)
}
