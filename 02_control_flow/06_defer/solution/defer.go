package defer_statement

import (
	"errors"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

// ProcessFile reads the content of a file and returns it as a string
// It uses defer to ensure the file is properly closed
func ProcessFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close() // Ensures file is closed when the function returns
	
	content, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	
	return string(content), nil
}

// RunTransaction executes a series of operations, stopping if any fails
// It uses defer to ensure proper cleanup
func RunTransaction(operations []func() error) error {
	// If we return an error in the cleanup, that becomes our return value,
	// which could mask the original error
	var err error
	
	// Track if we need to rollback (stop processing)
	shouldRollback := false
	
	// Defer our cleanup handler
	defer func() {
		if shouldRollback {
			// In a real application, you might want to do something
			// to roll back or clean up the transaction here
			fmt.Println("Transaction rolled back")
		} else {
			fmt.Println("Transaction committed")
		}
	}()
	
	// Execute operations until one fails
	for _, op := range operations {
		if opErr := op(); opErr != nil {
			shouldRollback = true
			err = opErr
			return err // Return immediately, defer will still execute
		}
	}
	
	return nil
}

// TrackTime measures and prints the execution time of an operation
func TrackTime(operation func()) {
	start := time.Now()
	
	// Defer the time tracking code so it runs after operation completes
	defer func() {
		elapsed := time.Since(start)
		fmt.Printf("Operation took %v\n", elapsed)
	}()
	
	// Execute the operation
	operation()
}

// FilterPanic executes a potentially risky function
// and calls the handler if a panic occurs
func FilterPanic(handler func(interface{}), riskyFunc func()) {
	// Defer panic recovery
	defer func() {
		if r := recover(); r != nil {
			// Call the provided handler with the panic value
			handler(r)
		}
	}()
	
	// Execute the potentially risky function
	riskyFunc()
}

// ModifyReturnValue demonstrates how defer can modify a named return value
func ModifyReturnValue() (result int) {
	// Set initial value
	result = 21
	
	// Defer a function that doubles the result
	defer func() {
		result *= 2
	}()
	
	// This value will be modified by the deferred function
	return result
}

// MultiResourceCleanup shows how to properly clean up multiple resources with defer
// It opens and reads two files while holding a mutex lock
func MultiResourceCleanup(file1Path, file2Path string, mu *sync.Mutex, shouldError bool) (string, error) {
	// Lock the mutex
	mu.Lock()
	defer mu.Unlock() // Ensure mutex is unlocked when the function exits
	
	// Open the first file
	file1, err := os.Open(file1Path)
	if err != nil {
		return "", err
	}
	defer file1.Close() // Ensure first file is closed
	
	// Open the second file
	file2, err := os.Open(file2Path)
	if err != nil {
		return "", err
	}
	defer file2.Close() // Ensure second file is closed
	
	// Simulate an error if needed (for testing)
	if shouldError {
		return "", errors.New("simulated error for testing")
	}
	
	// Read both files
	content1, err := io.ReadAll(file1)
	if err != nil {
		return "", err
	}
	
	content2, err := io.ReadAll(file2)
	if err != nil {
		return "", err
	}
	
	// Combine and return the contents
	return string(content1) + "\n" + string(content2), nil
}

// LogFunction logs the entry, exit, and any errors from a function
func LogFunction(name string, operation func() error) {
	fmt.Printf("ENTER: %s\n", name)
	
	// Log exit with defer, including any errors
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("PANIC in %s: %v\n", name, r)
			// Re-panic to propagate the panic up the call stack
			panic(r)
		}
	}()
	
	// Execute the operation and capture error
	err := operation()
	
	// Defer this after error capture so we have access to err
	defer func() {
		if err != nil {
			fmt.Printf("ERROR in %s: %v\n", name, err)
		} else {
			fmt.Printf("EXIT: %s\n", name)
		}
	}()
	
	// Note: The error will be available to the deferred function
	// Even though we don't return it here
} 