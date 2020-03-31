package main

import (
   "encoding/json"
   "fmt"
)

type groupList struct {
    Status int `json:"status"`
    Data []struct {
      Group string `json:"group"`
      Version string `json:"version"`
      Count int `json:"count"`
      ConsumeType string `json:"consumeType"`
      MessageModel string `json:"messageModel"`
      ConsumeTps int `json:"consumeTps"`
      DiffTotal int `json:"diffTotal"`
   } `json:"data"`
   ErrMsg string `json:"errMsg"`
}

var body string = "\"staus\":0, \"data\":[],\"errMsg\":\"Err\""

func main() {
    var g groupList
    _ = json.Unmarshal(body, &g)
    fmt.Println(g)
}