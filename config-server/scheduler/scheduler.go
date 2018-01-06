package scheduler

import (
	"os"

	"github.com/467754239/GolangNote/config-server/cfgLoader"
)

type Scheduler struct {
	reload_ch   chan uint8
	sig_ch      chan os.Signal
	cfgDataPath string
}

func New(cfgDataPath string) *Scheduler {
	return &Scheduler{
		reload_ch:   make(chan uint8, 1),
		sig_ch:      make(chan os.Signal, 1),
		cfgDataPath: cfgDataPath,
	}

}

func (this *Scheduler) setSigHandler() {

}

func (this *Scheduler) setDirWatcher() {

}

func (this *Scheduler) Run(cfgm *cfgLoader.CfgManager) {

}
