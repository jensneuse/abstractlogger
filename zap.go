package abstractlogger

import (
	"go.uber.org/zap"
)

// ZapLogger implements the Logging frontend using the popular logging backend zap
// It uses the LevelCheck helper to increase performance.
type ZapLogger struct {
	l          *zap.Logger
	levelCheck LevelCheck
}

func NewZapLogger(zapLogger *zap.Logger,level Level) *ZapLogger{
	return &ZapLogger{
		l:          zapLogger,
		levelCheck: NewLevelCheck(level),
	}
}

func (z *ZapLogger) DebugString(msg string, key, value string) {
	if !z.levelCheck.Check(DebugLevel){
		return
	}
	z.l.Debug(msg,zap.String(key,value))
}

func (z *ZapLogger) DebugInt(msg string, key string, value int) {
	if !z.levelCheck.Check(DebugLevel){
		return
	}
	z.l.Debug(msg,zap.Int(key,value))
}

func (z *ZapLogger) DebugInt64(msg string, key string, value int64) {
	if !z.levelCheck.Check(DebugLevel){
		return
	}
	z.l.Debug(msg,zap.Int64(key,value))
}

func (z *ZapLogger) DebugFloat32(msg string, key string, value float32) {
	if !z.levelCheck.Check(DebugLevel){
		return
	}
	z.l.Debug(msg,zap.Float32(key,value))
}

func (z *ZapLogger) DebugFloat64(msg string, key string, value float64) {
	if !z.levelCheck.Check(DebugLevel){
		return
	}
	z.l.Debug(msg,zap.Float64(key,value))
}

func (z *ZapLogger) DebugByteString(msg string, key string, value []byte) {
	if !z.levelCheck.Check(DebugLevel){
		return
	}
	z.l.Debug(msg,zap.ByteString(key,value))
}

func (z *ZapLogger) ErrorString(msg string, key, value string) {
	if !z.levelCheck.Check(ErrorLevel){
		return
	}
	z.l.Error(msg,zap.String(key,value))
}

func (z *ZapLogger) ErrorInt(msg string, key string, value int) {
	if !z.levelCheck.Check(ErrorLevel){
		return
	}
	z.l.Error(msg,zap.Int(key,value))
}

func (z *ZapLogger) ErrorInt64(msg string, key string, value int64) {
	if !z.levelCheck.Check(ErrorLevel){
		return
	}
	z.l.Error(msg,zap.Int64(key,value))
}

func (z *ZapLogger) ErrorFloat32(msg string, key string, value float32) {
	if !z.levelCheck.Check(ErrorLevel){
		return
	}
	z.l.Error(msg,zap.Float32(key,value))
}

func (z *ZapLogger) ErrorFloat64(msg string, key string, value float64) {
	if !z.levelCheck.Check(ErrorLevel){
		return
	}
	z.l.Error(msg,zap.Float64(key,value))
}

func (z *ZapLogger) ErrorByteString(msg string, key string, value []byte) {
	if !z.levelCheck.Check(ErrorLevel){
		return
	}
	z.l.Error(msg,zap.ByteString(key,value))
}

func (z *ZapLogger) FatalString(msg string, key, value string) {
	if !z.levelCheck.Check(FatalLevel){
		return
	}
	z.l.Error(msg,zap.String(key,value))
}

func (z *ZapLogger) FatalInt(msg string, key string, value int) {
	if !z.levelCheck.Check(FatalLevel){
		return
	}
	z.l.Error(msg,zap.Int(key,value))
}

func (z *ZapLogger) FatalInt64(msg string, key string, value int64) {
	if !z.levelCheck.Check(FatalLevel){
		return
	}
	z.l.Error(msg,zap.Int64(key,value))
}

func (z *ZapLogger) FatalFloat32(msg string, key string, value float32) {
	if !z.levelCheck.Check(FatalLevel){
		return
	}
	z.l.Error(msg,zap.Float32(key,value))
}

func (z *ZapLogger) FatalFloat64(msg string, key string, value float64) {
	if !z.levelCheck.Check(FatalLevel){
		return
	}
	z.l.Error(msg,zap.Float64(key,value))
}

