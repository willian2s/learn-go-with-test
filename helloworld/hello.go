package helloworld

const englishHelloPrefix = "Hello, "

// Hello returns a hello message.
func Hello(name string) string {
	if name == "" {
		return englishHelloPrefix + "World"
	}
	return englishHelloPrefix + name
}
