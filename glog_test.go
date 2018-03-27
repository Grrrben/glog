package glog

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"testing"
)

const testlog = "test.log"

func TestMain(m *testing.M) {
	runTests := m.Run()
	teardown()
	os.Exit(runTests)
}

func teardown() {
	var err = os.Remove(testlog)
	if err != nil {
		log.Fatal("Could not remove the logfile")
	}
}

func reset() {
	// delete file
	var err = os.Remove(testlog)
	if err != nil {
		log.Fatal("Could not remove the logfile")
	}
	// reset to test.log
	SetLogFile(testlog)
}

func TestSetLogFile(t *testing.T) {
	if logfile != "glog.log" {
		t.Error("Base logfile is incorrect")
	}

	SetLogFile("dir/test.log")

	if logfile != "dir/test.log" {
		t.Error("Could not change the location of the logfile")
	}

	// reset
	SetLogFile(testlog)
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

func TestGetLog(t *testing.T) {
	SetLogFile(testlog)
	getLog(prefix_info)

	if _, err := os.Stat(testlog); os.IsNotExist(err) {
		t.Error("Logfile test.log does not exists")
	}
}

func TestInfo(t *testing.T) {
	reset()
	Info("info message")
	file, err := os.Open(testlog)
	if err != nil {
		t.Error("Could not open logfile test.log")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		bytes := []byte(line)
		match, err := regexp.Match(`^I\s[0-9]{4}/[0-9]{2}/[0-9]{2}\s[0-9]{2}:[0-9]{2}:[0-9]{2}\sinfo\smessage$`, bytes)

		if err != nil {
			t.Errorf("Error in regxep: %s", err.Error())
		}

		if !match {
			t.Error("Info message not found")
		}
	}
}

func TestInfof(t *testing.T) {
	reset()
	Infof("info message %s %d", "number", 2)
	file, err := os.Open(testlog)
	if err != nil {
		t.Error("Could not open logfile test.log")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		bytes := []byte(line)
		pattern := `^I\s[0-9]{4}/[0-9]{2}/[0-9]{2}\s[0-9]{2}:[0-9]{2}:[0-9]{2}\sinfo\smessage\snumber\s2$`
		match, err := regexp.Match(pattern, bytes)

		if err != nil {
			t.Errorf("Error in regxep: %s", err.Error())
		}

		if !match {
			t.Errorf("Infof message not found in line %s", line)
		}
	}
}

func TestWarning(t *testing.T) {
	reset()
	Warning("warning message")
	file, err := os.Open(testlog)
	if err != nil {
		t.Error("Could not open logfile test.log")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		bytes := []byte(line)
		match, err := regexp.Match(`^W\s[0-9]{4}/[0-9]{2}/[0-9]{2}\s[0-9]{2}:[0-9]{2}:[0-9]{2}\swarning\smessage$`, bytes)

		if err != nil {
			t.Errorf("Error in regxep: %s", err.Error())
		}

		if !match {
			t.Error("Warning message not found")
		}
	}
}

func TestWarningf(t *testing.T) {
	reset()
	Warningf("warning message %s %d", "number", 2)
	file, err := os.Open(testlog)
	if err != nil {
		t.Error("Could not open logfile test.log")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		bytes := []byte(line)
		pattern := `^W\s[0-9]{4}/[0-9]{2}/[0-9]{2}\s[0-9]{2}:[0-9]{2}:[0-9]{2}\swarning\smessage\snumber\s2$`
		match, err := regexp.Match(pattern, bytes)

		if err != nil {
			t.Errorf("Error in regxep: %s", err.Error())
		}

		if !match {
			t.Errorf("Warningf message not found in line %s", line)
		}
	}
}

func TestError(t *testing.T) {
	reset()
	Error("error message")
	file, err := os.Open(testlog)
	if err != nil {
		t.Error("Could not open logfile test.log")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		bytes := []byte(line)
		match, err := regexp.Match(`^E\s[0-9]{4}/[0-9]{2}/[0-9]{2}\s[0-9]{2}:[0-9]{2}:[0-9]{2}\serror\smessage$`, bytes)

		if err != nil {
			t.Errorf("Error in regxep: %s", err.Error())
		}

		if !match {
			t.Error("Error message not found")
		}
	}
}

func TestErrorf(t *testing.T) {
	reset()
	Errorf("error message %s %d", "number", 2)
	file, err := os.Open(testlog)
	if err != nil {
		t.Error("Could not open logfile test.log")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		bytes := []byte(line)
		pattern := `^E\s[0-9]{4}/[0-9]{2}/[0-9]{2}\s[0-9]{2}:[0-9]{2}:[0-9]{2}\serror\smessage\snumber\s2$`
		match, err := regexp.Match(pattern, bytes)

		if err != nil {
			t.Errorf("Error in regxep: %s", err.Error())
		}

		if !match {
			t.Errorf("Errorf message not found in line %s", line)
		}
	}
}

func TestPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			if r != "panic message" {
				t.Errorf("Error in 'panic message': '%s'", r)
			}

			file, err := os.Open(testlog)
			if err != nil {
				t.Error("Could not open logfile test.log")
			}

			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				line := scanner.Text()
				bytes := []byte(line)
				match, err := regexp.Match(`^P\s[0-9]{4}/[0-9]{2}/[0-9]{2}\s[0-9]{2}:[0-9]{2}:[0-9]{2}\spanic\smessage$`, bytes)

				if err != nil {
					t.Errorf("Error in regxep: %s", err.Error())
				}

				if !match {
					t.Error("Panic message not found")
				}
			}
			file.Close()
		}
	}()
	reset()
	Panic("panic message")
}

