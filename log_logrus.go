// +build logrus

package gotopt

import (
	log "github.com/Sirupsen/logrus"
	"os"
)

func init() {
	if os.Getenv("GOTOPT_DEBUG") == "true" {
		log.SetLevel(log.DebugLevel)
	}
}

// Debugln logs the arguments only if the logging level is DEBUG or higher.
func Debugln(args ...interface{}) {
	log.Debugln(args...)
}

// Debugf logs the arguments only if the logging level is DEBUG or higher.
func Debugf(format string, args ...interface{}) {
	log.Debugf(format, args...)
}
