package logger

import (
    "runtime"
    "path"
    "fmt"
    "io"

    "github.com/sirupsen/logrus"
)

type writerHook struct {
    Writer []io.Writer
    LogLevels []logrus.Level
}

func init() {
    logger := logrus.New()
    logger.SetReportCaller(true)

    logger.Formatter = &logrus.TextFormatter {
        CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
            filename := path.Base(frame.File)
            return fmt.Sprintf("%s()", frame.Function),fmt.Sprintf("%s:%d", filename, frame.Line)
        },
        DisableColors: false,
        FullTimestamp: true,
    }
}
