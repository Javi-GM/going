package hello

import "fmt"

func Hello() string {
	return "Hello, World!"
}

func HelloWithName(name string) string {
	if name == "" {
		return Hello()
	}
	return fmt.Sprintf("Hello, %s!", name)
}