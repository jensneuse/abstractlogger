package abstractlogger

import (
	"errors"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest"
	"io/ioutil"
	"testing"
)

func BenchmarkNoopLogger(b *testing.B) {

	noopLog := NoopLogger
	zapLog := zapLogger(zapcore.InfoLevel, &zaptest.Discarder{})
	logrusLog := logrusLogger(ioutil.Discard, logrus.InfoLevel)

	abstractZap := NewZapLogger(zapLog, InfoLevel)
	abstractLogrus := NewLogrusLogger(logrusLog, InfoLevel)

	b.Run("with interface", func(b *testing.B) {
		b.Run("log level invalid", func(b *testing.B) {
			b.Run("noop", func(b *testing.B) {
				b.ResetTimer()
				b.ReportAllocs()
				b.RunParallel(func(pb *testing.PB) {
					for pb.Next() {
						noopLog.Debug("foo",
							String("bar", "baz"),
							Bool("bool", true),
							Int("int", 123),
							Any("any", String("foo", "bar")),
						)
					}
				})
			})
			b.Run("logrus", func(b *testing.B) {
				b.ResetTimer()
				b.ReportAllocs()
				b.RunParallel(func(pb *testing.PB) {
					for pb.Next() {
						logrusLog.WithField("bar", "baz").
							WithField("bool", true).
							WithField("int", 123).
							WithField("any", String("foo", "bar")).
							Debug("foo")
					}
				})
			})
			b.Run("zap", func(b *testing.B) {
				b.ResetTimer()
				b.ReportAllocs()
				b.RunParallel(func(pb *testing.PB) {
					for pb.Next() {
						zapLog.Debug("foo",
							zap.String("bar", "baz"),
							zap.Bool("bool", true),
							zap.Int("int", 123),
							zap.Any("any", String("foo", "bar")),
						)
					}
				})
			})
			b.Run("abstract zap", func(b *testing.B) {
				b.ResetTimer()
				b.ReportAllocs()
				b.RunParallel(func(pb *testing.PB) {
					for pb.Next() {
						abstractZap.Debug("foo",
							String("bar", "baz"),
							Bool("bool", true),
							Int("int", 123),
							Any("any", String("foo", "bar")),
						)
					}
				})
			})
			b.Run("abstract logrus", func(b *testing.B) {
				b.ResetTimer()
				b.ReportAllocs()
				b.RunParallel(func(pb *testing.PB) {
					for pb.Next() {
						abstractLogrus.Debug("foo",
							String("bar", "baz"),
							Bool("bool", true),
							Int("int", 123),
							Any("any", String("foo", "bar")),
						)
					}
				})
			})
		})
		b.Run("log level valid", func(b *testing.B) {
			b.Run("noop", func(b *testing.B) {
				b.ResetTimer()
				b.ReportAllocs()
				b.RunParallel(func(pb *testing.PB) {
					for pb.Next() {
						noopLog.Info("foo",
							String("bar", "baz"),
							Bool("bool", true),
							Int("int", 123),
							Any("any", String("foo", "bar")),
						)
					}
				})
			})
			b.Run("logrus", func(b *testing.B) {
				b.ResetTimer()
				b.ReportAllocs()
				b.RunParallel(func(pb *testing.PB) {
					for pb.Next() {
						logrusLog.WithField("bar", "baz").
							WithField("bool", true).
							WithField("int", 123).
							WithField("any", String("foo", "bar")).
							Info("foo")
					}
				})
			})
			b.Run("zap", func(b *testing.B) {
				b.ResetTimer()
				b.ReportAllocs()
				b.RunParallel(func(pb *testing.PB) {
					for pb.Next() {
						zapLog.Info("foo",
							zap.String("bar", "baz"),
							zap.Bool("bool", true),
							zap.Int("int", 123),
							zap.Any("any", String("foo", "bar")),
						)
					}
				})
			})
			b.Run("abstract zap", func(b *testing.B) {
				b.ResetTimer()
				b.ReportAllocs()
				b.RunParallel(func(pb *testing.PB) {
					for pb.Next() {
						abstractZap.Info("foo",
							String("bar", "baz"),
							Bool("bool", true),
							Int("int", 123),
							Any("any", String("foo", "bar")),
						)
					}
				})
			})
			b.Run("abstract logrus", func(b *testing.B) {
				b.ResetTimer()
				b.ReportAllocs()
				b.RunParallel(func(pb *testing.PB) {
					for pb.Next() {
						abstractLogrus.Info("foo",
							String("bar", "baz"),
							Bool("bool", true),
							Int("int", 123),
							Any("any", String("foo", "bar")),
						)
					}
				})
			})
		})
	})
	b.Run("without interface", func(b *testing.B) {
		b.Run("log level invalid", func(b *testing.B) {
			b.Run("noop", func(b *testing.B) {
				b.ResetTimer()
				b.ReportAllocs()
				b.RunParallel(func(pb *testing.PB) {
					for pb.Next() {
						noopLog.Debug("foo",
							String("bar", "baz"),
							Bool("bool", true),
							Int("int", 123),
						)
					}
				})
			})
			b.Run("logrus", func(b *testing.B) {
				b.ResetTimer()
				b.ReportAllocs()
				b.RunParallel(func(pb *testing.PB) {
					for pb.Next() {
						logrusLog.WithField("bar", "baz").
							WithField("bool", true).
							WithField("int", 123).
							Debug("foo")
					}
				})
			})
			b.Run("zap", func(b *testing.B) {
				b.ResetTimer()
				b.ReportAllocs()
				b.RunParallel(func(pb *testing.PB) {
					for pb.Next() {
						zapLog.Debug("foo",
							zap.String("bar", "baz"),
							zap.Bool("bool", true),
							zap.Int("int", 123),
						)
					}
				})
			})
			b.Run("abstract zap", func(b *testing.B) {
				b.ResetTimer()
				b.ReportAllocs()
				b.RunParallel(func(pb *testing.PB) {
					for pb.Next() {
						abstractZap.Debug("foo",
							String("bar", "baz"),
							Bool("bool", true),
							Int("int", 123),
						)
					}
				})
			})
			b.Run("abstract logrus", func(b *testing.B) {
				b.ResetTimer()
				b.ReportAllocs()
				b.RunParallel(func(pb *testing.PB) {
					for pb.Next() {
						abstractLogrus.Debug("foo",
							String("bar", "baz"),
							Bool("bool", true),
							Int("int", 123),
						)
					}
				})
			})
		})
		b.Run("log level valid", func(b *testing.B) {
			b.Run("noop", func(b *testing.B) {
				b.ResetTimer()
				b.ReportAllocs()
				b.RunParallel(func(pb *testing.PB) {
					for pb.Next() {
						noopLog.Info("foo",
							String("bar", "baz"),
							Bool("bool", true),
							Int("int", 123),
						)
					}
				})
			})
			b.Run("logrus", func(b *testing.B) {
				b.ResetTimer()
				b.ReportAllocs()
				b.RunParallel(func(pb *testing.PB) {
					for pb.Next() {
						logrusLog.WithField("bar", "baz").
							WithField("bool", true).
							WithField("int", 123).
							Info("foo")
					}
				})
			})
			b.Run("zap", func(b *testing.B) {
				b.ResetTimer()
				b.ReportAllocs()
				b.RunParallel(func(pb *testing.PB) {
					for pb.Next() {
						zapLog.Info("foo",
							zap.String("bar", "baz"),
							zap.Bool("bool", true),
							zap.Int("int", 123),
						)
					}
				})
			})
			b.Run("abstract zap", func(b *testing.B) {
				b.ResetTimer()
				b.ReportAllocs()
				b.RunParallel(func(pb *testing.PB) {
					for pb.Next() {
						abstractZap.Info("foo",
							String("bar", "baz"),
							Bool("bool", true),
							Int("int", 123),
						)
					}
				})
			})
			b.Run("abstract logrus", func(b *testing.B) {
				b.ResetTimer()
				b.ReportAllocs()
				b.RunParallel(func(pb *testing.PB) {
					for pb.Next() {
						abstractLogrus.Info("foo",
							String("bar", "baz"),
							Bool("bool", true),
							Int("int", 123),
						)
					}
				})
			})
		})
	})

	b.Run("zap common use case", func(b *testing.B) {

		err := errors.New("foo")

		b.ResetTimer()
		b.ReportAllocs()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				abstractZap.Debug("foo",
					String("bar", "baz"),
					Bool("bool", true),
					Int("int", 123),
					Error(err),
				)
			}
		})
	})
}
