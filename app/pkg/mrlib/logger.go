package mrlib

import (
    "auth-le-back/pkg/mrapp"
    "fmt"
    "log"
    "os"
    "runtime"
    "time"
)

const (
    datetime = "2006/01/02 15:04:05"
)

type Logger struct {
    level mrapp.LogLevel
    color bool
    infoLog *log.Logger
    errLog *log.Logger
}

// Make sure the Logger conforms with the mrapp.Logger interface
var _ mrapp.Logger = (*Logger)(nil)

func NewLogger(level string, color bool) *Logger {
    lvl := mrapp.ParseLogLevel(level)

    infoLog := log.New(os.Stdout, "", 0)
    errLog := log.New(os.Stderr, "", 0)

    return &Logger {
        level: lvl,
        color: color,
        infoLog: infoLog,
        errLog: errLog,
    }
}

func (l *Logger) Fatal(message any, args ...any) {
    var buf []byte

    l.formatHeader(&buf, "FATAL", 2)
    l.formatMessage(&buf, message)

    if len(args) == 0 {
        l.errLog.Fatal(string(buf))
    } else {
        l.errLog.Fatalf(string(buf), args...)
    }
}

func (l *Logger) Error(message any, args ...any) {
    if l.level >= mrapp.LogErrorLevel {
        l.logPrint(l.errLog, 3,"ERROR", message, args)
    }
}

func (l *Logger) Warn(message string, args ...any) {
    if l.level >= mrapp.LogWarnLevel {
        l.logPrint(l.infoLog, 3,"WARN", message, args)
    }
}

func (l *Logger) Info(message string, args ...any) {
    if l.level >= mrapp.LogInfoLevel {
        l.logPrint(l.infoLog, 0,"INFO", message, args)
    }
}

func (l *Logger) Debug(message any, args ...any) {
    if l.level >= mrapp.LogDebugLevel {
        l.logPrint(l.infoLog, 0, "DEBUG", message, args)
    }
}

func (l *Logger) logPrint(logger *log.Logger, calldepth int, prefix string, message any, args []any) {
    var buf []byte

    l.formatHeader(&buf, prefix, calldepth)
    l.formatMessage(&buf, message)

    if len(args) == 0 {
        logger.Print(string(buf))
    } else {
        logger.Printf(string(buf), args...)
    }
}

func (l *Logger) formatHeader(buf *[]byte, prefix string, calldepth int) {
    *buf = append(*buf, time.Now().Format(datetime)...)
    *buf = append(*buf, ' ')

    if calldepth > 0 {
        _, file, line, ok := runtime.Caller(calldepth)

        if !ok {
            file = "???"
            line = 0
        }

        *buf = append(*buf, fmt.Sprintf("%s:%d ", file, line)...)
    }

    *buf = append(*buf, prefix...)
    *buf = append(*buf, '\t')
}

func (l *Logger) formatMessage(buf *[]byte, message any) {
    switch msg := message.(type) {
        case error:
            *buf = append(*buf, msg.Error()...)
        case string:
            *buf = append(*buf, msg...)
        default:
            *buf = append(*buf, fmt.Sprintf("Message %v has unknown type %v", message, msg)...)
    }
}
