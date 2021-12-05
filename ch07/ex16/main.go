package main

import (
	"basic-go/ch07/ex16/eval"
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/result", calc)
	mux.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

func calc(w http.ResponseWriter, req *http.Request) {
	formula := req.URL.Query().Get("formula")
	if formula == "" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "no formula value.")
		return
	}
	fmt.Println(formula)

	expr, err := eval.Parse(formula)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "your input is invald. %s", err)
	}

	vars := make(map[eval.Var]bool)
	if err := expr.Check(vars); err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "your input is invald. %s", err)
	}
	env := eval.Env{}
	fmt.Fprintf(w, "answer: %g\n", expr.Eval(env))

}

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, `
	<html>
		<form action='/result' method='get'>
		<div>
		<label for='formula'>計算式</label>
		<input type='text' id='formula' name='formula'>
	</div>
	<input type='submit' value='計算する'></html>
	`)
}
