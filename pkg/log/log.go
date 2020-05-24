package log

import (
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
)

func Setup(service string) {

	logLevel := zerolog.InfoLevel
	switch strings.ToLower(os.Getenv("LOG_LEVEL")) {
	case "debug":
		logLevel = zerolog.DebugLevel
	case "info":
	case "warn":
		logLevel = zerolog.WarnLevel
	case "error":
		logLevel = zerolog.ErrorLevel
	case "fatal", "panic":
		logLevel = zerolog.PanicLevel
	}

	zerolog.SetGlobalLevel(logLevel)

	// callerHook := ZerologCallerHook{
	// 	RootPath: rootPath,
	// }

	// Add file and line to caller key
	zlog.Logger = zlog.With().
		Str("service", service).
		Caller().
		Logger()

	if strings.ToLower(os.Getenv("LOG_FORMAT")) != "json" {
		zlog.Logger = zlog.Output(zerolog.ConsoleWriter{
			Out:        os.Stderr,
			TimeFormat: time.RFC3339Nano,
		})
	}

	zerolog.TimeFieldFormat = time.RFC3339Nano

}

func IfError(err error) bool {

	if err != nil {

		_, file, line, _ := runtime.Caller(1)

		Error().
			Str("file", file).
			Int("line", line).
			Err(err).
			Msg("")

		return true
	}

	return false
}

func IfErrorWarn(err error, msg ...string) bool {
	if err != nil {

		_, file, line, _ := runtime.Caller(1)

		Warn().
			Str("file", file).
			Int("line", line).
			Err(err).
			Msg(strings.Join(msg, " "))

		return true
	}

	return false
}

func Debug() *zerolog.Event {
	return zlog.Debug()
}

func Info() *zerolog.Event {
	return zlog.Info()
}

func Warn() *zerolog.Event {
	return zlog.Warn()
}

func Fatal() *zerolog.Event {
	return zlog.Fatal()
}

func Error() *zerolog.Event {
	return zlog.Error()
}

func With() zerolog.Context {
	return zlog.With()
}
