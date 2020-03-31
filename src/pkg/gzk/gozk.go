package gzk

import (
	"fmt"
	"time"
	"github.com/samuel/go-zookeeper/zk"
)

var hosts = []string{"10.6.124.21:2181"}

func Connect() *zk.Conn {
	conn, _ , err := zk.Connect(hosts, time.Second * 3)
	if err != nil {
		panic(err)
	}
	return conn
}

// get 
func Get(conn *zk.Conn, path string) error {
	v, _, err := conn.Get(path)
    if err != nil {
        return err
    }
    fmt.Println("Get Value:", string(v))
    return nil
}

// create
func Create(conn *zk.Conn, path string, data []byte) error {
	var flags int32 = 0
	//flags有4种取值：
	//0:永久，除非手动删除
	//zk.FlagEphemeral = 1:短暂，session断开则改节点也被删除
	//zk.FlagSequence  = 2:会自动在节点后面添加序号
	//3:Ephemeral和Sequence，即，短暂且自动添加序号
	var acls = zk.WorldACL(zk.PermAll)	//控制访问权限模式
	p, err_create := conn.Create(path, data, flags, acls)
    if err_create != nil {
        fmt.Println("Create Error:", err_create)
        return err_create
    }
    fmt.Println("created:", p)
    return nil
}

// watchchild
func mirror(conn *zk.Conn, path string) (chan []string, chan error) {
    snapshots := make(chan []string)
    errors := make(chan error)
 
    go func() {
        for {
            snapshot, _, events, err := conn.ChildrenW(path)
            if err != nil {
                errors <- err
                return
            }
            snapshots <- snapshot
            evt := <-events
            if evt.Err != nil {
                errors <- evt.Err
                return
            }
        }
    }()
 
    return snapshots, errors
}
func WatchChild(conn *zk.Conn, path string) {
    snapshots, errors := mirror(conn, path)
    for i := 0; i < 5; i ++ {
    	fmt.Println("i:", i)
        select {
        case snapshot := <-snapshots:
            fmt.Printf("%+v\n", snapshot)
        case err := <-errors:
            fmt.Printf("%+v\n", err)
        }
    }
}

// get watch
func GetWatch(conn *zk.Conn, path string) {
	for i := 0; i < 5; i ++ {
		v, _, ech ,err := conn.GetW(path)
		if err != nil {
			fmt.Println("Get Watch Error:", err)
			return
		}
		fmt.Println("Value:", string(v))
		event := <-ech
		fmt.Println("path:", event.Path)
		fmt.Println("type:", event.Type.String())
		fmt.Println("state:", event.State.String())
	}
}

// Usage Example

// func main() {
//     conn := Connect()
//     defer conn.Close()
//     err := Get(conn, "/zkd")
//     fmt.Println("err:", err)
// }