package infrastructure

import (
	"fmt"
)

// ConsoleLogger simply logs to the console
type ConsoleLogger struct{}

// Log logs the given msg
func (l ConsoleLogger) Log(format string, args ...interface{}) {
	fmt.Printf(format, args)
}
