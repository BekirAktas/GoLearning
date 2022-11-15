package main

import (
	"fmt"
)

const (
	turkish = "Turkish"
	spanish = "Spanish"
 	englishHelloPrefix = "Hello, "
 	spanishHelloPrefix = "Hola, "
	turkishHelloPrefix = "Merhaba, "
)

func Hello(name string, language string) string{
	if name == "" {
		name = "World"
	}

	// if language == spanish {
	// 	return spanishHelloPrefix + name
	// }

	// if language == turkish {
	// 	return turkishHelloPrefix + name
	// }

	


	prefix, _ := greetingPrefix(language)
	return prefix + name
}

func greetingPrefix(language string)(string, int) {
	prefix := ""
	sayi:= 10;
	switch language {
		case turkish:
			prefix = turkishHelloPrefix
		case spanish:
			prefix = spanishHelloPrefix
		default:
			prefix = englishHelloPrefix
	}

	return prefix, sayi

}

func main() {
	fmt.Println(Hello("Deniz", ""))
}