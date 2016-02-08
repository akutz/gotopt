// +build !logrus

package gotopt

import (
	"log"
	"os"
)

var (
	debug = os.Getenv("GOTOPT_DEBUG") == "true"
)

// debugln logs the arguments only if the logging level is DEBUG or higher.
func debugln(args ...interface{}) {
	if !debug {
		return
	}
	log.Println(args...)
}

// debugf logs the arguments only if the logging level is DEBUG or higher.
func debugf(format string, args ...interface{}) {
	if !debug {
		return
	}
	log.Printf(format, args...)
}
