package cmder

import (
	_ "embed"
	"encoding/json"
	"log"
)

type CertificateConfig struct {
	Cert string `json:"cert"`
	Key  string `json:"key"`
}

type PersistenceConfig struct {
	Path     string `json:"path"`
	TaskName string `json:"taskName"`
}

type AgentConfig struct {
	ControllerAddress string            `json:"controllerAddress"`
	Persistence       PersistenceConfig `json:"persistence"`
	Cert              CertificateConfig `json:"cert"`
}

type ControllerConfig struct {
	AgentServerPort        int               `json:"agentServerPort"`
	AgentManagerServerPort int               `json:"agentManagerServerPort"`
	Cert                   CertificateConfig `json:"cert"`
}

type Config struct {
	Agent      AgentConfig      `json:"agent"`
	Controller ControllerConfig `json:"controller"`
}

var (
	//go:embed config.json
	configBuffer []byte

	Conf = getConfig()
)

func getConfig() *Config {
	cfg := &Config{}
	err := json.Unmarshal(configBuffer, cfg)
	if err != nil {
		log.Fatalf("failed to parse config: %v", err)
	}
	return cfg
}
