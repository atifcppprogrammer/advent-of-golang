package utilities

import (
	"fmt"
	"os"
	"testing"
)

func TestFileExists(t *testing.T) {
	t.Run("returns nil, error for when file does'nt exist", func(t *testing.T) {
		received, _ := FileExists("boogie-woogie.txt")
		expected := false
		if received != expected {
			t.Errorf("Expected %t but received %t", expected, received)
		}
	})
}

func TestReadLines(t *testing.T) {
	filePath := "test-temp.txt"
	file, _ := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	t.Run("reads all lines from file as []string", func(t *testing.T) {
		expected := []string{
			"line-0",
			"line-1",
			"line-2",
			"line-4",
		}

		for _, line := range expected {
			file.Write([]byte(fmt.Sprint(line, "\n")))
		}

		received, _ := ReadLines(filePath)
		if len(expected) != len(received) {
			t.Errorf(
				"Expected []string %q but received %q",
				expected,
				received,
			)
		}

		for index := range received {
			if received[index] != expected[index] {
				t.Errorf(
					"Expected string %q but received %q",
					expected[index],
					received[index],
				)
			}
		}
	})

	t.Cleanup(func() {
		file.Close()
		os.Remove(filePath)
	})
}
