package main

import "fmt"

const englishHelloPrefix = "Hello"
const spanishHelloPrefix = "Hola"
const frenchHelloPrefix = "Bonjour"
const portugueseHelloPrefix = "Olá"
const japaneseHelloPrefix = "こんにちは"
const italianHelloPrefix = "Ciao"
const chineseHelloPrefix = "你好"
const dutchHelloPrefix = "Hallo"
const malaioHelloPrefix = "Halo"
const koreanHelloPrefix = "안녕하세요"

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
	case "Italian":
		prefix = italianHelloPrefix
	case "Chinese":
		prefix = chineseHelloPrefix
	case "Malaio":
		prefix = malaioHelloPrefix
	case "Korean":
		prefix = koreanHelloPrefix
	case "Dutch":
		prefix = dutchHelloPrefix
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
