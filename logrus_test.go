package abstractlogger

import (
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"testing"
)

func BenchmarkWrappedLogrusLogger_String(b *testing.B) {
	logrusLogger := &logrus.Logger{
		Out:       ioutil.Discard,
		Formatter: new(logrus.JSONFormatter),
		Hooks:     make(logrus.LevelHooks),
		Level:     logrus.DebugLevel,
	}

	var logger Logger
	logger = NewLogrusLogger(logrusLogger, DebugLevel)

	b.ResetTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.DebugString("foo","bar","baz")
		}
	})
}

func BenchmarkWrappedLogrusLoggerLevelMiss_String(b *testing.B) {
	logrusLogger := &logrus.Logger{
		Out:       ioutil.Discard,
		Formatter: new(logrus.JSONFormatter),
		Hooks:     make(logrus.LevelHooks),
		Level:     logrus.InfoLevel,
	}

	var logger Logger
	logger = NewLogrusLogger(logrusLogger, InfoLevel)

	b.ResetTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.DebugString("foo","bar","baz")
		}
	})
}

func BenchmarkLogrusLogger_String(b *testing.B) {
	logger := &logrus.Logger{
		Out:       ioutil.Discard,
		Formatter: new(logrus.JSONFormatter),
		Hooks:     make(logrus.LevelHooks),
		Level:     logrus.DebugLevel,
	}


	b.ResetTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.WithField("foo","bar").Debug("baz")
		}
	})
}

func BenchmarkLogrusLoggerLevelMiss_String(b *testing.B) {
	logger := &logrus.Logger{
		Out:       ioutil.Discard,
		Formatter: new(logrus.JSONFormatter),
		Hooks:     make(logrus.LevelHooks),
		Level:     logrus.InfoLevel,
	}

	b.ResetTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.WithField("foo","bar").Debug("baz")
		}
	})
}