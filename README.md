# go-log-to-file
Simple go module to print log entries to a file

#### Import module:
```go
import "github.com/cavebird/go-log-to-file"
```

#### Initialize module:
  - Parameters:
    1. filename (string)  : Absolute path to filename to print the results to. Will be created if not exists. ".log" is appended.
    2. print    (bool)    : If true prints each entry to stdout after writing it
```go
LogToFile.Init("/my/path/myFileName", true)
```

#### Add entry to log:
- Parameters:
  1. log level    (string)  : Is printed first in the line. 
  2. log message  (string)  : Is printed last in the line.

```go
LogToFile.Add("INFO", "test log message")
```

#### Write to file
```go
LogToFile.WriteToFile()
```

#### Example
```go
func main(){
    LogToFile.Init("myFile", true)
    LogToFile.Add("INFO", "test log message")
    defer LogToFile.WriteToFile()
}
```
Will write ``` [INFO] 2022/04/21 16:35:58 test log message ``` to ``` myFile.log``` in directory where called
