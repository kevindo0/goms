package gmath

import (
	"fmt"
	"strconv"
	"math"
)

func SumInt(s []int) int {
	sum := 0
	for _, v := range s {
		sum += v
	}
	return sum
}

func Average(s []float64) float64 {
	var res float64 = 0
	length := len(s)
	if length > 0 {
		for _, v := range s {
			res += v
		}
		res /= float64(length)
	}
	res, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", res), 64)
	return res
}

func variance(s []float64) float64 {
	var res float64 = 0
	length := len(s)
	if length > 0 {
		avg := Average(s)
		for _, v := range s {
			res += math.Pow(v - avg, 2)
		}
		res /= float64(length)
	}
	return res
}

func Var(s []float64) float64 {
	res := variance(s)
	res, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", res), 64)
	return res
}

func Std(s []float64) float64 {
	res := variance(s)
	res = math.Sqrt(res)
	res, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", res), 64)
	return res
}