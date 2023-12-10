package server

import (
	"os"
	"path/filepath"

	"gecollsn.it/symparse/server/tools"
	"gopkg.in/yaml.v2"
)

type CServer struct {
	Frontend   ServerCfg `yaml:"frontend"`
	BackendCfg ServerCfg `yaml:"backend"`
}

type ServerCfg struct {
	Port         int `yaml:"port"`
	ReadTimeout  int `yaml:"readTimeout"`
	WriteTimeout int `yaml:"writeTimeout"`
}

var cserver *CServer = &CServer{}

func init() {
	cf, _ := os.ReadFile(filepath.Join(tools.Dirname(), "..", "..", "configs", "cserver.yml"))
	yaml.Unmarshal(cf, cserver)
}

func FrontCfg() ServerCfg {
	return cserver.Frontend
}

func BackendCfg() ServerCfg {
	return cserver.BackendCfg
}
