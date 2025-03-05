package interation

import (
	"fmt"
	"testing"
)

type CharacterTest struct {
	title     string
	character string
	count     int
	expected  string
}

func TestRepeat(t *testing.T) {

	tests := []CharacterTest{
		{title: "Repetição normal", character: "a", count: 5, expected: "aaaaa"},
		{title: "Repetição de 'b'", character: "b", count: 3, expected: "bbb"},
		{title: "Repetição com count zero", character: "c", count: 0, expected: ""},
		{title: "Repetição com count um", character: "d", count: 1, expected: "d"},
		{title: "Repetição de string maior que um caractere", character: "xy", count: 2, expected: "xyxy"},
		{title: "Repetição com count negativo", character: "z", count: -1, expected: ""},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			repeated := Repeat(test.character, test.count)

			if repeated != test.expected {
				t.Errorf("expected %q but got %q", test.expected, repeated)
			}
		})
	}
}

func ExampleRepeat() {
	repated := Repeat("a", 5)
	fmt.Println((repated))
	// Output: aaaaa
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
