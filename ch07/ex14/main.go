package main

import (
	"basic-go/ch07/ex14/eval"
	"fmt"
)

func main() {
	expr1, err1 := eval.Parse("pow(a,b)")
	fmt.Println(expr1, err1)
	vars := make(map[eval.Var]bool)
	if err := expr1.Check(vars); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(expr1.Eval(eval.Env{"a": 2.0, "b": 3.0}))
	}

	expr2, err2 := eval.Parse("min(a,b,1)")
	fmt.Println(expr2, err2)
	vars2 := make(map[eval.Var]bool)
	if err := expr2.Check(vars2); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(expr2.Eval(eval.Env{"a": 2.0, "b": 4.0}))
	}

	expr3, err3 := eval.Parse("min()")
	fmt.Println(expr3, err3)
	vars3 := make(map[eval.Var]bool)
	if err := expr3.Check(vars3); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(expr3.Eval(eval.Env{"a": 2.0, "b": 4.0}))
	}
}
