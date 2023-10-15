// Package logger.
//
// License: MIT
// Copyright: 2023, Denis Voytyuk
package logger

import (
	"context"
	"fmt"
	"log/slog"
	"os"
)

const errKey = "error"

const SlogLevelFatal slog.Level = 100

var _ FieldLogger[[]any, *Slog] = (*Slog)(nil)

func NewSlog(slog *slog.Logger) *Slog {
	return &Slog{
		Logger: slog,
	}
}

// Slog is a wrapper around slog.Logger that implements FieldLogger.
// It is intended to be used as a logrus migration path.
// It is not intended to be used as a general purpose logger.
// Note, this is an experimental approach and it's not recommended to be used
// in production.
//
// To migrate from logrus do the following steps:
//  1. Replace all logrus implementation references with logger.FieldLogger
//     (or any other suitable interface from this package)
//  2. Replace all logrus.New() calls with logger.NewSlog(slog.New())
//  3. Replace all calls to logrus.* with logger.* (e.g. logrus.WithField -> logger.WithField)
//  4. You will have to manually adjust remaining incopatibilities (e.g. this package does not
//     support logrus.Fields and they have to be replaced with logger.SlogFields).
//
// As this struct is a wrapper around slog.Logger, it is possible to use slog.Logger methods.
type Slog struct {
	*slog.Logger
}

func (s *Slog) clone() *Slog {
	r := *s
	return &r
}

func (s *Slog) Printf(format string, args ...any) {
	s.Logger.Log(
		context.Background(),
		slog.LevelInfo,
		fmt.Sprintf(format, args...),
	)
}

func (s *Slog) Print(args ...any) {
	s.Logger.Log(
		context.Background(),
		slog.LevelInfo,
		fmt.Sprint(args...),
	)
}

func (s *Slog) Fatalf(format string, args ...any) {
	s.Logger.Log(
		context.Background(),
		SlogLevelFatal,
		fmt.Sprintf(format, args...),
	)
	os.Exit(1)
}

func (s *Slog) Panicf(format string, args ...any) {
	rec := fmt.Sprintf(format, args...) // TODO: it is possible to catch the record in handler
	s.Logger.Log(
		context.Background(),
		SlogLevelFatal,
		rec,
	)
	panic(rec)
}

func (s *Slog) Fatal(args ...any) {
	s.Logger.Log(
		context.Background(),
		SlogLevelFatal,
		fmt.Sprint(args...),
	)
	os.Exit(1)
}

func (s *Slog) Panic(args ...any) {
	rec := fmt.Sprint(args...) // TODO: it is possible to catch the record in handler
	s.Logger.Log(
		context.Background(),
		SlogLevelFatal,
		rec,
	)
	panic(rec)
}

func (s *Slog) Debugf(format string, args ...any) {
	s.Logger.Log(
		context.Background(),
		slog.LevelDebug,
		fmt.Sprintf(format, args...),
	)
}

func (s *Slog) Infof(format string, args ...any) {
	s.Logger.Log(
		context.Background(),
		slog.LevelInfo,
		fmt.Sprintf(format, args...),
	)
}

func (s *Slog) Warnf(format string, args ...any) {
	s.Logger.Log(
		context.Background(),
		slog.LevelWarn,
		fmt.Sprintf(format, args...),
	)
}

func (s *Slog) Warningf(format string, args ...any) {
	s.Logger.Log(
		context.Background(),
		slog.LevelWarn,
		fmt.Sprintf(format, args...),
	)
}

func (s *Slog) Errorf(format string, args ...any) {
	s.Logger.Log(
		context.Background(),
		slog.LevelError,
		fmt.Sprintf(format, args...),
	)
}

func (s *Slog) Debug(args ...any) {
	s.Logger.Log(
		context.Background(),
		slog.LevelDebug,
		fmt.Sprint(args...),
	)
}

func (s *Slog) Info(args ...any) {
	s.Logger.Log(
		context.Background(),
		slog.LevelInfo,
		fmt.Sprint(args...),
	)
}

