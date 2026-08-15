
package parallel

import (
	"errors"
	"testing"
)

func TestWaitPropagates(t *testing.T) {
	boom := errors.New("boom")
	err := Run(func() error { return nil }, func() error { return boom })
	if !errors.Is(err, boom) {
		t.Fatalf("want boom, got %v", err)
	}
}
