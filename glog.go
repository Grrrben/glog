package glog

import (
	"log"
	"os"
	"sync"
)

// Prefixes for the log lines to determine the severity of the message
const (
	prefix_info    = "I "
	prefix_warning = "W "
	prefix_error   = "E "
	prefix_fatal   = "F "
)

// Exported constants for the severity level of the logged messages.
// Use these when calling SetLogLevel(int)
const (
	Log_level_error   = 0
	Log_level_warning = 1
	Log_level_info    = 2
)

// Should everything be logged, or should Info and Warning be dismissed?
var logLevel int = Log_level_info

// this is where your log would be placed.
var logfile string = "glog.log"

// safe to use glog concurrently
var mux sync.Mutex

// SetLogLevel tells glog to log or dismiss Warning and Info messages
// To be used with the Log_level_* constants
func SetLogLevel(level int) {
	if level < 0 || level > 2 {
		// impossible value, should be consistent with the severity levels
		logLevel = Log_level_info // include everything...
	} else {
		logLevel = level
	}
}

// SetLogFile sets a preferred path/and/file to the log.
func SetLogFile(lf string) {
	logfile = lf
}

// getLog gives the caller a pointer to the log File so message strings can be added to the file
// Should be closed after use.
func getLog(prefix string) *os.File {
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
	mux.Lock()
	defer mux.Unlock()
	if logLevel >= Log_level_info {
		f := getLog(prefix_info)
		defer f.Close()
		log.Println(msg)
	}
}

// Infof is identical to Info() but let's the user format the message string with any number of var's.
func Infof(msg string, vars ...interface{}) {
	mux.Lock()
	defer mux.Unlock()
	if logLevel >= Log_level_info {
		f := getLog(prefix_info)
		defer f.Close()
		log.Printf(msg, vars...)
	}
}

// Error prints a line to the log beginning with the E char to clarify that it's an Error message
func Error(msg string) {
	mux.Lock()
	defer mux.Unlock()
	f := getLog(prefix_error)
	defer f.Close()
	log.Println(msg)
}

// Errorf is identical to Error() but let's the user format the message string with any number of var's.
func Errorf(msg string, vars ...interface{}) {
	mux.Lock()
	defer mux.Unlock()
	f := getLog(prefix_error)
	defer f.Close()
	log.Printf(msg, vars...)
}

// Fatal prints a line to the log beginning with the F char to clarify that it's an Fatal message
// Fatal will terminate the program after logging.
func Fatal(msg string) {
	mux.Lock()
	defer mux.Unlock()
	f := getLog(prefix_fatal)
	defer f.Close()
	log.Println(msg)
}

// Fatalf is identical to Fatal() but let's the user format the message string with any number of var's.
func Fatalf(msg string, vars ...interface{}) {
	mux.Lock()
	defer mux.Unlock()
	f := getLog(prefix_fatal)
	defer f.Close()
	log.Printf(msg, vars...)
}

// Warning prints a line to the log beginning with the W char to clarify that it's an Warning message
func Warning(msg string) {
	mux.Lock()
	defer mux.Unlock()
	if logLevel >= Log_level_warning {
		f := getLog(prefix_warning)
		defer f.Close()
		log.Println(msg)
	}
}

// Warningf is identical to Warning() but let's the user format the message string with any number of var's.
func Warningf(msg string, vars ...interface{}) {
	mux.Lock()
	defer mux.Unlock()
	if logLevel >= Log_level_warning {
		f := getLog(prefix_warning)
		defer f.Close()
		log.Printf(msg, vars...)
	}
}

// Flush is not yet implemented, needed for the API's interface sake
func Flush() {}
