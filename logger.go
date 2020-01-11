package abstractlogger

// Logger is the interface to abstract away a general purpose logger
// For performance reasons the interface is very specific and doesn't come with a "variadic Fields" argument as most logger libraries do.
// This is to make the Logger function call easy to inline for the compiler to not have any overhead using this Logger.
type Logger interface {
	DebugLogger
	InfoLogger
	WarnLogger
	ErrorLogger
	FatalLogger
	PanicLogger
}

// NoopLogger satisfies the Logger interface while doing nothing
var NoopLogger Logger = Noop{}

type DebugLogger interface {
	DebugString(msg string, key, value string)
	DebugInt(msg string, key string, value int)
	DebugInt64(msg string, key string, value int64)
	DebugFloat32(msg string, key string, value float32)
	DebugFloat64(msg string, key string, value float64)
	DebugByteString(msg string, key string, value []byte)
}

type ErrorLogger interface {
	ErrorString(msg string, key, value string)
	ErrorInt(msg string, key string, value int)
	ErrorInt64(msg string, key string, value int64)
	ErrorFloat32(msg string, key string, value float32)
	ErrorFloat64(msg string, key string, value float64)
	ErrorByteString(msg string, key string, value []byte)
}

type FatalLogger interface {
	FatalString(msg string, key, value string)
	FatalInt(msg string, key string, value int)
	FatalInt64(msg string, key string, value int64)
	FatalFloat32(msg string, key string, value float32)
	FatalFloat64(msg string, key string, value float64)
	FatalByteString(msg string, key string, value []byte)
}

type InfoLogger interface {
	InfoString(msg string, key, value string)
	InfoInt(msg string, key string, value int)
	InfoInt64(msg string, key string, value int64)
	InfoFloat32(msg string, key string, value float32)
	InfoFloat64(msg string, key string, value float64)
	InfoByteString(msg string, key string, value []byte)
}

type PanicLogger interface {
	PanicString(msg string, key, value string)
	PanicInt(msg string, key string, value int)
	PanicInt64(msg string, key string, value int64)
	PanicFloat32(msg string, key string, value float32)
	PanicFloat64(msg string, key string, value float64)
	PanicByteString(msg string, key string, value []byte)
}

type WarnLogger interface {
	WarnString(msg string, key, value string)
	WarnInt(msg string, key string, value int)
	WarnInt64(msg string, key string, value int64)
	WarnFloat32(msg string, key string, value float32)
	WarnFloat64(msg string, key string, value float64)
	WarnByteString(msg string, key string, value []byte)
}

type Noop struct {
	
}

func (_ Noop) InfoString(msg string, key, value string) {
	
}

func (_ Noop) InfoInt(msg string, key string, value int) {
	
}

func (_ Noop) InfoInt64(msg string, key string, value int64) {
	
}

func (_ Noop) InfoFloat32(msg string, key string, value float32) {
	
}

func (_ Noop) InfoFloat64(msg string, key string, value float64) {
	
}

func (_ Noop) InfoByteString(msg string, key string, value []byte) {
	
}

func (_ Noop) ErrorString(msg string, key, value string) {
	
}

func (_ Noop) ErrorInt(msg string, key string, value int) {
	
}

func (_ Noop) ErrorInt64(msg string, key string, value int64) {
	
}

func (_ Noop) ErrorFloat32(msg string, key string, value float32) {
	
}

func (_ Noop) ErrorFloat64(msg string, key string, value float64) {
	
}

func (_ Noop) ErrorByteString(msg string, key string, value []byte) {
	
}

func (_ Noop) FatalString(msg string, key, value string) {
	
}

func (_ Noop) FatalInt(msg string, key string, value int) {
	
}

func (_ Noop) FatalInt64(msg string, key string, value int64) {
	
}

func (_ Noop) FatalFloat32(msg string, key string, value float32) {
	
}

func (_ Noop) FatalFloat64(msg string, key string, value float64) {
	
}

func (_ Noop) FatalByteString(msg string, key string, value []byte) {
	
}

func (_ Noop) PanicString(msg string, key, value string) {
	
}

func (_ Noop) PanicInt(msg string, key string, value int) {
	
}

func (_ Noop) PanicInt64(msg string, key string, value int64) {
	
}

func (_ Noop) PanicFloat32(msg string, key string, value float32) {
	
}

func (_ Noop) PanicFloat64(msg string, key string, value float64) {
	
}

func (_ Noop) PanicByteString(msg string, key string, value []byte) {
	
}

func (_ Noop) WarnString(msg string, key, value string) {
	
}

func (_ Noop) WarnInt(msg string, key string, value int) {
	
}

func (_ Noop) WarnInt64(msg string, key string, value int64) {
	
}

func (_ Noop) WarnFloat32(msg string, key string, value float32) {
	
}

func (_ Noop) WarnFloat64(msg string, key string, value float64) {
	
}

func (_ Noop) WarnByteString(msg string, key string, value []byte) {
	
}

func (_ Noop) DebugString(msg string, key, value string) {
	
}

func (_ Noop) DebugInt(msg string, key string, value int) {
	
}

func (_ Noop) DebugInt64(msg string, key string, value int64) {
	
}

func (_ Noop) DebugFloat32(msg string, key string, value float32) {
	
}

func (_ Noop) DebugFloat64(msg string, key string, value float64) {
	
}

func (_ Noop) DebugByteString(msg string, key string, value []byte) {
	
}