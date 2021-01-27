package log

import (
	"fmt"
	"io"
	stdlog "log"
	"os"
)

// DebugLogger logs to Stderr by default with `Debug` prefix, can be override.
var DebugLogger = stdlog.New(os.Stderr, "Debug: ", stdlog.Ldate|stdlog.Ltime|stdlog.Lshortfile)

// InfoLogger logs to Stdout by default, can be override.
var InfoLogger = stdlog.New(os.Stdout, "", stdlog.Ldate|stdlog.Ltime)

// Println calls Output to print to the DebugLogger.
// Arguments are handled in the manner of fmt.Println.
func Println(v ...interface{}) {
	DebugLogger.Output(2, fmt.Sprintln(v...))
}

// Printf calls Output to print to the DebugLogger.
// Arguments are handled in the manner of fmt.Printf.
func Printf(format string, v ...interface{}) {
	DebugLogger.Output(2, fmt.Sprintf(format, v...))
}

// Info calls Output to print to the InfoLogger.
// Arguments are handled in the manner of fmt.Printf.
func Info(v ...interface{}) {
	InfoLogger.Output(2, fmt.Sprintln(v...))
}

// Copied from https://golang.org/src/log/log.go

// SetOutput sets the output destination for the standard logger.
func SetOutput(w io.Writer) {
	DebugLogger.SetOutput(w)
}

// Flags returns the output flags for the standard logger.
// The flag bits are Ldate, Ltime, and so on.
func Flags() int {
	return DebugLogger.Flags()
}

// SetFlags sets the output flags for the standard logger.
// The flag bits are Ldate, Ltime, and so on.
func SetFlags(flag int) {
	DebugLogger.SetFlags(flag)
}

// Prefix returns the output prefix for the standard logger.
func Prefix() string {
	return DebugLogger.Prefix()
}

// SetPrefix sets the output prefix for the standard logger.
func SetPrefix(prefix string) {
	DebugLogger.SetPrefix(prefix)
}

// Writer returns the output destination for the standard logger.
func Writer() io.Writer {
	return DebugLogger.Writer()
}

// These functions write to the standard logger.

// Print calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Print.
func Print(v ...interface{}) {
	DebugLogger.Output(2, fmt.Sprint(v...))
}

// Fatal is equivalent to Print() followed by a call to os.Exit(1).
func Fatal(v ...interface{}) {
	DebugLogger.Output(2, fmt.Sprint(v...))
	os.Exit(1)
}

// Fatalf is equivalent to Printf() followed by a call to os.Exit(1).
func Fatalf(format string, v ...interface{}) {
	DebugLogger.Output(2, fmt.Sprintf(format, v...))
	os.Exit(1)
}

// Fatalln is equivalent to Println() followed by a call to os.Exit(1).
func Fatalln(v ...interface{}) {
	DebugLogger.Output(2, fmt.Sprintln(v...))
	os.Exit(1)
}

// Panic is equivalent to Print() followed by a call to panic().
func Panic(v ...interface{}) {
	s := fmt.Sprint(v...)
	DebugLogger.Output(2, s)
	panic(s)
}

// Panicf is equivalent to Printf() followed by a call to panic().
func Panicf(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	DebugLogger.Output(2, s)
	panic(s)
}

// Panicln is equivalent to Println() followed by a call to panic().
func Panicln(v ...interface{}) {
	s := fmt.Sprintln(v...)
	DebugLogger.Output(2, s)
	panic(s)
}

// Output writes the output for a logging event. The string s contains
// the text to print after the prefix specified by the flags of the
// Logger. A newline is appended if the last character of s is not
// already a newline. Calldepth is the count of the number of
// frames to skip when computing the file name and line number
// if Llongfile or Lshortfile is set; a value of 1 will print the details
// for the caller of Output.
func Output(calldepth int, s string) error {
	return DebugLogger.Output(calldepth+1, s) // +1 for this frame.
}
