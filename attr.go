package attr

import (
	"log/slog"
)

// Attr represents a key-value pair.
type Attr struct {
	a slog.Attr
}

// New returns a new Attr.
func New(key string, value any) Attr {
	return Attr{
		a: slog.Any(key, value),
	}
}

// KV returns the key and value.
func (a Attr) KV() (string, any) {
	return a.a.Key, a.a.Value.Any()
}

// SlogAttr returns the underlying slog.Attr.
func (a Attr) SlogAttr() slog.Attr {
	return a.a
}

// String implements fmt.Stringer.
// It returns the string representation of the underlying slog.Attr.
func (a Attr) String() string {
	return a.a.String()
}
