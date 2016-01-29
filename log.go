// +build !logrus

package gotopt

import (
	"log"
	"os"
)

var (
	debug = os.Getenv("GOTOPT_DEBUG") == "true"
)

// Debugln logs the arguments only if the logging level is DEBUG or higher.
func Debugln(args ...interface{}) {
	if !debug {
		return
	}
	log.Println(args)
}

// Debugf logs the arguments only if the logging level is DEBUG or higher.
func Debugf(format string, args ...interface{}) {
	if !debug {
		return
	}
	log.Printf(format, args...)
}
