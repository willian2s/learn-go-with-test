package main

import "fmt"

const englishHelloPrefix = "Hello"

func Hello(name string) string {
	if name == "" {
		return englishHelloPrefix + ", World!"
	}

	return fmt.Sprintf("%s, %s!", englishHelloPrefix, name)
}

func main() {
	fmt.Println(Hello("World"))
}
