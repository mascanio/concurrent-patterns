package concurrent_patterns

// OrDone returns an output channel that returns the values from the inCh channel.
// If the done or inCh channels are closed, the returned channel is also closed.
func OrDone[T any](done <-chan struct{}, recvCh <-chan T) <-chan T {
	orDoneCh := make(chan T)
	go func() {
		defer close(orDoneCh)
		for {
			select {
			case <-done:
				return
			case v, ok := <-recvCh:
				if !ok {
					return
				}
				select {
				case <-done:
					return
				case orDoneCh <- v:
				}
			}
		}
	}()
	return orDoneCh
}
