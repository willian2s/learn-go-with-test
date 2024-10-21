package interation

func Repeat(character string, repeatCount int) string {
	count := 5
	var repeat string

	if repeatCount != 0 {
		count = repeatCount
	}

	for i := 0; i < count; i++ {
		repeat += character
	}
	return repeat
}
