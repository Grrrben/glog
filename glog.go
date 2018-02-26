package glog

import (
	"log"
	"os"
)

const (
	prefix_info    = "I "
	prefix_warning = "W "
	prefix_error   = "E "
	prefix_fatal   = "F "
)

// this is where your log would be placed.
var logfile string = "glog.log"

// SetLogFile sets a preferred path/and/file to the log.
func SetLogFile(lf string) {
	logfile = lf
}

func getLog (prefix string) (*os.File) {
	f, err := os.OpenFile(logfile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	log.SetOutput(f)
	log.SetPrefix(prefix)
	return f
}

// Info prints a line to the log beginning with the I char to clarify that it's an Info message
func Info(msg string) {
	f := getLog(prefix_info)
	defer f.Close()
	log.Println(msg)
}

// Infof is identical to Info() but let's the user format the message string with any number of var's.
func Infof(msg string, vars ...interface{}) {
	f := getLog(prefix_info)
	defer f.Close()
	log.Printf(msg, vars...)
}

// Error prints a line to the log beginning with the E char to clarify that it's an Error message
func Error(msg string) {
	f := getLog(prefix_error)
	defer f.Close()
	log.Println(msg)
}

// Errorf is identical to Error() but let's the user format the message string with any number of var's.
func Errorf(msg string, vars ...interface{}) {
	f := getLog(prefix_error)
	defer f.Close()
	log.Printf(msg, vars...)
}

// Fatal prints a line to the log beginning with the F char to clarify that it's an Fatal message
// Fatal will terminate the program after logging.
func Fatal(msg string) {
	f := getLog(prefix_fatal)
	defer f.Close()
	log.Println(msg)
}

// Fatalf is identical to Fatal() but let's the user format the message string with any number of var's.
func Fatalf(msg string, vars ...interface{}) {
	f := getLog(prefix_fatal)
	defer f.Close()
	log.Printf(msg, vars...)
}

// Warning prints a line to the log beginning with the W char to clarify that it's an Warning message
func Warning(msg string) {
	f := getLog(prefix_warning)
	defer f.Close()
	log.Println(msg)
}

// Warningf is identical to Warning() but let's the user format the message string with any number of var's.
func Warningf(msg string, vars ...interface{}) {
	f := getLog(prefix_warning)
	defer f.Close()
	log.Printf(msg, vars...)
}

// Flush is not yet implemented, needed for the API's interface sake
func Flush() {}
