package glog

import "testing"

func TestSetLogFile(t *testing.T) {
	if logfile != "glog.log" {
		t.Error("Base logfile is incorrect")
	}

	SetLogFile("dir/test.log")

	if logfile != "dir/test.log" {
		t.Error("Could not change the location of the logfile")
	}
}

func TestSetLogLevel(t *testing.T) {
	if logLevel != Log_level_info {
		t.Errorf("Loglevel unexpected: want %d got %d", Log_level_info, logLevel)
	}

	SetLogLevel(Log_level_warning)
	if logLevel != Log_level_warning {
		t.Errorf("Loglevel unexpected: want %d got %d", Log_level_warning, logLevel)
	}

	SetLogLevel(Log_level_error)
	if logLevel != Log_level_error {
		t.Errorf("Loglevel unexpected: want %d got %d", Log_level_error, logLevel)
	}

	SetLogLevel(5)
	if logLevel != Log_level_info {
		t.Errorf("Loglevel unexpected: want %d got %d", Log_level_info, logLevel)
	}
}
