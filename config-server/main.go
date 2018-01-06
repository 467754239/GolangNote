package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/467754239/cfg-center/src/cfg-server/cfgLoader"
	log "github.com/auxten/logrus"
)

var (
	versionStr = "unknown"
)

type Config struct {
	IsDebug     bool
	ListenPort  int
	CfgDataPath string
}

func init_args(cfg *Config) error {
	flag.BoolVar(&cfg.IsDebug, "d", false, "Debug Swith")
	flag.IntVar(&cfg.ListenPort, "port", 2120, "The Tcp Port to listen")
	flag.StringVar(&cfg.CfgDataPath, "data", "config_data", "Config Data Dir Path")

	version := flag.Bool("v", false, "Show Version")

	flag.Parse()

	if *version {
		fmt.Println(versionStr)
		os.Exit(0)
	}

	if cfg.IsDebug {
		log.SetLevel(log.DebugLevel)
	}

	log.Debug("initing")

	return nil
}

func main() {
	var cfg Config
	err := init_args(&cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	cfg_mgr := cfgLoader.New(cfg.CfgDataPath)
	cfg_mgr.LoadCfg()

}
