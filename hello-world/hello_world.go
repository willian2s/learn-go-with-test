package main

import "fmt"

const englishHelloPrefix = "Hello"
const spanishHelloPrefix = "Hola"
const frenchHelloPrefix = "Bonjour"
const portugueseHelloPrefix = "Olá"
const japaneseHelloPrefix = "こんにちは"

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "Spanish":
		prefix = spanishHelloPrefix
	case "French":
		prefix = frenchHelloPrefix
	case "Portuguese":
		prefix = portugueseHelloPrefix
	case "Japanese":
		prefix = japaneseHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func Hello(name string, language string) string {
	if name == "" {
		return englishHelloPrefix + ", World!"
	}

	prefix := greetingPrefix(language)

	return fmt.Sprintf("%s, %s!", prefix, name)
}

func main() {
	fmt.Println(Hello("World", ""))
}
