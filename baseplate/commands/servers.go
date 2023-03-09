package commands

import (
	pb "cmder/protos/cmds"
)

type AgentManagerServer struct {
	pb.UnimplementedAgentManagerServer
}
