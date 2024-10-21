package helloworld

import "strings"

const (
	spanish    = "spanish"
	french     = "french"
	portuguese = "portuguese"

	englishHelloPrefix    = "Hello, "
	spanishHelloPrefix    = "Hola, "
	frenchHelloPrefix     = "Bonjour, "
	portugueseHelloPrefix = "Olá, "
)

// Hello returns a hello message.
func Hello(name, lang string) string {
	if name == "" {
		return englishHelloPrefix + "World"
	}

	return greeting(lang) + name
}

func greeting(lang string) (prefix string) {

	switch strings.ToLower(lang) {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case portuguese:
		prefix = portugueseHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}
