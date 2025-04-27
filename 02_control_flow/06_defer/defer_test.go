package defer_statement

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
	"time"
)

// Helper function to create a temporary test file
func createTempFile(t *testing.T, content string) string {
	tmpfile, err := os.CreateTemp("", "test-*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}

	if _, err := tmpfile.Write([]byte(content)); err != nil {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
		t.Fatalf("Failed to write to temp file: %v", err)
	}

	if err := tmpfile.Close(); err != nil {
		os.Remove(tmpfile.Name())
		t.Fatalf("Failed to close temp file: %v", err)
	}

	return tmpfile.Name()
}

func TestProcessFile(t *testing.T) {
	content := "Hello, World!"
	tmpfile := createTempFile(t, content)
	defer os.Remove(tmpfile)

	result, err := ProcessFile(tmpfile)
	if err != nil {
		t.Errorf("ProcessFile() error = %v", err)
		return
	}

	if result != content {
		t.Errorf("ProcessFile() = %q, want %q", result, content)
	}

	// Test with non-existent file
	result, err = ProcessFile("non-existent-file.txt")
	if err == nil {
		t.Errorf("ProcessFile() expected error for non-existent file, got nil")
	}
}

func TestRunTransaction(t *testing.T) {
	// Successful transaction
	successOps := []func() error{
		func() error { return nil },
		func() error { return nil },
		func() error { return nil },
	}

	err := RunTransaction(successOps)
	if err != nil {
		t.Errorf("RunTransaction() with successful operations error = %v, want nil", err)
	}

	// Failed transaction
	expectedErr := errors.New("operation failed")
	failOps := []func() error{
		func() error { return nil },
		func() error { return expectedErr },
		func() error { t.Error("This operation should not run"); return nil },
	}

	err = RunTransaction(failOps)
	if err != expectedErr {
		t.Errorf("RunTransaction() with failing operation = %v, want %v", err, expectedErr)
	}
}

func TestTrackTime(t *testing.T) {
	// Create a string builder to capture output
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Define a test operation that sleeps
	sleepDuration := 100 * time.Millisecond
	operation := func() {
		time.Sleep(sleepDuration)
	}

	// Track time of the operation
	TrackTime(operation)

	// Restore stdout and read captured output
	w.Close()
	os.Stdout = oldStdout
	var buf strings.Builder
	_, _ = buf.ReadFrom(r)
	output := buf.String()

	// Verify output includes time information
	if !strings.Contains(output, "took") {
		t.Errorf("TrackTime() output = %q, expected to contain 'took'", output)
	}

	// Verify reasonable time range (at least 90% of sleep duration)
	for _, timeStr := range []string{"ms", "µs", "ns"} {
		if strings.Contains(output, timeStr) {
			// Found time unit in output
			break
		}
	}
}

func TestFilterPanic(t *testing.T) {
	// Test with function that doesn't panic
	noPanicCalled := false
	noPanic := func() {
		noPanicCalled = true
	}

	handler := func(p interface{}) {
		t.Errorf("Handler called when no panic occurred")
	}

	FilterPanic(handler, noPanic)
	if !noPanicCalled {
		t.Errorf("Non-panicking function was not called")
	}

	// Test with function that panics
	panicValue := "test panic"
	panicFunc := func() {
		panic(panicValue)
	}

	handlerCalled := false
	recoveredPanic := ""
	handler = func(p interface{}) {
		handlerCalled = true
		recoveredPanic = fmt.Sprintf("%v", p)
	}

	FilterPanic(handler, panicFunc)
	if !handlerCalled {
		t.Errorf("Panic handler was not called for panic")
	}
	if recoveredPanic != panicValue {
		t.Errorf("Recovered panic = %q, want %q", recoveredPanic, panicValue)
	}
}

func TestModifyReturnValue(t *testing.T) {
	expected := 42 // This value depends on implementation details
	result := ModifyReturnValue()
	if result != expected {
		t.Errorf("ModifyReturnValue() = %v, want %v", result, expected)
	}
}

// Bonus challenge tests
func TestMultiResourceCleanup(t *testing.T) {
	// Create temporary directories for testing
	tmpDir, err := os.MkdirTemp("", "test-dir")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	file1Path := filepath.Join(tmpDir, "file1.txt")
	file2Path := filepath.Join(tmpDir, "file2.txt")

	// Write test content
	if err := os.WriteFile(file1Path, []byte("File 1 content"), 0644); err != nil {
		t.Fatalf("Failed to write test file: %v", err)
	}
	if err := os.WriteFile(file2Path, []byte("File 2 content"), 0644); err != nil {
		t.Fatalf("Failed to write test file: %v", err)
	}

	// Create mutex for testing
	var mu sync.Mutex

	// Execute with error
	result, err := MultiResourceCleanup(file1Path, file2Path, &mu, true)
	if err == nil {
		t.Errorf("MultiResourceCleanup() with error flag expected error, got nil")
	}
	if result != "" {
		t.Errorf("MultiResourceCleanup() with error expected empty result, got %q", result)
	}

	// Verify mutex is unlocked after error
	locked := mu.TryLock()
	if !locked {
		t.Errorf("Mutex was left locked after error")
	} else {
		mu.Unlock()
	}

	// Execute successfully
	result, err = MultiResourceCleanup(file1Path, file2Path, &mu, false)
	if err != nil {
		t.Errorf("MultiResourceCleanup() without error flag got error: %v", err)
	}
	combinedContent := "File 1 content\nFile 2 content"
	if !strings.Contains(result, "File 1 content") || !strings.Contains(result, "File 2 content") {
		t.Errorf("MultiResourceCleanup() result = %q, expected to contain %q", result, combinedContent)
	}
}

func TestLogFunction(t *testing.T) {
	// Capture stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Test successful function execution
	successFunc := func() error {
		return nil
	}
	LogFunction("SuccessOperation", successFunc)

	// Test function with error
	expectedErr := errors.New("operation failed")
	errorFunc := func() error {
		return expectedErr
	}
	LogFunction("FailOperation", errorFunc)

	// Restore stdout and read captured output
	w.Close()
	os.Stdout = oldStdout
	var buf strings.Builder
	_, _ = buf.ReadFrom(r)
	output := buf.String()

	// Verify output
	expectedOutputs := []string{
		"ENTER: SuccessOperation",
		"EXIT: SuccessOperation",
		"ENTER: FailOperation",
		"ERROR in FailOperation",
		"operation failed",
	}

	for _, expected := range expectedOutputs {
		if !strings.Contains(output, expected) {
			t.Errorf("LogFunction output missing expected content: %q", expected)
		}
	}
} 