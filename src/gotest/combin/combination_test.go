package gtest

import (
    "fmt"
    "os"
    "math/rand"
    "testing"
)

func TestMain(m *testing.M) {
    fmt.Println("Init...")
    os.Exit(m.Run())
}

// 单元测试
// 测试全局函数，以TestFunction命名
// 测试类成员函数，以TestClass_Function命名
func TestCombination(t *testing.T) {
    // 这里定义一个临时的结构体来存储测试case的参数以及期望的返回值
    for _, unit := range []struct {
        m        int
        n        int
        expected int
    }{
        {1, 0, 1},
        {4, 1, 4},
        {4, 2, 6},
        {4, 3, 4},
        {4, 4, 1},
        {10, 1, 10},
        {10, 3, 120},
        {10, 7, 120},
    } {
        // 调用排列组合函数，与期望的结果比对，如果不一致输出错误
        if actually := combination(unit.m, unit.n); actually != unit.expected {
            t.Errorf("combination: [%v], actually: [%v]", unit, actually)
        }
    }
}

// 性能测试
func BenchmarkCombination(b *testing.B) {
    // b.N会根据函数的运行时间取一个合适的值
    for i := 0; i < b.N; i++ {
        combination(i+1, rand.Intn(i+1))
    }
}

// 并发性能测试
func BenchmarkCombinationParallel(b *testing.B) {
    // 测试一个对象或者函数在多线程的场景下面是否安全
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            m := rand.Intn(100) + 1
            n := rand.Intn(m)
            combination(m, n)
        }
    })
}