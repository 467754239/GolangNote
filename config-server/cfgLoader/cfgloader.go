package cfgloader

import (
	"sync"

	log "github.com/auxten/logrus"

	simplejson "github.com/bitly/go-simplejson"
)

type CfgManager struct {
	cfg_file_path string
	cfg_data      *simplejson.Json
	cfg_data_lock sync.RWMutex
}

func New(cfgDataPath string) {
	return &CfgManager{
		cfg_file_path: cfgDataPath,
	}

}

func (this *CfgManager) LoadCfg(js *simplejson.Json, err error) {
	log.Info("Reloading ", this.cfg_file_path)

	return js, nil
}
