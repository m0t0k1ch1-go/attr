package attr_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/m0t0k1ch1-go/attr"
)

func TestAttr(t *testing.T) {
	var a attr.Attr
	require.Implements(t, (*fmt.Stringer)(nil), &a)
}
