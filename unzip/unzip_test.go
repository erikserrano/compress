package unzip

import "testing"

// TestUnzip test Unzip function
func TestUnzip(t *testing.T) {
	files, err := Unzip("C:\\Users\\serrer01\\Downloads\\RR\\go\\zips\\I_07044_032_2016010_000510.zip", "C:\\Users\\serrer01\\Downloads\\DROP\\")
	if err != nil {
		t.Fatal(err)
	} else if len(files) != 56 {
		t.Fatalf("Expected %v items but got %v", 56, len(files))
	}
}

func BenchmarkUnzip(b *testing.B) {
	files, err := Unzip("C:\\Users\\serrer01\\Downloads\\RR\\go\\zips\\I_07044_032_2016010_000510.zip", "C:\\Users\\serrer01\\Downloads\\DROP\\")
	if err != nil {
		b.Fatal(err)
	} else if len(files) != 56 {
		b.Fatalf("Expected %v items but got %v", 56, len(files))
	}
}
