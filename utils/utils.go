package utils

import (
	"fmt"
)

const SYMMER_VERSION = "1.0.0"

// Returns the given argument as bold white
func WhiteText(v string) string {
	return "\033[1;99m" + v + "\033[0m"
}

// Returns the given argument as bold red
func RedText(v string) string {
	return "\033[1;91m" + v + "\033[0m"
}

// Returns the given argument as bold green
func GreenText(v string) string {
	return "\033[1;92m" + v + "\033[0m"
}

// Prints the given value with symmer: prefixed on the value.
func SymmerPrint(v any) {
	fmt.Printf("%v: %v\n", WhiteText("symmer"), v)
}

// Prints the given value with error: prefixed on the value.
// This command should only be used to convey an error.
func SymmerError(v any) {
	fmt.Printf("%v: %v\n", RedText("symmer"), v)
}
