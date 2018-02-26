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
func Info(line string) {
	f := getLog(prefix_info)
	defer f.Close()
	log.Println(line)
}

// Infof is identical to Info() but let's the user format the message string with any number of var's.
func Infof(line string, vars ...interface{}) {
	f := getLog(prefix_info)
	defer f.Close()
	log.Printf(line, vars...)
}

// Error prints a line to the log beginning with the E char to clarify that it's an Error message
func Error(line string) {
	f := getLog(prefix_error)
	defer f.Close()
	log.Println(line)
}

// Errorf is identical to Error() but let's the user format the message string with any number of var's.
func Errorf(line string, vars ...interface{}) {
	f := getLog(prefix_error)
	defer f.Close()
	log.Printf(line, vars...)
}

// Fatal prints a line to the log beginning with the F char to clarify that it's an Fatal message
// Fatal will terminate the program after logging.
func Fatal(line string) {
	f := getLog(prefix_fatal)
	defer f.Close()
	log.Println(line)
}

// Fatalf is identical to Fatal() but let's the user format the message string with any number of var's.
func Fatalf(line string, vars ...interface{}) {
	f := getLog(prefix_fatal)
	defer f.Close()
	log.Printf(line, vars...)
}

// Warning prints a line to the log beginning with the W char to clarify that it's an Warning message
func Warning(line string) {
	f := getLog(prefix_warning)
	defer f.Close()
	log.Println(line)
}

// Warningf is identical to Warning() but let's the user format the message string with any number of var's.
func Warningf(line string, vars ...interface{}) {
	f := getLog(prefix_warning)
	defer f.Close()
	log.Printf(line, vars...)
}
