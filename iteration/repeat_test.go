package iteration

import "testing"

func TestFiveTimes(t *testing.T) {
	repeated := FiveTimes("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected '%s' but got '%s'", expected, repeated)
	}
}

func BenchmarkFiveTimes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FiveTimes("a")
	}
}
