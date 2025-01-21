package main

import "fmt"

const englishHelloPrefix = "Hello"
const spanishHelloPrefix = "Hola"
const frenchHelloPrefix = "Bonjour"
const portugueseHelloPrefix = "Olá"
const japaneseHelloPrefix = "こんにちは"

func Hello(name string, language string) string {
	if name == "" {
		return englishHelloPrefix + ", World!"
	}

	prefix := englishHelloPrefix
	switch language {
	case "Spanish":
		prefix = spanishHelloPrefix
	case "French":
		prefix = frenchHelloPrefix
	case "Portuguese":
		prefix = portugueseHelloPrefix
	case "Japanese":
		prefix = japaneseHelloPrefix
	}

	return fmt.Sprintf("%s, %s!", prefix, name)
}

func main() {
	fmt.Println(Hello("World", ""))
}
