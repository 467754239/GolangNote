package cfgLoader

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/467754239/GolangNote/config-server/cfgTemplate"
	"github.com/467754239/GolangNote/config-server/mapFlattener"
	log "github.com/auxten/logrus"
	"github.com/coreos/etcd/pkg/fileutil"

	simplejson "github.com/bitly/go-simplejson"
	ym "github.com/ghodss/yaml"
)

type cfgData map[string]interface{}

const (
	CONF_EXT   = ".yaml"
	GLOBAL_VAR = "GLOBAL_VAR.yaml"
)

type CfgManager struct {
	cfg_file_path string
	cfg_data      *simplejson.Json
	cfg_data_lock sync.RWMutex
}

func New(cfgDataPath string) *CfgManager {
	return &CfgManager{
		cfg_file_path: cfgDataPath,
	}
}

type Fmt int

const (
	Raw Fmt = 1 + iota
	Flat
)

const NIL_JSON string = "null"

func (cfgm *CfgManager) GetCfg_json(conf_path_list []string, format Fmt) (js_ret []byte) {
	log.Debug("GetCfg_json ", conf_path_list)

	cfgm.cfg_data_lock.RLock()
	defer cfgm.cfg_data_lock.RUnlock()

	jin := cfgm.cfg_data
	for _, p := range conf_path_list {
		jin = jin.Get(p)
	}

	if format == Flat {
		m, err := jin.Map()
		if err != nil {
			log.Error(err)
			return []byte(NIL_JSON)
		}
		flat_m := mapFlattener.Flatten(m)
		js_ret, _ = json.Marshal(flat_m)
	} else if format == Raw {
		js_ret, _ = jin.MarshalJSON()
	}
	return js_ret
}

func (this *CfgManager) LoadCfg() (js *simplejson.Json, err error) {
	log.Info("Reloading ", this.cfg_file_path)

	js = simplejson.New()
	g_var_path := this.cfg_file_path + "/" + GLOBAL_VAR

	cfg_t := cfgTemplate.New()
	if fileutil.Exist(g_var_path) {
		if err = cfg_t.LoadVars(g_var_path); err != nil {
			log.Info("No global var yaml loaded")
		}
	}

	var scan_func = func(path_one string, info os.FileInfo, e error) (err error) {

		if e == nil && info.Mode().IsRegular() && strings.ToLower(filepath.Ext(info.Name())) == CONF_EXT {

			var (
				cfg_file_content, jsonBytes []byte
				rel_path                    string
			)

			if cfg_file_content, err = ioutil.ReadFile(path_one); err == nil {
				if rel_path, err = filepath.Rel(this.cfg_file_path, path_one); err == nil {
					rel_path_without_ext := rel_path[:len(rel_path)-len(CONF_EXT)]
					path_list := strings.Split(rel_path_without_ext, string(os.PathSeparator))

					// Init template
					inited_content := cfg_t.Translate(cfg_file_content)
					log.Debug(string(inited_content))

					// YAML objects are not completely compatible with JSON objects (e.g. you
					// can have non-string keys in YAML). So, convert the YAML-compatible object
					// to a JSON-compatible object, failing with an error if irrecoverable
					// incompatibilties happen along the way.
					if jsonBytes, err = ym.YAMLToJSON(inited_content); err == nil {
						jsFromCfg := simplejson.New()

						if err = jsFromCfg.UnmarshalJSON(jsonBytes); err == nil {
							js.SetPath(path_list, jsFromCfg.Interface())

							jsout, err := js.EncodePretty()
							log.Debug(string(jsout), err)

							return nil
						}
					}
				}
				log.Error(fmt.Sprintf("Error parsing %s, Skipped", path_one))
				log.Error(err.Error())
			}
		}
		return nil // 永远返回nil,这样就不会因为解析一个文件出错导致不正常运行
	}

	//todo if cfgm.cfg_file_path is a file
	err = filepath.Walk(this.cfg_file_path, scan_func)
	if err != nil {
		log.Error(err.Error())
		return js, err
	}

	jsout, err := js.EncodePretty()
	log.Debug(string(jsout))

	//todo load cfg_data
	this.cfg_data_lock.Lock()
	this.cfg_data = js
	this.cfg_data_lock.Unlock()

	return js, err
}
