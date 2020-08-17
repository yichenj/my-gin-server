package applog

import (
	"io/ioutil"
	"log"
	"os"
)

type Level int

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
)

var (
	Debug   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func SetupLoggers(logLevel string) {
	debugOutput := ioutil.Discard
	infoOutput := ioutil.Discard
	warningOutput := ioutil.Discard
	errorOutput := ioutil.Discard

	debugFlags := 0
	infoFlags := 0
	warningFlags := 0
	errorFlags := 0

	switch logLevel {
	case "DEBUG":
		debugOutput = os.Stdout
		debugFlags = log.Ldate | log.Ltime | log.Lshortfile
		fallthrough
	case "INFO":
		infoOutput = os.Stdout
		infoFlags = log.Ldate | log.Ltime | log.Lshortfile
		fallthrough
	case "WARNING":
		warningOutput = os.Stdout
		warningFlags = log.Ldate | log.Ltime | log.Lshortfile
		fallthrough
	case "ERROR":
		errorOutput = os.Stderr
		errorFlags = log.Ldate | log.Ltime | log.Lshortfile
	}

	Debug = log.New(debugOutput, "[DEBUG] ", debugFlags)
	Info = log.New(infoOutput, "[INFO] ", infoFlags)
	Warning = log.New(warningOutput, "[WARNING] ", warningFlags)
	Error = log.New(errorOutput, "[ERROR] ", errorFlags)
}
