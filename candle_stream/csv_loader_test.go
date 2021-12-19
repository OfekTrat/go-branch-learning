package candlestream

import "testing"

// func TestReadLine(t *testing.T) {

// }

func TestLoadingCsv(t *testing.T) {
	path := ".\\test_data\\test.csv"
	stream := LoadCandleStreamFromCsv(path)

	if length := stream.Length(); length != 2 {
		t.Logf("Wrong Length, Got: %v, Expected: %v", length, 2)
		t.Log(stream.candles)
		t.Error("AssertionError")
	}
	c0 := stream.Get(0)
	c1 := stream.Get(1)

	if c0.Get("a") != 1.1 {
		t.Error("AssertionError")
	}
	if c1.Get("b") != 5.666 {
		t.Error("AssertionError")
	}

}
