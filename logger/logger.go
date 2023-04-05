// Package logger provides interfaces for logging with various levels of verbosity and functionality.
//
// License: MIT
// Copyright: 2023, Denis Voytyuk
package logger

// PrimitiveLogger is the simplest logger interface that only provides the Printf and Print methods.
type PrimitiveLogger interface {
	Printf(format string, args ...any)
	Print(args ...any)
}

// BasicLogger is a logger interface that extends PrimitiveLogger and provides additional methods for logging fatal errors and panics.
type BasicLogger interface {
	PrimitiveLogger

	Fatalf(format string, args ...any)
	Panicf(format string, args ...any)

	Fatal(args ...any)
	Panic(args ...any)
}

// LevelLogger is a logger interface that extends BasicLogger and provides additional methods for logging messages with various levels of severity.
type LevelLogger interface {
	BasicLogger

	Debugf(format string, args ...any)
	Infof(format string, args ...any)
	Warnf(format string, args ...any)
	Warningf(format string, args ...any)
	Errorf(format string, args ...any)

	Debug(args ...any)
	Info(args ...any)
	Warn(args ...any)
	Warning(args ...any)
	Error(args ...any)
}

// TraceLogger is a logger interface that extends LevelLogger and provides additional methods for logging messages with the lowest level of severity.
type TraceLogger interface {
	LevelLogger

	Tracef(format string, args ...any)
	Trace(args ...any)
}

// FieldLogger is a logger interface that extends LevelLogger and provides additional methods for logging messages with structured data.
// For use with logrus, the T and U type parameters should be set to logrus.Fields and *logrus.Entry respectively.
type FieldLogger[T any, U any] interface {
	LevelLogger

	WithField(key string, value any) U
	WithFields(fields T) U
	WithError(err error) U
}

// TraceFieldLogger extends TraceLogger and FieldLogger with methods for adding fields to trace messages.
// For use with logrus, the T and U type parameters should be set to logrus.Fields and *logrus.Entry respectively.
type TraceFieldLogger[T any, U any] interface {
	TraceLogger
	FieldLogger[T, U]
}
