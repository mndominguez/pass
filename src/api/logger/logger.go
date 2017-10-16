package logger

import (
	log "github.com/Sirupsen/logrus"
	"os"
	"runtime/debug"
)

var logger *log.Logger

func init() {
	log.SetFormatter(&log.TextFormatter{})

	level := os.Getenv("logger_level")

	switch level {
	case "INFO":
		log.SetLevel(log.InfoLevel)
	case "DEBUG":
		log.SetLevel(log.DebugLevel)
	case "WARN":
		log.SetLevel(log.WarnLevel)
	}

}

func Info(msg string) {
	log.WithFields(log.Fields{}).Info(msg)
}

func Debug(msg string) {
	log.WithFields(log.Fields{}).Debug(msg)
}

func Error(msg string, err error) {
	msg = msg + " - ERROR: %v"
	log.WithFields(log.Fields{}).Errorf(msg, err)
}

func Panic(msg string, err error) {
	msg = msg + " - PANIC: %v"
	log.WithFields(log.Fields{}).Panicf(msg, err)
}

func Warn(msg string) {
	log.WithFields(log.Fields{}).Warn(msg)
}

func Print(e interface{}) {
	log.Printf("%s: %s", e, debug.Stack())
}
