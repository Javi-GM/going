package hello

func Hello() string {
	return "Hello, World!"
}

func HelloWithName(name string) string {
	if name == "" {
		return Hello()
	}
	return "Hello, " +  name + "!"
}