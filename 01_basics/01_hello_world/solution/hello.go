package hello

// Hello returns a greeting message
func Hello() string {
	return "Hello, World!"
}

// HelloWithName returns a personalized greeting message
// If name is empty, it returns the default greeting
func HelloWithName(name string) string {
	if name == "" {
		return "Hello, World!"
	}
	return "Hello, " + name + "!"
} 