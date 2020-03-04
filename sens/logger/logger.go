package logger

import (
	"log"
	"os"
)

type logger struct{}

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
}

func InitLogger() {
}

func Error(v ...interface{}) {
	log.Println(v...)
}

func Debug(v ...interface{}) {
	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "DEBUG" {
		log.Println(v...)
	}
}
