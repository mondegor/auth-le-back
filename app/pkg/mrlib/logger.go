package mrlib

import (
    "auth-le-back/pkg/mrapp"
    "fmt"
    "log"
    "os"
    "strings"
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

    infoLog := log.New(os.Stdout, "", log.Ldate | log.Ltime)
    errLog := log.New(os.Stderr, "", log.Ldate | log.Ltime | log.Lshortfile)

    return &Logger {
        level: lvl,
        color: color,
        infoLog: infoLog,
        errLog: errLog,
    }
}

func (l *Logger) Fatal(message any, args ...any) {
    if len(args) == 0 {
        l.errLog.Fatal(getPrefix("fatal"), getMessage(message))
    } else {
        l.errLog.Fatalf(getPrefix("fatal") + getMessage(message), args...)
    }
}

func (l *Logger) Error(message any, args ...any) {
    if l.level >= mrapp.LogErrorLevel {
        logPrint(l.errLog, "error", message, args)
    }
}

func (l *Logger) Warn(message string, args ...any) {
    if l.level >= mrapp.LogWarnLevel {
        logPrint(l.infoLog, "warn", message, args)
    }
}

func (l *Logger) Info(message string, args ...any) {
    if l.level >= mrapp.LogInfoLevel {
        logPrint(l.infoLog, "info", message, args)
    }
}

func (l *Logger) Debug(message any, args ...any) {
    if l.level >= mrapp.LogDebugLevel {
        logPrint(l.infoLog, "debug", message, args)
    }
}

func getPrefix(prefix string) string {
    return strings.ToUpper(prefix) + "\t"
}

func getMessage(message any) string {
    switch msg := message.(type) {
        case error:
            return msg.Error()
        case string:
            return msg
        default:
            return fmt.Sprintf("Message %v has unknown type %v", message, msg)
    }
}

func logPrint(logger *log.Logger, prefix string, message any, args []any) {
    if len(args) == 0 {
        logger.Print(getPrefix(prefix), getMessage(message))
    } else {
        logger.Printf(getPrefix(prefix) + getMessage(message), args...)
    }
}


