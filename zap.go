package abstractlogger

import (
	"go.uber.org/zap"
)

func NewZapLogger(zapLogger *zap.Logger,level Level) *ZapLogger{
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

func (z *ZapLogger) Debug(msg string) {
	if !z.levelCheck.Check(DebugLevel){
		return
	}
	z.l.Debug(msg)
}

func (z *ZapLogger) field(field Field) zap.Field {
	switch field.Kind {
	case StringField:
		return zap.String(field.Key,field.StringValue)
	case IntField:
		return zap.Int(field.Key,int(field.IntValue))
	case BoolField:
		return zap.Bool(field.Key,field.IntValue != 0)
	case ByteStringField:
		return zap.ByteString(field.Key,field.ByteValue)
	default:
		return zap.Skip()
	}
}

func (z *ZapLogger) DebugField(msg string, field Field) {
	if !z.levelCheck.Check(DebugLevel){
		return
	}
	z.l.Debug(msg,z.field(field))
}

func (z *ZapLogger) DebugField2(msg string, field1 Field, field2 Field) {
	if !z.levelCheck.Check(DebugLevel){
		return
	}
	z.l.Debug(msg,
		z.field(field1),
		z.field(field2),
	)
}

func (z *ZapLogger) DebugField3(msg string, field1 Field, field2 Field, field3 Field) {
	if !z.levelCheck.Check(DebugLevel){
		return
	}
	z.l.Debug(msg,
		z.field(field1),
		z.field(field2),
		z.field(field3),
	)
}

func (z *ZapLogger) DebugField4(msg string, field1 Field, field2 Field, field3 Field, field4 Field) {
	if !z.levelCheck.Check(DebugLevel){
		return
	}
	z.l.Debug(msg,
		z.field(field1),
		z.field(field2),
		z.field(field3),
		z.field(field4),
	)
}

func (z *ZapLogger) DebugField5(msg string, field1 Field, field2 Field, field3 Field, field4 Field, field5 Field) {
	if !z.levelCheck.Check(DebugLevel){
		return
	}
	z.l.Debug(msg,
		z.field(field1),
		z.field(field2),
		z.field(field3),
		z.field(field4),
		z.field(field5),
	)
}

func (z *ZapLogger) DebugField6(msg string, field1 Field, field2 Field, field3 Field, field4 Field, field5 Field, field6 Field) {
	if !z.levelCheck.Check(DebugLevel){
		return
	}
	z.l.Debug(msg,
		z.field(field1),
		z.field(field2),
		z.field(field3),
		z.field(field4),
		z.field(field5),
		z.field(field6),
	)
}

func (z *ZapLogger) Info(msg string) {
	if !z.levelCheck.Check(InfoLevel){
		return
	}
	z.l.Info(msg)
}

func (z *ZapLogger) InfoField(msg string, field Field) {
	if !z.levelCheck.Check(InfoLevel){
		return
	}
	z.l.Info(msg,
		z.field(field),
	)
}

func (z *ZapLogger) InfoField2(msg string, field1 Field, field2 Field) {
	if !z.levelCheck.Check(InfoLevel){
		return
	}
	z.l.Info(msg,
		z.field(field1),
		z.field(field2),
	)
}

func (z *ZapLogger) InfoField3(msg string, field1 Field, field2 Field, field3 Field) {
	if !z.levelCheck.Check(InfoLevel){
		return
	}
	z.l.Info(msg,
		z.field(field1),
		z.field(field2),
		z.field(field3),
	)
}

func (z *ZapLogger) InfoField4(msg string, field1 Field, field2 Field, field3 Field, field4 Field) {
	if !z.levelCheck.Check(InfoLevel){
		return
	}
	z.l.Info(msg,
		z.field(field1),
		z.field(field2),
		z.field(field3),
		z.field(field4),
	)
}

func (z *ZapLogger) InfoField5(msg string, field1 Field, field2 Field, field3 Field, field4 Field, field5 Field) {
	if !z.levelCheck.Check(InfoLevel){
		return
	}
	z.l.Info(msg,
		z.field(field1),
		z.field(field2),
		z.field(field3),
		z.field(field4),
		z.field(field5),
	)
}

func (z *ZapLogger) InfoField6(msg string, field1 Field, field2 Field, field3 Field, field4 Field, field5 Field, field6 Field) {
	if !z.levelCheck.Check(InfoLevel){
		return
	}
	z.l.Info(msg,
		z.field(field1),
		z.field(field2),
		z.field(field3),
		z.field(field4),
		z.field(field5),
		z.field(field6),
	)
}

func (z *ZapLogger) Warn(msg string) {
	if !z.levelCheck.Check(WarnLevel){
		return
	}
	z.l.Warn(msg)
}

func (z *ZapLogger) WarnField(msg string, field Field) {
	if !z.levelCheck.Check(WarnLevel){
		return
	}
	z.l.Warn(msg,
		z.field(field),
	)
}

func (z *ZapLogger) WarnField2(msg string, field1 Field, field2 Field) {
	if !z.levelCheck.Check(WarnLevel){
		return
	}
	z.l.Warn(msg,
		z.field(field1),
		z.field(field2),
	)
}

func (z *ZapLogger) WarnField3(msg string, field1 Field, field2 Field, field3 Field) {
	if !z.levelCheck.Check(WarnLevel){
		return
	}
	z.l.Warn(msg,
		z.field(field1),
		z.field(field2),
		z.field(field3),
	)
}

func (z *ZapLogger) WarnField4(msg string, field1 Field, field2 Field, field3 Field, field4 Field) {
	if !z.levelCheck.Check(WarnLevel){
		return
	}
	z.l.Warn(msg,
		z.field(field1),
		z.field(field2),
		z.field(field3),
		z.field(field4),
	)
}

func (z *ZapLogger) WarnField5(msg string, field1 Field, field2 Field, field3 Field, field4 Field, field5 Field) {
	if !z.levelCheck.Check(WarnLevel){
		return
	}
	z.l.Warn(msg,
		z.field(field1),
		z.field(field2),
		z.field(field3),
		z.field(field4),
		z.field(field5),
	)
}

func (z *ZapLogger) WarnField6(msg string, field1 Field, field2 Field, field3 Field, field4 Field, field5 Field, field6 Field) {
	if !z.levelCheck.Check(WarnLevel){
		return
	}
	z.l.Warn(msg,
		z.field(field1),
		z.field(field2),
		z.field(field3),
		z.field(field4),
		z.field(field5),
		z.field(field6),
	)
}

func (z *ZapLogger) Error(msg string) {
	if !z.levelCheck.Check(ErrorLevel){
		return
	}
	z.l.Error(msg)
}

func (z *ZapLogger) ErrorField(msg string, field Field) {
	if !z.levelCheck.Check(ErrorLevel){
		return
	}
	z.l.Error(msg,
		z.field(field),
	)
}

func (z *ZapLogger) ErrorField2(msg string, field1 Field, field2 Field) {
	if !z.levelCheck.Check(ErrorLevel){
		return
	}
	z.l.Error(msg,
		z.field(field1),
		z.field(field2),
	)
}

func (z *ZapLogger) ErrorField3(msg string, field1 Field, field2 Field, field3 Field) {
	if !z.levelCheck.Check(ErrorLevel){
		return
	}
	z.l.Error(msg,
		z.field(field1),
		z.field(field2),
		z.field(field3),
	)
}

func (z *ZapLogger) ErrorField4(msg string, field1 Field, field2 Field, field3 Field, field4 Field) {
	if !z.levelCheck.Check(ErrorLevel){
		return
	}
	z.l.Error(msg,
		z.field(field1),
		z.field(field2),
		z.field(field3),
		z.field(field4),
	)
}

func (z *ZapLogger) ErrorField5(msg string, field1 Field, field2 Field, field3 Field, field4 Field, field5 Field) {
	if !z.levelCheck.Check(ErrorLevel){
		return
	}
	z.l.Error(msg,
		z.field(field1),
		z.field(field2),
		z.field(field3),
		z.field(field4),
		z.field(field5),
	)
}

func (z *ZapLogger) ErrorField6(msg string, field1 Field, field2 Field, field3 Field, field4 Field, field5 Field, field6 Field) {
	if !z.levelCheck.Check(ErrorLevel){
		return
	}
	z.l.Error(msg,
		z.field(field1),
		z.field(field2),
		z.field(field3),
		z.field(field4),
		z.field(field5),
		z.field(field6),
	)
}

func (z *ZapLogger) Fatal(msg string) {
	if !z.levelCheck.Check(FatalLevel){
		return
	}
	z.l.Fatal(msg)
}

func (z *ZapLogger) FatalField(msg string, field Field) {
	if !z.levelCheck.Check(FatalLevel){
		return
	}
	z.l.Fatal(msg,
		z.field(field),
	)
}

func (z *ZapLogger) FatalField2(msg string, field1 Field, field2 Field) {
	if !z.levelCheck.Check(FatalLevel){
		return
	}
	z.l.Fatal(msg,
		z.field(field1),
		z.field(field2),
	)
}

func (z *ZapLogger) FatalField3(msg string, field1 Field, field2 Field, field3 Field) {
	if !z.levelCheck.Check(FatalLevel){
		return
	}
	z.l.Fatal(msg,
		z.field(field1),
		z.field(field2),
		z.field(field3),
	)
}

func (z *ZapLogger) FatalField4(msg string, field1 Field, field2 Field, field3 Field, field4 Field) {
	if !z.levelCheck.Check(FatalLevel){
		return
	}
	z.l.Fatal(msg,
		z.field(field1),
		z.field(field2),
		z.field(field3),
		z.field(field4),
	)
}

func (z *ZapLogger) FatalField5(msg string, field1 Field, field2 Field, field3 Field, field4 Field, field5 Field) {
	if !z.levelCheck.Check(FatalLevel){
		return
	}
	z.l.Fatal(msg,
		z.field(field1),
		z.field(field2),
		z.field(field3),
		z.field(field4),
		z.field(field5),
	)
}

func (z *ZapLogger) FatalField6(msg string, field1 Field, field2 Field, field3 Field, field4 Field, field5 Field, field6 Field) {
	if !z.levelCheck.Check(FatalLevel){
		return
	}
	z.l.Fatal(msg,
		z.field(field1),
		z.field(field2),
		z.field(field3),
		z.field(field4),
		z.field(field5),
		z.field(field6),
	)
}

func (z *ZapLogger) Panic(msg string) {
	if !z.levelCheck.Check(PanicLevel){
		return
	}
	z.l.Panic(msg)
}

func (z *ZapLogger) PanicField(msg string, field Field) {
	if !z.levelCheck.Check(PanicLevel){
		return
	}
	z.l.Panic(msg,
		z.field(field),
	)
}

func (z *ZapLogger) PanicField2(msg string, field1 Field, field2 Field) {
	if !z.levelCheck.Check(PanicLevel){
		return
	}
	z.l.Panic(msg,
		z.field(field1),
		z.field(field2),
	)
}

func (z *ZapLogger) PanicField3(msg string, field1 Field, field2 Field, field3 Field) {
	if !z.levelCheck.Check(PanicLevel){
		return
	}
	z.l.Panic(msg,
		z.field(field1),
		z.field(field2),
		z.field(field3),
	)
}

func (z *ZapLogger) PanicField4(msg string, field1 Field, field2 Field, field3 Field, field4 Field) {
	if !z.levelCheck.Check(PanicLevel){
		return
	}
	z.l.Panic(msg,
		z.field(field1),
		z.field(field2),
		z.field(field3),
		z.field(field4),
	)
}

func (z *ZapLogger) PanicField5(msg string, field1 Field, field2 Field, field3 Field, field4 Field, field5 Field) {
	if !z.levelCheck.Check(PanicLevel){
		return
	}
	z.l.Panic(msg,
		z.field(field1),
		z.field(field2),
		z.field(field3),
		z.field(field4),
		z.field(field5),
	)
}

func (z *ZapLogger) PanicField6(msg string, field1 Field, field2 Field, field3 Field, field4 Field, field5 Field, field6 Field) {
	if !z.levelCheck.Check(PanicLevel){
		return
	}
	z.l.Panic(msg,
		z.field(field1),
		z.field(field2),
		z.field(field3),
		z.field(field4),
		z.field(field5),
		z.field(field6),
	)
}