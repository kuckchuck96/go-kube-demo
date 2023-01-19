package logger

import (
	"log"
	"os"
	"path"
)

type (
	Logger interface {
		Print(any)
	}

	logger struct {
		log *log.Logger
	}
)

func NewLogger(fileName string) (Logger, error) {
	lg := log.Default()

	if err := os.MkdirAll("logs", os.ModePerm); err != nil {
		return nil, err
	}

	f, err := os.Create(path.Join("logs", fileName))
	if err != nil {
		return nil, err
	}

	lg.SetOutput(f)
	return &logger{
		log: lg,
	}, nil
}

func (logger *logger) Print(v any) {
	logger.log.Println(v)
}
