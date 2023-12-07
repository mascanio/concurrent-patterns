package concurrent_patterns

import (
	"testing"

	g "github.com/mascanio/generators"
)

func TestOrDone(t *testing.T) {
	done := make(chan struct{})
	defer close(done)
	inCh := g.Take(done, g.Repeat(done, 1), 3)
	orDoneCh := OrDone(done, inCh)

	i := 0
	for v := range orDoneCh {
		if v != 1 {
			t.Errorf("Expected value 1, got %d", v)
		}
		i++
	}
	if i != 3 {
		t.Errorf("Expected 3 repeats, got %d", i)
	}

	_, ok := <-orDoneCh
	if ok {
		t.Errorf("Expected channel to be closed")
	}
}

func TestOrDoneCloseDone(t *testing.T) {
	done := make(chan struct{})
	inCh := g.Take(done, g.Repeat(done, 1), 3)
	orDoneCh := OrDone(done, inCh)

	v, ok := <-orDoneCh
	if !ok {
		t.Errorf("Expected channel to be open")
	}
	if v != 1 {
		t.Errorf("Expected value 1, got %d", v)
	}
	close(done)
	_, ok = <-orDoneCh
	if ok {
		t.Errorf("Expected channel to be closed")
	}
}
