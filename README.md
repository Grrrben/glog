# glog
A simple log tool for usage in Golang applications.

[![Build Status](https://travis-ci.org/Grrrben/glog.svg?branch=master)](https://travis-ci.org/Grrrben/glog)

Just a wrapper on the standard go log package. It includes information on the serverity of the message:

I(nfo)  
W(arning)  
E(rror)
P(anic)
F(atal)

## Usage

Just import it:
```
import (
	"github.com/grrrben/glog"
)
```

Don't forget that `go get` command...  
And then use it like you would use any sane logging tool:

```
glog.SetLogFile("log/glogfile.log")

glog.Warning("warning")
glog.Error("this is an Error")
glog.Warningf("warning: %s %d", "string", 8)
```

Output would be something like:

```
W 2018/02/25 13:42:50 warning
E 2018/02/25 13:42:50 this is an Error
W 2018/02/25 13:42:50 warning: string 8
```

__Setting the log file__
`SetLogFile()` is used to set a `path/plus/file.log` that the logger will use. If `glog.SetLogFile(path string)` is not set, the log wil appear as `glog.log` in your projects root directory.

__Muting Info and Warning messages__
Warning and Info messages can be muted by setting the log level with `glog.SetLogLevel(int)`.  
The param given to the method should be one of the `Log_level_*` constants:

```
// Exported constants for the severity level of the logged messages.
// Use these when calling SetLogLevel(int)
const (
	Log_level_error   = 0
	Log_level_warning = 1
	Log_level_info    = 2
)

// turn of info and warning:
glog.SetLogLevel(glog.Log_level_error)

```

The default behaviour is Log_level_info, where every line is logged.

__Concurrency__

`glog` is safe to use in concurrent programs. Just set the `go` keyword before calling a log message:

```
go glog.Warning("A concurrent warning")
go glog.Error("this is an Error in a goroutine")
go glog.Warningf("warning: %s %d", "goroutine", 8)
```

## Todo

- Implement a log stack with a Flush() method