func (s *Slog) Warn(args ...any) {
	s.Logger.Log(
		context.Background(),
		slog.LevelWarn,
		fmt.Sprint(args...),
	)
}

func (s *Slog) Warning(args ...any) {
	s.Logger.Log(
		context.Background(),
		slog.LevelWarn,
		fmt.Sprint(args...),
	)
}

func (s *Slog) Error(args ...any) {
	s.Logger.Log(
		context.Background(),
		slog.LevelError,
		fmt.Sprint(args...),
	)
}

func (s *Slog) WithField(key string, value any) *Slog {
	r := s.clone()
	r.Logger = r.With(key, value)
	return r
}

func (s *Slog) WithFields(fields []any) *Slog {
	r := s.clone()
	r.Logger = r.With(fields...)
	return r
}

func (s *Slog) WithError(err error) *Slog {
	r := s.clone()
	r.Logger = r.With(errKey, err)
	return r
}

// SlogFields is a helper function that converts a list of key-value pairs to a slice of them.
// Use this function in conjuction with WithFields to pass a slice of key-value pairs to it.
func SlogFields(args ...any) []any {
	return args
}

func Printf(format string, args ...any) {
	slog.Log(
		context.Background(),
		slog.LevelInfo,
		fmt.Sprintf(format, args...),
	)
}

func Print(args ...any) {
	slog.Log(
		context.Background(),
		slog.LevelInfo,
		fmt.Sprint(args...),
	)
}

func Fatalf(format string, args ...any) {
	slog.Log(
		context.Background(),
		SlogLevelFatal,
		fmt.Sprintf(format, args...),
	)
	os.Exit(1)
}

func Panicf(format string, args ...any) {
	rec := fmt.Sprintf(format, args...) // TODO: it is possible to catch the record in handler
	slog.Log(
		context.Background(),
		SlogLevelFatal,
		rec,
	)
	panic(rec)
}

func Fatal(args ...any) {
	slog.Log(
		context.Background(),
		SlogLevelFatal,
		fmt.Sprint(args...),
	)
	os.Exit(1)
}

func Panic(args ...any) {
	rec := fmt.Sprint(args...) // TODO: it is possible to catch the record in handler
	slog.Log(
		context.Background(),
		SlogLevelFatal,
		rec,
	)
	panic(rec)
}

func Debugf(format string, args ...any) {
	slog.Log(
		context.Background(),
		slog.LevelDebug,
		fmt.Sprintf(format, args...),
	)
}

func Infof(format string, args ...any) {
	slog.Log(
		context.Background(),
		slog.LevelInfo,
		fmt.Sprintf(format, args...),
	)
}

func Warnf(format string, args ...any) {
	slog.Log(
		context.Background(),
		slog.LevelWarn,
		fmt.Sprintf(format, args...),
	)
}

func Warningf(format string, args ...any) {
	slog.Log(
		context.Background(),
		slog.LevelWarn,
		fmt.Sprintf(format, args...),
	)
}

func Errorf(format string, args ...any) {
	slog.Log(
		context.Background(),
		slog.LevelError,
		fmt.Sprintf(format, args...),
	)
}

func Debug(args ...any) {
	slog.Log(
		context.Background(),
		slog.LevelDebug,
		fmt.Sprint(args...),
	)
}

func Info(args ...any) {
	slog.Log(
		context.Background(),
		slog.LevelInfo,
		fmt.Sprint(args...),
	)
}

func Warn(args ...any) {
	slog.Log(
		context.Background(),
		slog.LevelWarn,
		fmt.Sprint(args...),
	)
}

func Warning(args ...any) {
	slog.Log(
		context.Background(),
		slog.LevelWarn,
		fmt.Sprint(args...),
	)
}

func Error(args ...any) {
	slog.Log(
		context.Background(),
		slog.LevelError,
		fmt.Sprint(args...),
	)
}

func WithField(key string, value any) *Slog {
	return NewSlog(slog.With(key, value))
}

func WithFields(fields []any) *Slog {
	return NewSlog(slog.With(fields...))
}

func WithError(err error) *Slog {
	return NewSlog(slog.With(errKey, err))
}
