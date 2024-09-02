package utils

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func init() {
	Logger = logrus.New()

	// Set JSON formatter
	Logger.SetFormatter(&logrus.JSONFormatter{})

	// Output to stdout by default
	Logger.SetOutput(os.Stdout)

	// Set log level to Info
	Logger.SetLevel(logrus.InfoLevel)
}
func HandleError(err error, message string) {
	if err != nil {
		Logger.WithFields(logrus.Fields{
			"error": err,
		}).Error(message)
	}
}

func Info(message string, fields map[string]interface{}) {
	Logger.WithFields(logrus.Fields{
		"level": "info",
		"extra": fields,
	}).Info(message)
}
func Warn(message string, fields map[string]interface{}) {
	Logger.WithFields(logrus.Fields{
		"level": "warn",
		"extra": fields,
	}).Warn(message)
}

func Error(message string, fields map[string]interface{}) {
	Logger.WithFields(logrus.Fields{
		"level": "error",
		"extra": fields,
	}).Error(message)
}

func LogHTTPRequest(method, url string, statusCode int, duration time.Duration, fields map[string]interface{}) {
	Logger.WithFields(logrus.Fields{
		"method":      method,
		"url":         url,
		"status_code": statusCode,
		"duration":    duration.Seconds(), 
		"extra":       fields,
	}).Info("HTTP Request")
}

func LogHTTPRequestError(method, url string, statusCode int, duration time.Duration, err error, fields map[string]interface{}) {
	Logger.WithFields(logrus.Fields{
		"method":      method,
		"url":         url,
		"status_code": statusCode,
		"duration":    duration.Seconds(),
		"error":       err,
		"extra":       fields,
	}).Error("HTTP Request Error")
}
