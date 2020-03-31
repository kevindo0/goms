package main
import (
    "io"
    "fmt"
    "log"
    "net"
    "bufio"
    "reflect"
)

// 串行指定读取客户端返回内容大小
func Serial(conn net.Conn) {
    buf := make([]byte, 1024) //定义一个切片的长度是1024。

    n,err := conn.Read(buf) //接收到的内容大小。

    if err != nil && err != io.EOF {  //io.EOF在网络编程中表示对端把链接关闭了。
        log.Fatal(err)
    }
    fmt.Println(string(buf[:n])) //将接受的内容都读取出来。
}

// 按照指定方式循环读取
func Circle(conn net.Conn) {
    buf := make([]byte, 6)
    for {
        n, err := conn.Read(buf) //接收到的内容大小
        if err == io.EOF {
            break
        }
        fmt.Println(n, string(buf[:n]))
    }
}

func ReadByLine(conn net.Conn) {
    r := bufio.NewReader(conn)
    for {
        line, err := r.ReadString('\n')
        if err == io.EOF {
            break
        }
        fmt.Println(line)
    }
}

func main() {
    addr := ":10001"
    //拨号操作，需要指定协议。
    conn,err := net.Dial("tcp",addr)
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    fmt.Println("访问公网IP地址是：",conn.RemoteAddr().String()) /*获取“conn”中的公网地址。注意：最好是加上后面的String方法，因为他们的那些是不一样的哟·当然你打印的时候
        可以不加输出结果是一样的，但是你的内心是不一样的哟！*/
    fmt.Printf("客户端链接的地址及端口是：%v\n",conn.LocalAddr()) //获取到本地的访问地址和端口。
    fmt.Println("“conn.LocalAddr()”所对应的数据类型是：",reflect.TypeOf(conn.LocalAddr()))
    fmt.Println("“conn.RemoteAddr().String()”所对应的数据类型是：",reflect.TypeOf(conn.RemoteAddr().String()))
    Serial(conn)
    // Circle(conn)
    // ReadByLine(conn)
}