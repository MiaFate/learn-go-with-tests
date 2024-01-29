package main

import "fmt"

const (
	french     = "French"
	spanish    = "Spanish"
	portuguese = "Portuguese"

	englishHelloPrefix    = "Hello, "
	spanishHelloPrefix    = "Hola, "
	frenchHelloPrefix     = "Bonjour, "
	portugueseHelloPrefix = "Olá, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingsPrefix(language) + name
}

func greetingsPrefix(language string) (prefix string) {

	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case portuguese:
		prefix = portugueseHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Mia", ""))
}
