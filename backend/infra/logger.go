package infra

import (
	"io"
	"log"
	"os"

	"github.com/zyzmoz/DigitalMarket/usecases"
)

type Logger struct{}

func NewLogger() usecases.Logger {
	return &Logger{}
}

func (l *Logger) LogError(fmt string, v ...interface{}) {
	file, err := os.OpenFile("./logs/error.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		log.Printf("%s", err)
	}
	defer file.Close()

	log.SetOutput(io.MultiWriter(file, os.Stdout))
	log.SetFlags(log.Ldate | log.Ltime)

	log.Printf(fmt, v...)
}

func (l Logger) LogAccess(fmt string, v ...interface{}) {
	file, err := os.OpenFile("./logs/access.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Printf("%s", err)
	}
	defer file.Close()

	log.SetOutput(io.MultiWriter(file, os.Stdout))
	log.SetFlags(log.Ldate | log.Ltime)

	log.Printf(fmt, v...)
}
