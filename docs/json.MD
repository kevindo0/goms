### 一、功能点
#### 1、 json tag定制struct编码
* `json:"-"`忽略字段
```go
struct {
        Name string `json:"-"`
        No   int
    }{"Hello", 112}
// {"No":112}
```
* `json:“myName”`制定字段名称
```go
struct {
        Name string `json:"myName"`
        No   int
    }{"Hello", 112}
// {"myName":"Hello","No":112}
```
* `json:",omitempty”`或`json:“myName,omitempty”`忽略空值字段
```go
struct {
    Name string `json:"myName,omitempty"`
    No   int
}{"Hello", 112}
// {"myName":"Hello","No":112}
  
struct {
    Name string `json:"myName,omitempty"`
    No   int
}{}
// {"No":0}
```
* `json:",string"`
```go
struct {
    Name string
    No   int `json:",string"`
}{"Hello", 112}
// {"Name":"Hello","No":"112"}
```
* 内嵌匿名结构会自动展开,可使用`json:“myName”`取消自动展开
```go
type Foo struct {
    Name string
}
type Bar struct {
    Foo
    No int
}
v := Bar{Foo{"Hello"}, 112}
// {"Name":"Hello","No":112}
  
type Foo struct {
    Name string
}
type Bar struct {
    Foo `json:"Foo"`
    No  int
}
v := Bar{Foo{"Hello"}, 112}
// {"Foo":{"Name":"Hello"},"No":112}
```
