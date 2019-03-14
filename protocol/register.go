package protocol

import (
	"fmt"
	"net"
	"strconv"
	"time"

	consul_api "github.com/hashicorp/consul/api"
)

func RegistConsul(consulAddr string, g *GrpcConfig) error {
	if g.Dcsa <= 0 {
		g.Dcsa = 10
	}
	if g.Interval <= 0 {
		g.Interval = 5
	}
	return register(consulAddr, g)
}

func register(consulAddr string, g *GrpcConfig) error {
	config := consul_api.DefaultConfig()
	config.Address = consulAddr
	client, err := consul_api.NewClient(config)
	if err != nil {
		return err
	}
	agent := client.Agent()

	if len(g.Addr) <= 0 {
		g.Addr = localIP()
	}

	id := g.Addr + ":" + strconv.Itoa(g.Port)
	check := consul_api.AgentServiceCheck{
		Interval: (time.Duration(g.Dcsa) * time.Second).String(),
		GRPC:     fmt.Sprintf("%s:%d/%s", g.Addr, g.Port, g.Name),
		DeregisterCriticalServiceAfter: (time.Duration(g.Interval) * time.Second).String(),
	}
	reg := consul_api.AgentServiceRegistration{
		ID:      id,
		Name:    g.Name,
		Tags:    g.Tags,
		Port:    g.Port,
		Address: g.Addr,
		Meta:    g.Meta,
		Check:   &check,
	}
	if err := agent.ServiceRegister(&reg); err != nil {
		return err
	}
	return nil
}

func localIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
