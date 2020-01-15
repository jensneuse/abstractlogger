package abstractlogger

import (
	"go.uber.org/zap"
)

func NewZapLogger(zapLogger *zap.Logger, level Level) *ZapLogger {
	return &ZapLogger{
		l:          zapLogger,
		levelCheck: NewLevelCheck(level),
	}
}

// ZapLogger implements the Logging frontend using the popular logging backend zap
// It uses the LevelCheck helper to increase performance.
type ZapLogger struct {
	l          *zap.Logger
	levelCheck LevelCheck
}

func (z *ZapLogger) field(field Field) zap.Field {
	switch field.Kind {
	case StringField:
		return zap.String(field.Key, field.StringValue)
	case IntField:
		return zap.Int(field.Key, int(field.IntValue))
	case BoolField:
		return zap.Bool(field.Key, field.IntValue != 0)
	case ByteStringField:
		return zap.ByteString(field.Key, field.ByteValue)
	case ErrorField:
		return zap.Error(field.ErrorValue)
	case NamedErrorField:
		return zap.NamedError(field.Key,field.ErrorValue)
	default:
		return zap.Any(field.Key, field.IfaceValue)
	}
}

func (z *ZapLogger) fields(fields []Field) []zap.Field {
	out := make([]zap.Field, len(fields))
	for i := range fields {
		out[i] = z.field(fields[i])
	}
	return out
}

func (z *ZapLogger) Debug(msg string, fields ...Field) {
	if !z.levelCheck.Check(DebugLevel) {
		return
	}
	z.l.Debug(msg, z.fields(fields)...)
}

func (z *ZapLogger) Info(msg string, fields ...Field) {
	if !z.levelCheck.Check(InfoLevel) {
		return
	}
	z.l.Info(msg, z.fields(fields)...)
}

func (z *ZapLogger) Warn(msg string, fields ...Field) {
	if !z.levelCheck.Check(WarnLevel) {
		return
	}
	z.l.Warn(msg, z.fields(fields)...)
}

func (z *ZapLogger) Error(msg string, fields ...Field) {
	if !z.levelCheck.Check(ErrorLevel) {
		return
	}
	z.l.Error(msg, z.fields(fields)...)
}

func (z *ZapLogger) Fatal(msg string, fields ...Field) {
	if !z.levelCheck.Check(FatalLevel) {
		return
	}
	z.l.Fatal(msg, z.fields(fields)...)
}

func (z *ZapLogger) Panic(msg string, fields ...Field) {
	if !z.levelCheck.Check(PanicLevel) {
		return
	}
	z.l.Panic(msg, z.fields(fields)...)
}
