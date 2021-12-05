package main

import (
	"basic-go/ch07/ex15/eval"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var formula string
	fmt.Println("enter the formura u wanna calculate!")
	fmt.Scan(&formula)

	expr, err := eval.Parse(formula)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	vars := make(map[eval.Var]bool)
	if err := expr.Check(vars); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	env := eval.Env{}
	for {
		if generateEnv(env, vars) == nil {
			break
		}
		fmt.Println("your input is invalid.")
	}

	fmt.Printf("answer: %g\n", expr.Eval(env))
}

//mapは、参照型なので値渡しでも参照渡しでも何でもいいっぽい(第一引数の話)
func generateEnv(env eval.Env, vars map[eval.Var]bool) error {
	for _var := range vars {
		fmt.Printf("what is the value of the variable %s?\n", _var)
		var value string
		fmt.Scan(&value)
		value_f, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
		env[_var] = value_f
	}
	return nil
}
