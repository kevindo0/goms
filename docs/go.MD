### 1、 知识点
#### 1、 字符串转成json数据
```golang
//json str 转map
jsonStr := `{"host": "http://localhost:9090"}`
var dat map[string]interface{}
if err := json.Unmarshal([]byte(jsonStr), &dat); err == nil {
    fmt.Println("=======json str 转map=======")
    fmt.Println(dat)
}
```
