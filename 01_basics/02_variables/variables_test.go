package variables

import (
	"testing"
)

func TestDeclareInteger(t *testing.T) {
	expected := 42
	actual := DeclareInteger()
	
	if actual != expected {
		t.Errorf("Expected %d but got %d", expected, actual)
	}
}

func TestDeclareFloat(t *testing.T) {
	expected := 3.14
	actual := DeclareFloat()
	
	if actual != expected {
		t.Errorf("Expected %f but got %f", expected, actual)
	}
}

func TestDeclareString(t *testing.T) {
	expected := "Go is fun"
	actual := DeclareString()
	
	if actual != expected {
		t.Errorf("Expected %q but got %q", expected, actual)
	}
}

func TestDeclareBool(t *testing.T) {
	expected := true
	actual := DeclareBool()
	
	if actual != expected {
		t.Errorf("Expected %t but got %t", expected, actual)
	}
}

func TestMultipleDeclarations(t *testing.T) {
	expectedX, expectedY, expectedZ := 10, 15.5, "multiple"
	actualX, actualY, actualZ := MultipleDeclarations()
	
	if actualX != expectedX {
		t.Errorf("For x: Expected %d but got %d", expectedX, actualX)
	}
	
	if actualY != expectedY {
		t.Errorf("For y: Expected %f but got %f", expectedY, actualY)
	}
	
	if actualZ != expectedZ {
		t.Errorf("For z: Expected %q but got %q", expectedZ, actualZ)
	}
}

// Bonus challenge test
// func TestSwapVariables(t *testing.T) {
// 	testCases := []struct {
// 		a        int
// 		b        int
// 		expectedA int
// 		expectedB int
// 	}{
// 		{5, 10, 10, 5},
// 		{-3, 8, 8, -3},
// 		{0, 42, 42, 0},
// 	}
// 
// 	for _, tc := range testCases {
// 		a, b := tc.a, tc.b
// 		a, b = SwapVariables(a, b)
// 		
// 		if a != tc.expectedA || b != tc.expectedB {
// 			t.Errorf("SwapVariables(%d, %d) = (%d, %d); want (%d, %d)",
// 				tc.a, tc.b, a, b, tc.expectedA, tc.expectedB)
// 		}
// 	}
// } 