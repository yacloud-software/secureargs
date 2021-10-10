package client

import (
	"flag"
	"fmt"
	"golang.conradwood.net/go-easyops/prometheus"
	pb "golang.conradwood.net/apis/registry"
	"golang.conradwood.net/go-easyops/cmdline"
	"golang.conradwood.net/go-easyops/common"
	pp "golang.conradwood.net/go-easyops/profiling"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/naming"
	"math/rand"
	"sync"
	"time"
)

// we keep a local "mirror" copy of the registry
// so that, in case the registry goes away for some time
// we can still do our load balancing.
// we also do that because we need to regularly requery the registry
// to find updates to our target locations.

// to enable periodic updates, call notifyTargetChanges(servicename/channel)
// and this thing will inform the channel of changes

// notifyTargetChanges() will block until at least one target becomes
// available.
// This should prevent services from coming up and immediately failing
// because they're missing a backend.
// this might cause trouble during testing, if so a variable is provided to
// change that behavior

var (
	normal_sleep_time = flag.Int("ge_dialer_sleep_time", 20, "interval in seconds before querying the registry for changes (should be lower than ge_max_block)")
	log_on_block      = flag.Bool("ge_dial_log_on_block", true, "set this to false if you want no log messages on blocking behaviour [deprecated].")
	randomer          = rand.New(rand.NewSource(time.Now().Unix()))
	targetNotifiers   []*targetNotifier
	notifierLock      sync.Mutex
	blockCtr          = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "grpc_loadbalancer_no_connection",
			Help: "counter incremented each time the loadbalancer has no instances",
		},
		[]string{"servicename"},
	)

	failedQueryCtr = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "grpc_loadbalancer_registry_failures",
			Help: "counter incremented each time the loadbalancer fails to query the registry",
		},
		[]string{},
	)

	// will be lazily initialised...
	registryClient pb.RegistryClient
)

const (
	// if we have a long rpc call (e.g. 3 seconds)
	// then we can run into timeouts, during the
	// interceptor auth phase
	// imho - the auth should be handled by "normal"
	// loadbalancer function
	// (it seems) cnw 19/5/2018
	CONST_CALL_TIMEOUT = 4
)

type targetNotifier struct {
	serviceName string
	lock        sync.Mutex // lock updates to watchers
	watchers    []*targetWatcher
	blocked     bool
}

// true if ALL watchers are closed.
// false if at least one watcher is open
func (t *targetNotifier) IsActive() bool {
	t.lock.Lock()
	defer t.lock.Unlock()
	for _, x := range t.watchers {
		if !x.closed {
			return true
		}
	}
	return false
}

func GetOrCreateNotifier(service string) *targetNotifier {
	notifierLock.Lock()
	defer notifierLock.Unlock()
	for _, tn := range targetNotifiers {
		if tn.serviceName == service {
			return tn
		}
	}
	common.AddServiceName(service)
	tn := &targetNotifier{serviceName: service}
	targetNotifiers = append(targetNotifiers, tn)
	go tn.requery(true)
	return tn
}

func notifyTargetChanges(ch *targetWatcher) error {
	if *dialer_debug {
		fmt.Printf("[go-easyops] New notifier requested for %s\n", ch.serviceName)
	}
	tn := GetOrCreateNotifier(ch.serviceName)
	tn.AddWatcher(ch)
	tn.requery(false)
	return nil
}

// get a list of instances from registry
func queryForActiveInstances(serviceName string) ([]string, error) {
	totalQueryCtr.With(prometheus.Labels{"servicename": serviceName}).Inc()
	if *dialer_debug {
		fmt.Printf("[go-easyops] Resolving service address \"%s\" via registry %s...\n", serviceName, cmdline.GetClientRegistryAddress())
	}
	var err error
	var serverAddresses []string

	// check & create if nessary the registryClient
	// we only want to Dial() at startup
	// we need no lock, because, in the worst case we dial twice, so what?
	if registryClient == nil {
		registryAddress := cmdline.GetClientRegistryAddress()
		if *dialer_debug {
			fmt.Printf("[go-easyops] Dialing via registry at %s...\n", registryAddress)
		}

		conn, err := grpc.Dial(
			registryAddress,
			grpc.WithInsecure(),
			//			grpc.WithUnaryInterceptor(unaryClientInterceptor),
			//			grpc.WithStreamInterceptor(unaryStreamInterceptor),
			grpc.WithTimeout(time.Duration(CONST_CALL_TIMEOUT)*time.Second),
		)
		if err != nil {
			if *dialer_debug {
				fmt.Printf("Failed to connect to registry at %s\n", cmdline.GetClientRegistryAddress())
			}
			return nil, err
		}
		registryClient = pb.NewRegistryClient(conn)
	}
	list, err := registryClient.GetTarget(
		context.Background(),
		&pb.GetTargetRequest{
			Name:    serviceName,
			ApiType: pb.Apitype_grpc,
		},
	)

	// error getting stuff from registry
	if err != nil {
		if *dialer_debug {
			fmt.Printf("[go-easyops] error retrieving hosts for %s: %s\n", serviceName, err)
		}
		return nil, err
	}

	// we are only interested in the addresses in form of 'host:port', because that's what the grpc LB wants
	// so we're converting them here...
	for _, getResponse := range list.Service {
		location := getResponse.Location
		if location != nil {
			for _, svcaddr := range location.Address {
				x := fmt.Sprintf("%s:%d", svcaddr.Host, svcaddr.Port)
				serverAddresses = append(serverAddresses, x)
			}
		}
	}
	if *dialer_debug {
		fmt.Printf("[go-easyops] %d Instances of %s: %s\n", len(serverAddresses), serviceName, serverAddresses)
	}

	return serverAddresses, nil

}