func (z *ZapLogger) FatalByteString(msg string, key string, value []byte) {
	if !z.levelCheck.Check(FatalLevel){
		return
	}
	z.l.Error(msg,zap.ByteString(key,value))
}

func (z *ZapLogger) InfoString(msg string, key, value string) {
	if !z.levelCheck.Check(InfoLevel){
		return
	}
	z.l.Error(msg,zap.String(key,value))
}

func (z *ZapLogger) InfoInt(msg string, key string, value int) {
	if !z.levelCheck.Check(InfoLevel){
		return
	}
	z.l.Error(msg,zap.Int(key,value))
}

func (z *ZapLogger) InfoInt64(msg string, key string, value int64) {
	if !z.levelCheck.Check(InfoLevel){
		return
	}
	z.l.Error(msg,zap.Int64(key,value))
}

func (z *ZapLogger) InfoFloat32(msg string, key string, value float32) {
	if !z.levelCheck.Check(InfoLevel){
		return
	}
	z.l.Error(msg,zap.Float32(key,value))
}

func (z *ZapLogger) InfoFloat64(msg string, key string, value float64) {
	if !z.levelCheck.Check(InfoLevel){
		return
	}
	z.l.Error(msg,zap.Float64(key,value))
}

func (z *ZapLogger) InfoByteString(msg string, key string, value []byte) {
	if !z.levelCheck.Check(InfoLevel){
		return
	}
	z.l.Error(msg,zap.ByteString(key,value))
}

func (z *ZapLogger) PanicString(msg string, key, value string) {
	if !z.levelCheck.Check(PanicLevel){
		return
	}
	z.l.Error(msg,zap.String(key,value))
}

func (z *ZapLogger) PanicInt(msg string, key string, value int) {
	if !z.levelCheck.Check(PanicLevel){
		return
	}
	z.l.Error(msg,zap.Int(key,value))
}

func (z *ZapLogger) PanicInt64(msg string, key string, value int64) {
	if !z.levelCheck.Check(PanicLevel){
		return
	}
	z.l.Error(msg,zap.Int64(key,value))
}

func (z *ZapLogger) PanicFloat32(msg string, key string, value float32) {
	if !z.levelCheck.Check(PanicLevel){
		return
	}
	z.l.Error(msg,zap.Float32(key,value))
}

func (z *ZapLogger) PanicFloat64(msg string, key string, value float64) {
	if !z.levelCheck.Check(PanicLevel){
		return
	}
	z.l.Error(msg,zap.Float64(key,value))
}

func (z *ZapLogger) PanicByteString(msg string, key string, value []byte) {
	if !z.levelCheck.Check(PanicLevel){
		return
	}
	z.l.Error(msg,zap.ByteString(key,value))
}

func (z *ZapLogger) WarnString(msg string, key, value string) {
	if !z.levelCheck.Check(WarnLevel){
		return
	}
	z.l.Error(msg,zap.String(key,value))
}

func (z *ZapLogger) WarnInt(msg string, key string, value int) {
	if !z.levelCheck.Check(WarnLevel){
		return
	}
	z.l.Error(msg,zap.Int(key,value))
}

func (z *ZapLogger) WarnInt64(msg string, key string, value int64) {
	if !z.levelCheck.Check(WarnLevel){
		return
	}
	z.l.Error(msg,zap.Int64(key,value))
}

func (z *ZapLogger) WarnFloat32(msg string, key string, value float32) {
	if !z.levelCheck.Check(WarnLevel){
		return
	}
	z.l.Error(msg,zap.Float32(key,value))
}

func (z *ZapLogger) WarnFloat64(msg string, key string, value float64) {
	if !z.levelCheck.Check(WarnLevel){
		return
	}
	z.l.Error(msg,zap.Float64(key,value))
}

func (z *ZapLogger) WarnByteString(msg string, key string, value []byte) {
	if !z.levelCheck.Check(WarnLevel){
		return
	}
	z.l.Error(msg,zap.ByteString(key,value))
}

