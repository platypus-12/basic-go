package evalk_test

import (
	"fmt"
	"math"
	"testing"

	"gopl.io/ch7/eval"
)

func TestEval(t *testing.T) {
	tests := []struct {
		expr string
		env  eval.Env
		want string
	}{
		{"sqrt(A / pi)", eval.Env{"A": 87616, "pi": math.Pi}, "167"},
		{"pow(x, 3) + pow(y, 3)", eval.Env{"x": 12, "y": 1}, "1729"},
		{"5 / 9 * (F - 32)", eval.Env{"F": -40}, "-40"},
		{"5 / 9 * (F-32)", eval.Env{"F": 32}, "0"},
		{"5/9 * (F-32)", eval.Env{"F": 212}, "100"},
	}

	var prevExpr string
	for _, test := range tests {
		if test.expr != prevExpr {
			fmt.Printf("\n%s\n", test.expr)
			prevExpr = test.expr
		}
		expr, err := eval.Parse(test.expr)
		if err != nil {
			t.Error(err)
			continue
		}
		got := fmt.Sprintf("%.6g", expr.Eval(test.env))
		fmt.Printf("\t%v => %s\n", test.env, got)
		if got != test.want {
			t.Errorf("%s.Eval() in %v = %q, want %q\n", test.expr, test.env, got, test.want)
		}

	}
	//t.Errorf("das")
}
