package zip

import (
    "testing"
    "sync"
    "fmt"
)

func TestZip(t *testing.T) {
    var wg sync.WaitGroup
    for i := 0; i < 0; i++ {
        wg.Add(1)
        go func(k int) {
            defer wg.Done()
            if err := Zip("E:/GoogleAPPEngine/Java/libs/", fmt.Sprintf("test-%v.zip", k), "C:/Users/Areté/Downloads/"); err != nil {
                t.Fatal(err.Error())
            }
        }(i)
    }
    wg.Wait()
}

func BenchmarkZip(b *testing.B) {
    for i := 0; i < b.N; i++ {
        if err := Zip("E:/GoogleAPPEngine/Java/libs/", fmt.Sprintf("test-%v.zip", i), "C:/Users/Areté/Downloads/"); err != nil {
            b.Fatal(err.Error())
        }
    }
}
