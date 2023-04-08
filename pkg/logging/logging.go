package logging

import (
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

const (
	logDir  = "log"
	logFile = "all.log"
)

type writerHook struct {
	Writer    []io.Writer
	LogLevels []logrus.Level
}

func (hook *writerHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		return err
	}
	for _, w := range hook.Writer {
		w.Write([]byte(line))
	}
	return err
}

func (hook *writerHook) Levels() []logrus.Level {
	return hook.LogLevels
}

type Logger struct {
	*logrus.Logger
}

func NewLogger() (Logger, error) {
	logger := logrus.New()
	logger.SetReportCaller(true)
	logger.Formatter = &logrus.JSONFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			var filename string
			projectDir := "2023_1_ContentDealers"
			index := strings.Index(frame.File, projectDir)
			if index == -1 {
				filename = path.Base(frame.File)
			} else {
				// вывод пути от директории проекта
				filename = path.Clean(frame.File[index+len(projectDir):])
			}

			return fmt.Sprintf("%s()", frame.Function), fmt.Sprintf("%s:%d", filename, frame.Line)
		},
		PrettyPrint: true,
		DataKey:     "extra",
	}

	err := os.MkdirAll(logDir, 0644)
	if err != nil {
		panic(err)
	}

	logPath := fmt.Sprintf("%s/%s", logDir, logFile)
	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0640)
	if err != nil {
		return Logger{}, err
	}

	// используем только hooks, обычный вывод не нужен
	logger.SetOutput(io.Discard)

	// регистрируем hook
	logger.AddHook(&writerHook{
		Writer:    []io.Writer{file, os.Stdout},
		LogLevels: logrus.AllLevels,
	})

	logger.SetLevel(logrus.TraceLevel)

	return Logger{Logger: logger}, nil
}
