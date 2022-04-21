package LogToFile

import (
    "fmt"
    "log"
    "os"
)

var (
    entries     []LogEntry
    fileName    string
    printOn     bool
)

type LogEntry struct {
    Level   string
    Message string
}

// Initialize LogTofile entries list and parameters
func Init(file string, print bool){
    entries     = []LogEntry{}
    fileName    = file
    printOn     = print
}

// Add a LogEntry to the array
func Add(lvl, msg string){
   entries = append(entries, LogEntry{lvl, msg})
}

func WriteToFile() {
    // open file and create if non-existent
    file, err := os.OpenFile(fileName + ".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    for _, entry := range entries {
        logger := log.New(file, "[" + entry.Level + "] ", log.LstdFlags)
        logger.Println(entry.Message)
        // Print if enabled
        if printOn { fmt.Println(entry.Level + " : " + entry.Message) }
    }
    // clear entries after writing them
    clear()
}

func clear() {
    entries = []LogEntry{}
}
