package unzip

import "testing"
import "sync"

// TestUnzip test Unzip function
func TestUnzip(t *testing.T) {
	files, err := Unzip("C:\\Users\\serrer01\\Downloads\\RR\\go\\zips\\I_07044_032_2016010_000510.zip", "C:\\Users\\serrer01\\Downloads\\DROP\\")
	if err != nil {
		t.Fatal(err)
	} else if len(files) != 56 {
		t.Fatalf("Expected %v items but got %v", 56, len(files))
	}
}

// BenchmarkUnzip benchmark Unzip function
func BenchmarkUnzip(b *testing.B) {
	files := []string{
		"C:\\Users\\serrer01\\Downloads\\RR\\go\\zips\\I_07044_032_2016010_000510.zip",
		"C:\\Users\\serrer01\\Downloads\\RR\\go\\zips\\I_07044_032_2016010_000511.zip",
		"C:\\Users\\serrer01\\Downloads\\RR\\go\\zips\\I_07044_032_2016010_000513.zip",
		"C:\\Users\\serrer01\\Downloads\\RR\\go\\zips\\I_07044_032_2016010_000514.zip",
		"C:\\Users\\serrer01\\Downloads\\RR\\go\\zips\\I_07044_032_2016010_000517.zip",
	}
	var wg sync.WaitGroup

	for _, path := range files {
		wg.Add(1)
		go func(sipPath string) {
			defer wg.Done()

			_, err := Unzip(sipPath, "C:\\Users\\serrer01\\Downloads\\DROP\\")
			if err != nil {
				b.Fatal(err)
			}
		}(path)
	}
	wg.Wait()
}
