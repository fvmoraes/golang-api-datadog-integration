package logs

import (
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace"
)

const myTimeFormat = "02/01/2006 15:04:05"

func PopulateErrorLogFile(traceSpan ddtrace.Span, dataMessage string, logType string, logInputed string, logError error) {
	if logInputed != "" {
		WriteLogFile(traceSpan, logType, logInputed, dataMessage)
	}
	if logError != nil {
		WriteLogFile(traceSpan, logType, logError.Error(), dataMessage)
	}
}

func WriteLogFile(traceSpan ddtrace.Span, logType string, information string, dataMessage string) {
	var log = logrus.New()
	log.Formatter = new(logrus.JSONFormatter)
	log.Level = logrus.TraceLevel
	trace := ddtrace.Span.Context(traceSpan).TraceID()
	span := ddtrace.Span.Context(traceSpan).SpanID()
	logFile, _ := os.OpenFile("/home/franklin.moraes@poa01.local/go/src/api-sample/logs/logg.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	log.Out = logFile
	if logType == "INFO" {
		log.WithFields(logrus.Fields{
			"env":         "dev",
			"type":        "log",
			"data":        dataMessage,
			"dd.trace_id": trace,
			"dd.span_id":  span,
		}).Info(information)
	}
	if logType == "ERROR" {
		log.WithFields(logrus.Fields{
			"env":         "dev",
			"type":        "log",
			"data":        dataMessage,
			"dd.trace_id": trace,
			"dd.span_id":  span,
		}).Error(information)
	}
	if logType == "WARNING" {
		log.WithFields(logrus.Fields{
			"env":         "dev",
			"type":        "log",
			"data":        dataMessage,
			"dd.trace_id": trace,
			"dd.span_id":  span,
		}).Warning(information)
	}
	defer logFile.Close()
}
