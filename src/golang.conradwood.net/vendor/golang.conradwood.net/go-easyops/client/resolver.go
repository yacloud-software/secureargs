package client

import (
	"google.golang.org/grpc/naming"
)

type RegistryResolver struct{}

// Conceptually, the grpc LB calls 'resolve' when it starts up. It expects a 'watcher'
// (implemented in 'targetWatcher' below)
// the grpc LB blocks on calls to Next() of the watcher.
// The watcher is expected to deliver updates (ADD/REMOVE targets) whenever it returns from Next()
// the grpc LB updates its list of targets and routes RPC calls accordingly

// the naming.Resolve (which apparently is deprecated and resolver.XX should be used, which I cannot find?)
func (r *RegistryResolver) Resolve(serviceName string) (naming.Watcher, error) {
	//panic("wtf")

	// TODO: investigate chan size and blocking behaviour. is is safe to block on this channel?
	var ch chan []*naming.Update = make(chan []*naming.Update, 1)
	sw := &targetWatcher{closed: false, updates: ch, serviceName: serviceName}
	notifyTargetChanges(sw)
	return sw, nil
}

type targetWatcher struct {
	closed         bool
	updates        chan []*naming.Update
	currentTargets []string
	serviceName    string
}

func (w *targetWatcher) Next() ([]*naming.Update, error) {
	return <-w.updates, nil
}

// cnw: I suspect this is a memoryleak. we close and never reuse them.
// there is no way of "open" again. But on the other hand, it is unclear
// at which point we may remove this channel, isn't it?
// without this mechanism the result is a panic() on close() connection,
// so I think this is better than without
func (w *targetWatcher) Close() {
	w.closed = true
	close(w.updates)
}
