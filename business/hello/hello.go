package hello

import "fmt"

// Greet will take a name and create a greeting.
func Greet(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}
