package commands

import (
	"context"
	"log"
	"cmder/controller/agents"
	pb "cmder/protos/cmds"
)

func (s *AgentManagerServer) RunShellCommand(ctx context.Context, req *pb.ShellCommandRequest) (*pb.ShellCommandResponse, error) {
	log.Printf("[%d] sending shell command", req.AgentId)
	agent, err := agents.GetAgent(req.AgentId)
	if err != nil {
		return nil, err
	}
	resp, err := agent.RunShellCommand(req)
	if err != nil {
		return nil, err
	}
	log.Printf("received shell response")
	return resp, nil
}
