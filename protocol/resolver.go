package protocol

import (
	"net"
	"strconv"
	"strings"
	"sync"

	consul_api "github.com/hashicorp/consul/api"
	"google.golang.org/grpc/resolver"
	//log "github.com/sirupsen/logrus"
)

const scheme = "consul"

type consulResolverBuilder struct{}

func (*consulResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {

	var addr, service string
	if ss := strings.Split(target.Endpoint, "/"); len(ss) >= 2 {
		addr, service = ss[0], ss[1]
	} else {
		addr = target.Endpoint
	}

	config := consul_api.DefaultConfig()
	config.Address = addr

	client, err := consul_api.NewClient(config)
	if err != nil {
		return nil, err
	}

	r := &consulResolver{
		target:       target,
		cc:           cc,
		consulClient: client,
		service:      service,
	}
	r.start()
	return r, nil
}

func (*consulResolverBuilder) Scheme() string {
	return scheme
}

type consulResolver struct {
	target resolver.Target
	cc     resolver.ClientConn

	consulClient *consul_api.Client
	service      string
	lastIndex    uint64

	done bool
	wg   sync.WaitGroup
}

func (r *consulResolver) ResolveNow(o resolver.ResolveNowOptions) {
	r.resolveOnce(o)
}

func (r *consulResolver) Close() {
	r.done = true
	//r.wg.Wait()
}

func (r *consulResolver) start() {
	r.done = false
	go r.watchAddrUpdates()
}

func (r *consulResolver) watchAddrUpdates() {
	r.wg.Add(1)
	defer r.wg.Done()

	o := resolver.ResolveNowOptions{}
	for r.done != true {
		r.resolveOnce(o)
	}
}

func (r *consulResolver) resolveOnce(o resolver.ResolveNowOptions) {
	services, meta, err := r.consulClient.Health().Service(r.service, "", true, &consul_api.QueryOptions{
		WaitIndex: r.lastIndex,
	})
	if err != nil {
		return
	}
	r.lastIndex = meta.LastIndex
	var newAddrs []resolver.Address
	for _, s := range services {
		host := net.JoinHostPort(s.Service.Address, strconv.Itoa(s.Service.Port))
		newAddrs = append(newAddrs, resolver.Address{Addr: host, Metadata: s.Node})
		//log.Printf("new host:%#v, node:%#v", host, s.Node)
	}
	if len(newAddrs) > 0 {
		r.cc.NewAddress(newAddrs)
	}
}

func init() {
	resolver.Register(&consulResolverBuilder{})
}
