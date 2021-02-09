package internal

import (
	"fmt"
	"log"
)

// Errorf should be used to naively panics if an error is not nil.
func Errorf(err error, format string, args ...interface{}) {
	if err == nil {
		return
	}

	log.Printf("%s", fmt.Sprintf("%s. Error: %s", fmt.Sprintf(format, args...), err))
}

// Infof should be used to describe the example commands that are about to run.
func Infof(format string, args ...interface{}) {
	log.Printf("%s", fmt.Sprintf(format, args...))
}

// Warningf should be used to display a warning
func Warningf(format string, args ...interface{}) {
	log.Printf("%s", fmt.Sprintf(format, args...))
}
