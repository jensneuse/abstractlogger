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

// DebugLogger is for logging debug messages
type DebugLogger interface {
	// Debug is for logging debug information including variadic Fields
	// use the non variadic methods in case you'd like to trade usability for performance
	Debug(msg string,fields ...Field)
	DebugField(msg string, field Field)
	DebugField2(msg string, field1 Field, field2 Field)
	DebugField3(msg string, field1 Field, field2 Field, field3 Field)
	DebugField4(msg string, field1 Field, field2 Field, field3 Field, field4 Field)
	DebugField5(msg string, field1 Field, field2 Field, field3 Field, field4 Field, field5 Field)
	DebugField6(msg string, field1 Field, field2 Field, field3 Field, field4 Field, field5 Field, field6 Field)
}

// ErrorLogger is for logging errors
type ErrorLogger interface {
	// Error is for logging errors including variadic Fields
	// use the non variadic methods in case you'd like to trade usability for performance
	Error(msg string,fields ...Field)
	ErrorField(msg string, field Field)
	ErrorField2(msg string, field1 Field, field2 Field)
	ErrorField3(msg string, field1 Field, field2 Field, field3 Field)
	ErrorField4(msg string, field1 Field, field2 Field, field3 Field, field4 Field)
	ErrorField5(msg string, field1 Field, field2 Field, field3 Field, field4 Field, field5 Field)
	ErrorField6(msg string, field1 Field, field2 Field, field3 Field, field4 Field, field5 Field, field6 Field)
}

type FatalLogger interface {
	// Fatal is for logging a fatal error including variadic Fields
	// use the non variadic methods in case you'd like to trade usability for performance
	Fatal(msg string,fields ...Field)
	FatalField(msg string, field Field)
	FatalField2(msg string, field1 Field, field2 Field)
	FatalField3(msg string, field1 Field, field2 Field, field3 Field)
	FatalField4(msg string, field1 Field, field2 Field, field3 Field, field4 Field)
	FatalField5(msg string, field1 Field, field2 Field, field3 Field, field4 Field, field5 Field)
	FatalField6(msg string, field1 Field, field2 Field, field3 Field, field4 Field, field5 Field, field6 Field)
}

type InfoLogger interface {
	// Info is for logging info including variadic Fields
	// use the non variadic methods in case you'd like to trade usability for performance
	Info(msg string,fields ...Field)
	InfoField(msg string, field Field)
	InfoField2(msg string, field1 Field, field2 Field)
	InfoField3(msg string, field1 Field, field2 Field, field3 Field)
	InfoField4(msg string, field1 Field, field2 Field, field3 Field, field4 Field)
	InfoField5(msg string, field1 Field, field2 Field, field3 Field, field4 Field, field5 Field)
	InfoField6(msg string, field1 Field, field2 Field, field3 Field, field4 Field, field5 Field, field6 Field)
}

type PanicLogger interface {
	// Panic is for logging a panic including variadic Fields
	// use the non variadic methods in case you'd like to trade usability for performance
	Panic(msg string,fields ...Field)
	PanicField(msg string, field Field)
	PanicField2(msg string, field1 Field, field2 Field)
	PanicField3(msg string, field1 Field, field2 Field, field3 Field)
	PanicField4(msg string, field1 Field, field2 Field, field3 Field, field4 Field)
	PanicField5(msg string, field1 Field, field2 Field, field3 Field, field4 Field, field5 Field)
	PanicField6(msg string, field1 Field, field2 Field, field3 Field, field4 Field, field5 Field, field6 Field)
}

type WarnLogger interface {
	// Warn is for logging warnings including variadic Fields
	// use the non variadic methods in case you'd like to trade usability for performance
	Warn(msg string,fields ...Field)
	WarnField(msg string, field Field)
	WarnField2(msg string, field1 Field, field2 Field)
	WarnField3(msg string, field1 Field, field2 Field, field3 Field)
	WarnField4(msg string, field1 Field, field2 Field, field3 Field, field4 Field)
	WarnField5(msg string, field1 Field, field2 Field, field3 Field, field4 Field, field5 Field)
	WarnField6(msg string, field1 Field, field2 Field, field3 Field, field4 Field, field5 Field, field6 Field)
}

type Noop struct {}

func (_ Noop) Debug(msg string, fields ...Field) {
	
}

func (_ Noop) DebugField(msg string, field Field) {
	
}

