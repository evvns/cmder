package commands

import (
	"context"
	"log"
	"cmder/controller/agents"
	pb "cmder/protos/cmds"
)

func (s *AgentManagerServer) Screenshot(ctx context.Context, req *pb.ScreenshotRequest) (*pb.ScreenshotResponse, error) {
	log.Printf("[%d] sending screenshot command", req.AgentId)
	agent, err := agents.GetAgent(req.AgentId)
	if err != nil {
		return nil, err
	}
	resp, err := agent.Screenshot(req)
	if err != nil {
		return nil, err
	}
	log.Printf("received screenshot response")
	return resp, nil
}
