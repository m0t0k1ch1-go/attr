package attr_test

import (
	"errors"
	"fmt"
	"log/slog"
	"math"
	"strconv"
	"testing"

	"github.com/getsentry/sentry-go"
	sentryattr "github.com/getsentry/sentry-go/attribute"
	"github.com/stretchr/testify/require"

	"github.com/m0t0k1ch1-go/attr"
)

func TestAttr(t *testing.T) {
	var a attr.Attr
	require.Implements(t, (*fmt.Stringer)(nil), &a)
}

func TestAttr_KV(t *testing.T) {
	type output struct {
		k string
		v any
	}

	tcs := []struct {
		name string
		in   attr.Attr
		want output
	}{
		{
			"bool",
			attr.Bool("bool", true),
			output{"bool", true},
		},
		{
			"int",
			attr.Int("int", 1),
			output{"int", 1},
		},
		{
			"int64",
			attr.Int64("int64", 1),
			output{"int64", int64(1)},
		},
		{
			"uint64",
			attr.Uint64("uint64", 1),
			output{"uint64", uint64(1)},
		},
		{
			"float64",
			attr.Float64("float64", 1.1),
			output{"float64", float64(1.1)},
		},
		{
			"string",
			attr.String("string", "m0t0k1ch1"),
			output{"string", "m0t0k1ch1"},
		},
		{
			"error",
			attr.Error("error", errors.New("something went wrong")),
			output{"error", errors.New("something went wrong")},
		},
		{
			"sentry.Level",
			attr.SentryLevel("sentry.level", sentry.LevelDebug),
			output{"sentry.level", sentry.LevelDebug},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			k, v := tc.in.KV()
			require.Equal(t, tc.want.k, k)
			require.Equal(t, tc.want.v, v)
		})
	}
}

func TestAttr_String(t *testing.T) {
	tcs := []struct {
		name string
		in   attr.Attr
		want string
	}{
		{
			"bool",
			attr.Bool("bool", true),
			"bool=true",
		},
		{
			"int",
			attr.Int("int", 1),
			"int=1",
		},
		{
			"int64",
			attr.Int64("int64", 1),
			"int64=1",
		},
		{
			"uint64",
			attr.Uint64("uint64", 1),
			"uint64=1",
		},
		{
			"float64",
			attr.Float64("float64", 1.1),
			"float64=1.1",
		},
		{
			"string",
			attr.String("string", "m0t0k1ch1"),
			"string=m0t0k1ch1",
		},
		{
			"error",
			attr.Error("error", errors.New("something went wrong")),
			"error=something went wrong",
		},
		{
			"sentry.Level",
			attr.SentryLevel("sentry.level", sentry.LevelDebug),
			"sentry.level=debug",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.want, tc.in.String())
		})
	}
}

func TestAttr_SlogAttr(t *testing.T) {
	tcs := []struct {
		name string
		in   attr.Attr
		want slog.Attr
	}{
		{
			"bool",
			attr.Bool("bool", true),
			slog.Bool("bool", true),
		},
		{
			"int",
			attr.Int("int", 1),
			slog.Int("int", 1),
		},
		{
			"int64",
			attr.Int64("int64", 1),
			slog.Int64("int64", 1),
		},
		{
			"uint64",
			attr.Uint64("uint64", 1),
			slog.Uint64("uint64", 1),
		},
		{
			"float64",
			attr.Float64("float64", 1.1),
			slog.Float64("float64", 1.1),
		},
		{
			"string",
			attr.String("string", "m0t0k1ch1"),
			slog.String("string", "m0t0k1ch1"),
		},
		{
			"error",
			attr.Error("error", errors.New("something went wrong")),
			slog.Any("error", errors.New("something went wrong")),
		},
		{
			"sentry.Level",
			attr.SentryLevel("sentry.level", sentry.LevelDebug),
			slog.Any("sentry.level", sentry.LevelDebug),
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.want, tc.in.SlogAttr())
		})
	}
}

func TestAttr_SentryAttr(t *testing.T) {
	tcs := []struct {
		name string
		in   attr.Attr
		want sentryattr.Builder
	}{
		{
			"bool",
			attr.Bool("bool", true),
			sentryattr.Bool("bool", true),
		},
		{
			"int",
			attr.Int("int", 1),
			sentryattr.Int("int", 1),
		},
		{
			"int64",
			attr.Int64("int64", 1),
			sentryattr.Int64("int64", 1),
		},
		{
			"uint64: max int64",
			attr.Uint64("uint64", math.MaxInt64),
			sentryattr.Int64("uint64", math.MaxInt64),
		},
		{
			"uint64: max int64 + 1",
			attr.Uint64("uint64", math.MaxInt64+1),
			sentryattr.String("uint64", strconv.FormatUint(math.MaxInt64+1, 10)),
		},
		{
			"float64",
			attr.Float64("float64", 1.1),
			sentryattr.Float64("float64", 1.1),
		},
		{
			"string",
			attr.String("string", "m0t0k1ch1"),
			sentryattr.String("string", "m0t0k1ch1"),
		},
		{
			"error",
			attr.Error("error", errors.New("something went wrong")),
			sentryattr.String("error", "something went wrong"),
		},
		{
			"sentry.Level",
			attr.SentryLevel("sentry.level", sentry.LevelDebug),
			sentryattr.String("sentry.level", string(sentry.LevelDebug)),
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.want, tc.in.SentryAttr())
		})
	}
}