func (_ Noop) DebugField2(msg string, field1 Field, field2 Field) {
	
}

func (_ Noop) DebugField3(msg string, field1 Field, field2 Field, field3 Field) {
	
}

func (_ Noop) DebugField4(msg string, field1 Field, field2 Field, field3 Field, field4 Field) {
	
}

func (_ Noop) DebugField5(msg string, field1 Field, field2 Field, field3 Field, field4 Field, field5 Field) {
	
}

func (_ Noop) DebugField6(msg string, field1 Field, field2 Field, field3 Field, field4 Field, field5 Field, field6 Field) {
	
}

func (_ Noop) Info(msg string, fields ...Field) {
	
}

func (_ Noop) InfoField(msg string, field Field) {
	
}

func (_ Noop) InfoField2(msg string, field1 Field, field2 Field) {
	
}

func (_ Noop) InfoField3(msg string, field1 Field, field2 Field, field3 Field) {
	
}

func (_ Noop) InfoField4(msg string, field1 Field, field2 Field, field3 Field, field4 Field) {
	
}

func (_ Noop) InfoField5(msg string, field1 Field, field2 Field, field3 Field, field4 Field, field5 Field) {
	
}

func (_ Noop) InfoField6(msg string, field1 Field, field2 Field, field3 Field, field4 Field, field5 Field, field6 Field) {
	
}

func (_ Noop) Warn(msg string, fields ...Field) {
	
}

func (_ Noop) WarnField(msg string, field Field) {
	
}

func (_ Noop) WarnField2(msg string, field1 Field, field2 Field) {
	
}

func (_ Noop) WarnField3(msg string, field1 Field, field2 Field, field3 Field) {
	
}

func (_ Noop) WarnField4(msg string, field1 Field, field2 Field, field3 Field, field4 Field) {
	
}

func (_ Noop) WarnField5(msg string, field1 Field, field2 Field, field3 Field, field4 Field, field5 Field) {
	
}

func (_ Noop) WarnField6(msg string, field1 Field, field2 Field, field3 Field, field4 Field, field5 Field, field6 Field) {
	
}

func (_ Noop) Error(msg string, fields ...Field) {
	
}

func (_ Noop) ErrorField(msg string, field Field) {
	
}

func (_ Noop) ErrorField2(msg string, field1 Field, field2 Field) {
	
}

func (_ Noop) ErrorField3(msg string, field1 Field, field2 Field, field3 Field) {
	
}

func (_ Noop) ErrorField4(msg string, field1 Field, field2 Field, field3 Field, field4 Field) {
	
}

func (_ Noop) ErrorField5(msg string, field1 Field, field2 Field, field3 Field, field4 Field, field5 Field) {
	
}

func (_ Noop) ErrorField6(msg string, field1 Field, field2 Field, field3 Field, field4 Field, field5 Field, field6 Field) {
	
}

func (_ Noop) Fatal(msg string, fields ...Field) {
	
}

func (_ Noop) FatalField(msg string, field Field) {
	
}

func (_ Noop) FatalField2(msg string, field1 Field, field2 Field) {
	
}

func (_ Noop) FatalField3(msg string, field1 Field, field2 Field, field3 Field) {
	
}

func (_ Noop) FatalField4(msg string, field1 Field, field2 Field, field3 Field, field4 Field) {
	
}

func (_ Noop) FatalField5(msg string, field1 Field, field2 Field, field3 Field, field4 Field, field5 Field) {
	
}

func (_ Noop) FatalField6(msg string, field1 Field, field2 Field, field3 Field, field4 Field, field5 Field, field6 Field) {
	
}

func (_ Noop) Panic(msg string, fields ...Field) {
	
}

func (_ Noop) PanicField(msg string, field Field) {
	
}

func (_ Noop) PanicField2(msg string, field1 Field, field2 Field) {
	
}

func (_ Noop) PanicField3(msg string, field1 Field, field2 Field, field3 Field) {
	
}

func (_ Noop) PanicField4(msg string, field1 Field, field2 Field, field3 Field, field4 Field) {
	
}

func (_ Noop) PanicField5(msg string, field1 Field, field2 Field, field3 Field, field4 Field, field5 Field) {
	
}

func (_ Noop) PanicField6(msg string, field1 Field, field2 Field, field3 Field, field4 Field, field5 Field, field6 Field) {
	
}
