package model

import "goms/gotest/app1/utils"

type OrderInfo struct {
	Id         int
	TotalMoney int
}

func OrderById(id int) OrderInfo {
	return OrderInfo{Id: id, TotalMoney: utils.RandomInt(10, 10000)}
}

type OrderInfoList []OrderInfo
