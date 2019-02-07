package lib

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	agent2 "softnet/pkg/agent"
	"softnet/pkg/agent/proto"
)

func StartSandboxServer(db *gorm.DB) {
	agent := agent2.NewSandboxCounterAgent(db)
	s := grpc.NewServer()
	proto.RegisterSandboxAgentServer(s, agent)
	reflection.Register(s)
	lis, err := net.Listen("tcp",
		fmt.Sprintf(":%s", "4000"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("GRPC Running on :%s\n", "4000")
	s.Serve(lis)
}
