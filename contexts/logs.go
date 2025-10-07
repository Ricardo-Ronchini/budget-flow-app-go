package contexts

import (
	"os"

	"github.com/sirupsen/logrus"
)

type Logs struct {
	Logger *logrus.Logger
}

func NewLogs() *Logs {
	logger := logrus.New()

	// Saída padrão no terminal
	logger.Out = os.Stdout

	// Formato dos logs
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		DisableColors: false,
	})

	// Definir nível no .env
	logger.SetLevel(logrus.DebugLevel)

	return &Logs{Logger: logger}
}

func (l *Logs) Info(v ...any) {
	l.Logger.Info(v...)
}

func (l *Logs) Error(v ...any) {
	l.Logger.Error(v...)
}

func (l *Logs) Warn(v ...any) {
	l.Logger.Warn(v...)
}

func (l *Logs) Debug(msg string, fields logrus.Fields) {
	l.Logger.WithFields(fields).Debug(msg)
}
