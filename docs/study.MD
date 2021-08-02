#### 学习目录
1. copy build in func
2. make 和 new的区别
3. sync/context
4. [channel/select](https://studygolang.com/articles/11320)
5. [golang nil](https://zhuanlan.zhihu.com/p/151140497)
6. uintptr
7. interface

#### 1.copy
func copy(dst, src []byte) int
```golang
//1.不同类型的切片无法复制
//2.如果s1的长度大于s2的长度，将s2中对应位置上的值替换s1中对应位置的值
//3.如果s1的长度小于s2的长度，多余的将不做替换
func main() {
    s1 := []int{1, 2, 3}
    s2 := []int{4, 5}
    s3 := []int{6, 7, 8, 9}
    copy(s1, s2)
    fmt.Println(s1) //[4 5 3]
    copy(s2, s3)
    fmt.Println(s2) //[6 7]
}
```
#### 2、make 和 new的区别
func make(t Type, size ...IntegerType) Type    
make: 被用来分配引用类型的内存： map, slice, channel    
new: 被用来分配除了引用类型的所有其他类型的内存： int, string, array等  
**异同**：  
make返回的还是这三个引用类型本身；而new返回的是指向类型的指针。
```golang
//创建一个初始元素长度为5的数组切片，元素初始值为0： 
mySlice1 := make([]int, 5) 
//创建一个初始元素长度为5的数组切片，元素初始值为0，并预留10个元素的存储空间： 
mySlice2 := make([]int, 5, 10) 
//切片字面量创建长度为5容量为5的切片,需要注意的是 [ ] 里面不要写数组的容量，因为如果写了个数以后就是数组了，而不是切片了。
mySlice3 := []int{10,20,30,40,50}
```