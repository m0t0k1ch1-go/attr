package attr

import (
	"fmt"
	"log/slog"
	"math"
	"strconv"

	"github.com/getsentry/sentry-go"
	sentryattr "github.com/getsentry/sentry-go/attribute"
)

// Attr represents a key-value pair.
type Attr struct {
	k string
	v any
}

func new(k string, v any) Attr {
	return Attr{
		k: k,
		v: v,
	}
}

// Bool returns a new [Attr] with the given key and bool value.
func Bool(k string, v bool) Attr {
	return new(k, v)
}

// Int returns a new [Attr] with the given key and int value.
func Int(k string, v int) Attr {
	return new(k, v)
}

// Int64 returns a new [Attr] with the given key and int64 value.
func Int64(k string, v int64) Attr {
	return new(k, v)
}

// Uint64 returns a new [Attr] with the given key and uint64 value.
func Uint64(k string, v uint64) Attr {
	return new(k, v)
}

// Float64 returns a new [Attr] with the given key and float64 value.
func Float64(k string, v float64) Attr {
	return new(k, v)
}

// String returns a new [Attr] with the given key and string value.
func String(k string, v string) Attr {
	return new(k, v)
}

// Error returns a new [Attr] with the given key and error value.
func Error(k string, v error) Attr {
	return new(k, v)
}

// SentryLevel returns a new [Attr] with the given key and [sentry.Level] value.
func SentryLevel(k string, v sentry.Level) Attr {
	return new(k, v)
}

// K returns the key.
func (a Attr) K() string {
	return a.k
}

// V returns the value.
func (a Attr) V() any {
	return a.v
}

// KV returns the key and value.
func (a Attr) KV() (string, any) {
	return a.K(), a.V()
}

// String implements [fmt.Stringer].
// It returns the [Attr] as a string in the format "key=value".
func (a Attr) String() string {
	return a.k + "=" + a.valueString()
}

func (a Attr) valueString() string {
	switch v := a.v.(type) {
	case bool:
		return strconv.FormatBool(v)
	case int:
		return strconv.Itoa(v)
	case int64:
		return strconv.FormatInt(v, 10)
	case uint64:
		return strconv.FormatUint(v, 10)
	case float64:
		return strconv.FormatFloat(v, 'g', -1, 64)
	case string:
		return v
	case error:
		return v.Error()
	case sentry.Level:
		return string(v)
	default:
		return fmt.Sprintf("%v", v)
	}
}

// SlogAttr returns the [Attr] as a [slog.Attr].
func (a Attr) SlogAttr() slog.Attr {
	switch v := a.v.(type) {
	case bool:
		return slog.Bool(a.k, v)
	case int:
		return slog.Int(a.k, v)
	case int64:
		return slog.Int64(a.k, v)
	case uint64:
		return slog.Uint64(a.k, v)
	case float64:
		return slog.Float64(a.k, v)
	case string:
		return slog.String(a.k, v)
	default:
		return slog.Any(a.k, v)
	}
}

// SentryAttr returns the [Attr] as a [sentryattr.Builder].
func (a Attr) SentryAttr() sentryattr.Builder {
	switch v := a.v.(type) {
	case bool:
		return sentryattr.Bool(a.k, v)
	case int:
		return sentryattr.Int(a.k, v)
	case int64:
		return sentryattr.Int64(a.k, v)
	case uint64:
		if v <= math.MaxInt64 {
			return sentryattr.Int64(a.k, int64(v))
		} else {
			return sentryattr.String(a.k, a.valueString())
		}
	case float64:
		return sentryattr.Float64(a.k, v)
	case string:
		return sentryattr.String(a.k, v)
	default:
		return sentryattr.String(a.k, a.valueString())
	}
}
