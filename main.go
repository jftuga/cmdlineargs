/*
main.go
-John Taylor
2021-05-13

Output all given command line arguments to STDOUT unless the CMDLINEARGS_FILE environment variable is set.
In that case, append the command line arguments to the file stored in this value.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

const version string = "1.1.0"

func main() {
	var f *os.File
	var err error
	saveToFile := os.Getenv("CMDLINEARGS_FILE")
	if len(saveToFile) > 0 {
		f, err = os.OpenFile(saveToFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening file: %s\n", saveToFile)
			os.Exit(1)
		}
	} else {
		f = os.Stdout
	}
	writer := bufio.NewWriter(f)
	writer.WriteString("\n========================================================================\n")
	dt := time.Now()
	writer.WriteString(fmt.Sprintf("%s\n", dt.Format(time.RFC3339)))
	//writer.WriteString(fmt.Sprintf("length: %d\n\n", len(os.Args)))
	for i, arg := range os.Args {
		writer.WriteString(fmt.Sprintf("[%2d] %s\n", i, arg))
	}
	writer.Flush()
	f.Close()
}
