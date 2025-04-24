package hello

import "testing"

func TestHello(t *testing.T) {
	expected := "Hello, World!"
	actual := Hello()
	
	if actual != expected {
		t.Errorf("Expected %q but got %q", expected, actual)
	}
}

func TestHelloWithName(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty string", "", "Hello, World!"},
		{"with name", "Alice", "Hello, Alice!"},
		{"with another name", "Bob", "Hello, Bob!"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := HelloWithName(test.input)
			if actual != test.expected {
				t.Errorf("Expected %q but got %q", test.expected, actual)
			}
		})
	}
} 