// this thing runs in the background, one per servicename
func (tn *targetNotifier) requery(repeat bool) {
	for {
		for !tn.IsActive() {
			time.Sleep(time.Duration(5) * time.Second)
		}
		tn.lock.Lock()
		serverAddresses, err := queryForActiveInstances(tn.serviceName)
		if err != nil {
			failedQueryCtr.With(prometheus.Labels{}).Inc()
			// we will NOT update the local copy of registry - we keep going...
			tn.lock.Unlock()
			if !repeat {
				// never ever exit this thread if we are in 'repeat' mode
				return
			}
			// attempt requery 5 seconds later
			time.Sleep(time.Duration(randomer.Intn(5)) * time.Second)
			continue

		}
		lastBlock := tn.blocked
		if len(serverAddresses) == 0 {
			blockCtr.With(prometheus.Labels{
				"servicename": tn.serviceName,
			}).Inc()
			tn.blocked = true
		} else {
			tn.blocked = false
		}
		newBlock := lastBlock != tn.blocked
		if *log_on_block && newBlock {
			if tn.blocked {
				fmt.Printf("[go-easyops] WARNING load balancer blocked: no instance found for service %s\n", tn.serviceName)
			} else {
				fmt.Printf("[go-easyops] CANCEL WARNING load balancer unblocked: %d instances found for service %s\n", len(serverAddresses), tn.serviceName)
			}
		}

		// watchers can change, so take a snapshot whilst lock is held
		curWatchers := make([]*targetWatcher, len(tn.watchers))
		i := 0
		for _, w := range tn.watchers {
			curWatchers[i] = w
			i++
		}
		tn.lock.Unlock()
		// use snapshot to send updates to each watcher
		for _, watcher := range curWatchers {
			sendDiff(serverAddresses, watcher)
		}
		if !repeat {
			return
		}
		if tn.blocked {
			common.AddBlockedServiceName(tn.serviceName)
			time.Sleep(time.Duration(1+randomer.Intn(1)) * time.Second)
		} else {
			common.RemoveBlockedServiceName(tn.serviceName)
			time.Sleep(time.Duration(1+randomer.Intn(*normal_sleep_time)) * time.Second)
		}
	}
}

// work out changes (compared to this particular watcher)
// we have a list of serveraddresses from most recent lookup and a watcher to update
// each watcher has its own list of what it has been informed previously, and since
// we can only send ADD/REMOVE events (not a whole list) we got to remember what we sent previously.
// (tw.currentTargets is what WE think we sent to the watcher)
func sendDiff(serviceAddresses []string, tw *targetWatcher) {
	if tw.closed { // nothing to do on closed watchers
		return
	}
	// take care of NEW services
	var ups []*naming.Update
	for _, sa := range serviceAddresses {
		if !isInArray(tw.currentTargets, sa) {
			ups = append(ups, &naming.Update{Op: naming.Add, Addr: sa, Metadata: ""})
			if *dialer_debug {
				fmt.Printf("[go-easyops] notifying watcher of NEW target %s\n", sa)
			}
			tw.currentTargets = append(tw.currentTargets, sa)
		}
	}

	// take care of REMOVED services
	var nts []string
	for _, sa := range tw.currentTargets {
		if !isInArray(serviceAddresses, sa) {
			ups = append(ups, &naming.Update{Op: naming.Delete, Addr: sa, Metadata: ""})
			if *dialer_debug {
				fmt.Printf("[go-easyops] notifying watcher of REMOVED target %s\n", sa)
			}
		} else {
			nts = append(nts, sa)
		}
	}
	tw.currentTargets = nts
	// this should be truly async I guess
	tw.updates <- ups

}

func isInArray(s []string, find string) bool {
	for _, x := range s {
		if x == find {
			return true
		}
	}
	return false
}

func (tn *targetNotifier) AddWatcher(watcher *targetWatcher) {
	tn.lock.Lock()
	tn.watchers = append(tn.watchers, watcher)
	tn.lock.Unlock()
}
func unaryClientInterceptor(
	ctx context.Context,
	method string,
	req interface{},
	reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	pp.ClientRpcEntered()
	err := invoker(ctx, method, req, reply, cc, opts...)
	pp.ClientRpcDone()
	return err
}
func unaryStreamInterceptor(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	pp.ClientRpcEntered()
	cs, err := streamer(ctx, desc, cc, method, opts...)
	pp.ClientRpcDone()
	return cs, err
}
