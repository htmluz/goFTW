package main

import "fmt"

const (
	spanish    = "Spanish"
	portuguese = "Portuguese"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	portuguesePrefix   = "Ola, "
)

func greetingPrefix(language string) (prefix string) {
	switch language {
	case portuguese:
		prefix = portuguesePrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
