package main

import (
	"github.com/467754239/GolangNote/config-server/cfgLoader"
	"github.com/467754239/GolangNote/config-server/httpApi"
	"github.com/467754239/GolangNote/config-server/scheduler"
)

func main() {
	CfgDataPath := "config_data"

	cfg_mgr := cfgLoader.New(CfgDataPath)
	cfg_mgr.LoadCfg()

	scheduler := scheduler.New(CfgDataPath)
	go scheduler.Run(cfg_mgr)

	httpApi.HTTPServerStart(2120, cfg_mgr)

}
