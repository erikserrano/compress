package zip

import (
	"fmt"
	"os"
	"sync"
	"testing"
)

func showZipped(info []os.FileInfo) {
	for _, val := range info {
		println(val.Name())
	}
}

func zipDirectory(inputPath, outputPath string, length int) error {
	var wg sync.WaitGroup
	var lastError error
	for i := 0; i < length; i++ {
		wg.Add(1)
		go func(k int) {
			defer wg.Done()
			if files, err := Zip(inputPath, fmt.Sprintf("test-%v.zip", k), outputPath); err != nil {
				lastError = err
			} else {
				showZipped(files)
			}
		}(i)
	}
	wg.Wait()
	return lastError
}

func TestZip(t *testing.T) {
	if err := zipDirectory("/Users/sero/gocode/pkg/", "/Users/sero/Downloads/", 10); err != nil {
		t.Errorf("%s", err)
	}
}

func BenchmarkZip(b *testing.B) {
	if err := zipDirectory("/Users/sero/gocode/pkg/", "/Users/sero/Downloads/", b.N); err != nil {
		b.Errorf("%s", err)
	}
}
