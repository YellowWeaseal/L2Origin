package middleware

import (
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"time"
)

func NewLogger(logFile string, level string) *logrus.Logger {
	var logger = logrus.New()
	f, err := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		logger.Fatal(err)
	}
	mw := io.MultiWriter(os.Stdout, f)
	logrus.SetOutput(mw)

	var loggerLevel logrus.Level
	switch level {
	case "info":
		loggerLevel = logrus.InfoLevel
		break
	case "debug":
		loggerLevel = logrus.DebugLevel
		break
	case "warn":
		loggerLevel = logrus.WarnLevel
		break
	case "error":
		loggerLevel = logrus.ErrorLevel
		break
	}
	logger.SetLevel(loggerLevel)
	return logger
}
func LoggerMiddleware(log *logrus.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
			latencyTime := time.Now()
			next.ServeHTTP(resp, req)
			log.WithFields(logrus.Fields{
				"method":  req.Method,
				"uri":     req.RequestURI,
				"host":    req.Host,
				"ip":      req.RemoteAddr,
				"latency": time.Since(latencyTime).Seconds(),
			}).Info()
		})
	}
}
