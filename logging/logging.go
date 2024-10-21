package main

import (
	"log"
	"os"
)

type LogLevel int

const (
	ErrorLogLevel   LogLevel = 0
	WarningLogLevel LogLevel = 1
	InfoLogLevel    LogLevel = 2
	DebugLogLevel   LogLevel = 3
)

type (
	ILogger interface {
		Debug(string, any)
		Info(string, any)
		Error(string, any)
		Warning(string, any)
	}
	logger struct {
		logLevel       LogLevel
		internalLogger *log.Logger
	}
)

func (l *logger) Debug(message string, args any) {
	if l.logLevel >= DebugLogLevel {
		l.internalLogger.SetPrefix("DEBUG")
		l.internalLogger.Printf(message, args)
	}
}

func (l *logger) Info(message string, args any) {
	if l.logLevel >= InfoLogLevel {
		l.internalLogger.SetPrefix("INFO")
		l.internalLogger.Printf(message, args)
	}
}
func (l *logger) Warning(message string, args any) {
	if l.logLevel >= WarningLogLevel {
		l.internalLogger.SetPrefix("WARNING")
		l.internalLogger.Printf(message, args)
	}
}
func (l *logger) Error(message string, args any) {
	if l.logLevel >= ErrorLogLevel {
		l.internalLogger.SetPrefix("ERROR")
		l.internalLogger.Printf(message, args)
	}
}

func New() ILogger {
	return NewWithLogLevel(InfoLogLevel)
}

func NewWithLogLevel(level LogLevel) ILogger {
	return &logger{
		logLevel:       level,
		internalLogger: log.New(os.Stdout, "DEBUG", log.Flags()),
	}
}
