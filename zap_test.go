package abstractlogger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest"
	"testing"
)

func newZapLogger(lvl zapcore.Level) *zap.Logger {
	ec := zap.NewProductionEncoderConfig()
	ec.EncodeDuration = zapcore.NanosDurationEncoder
	ec.EncodeTime = zapcore.EpochNanosTimeEncoder
	enc := zapcore.NewJSONEncoder(ec)
	return zap.New(zapcore.NewCore(
		enc,
		&zaptest.Discarder{},
		lvl,
	))
}

func BenchmarkWrappedZapLoggerString(b *testing.B) {
	var logger Logger
	logger = NewZapLogger(newZapLogger(zapcore.DebugLevel), DebugLevel)

	b.ResetTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.DebugString("foo", "bar", "baz")
		}
	})
}

func BenchmarkWrappedZapLoggerStringNonParallel(b *testing.B) {
	var logger Logger
	logger = NewZapLogger(newZapLogger(zapcore.DebugLevel), DebugLevel)

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		logger.DebugString("foo", "bar", "baz")
	}
}

func BenchmarkWrappedZapLoggerLevelMiss(b *testing.B) {
	var logger Logger
	logger = NewZapLogger(newZapLogger(zapcore.DebugLevel), InfoLevel)

	b.ResetTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.DebugString("foo", "bar", "baz")
		}
	})
}

func BenchmarkWrappedZapLoggerBytes(b *testing.B) {
	var logger Logger
	logger = NewZapLogger(newZapLogger(zapcore.DebugLevel), DebugLevel)

	value := []byte("baz")

	b.ResetTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.DebugByteString("foo", "bar", value)
		}
	})
}

func BenchmarkWrappedZapLoggerLevelMissBytes(b *testing.B) {
	var logger Logger
	logger = NewZapLogger(newZapLogger(zapcore.InfoLevel), InfoLevel)

	value := []byte("baz")

	b.ResetTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.DebugByteString("foo", "bar", value)
		}
	})
}

func BenchmarkZapString(b *testing.B) {
	logger := newZapLogger(zapcore.DebugLevel)

	b.ResetTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Debug("foo", zap.String("bar", "baz"))
		}
	})
}

func BenchmarkZapStringNonParallel(b *testing.B) {
	logger := newZapLogger(zapcore.DebugLevel)

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		logger.Debug("foo", zap.String("bar", "baz"))
	}
}

func BenchmarkZapLevelMiss(b *testing.B) {
	logger := newZapLogger(zapcore.InfoLevel)

	b.ResetTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Debug("foo", zap.String("bar", "baz"))
		}
	})
}
