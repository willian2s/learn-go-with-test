package interation

import (
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("repeat a character provides repeat count 0", func(t *testing.T) {
		repeated := Repeat("a", 0)
		expected := strings.Repeat("a", 5)

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})

	t.Run("repeat a character", func(t *testing.T) {
		repeated := Repeat("a", 15)
		expected := strings.Repeat("a", 15)

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 0)
	}
}
