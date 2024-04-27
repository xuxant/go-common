package logs

import (
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const CtxLogKey = "data-management"

var ContextLogger = ZerologMiddleware

type Level string

func logLevel(logLevel string) zerolog.Level {
	switch strings.ToLower(logLevel) {
	case "trace":
		return zerolog.TraceLevel

	case "debug":
		return zerolog.DebugLevel

	case "info":
		return zerolog.InfoLevel

	case "warn":
		return zerolog.WarnLevel

	case "error":
		return zerolog.ErrorLevel

	case "fatal":
		return zerolog.FatalLevel

	case "panic":
		return zerolog.PanicLevel
	}
	return zerolog.NoLevel
}

func LogLevelStr(logLevel zerolog.Level) string {
	switch logLevel {
	case zerolog.TraceLevel:
		return "trace"

	case zerolog.DebugLevel:
		return "debug"

	case zerolog.InfoLevel:
		return "info"

	case zerolog.WarnLevel:
		return "warn"

	case zerolog.ErrorLevel:
		return "error"

	case zerolog.FatalLevel:
		return "fatal"

	case zerolog.PanicLevel:
		return "panic"
	}
	return "unknown"
}

func ZerologMiddleware(parent zerolog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestLogger := parent.With().Logger()
		c.Set(CtxLogKey, requestLogger)
		c.Next()
	}
}

func LoggerForRequest(c *gin.Context) zerolog.Logger {
	logger, _ := c.Get(CtxLogKey)
	return logger.(zerolog.Logger)
}

func InitializeLogger(loggingLevel string, preetyfy bool) zerolog.Logger {
	level := logLevel(loggingLevel)
	zerolog.SetGlobalLevel(level)
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	if preetyfy {
		logger = log.Output(
			zerolog.ConsoleWriter{
				Out:        os.Stderr,
				TimeFormat: time.RFC3339,
			},
		)
	}

	logger.Info().Msgf("Logging Level: %s", loggingLevel)
	return logger
}
