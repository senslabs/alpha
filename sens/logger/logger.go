package logger

import (
	"log"
	"os"
	"runtime"
)

type Logger interface {
	Debug(v ...interface{})
	Error(v ...interface{})
	Debugf(format string, v ...interface{})
	Errorf(format string, v ...interface{})
}

type ConsoleLogger struct{}
type FileLogger struct{}
type FluentLogger struct{}

var logger Logger

func InitConsoleLogger() {
	logger = &ConsoleLogger{}
}

func InitFileLogger(name string) {
	err := os.MkdirAll("/var/sens/logs", 0666)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.OpenFile("/var/sens/logs/"+name+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	log.SetOutput(f)
	log.SetFlags(log.Lshortfile)
	logger = &FileLogger{}
}

func InitFluentLogger() {
	logger = &FluentLogger{}
}

// -- INIT ENDS HERE -- //

func (this *ConsoleLogger) Error(v ...interface{}) {
	log.Println(v...)
}

func (this *ConsoleLogger) Debug(v ...interface{}) {
	log.Println(v...)
}

func (this *ConsoleLogger) Errorf(format string, v ...interface{}) {
	log.Printf(format, v...)
}

func (this *ConsoleLogger) Debugf(format string, v ...interface{}) {
	log.Printf(format, v...)
}

func (this *FileLogger) Error(v ...interface{}) {
	log.Println(v...)
}

func (this *FileLogger) Debug(v ...interface{}) {
	log.Println(v...)
}

func (this *FileLogger) Errorf(format string, v ...interface{}) {
	log.Printf(format, v...)
}

func (this *FileLogger) Debugf(format string, v ...interface{}) {
	log.Printf(format, v...)
}

func (this *FluentLogger) Error(v ...interface{}) {
}

func (this *FluentLogger) Debug(v ...interface{}) {
}

func (this *FluentLogger) Errorf(format string, v ...interface{}) {
}

func (this *FluentLogger) Debugf(format string, v ...interface{}) {
}

func LogMeta(level string) {
	pc, file, line, ok := runtime.Caller(2)
	if ok {
		log.Printf("%s: %s:%d in func: %v", level, file, line, runtime.FuncForPC(pc).Name())
	}
}

func InitLogger(arg interface{}) {
	logStore := os.Getenv("LOG_STORE")
	if logStore == "file" {
		InitFileLogger(arg.(string))
	} else if logStore == "fluentd" {
		//Prabhu
	} else {
		InitConsoleLogger()
	}
}

func Error(v ...interface{}) {
	LogMeta("Error")
	logger.Error(v...)
}

func Debug(v ...interface{}) {
	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "DEBUG" {
		LogMeta("Debug")
		logger.Debug(v...)
	}
}

func Errorf(format string, v ...interface{}) {
	LogMeta("Error")
	logger.Errorf(format, v...)
}

func Debugf(format string, v ...interface{}) {
	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "DEBUG" {
		LogMeta("Debug")
		logger.Debugf(format, v...)
	}
}
