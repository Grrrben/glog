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

func Info(line string) {
	f := getLog(prefix_info)
	defer f.Close()
	log.Println(line)
}

func Infof(line string, vars ...interface{}) {
	f := getLog(prefix_info)
	defer f.Close()
	log.Printf(line, vars...)
}

func Error(line string) {
	f := getLog(prefix_error)
	defer f.Close()
	log.Println(line)
}

func Errorf(line string, vars ...interface{}) {
	f := getLog(prefix_error)
	defer f.Close()
	log.Printf(line, vars...)
}

func Fatal(line string) {
	f := getLog(prefix_fatal)
	defer f.Close()
	log.Println(line)
}

func Fatalf(line string, vars ...interface{}) {
	f := getLog(prefix_fatal)
	defer f.Close()
	log.Printf(line, vars...)
}

func Warning(line string) {
	f := getLog(prefix_warning)
	defer f.Close()
	log.Println(line)
}

func Warningf(line string, vars ...interface{}) {
	f := getLog(prefix_warning)
	defer f.Close()
	log.Printf(line, vars...)
}
