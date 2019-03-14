package main

import (
	"flag"
	"net"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"grpcserver/internal"
	"local.com/abc/game/protocol"
	"local.com/abc/game/util"
)

func main() {
	var name string
	flag.StringVar(&name, "conf", "app.yaml", "config file name")
	flag.Parse()
	defer util.PrintPanicStack()
	// open profiling
	config := internal.InitConfig(name)

	lis, err := net.Listen("tcp", config.Grpc.Listen)
	if err != nil {
		panic(err)
	}

	gs := grpc.NewServer()
	s := &internal.Server{}
	s.Init(config)
	protocol.RegisterGameServer(gs, s)
	protocol.RegisterGrpcServer(gs)

	err = protocol.RegistConsul(config.Consul.Addr, &config.Grpc)
	if err != nil {
		panic(err)
	}
	log.Info("starting service at:", lis.Addr())
	signal := util.NewAppSignal()
	signal.Run(func() {
		go gs.Serve(lis)
	})
}
