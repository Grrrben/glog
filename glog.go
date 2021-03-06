package glog

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

// Exported constants for the severity level of the logged messages.
// Use these when calling Setlogger.logLevel(int)
const (
	Log_level_error = iota
	Log_level_warning
	Log_level_info
)

// Prefixes for the log lines to determine the severity of the message
const (
	prefix_info    = "I "
	prefix_warning = "W "
	prefix_error   = "E "
	prefix_fatal   = "F "
	prefix_panic   = "P "
)

var logger struct {
	logLevel     int    // Should everything be logged, or should Info and Warning be dismissed?
	logfile      string // this is where your log would be placed.
	sync.RWMutex        // for usage in goroutines
}

func init() {
	logger.logLevel = Log_level_info
	logger.logfile = "glog.log"
}

// SetLogLevel tells glog to log or dismiss Warning and Info messages
// To be used with the Log_level_* constants
func SetLogLevel(level int) {
	if level < Log_level_error || level > Log_level_info {
		// impossible value, should be consistent with the severity levels
		logger.logLevel = Log_level_info // include everything...
	} else {
		logger.logLevel = level
	}
}

// SetLogfile sets a preferred path/and/file to the log.
func SetLogFile(lf string) {
	logger.logfile = lf
}

// getLog gives the caller a pointer to the log File so message strings can be added to the file
// Should be closed after use.
func getLog(prefix string) *os.File {
	f, err := os.OpenFile(logger.logfile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	log.SetOutput(f)
	log.SetPrefix(prefix)
	return f
}

// SetOutput sets the output destination for the logger.
func SetOutput(w io.Writer) {
	logger.Lock()
	defer logger.Unlock()
	log.SetOutput(w)
}

// Info prints a line to the log beginning with the I char to clarify that it's an Info message
func Info(msg ...interface{}) {
	logger.Lock()
	defer logger.Unlock()
	if logger.logLevel >= Log_level_info {
		f := getLog(prefix_info)
		defer f.Close()
		for _, m := range msg {
			log.Println(m)
		}
	}
}

// Infof is identical to Info() but let's the user format the message string with any number of var's.
func Infof(msg string, vars ...interface{}) {
	logger.Lock()
	defer logger.Unlock()
	if logger.logLevel >= Log_level_info {
		f := getLog(prefix_info)
		defer f.Close()
		log.Printf(msg, vars...)
	}
}

// Error prints a line to the log beginning with the E char to clarify that it's an Error message
func Error(msg ...interface{}) {
	logger.Lock()
	defer logger.Unlock()
	f := getLog(prefix_error)
	defer f.Close()
	for _, m := range msg {
		log.Println(m)
	}
}

// Errorf is identical to Error() but let's the user format the message string with any number of var's.
func Errorf(msg string, vars ...interface{}) {
	logger.Lock()
	defer logger.Unlock()
	f := getLog(prefix_error)
	defer f.Close()
	log.Printf(msg, vars...)
}

// Fatal prints a line to the log beginning with the F char to clarify that it's an Fatal message
// Fatal will terminate the program after logging.
func Fatal(msg ...interface{}) {
	logger.Lock()
	defer logger.Unlock()
	f := getLog(prefix_fatal)
	defer f.Close()
	for _, m := range msg {
		log.Println(m)
	}
	os.Exit(1)
}

// Fatalf is identical to Fatal() but let's the user format the message string with any number of var's.
func Fatalf(msg string, vars ...interface{}) {
	logger.Lock()
	defer logger.Unlock()
	f := getLog(prefix_fatal)
	defer f.Close()
	log.Printf(msg, vars...)
	os.Exit(1)
}

// Panic prints a line to the log beginning with the P char to clarify that it's an Panic message
// Panic will panic after logging.
func Panic(msg ...interface{}) {
	var m interface{}
	logger.Lock()
	defer logger.Unlock()
	f := getLog(prefix_panic)
	defer f.Close()
	for _, m = range msg {
		log.Println(m)
	}
	panic(m)
}

// Panicf is identical to Panic() but let's the user format the message string with any number of var's.
func Panicf(msg string, vars ...interface{}) {
	logger.Lock()
	defer logger.Unlock()
	f := getLog(prefix_panic)
	defer f.Close()
	m := fmt.Sprintf(msg, vars...)
	log.Println(m)
	panic(m)
}

// Warning prints a line to the log beginning with the W char to clarify that it's an Warning message
func Warning(msg ...interface{}) {
	logger.Lock()
	defer logger.Unlock()
	if logger.logLevel >= Log_level_warning {
		f := getLog(prefix_warning)
		defer f.Close()
		for _, m := range msg {
			log.Println(m)
		}
	}
}

// Warningf is identical to Warning() but let's the user format the message string with any number of var's.
func Warningf(msg string, vars ...interface{}) {
	logger.Lock()
	defer logger.Unlock()
	if logger.logLevel >= Log_level_warning {
		f := getLog(prefix_warning)
		defer f.Close()
		log.Printf(msg, vars...)
	}
}

// Flush is not yet implemented, needed for the API's interface sake
func Flush() {}
