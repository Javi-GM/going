package variables

const floatnumber float64 = 3.14

func DeclareInteger() int {
	integer := 42

	return integer
}

func DeclareFloat() float64 {
	return floatnumber
}

func DeclareString() string {
	return "Go is fun"
}

func DeclareBool() bool {
	return true
}

func MultipleDeclarations() (int, float64, string) {
	intVar, floatVar, stringVar := 10, 15.5, "multiple"

	return intVar, floatVar, stringVar
}

