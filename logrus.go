package abstractlogger

import (
	"github.com/sirupsen/logrus"
)

// LogrusLogger implements the Logger frontend using the popular logrus library as a backend
// It makes use of the LevelCheck helper to increase performance
type LogrusLogger struct {
	l          *logrus.Logger
	levelCheck LevelCheck
}

func NewLogrusLogger(l *logrus.Logger, level Level) *LogrusLogger {
	return &LogrusLogger{
		l:          l,
		levelCheck: NewLevelCheck(level),
	}
}

func (l *LogrusLogger) fields(fields []Field) logrus.Fields {
	out := make(logrus.Fields, len(fields))
	for i := range fields {
		switch fields[i].kind {
		case StringField:
			out[fields[i].key] = fields[i].stringValue
		case ByteStringField:
			out[fields[i].key] = string(fields[i].byteValue)
		case IntField:
			out[fields[i].key] = fields[i].intValue
		case BoolField:
			out[fields[i].key] = fields[i].intValue != 0
		case ErrorField,NamedErrorField:
			out[fields[i].key] = fields[i].errorValue
		case StringsField:
			out[fields[i].key] = fields[i].stringsValue
		default:
			out[fields[i].key] = fields[i].interfaceValue
		}
	}
	return out
}

func (l *LogrusLogger) Debug(msg string,fields ...Field) {
	if !l.levelCheck.Check(DebugLevel) {
		return
	}
	l.l.WithFields(l.fields(fields)).Debug(msg)
}

func (l *LogrusLogger) Info(msg string,fields ...Field) {
	if !l.levelCheck.Check(InfoLevel) {
		return
	}
	l.l.WithFields(l.fields(fields)).Info(msg)
}

func (l *LogrusLogger) Warn(msg string,fields ...Field) {
	if !l.levelCheck.Check(WarnLevel) {
		return
	}
	l.l.WithFields(l.fields(fields)).Warn(msg)
}


func (l *LogrusLogger) Error(msg string,fields ...Field) {
	if !l.levelCheck.Check(ErrorLevel) {
		return
	}
	l.l.WithFields(l.fields(fields)).Error(msg)
}

func (l *LogrusLogger) Fatal(msg string,fields ...Field) {
	if !l.levelCheck.Check(FatalLevel) {
		return
	}
	l.l.WithFields(l.fields(fields)).Fatal(msg)
}


func (l *LogrusLogger) Panic(msg string,fields ...Field) {
	if !l.levelCheck.Check(PanicLevel) {
		return
	}
	l.l.WithFields(l.fields(fields)).Panic(msg)
}