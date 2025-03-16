package logger_test

import (
	"bytes"
	"log/slog"
	"testing"

	qt "github.com/frankban/quicktest"

	"github.com/go-extras/go-kit/logger"
)

func TestPrintf(t *testing.T) {
	c := qt.New(t)

	var b bytes.Buffer
	h := slog.NewTextHandler(&b, &slog.HandlerOptions{
		ReplaceAttr: func(_groups []string, a slog.Attr) slog.Attr {
			if a.Key == "time" {
				return slog.Attr{
					Key:   "time",
					Value: slog.StringValue("now"),
				}
			}
			return a
		},
	})

	sl := logger.NewSlog(slog.New(h))
	sl.Printf("test %s", "message")

	c.Assert(b.String(), qt.Equals, "time=now level=INFO msg=\"test message\"\n")
}

func TestWithFields(t *testing.T) {
	c := qt.New(t)

	var b bytes.Buffer
	h := slog.NewTextHandler(&b, &slog.HandlerOptions{
		ReplaceAttr: func(_groups []string, a slog.Attr) slog.Attr {
			if a.Key == "time" {
				return slog.Attr{
					Key:   "time",
					Value: slog.StringValue("now"),
				}
			}
			return a
		},
	})

	sl := logger.NewSlog(slog.New(h))
	sl.WithFields(logger.SlogFields("mykey", "test key")).Print("test message")

	c.Assert(b.String(), qt.Equals, "time=now level=INFO msg=\"test message\" mykey=\"test key\"\n")
}

func TestWithField(t *testing.T) {
	c := qt.New(t)

	var b bytes.Buffer
	h := slog.NewTextHandler(&b, &slog.HandlerOptions{
		ReplaceAttr: func(_groups []string, a slog.Attr) slog.Attr {
			if a.Key == "time" {
				return slog.Attr{
					Key:   "time",
					Value: slog.StringValue("now"),
				}
			}
			return a
		},
	})

	sl := logger.NewSlog(slog.New(h))
	sl.WithField("mykey", "test key").Print("test message")

	c.Assert(b.String(), qt.Equals, "time=now level=INFO msg=\"test message\" mykey=\"test key\"\n")
}
