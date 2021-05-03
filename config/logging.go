package config

import (
	"fmt"
	"log"
	"os"
)

// var (
// WarningFile string
// InfoFile    string
// ErrorFile   string

// WarningLogger *log.Logger
// InfoLogger    *log.Logger
// ErrorLogger   *log.Logger
// )
type LogType struct{}

var (
	Log *LogType
)

const (
	LogFile     = "logs.txt"
	InfoFile    = "infos.txt"
	WarningFile = "warnings.txt"
	ErrorFile   = "errors.txt"

	LogDir = "./storage/logs/"
)

func init() {
	Log = &LogType{}

	fmt.Println("logging package initialized")

	// file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	// WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	// ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func (l *LogType) Info(msg string) {

	file, err := os.OpenFile(LogDir+InfoFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	logger := log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println(msg)
}

func (l *LogType) Error(msg string) {

	file, err := os.OpenFile(LogDir+ErrorFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	logger := log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println(msg)
}

func (l *LogType) Warning(msg string) {

	file, err := os.OpenFile(LogDir+WarningFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	logger := log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println(msg)
}
