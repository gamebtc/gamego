package protocol

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/dodoZeng/grpclb/balancer/ketama"
	"github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"github.com/hashicorp/consul/api"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/codes"
)

type ServicePool interface {
	GetService(id int32) *grpc.ClientConn
	GetValue(k string) []byte
	SetValue(k string, v []byte) bool
}

// all services
type service_pool struct {
	sync.RWMutex
	addr     string
	prefix   string
	services map[int32]*grpc.ClientConn
	retry    bool
}

func NewServicePool(addr string, prefix string, services []int32, retry bool) ServicePool {
	p := new(service_pool)
	p.addr = addr
	p.prefix = prefix
	p.retry = retry
	p.services = make(map[int32]*grpc.ClientConn, 100)
	if services != nil {
		for _, id := range services {
			p.addService(id)
		}
	}
	return p
}

func (p *service_pool) addService(id int32) *grpc.ClientConn {
	max := uint(1)
	if p.retry {
		max = 3
	}
	// 开启 grpc 中间件的重试功能
	opt := grpc.WithUnaryInterceptor(
		grpc_retry.UnaryClientInterceptor(
			// 重试次数
			grpc_retry.WithMax(max),
			// 重试间隔时间
			grpc_retry.WithBackoff(grpc_retry.BackoffLinear(time.Duration(5)*time.Millisecond)),
			// 超时时间
			grpc_retry.WithPerRetryTimeout(time.Duration(5)*time.Second),
			// 返回码为如下值时重试 (codes.DeadlineExceeded)
			grpc_retry.WithCodes(codes.ResourceExhausted, codes.Unavailable),
		),
	)

	serviceName := p.prefix + strconv.Itoa(int(id))
	target := fmt.Sprintf("consul:///%s/%s", p.addr, serviceName)
	log.Debugf("addService: serviceName:%v, target:%v", serviceName, target)
	conn, err := grpc.Dial(
		target,
		grpc.WithInsecure(),
		opt,
		//grpc.WithBalancerName(roundrobin.Name),
		grpc.WithBalancerName(ketama.Name),
	)
	if err != nil {
		log.Errorf("addService err:%v, serviceName:%v, target:%v", err, serviceName, target)
		return nil
	}
	p.services[id] = conn
	return conn
}

// get a service in round-robin style
// especially useful for load-balance with state-less services
func (p *service_pool) GetService(id int32) (conn *grpc.ClientConn) {
	p.RLock()
	if conn = p.services[id]; conn == nil {
		p.RUnlock()
		p.Lock()
		defer p.Unlock()
		if conn = p.services[id]; conn == nil {
			conn = p.addService(id)
		}
	} else {
		p.RUnlock()
	}
	return
}

func (p *service_pool) GetValue(k string) []byte {
	return ConsulGetValue(p.addr, k)
}

func (p *service_pool) SetValue(k string, v []byte) bool {
	return ConsulSetValue(p.addr, k, v)
}

func ConsulGetValue(addr string, k string) []byte {
	config := api.DefaultConfig()
	config.Address = addr
	if client, err := api.NewClient(config); err == nil {
		kv := client.KV()
		if pair, _, err := kv.Get(k, nil); err == nil && pair != nil {
			return pair.Value
		}
	}
	return nil
}

func ConsulSetValue(addr string, k string, v []byte) bool {
	config := api.DefaultConfig()
	config.Address = addr
	if client, err := api.NewClient(config); err != nil {
		kv := client.KV()
		pair := &api.KVPair{Key: k, Value: v}
		if _, err = kv.Put(pair, nil); err == nil {
			return true
		}
	}
	return false
}

func ConsulRemove(addr string, k string)bool{
	config := api.DefaultConfig()
	config.Address = addr
	if client, err := api.NewClient(config); err != nil {
		kv := client.KV()
		if _, err = kv.Delete(k, nil); err == nil {
			return true
		}
	}
	return false
}