package utils


import (
    "fmt"
    "encoding/binary"
    "math/rand"
    "time"
    "sync"
    "encoding/hex"
    "strings"
)
type ID string

func RandomIDGenerator() func() ID {
    buf := make([]byte, 8)
    var seed int64
    if _, err := rand.Read(buf); err == nil {
        seed = int64(binary.BigEndian.Uint64(buf))
    } else {
        seed = int64(time.Now().Nanosecond())
    }
    fmt.Printf("%d\n", seed)
    var (
        mu sync.Mutex
        rng = rand.New(rand.NewSource(seed))
    )
    return func() ID {
        mu.Lock()
        defer mu.Unlock()
        id := make([]byte, 16)
        rng.Read(id)
        return encodeID(id)
    }
}

func encodeID(b []byte) ID {
    id := hex.EncodeToString(b)
    id = strings.TrimLeft(id, "0")
    if id == "" {
        id = "0"
    }
    return ID(id)
}