package livetime

import "fmt"

func Seconds2Time(interval int) string {
	h := interval / 3600
	m := interval % 3600 / 60
	s := interval % 60
	res := fmt.Sprintf("%02d:%02d:%02d", h, m, s)
	return res
}

func Seconds2TimeMinite(interval int) string {
	m := interval / 60
	s := interval % 60
	res := fmt.Sprintf("%02d:%02d", m, s)
	return res
}
