package infrastructure

import (
	"fmt"
)

// ConsoleLogger simply logs to the console
type ConsoleLogger int

// Log logs the given msg
func (l ConsoleLogger) Log(msg string) {
	fmt.Println(msg)
}
