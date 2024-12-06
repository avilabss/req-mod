package req

import (
	"testing"

	"github.com/avilabss/req-mod/internal/tests"
)

func TestPeekDrain(t *testing.T) {
	a := autoDecodeReadCloser{peek: []byte("test")}
	p := make([]byte, 2)
	n, _ := a.peekDrain(p)
	tests.AssertEqual(t, 2, n)
	tests.AssertEqual(t, true, a.peek != nil)
	n, _ = a.peekDrain(p)
	tests.AssertEqual(t, 2, n)
	tests.AssertEqual(t, true, a.peek == nil)
}
