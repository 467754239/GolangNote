package cfgTemplate

import (
	"bytes"
	"io/ioutil"
	"text/template"

	log "github.com/auxten/logrus"
	"github.com/bitly/go-simplejson"
	ym "github.com/ghodss/yaml"
)

type CfgTemplate struct {
	vars map[string]interface{}
}

func New() *CfgTemplate {
	return &CfgTemplate{}
}

func (this *CfgTemplate) Translate(templ []byte) (ret []byte) {
	tmpl, err := template.New("tpl").Parse(string(templ))
	if err != nil {
		log.Info("Loading template error, using raw template as config")
		ret = templ
		return ret
	}
	output := new(bytes.Buffer)
	err = tmpl.Execute(output, this.vars)
	if err != nil {
		log.Info("Translating template error, using raw template as config")
		ret = templ
		return ret
	}

	return output.Bytes()

}

/*
	1. 读取yaml文件内容为 var_content
	2. 将文件内容var_content 序列化成JSON
	3. 反序列化JSON
*/
func (this *CfgTemplate) LoadVars(var_path string) (err error) {
	var var_content, jsonBytes []byte
	vars_json := simplejson.New()
	if var_content, err = ioutil.ReadFile(var_path); err == nil {
		if jsonBytes, err = ym.YAMLToJSON(var_content); err == nil {
			if err = vars_json.UnmarshalJSON(jsonBytes); err == nil {
				if this.vars, err = vars_json.Map(); err == nil {
					js_out, err := vars_json.EncodePretty()
					log.Debug(string(js_out), err)
					return nil
				}
			}
		}
	}
	log.Error(err)
	return err

}
