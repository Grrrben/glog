# glog
A simple log tool for usage in Golang applications.

[![Build Status](https://travis-ci.org/Grrrben/glog.svg?branch=master)](https://travis-ci.org/Grrrben/glog)

Just a wrapper on the standard go log package. It includes information on the serverity of the message:

I(nfo)  
W(arning)  
E(rror)
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
```

The default behaviour is Log_level)_info, where every line is logged.

## Todo

- Implement a log stack with a Flush() method