func TestPanicf(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			if r != "panic message 2" {
				t.Errorf("Error in 'panic message': '%s'", r)
			}

			file, err := os.Open(testlog)
			if err != nil {
				t.Error("Could not open logfile test.log")
			}

			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				line := scanner.Text()
				bytes := []byte(line)
				match, err := regexp.Match(`^P\s[0-9]{4}/[0-9]{2}/[0-9]{2}\s[0-9]{2}:[0-9]{2}:[0-9]{2}\spanic\smessage\s2$`, bytes)

				if err != nil {
					t.Errorf("Error in regxep: %s", err.Error())
				}

				if !match {
					t.Error("Panic message not found")
				}
			}
			file.Close()
		}
	}()
	reset()
	Panicf("panic message %d", 2)
}

func TestMultipleLines(t *testing.T) {
	reset()
	addSomeLines()

	file, err := os.Open(testlog)
	if err != nil {
		t.Error("Could not open logfile test.log")
	}
	defer file.Close()

	i := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i++
	}

	if i != 6 {
		t.Errorf("Error, expected 6 loglines, got %d", i)
	}
}

func TestMultipleLinesWithInfoDisabled(t *testing.T) {
	reset()
	SetLogLevel(Log_level_warning)
	addSomeLines()

	file, err := os.Open(testlog)
	if err != nil {
		t.Error("Could not open logfile test.log")
	}
	defer file.Close()

	i := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i++
	}

	if i != 4 {
		t.Errorf("Error, expected 4 loglines, got %d", i)
	}
}

func TestMultipleLinesWithWarningDisabled(t *testing.T) {
	reset()
	SetLogLevel(Log_level_error)
	addSomeLines()

	file, err := os.Open(testlog)
	if err != nil {
		t.Error("Could not open logfile test.log")
	}
	defer file.Close()

	i := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i++
	}

	if i != 2 {
		t.Errorf("Error, expected 2 loglines, got %d", i)
	}
}

// addSomeLines just add's a log message for every available action
func addSomeLines() {
	Info("info message")
	Warning("warning message")
	Error("error message")

	Infof("info message %d", 2)
	Warningf("warning message %d", 2)
	Errorf("error message %d", 2)
}
