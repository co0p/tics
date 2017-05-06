package infrastructure

import (
	"golang.org/x/net/context"

	"google.golang.org/appengine/log"
)

// GaeLogger simply logs to the console
type GaeLogger struct {
	Ctx context.Context
}

// Log logs the given msg
func (l GaeLogger) Log(format string, args ...interface{}) {
	log.Infof(l.Ctx, format, args)
}
