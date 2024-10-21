package interation

const repeatCount = 5

func Repeat(character string) string {
	var repeat string
	for i := 0; i < repeatCount; i++ {
		repeat += character
	}
	return repeat
}
