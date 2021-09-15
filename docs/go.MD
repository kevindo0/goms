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

### 2、 难点
#### 1. map
#### 2. channel
1. [channel实现原理剖析](https://my.oschina.net/renhc/blog/2246871)

#### 3. interface
#### 4. defer原理
#### 5. 调试工具dlv
#### 6. reflect
