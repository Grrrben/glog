# glog
A simple log tool for usage in Golang applications.  

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

## Todo

- Set severity level
- Write some test
- Document the methods
