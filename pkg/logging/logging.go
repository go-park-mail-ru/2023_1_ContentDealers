package logging

import (
	"context"
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

type writerHook struct {
	Writer      []io.Writer
	LogLevels   []logrus.Level
	serviceName string
}

func (hook *writerHook) Fire(entry *logrus.Entry) error {
	entry.Data["service"] = hook.serviceName
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

func (lg *Logger) WithRequestID(ctx context.Context) *logrus.Entry {
	reqID, ok := ctx.Value("requestID").(string)
	if !ok {
		reqID = "unknown"
	}
	return lg.WithField("request_id", reqID)

}

func NewLogger(cfg LoggingConfig, serviceName string) (Logger, error) {
	logger := logrus.New()
	if cfg.Dir == "" && cfg.Filename == "" && cfg.ProjectDir == "" {
		return Logger{Logger: logger}, nil
	}
	logger.SetReportCaller(true)
	logger.Formatter = &logrus.JSONFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			var filename string
			projectDir := cfg.ProjectDir
			index := strings.Index(frame.File, projectDir)
			if index == -1 {
				// попадает случай, если cfg.ProjectDir пустой
				filename = frame.File
			} else {
				// вывод пути от директории проекта
				filename = path.Clean(frame.File[index+len(projectDir):])[1:]
			}

			return fmt.Sprintf("%s()", frame.Function), fmt.Sprintf("%s:%d", filename, frame.Line)
		},
		DataKey:         "extra",
		TimestampFormat: "2006-01-02 15:04:05",
		FieldMap: logrus.FieldMap{
			// поля выводятся в алфавитном порядке
			logrus.FieldKeyTime: "__time",
			logrus.FieldKeyMsg:  "_msg",
			"service":           serviceName,
		},
	}

	// nolint:gomnd
	err := os.MkdirAll(cfg.Dir, 0777)
	if err != nil {
		panic(err)
	}

	logPath := fmt.Sprintf("%s/%s", cfg.Dir, cfg.Filename)
	// nolint:gomnd
	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0640)
	if err != nil {
		return Logger{}, err
	}

	// используем только hooks, обычный вывод не нужен
	logger.SetOutput(io.Discard)

	// определяем уровни логирования
	var levels []logrus.Level
	if len(cfg.Levels) != 0 {
		if cfg.Levels[0] == "all" {
			levels = logrus.AllLevels
		} else {
			for _, levelString := range cfg.Levels {
				level, err := logrus.ParseLevel(levelString)
				if err != nil {
					return Logger{}, err
				}
				levels = append(levels, level)
			}
		}
	}

	// регистрируем hook
	logger.AddHook(&writerHook{
		Writer:      []io.Writer{file, os.Stdout},
		LogLevels:   levels,
		serviceName: serviceName,
	})

	logger.SetLevel(logrus.TraceLevel)

	return Logger{Logger: logger}, nil
}
