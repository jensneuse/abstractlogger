package abstractlogger

import (
	"bytes"
	"errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"testing"
	"time"
)

func zapLogger(lvl zapcore.Level, syncer zapcore.WriteSyncer) *zap.Logger {
	ec := zap.NewProductionEncoderConfig()
	ec.EncodeDuration = zapcore.NanosDurationEncoder
	ec.EncodeTime = ZeroTimeEncoder
	enc := zapcore.NewJSONEncoder(ec)
	return zap.New(zapcore.NewCore(
		enc,
		syncer,
		lvl,
	))
}

func ZeroTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendInt64(0)
}

func TestZapLogger(t *testing.T) {

	level := zapcore.DebugLevel

	var directOut bytes.Buffer
	direct := zapLogger(level,zapcore.AddSync(&directOut))

	var wrappedOut bytes.Buffer
	wrapped := zapLogger(level,zapcore.AddSync(&wrappedOut))
	indirect := NewZapLogger(wrapped, DebugLevel)

	direct.Debug("baz",zap.String("foo","bar"))
	indirect.Debug("baz", String("foo", "bar"))

	direct.Info("baz",zap.String("foo","bar"))
	indirect.Info("baz", String("foo", "bar"))

	direct.Debug("baz",zap.Strings("foo",[]string{"foo","bar"}))
	indirect.Debug("baz", Strings("foo", []string{"foo","bar"}))

	direct.Warn("baz",zap.String("foo","bar"))
	indirect.Warn("baz", String("foo", "bar"))

	direct.Error("baz",zap.String("foo","bar"),zap.Error(errors.New("err")))
	indirect.Error("baz", String("foo", "bar"),Error(errors.New("err")))

	direct.Sugar().Debug("foo")
	direct.Sugar().Debugf("foo %s","bar")

	debugLogger := indirect.LevelLogger(DebugLevel)
	debugLogger.Println("foo")
	debugLogger.Printf("foo %s","bar")

	if directOut.String() != wrappedOut.String() {
		t.Fatalf("direct:\n%s\n\nindirect:\n%s\n",directOut.String(),wrappedOut.String())
	}
}
