package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(max())
	fmt.Println(max(2, 3, 1, 5, -4))
	fmt.Println(min())
	fmt.Println(min(-3, 5, 3, -11))

	//fmt.Println(maxa())
	fmt.Println(maxa(2, 3, 1, 5, -4))
	//fmt.Println(mina())
	fmt.Println(mina(-3, 5, 3, -11))
}

func max(vals ...int) (int, error) {
	if len(vals) == 0 {
		err := fmt.Errorf("引数は1つ以上必要です")
		return 0, err
	}
	_max := math.MinInt
	for _, val := range vals {
		if _max < val {
			_max = val
		}
	}
	return _max, nil
}

func min(vals ...int) (int, error) {
	if len(vals) == 0 {
		err := fmt.Errorf("引数は1つ以上必要です")
		return 0, err
	}
	_min := math.MaxInt
	for _, val := range vals {
		if _min > val {
			_min = val
		}
	}
	return _min, nil
}

func maxa(first int, vals ...int) int {
	_max := first
	for _, val := range vals {
		if _max < val {
			_max = val
		}
	}
	return _max
}

func mina(first int, vals ...int) int {
	_min := first
	for _, val := range vals {
		if _min > val {
			_min = val
		}
	}
	return _min
}
